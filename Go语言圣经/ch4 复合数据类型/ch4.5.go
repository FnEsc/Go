/*
4.5 JSON
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const templ = `{{.TotalCount}} issues
{{range .Items}}--------
Number: 	{{.Number}}
User: 		{{.User.Login}}
Title:		{{.Title | printf "%.64s}}
Age: 		{{.CreareAt | daysAgo}} days
{{end}}
`

func main() {
	// 数据结构 vs JSON格式
	type Movie struct {
		Title  string
		Year   int  `json:"released"`        // 成员Tag
		Color  bool `json:"color,omitempty"` // 成员Tag
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}

	// data, err := json.Marshal(movies) // Marshal 函数返回一个编码后的字节slice, 没有空白缩进
	data, err := json.MarshalIndent(movies, "", "	") // MarshalIndent 行缩进和层级缩进
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
