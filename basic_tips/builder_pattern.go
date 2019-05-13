package main

import "fmt"

type Degree float64

// 定数の定義
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
	FahrenheitDegree(Degree) Builder
	Decision(Decision) Builder
	Build() Interface
}

type Interface interface {
	Observe() error
	Stop() error
}

type forecastBuilder struct {
	weather  Weather
	sense    Sense
	degree   Degree
	decision Decision
}

type forecast struct {
	params forecastBuilder
}

//デフォルト値の指定
func NewBuilder() *forecastBuilder {
	return &forecastBuilder{
		weather:  RainyWeather,
		sense:    HotSense,
		degree:   Celsius,
		decision: GoOut,
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

func (b *forecastBuilder) Decision(decision Decision) Builder {
	b.decision = decision
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
	forecast := NewBuilder().Weather(CloudWeather).Sense(HotSense).FahrenheitDegree(Fahrenheit).Decision(GoOut).Build()
	forecast.Observe()
	forecast.Stop()
}
