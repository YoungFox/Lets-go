package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func xx(w http.ResponseWriter, r *http.Request) {

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("111111111111111111111")
	u, _ := url.Parse("")
	fmt.Println(u)
	proxy := httputil.NewSingleHostReverseProxy(u)
	fmt.Println(r.Header, r.URL, r.Host)
	r.Host = r.URL.Host
	proxy.ServeHTTP(w, r)

	// 	fmt.Println((*x).Status)
	// 	fmt.Println(y)
	// 	proxy.ServeHTTP(w, r)
	// fmt.Println((*x).Body)
	// js, _ := json.Marshal(x.Body)
	// fmt.Println(js)
	// fmt.Println(err)

	// content, _ := ioutil.ReadAll(x.Body)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(content)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("")))

	http.HandleFunc("/h5/", Redirect)

	http.HandleFunc("/api/", Redirect)

	http.ListenAndServe(":8080", nil)

}
