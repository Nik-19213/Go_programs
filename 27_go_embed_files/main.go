package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed text.txt
var s string

//go:embed assets/*
var assets embed.FS

func main() {
	fmt.Println(s)

	fs := http.FileServer(http.FS(assets))
	http.ListenAndServe(":8000", fs)
}
