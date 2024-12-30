package main
import (
	"fmt"
	"strconv")

func calPoints(operations []string) int{
	var sum int
	var stack []int
	for _, op := range operations {
		opInt, _ := strconv.Atoi(op)
		switch op {
			case "C":
				sum -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			case "D":
				sum += stack[len(stack)-1]*2
				stack = append(stack, stack[len(stack)-1]*2)
			case "+":
				sum += stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, stack[len(stack)-1] + stack[len(stack)-2])
			default:
				sum += opInt
				stack = append(stack, opInt)
		}
	}
	return sum	
}

func main() {
	ops := []string{"5","2","C","D","+"}
	fmt.Println(calPoints(ops))
}