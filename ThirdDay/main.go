package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) //取得したデータはバイト型なので文字列型に変換する作業が必要になる

}
