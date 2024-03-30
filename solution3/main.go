package main

import (
	"fmt"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	return sum
}

func FindNextIndex(A []int, n int, idx int) int {
	for i := idx + 1; i < n; i++ {
		if A[i] > A[idx] {
			return i
		}
	}
	return n
}

func CheckLastInc(A []int) int {
	flag := 1
	for i := 2; i < len(A); i++ {
		if A[0] != A[i] {
			flag = 0
		}
	}
	if A[0] == A[1] && flag == 1 {
		return 2
	}
	if A[0] != A[1] && flag == 1 && A[1] == A[2] {
		return 1
	}

	return 0
}

var dp = make(map[string]int)

func Backtrack(A []int, n int, C1 int, C2 int, index int) int {
	if A[0] == A[n-1] {
		return 0
	}

	key := fmt.Sprintf("%v-%d", A, index)

	if val, ok := dp[key]; ok {
		return val
	}

	result := 0
	nId := FindNextIndex(A, n, 1)
	if nId == n {
		if check := CheckLastInc(A); check != 0 {
			if check == 1 {
				result = C1
			} else if check == 2 {
				result = C2
			}
		} else {
			temp := make([]int, len(A))
			copy(temp, A)
			temp[0]++
			temp[n-1]++
			result = C2 + Backtrack(temp, n, C1, C2, index+1)
			A[0]++
			result = Min(result, C1+Backtrack(A, n, C1, C2, index+1))
		}

	} else {
		diff := Min(A[nId]-A[1], (A[nId]-A[0])/(nId-1))
		if diff == 0 {
			diff = A[nId] - A[1]
		}
		A[0] += diff
		A[1] += diff
		sort.Ints(A)
		result = C2*diff + Backtrack(A, n, C1, C2, index+1)
	}

	// Lưu trữ kết quả cho khóa này vào map dp
	dp[key] = result

	return result
}

func Solution(A []int, C1 int, C2 int) int {
	n := len(A)
	sort.Ints(A)
	if n == 2 {
		return (A[1] - A[0]) * C1
	}
	if A[0] == A[n-1] {
		return 0
	}
	if C1*2 <= C2 {
		return C1 * (A[n-1]*n - Sum(A))
	}

	res := Backtrack(A, n, C1, C2, 0)
	return res % 1000000000
}

func main() {

	test1 := []int{500000, 0, 500000, 500000}
	C1 := 10000
	C2 := 3000
	fmt.Println(Solution(test1, C1, C2))
}
