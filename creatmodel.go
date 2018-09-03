package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	command := "./word2vec"
	params := []string{"-train", "/home/jhz126/Desktop/docTrainEnd.txt", "-output", "docTrainvectors.bin", "-cbow", "1",
		"-size", "100", "-window", "8", "-negative", "25", "-hs", "0", "-sample", "1e-4",
		"-threads", "20", "-binary", "1", "-iter", "15"}
	execCommand(command, params)
	// commandDist := "./distance"
	// paramsDist := []string{"vectors.bin"}
	// execCommand(commandDist, paramsDist)
}
func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return false
	}
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		//        enc := mahonia.NewDecoder("UTF-8")
		//        goStr := enc.ConvertString(line)
		fmt.Println(line)
	}
	cmd.Wait()
	return true
}
