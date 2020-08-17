package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var bigArray [size]int
	_ = bigArray

	// 2. print the memory usage (use the report func).
	report("after declaring an array")

	// 3. copy the array to a new array.
	bigArrayCopy := bigArray
	_ = bigArrayCopy

	// 4. print the memory usage
	report("after copying the array")

	// 5. pass the array to the passArray function
	passArray(bigArray)

	// 6. convert one of the arrays to a slice
	bigSliceFromArray := bigArray[:]
	_ = bigSliceFromArray

	// 7. slice only the first 1000 elements of the array
	slice1k := bigSliceFromArray[:1e3]
	_ = slice1k

	// 8. slice only the elements of the array between 1000 and 10000
	slice9k := bigSliceFromArray[1e3:1e5]
	_ = slice9k

	// 9. print the memory usage (report func)
	report("after slicings")

	// 10. pass the one of the slices to the passSlice function
	passSlice(bigSliceFromArray)

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("bigArray         : %d bytes\n", unsafe.Sizeof(bigArray))
	fmt.Printf("bigArrayCopy     : %d bytes\n", unsafe.Sizeof(bigArrayCopy))
	fmt.Printf("bigSliceFromArray: %d bytes\n", unsafe.Sizeof(bigSliceFromArray))
	fmt.Printf("slice1k          : %d bytes\n", unsafe.Sizeof(slice1k))
	fmt.Printf("slice9k          : %d bytes\n", unsafe.Sizeof(slice9k))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
