package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int `a:"1111"b:"3333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fmt.Println(rt)
	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fieldName.Tag)
	}

	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}
	fmt.Println("type_Name:", rt.Name())
	fmt.Println("type_NumField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String())
	fmt.Println("type.Kind.String:", rt.Kind().String())
	fmt.Println("type_String()=", rt.String())

	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v \n", i, rt.Field(i).Name)
	}
}
