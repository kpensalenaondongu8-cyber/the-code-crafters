// package main

// import "fmt"

// func main() {
// 	name := "Thomas"
// 	age := 25

// 	fmt.Println("my name is", name)
// 	fmt.Println("I am",  age,"years old")
// }

// package main

// import "fmt"

// func main() {
// 	city := "Paris"
// 	city = "Tokyo"

// 	fmt.Println(city)
// }
// package main

// import "fmt"

// func main() {
// 	var width int = 10
// 	var height int = 5
// 	area := width * height
// 	fmt.Println(area)
// }
// package main

// import "fmt"

// func main() {
// 	a := 10
// 	b := 20

// 	a, b = b, a
// 	fmt.Println(a, b)
// }

// package main

// import "fmt"

// func main() {

// 	var userName string 

// 	fmt.Println("Enter your name:")
// 	fmt.Scan(&userName)

// 	fmt.Println("Hello", userName)
// }

// package main

// import (
// 	"strconv"
// 	"fmt"
// )

// func main() {
// 	num := "25"
//     num2 := 2
//  file, err := strconv.Atoi(num)

//  if err != nil {
// 	fmt.Println("Error could'nt convert")
//  }
//   num3 := file * num2
//   fmt.Println(num3)



// }

// package main

// import "fmt"

// func main() {
// 	counter := [5]int{1, 2, 3, 4, 5} 

// 	for i := 1; i <= len(counter); i++ {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	num1 := 10
// 	num2 := 5
// 	operator := "+"

// 	if operator == "+" {
// 		fmt.Println(num1 + num2)
// 	}


// }

// package main

// import "fmt"

// func main() {
// 	word := "hello"

// 	replace := word[0:1]

// 	fmt.Println(replace)
// }




// package main

// import "fmt"

// func main() {
// 	numbers := []int{10,20,30}
// 	numbers[1] = 99
// 	fmt.Println(numbers)
// }

// package main

// import "fmt"

// func main() {
// 	numbers := []int{1,2,3}

// 	numbers = append(numbers, 4)

// 	fmt.Println(numbers)
// }


// package main

// import "fmt"

// func main() {
// 	numbers := []int{10,20,30}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }


// package main

// import "fmt"

// func main() {
// 	numbers := []int{10,20,30}

// 	for i, value := range numbers {
// 		fmt.Println(i,
// 			 value)
// 	}
// }


// package main

// import (
// 	"strings"
// 	"fmt"
// )
// func main() {
// 	sentence := "go is fun, but hard to learn shaa"

// 	words := strings.Fields(sentence)

// 	fmt.Println(words)
// }


//removing an element in a slice
// package main

// import "fmt"

// func main() {
// 	numbers := []int{10,20,30,40}

// 	numbers = append(numbers[:1], numbers[2:]...)
// 	fmt.Println(numbers)
// }

//slicing a slice

// package main

// import "fmt"

// func main() {
// 	numbers := []int{10,20,30,40}
// 	fmt.Println(numbers[:1])
// }


// package main

// import "fmt"

// func main() {
// 	num := []int{5, 10, 15, 20}

// 	fmt.Println(num[2])
// }

// package main

// import "fmt"

// func main() {
// 	word := []int{}

// 	sep := append(word, 1,2,3)
// 	fmt.Println(sep)
// }

// package main

// import (

// 	"fmt"
// )
// func main() {
// 	word := []string{"Go", "is", "awesome"}
// for i, value := range word {
//  		fmt.Println(i, value)
	 
// }
// }

// package main

// import "fmt"

// func main() {
// 	num1 := []int{10,20,30,40,50}
// 	fmt.Println(num1[1:4])
// }


// package main

// import "fmt"

// func main() {
// 	words := []string{"i", "love", "go"}

// 	sept := append(words, "really")

// 	fmt.Println(sept)
// }
// package main

// import "fmt"

// func main() {
// 	numbers := []int{10, 20, 30, 40}

// 	for i := 1; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }


// package main


// import "fmt"

// func main() {
// 	numbers := []int{10,20,30}

// 	for i, value := range numbers {
// 		fmt.Println(i, value)
// 	}
// }



//if u only need the value
// package main

// import "fmt"

// func main() {
// 	numbers := []int{10, 20, 30, 40}

// 	for _, i := range numbers {
// 		fmt.Println(i)
// 	}
// }

// when u need only the index
// package main

// import "fmt"

// func main() {
// 	numbers := []int{10, 20, 30, 40, 50}
//     for i := range numbers {
// 		fmt.Println(i)
// 	}

// }


//modifying elements inside a loop

// package main

// import "fmt"

// func main() {
// 	numbers := []int{10, 20, 30, 40}

// 	for i := range numbers {
// 		numbers[i] = numbers[i] * 2
// 	}
// 	fmt.Println(numbers)
// }

// looping over strings slice

// package main

// import "fmt"

// func main() {
// 	words := []string{"Go", "is", "fun"}

// 	for _, word := range words {
// 		fmt.Println( word)
// 	}
// }


//transform words while looping
// package main

// import (
// 	"strings"
// 	"fmt"
// )
// func main() {

// 	words := []string{"go", "is", "fun"}

// 	for  i := range words {
// 		words[i] = strings.ToUpper(words[i])
// 	}
// 	fmt.Println( words)
// }


// package main

// import "fmt"

// func main() {
// 	numbers := []int{5, 10, 15, 20}

// 	for _, r := range numbers {
// 		fmt.Println(r)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	words := []string{"Go", "is", "awesome"}

// 	for i, value := range words {
// 		fmt.Println(i, value)
// 	}
// }

// package main

// import (
// 	"strings"
// 	"fmt"
// )
// func main() {
// 	words := []string{"go", "is", "fun"}
// 	for i := range words {
// 	words[i] = strings.ToUpper(words[i]) 	
// 	}

// 		fmt.Println(words)

// }

package main

import "fmt"

func main() {
	numbers := []int{2,4,6,8}

	for i := range numbers {
		numbers[i] = numbers[i] * 2
			fmt.Println(numbers)

	}
}