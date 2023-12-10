package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%.3f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3f°F", f) }

func (ft Feet) String() string { return fmt.Sprintf("%.3fft", ft) }
func (m Meter) String() string { return fmt.Sprintf("%.3fm", m) }

func (lb Pound) String() string    { return fmt.Sprintf("%.3flb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%.3fkg", kg) }

func CTof(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToc(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FtTom(ft Feet) Meter { return Meter(ft / 3.28084) }
func MToft(m Meter) Feet  { return Feet(m * 3.28084) }

func LbTokg(lb Pound) Kilogram { return Kilogram(lb / 2.2046) }
func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2046) }
