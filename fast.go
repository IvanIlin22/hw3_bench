package main

import (
	"io"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/mailru/easyjson"
	"sync"
)

//easyjson:json
type UserWK struct {
	Browsers []string `json:"browsers"`
	Company  string   `json:"company"`
	Country  string   `json:"country"`
	Email    string   `json:"email"`
	Job      string   `json:"job"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}

var dataPool = sync.Pool{
	New: func() interface{} {
		return &UserWK{}
	},
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	foundUsers := ""

	lines := strings.Split(string(fileContents), "\n")

	seenBrowsers := []string{}

	fmt.Fprintln(out, "found users:")


	for i, line := range lines {
		user := dataPool.Get().(*UserWK)
		err := easyjson.Unmarshal([]byte(line), user)
		if err != nil {
			panic(err)
		}


		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {

			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
				}
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
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
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}