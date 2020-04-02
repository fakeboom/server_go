package main

import (
	"fmt"
	"log"
	"net/http"
)

// 返回静态页面
func handleGet(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "world")
	//writer.Write([]byte("world"))
}

func main() {
	http.HandleFunc("/hello", handleGet)

	fmt.Println("Running at port 80 ...")

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
