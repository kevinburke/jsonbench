package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kevinburke/jsonbench/gojson/src/jsonbench"
)

var codeJSON []byte
var codeStruct jsonbench.JSONBlob

func codeInit() {
	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	codeJSON = data

	if err := json.Unmarshal(codeJSON, &codeStruct); err != nil {
		panic("unmarshal code.json: " + err.Error())
	}

	if data, err = json.Marshal(&codeStruct); err != nil {
		panic("marshal code.json: " + err.Error())
	}

	if !bytes.Equal(data, codeJSON) {
		println("different lengths", len(data), len(codeJSON))
		for i := 0; i < len(data) && i < len(codeJSON); i++ {
			if data[i] != codeJSON[i] {
				println("re-marshal: changed at byte", i)
				println("orig: ", string(codeJSON[i-10:i+10]))
				println("new: ", string(data[i-10:i+10]))
				break
			}
		}
		panic("re-marshal code.json: different result")
	}
}

func BenchmarkCodeMarshal(b *testing.B) {
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(&codeStruct); err != nil {
			b.Fatal("Marshal:", err)
		}
	}
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkUnmarshal(b *testing.B) {
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		var r jsonbench.JSONBlob
		if err := json.Unmarshal(codeJSON, &r); err != nil {
			b.Fatal("Unmarshal:", err)
		}
	}
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkUnmarshalReuse(b *testing.B) {
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	var r jsonbench.JSONBlob
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(codeJSON, &r); err != nil {
			b.Fatal("Unmarshal:", err)
		}
	}
}
