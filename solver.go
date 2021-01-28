package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"strconv"

	"github.com/baptr/factory-solver/configpb"
	"github.com/golang/protobuf/proto"
)

type Solver struct {
	byResult  map[string][]*configpb.Recipe
	buildings map[configpb.ProductionType][]*configpb.Building
	bonuses   map[configpb.ProductionType]float64
	fuels     map[string]*configpb.Fuel

	prodPerMin map[string]float64
	resPerMin  map[string]float64
	powerW     float64
}

func newSolver(cfg *configpb.Config) (*Solver, error) {
	byName := make(map[string]*configpb.Recipe)
	byResult := make(map[string][]*configpb.Recipe)
	var unnamed []*configpb.Recipe
	for _, r := range cfg.Recipe {
		if r.Name == "" {
			if len(r.Result) != 1 {
				return nil, fmt.Errorf("unnamed recipe with complicated results: %v", r)
			}
			unnamed = append(unnamed, r)
		} else {
			if old, ok := byName[r.Name]; ok {
				return nil, fmt.Errorf("duplicate recipe name %q.\nOld: %v\nNew: %v", r.Name, old, r)
			}
			byName[r.Name] = r
		}
		for _, out := range r.Result {
			byResult[out.Item] = append(byResult[out.Item], r)
		}
	}
	for _, r := range unnamed {
		name := r.Result[0].Item
		if old, ok := byName[name]; ok {
			return nil, fmt.Errorf("unnamed recipe overlaps existing recipe %q.\nOld: %v\nNew: %v", name, old, r)
		}
		r.Name = name
		byName[name] = r
	}
	log.Printf("Loaded %d recipes", len(byName))

	for res, recipes := range byResult {
		if len(recipes) != 1 {
			log.Printf("WARNING: Multiple recipes produce %q: %v", res, recipes)
		}
	}
	unsourced := make(map[string][]*configpb.Recipe)
	for _, r := range byName {
		for _, i := range r.Input {
			if len(byResult[i.Item]) == 0 {
				unsourced[i.Item] = append(unsourced[i.Item], r)
			}
		}
	}
	for i, recipes := range unsourced {
		var names []string
		for _, r := range recipes {
			names = append(names, r.Name)
		}
		log.Printf("Unsourced component %q in recipes: %q", i, names)
	}

	bonuses := make(map[configpb.ProductionType]float64)
	for _, b := range cfg.Efficiency {
		for _, t := range b.Type {
			if bonuses[t] != 0 {
				log.Printf("WARNING: Overwriting existing %q bonus %f with %f", t, bonuses[t], b.Multiplier)
			}
			bonuses[t] = b.Multiplier
		}
	}

	buildings := make(map[configpb.ProductionType][]*configpb.Building)
	for _, b := range cfg.Building {
		buildings[b.Type] = append(buildings[b.Type], b)
	}
	log.Printf("Loaded %d buildings", len(buildings))

	fuels := make(map[string]*configpb.Fuel)
	for _, f := range cfg.Fuel {
		if old := fuels[f.Item]; old != nil {
			log.Printf("WARNING: Duplicate fuel definition for %q:\nOLD: %v\nNEW:  %v", f.Item, old, f)
		}
		fuels[f.Item] = f
	}

	return &Solver{
		byResult:  byResult,
		buildings: buildings,
		bonuses:   bonuses,
		fuels:     fuels,

		resPerMin:  make(map[string]float64),
		prodPerMin: make(map[string]float64),
	}, nil
}

func (s *Solver) step(res string, rate float64, depth int) error {
	s.resPerMin[res] -= rate
	if s.resPerMin[res] > -0.01 {
		// Already enough present.
		return nil
	}

	rs := s.byResult[res]
	if len(rs) == 0 {
		return fmt.Errorf("no recipe produces %q", res)
	}
	r := rs[0]

	var marker string
	if depth > 0 {
		marker = fmt.Sprintf("%*s", (depth-1)*2, "\\_")
	}

	util := rate / s.outPerMin(r, res)
	s.prodPerMin[r.Name] += util
	b, err := s.buildingFor(r)
	if err != nil {
		log.Printf("WARNING: Power calculation will be inaccurate: %v", err)
	} else {
		buildings := util / b.Efficiency
		power := float64(b.ActiveWattsUsed)*buildings + float64(b.IdleWattsUsed)*(math.Ceil(buildings)-buildings)
		s.powerW -= power
	}
	fmt.Printf("%*s*%6.2f x %s %v (%.f %s/min)\n", depth*2, marker, util, r.Name, r.Type, rate, res)

	switch r.Type {
	case configpb.ProductionType_VEIN_HARVESTED,
		configpb.ProductionType_LIQUID_HARVESTED,
		configpb.ProductionType_OIL_HARVESTED:
		// Leave harvest in the negative so they show up as I/O
	default:
		for _, i := range r.Result {
			s.resPerMin[i.Item] += util * s.outPerMin(r, i.Item)
		}
	}
	for _, i := range r.Input {
		need := util * inPerMin(r, i.Item)
		if err := s.step(i.Item, need, depth+1); err != nil {
			return fmt.Errorf("producing %s: %v", res, err)
		}
	}
	return nil
}

func (s *Solver) burn(f *configpb.Fuel, perMin float64) {
	bOpts := s.buildings[f.Type]
	if len(bOpts) == 0 {
		log.Printf("WARNING: No buildings defined to burn %v fuel (%q)", f.Type, f.Item)
		return
	}
	b := bOpts[0]
	if len(bOpts) > 1 {
		log.Printf("WARNING: Defaulting to first candidate burner %s of %v", b.Name, bOpts)
	}
	totalWatts := float64(f.Joules) * perMin / 60
	util := totalWatts / float64(-b.ActiveWattsUsed)
	fmt.Printf("Burn(%s x %.2f) = %sW = %.2f x %s\n", f.Item, perMin, si(totalWatts), util, b.Name)
}

func si(f float64) string {
	if math.Abs(f) >= 1e9 {
		return fmt.Sprintf("%.1fG", f/1e9)
	}
	if math.Abs(f) >= 1e6 {
		return fmt.Sprintf("%.1fM", f/1e6)
	}
	if math.Abs(f) >= 1e3 {
		return fmt.Sprintf("%.1fk", f/1e3)
	}
	return fmt.Sprintf("%.f", f)
}

func (s *Solver) buildingFor(r *configpb.Recipe) (*configpb.Building, error) {
	opts := s.buildings[r.Type]
	if len(opts) == 0 {
		return nil, fmt.Errorf("no building defined for type %v needed by recipe %q", r.Type, r.Name)
	}
	// TODO: Bin pack the desired rate?
	// TODO: Add options for space/power optimization?
	return opts[0], nil
}

func (s *Solver) outPerMin(r *configpb.Recipe, res string) float64 {
	bonus := s.bonuses[r.Type]
	if bonus == 0 {
		bonus = 1
	}
	for _, i := range r.Result {
		if i.Item == res {
			return perMin(r) * float64(i.Quantity) * bonus
		}
	}
	log.Fatalf("outPerMin: unable to find result %q in recipe %v", res, r)
	return 0
}

func inPerMin(r *configpb.Recipe, res string) float64 {
	for _, i := range r.Input {
		if i.Item == res {
			return perMin(r) * float64(i.Quantity)
		}
	}
	log.Fatalf("inPerMin: unable to find input %q in recipe %v", res, r)
	return 0
}

func perMin(r *configpb.Recipe) float64 {
	if r.GetTiming() == nil {
		log.Fatalf("Recipe %q has no timing information: %v", r.Name, r)
	}
	if d := r.GetDuration(); d != nil {
		return 60.0 / secs(d)
	}
	if s := r.GetPerSecond(); s != 0 {
		return 60.0 * s
	}
	if m := r.GetPerMinute(); m != 0 {
		return m
	}
	log.Fatalf("Unhandled timing type in recipe %q: %v", r.Name, r)
	return 0
}

func secs(d *configpb.Duration) float64 {
	return float64(d.Seconds) + float64(d.Millis)/1000
}

// TODOs:
// - Produce: Information matrix x3
// - Produce: as many Information matrix as possible
//		- Given: 4.2/s crude oil
func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <game textproto> <target resource> <per minute>", path.Base(os.Args[0]))
	}
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open config textproto: %v", err)
	}
	cfg := &configpb.Config{}
	if err := proto.UnmarshalText(string(f), cfg); err != nil {
		log.Fatalf("Failed to unmarshal config proto: %v", err)
	}

	target := os.Args[2]
	rate, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("Invalid target rate: %v", err)
	}

	s, err := newSolver(cfg)
	if err != nil {
		log.Fatalf("Config validation failed: %v", err)
	}

	fmt.Printf("Solution:\n")
	if err := s.step(target, rate, 0); err != nil {
		log.Fatalf("Failed to find a solution: %v", err)
	}
	s.resPerMin[target] += rate
	fmt.Println()

	fmt.Printf("Total production:\n")
	var prodNames []string
	for r := range s.prodPerMin {
		prodNames = append(prodNames, r)
	}
	// TODO: Topological sort would feel better.
	sort.Slice(prodNames, func(i, j int) bool {
		return s.prodPerMin[prodNames[i]] > s.prodPerMin[prodNames[j]]
	})
	for _, r := range prodNames {
		fmt.Printf("%7.2f x %s\n", s.prodPerMin[r], r)
	}
	fmt.Println()

	fmt.Printf("Final resources/min:\n")
	var resNames []string
	for r := range s.resPerMin {
		resNames = append(resNames, r)
	}
	sort.Slice(resNames, func(i, j int) bool {
		return s.resPerMin[resNames[i]] > s.resPerMin[resNames[j]]
	})
	for _, r := range resNames {
		c := s.resPerMin[r]
		if c == 0 {
			continue
		}
		fmt.Printf("%7.2f x %s\n", c, r)
	}

	fmt.Printf("\nBuilding Power: %sW\n", si(s.powerW))
	for r, v := range s.resPerMin {
		if v > 0 {
			f := s.fuels[r]
			if f == nil {
				continue // not a fuel
			}
			// Propose burning the excess
			s.burn(f, v)
		}
	}
}
