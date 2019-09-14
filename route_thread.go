package main

import (
	"fmt"
	"github.com/polandtyler/chit_chat/data"
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func newThread(writer http.ResponseWriter, r *http.Request) {
	_, err := session(writer, r)
	if err != nil {
		http.Redirect(writer, r, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(writer http.ResponseWriter, r *http.Request) {
	sess, err := session(writer, r)
	if err != nil {
		http.Redirect(writer, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, r, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUuid(uuid)
	if err != nil {
		errorMessage(writer, request, "Cannot read thread")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(writer, &thread, "layout", "private.navbar", "public.navbar")
		}
	}
}

// POST /thread/post
// Create a post
func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUuid(uuid)
		if err != nil {
			errorMessage(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}

}
