package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[0:]
	i, e, lm := 0, 0, 0
	i, er := strconv.Atoi(args[1])
	if er != nil {
		print("error: Not a num:", args[1], "==", er)
		return
	}
	e, er = strconv.Atoi(args[2])
	if er != nil {
		print("error: Not a num:", args[2], "==", er)
		return
	}
	lm, er = strconv.Atoi(args[3])
	if er != nil {
		print("error: Not a num:", args[3], "==", er)
		return
	}
	sum_end_or_lim(i, e, lm)
}

func sum_end_or_lim(i, e, lim int) {
	if sum := sum_loop(i, e); sum < lim {
		print("Sum:", sum)
		return
	}
	print("Lim: ", lim)
	return
}

func sum_loop(i, e int) int {
	sum := 0
	for ; i < e; i++ {
		sum += i
	}
	return sum
}

func looptest() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		print(i)
	}
}

func printHello() {
	fmt.Println("hello", "sandeep")
}
