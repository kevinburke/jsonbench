/*
 * This was used to parse jsonbench_sample.txt and pretty print the result.
 *
 * Note jsonbench_sample.txt contains two distinct structures, which have been
 * combined into one in go/data.json and go/pretty.json.
 */
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("jsonbench_sample.txt")
	var d json.RawMessage
	dec := json.NewDecoder(f)
	for {
		err := dec.Decode(&d)
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		err = json.Indent(&out, d, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		out.WriteTo(os.Stdout)
	}
}
