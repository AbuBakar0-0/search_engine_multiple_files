package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func main() {
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		fmt.Println("ERROR FILE IS EMPTY")
		log.Fatal(err)
	}
	lines := ""
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fmt.Println(file.Name(), filepath.Ext(file.Name()))
			_, error2 := os.Stat("./files/" + file.Name())

			if errors.Is(error2, os.ErrNotExist) {
				fmt.Println("file does not exist")
			} else {
				fmt.Println("file exists")
			}

			content, err := ioutil.ReadFile("./files/" + file.Name())

			if err != nil {
				log.Fatal(err)
			}

			lines += string(content)
		}
	}

	fmt.Println(lines)

}
