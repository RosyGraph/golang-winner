package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Feet float64
type Meters float64

type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroK = 0
	FreezingK     = 273.15
	BoilingK      = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }

func (f Feet) String() string   { return fmt.Sprintf("%.2f\"", f) }
func (m Meters) String() string { return fmt.Sprintf("%.2fm", m) }

func (p Pounds) String() string    { return fmt.Sprintf("%.2flbs", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%.2fkm", k) }
