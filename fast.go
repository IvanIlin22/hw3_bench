package main

import (
	"io"
	"sync"
	"os"
	"bufio"
	"strings"
	"fmt"
	"github.com/mailru/easyjson"
	"bytes"
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

	file, _ := os.Open(filePath) // For read access.
	defer file.Close()

	var seenBrowsers []string
	var writeUniq bool
	var notSeenBefore bool
	var isAndroid bool
	var isMSIE bool

	in := bufio.NewScanner(file)

	fmt.Fprintln(out, fmt.Sprintf("found users:"))

	count := 0
	for in.Scan() {
		count ++
		row := in.Bytes()

		if bytes.Contains(row, []byte("Android")) == false && bytes.Contains(row, []byte("MSIE")) == false {
			continue
		}

		user := dataPool.Get().(*UserWK)
		easyjson.Unmarshal(row, user)

		isAndroid = false
		isMSIE = false

		for _, browser := range user.Browsers {
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

		dataPool.Put(user)
		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email,  "@", " [at] ", -1)
		fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", count - 1, user.Name, email))
	}


	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}