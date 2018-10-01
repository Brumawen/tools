package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	path := flag.String("f", "", "Path to the source file to split.")
	start := flag.Int("s", 0, "Starting offset.")
	length := flag.Int("l", 0, "Length of data to write to the destination file.")
	dest := flag.String("d", "", "Path to the destination file.")
	flag.Parse()

	if *path == "" {
		fmt.Println("Source file path not specified.")
		return
	}
	dp := *dest
	if dp == "" {
		dp = *path + ".split"
	}
	b, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := *start
	if s < 0 {
		s = 0
	}
	l := *length
	if l <= 0 || s+l > len(b) {
		l = len(b) - s
	}
	e := s + l

	err = ioutil.WriteFile(dp, b[s:e], 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
}
