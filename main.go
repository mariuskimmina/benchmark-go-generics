package main

import (
	"fmt"
	"reflect"
)


func main() {
    fmt.Println(addInt(1, 1))
    fmt.Println(addFloat(1, 1))
    fmt.Println(addString("1", "1"))
    fmt.Println(addTypeSwitch(1, 1))
    fmt.Println(addGenerics(1, 1))
    fmt.Println(addGenericsTypeSet(1, 1))
    fmt.Println(addReflect(1, 1))
}

func addInt(a, b int) int {
    return a + b
}

func addFloat(a, b float64) float64 {
    return a + b
}

func addString(a, b string) string {
    return a + b
}

func addReflect(a, b interface{}) interface{} {
    if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
        return a.(int) + b.(int)
    }
    if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
        return a.(float64) + b.(float64)
    }
    if reflect.TypeOf(a).Kind() == reflect.String && reflect.TypeOf(b).Kind() == reflect.String {
        return a.(string) + b.(string)
    }

    return nil
}
func addTypeSwitch(a, b interface{}) interface{} {
    switch a.(type) {
    case int:
        return a.(int) + b.(int)
    case float64:
        return a.(float64) + b.(float64)
    case string:
        return a.(string) + b.(string)
    }
    return nil
}

func addGenerics[T int | float64 | string](a, b T) T {
    return a + b
}

type TypeSet interface {
    int | float64 | string
}

func addGenericsTypeSet[T TypeSet](a, b T) T {
    return a + b
}

