package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func HttpPostByJSON(accessUrl string, json string, redo int) error {
	fmt.Println(json)
	fmt.Println("post write json bytes:" + strconv.Itoa(len(json)))
	for j := 1; j <= redo; j++ {
		req_new := bytes.NewBuffer([]byte(json))
		request, err := http.NewRequest("POST", accessUrl, req_new)
		if err == nil {
			request.Header.Set("Content-Type", "application/json;charset=UTF-8")
			client := http.Client{}
			response, err1 := client.Do(request)
			if err1 == nil {
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					fmt.Println("Unknown error in sending Email")
				} else {
					resp := string(body)
					if strings.Contains(resp, "\"code\":\"200\"") {
						return nil
					} else {
						fmt.Println(string(body))
					}
				}
			} else {
				fmt.Println(err1)
			}
		} else {
			fmt.Println(err)
		}
	}
	return errors.New("Fail to send email notification")
}
