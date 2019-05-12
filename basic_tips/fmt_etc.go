package main

import "fmt"

type Value struct {
    Old string
    New	string
}

func (v *Value) String() string{
	return fmt.Sprintf("Old is [%s]\nNew is [%s]", v.Old, v.New)
}

type GuideLine struct {
    *Value
    Accept bool
}

func main(){
	val := &Value{
        Old: "effort",
        New:  "action",
	}
	gl := &GuideLine{
        Value: val,
        Accept: true,
	}
	fmt.Println(gl.String())
}