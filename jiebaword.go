package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/yanyiwu/gojieba"
)

func main() {
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	file, err := ioutil.ReadFile("/home/jhz126/Desktop/docTrainFilter.txt")

	//file, err := ioutil.ReadFile("/home/jhz126/Desktop/123.txt")
	if err != nil {
		fmt.Println("read file is failed")
	}
	words = x.Cut(string(file), use_hmm)
	wordsDiv := strings.Join(words, " ")
	dl := []byte(wordsDiv)
	errW := ioutil.WriteFile("/home/jhz126/Desktop/docTrainEnd.txt", dl, 0644)
	if errW != nil {
		fmt.Println("write file is failed")
	}
}
