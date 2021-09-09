package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello world")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("0000", r.Form.Get("username"))
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))
		fmt.Println(timestamp, token)
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("username")); m {
			fmt.Println("是中文名字")
		} else {
			fmt.Println("不是中文名字")
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
