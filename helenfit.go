package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

// "This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error."
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析參數，預設是不會解析的
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, `
	<h1>HELENFIT.COM Welcomes You!</h1>
	<h2>The Beggining.</h2>
	
	`) //這個寫入到 w 的是輸出到客戶端的
}

func handler(w http.ResponseWriter, r *http.Request) {
	subsliceofPath := r.URL.Path[1:]
	if subsliceofPath == "" {
		subsliceofPath = "www.helenfit.com"
	}
	fmt.Fprintf(w, `
	<h1>🇹🇼 Hi there, I love %s!</h1>
	<hr>
	<h2>-dev branch-</h2>
	`, subsliceofPath)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	bodyContent := `
	<h1>%s</h1>
	<hr>
	<div>
	<p>www.helenfit.com - <b>%s</b></p>
	</div>
	
	`
	fmt.Fprintf(w, bodyContent, p.Title, p.Body)
}

func main() {
	// http.HandleFunc("/", sayhelloName)       //設定存取的路由
	http.HandleFunc("/", handler)            //設定存取的路由
	http.HandleFunc("/view/", viewHandler)   //設定存取的路由
	err := http.ListenAndServe(":8877", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
