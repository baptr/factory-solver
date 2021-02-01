# Usage

`go run github.com/baptr/factory-solver <game config proto path> "<desired resource>" <rate per minute>`

# Customization

For now, it uses the first recipe in the provided config textproto that
produces a required resource. If you don't have access to a given rare recipe
or resource, you can move it further down the list, or delete it entirely from
the textproto.

Similarly, resource harvesting efficiency values are hardcoded in the textproto
for now, so modify to match your game.

This will be improved in the future.

# Examples

## 30x Gravity matrix/min

Using a bunch of rare materials, how do I run 6 green science research stations?

```
$ go run github.com/baptr/factory-solver games/dyson-sphere-program.textproto "Gravity matrix" 30
Solution:
*  6.00 x Gravity matrix RESEARCHED (30 Gravity matrix/min)
\_*  1.50 x Graviton lens ASSEMBLED (15 Graviton lens/min)
| \_*  2.00 x Diamond SMELTED (60 Diamond/min)
| | \_*  4.00 x X-ray cracking REFINED (60 Energetic graphite/min)
| |   \_*  2.00 x Plasma refining REFINED (60 Refined oil/min)
| |   | \_*  0.67 x Crude oil seep OIL_HARVESTED (60 Crude oil/min)
| \_*  2.00 x Strange matter COLLIDED (15 Strange matter/min)
|   \_*  2.00 x Particle container ASSEMBLED (30 Particle container/min)
|   | \_*  2.00 x Electromagnetic turbine ASSEMBLED (60 Electromagnetic turbine/min)
|   | | \_*  4.00 x Electric motor ASSEMBLED (120 Electric motor/min)
|   | | | \_*  4.00 x Iron ingot SMELTED (240 Iron ingot/min)
|   | | | | \_*  5.33 x Iron vein VEIN_HARVESTED (240 Iron ore/min)
|   | | | \_*  2.00 x Gear ASSEMBLED (120 Gear/min)
|   | | | | \_*  2.00 x Iron ingot SMELTED (120 Iron ingot/min)
|   | | | |   \_*  2.67 x Iron vein VEIN_HARVESTED (120 Iron ore/min)
|   | | | \_*  1.00 x Magnetic coil ASSEMBLED (120 Magnetic coil/min)
|   | | |   \_*  3.00 x Magnet SMELTED (120 Magnet/min)
|   | | |   | \_*  2.67 x Iron vein VEIN_HARVESTED (120 Iron ore/min)
|   | | |   \_*  1.00 x Copper ingot SMELTED (60 Copper ingot/min)
|   | | |     \_*  1.33 x Copper vein VEIN_HARVESTED (60 Copper ore/min)
|   | | \_*  1.00 x Magnetic coil ASSEMBLED (120 Magnetic coil/min)
|   | |   \_*  3.00 x Magnet SMELTED (120 Magnet/min)
|   | |   | \_*  2.67 x Iron vein VEIN_HARVESTED (120 Iron ore/min)
|   | |   \_*  1.00 x Copper ingot SMELTED (60 Copper ingot/min)
|   | |     \_*  1.33 x Copper vein VEIN_HARVESTED (60 Copper ore/min)
|   | \_*  1.00 x Copper ingot SMELTED (60 Copper ingot/min)
|   | | \_*  1.33 x Copper vein VEIN_HARVESTED (60 Copper ore/min)
|   | \_*  1.00 x Fire Ice graphene CHEMICAL (60 Graphene/min)
|   |   \_*  1.33 x Fire ice vein VEIN_HARVESTED (60 Fire ice/min)
|   \_*  0.50 x Iron ingot SMELTED (30 Iron ingot/min)
|   | \_*  0.67 x Iron vein VEIN_HARVESTED (30 Iron ore/min)
|   \_*  2.50 x Deuterium COLLIDED (150 Deuterium/min)
|     \_*  0.25 x Hydrogen gas giant ORBIT_HARVESTED (180 Hydrogen/min)
\_*  1.50 x Quantum chip ASSEMBLED (15 Quantum chip/min)
  \_*  1.50 x Processor ASSEMBLED (30 Processor/min)
  | \_*  0.50 x Circuit board ASSEMBLED (60 Circuit board/min)
  | | \_*  1.00 x Iron ingot SMELTED (60 Iron ingot/min)
  | | | \_*  1.33 x Iron vein VEIN_HARVESTED (60 Iron ore/min)
  | | \_*  0.50 x Copper ingot SMELTED (30 Copper ingot/min)
  | |   \_*  0.67 x Copper vein VEIN_HARVESTED (30 Copper ore/min)
  | \_*  2.00 x Microcrystalline component ASSEMBLED (60 Microcrystalline component/min)
  |   \_*  4.00 x High-purity silicon SMELTED (120 High-purity silicon/min)
  |   | \_*  5.33 x Silicon vein VEIN_HARVESTED (240 Silicon ore/min)
  |   \_*  1.00 x Copper ingot SMELTED (60 Copper ingot/min)
  |     \_*  1.33 x Copper vein VEIN_HARVESTED (60 Copper ore/min)
  \_*  6.00 x Plane filter ASSEMBLED (30 Plane filter/min)
    \_*  2.00 x Casimir crystal ASSEMBLED (30 Casimir crystal/min)
    | \_*  2.00 x Titanium crystal ASSEMBLED (30 Titanium crystal/min)
    | | \_*  0.67 x Organic crystal vein VEIN_HARVESTED (30 Organic crystal/min)
    | | \_*  3.00 x Titanium ingot SMELTED (90 Titanium ingot/min)
    | |   \_*  4.00 x Titanium vein VEIN_HARVESTED (180 Titanium ore/min)
    | \_*  1.00 x Fire Ice graphene CHEMICAL (60 Graphene/min)
    | | \_*  1.33 x Fire ice vein VEIN_HARVESTED (60 Fire ice/min)
    | \_*  0.50 x Hydrogen gas giant ORBIT_HARVESTED (360 Hydrogen/min)
    \_*  2.50 x Titanium glass ASSEMBLED (60 Titanium glass/min)
      \_*  2.00 x Glass SMELTED (60 Glass/min)
      | \_*  2.67 x Stone vein VEIN_HARVESTED (120 Stone ore/min)
      \_*  2.00 x Titanium ingot SMELTED (60 Titanium ingot/min)
      | \_*  2.67 x Titanium vein VEIN_HARVESTED (120 Titanium ore/min)
      \_*  0.80 x Water pump LIQUID_HARVESTED (60 Water/min)

Total production:
  15.33 x Iron vein
   7.50 x Iron ingot
   6.67 x Titanium vein
   6.00 x Magnet
   6.00 x Plane filter
   6.00 x Copper vein
   6.00 x Gravity matrix
   5.33 x Silicon vein
   5.00 x Titanium ingot
   4.50 x Copper ingot
   4.00 x High-purity silicon
   4.00 x X-ray cracking
   4.00 x Electric motor
   2.67 x Stone vein
   2.67 x Fire ice vein
   2.50 x Titanium glass
   2.50 x Deuterium
   2.00 x Plasma refining
   2.00 x Strange matter
   2.00 x Diamond
   2.00 x Microcrystalline component
   2.00 x Gear
   2.00 x Casimir crystal
   2.00 x Fire Ice graphene
   2.00 x Electromagnetic turbine
   2.00 x Particle container
   2.00 x Magnetic coil
   2.00 x Titanium crystal
   2.00 x Glass
   1.50 x Graviton lens
   1.50 x Quantum chip
   1.50 x Processor
   0.80 x Water pump
   0.75 x Hydrogen gas giant
   0.67 x Organic crystal vein
   0.67 x Crude oil seep
   0.50 x Circuit board

Final resources/min:
  30.00 x Gravity matrix
 -30.00 x Organic crystal
 -60.00 x Crude oil
 -60.00 x Water
-120.00 x Fire ice
-120.00 x Stone ore
-240.00 x Silicon ore
-270.00 x Copper ore
-300.00 x Titanium ore
-510.00 x Hydrogen
-690.00 x Iron ore

Building Power: -126.8MW
```

## Oil power calculation

How many burners do I need to use up the output of 10 Plasma refining, assuming 100% power load?

```
$ go run github.com/baptr/factory-solver games/dyson-sphere-program.textproto "Refined oil" 300
Solution:
* 10.00 x Plasma refining REFINED (300 Refined oil/min)
\_*  3.33 x Crude oil seep OIL_HARVESTED (300 Crude oil/min)

Total production:
  10.00 x Plasma refining
   3.33 x Crude oil seep

Final resources/min:
 300.00 x Refined oil
 150.00 x Hydrogen
-300.00 x Crude oil

Building Power: -12.4MW
Burn(Refined oil x 300.00/min) = 22.0MW = 10.19 x Thermal power station = 17.6MW @ 80%
Burn(Hydrogen x 150.00/min) = 20.0MW = 9.26 x Thermal power station = 16.0MW @ 80%
```
