// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func ToUpper(s string) string {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
// }
// func main() {
// 	fmt.Println(ToUpper("madam"))
// 	fmt.Println(ToUpper("thOmaS"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func FixArticle(s string) string {
// 	words := strings.Fields(s)
// 	vowels := "aeiouh"
// 	for i := 0; i < len(words)-1; i++ {
// 		if strings.ToLower(words[i]) == "a" {
// 			first := strings.ToLower(string(words[i+1][0]))

// 			if strings.Contains(vowels, first) {
// 				words[i] = "an"
// 			}
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
// func main() {
// 	fmt.Println(FixArticle("a man a apple"))
// 	fmt.Println(FixArticle("a anchor"))
// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main() {
// 	var num1 string = "25"

// 	num2, _ := strconv.Atoi(num1)

// 	fmt.Println(num2)
// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main() {
// 	var num1 int = 100

// 	num2 := strconv.Itoa(num1)
// 	fmt.Println(num2)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		var pi string = "3.14"
//		pi2, _ := strconv.ParseFloat(pi, 64)
//		fmt.Println(pi2)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num1 := "10"
// 	num2 := "20"

// 	num3, _ := strconv.Atoi(num1)
// 	num4, _ := strconv.Atoi(num2)

//		num5 := num3 + num4
//		fmt.Println(num5)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num1 float64 = 2.71828

//		var num2 string = strconv.FormatFloat(num1, 'f', -1, 64)
//		fmt.Println(num2)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num1 string = "true"

//		num2, _ := strconv.ParseBool(num1)
//		fmt.Println(num2)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	a := "12"
// 	b := "8"

//		c, _ := strconv.Atoi(a)
//		d, _ := strconv.Atoi(b)
//		e := c + d
//		fmt.Println("12 + 8 =", e)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	count := 5

//		word := strconv.Itoa(count)
//		fmt.Println("you have", word, "messages")
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		num := "36.6"
//		word, _ := strconv.ParseFloat(num, 64)
//		fmt.Println(word)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	var word string = "(cap, 3)"
// 	words := strings.Fields(word)
// 	for x := 0; x < len(words); x++ {
// 		if words[x] == "(cap," {

// 			words[x+1] = strings.Trim(words[x+1], ")")
// 			num, _ := strconv.Atoi(words[x+1])
// 			fmt.Println(num)

// 		}
// 	}

// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	s := "50"
// 	//n := s + 10
// 	b, _ := strconv.Atoi(s)

// 	n := b + 10

//		fmt.Println(n)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func ConvertToInts(nums []string) []int {
// 	intSlice := make([]int, 0, len(nums))

// 	for _, i := range nums {
// 		file, err := strconv.Atoi(i)
// 		if err != nil {
// 			fmt.Println("error file", i)
// 			continue

// 		}
// 		intSlice = append(intSlice, file)
// 	}
// 	return intSlice

// }
// func main() {

// 	data := []string{"10", "20", "30"}

// 	result := ConvertToInts(data)

//		fmt.Println(result)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func ToUpper(s string) string {
//		if len(s) == 0 {
//			return s
//		}
//		return strings.ToUpper(string(s))
//	}
//
//	func main() {
//		fmt.Println(ToUpper("happy"))
//		fmt.Println(ToUpper("madam"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func Punctuations(s string) string {
//		s = strings.ReplaceAll(s, " ,", ",")
//		s = strings.ReplaceAll(s, " .", ".")
//		s = strings.ReplaceAll(s, ". ", ".")
//		s = strings.ReplaceAll(s, "! ", "!")
//		s = strings.ReplaceAll(s, " !", "!")
//		return strings.TrimSpace(s)
//	}
//
//	func main() {
//		fmt.Println(Punctuations("hello ..."))
//		fmt.Println(Punctuations("hello !!"))
//		fmt.Println(Punctuations("hello! ! !"))
//		fmt.Println(Punctuations("!! hello !!"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func upper(s []string, n int) []string {

// 	for i := n; i < len(s); i++ {
// 		s[i] = strings.ToUpper(s[i])
// 	}
// 	return s
// }

//	func main() {
//		fmt.Println(upper([]string{"hello", "thomas", "how", "are", "you", "doing"}, 4))
//	}
