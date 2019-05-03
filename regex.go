package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("/home/jhz126/Desktop/docTrain.txt")
	if err != nil {
		fmt.Println("read file is failed")
	}
	src := string(file)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	fmt.Println(src)
	dl := []byte(src)
	errW := ioutil.WriteFile("/home/jhz126/Desktop/docTrainFilter.txt", dl, 0644)
	if errW != nil {
		fmt.Println("write file is failed")
	}
}
