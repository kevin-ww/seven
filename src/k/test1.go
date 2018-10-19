package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"reflect"
)

type T struct{}

func (t *T) Foo() string {
	fmt.Println("foo from struct t")
	return `kevin hahaha`
}

func (t *T) withparameters(s string, i int) (string, int) {
	fmt.Printf("%v %v \n", s, i)
	return `kevin hehehe`, 100
}

func (t *T) f1() {
	fmt.Printf("f1 from struct t")
}

type aImpl interface {
	f1()
}

func main() {
	//var t T
	//reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

	//var a aImpl

	t := &T{}

	core(t)

}

func core(impl interface{}) {

	a := impl.(aImpl)

	a.f1()

	v := reflect.ValueOf(a)

	fmt.Printf("%v \n", v)

	method := reflect.ValueOf(impl).NumMethod()

	fmt.Printf("%v \n", method)

	response := reflect.ValueOf(impl).MethodByName(`Foo`).Call(nil)

	fmt.Printf("%v \n", response)

	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf("kevin")
	inputs[1] = reflect.ValueOf(123)
	m := reflect.ValueOf(impl).MethodByName("withparameters")
	m.Addr()
	m.Pointer()
	m.Kind()

	//Kind.

	fmt.Printf("%v",m.Pointer())

	value := reflect.ValueOf(m)

	fmt.Printf("%v \n",value)

	if reflect.ValueOf(m).IsNil() {
		log.Fatal(`no such method`)
	}

	var outputs []reflect.Value

	outputs = m.Call(inputs)

	for _, val := range outputs {
		fmt.Printf("%v\n", val)
	}
	//fmt.Printf("%v\n",outputs)

}
