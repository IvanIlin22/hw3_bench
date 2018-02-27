package main

import (
	"io"
	"io/ioutil"
	"sync"
	"bytes"
	"fmt"
	"github.com/mailru/easyjson"
	"strings"
)

//easyjson:json
type UserWK struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
}

var dataPool = sync.Pool{
	New: func() interface{} {
		return &UserWK{}
	},
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {

	lines, err := readLines(filePath)

	if err != nil {
		panic(err)
	}

	var seenBrowsers []string

	var writeUniq bool
	var notSeenBefore bool
	var isAndroid bool
	var isMSIE bool

	fmt.Fprintln(out, fmt.Sprintf("found users:"))
	for i, line := range lines {
		userWK := dataPool.Get().(*UserWK)
		easyjson.Unmarshal([]byte(line), userWK)

		isAndroid = false
		isMSIE = false

		for _, browser := range userWK.Browsers {
			notSeenBefore = true
			writeUniq = false

			if strings.Contains(browser, "Android") {
				isAndroid = true
				writeUniq = true
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				writeUniq = true
			}

			if writeUniq == true {

				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}

				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
				}

			}
		}

		dataPool.Put(userWK)

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(userWK.Email,  "@", " [at] ", -1)
		fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", i, userWK.Name, email))
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}


func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}