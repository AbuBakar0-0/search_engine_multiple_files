package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
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

type fileStruct struct {
	name       string
	path       string
	lineNumber int
}

func main() {

	//fStruct := make([]fileStruct, 0)

	search := "engine"

	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		fmt.Println("ERROR FILE IS EMPTY")
		log.Fatal(err)
	}

	lines := make(map[fileStruct]string)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			//fmt.Println(file.Name(), filepath.Ext(file.Name()))
			// _, error2 := os.Stat("./files/" + file.Name())

			// if errors.Is(error2, os.ErrNotExist) {
			// 	fmt.Println("file does not exist")
			// } else {
			// 	fmt.Println("file exists")
			// }

			content, err := ioutil.ReadFile("./files/" + file.Name())

			if err != nil {
				log.Fatal(err)
			}
			lines[fileStruct{name: file.Name(), path: "./files/" + file.Name()}] += string(content) + "\n"
		}
	}
	// for k, v := range lines {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step2 : Create Arrlines

	arrlines := make(map[fileStruct][]string)

	for k, v := range lines {
		arrlines[k] = append(arrlines[k], strings.Split(v, "\n")...)
	}

	// for k, v := range arrlines {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 3 : Create LineMap

	linemap := make(map[fileStruct][]string)

	for k, v := range arrlines {
		for i := 0; i < len(v); i++ {
			linemap[fileStruct{name: k.name, path: k.path, lineNumber: i + 1}] = strings.Fields(v[i])
		}
	}

	// for k, v := range linemap {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 4 : Create token map

	tokenMap := make(map[string][]fileStruct)

	lineNumbers := make([]fileStruct, 0)

	for k, v := range linemap {
		for i := 0; i < len(v); i++ {
			if search == v[i] {
				lineNumbers = append(lineNumbers, k)
			}
		}
	}

	tokenMap[search] = lineNumbers

	// for k, v := range tokenMap {
	// 	for i := 0; i < len(v); i++ {
	// 		fmt.Println(k, " NAME : ", v[i].name, " => Path : => ", v[i].path, " Line Number : ", v[i].lineNumber)
	// 	}
	// }

	//Step 5: Token Array

	tokens := make([]string, 0)

	for _, v := range linemap {
		for i := 0; i < len(v); i++ {
			if !contains(tokens, v[i]) {
				tokens = append(tokens, v[i])
			}
		}
	}

	// for k, v := range tokens {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 6: Create Trie

	trie := initTrie()

	for i := 0; i < len(tokens); i++ {
		trie.insert(tokens[i])
	}

	found := trie.find(search)
	if found {
		fmt.Printf("Word \"%s\" found in trie\n", search)
		fmt.Println("FOUND : ")

		//Step 8: looking for word in token map

		for i := 0; i < len(tokenMap[search]); i++ {
			fmt.Println(" NAME : ", tokenMap[search][i].name, " => Path : => ", tokenMap[search][i].path, " Line Number : ", tokenMap[search][i].lineNumber, " -> ", linemap[tokenMap[search][i]])
		}

	} else {
		fmt.Printf("Word \"%s\" not found in trie\n", search)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
