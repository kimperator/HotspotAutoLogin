package main

import (
	"io"
	"net/http"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
}
