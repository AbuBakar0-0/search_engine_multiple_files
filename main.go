package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		fmt.Println("ERROR FILE IS EMPTY")
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), filepath.Ext(file.Name()))
	}

	_, error := os.Stat("./files/")

	// if errors.Is(err, os.ErrNotExist) {
	// 	fmt.Println("file does not exist")
	// } else {
	// 	fmt.Println("file exists")
	// }

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := string(content)

}
