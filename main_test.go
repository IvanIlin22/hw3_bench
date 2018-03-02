package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	SlowSearch(ioutil.Discard)
	FastSearch(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	SlowSearch(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch(ioutil.Discard)
	}
}


//func BenchmarkCopyPreAllocate(b *testing.B) {
//	var result string
//	for n := 0; n < b.N; n++ {
//		bs := make([]byte, cnt*len(sss))
//		bl := 0
//		for i := 0; i < cnt; i++ {
//			bl += copy(bs[bl:], sss)
//		}
//		result = string(bs)
//	}
//	b.StopTimer()
//	if result != expected {
//		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
//	}
//}
//
//func BenchmarkAppendPreAllocate(b *testing.B) {
//	var result string
//	for n := 0; n < b.N; n++ {
//		data := make([]byte, 0, cnt*len(sss))
//		for i := 0; i < cnt; i++ {
//			data = append(data, sss...)
//		}
//		result = string(data)
//	}
//	b.StopTimer()
//	if result != expected {
//		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
//	}
//}