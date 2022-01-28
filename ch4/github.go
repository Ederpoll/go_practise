package main

import "time"

type Issue struct {
	Number    int
	HTMLURL   string `json: "html_url"`
	Title     string
	Stage     string
	User      *User
	CreatedAt time.Time `json: "created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json: "html_url"`
}

const IssuesURL = "https://api.github.com/search/issues"
