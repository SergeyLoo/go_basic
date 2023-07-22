package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	perem2 := 35
	perem1 := "104"
	perem3, _ := strconv.Atoi(perem1)
	perem4 := strconv.Itoa(perem2)
	fmt.Println(reflect.TypeOf(perem3), perem3, reflect.TypeOf(perem4), perem2)

}
