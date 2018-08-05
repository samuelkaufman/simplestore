package main

import (
	"net/http"

	"github.com/samuelkaufman/simplestore/pkg/simplestoreapp"
)

func init() {
	simpleStore := simplestoreapp.New()
	http.Handle("/", simpleStore)
}
