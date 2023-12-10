package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

const (
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

const (
	AbsoluteZeroK = 0
	FreezingK     = 273.15
	BoilingK      = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func CTof(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CTok(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func FToc(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FTok(f Fahrenheit) Kelvin  { return Kelvin((f-32)*5/9 + 273.15) }
func KToc(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func KTof(k Kelvin) Fahrenheit  { return Fahrenheit((k-273.15)*9/5 + 32) }
