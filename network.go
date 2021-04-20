package main

import "net/http"

func get_response(url string) http.Response {
	response, _ := http.Get("https://www.wikipedia.org")
	return *response
}
