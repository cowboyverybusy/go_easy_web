package main

import (
	"bytes"
	"fmt"
	"go_easy_web/internal/data"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
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
	// err = ts.Execute(w, user)

	//将模板渲染分为两个阶段。首先，我们应该通过将模板写入缓冲区来进行“试验”渲染。如果失败，我们可以用一条错误消息响应用户。
	// 但如果它能工作，我们就可以将缓冲区的内容写入http.ResponseWriter。
	buf := new(bytes.Buffer)
	err = ts.Execute(buf, user)
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	mid, _ := strconv.Atoi(id)
	movie := &data.Movie{
		ID:          mid, //如果没有传递参数，则这里为0（空值），不展示这个字段
		Title:       "独行月球",
		CreateAt:    time.Now(),
		Version:     1,
		TicketPrice: 43.5,
		Runtime:     20,
	}
	data := envelope{"movie": movie}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		// app.errorLog.Println("showMovie err:", err.Error())
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}

}

//测试助手函数background能否捕获恐慌，已经执行goroutine
func (app *application) outputWord(w http.ResponseWriter, r *http.Request) {
	//⭐后台执行的协程出现恐慌的时候会导致整个程序终止，所以我们需要捕获恐慌，并自动恢复。
	// go func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			app.errorLog.Println(err)
	// 		}
	// 	}()
	// 	fmt.Println("hello world")
	// 	panic("it is panic!")
	// }()

	app.background(func() {
		// panic("it is panic!")
		fmt.Println("hello world")
	})
	time.Sleep(2 * time.Second)
	fmt.Println("结束")
}
