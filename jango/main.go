// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("Sample.txt")
// 	if err != nil {
// 		log.Fatal("Bullshit", err)
// 	}
// 	text := []byte("Hello, World!")

// 	file.Write(text)
// 	file.WriteString("this is a string")// write data aas a string
// 	file.Close()

//		//fmt.Println("WOOOHOOOHOOO... ya Genius bro")
//	}
// package main

// import (
// 	"fmt"
// 	"os"

// )

//	func main() {
//		file, err := os.Open("Sample.txt")
//		if err != nil {
//			panic(err)
//		}
//		data := make([]byte, 29)
//		file.Read(data) // read the data from file to byte array
//		fmt.Printf("The file data is %s\n", string(data)) //
//		file.Close()
//	}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// file, err := os.Create("data.txt")
// if err != nil {
// 	log.Fatal(err)
// 	defer file.Close()
// }

// file, err := os.Open("data.txt")
// if err != nil {
// 	log.Fatal("Bullshit mheeen, ya wack", err)
// }
// defer file.Close()

// file, err := os.Open("data.txt")
// if err != nil {
// 	log.Fatal(err)
// 	word := []byte("Hello, go file operation successful")
// 	file.Write(word)
// 	file.Close()
// }

// defer file.Close()
//
//		fmt.Println("you succeeded in creating the file")
//	}
// package main

// import (
// 	"fmt"
// )

// func countLetters(s string) int {
// 	count := 0

// 	for _, r := range s {

// 		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(countLetters("Hello123!"))
// }

// package main

// import (
// 	"fmt"
// )

// func reverse(s string) string {
// 	runes := []rune(s)

// 	left := 0
// 	right := len(runes) - 1

// 	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
// 		runes[left], runes[right] = runes[right], runes[left]
// 	}
// 	left++
// 	right--

// 	return string(runes)
// }

// func main() {

//		fmt.Println(reverse("Rider"))
//		fmt.Println(reverse("Thomas"))
//		fmt.Println(reverse("jack sparrow"))
//		fmt.Println(reverse("go-reloaded"))
//		fmt.Println(reverse("Available"))
//	}
// package main

// import (
// 	"strings"
// 	"fmt"
// )

// func fixSpaces(s rune) string {
//   puncts := ",.!?:;\""
//   return strings.ContainsRune(puncts, r)
// }
// func main() {
// 	fmt.Println(fixSpaces(" Hello "))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func countPunct(s string) int {

// 	count := 0
// 	puncts := ",.?';:!/*"
// 	hk := `"`
// 	num := puncts + hk

// 	for _, i := range s {
// 		if strings.ContainsRune(num, i) {
// 			count++
// 		}

//		}
//		return count
//	}
//
//	func main() {
//		fmt.Println(countPunct("Hello, world!"))
//		fmt.Println(countPunct("|1 * 3"))
//		fmt.Println(countPunct("normally"))
//		fmt.Println(countPunct(":;'/"))
//	}
// package main

// import "fmt"

// func main() {
// 	scores := map[string]int{
// 		"Alice":     90,
// 		"Geraldine": 100,
// 		"Thomas": 150,
// 		"suur": 80,

// 		value, ok := scores["Alice"]
// if ok {
// 			fmt.Println("found:", value)
// } else {
// 			fmt.Println("Not found")
// 		}
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	grades := map[string]map[string]int{

// 		"Alice": {
// 			"Math":    90,
// 			"English": 80,
// 		},

//			"Thomas": {
//				"Math":    60,
//				"English": 80,
//			},
//			"James": {
//				"Math":    70,
//				"English": 60,
//			},
//			"Nathan": {
//				"Math":    100,
//				"English": 60,
//			},
//			"O'rielly": {
//				"Math":    100,
//				"English": 100,
//			},
//		}
//		fmt.Println(grades["Nathan"]["Math"])
//		fmt.Println(grades["James"]["Math"])
//		fmt.Println(grades["Thomas"]["Math"])
//	}
// package main

// import "fmt"

// func CharCount(s string) map[rune]int {

// 	words := make(map[rune]int)
// 	for _, r := range s {
// 		words[r]++

// 	}

// 	return words

// }
//
//	func main() {
//		fmt.Println(CharCount("Banana"))
//	}

package main

import "fmt"

func main() {

	var x uint8 = 13

	fmt.Printf("%b\n", x)

}
