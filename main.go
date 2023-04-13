package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func getId(url string) string {
	
	parts := strings.Split(url, "/")
	if parts[2] != "open.spotify.com"{
		panic(errors.New("bad url"))
	}

	id := strings.Split(parts[4], "?")

	return id[0]
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	id := getId(os.Args[1])
	fmt.Println("ID is " + id + " (It's copied to your clipboard)")
    
	clipboard.Write(clipboard.FmtText, []byte(id))
}

