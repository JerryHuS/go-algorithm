package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
}

func scanf() {
	//按行获取
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)

	//循环获取
	var strInput string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		strInput = in.Text()
	}
	fmt.Println(strInput)

	//循环获取，遇到指定符号停止
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil || strings.TrimSuffix(input, "\n") == "" {
			break
		}
	}
}
