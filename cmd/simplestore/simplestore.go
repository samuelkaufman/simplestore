package main

import (
	"net/http"

	"github.com/samuelkaufman/simplestoreapp"
)

func init() {
	simpleStore := simplestoreapp.New()
	http.Handle("/", simpleStore)
}
