package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	resolution = "1920x1080"
	url        = "https://source.unsplash.com/random/"
)

func main() {
	if len(os.Args) < 2 {
	        fmt.Println("[*] Author: Abdullah Khabir https://abdullah.solutions")
		fmt.Println("[*] Download random wallpaper from unsplash based on keywords...")
		fmt.Println("[*] Usage: unsplashwall <keywords>")
		os.Exit(0)
	}

	keywords := os.Args[1]
	wallsDir := os.Getenv("HOME") + "/Pictures/unsplash"
	wallName := "unsplash-" + time.Now().Format("02-01-15-04-05") + ".jpg"
	filePath := wallsDir + "/" + wallName

	if err := os.MkdirAll(wallsDir, 0755); err != nil {
		fmt.Println("Error creating directory: ", err)
		os.Exit(1)
	}

	resp, err := http.Get(url + resolution + "/?" + keywords)
	if err != nil {
		fmt.Println("Error downloading image: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		os.Exit(1)
	}

	fmt.Println("[*] Author: Abdullah Khabir https://abdullah.solutions")
	fmt.Println("Here is your result: ", filePath)
}
