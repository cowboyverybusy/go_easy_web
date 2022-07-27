package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

/*
方式一
自己声明servemux。这种方式比较安全，推荐使用
*/
func main() {
	//指定运行端口：go run .\main.go -addr=":4005"，默认是4000。
	//查看选项：go run .\main.go -help
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages
	//生成日志文件的话，执行： go run main.go >>info.log 2>>error.log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/say", say)
	mux.HandleFunc("/showHome", showHome)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	//想要隐藏某个目录中的文件，只需要在这个目录创建一个空白的index.html
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// log.Printf("Starting server on :%s\n", *addr)
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	// log.Fatal(err)
	errorLog.Fatal(err)
}

/*
方式二
没有声明servemux，其实是net/http底层帮你初始化了默认的servemux：设置了全局的变量DefaultServeMux：
var DefaultServeMux = NewServeMux()
虽然这种方式代码稍微短一些，但是还是推荐使用方式一。因为DefaultServeMux是全局变量，
任何包都可以访问它并注册路由，包括应用程序导入的任何第三方包。如果其中一个第三方软件包被破坏，他们可能会使用DefaultServeMux向web公开恶意处理程序。
因此，为了安全起见，通常最好避免使用DefaultServeMux和相应的辅助函数。使用您自己的局部范围的servemux
*/
func main2() {
	http.HandleFunc("/hello", hello)
	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", nil)
	log.Fatal(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world,couwboy~~~~"))
}

func say(w http.ResponseWriter, r *http.Request) {
	//获取参数的两种方式
	name := r.FormValue("name")
	desc := r.URL.Query().Get("desc")
	s := fmt.Sprintf("Hey,%s,you are %s", name, desc)
	w.Write([]byte(s))
}

//跳转到html页面(模板循序，a页面调用b模板，b模板又调用a页面中的模板)
func showHome(w http.ResponseWriter, r *http.Request) {
	//相对路径
	// ts, err := template.ParseFiles("../../ui/html/home.html")
	files := []string{
		"../../ui/html/home.html",
		"../../ui/html/base.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
