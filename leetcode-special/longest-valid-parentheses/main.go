package main
import "fmt"
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	max, left := 0, 0
	stack := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '('{
			left++
		} else if left > 0 {
			stack[i] = 2 + stack[i-1]
			if i-stack[i] > 0 {
				stack[i] += stack[i-stack[i]]
			}
			if max < stack[i] {
				max = stack[i]
			}
			left--
		}
	}
	return max
}

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses(s))
}