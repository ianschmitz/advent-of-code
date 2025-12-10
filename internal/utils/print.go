package utils

import "fmt"

func PrettyPrintlnStruct(obj any) {
	fmt.Printf("%+v\n", obj)
}
