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
	"github.com/mailru/easyjson/jwriter"
	"github.com/mailru/easyjson/jlexer"
	"encoding/json"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
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


func easyjson3486653aDecodeHw3BenchTemp(in *jlexer.Lexer, out *UserWK) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3486653aEncodeHw3BenchTemp(out *jwriter.Writer, in UserWK) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWK) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3486653aEncodeHw3BenchTemp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWK) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3486653aEncodeHw3BenchTemp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWK) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeHw3BenchTemp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWK) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3486653aDecodeHw3BenchTemp(l, v)
}

