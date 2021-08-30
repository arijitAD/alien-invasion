package main

import (
	"fmt"
	"os"
)

const (
	title = "$RAW_TITLE"
	body = "$RAW_BODY"
)

func main() {
	fmt.Println("Title: ", os.Getenv(title))
	fmt.Println("Body: ", os.Getenv(body))
}