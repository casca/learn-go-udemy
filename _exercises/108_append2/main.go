package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

func main() {
	var (
		toppings    []string
		departures  []time.Time
		graduations []int
		lights      []bool
	)

	toppings = append(toppings, "pepperoni", "onions", "extra cheese")
	fmt.Printf("toppings    : %s\n", toppings)
	departures = append(departures, time.Now(), time.Now().Add(time.Minute))
	fmt.Printf("departures  : %s\n", departures)
	graduations = append(graduations, 1998, 2005, 2018)
	fmt.Printf("graduations : %d\n", graduations)
	lights = append(lights, true, false, true)
	fmt.Printf("lights      : %t\n", lights)
}
