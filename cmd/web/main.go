package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//指定运行端口：go run .\main.go -addr=":4005"，默认是4000。
	//查看选项：go run .\main.go -help
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages
	//生成日志文件的话，执行： go run main.go >>info.log 2>>error.log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	mux := app.routes()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
		// Add Idle, Read and Write timeouts to the server.
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// log.Printf("Starting server on :%s\n", *addr)
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()

	// log.Fatal(err)
	errorLog.Fatal(err)
}
