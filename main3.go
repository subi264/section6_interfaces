package main

import (
	"fmt"
	"net/http"
	"os"
)
type logWriterstruct{}


func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	
	lw := logWriter{}

	io.Copy(os.Stdout, resp.Body)


func (logWriter) Write(bs []byte) (int,error) {
	//return 1,nil
	fmt.Println(string(bs))

	return len(bs),nil

}