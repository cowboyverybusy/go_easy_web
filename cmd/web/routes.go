package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", app.Hello)
	mux.HandleFunc("/say", app.Say)
	// mux.HandleFunc("/showHome", app.showHome)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	//想要隐藏某个目录中的文件，只需要在这个目录创建一个空白的index.html
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
