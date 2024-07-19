package mystrings

import "fmt"

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result

}

func sayHello() {
	fmt.Println("hello!");
}
