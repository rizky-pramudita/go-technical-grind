package main
import "fmt"
func longestValidParentheses(s string) int {
    charArr := []rune(s)
	for i := 0; i < len(charArr); i++ {
		fmt.Println(charArr[i])
	}
	return 0
}

func main() {
	s := "(()"
	fmt.Println(longestValidParentheses(s))
}