package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strconv"

	"github.com/baptr/factory-solver/configpb"
	"github.com/golang/protobuf/proto"
)

type Solver struct {
	byResult map[string][]*configpb.Recipe

	resPerMin map[string]float64
}

func newSolver(cfg *configpb.Config) (*Solver, error) {
	byName := make(map[string]*configpb.Recipe)
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

	byResult := make(map[string][]*configpb.Recipe)
	for _, r := range byName {
		for _, out := range r.Result {
			byResult[out.Item] = append(byResult[out.Item], r)
		}
	}
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

	return &Solver{
		byResult: byResult,

		resPerMin: make(map[string]float64),
	}, nil
}

func (s *Solver) step(res string, rate float64, depth int) {
	rs := s.byResult[res]
	if len(rs) == 0 {
		return
	}
	r := rs[0]

	buildings := rate / outPerMin(r, res)
	fmt.Printf("%*s%6.2f x %s %v (%.f %s/min)\n", depth*2, "", buildings, r.Name, r.Type, rate, res)
	for _, i := range r.Result {
		s.resPerMin[i.Item] += buildings * outPerMin(r, i.Item)
	}
	for _, i := range r.Input {
		need := buildings * inPerMin(r, i.Item)
		s.resPerMin[i.Item] -= need
		s.step(i.Item, need, depth+1)
	}
	return
}

func outPerMin(r *configpb.Recipe, res string) float64 {
	for _, i := range r.Result {
		if i.Item == res {
			return perMin(r) * float64(i.Quantity)
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
	s.step(target, rate, 0)
	fmt.Println()

	fmt.Printf("Final resources/min:\n")
	var resNames []string
	for r, c := range s.resPerMin {
		if c == 0 {
			continue
		}
		resNames = append(resNames, r)
	}
	sort.Slice(resNames, func(i, j int) bool {
		return s.resPerMin[resNames[i]] < s.resPerMin[resNames[j]]
	})
	for _, r := range resNames {
		fmt.Printf("%7.2f x %s\n", s.resPerMin[r], r)
	}
}
