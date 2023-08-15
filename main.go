package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader(`
	{
		"model": "gpt-3.5-turbo-16k",
		"messages": [{"role": "user", "content": "你是谁?"}],
		"temperature": 0.7,
		"stream": true
	}
`)
	base_url := "**************************"
	if req, err := http.NewRequest("POST", base_url, reader); err != nil {
		fmt.Println(err)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer sk-ql6*************K17P5wPl3LWlR2Ek")
		client := http.Client{}
		if resp, err := client.Do(req); err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			answer_reader := bufio.NewReader(resp.Body)
			for {
				line, err := answer_reader.ReadString('\n')
				if err != nil {
					break
				} else {
					if len(line) < 10 {
						continue
					}
					//fmt.Println(line)
					as_s := strings.Index(line, "\"content\":\"") + len("\"content\":\"")
					as_e := strings.Index(line[as_s:], "\"")
					if as_e == -1 {
						break
					}
					as := line[as_s : as_s+as_e]
					fmt.Printf("%s", as)
				}
			}
		}
	}
}
