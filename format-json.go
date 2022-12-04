package main

import (
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"github.com/tidwall/gjson"
)


func main() {
	format := "-%s &color(Red){■};[[%s>https://youtu.be/%s]]\n"
	timeLayout := "2006-01-02T15:04:05Z"
	timeFormat := "2006/01/02"
	json, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
	}
	items := gjson.Get(string(json), "items")
	for _, item := range items.Array() {
		d, err := time.Parse(timeLayout, item.Get("snippet.publishedAt").String())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(format, d.Format(timeFormat), item.Get("snippet.title"), item.Get("id.videoId"))
	}
}
