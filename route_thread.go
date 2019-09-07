package main

import (
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

