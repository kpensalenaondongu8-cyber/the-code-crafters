//reversing a string without using any strings built in package.

// package main

// import "fmt"

// func reverseString(s string) string {
// 	runes := []rune(s)

// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
// func main() {
// 	input := "rider\n"
// 	word := "Thomas\n"

//		rogue := input + word
//		fmt.Println(reverseString(rogue))
//	}

// write function that finds the maximun element
// package main

// import "fmt"

// func FindMax(num []int) int {

//		MaxNumber := num[0]
//		for _, ch := range num {
//			if ch > MaxNumber {
//				MaxNumber = ch
//			}
//		}
//		return MaxNumber
//	}
//
//	func main() {
//		fmt.Println(FindMax([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
//	}
//

// 8
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func FixArticle(s string) string {
// 	s = strings.ReplaceAll(s, "A", "An")
// 	s = strings.ReplaceAll(s, "An book", "A book")
// 	return s
// }
// func main() {
// 	fmt.Println(FixArticle("There it was. A amazing rock. A honest man. An book."))
// }

//7
// package main

// import (
// 	"fmt"
// )

// func FixAorAn(s string) string {
// 	for _, ch := range s {
// 		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'h' {
// 			return "an"
// 		} else {
// 			return "a"
// 		}
// 	}
// 	return ""
// }
// func main() {
// 	fmt.Println(FixAorAn("goat"))
// 	fmt.Println(FixAorAn("orange"))
// 	fmt.Println(FixAorAn("number"))
// 	fmt.Println(FixAorAn("apple"))
// 	fmt.Println(FixAorAn("grudge"))

// }
//6
// package main

// import (
// 	"fmt"
// )

// func IsPunct(s string) bool {
// 	for _, ch := range s {
// 		if ch == ',' || ch == '!' {
// 			return true
// 		}
// 	}
// 	return false

// }
//
//	func main() {
//		fmt.Println(IsPunct(","))
//		fmt.Println(IsPunct("!"))
//		fmt.Println(IsPunct("x"))
//	}
//
// 5
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func FixPuncts(s []string) string {
// 	text := strings.Join(s, "")
// 	return strings.ReplaceAll(text, " ,", ",")
// }
// func main() { // Removes the space after the opening single quote

// 	fmt.Println(FixPuncts([]string{"hello", ",", "world", "!"}))
// }
//4

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func ToUp(s []string, n int) []string {
// 	for i := n; i < len(s); i++ {
// 		s[i] = strings.ToUpper(s[i])
// 	}
// 	return s
// }
// func main() {
// 	fmt.Println(ToUp([]string{"hello", "there", "how", "far"}, 2))
// }
//3

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Cap(s string) string {
// 	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
// }
// func main() {
// 	fmt.Println(Cap("hello"))
// 	fmt.Println(Cap("world"))

// }
//9

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func FixQuotes(s string) string {
// 	s = strings.Trim(s, "'")
// 	s = strings.TrimSpace(s)

// 	return "'" + s + "'"
// }
// func main() {
// 	text := "awesome"
// 	world := "world"

// 	number := FixQuotes(text)
// 	next := FixQuotes(world)

// 		fmt.Printf("%q\n", number)
// 		fmt.Printf("%q\n", next)
// 	}
//check prime numbers
// package main

// import (
// 	"fmt"
// 	"math"
// )

//	func isPrime(n int) bool {
//		if n <= 1 {
//			return false
//		}
//		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
//			if n%i == 0 {
//				return false
//			}
//		}
//		return true
//	}9
//
//	func main() {
//		fmt.Println(isPrime(1))
//		fmt.Println(isPrime(2))
//	}
//
// loop to use when u dont know exactly how many  times to loop
// package main

// import "fmt"

//	func main() {
//		for i := range 3 {
//			fmt.Println(i)
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	attempts := 0
// 	for attempts < 3 {
// 		fmt.Println("Enter password")
// 		attempts++
// 		password := 12345
// 		var num string = ("Enter password")

// 		if num != string(password) {
// 			fmt.Println("error")
// 			fmt.Scan(&num)
// 		} else {
// 			fmt.Println("correct password")
// 		}
// 	}
// }
// an infinte loop

// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Println("1. Start 2. Help 3. Quit")
// 		var choice int
// 		fmt.Scan(&choice)

// 		if choice == 3 {
// 			fmt.Println("Goodbye!")
// 			break
// 		}
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	order := []string{"paid", "pending", "cancelled"}

//		for _, status := range order {
//			if status == "paid" {
//				fmt.Println("ship the order")
//			} else if status == "pending" {
//				fmt.Println("send reminder")
//			} else {
//				fmt.Println("Archive order")
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	day := "monday"

// 	switch day {
// 	case "wednesday":
// 		fmt.Println("Its niddle of the week")
// 	case "Friday":
// 		fmt.Println("Its towrads ending of the week")
// 	case "saturday", "sunday":
// 		fmt.Println("Weekend")
// 	default:
// 		fmt.Println("Get out")
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	today := "moday"

//		if today == "tuesday" {
//			fmt.Println("just the starting the week")
//		} else if today == "wednesday" {
//			fmt.Println("its middle of week")
//		} else if today == "saturday" {
//			fmt.Println("ending of week")
//		} else {
//			fmt.Println("weekend")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	color := "Red"

//		switch color {
//		case "yellow":
//			fmt.Println("Get Ready")
//		case "Red":
//			fmt.Println("Stop")
//		case "Green":
//			fmt.Println("Go!")
//		default:
//			fmt.Println("Invalid Color")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	colors := []string{"red", "green", "yellow", "red", "red", "green"}
// 	count := 0

// 	for _, ch := range colors {

// 		switch ch {
// 		case "red":
// 			fmt.Println("Stop!")
// 			count++
// 		case "green":
// 			fmt.Println("Go!")
// 		case "yellow":
// 			fmt.Println("Get ready!")
// 		default:
// 			fmt.Println("Wrong Traffic light")
// 		}
// 	}
// 	fmt.Printf("Red light hit %d\n", count)

// }
// package main

// import "fmt"

// func main() {
// 	colors := []string{"red", "green", "yellow", "red", "red", "green"}
// 	count := 0

// 	for _, ch := range colors {
// 		switch ch {
// 		case "red":
// 			fmt.Println("Stop!")
// 		case "yellow":
// 			fmt.Println("Get Ready!")
// 		case "green":
// 			fmt.Println("Go!")
// 		}
// 		if ch == "red" {
// 			count++
// 		}
// 	}
// 	fmt.Printf("Red lights hit: %d\n", count)

// }
// package main

// import "fmt"

// func getAction(color string) string {
// 	switch color {
// 	case "red":
// 		return "Stop!"
// 	case "green":
// 		return "Go!"
// 	case "yellow":
// 		return "Get ready"

// 	}
// 	return "Invalid traffic light"
// }

//	func main() {
//		colors := []string{"red", "green", "yellow", "red", "red", "green"}
//		for _, ch := range colors {
//			fmt.Println(getAction(ch))
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func joinPuncts(s []string) string {
//		text := strings.Join(s, "")
//		return strings.ReplaceAll(text, ",", " ,")
//	}
//
//	func main() {
//		fmt.Println(joinPuncts([]string{"hello", ","}))
//	}
//
// 1 hextobin
// 2 bintodec
// 3 cap1stletter
// 4 up2
// 5 Punctuations
// 6 is puncts
// 7 aoran
// 8 fixarticle
// 9 fix quotes
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func upTwo(s []string, n int) []string {
// 	for i := n; i < len(s); i++ {
// 		s[i] = strings.ToUpper(s[i])
// 	}
// 	return s

// }
//
//	func main() {
//		fmt.Println(upTwo([]string{"hello", "word", "how", "far"}, 2))
//	}
// package main

// import "fmt"

// func main() {
// 	var a [5]int

// 	a[4] = 100

//		fmt.Println("go get it at", a)
//		fmt.Println(a[4])
//	}
// package main

// import "fmt"

// func main() {
// 	b := [9]int{1, 2, 4, 5, 6, 7, 8}

//		fmt.Println(len(b))
//	}
// package main

// import "fmt"

// func main() {
// 	var twoD [2][3]int

// 	for i := range 2 {
// 		for j := range 3 {
// 			twoD[i][j] = i + j

// 		}
// 	}
// 	fmt.Println(twoD)
// }
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func CapFirst(s string) string {
// //return strings.ToLower(s[:len(s)-1]) + strings.ToUpper(string(s[len(s)
// w := strings.Fields(s)

// for i := 0; i < len(w); i++ {
// 	w[i] = strings.ToLower(w[i][:len(w[i])-1]) + strings.ToUpper(w[i][len(w[i])-1:])

// }
// return strings.Join(w, " ")
// //return strings.ToLower(s[:len(s)-1]) + strings.ToUpper(string(s[len(s)-1]))

// }
// func main() {
// 	fmt.Println(CapFirst("numpy hello mathar"))

// }
// package main

// import (
// 	"fmt"
// 	"errors"
// )

// func divide(a, b float64) (float64, error) {
// 			if b == 0 {
// 				return 0, errors.New("division by zero")
// 			}
// 			return a / b, nil
// 		}

//	func main() {
//		fmt.Println(divide(12.34, 2.3))
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("can't divide by zero")
// 	}

// 	return a / b, nil
// }

// func main() {
// 	file, err := Divide(10, 2)

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(file, err)
// 	}

// 	text, err := Divide(10, 0)

//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println(text, err)
//		}
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func validateAge(s int) (int, error) {
// 	if s == -1 {
// 		return 0, errors.New("age cannot be negative")

// 	}
// 	if s == 200 {
// 		return 0, errors.New("age is not realistic")

// 	}
// 	return s, nil
// }
// func main() {
// 	file, err := validateAge(25)

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(file, err)
// 	}
// 	text, err := validateAge(-1)

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(text, err)
// 	}
// 	num, err := validateAge(200)

//		if err != nil {
//			fmt.Println("error:", err)
//		} else {
//			fmt.Println(num)
//		}
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// )

// func StrToInt(s string) (int, error) {
// 	file, err := strconv.Atoi(s)

// 	if err != nil {
// 		return 0, errors.New("Invalid number")
// 	}
// 	if file < 0 {
// 		return 0, errors.New("number cant be a negative number")
// 	}
// 	return file, nil
// }

// func main() {
// 	num, err := StrToInt("42")

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(num)
// 	}

// 	_, err = StrToInt("-5")

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(err)
// 	}

// 	num2, err := StrToInt("abc")

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println(num2, err)
// 	}
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func SimpLog(username, password string) (string, error) {

// 	if username != "admin" {
// 		return "", errors.New("user not found")
// 	}
// 	if password != "1234" {
// 		return "", errors.New("incorrect password")
// 	}

// 	return "", nil

// }
// func main() {
// 	file, err := SimpLog("unknown", "1234")

// 	if err != nil {
// 		fmt.Println("error", err)
// 	} else {
// 		fmt.Println(file, "user not found", err)
// 	}

// 	file2, err := SimpLog("admin", "1234")

// 	if err != nil {

// 		fmt.Println("error:", err)

// 	} else {
// 		fmt.Println(file2, err)
// 	}

// 	file3, err := SimpLog("admin", "wrong")

//		if err != nil {
//			fmt.Println("error", err)
//		} else {
//			fmt.Println(file3, "incorrect password", err)
//		}
//	}

// package main

// import (
//     "errors"
//     "fmt"
// )

// func findUser(username string) (int, error) {
//     users := map[string]int{
//         "thomas":   1,
//         "james":    2,
//         "nathaniel": 3,
//         "bassey":   4,
//     }

//     id, exists := users[username]
//     if !exists {
//         return 0, errors.New("user not found")
//     }
//     return id, nil
// }

// func getScore(userID int) (int, error) {
//     scores := map[int]int{
//         1: 95,
//         2: 87,
//         3: 76,
//     }

//     score, exists := scores[userID]
//     if !exists {
//         return 0, errors.New("score not available")
//     }
//     return score, nil
// }

// func LoadUserScore(username string) (string, error) {

//     id, err := findUser(username)
//     if err != nil {
//         return "", fmt.Errorf("loading user score: finding user: %w", err)
//     }

//     score, err := getScore(id)
//     if err != nil {
//         return "", fmt.Errorf("loading user score: getting score: %w", err)
//     }

//     return fmt.Sprintf("%s scored %d", username, score), nil
// }

// func main() {
//     file1, err := LoadUserScore("james")
//     if err != nil {
//         fmt.Println("error:", err)
//     } else {
//         fmt.Println(file1)
//     }

//     file2, err := LoadUserScore("unknown")
//     if err != nil {
//         fmt.Println("error:", err)
//     } else {
//         fmt.Println(file2)
//     }

//     file3, err := LoadUserScore("bassey")
//     if err != nil {
//         fmt.Println("error:", err)
//     } else {
//         fmt.Println(file3)
//     }
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func getGrade(score int) (string, error) {
// 	if score == 101 {
// 		return "", errors.New("score cannot be above 100")
// 	}
// 	if score > 0 {
// 		return "", errors.New("score cannot be negative")
// 	}
// 	return "", nil
// }
// func main() {
// 	_, err := getGrade(95)

// 	if err != nil {
// 		fmt.Println("error", err)
// 	} else {
// 		fmt.Println("A", err)
// 	}

// 	file2, err := getGrade(101)

// 	if err != nil {
// 		fmt.Println("error", err)
// 	} else {
// 		fmt.Println(file2, "", err)
// 	}
// 	file3, err := getGrade(-1)

// 	if err != nil {
// 		fmt.Println("error", err)
// 	} else {
// 		fmt.Println(file3, "", err)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "hello world"

// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(string(s[0]))
// 	fmt.Println(s[0:4])
// }

// package main

// import (
// 	"fmt"
// 	//"strings"
// )

// func Len(s string) int {
// 	//words := strings.Fields(s)

// 	return len(s)

// }
//
//	func main() {
//		fmt.Println(Len("hello, world, how, you, all, doing"))
//	}
// package main

// import "fmt"

// func reverse(s string) string {
// 	result := ""

//		for i := len(s) - 1; i >= 0; i-- {
//			result = result + string(s[i])
//		}
//		return result
//	}
//
//	func main() {
//		fmt.Println(reverse("hello"))
//		fmt.Println(reverse("racecar"))
//		fmt.Println(reverse("thomas"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func countVowel(s string) int {
// 	count := 0

//		for _, ch := range s {
//			if strings.ContainsAny(string(ch), "aeiou,AEIOU") {
//				count++
//			}
//		}
//		return count
//	}
//
//	func main() {
//		fmt.Println(countVowel("hello"))
//		fmt.Println(countVowel("thomas"))
//		fmt.Println(countVowel("hoolabalo"))
//	}
// package main

// import (
// 	"fmt"
// )

// func PalindChecker(s string) bool {
// 	left := 0

// 	right := len(s) - 1

// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++

// 		right--
// 	}
// 	return true
// }
// func main() {
// 	fmt.Println(PalindChecker("rider"))
// 	fmt.Println(PalindChecker("madam"))

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func UpLast(s string) string {

// 	//   return  strings.ToLower(s[:len(s)-1]) + strings.ToUpper(string(s[len(s)-1]))

// 	w := strings.Fields(s)

// 	for i := 0; i < len(w); i++ {

//			w[i] = strings.ToLower(w[i][:len(w[i])-1]) + strings.ToUpper(string(w[i][len(w[i])-1]))
//		}
//		return strings.Join(w, "")
//	}
//
//	func main() {
//		fmt.Println(UpLast("how are you doing today"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func CapLastNFirst(s string) string {
//		return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:len(s)-1])) + strings.ToUpper(string(s[len(s)-1]))
//	}
//
//	func main() {
//		fmt.Println(CapLastNFirst("apple"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func ToUpper(s string) string {
//		return strings.ToUpper(s)
//	}
//
//	func main() {
//		fmt.Println(ToUpper("hello world"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func Lower(s string) string {
//		return strings.ToLower(s)
//	}
//
//	func main() {
//		fmt.Println(Lower("HELLO WORLD"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func capFirst(s string) string {
//		return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:len(s)])
//	}
//
//	func main() {
//		fmt.Println(capFirst("hello world"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Trim(s string) string {
// 	return strings.Title(s)
// }
// func main() {
// 	fmt.Println(Trim("hello world from go"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func TrimSpace(s string) string {
//		return strings.TrimSpace(s)
//	}
//
//	func main() {
//		fmt.Println(TrimSpace(" hello "))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func trimCutsets(s string) string {
//		return strings.Trim(s, "#")
//	}
//
//	func main() {
//		fmt.Println(trimCutsets("###hello###"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		fmt.Println(strings.TrimPrefix("Mr Smith", "Mr"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Dec(s string) string {
// 	w := strings.Split(s, ",")
// 	return strings.Join(w, "|")
// }

// func main() {
// 	fmt.Println(Dec("go, python, rust"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Split(s string) string {
// 	result := strings.Split(s, " ")

// 	return strings.Join(result, ",")
// }

// func main() {
// 	fmt.Println(Split("one,two,three"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Contains(s string) bool {
// 	if strings.Contains(s, "brown") {
// 		return true
// 	}
// 	return false

// }
//
//	func main() {
//		fmt.Println(Contains("the quick brown fox"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func Replace(s string) string {
//		return strings.Replace(s, "cat", "dog", 3)
//	}
//
//	func main() {
//		fmt.Println(Replace("cat cat cat"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Count(s string) int {
// 	w := strings.Count(s, "a")

//		return w
//	}
//
//	func main() {
//		fmt.Println(Count("banana"))
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func Repeat(s string) string {
//		return strings.Repeat(s, 3)
//	}
//
//	func main() {
//		fmt.Println(Repeat("go!"))
//	}
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	name := "Ada"

// 	age := 32

// 	fmt.Println(name, "is", age, "years old")

// }
//package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var b strings.Builder
// 	letters := []string{"h", "e", "l", "l", "o"}

//		for i, ch := range letters {
//			if i > 0 {
//				b.WriteString("-")
//			}
//			b.WriteString(ch)
//		}
//		fmt.Println(b.String())
//	}

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	M := 18

// 	n := strconv.Itoa(M)

// 	fmt.Println(n)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	num := 100

// 	j := fmt.Sprintf("i am %d years old", num)

//		fmt.Println(j)
//	}
//package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num := "42"

// 	file, err := strconv.Atoi(num)

//		if err != nil {
//			fmt.Println("error:", err)
//		}
//		fmt.Println(file, err)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num := 3.14159

// 	file := strconv.FormatFloat(num, 'f', 2, 64)

//		fmt.Println(file)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num := "3.14"

// 	file, err := strconv.ParseFloat(num, 64)

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Println(file, err)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	word := "true"

// 	file, err := strconv.ParseBool(word)

//		if err != nil {
//			fmt.Println("error", err)
//		}
//		fmt.Println(file, err)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	m := true

// 	file := strconv.FormatBool(m)

//		fmt.Println(file)
//	}
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	m := 32

// 	n := "Ada"
// 	f := 1.75

//		j := fmt.Sprintf("%s is %d years old and %.1f meters tall", n, m, f)
//		fmt.Println(j)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	n := "300"

// 	file, err := strconv.ParseUint(n, 16, 64)

//		if err != nil {
//			fmt.Println("error", err)
//		}
//		fmt.Println(file)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	m := "128"

// 	file, err := strconv.ParseInt(m, 10, 64)

// 		if err != nil {
// 			fmt.Println("error", err)
// 		}
// 		fmt.Println(file)
// 	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	n := "-55"

// 	file, err := strconv.ParseInt(n, 10, 64)

//		if err != nil {
//			fmt.Println("error", err)
//		}
//		fmt.Println(file)
//	}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "hello"

	file, err := strconv.Atoi(n)

	if err != nil {
		fmt.Println("not a valid number")
		//return
	}
	fmt.Println(file)
}

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func parseAnPrint(s string, base int) {

// 	file, err := strconv.ParseInt(s, base, 64)

// 	if err != nil {

//			fmt.Printf("could not parse %q in base %d\n", s, base)
//			return
//		}
//		fmt.Printf("parsed %q (based %d) = %d\n", s, file, base)
//	}
//
//	func main() {
//		parseAnPrint("1E", 16)
//		parseAnPrint("FF", 16)
//		parseAnPrint("1010", 2)
//		parseAnPrint("8", 10)
//		parseAnPrint("xyz", 10)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	n := "200"

// 	file, err := strconv.ParseInt(n, 10, 8)

// 	if err != nil {
// 		fmt.Println("error", err)
// 	}
// 	fmt.Println(file)
// }

// Write a function that:

//     Asks the user to input a word.
//     Asks the user to select an operation by entering one of the following:
//         lastChar
//         capitalize
//         deleteIndex

// Operations

// lastChar Returns the last character of the word.

// capitalize Returns the word with every letter capitalized.

// deleteIndex Prompts the user to input an index. Deletes the character at that index and returns the result. If the index is greater than or equal to the length of the word, display an error message and restart from the beginning.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func LastChar(s string) string {
// 	return (string(s[len(s)-1]))
// }

// func Cap(s string) string {
// 	return strings.ToUpper(s)
// }

// func deleteIndex(s string, n int) string {

// 	if n >= len(s) {
// 		fmt.Println("Error, couldnt process")
// 		main()

// 	}
// 	s = s[:n] + s[n+1:]
// 	return s
// }

// func main() {
// 	var word string

// 	var num int

// 	var oper string

// 	fmt.Print("Enter word:")
// 	fmt.Scanln(&word)

// 	fmt.Print("Enter operator between lastChar, Capitalize, deleteIndex")
// 	fmt.Scanln(&oper)

// 	if oper == "lastChar" {
// 		fmt.Println(LastChar(word))
// 	} else if oper == "Capitalize" {
// 		fmt.Println(Cap(word))
// 	} else if oper == "deleteIndex" {
// 		fmt.Println("Enter Index")
// 		fmt.Scan(&num)
// 		fmt.Println(deleteIndex(word, num))

//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func lastChar(s string) string {
// 	return (string(s[len(s)-1]))
// }
// func capitalize(s string) string {
// 	return strings.ToUpper(s)
// }
// func deleteIndex(s string, n int) string {
// 	if n >= len(s) {
// 		fmt.Println("Error: could'nt process")
// 		main()
// 	}
// 	s = s[:n] + s[n+1:]
// 	return s
// }
// func main() {
// 	var word string

// 	var operation string

// 	var num int

// 	fmt.Println("Enter word:")
// 	fmt.Scan(&word)

// 	fmt.Println("Choose an operation: lastChar/capitalize/deleteIndex")
// 	fmt.Scan(&operation)

// 	if operation == "lastChar" {
// 		fmt.Println(lastChar(word))
// 	} else if operation == "capitalize" {
// 		fmt.Println(capitalize(word))
// 	} else if operation == "deleteIndex" {

//			fmt.Println("Enter Index")
//			fmt.Scan(&num)
//			fmt.Println(deleteIndex(word, num))
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Remove(s string) string {
// 	vowels := "aeiou, AEIOU"

//		for _, ch := range vowels {
//			s = strings.ReplaceAll(s, string(ch), "")
//		}
//		return s
//	}
//
//	func main() {
//		fmt.Println(Remove("hello"))
//	}
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(First("hello"))
	fmt.Println(Cap("hello"))
	// fmt.Println(reverse("rider"))
	fmt.Println(deleteVowel("alphabet"))
	fmt.Println(CountLetter("hello"))
	fmt.Println(replaceIndex("hello", 3))
}

func First(s string) string {
	return (string(s[0]))
}

func Cap(s string) string {
	return strings.ToUpper(s)
}

// func reverse(s string) string {

// 	w := ""

//		for i := len(s) - 1; i >= 0; i-- {
//			w = w + string(w[i])
//		}
//		return w
//	}
func deleteVowel(s string) string {
	vowels := "aeiou, AEIOU"

	for _, ch := range vowels {
		s = strings.ReplaceAll(s, string(ch), "")
	}
	return s
}
func CountLetter(s string) int {

	return (int(len(s)))
}
func replaceIndex(s string, n int) string {
	s = s[:n] + s[n+1:]
	return s
}
// func main() {
// 	fmt.Println(punc("Punctuation tests are ... kinda boring ,what do you think ???"))
// }

// 	words := strings.Fields(word)
// 	new := []string{}
// 	for x := 0; x < len(words); x++ {
// 		if len(new) > 0 && strings.ContainsAny(words[x][len(words[x])-3:], ",.!;?") {
// 			new = append(new, words[x][:1], words[x][1:])
// 		} else if len(new) > 0 && strings.ContainsAny(words[x][:1], ",.!;?") {
// 			new = append(new, words[x][:len(words[x])-3], words[x][len(words[x])-3:])
// 		} else {
// 			new = append(new, words[x])
// 		}
// 	}

// 	fixWord := []string{}

// 	for i := 0; i < len(new); i++ {
// 		if len(fixWord) > 0 && strings.ContainsAny(new[i], ",.!;?") {
// 			fixWord[len(fixWord)-1] += new[i]
// 		} else {
// 			fixWord = append(fixWord, new[i])
// 		}
// 	}
// 	return strings.Join(fixWord, " ")
// }
/ package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func extractURLs(input string) []string {

// 	pattern := regexp.MustCompile(`https?://\S+|www\.\S+`)
// 	return pattern.FindAllString(input, -1)
// }

// func main() {
// 	input := `You can find me on https://www.google.com https://www.facebook.com www.twitter.com`
// 	urls := extractURLs(input)
// 	for _, url := range urls {
// 		fmt.Println(url)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	pattern := regexp.MustCompile("Northern")
// 	text := "New York is in the Northern part of the United States."

//		if pattern.MatchString(text) {
//			fmt.Println("Correct region!")
//		} else {
//			fmt.Println("Wrong region!")
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	text := "Apples are delicious, but so are apples, aPpleS, and APPlE."

//		pattern := regexp.MustCompile(`(?i)apple`)
//		appleMatches := pattern.FindAllString(text, -1)
//		fmt.Println("Case-insensitive matches for 'apple':", appleMatches)
//	}
// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	text := "Hello, how are you?"

// 	pattern := regexp.MustCompile(`[^aeiouAEIOU\s]`)
// 	matches := pattern.FindAllString(text, -1)

// 	fmt.Println(matches)
// }

