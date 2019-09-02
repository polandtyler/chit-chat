package main

import (
	"github.com/polandtyler/chit_chat/data"
	"net/http"
)

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)

	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//files := []string{
	//	"templates/layout.html",
	//	"templates/navbar.html",
	//	"templates/index.html",
	//}
	//
	//templates := template.Must(template.ParseFiles(files...))
	//threads, err := data.Threads()
	//if err == nil {
	//	templates.ExecuteTemplate(w, "layout", threads)
	//}
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public_navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private_navbar", "index")
		}
	}
}
