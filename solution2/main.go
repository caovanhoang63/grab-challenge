package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

}

func Solution(S string) string {
	if strings.Trim(S, "0") == "" {
		return "0"
	}
	if Reverse(S) == S {

		return strings.Trim(S, "0")
	}
	a, _ := strconv.Atoi(S)
	freq := make(map[int]int)
	temp := a
	max := 0
	for temp > 0 {
		d := temp % 10
		freq[d]++
		temp /= 10
		if d > max {
			max = d
		}
	}
	var result string

	keys := make([]int, len(freq))
	ik := 0
	for k := range freq {
		keys[ik] = k
		ik++
	}
	sort.Ints(keys)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		leng := len(result)
		kChar := strconv.Itoa(k)
		if leng == 0 || leng%2 == 0 {
			insert := ""
			for i := 0; i < freq[k]; i++ {
				insert += kChar
			}
			result = result[:len(result)/2] + insert + result[len(result)/2:]
		} else {
			n := len(result)
			insert := ""
			for i := 0; i < freq[k]/2; i++ {
				insert += kChar
			}
			result = result[:n/2] + insert + result[n/2:n/2+1] + insert + result[n/2+1:]
		}
	}
	result = strings.Trim(result, "0")

	return result
}

func main() {
	test1 := "54321"
	fmt.Println(Solution(test1))
}
