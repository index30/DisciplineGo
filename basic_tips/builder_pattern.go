package main

import "fmt"

type Degree float64

const (
	Celsius    Degree = 0  // 摂氏
	Fahrenheit        = 32 // 華氏
)

type Weather string

const (
	SunnyWeather Weather = "sunny"
	CloudWeather         = "cloud"
	RainyWeather         = "rainy"
)

type Sense string

const (
	HotSense  Sense = "hot"
	ColdSense       = "cold"
)

type Builder interface {
	Weather(Weather) Builder
	Sense(Sense) Builder
	FahrenheitDegree(Degree) Builder
	Build() Interface
}

type Interface interface {
	Observe() error
	Stop() error
}

type forecastBuilder struct {
	weather Weather
	sense   Sense
	degree  Degree
}

type forecast struct {
	params forecastBuilder
}

func NewBuilder() *forecastBuilder {
	return &forecastBuilder{
		weather: RainyWeather,
		sense:   HotSense,
		degree:  Celsius,
	}
}

func (b *forecastBuilder) Weather(weather Weather) Builder {
	b.weather = weather
	return b
}

func (b *forecastBuilder) Sense(sense Sense) Builder {
	b.sense = sense
	return b
}

func (b *forecastBuilder) FahrenheitDegree(degree Degree) Builder {
	b.degree = degree
	return b
}

func (b *forecastBuilder) Build() Interface {
	return &forecast{
		params: *b,
	}
}

func (f *forecast) Observe() error {
	fmt.Printf("Observe: %#+v\n", f.params)
	return nil
}

func (f *forecast) Stop() error {
	fmt.Printf("Stop: %#+v\n", f.params)
	return nil
}

func main() {
	forecast := NewBuilder().Weather(CloudWeather).Sense(HotSense).FahrenheitDegree(Fahrenheit).Build()
	forecast.Observe()
	forecast.Stop()
}
