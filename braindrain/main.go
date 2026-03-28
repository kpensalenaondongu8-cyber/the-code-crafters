// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		for j := 0; j <= 10; j++ {

// 			fmt.Println(i, "x", j, "=", i*j)
// 		}
// 	}
//

// package main

// import "fmt"

// func main() {

// 	maxTrials := 4

// 	for i := 1; i <= maxTrials; i++ {
// 		fmt.Printf("Trial %d/%d - Enter pin: ", i, maxTrials)

// 		var pin int
// 		fmt.Scan(&pin)
// 		if pin == 123 {
// 			fmt.Println("Success! Proceeding to withdrawal...")
// 			break
// 		} else {
// 			fmt.Println("Wrong pin.")
// 		}

//			if i == maxTrials {
//				fmt.Println("Account locked. Too many failed attempts.")
//			}
//		}
//	}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Create the scanner (talking to standard input/keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")

	// 2. This waits for you to type and hit Enter
	if scanner.Scan() {
		// 3. .Text() gives you the string WITHOUT the "Enter" newline
		input := scanner.Text()
		fmt.Println("hello,", input, "how are you doing")
	}
}
