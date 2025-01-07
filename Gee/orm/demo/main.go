package main

import (
	"fmt"
	"reflect"
)

//
//func main() {
//	var x float64 = 3.4
//	fmt.Println("type:", reflect.TypeOf(x))
//}

type MyStrucy struct {
	Name string
}

func (s *MyStrucy) Talk() {
	fmt.Println("Hello World")
}
func main() {
	tmp := &MyStrucy{"penge"}
	value := reflect.ValueOf(tmp)
	method := value.MethodByName("Talk")
	method.Call(nil)
}
