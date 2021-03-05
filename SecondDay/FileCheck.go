package main

import (
	"fmt"
	"os"
)

func exists(name string) bool {
	_, err := os.Stat(name)

	return !os.IsNotExist(err)

	/*IaExistではなくIsNotExistを反転させるのかというと
	前者ではErrExistはエラーとして返ってこないので
	https://golang.org/src/os/error_plan9.go
	の11行目のdoes not existと~~os/src/error.go
	の17行目のErrNotExist = errors.New("file does not exist)とを比較するため*/
}
func main() {
	if exists("data.db") == false {
		fmt.Println("err")
	}
}
