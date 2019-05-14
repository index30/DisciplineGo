package main

import "fmt"

type Degreeval string

// 定数の定義
const (
	Celsius    Degreeval = "celsius"    // 摂氏
	Fahrenheit           = "fahrenheit" // 華氏
)

type Degree float64

const (
	Observed Degree = 0
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

//追加要素
/*
Builderに束縛されているので, 抜け漏れがある場合ビルド時に分かる
*/
type Decision string

const (
	GoOut Decision = "fishing"
	Home           = "game"
)

// メソッド宣言
type Builder interface {
	Weather(Weather) Builder
	Sense(Sense) Builder
	CelsiusDegree(Degree) Builder
	Degreeval(Degreeval) Builder
	Decision(Decision) Builder
	Build() Interface
}

type Interface interface {
	Observe() error
	Stop() error
}

type forecastBuilder struct {
	weather   Weather
	sense     Sense
	degreeval Degreeval
	degree    Degree
	decision  Decision
}

type forecast struct {
	params forecastBuilder
}

//デフォルト値の指定
func NewBuilder() *forecastBuilder {
	return &forecastBuilder{
		weather:   RainyWeather,
		sense:     HotSense,
		degreeval: Celsius,
		degree:    Observed,
		decision:  GoOut,
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

func (b *forecastBuilder) CelsiusDegree(degree Degree) Builder {
	b.degree = degree
	return b
}

func (b *forecastBuilder) Degreeval(degreeval Degreeval) Builder {
	b.degreeval = degreeval
	return b
}

func (b *forecastBuilder) Decision(decision Decision) Builder {
	b.decision = decision
	return b
}

func (b *forecastBuilder) Build() Interface {
	if b.degreeval == "fahrenheit" {
		b.degree = b.degree - 32
	}
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
	forecast := NewBuilder().Weather(CloudWeather).Sense(HotSense).CelsiusDegree(52).Degreeval(Fahrenheit).Decision(GoOut).Build()
	forecast.Observe()
	forecast.Stop()
}
