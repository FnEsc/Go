package main

import (
	"fmt"
	"regexp"
)

func re_match_byte() {
	//func Match(pattern string, b []byte) (matched bool, err error)
	matched, err := regexp.Match("^abc.*z$", []byte("abcdefgz"))
	fmt.Println(matched, err) //true nil

	matched, err = regexp.Match("^abc.*z$", []byte("bcdefgz"))
	fmt.Println(matched, err) //false nil
}
