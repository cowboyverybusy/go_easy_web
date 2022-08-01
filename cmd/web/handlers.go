package main

import (
	"fmt"
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
