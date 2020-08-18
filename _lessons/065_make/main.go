package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	// var upTasks []string
	// upTasks := make([]string, len(tasks))
	upTasks := make([]string, 0, len(tasks))
	// upTasks = upTasks[:cap(upTasks)]
	s.Show("upTasks", upTasks)

	// for i, task := range tasks {
	// 	upTasks[i] = strings.ToUpper(task)
	// 	s.Show("upTasks", upTasks)
	// }
	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
