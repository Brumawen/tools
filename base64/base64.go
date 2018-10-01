package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	path := flag.String("file", "", "Path to the source file to convert.")
	decode := flag.Bool("decode", false, "Decodes the file from text to binary.")
	dest := flag.String("dest", "", "Path to the destination file path.")
	flag.Parse()

	if *path == "" {
		fmt.Println("Source file path not specified.")
		return
	}
	dp := *dest
	if dp == "" {
		dp = *path
		if *decode {
			dp = dp + ".bin"
		} else {
			dp = dp + ".b64"
		}
	}
	b, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if *decode {
		if r, err := base64.StdEncoding.DecodeString(string(b)); err != nil {
			fmt.Println(err.Error())
		} else {
			if err := ioutil.WriteFile(dp, r, 0666); err != nil {
				fmt.Println(err.Error())
			}
		}
	} else {
		b64 := base64.StdEncoding.EncodeToString(b)
		if err = ioutil.WriteFile(dp, []byte(b64), 0666); err != nil {
			fmt.Println(err.Error())
		}
	}
}
