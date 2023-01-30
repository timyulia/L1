package main

import (
	"fmt"
	"reflect"
)

func deter(something interface{}) {
	fmt.Printf("something is %s\n", reflect.TypeOf(something)) //специальная функция для определения типа
}

func deter2(something interface{}) {
	switch something.(type) { //переключение по возможным значениям типа
	case int:
		fmt.Println("something is int")
	case string:
		fmt.Println("something is string")
	case bool:
		fmt.Println("something is bool")
	case chan int:
		fmt.Println("something is chan int")
	}
}

func main() {
	var someInt interface{} = 23
	var someString interface{} = "texas"
	var someBool interface{} = true
	var someChannel interface{} = make(chan int)

	deter(someInt)
	deter2(someInt)
	deter(someString)
	deter2(someString)
	deter(someBool)
	deter2(someBool)
	deter(someChannel)
	deter2(someChannel)
}
