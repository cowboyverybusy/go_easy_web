package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//MarshalIndent 能格式化展示json。但是这个开销更大，正式环境推荐直接使用json.Marshal()
	js, err := json.MarshalIndent(data, "", "\t")
	// js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// The background() helper accepts an arbitrary function as a parameter.
// ⭐后台执行的协程出现恐慌的时候会导致整个程序终止，所以我们需要捕获恐慌，并自动恢复。
// 可能有多个地方用到后台协程，所以封装成助手函数
func (app *application) background(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				app.errorLog.Println(err)
			}
		}()
		fn()
	}()
}

//
