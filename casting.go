package main

import (
	"fmt"
	"strconv"
)

var num int
var i int
var k int64
var f float64
var s string
var err error

func main() {
	fmt.Sscanf("57", "%d", &num)
	i, err = strconv.Atoi("350") // i == 350
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(num, i, k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(340)
	fmt.Println(k, f, s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)
}

