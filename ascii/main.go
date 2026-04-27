// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Enter 2 arguments")
// 		return
// 	}

// 	file, err := os.ReadFile("sample.txt")

// 	if err != nil {
// 		fmt.Println("error")
// 		return
// 	}
// 	word := strings.Split(string(file), "\n")

// 	input := os.Args[1]

// 	words := strings.Split(input, "\\n")

// 	for _, ch := range words {

// 		for i := 0; i < 8; i++ {

// 			for _, j := range ch {

// 				position := (int(j-' ') * 9) + i + 1

// 				fmt.Print(word[position])
// 			}
// 			fmt.Println()
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// 	func main() {
// 		if len(os.Args) != 2 {
// 			fmt.Println("Usage")
// 			return
// 		}
// 		input := os.Args[1]
// 		w := strings.Title(strings.ToLower(input))
// 		fmt.Println(w)
// 	}
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Enter 2 Arguments")
// 		return
// 	}
// 	input := os.Args[1]

// 	w := strings.Fields(input)

// 	max := ""

// 	for i := 0; i < len(w); i++ {

// 		if len(w[i]) > len(max) {
// 			max = w[i]
// 		}
// 	}
// 	fmt.Println(max)

// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Enter 2 Arguments")
// 		return
// 	}
// 	input := os.Args[1]
// 	count := make(map[string]int)

//		for _, ch := range input {
//			count[string(ch)]++
//		}
//		for key, value := range count {
//			fmt.Println(key, ":", value)
//		}
//	}
package main

import (
	"fmt"
)

func main() {
	phone := map[string]string{"Thomas": "080-1234", "Gerald": "080-5678", "jane": "080-9999"}
	fmt.Println(phone)
}
