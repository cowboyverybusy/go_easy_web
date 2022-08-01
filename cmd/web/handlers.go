package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	UserId   int
	UserName string
	UserDesc string
}

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("访问成功")
	w.Write([]byte("hello world,couwboy~~~~"))
}

func (app *application) Say(w http.ResponseWriter, r *http.Request) {
	//获取参数的两种方式
	name := r.FormValue("name")
	desc := r.URL.Query().Get("desc")
	s := fmt.Sprintf("Hey,%s,you are %s", name, desc)
	w.Write([]byte(s))
}

//跳转到html页面(模板循序，a页面调用b模板，b模板又调用a页面中的模板)
func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
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
	user := &User{}
	user.UserId = 1
	user.UserName = "cowboy"
	user.UserDesc = "is a very busy but robust boy"
	// user.UserDesc = "<script>alert('xss attack')</script>"
	//传入user
	err = ts.Execute(w, user)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
