package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	// v := url.Values{}
// 	// v.Set("name", "Ava")
// 	// v.Add("friend", "Jess")
// 	// v.Add("friend", "Sarah")
// 	// v.Add("friend", "Zoe")
// 	// // v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
// 	// fmt.Println(v.Get("name"))
// 	// fmt.Println(v.Get("friend"))
// 	// fmt.Println(v["friend"])

// 	fmt.Println("method:", r.Method) //获取请求的方法
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.gtpl")
// 		log.Println(t.Execute(w, nil))
// 	} else {
// 		r.ParseForm()

// 		//请求的是登录数据，那么执行登录的逻辑判断
// 		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
// 		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
// 		// template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

// 		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
// 		_ = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
// 	}
// }

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
