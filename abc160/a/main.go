package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := newScanner()
	var s = sc.readString()
	var s2, s3, s4, s5 rune
	for i, v := range s {
		if i == 2 {
			s2 = v
		}
		if i == 3 {
			s3 = v
		}
		if i == 4 {
			s4 = v
		}
		if i == 5 {
			s5 = v
		}
	}
	if s2 == s3 && s4 == s5 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

//////////////////
//////////////////
//////////////////

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func pow(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("exponent (%d) should not be a minus", exponent))
	}

	answer := 1
	for i := 0; i < exponent; i++ {
		answer *= base
	}
	return answer
}

func ceil(divident, dividor int) int {
	if dividor == 0 {
		panic("dividor should not be 0")
	}

	quo := divident / dividor
	rem := divident % dividor

	if rem != 0 {
		if (divident > 0 && dividor > 0) ||
			(divident < 0 && dividor < 0) {
			return quo + 1
		}
	}
	return quo
}

// -   sortutil

func reverseInts(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

func reverseStrings(a []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
}

// -   io

type scanner struct {
	bufScanner *bufio.Scanner
}

func newScanner() *scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &scanner{bufSc}
}

func (sc *scanner) readString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

func (sc *scanner) readInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}
