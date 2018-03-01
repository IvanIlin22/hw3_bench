package main

import (
	"io"
	"sync"
	"os"
	"bufio"
	"github.com/mailru/easyjson"
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

	//var seenBrowsers []string
	//var writeUniq bool
	//var notSeenBefore bool
	//var isAndroid bool
	//var isMSIE bool


	in := bufio.NewScanner(file)

	//count := 0
	//
	userWK := dataPool.Get().(*UserWK)

	for in.Scan() {

		temp := in.Bytes()

		easyjson.Unmarshal(temp, &UserWK{})

		//isAndroid = false
		//isMSIE = false
		//
		//for _, browser := range userWK.Browsers {
		//	notSeenBefore = true
		//	writeUniq = false
		//
		//	if strings.Contains(browser, "Android") {
		//		isAndroid = true
		//		writeUniq = true
		//	}
		//
		//	if strings.Contains(browser, "MSIE") {
		//		isMSIE = true
		//		writeUniq = true
		//	}
		//
		//	if writeUniq == true {
		//
		//		for _, item := range seenBrowsers {
		//			if item == browser {
		//				notSeenBefore = false
		//			}
		//		}
		//
		//		if notSeenBefore {
		//			seenBrowsers = append(seenBrowsers, browser)
		//		}
		//
		//	}
		//}
		//
		dataPool.Put(userWK)
		//count ++
		//if !(isAndroid && isMSIE) {
		//	continue
		//}
		//
		//email := strings.Replace(userWK.Email,  "@", " [at] ", -1)
		//fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", count, userWK.Name, email))
	}


	//fmt.Fprintln(out, fmt.Sprintf("found users:"))
	//
	//fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}