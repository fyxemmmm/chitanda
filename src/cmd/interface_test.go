package main

import (
	"fmt"
	"testing"
)

type Animal interface {
	eat()
}

type myType func(s string) string

func TestInterface(b *testing.T)  {
	new(myType).eat()
	new(myType).eat()
	new(myType).eat2()
	//i := get()
	//ref_i := reflect.ValueOf(i)
	//m := new(myType)
	//m_ref := reflect.ValueOf(m).Elem()
	//if ref_i.Type().ConvertibleTo(m_ref.Type()) {
	//	m_ref.Set(ref_i)
	//	fmt.Println("aaa")
	//	m_ref.Interface().(Animal).eat()
	//}
}

func (a myType) eat() {
	//fmt.Println("eattttt")
	fmt.Printf("%p\n", a)
}

func (a myType) eat2() {
	//fmt.Println("eattttt")
	fmt.Printf("%p\n", a)
}

func get() interface{} {
	return func(s string) string {
		fmt.Println("in i ...")
		return "feixiang"
	}
}
