package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
 {
  "Title": "mobydick",
  "Price": 10,
  "Released": 118281600
 },
 {
  "Title": "odyssey",
  "Price": 15,
  "Released": 733622400
 },
 {
  "Title": "hobbit",
  "Price": 25,
  "Released": -62135596800
 }
]`

func main() {
	// l := list{
	// 	{Title: "mobydick", Price: 10, Released: toTimestamp(118281600)},
	// 	{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
	// 	{Title: "hobbit", Price: 25},
	// }

	// data, err := json.MarshalIndent(l, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	var l list
	err := json.Unmarshal([]byte(data), &l)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)

}
