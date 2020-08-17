package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("topgoer.com是个不错的网站我是9091里面的")
	// w.Write({"1223"})
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "12344")
}
func main() {
	http.HandleFunc("/topgoer", sayHello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
