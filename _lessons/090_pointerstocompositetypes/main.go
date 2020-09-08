package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--- ARRAYS")
	arrays()

	fmt.Println("--- SLICES")
	slices()

	fmt.Println("--- MAPS")
	maps()

	fmt.Println("--- STRUCTS")
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)
	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("structs()      : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("structs()      : %p %+v\n", &myHouse, myHouse)
}

func addRoomPtr(h *house) {
	h.rooms++ // same as (*h).rooms++
	fmt.Printf("addRoomPtr()   : %p %+v\n", h, h)
	fmt.Printf("&myHouse.name  : %p\n", &h.name) // same as &(h.name)
	fmt.Printf("&myHouse.rooms : %p\n", &h.rooms)
}

func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom()      : %p %+v\n", &h, h)
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)

	// &confused["one"] // not allowed, Go runtime manages the address of map elements
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	// fmt.Println(dirs)
	fmt.Printf("slices list    : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list    : %p %q\n", &dirs, dirs)
}

func upPtr(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "HEISEN BUG")

	fmt.Printf("upPtr list     : %p %q\n", list, list)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
		fmt.Printf("up.list[%d]     : %p\n", i, &list[i])
	}

	list = append(list, "HEISEN BUG")

	fmt.Printf("up list        : %p %q\n", &list, list)
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("array nums     : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums      : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("incr.nums[%d]   : %p\n", i, &nums[i])
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPts nums: %p\n", &nums)
	for i := range nums {
		nums[i]++
		// (*nums)[i]++
	}
}
