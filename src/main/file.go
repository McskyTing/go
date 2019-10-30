package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return err == nil || os.IsExist(err)
}

func ReadFileByLine(filename string) []string {
	if !FileExist(filename) {
		return []string{}
	}
	f, err := os.Open(filename)
	fmt.Println(err)
	defer f.Close()
	rd := bufio.NewReader(f)
	result := []string{}
	for {
		line, err := rd.ReadString('\n')
		if err == nil || io.EOF == err {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				result = append(result, line)
			}
		}
		if io.EOF == err {
			break
		}
		fmt.Println(err)
	}
	return result
}

func WriteFileByLine(filename string, data []string) {
	f, err := os.Create(filename)
	fmt.Println(err)
	defer f.Close()
	wt := bufio.NewWriter(f)
	for i := range data {
		_, err := wt.WriteString(data[i])
		if io.EOF == err {
			break
		}
		fmt.Println(err)
	}
	wt.Flush()
}

func DeleteFile(filename string) {
	if FileExist(filename) {
		err := os.Remove(filename)
		fmt.Println(err)
	}
}
