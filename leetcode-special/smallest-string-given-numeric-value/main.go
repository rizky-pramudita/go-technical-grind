package main
import ("fmt")

func isValid(input string) bool {
    closeMap := {{")", "("}, {"]", "["}, {"}", "{"}}
    leftArray := ([]string)

    for i:=0;i<len(input);i++{
		if input[i] == "(" || 
    }
	return true
}

func main(){
    input := "(())"
    fmt.Println(isValid(input))
}
