package main

import "fmt"

type Degree float64
type Sense string
type Weather string

const (
	Observed      Degree  = 0
	HotSense      Sense   = "hot"
	ColdSense             = "cold"
	SunnyWeather  Weather = "sunny"
	CloudyWeather         = "cloudy"
	RainyWeather          = "rainy"
)

type Builder interface {
	Degree(Degree) Builder
	Sense(Sense) Builder
	Weather(Weather) Builder
	Build() Interface
}

type Interface interface {
	Observe() error
	Stop() error
}

type CelsiusBuilder struct {
	degree  Degree
	sense   Sense
	weather Weather
}

type celsius struct {
	params CelsiusBuilder
}

func NewCelsiusBuilder() *CelsiusBuilder {
	return &CelsiusBuilder{
		degree:  Observed,
		sense:   HotSense,
		weather: SunnyWeather,
	}
}

func (c *CelsiusBuilder) Degree(degree Degree) Builder {
	c.degree = degree
	return c
}

func (c *CelsiusBuilder) Sense(sense Sense) Builder {
	c.sense = sense
	return c
}

func (c *CelsiusBuilder) Weather(weather Weather) Builder {
	c.weather = weather
	return c
}

func (c *CelsiusBuilder) Build() Interface {
	return &celsius{
		params: *c,
	}
}

func (c *celsius) Observe() error {
	fmt.Printf("Observe: %#+v\n", c.params)
	return nil
}

func (c *celsius) Stop() error {
	fmt.Printf("Stop: %#+v\n", c.params)
	return nil
}

type FahrenheitBuilder struct {
	degree  Degree
	sense   Sense
	weather Weather
}

type fahrenheit struct {
	params FahrenheitBuilder
}

func NewFahrenheitBuilder() *FahrenheitBuilder {
	return &FahrenheitBuilder{
		degree:  Observed,
		sense:   HotSense,
		weather: SunnyWeather,
	}
}

func (f *FahrenheitBuilder) Degree(degree Degree) Builder {
	f.degree = degree
	return f
}

func (f *FahrenheitBuilder) Sense(sense Sense) Builder {
	f.sense = sense
	return f
}

func (f *FahrenheitBuilder) Weather(weather Weather) Builder {
	f.weather = weather
	return f
}

func (f *FahrenheitBuilder) Build() Interface {
	return &fahrenheit{
		params: *f,
	}
}

func (f *fahrenheit) Observe() error {
	fmt.Printf("Observe: %#+v\n", f.params)
	return nil
}

func (f *fahrenheit) Stop() error {
	fmt.Printf("Stop: %#+v\n", f.params)
	return nil
}

func main() {
	celsius := NewCelsiusBuilder().Weather(SunnyWeather).Sense(HotSense).Degree(25).Build()
	fahrenheit := NewFahrenheitBuilder().Weather(CloudyWeather).Sense(ColdSense).Degree(10).Build()
	celsius.Observe()
	fahrenheit.Observe()
	celsius.Stop()
	fahrenheit.Stop()
}
