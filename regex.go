package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	//resp, err := http.Get("http://www.sohu.com/a/232886309_108002")
	file, err := ioutil.ReadFile("/home/jhz126/Desktop/docTrain.txt")
	// resp, err := http.Get("http://www.163.com")
	// if err != nil {
	// 	fmt.Println("http get error.")
	// }
	if err != nil {
		fmt.Println("read file is failed")
	}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	src := string(file)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
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
	//fmt.Println(strings.TrimSpace(src))
	errW := ioutil.WriteFile("/home/jhz126/Desktop/docTrainFilter.txt", dl, 0644)
	if errW != nil {
		fmt.Println("write file is failed")
	}
}
