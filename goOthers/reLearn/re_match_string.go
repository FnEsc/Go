package main

import (
	"fmt"
	"regexp"
)

func re_match_string() {
	//func MatchString(pattern string, s string) (matched bool, err error)
	matched, err := regexp.MatchString("^abc.*z$", "abcdefgz")
	fmt.Println(matched, err) //true <nil>

	matched, err = regexp.MatchString("^abc.*z$", "bcdefgz")
	fmt.Println(matched, err) //false <nil>
}
