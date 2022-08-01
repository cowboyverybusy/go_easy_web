package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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

type User struct {
	UserId   int
	UserName string
	UserDesc string
	UserAddr string
	Likes    []Like
}
type Like struct {
	Name  string
	Score int
}

//模板可以调用结构体的方法
func (u *User) GetUserAge() int {
	return 10 * u.UserId
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
	user.UserId = 9
	user.UserName = "cowboy"
	user.UserDesc = "is a very busy but robust boy"
	user.UserAddr = "guangdong"
	user.Likes = []Like{
		{"play basketball", 60},
		{"study", 80},
		{"swim", 40},
	}
	// user.UserDesc = "<script>alert('xss attack')</script>"
	//传入user(如果要传递多个结构体，需要定义一个大的结构体（最好新建一个文件），包含若干小结构体)
	err = ts.Execute(w, user)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
