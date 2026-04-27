// package main

// import "fmt"

// func main() {
// 	func() {
// 		fmt.Println("Hey, I have no name")
// 	}()
// 	greeting := "Hey, there."

// 	secondWay := func(name string) {
// 		greeting = "Hey" + name
// 	}
// 	secondWay("Pavan")
// 	fmt.Println(greeting)
// }
// package main

// import (
//   "fmt"
// )

// func Hello(name string) string {
//   return "This is being returned with no parentheses. " + name
// }

// func MultipleReturns(a int64, b int64) (int64, int64) {
//   return a + b, a - b
// }

// func MultipleReturns2(a int64, b int64) (c int64, d int64) {
//   x, y := a+b, a-b
//   c = x
//   d = y
//   return
// }

// func printResult(f func(int64, int64) (int64, int64), a int64, b int64) {
//   c, d := f(a, b)
//   fmt.Println(c, d)
// }

// func returnFunc() func() {
//   return func() {
//     fmt.Println("This is a function returned by another function.")
//   }
// }

// func main() {
//   returnFunc()()
//   printResult(MultipleReturns, 10, 3)
//   fmt.Println(Hello("nice right"))
//   fmt.Println(MultipleReturns(10, 3))
//   fmt.Println(MultipleReturns2(10, 3))
// }
// package main

// import "fmt"

// func sum(nums ...int) int {
//    total := 0
//     for _, n := range nums {
//         total += n
//     }
//     return total
// }

// func main() {
//     fmt.Println(sum(1, 2))
//     fmt.Println(sum(1, 2, 3, 4))
// }
// package main

// import "fmt"

// func Sub(num...int) int {
// 	count := 0

// 	for _, ch := range num {
// 		count += ch
// 	}
// 	return count
// }
// func Sum(num1...[]int)int {
// 	count := 0

// 	for _, ch := range num1 {
// 		total := Sum(num1)
// 	}
// 	return total

// }
// func main() {
// 	fmt.Println(Sub(1, 2, 3, 4))
// 	fmt.Println(Sub(1, 2, 3, 4, 5, 6))
// 	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
// }
// 	package main
// import "fmt"

// func main() {
//     i := 0
// Loop:
//     if i < 3 {
//         fmt.Println(i)
//         i++
//         goto Loop
//     }
//     fmt.Println("Done")
// }

// package main

// import "fmt"

// func main() {
// 	arr1 := [3]int{}

// 	arr2 := [5]float64{}

// 	fmt.Println(arr1, arr2)
// }

// package main

// import "fmt"

// func main() {
// 	text := false

// 	fmt.Printf("%T", text)
// }

// package main

// import "fmt"

// func main() {
// 	 num := [3]int{3, 2, 1}

// 	 fmt.Println(num[2], num[1], num[0])
// }
// multy dimensional arrays
// package main

// import "fmt"

// func main() {
//   // Create an array of three arrays containing two integers
//   var twoDim1 [3][2]int

//   // Accessing the first item of the first array
//   twoDim1[0][0] = 1

//   // Accessing the second item of the second array
//   twoDim1[1][1] = 2

//   // Accessing the first item of the third array
//   twoDim1[2][0] = 3

//   fmt.Println(twoDim1)

//   // Create an array of two arrays containing two floating-point numbers
//   twoDim2 := [2][2]float64{{3.14, 2.72}, {2.1, 3.7}}

//   fmt.Println(twoDim2)
// }
// package main

// import "fmt"

// func main() {
// 	arr1 := [2][7]int{{1, 2, 3, 4, 5, 6, 7}, {4, 5, 6, 7, 8, 9, 1}}
// 	arr2 := [3][5]float64{{1.2, 3.4, 5.6}, {2.3, 4.5, 6.7, 9.2}, {2.3, 4.3, 5.1}}
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// package main

// import "fmt"

// type Pizza struct {
//   Name        string
//   Size        string
//   Toppings    []string
//   IsDelicious bool
// }
// func pizzaStyle(p Pizza) string {
//   return p.Name + " pizza is a " + p.Size + " pizza with toppings of " + fmt.Sprint(p.Toppings)
// }

// func main() {
//   myPizza := Pizza{
//     Name:        "Margherita",
//     Size:        "medium",
//     Toppings:    []string{"tomatoes", "mozzarella", "basil"},
//     IsDelicious: true,
//   }
//   fmt.Println(pizzaStyle(myPizza))
// }

// struct in another struct
// package main

// import "fmt"

// type Pizza struct {
//   Name        string
//   Size        string
//   Toppings    []string
//   IsDelicious bool
// }

// func pizzaStyle(p Pizza) string {
//   return p.Name + " pizza is a " + p.Size + " pizza with toppings of: " + fmt.Sprint(p.Toppings)
// }

// type Restaurant struct {
//   Name      string
//   Rating    int
//   PizzaMenu []Pizza
// }

// func restaurantInfo(r Restaurant) string {
//   return r.Name + " has a rating of " + fmt.Sprint(r.Rating) + " and serves the following pizzas: " + fmt.Sprint(r.PizzaMenu)
// }

// func main() {
//   myPizza := Pizza{
//     Name:        "Margherita",
//     Size:        "medium",
//     Toppings:    []string{"tomatoes", "mozzarella", "basil"},
//     IsDelicious: true,
//   }

//   myRestaurant := Restaurant{
//     Name:      "Pizzeria del Corso",
//     Rating:    4,
//     PizzaMenu: []Pizza{myPizza},
//   }

//   fmt.Println(pizzaStyle(myPizza))
//   fmt.Println(restaurantInfo(myRestaurant))
// }

// package main

// import "fmt"

// 	type Address struct {
// 		Street string
// 		City string
// 		Counttry string

// 	}

// 	func address(b Address) string {
// 		return b.Street + ", " + b.City
// 	}
// func main() {
// 	Street := "Gyado Junction"
// 	City := "Gboko"

// fmt.Printf("i am thomas and i live in %s at %s\n", Street, City)
// }
// package main

// import "fmt"

// type Address struct {
//   Street  string
//   City    string
//   Country string
// }

// func (a Address) printAddress() string {
//   return a.Street + ", " + a.City
// }

// type Restaurant struct {
//   Name      string
//   Rating    int
//   PizzaMenu []Pizza
//   Address
// }

// func restaurantInfo(r Restaurant) string {
//   return r.Name + ", located at " + r.printAddress() + ", " + r.Country + " has a rating of " + fmt.Sprint(r.Rating) + " and serves the following pizzas: " + fmt.Sprint(r.PizzaMenu)
// }

// func main() {
//   myPizza := Pizza{
//     Name:        "Margherita",
//     Size:        "medium",
//     Toppings:    []string{"tomatoes", "mozzarella", "basil"},
//     IsDelicious: true,
//   }

//   myAddress := Address{
//     Street: "1st Avenue",
//     City:   "New York",
//     Country: "USA",
//   }

//   myRestaurant := Restaurant{
//     Name:      "Pizzeria del Corso",
//     Rating:    4,
//     PizzaMenu: []Pizza{myPizza},
//     Address:   myAddress,
//   }

//   fmt.Println(pizzaStyle(myPizza))
//   fmt.Println(restaurantInfo(myRestaurant))
// }
// package main

// import "fmt"

// func main() {
// 	for i:= 1; i <= 10; i++ {
// 		for j := 0; j <= 10; j++{
// 			fmt.Println(i, "X", j, "=", i*j)

// 		}
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	//var complex complex64 = complex(1, 2)

// 	complex := 1 +2i
// 	fmt.Println(complex)
// }

// package main

// import (
//     "fmt"
//     "net/http"
// )

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello, World!")
//     })

//     fmt.Println("Server started on :8080")
//     http.ListenAndServe(":8080", nil)
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Person struct {
//     Name    string `json:"name"`
//     Age     int    `json:"age"`
//     Address string `json:"address"`
// }

// func main() {
//     // Creating a JSON object
//     person := Person{Name: "Thomas", Age: 25, Address: "123 Ring Road, GRA GBOKO"}

//     // Encoding the object to JSON
//     jsonData, err := json.Marshal(person)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println(string(jsonData))

//     // Decoding JSON to an object
//     var decodedPerson Person
//     if err := json.Unmarshal(jsonData, &decodedPerson); err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("Decoded:", decodedPerson)
// }
// package main

// import "fmt"

// func  TwoReturn(num1 int64, num2 int64) (a int64, b int64) {
// 	x, y := num1+num2, num1-num2
// 	a = y
// 	b = x
// 	return
// }
// func main() {
// 	fmt.Println(TwoReturn(3, 7))
// }

// package main

// import (
// 	"math"
//   "fmt"
// )

// // GenericMap is a versatile function that applies another function to each item in a list.
// func GenericMap[T any, U any](list []T, f func(T) U) []U {
//   result := make([]U, len(list))
//   for i, item := range list {
//     result[i] = f(item)
//   }
//   return result
// }

// func main() {
//   // Mapping integers to their squares
//   intList := []int{1, 2, 3, 4, 5}
//   squaredList := GenericMap(intList, func(x int) int {
//     return x * x
//   })
//   fmt.Println("Squared integers:", squaredList)

//   // Mapping strings to their lengths
//   stringList := []string{"apple", "banana", "cherry"}
//   lengthList := GenericMap(stringList, func(s string) int {
//     return len(s)
//   })
//   fmt.Println("String lengths:", lengthList)

//   // Mapping floats to their square roots
//   floatList := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
//   sqrtList := GenericMap(floatList, func(x float64) float64 {
//     return math.Sqrt(x)
//   })
//   fmt.Println("Square roots:", sqrtList)
// }
//goroutine afunction and method that runs concurrently in a go program
// package main
// import (
//   "fmt"
//   "time"
// )

// func myGoroutine() {
//   fmt.Println("This is my first goroutine")
// }

// func main() {
//   go myGoroutine()
//   time.Sleep(10 * time.Second)
//   fmt.Println("This is the main function")
// }
// package main
// import (
//   "fmt"
//   "time"
// )

// // define a function
// func print(text string) {
//   fmt.Println(text)
// }

// func main() {
//   // call goroutine
//   go print("This text is from the goroutine.")
//     time.Sleep(time.Second * 4)

//   // call function
//   print("This text is from the main function.")
// }

// package main

// import "fmt"

// func Bio(s string) string {
// 	type bio struct {
// 		name  string
// 		Age  int
// 	}
// fmt.Sprintf("name", "" + "", "Age")
// return ""
// }
// func main() {
// 	name := "Thomas"
//     Age := 25

// 	fmt.Println("Hello", name, "you will be", Age, "next year")
// }

// package main

// import "fmt"

// func main() {
// 	var pin int
// 	var count int

// 	for {
// 		fmt.Println("Enter Pin:")
// 		fmt.Scan(&pin)
// 		if pin == 1234 {
// 			fmt.Println("Correct Pin Process withdrawal")
// 			break
// 		} else {
// 			fmt.Println("Wrong pin Try again mate")
// 			continue
// 		}

// 	}
// }
// package main
// import "fmt"

// func main() {
//   for i := 0; i < 20; i++ {
//     if i % 2 == 0 {
//       continue
//     }

//     if i == 10 {
//       break
//     }

//     fmt.Println(i)
//   }
// }
// creating maps

//maps are buitl-in datastructures that store undordered keys(value pairs)

// package main

// import "fmt"

// func main() {
// 	//create a map using the make()
// 	gradebook1 := make(map[string]float32)
//      // create a map using the map literal
// 	gradebook2 := map[string]float32{}
//      // initialilize a map using a map literal
// 	gradebook3 := map[string]float32{"thomas": 85.8, "jasmine": 67.9}

// 	fmt.Println(gradebook1)
// 	fmt.Println(gradebook2)
// 	fmt.Println(gradebook3)
// }

// Accessing a value data type
// package main

// import "fmt"

// func main() {
// 	Names := map[string]int {"Thomas": 25, "Jasmine": 23, "Gerald": 24}
// 	fmt.Println(Names["Jasmine"])
// }

// Adding new items to a map

// package main

// import "fmt"

// func main() {
// 	Country := map[string]int{}

// 	Country["Nigeria"] = 3

// 	Country["Usa"] = 4
// 	Country["Germany"] = 2
// 	Country["Afghan"] = 1

// 	fmt.Println(Country)
// }

//removing items from a map

// package main

// import "fmt"

// func main() {
// 	Country := map[string]int{"Nigeria": 12, "Japan": 34, "Spain": 54, "Pakistan":25}

// 	delete(Country, "Nigeria")
// 	fmt.Println(Country)
// }
// iterating over  map

// package main

// import "fmt"

// func main() {
// 	Country := map[string]int {"Nigeria": 23, "Japan":13, "Germany":25}
// 	for j, i := range Country{
// 		//fmt.Println(j, i)
// 		fmt.Printf("(%s, %d)\n", j, i)
// 	}
// }
//math functions
// package main

// import (
// 	"math"
// 	"fmt"
// )
// func main() {
// 	number := math.Round(2.9)
// 	fmt.Printf("%.1f\n", number)
// }

// .Hypot()
//     Returns the square root of the sum of squares of two numbers, avoiding overflow and underflow issues.
// Abs()
//     Returns the magnitude of a numeric value by removing any negative sign.
// Acos()
//     Returns the inverse of the cosine value of a number.
// Acosh()
//     Returns the inverse hyperbolic cosine of a number.
// Asin()
//     Returns the inverse sine of a number.
// Asinh()
//     Computes the inverse hyperbolic sine (arsinh) of a numeric value.
// Atan()
//     Returns the arctangent of the given value.
// Atan2()
//     Returns the arctangent value of the x/y value.
// Atanh()
//     Returns the inverse hyperbolic tangent of the given value.
// Cbrt()
//     Returns the cube root of a given number of type float64.
// Ceil()
//     Returns a given decimal number rounded up to the next highest integer.
// Copysign()
//     Returns a value with the magnitude and sign of the given arguments.
// Cos()
//     Returns the cosine of the given angle.
// Cosh()
//     Returns the hyperbolic cosine of a given value.
// Dim()
//     Returns the maximum of the difference between two arguments.
// Exp()
//     Returns the value of e raised to the power of the parameter x.
// Exp2()
//     Returns the value of a base-2 exponential of a given number.
// Expm1()
//     Returns the value of e raised to the power of the given parameter minus 1.
// Floor()
//     Returns the given decimal number rounded down to the nearest whole number.
// Log()
//     Returns the natural logarithm of a given number.
// Log10()
//     Returns the base-10 logarithm of a given number.
// Log2()
//     Returns the base-2 logarithm of a given number.
// Max()
//     Returns the maximum value of two specified numbers.
// Min()
//     Returns the minimum value of two specified numbers.
// Mod()
//     Returns the floating-point remainder of dividing x by y.
// Modf()
//     Returns the integer and fractional part of a floating-point number.
// Pow()
//     Calculates exponential results by raising one number to the power of another.
// Remainder()
//     Returns the remainder of the division of two given values.
// Round()
//     Rounds the given number to the nearest integer.
// RoundToEven()
//     Rounds a floating-point number to the nearest even integer.
// Sin()
//     Used to calculate the sine of an angle.
// Sinh()
//     Returns the hyperbolic sine of the given number.
// Sqrt()
//     Returns the square root of a given number.
// Tan()
//     Returns the tangent of the given angle.
// Tanh()
//     Computes the hyperbolic tangent function for a given numeric input.
// Trunc()
//     Returns the integer value of the given number.
// package main

// import (
// 	"fmt"
// 	"math"
// )
// func main() {
// 	number := math.Cbrt(729)
// 	fmt.Printf("(%1.f)\n", number)
// 	fmt.Println(number)
// }
//accessing values through a pointer

// package main

// import "fmt"

// func main() {
// 	var num int = 10

// 	var ptr *int = &num

// 	fmt.Println(ptr)
// }

//modifying values through a pointer

// package main

// import "fmt"

// func main() {
// 	var num int = 10

// 	var ptr *int = &num

// 	*ptr = 20
// 	fmt.Println(ptr)
// }
// checking for null pointers
// package main

// import "fmt"

// func main() {
// 	var ptr *int
// 	fmt.Println(ptr)
// }

// package main

// import "fmt"

// func main() {
// 	var num int = 10

// 	var ptr *int = &num

// 	fmt.Println(num)
// 	fmt.Println(ptr)
// 	fmt.Println(*ptr)
// 	*ptr = 20
// 	fmt.Println(ptr)
// 	fmt.Println(num)
// }

//slice

// package main

// import "fmt"

// func main() {

//   // Declare and initialize a slice in one line
//   n := []int{1, 2, 3, 4, 5}
//   fmt.Println("New Slice:", n, "Length:", len(n), "Capacity:", cap(n), "\n")

//   // Create a slice from an existing array
//   a := [7]int{1, 2, 3, 4, 5, 6, 7}
//   s := a[2:5]
//   fmt.Println("Array:", a, "Length:", len(a), "Capacity:", cap(a))
//   fmt.Println("Slice from Array:", s, "Length:", len(s), "Capacity:", cap(s), "\n")

//   // Create an empty slice, setting length and capacity
//   m := make([]int, 3, 8)
//   fmt.Println("Empty Slice:", m, "Length:", len(m), "Capacity:", cap(m), "\n")

// }
//strings: this are read only sequence of bytes we have the raw string literals(``) and the interpreted
//string literals("") the raw ignores escape sequence while the interpreted doesnt  eg
// package main

// import "fmt"

// func main() {
// 	var name string = "thomas\njasmine\ncosmas\njohn"
// 	var name1 string = `thomas
// jasmine
// cosmas
// john	`
// 	var name2 string = `"thomas"`
// 	var name3 string = "thomas\""

// 	fmt.Println(name)
// 	fmt.Println(name1)
// 	fmt.Println(name2)
// 	fmt.Println(name3)
// }

// package main

// import "fmt"

// func main() {
// 	name := "there is no \"I\" in Team"
// 	fmt.Println(name)
// }

// package main
// import (
//   "fmt"
//   "strings"
// )
// func main() {
//   name := "Codecademy"
//   find := "d"
//   fmt.Println(find, "is in", name, "at index", strings.Index(name,find))
// }
// package main

// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	var name string = "codecademy"
// 	find := "d"

// 	fmt.Println("d in", name, "is on index", strings.Index(name, find))
// }

// Strings

// Compare()
//     Compares two strings in lexicographical order.
// Contains()
//     Returns a boolean value indicating whether a given substring is present or not in a given string.
// Count()
//     Returns the number of times a substring occurs in a given string.
// Cut()
//     Slices a string around a separator.
// CutPrefix()
//     Returns a given string with the specified prefix removed and a boolean value confirming if the prefix was present.
// CutSuffix()
//     Returns a given string with the specified suffix removed and a boolean value confirming if the suffix was present.
// Fields()
//     Splits a string into substrings based on whitespace and returns a slice of the substrings.
// HasPrefix()
//     Returns a boolean value indicating whether a given string begins with a given prefix.
// HasSuffix()
//     Checks if a given suffix exists on the specified string. Returns true if the string has the given suffix, else it returns false.
// Index()
//     Returns the index value of the first occurrence of a substring in the original string.
// Join()
//     Returns the concatenated elements of a string slice into a single string.
// LastIndex()
//     Returns the index value of the last occurrence of a substring in the original string.
// Map()
//     Returns a copy of a provided string with its characters modified according to a given mapping function.
// Replace()
//     Replaces occurrences of a specified substring within a given string with another substring.
// ToLower()
//     Converts a string to lowercase.
// ToUpper()
//     Returns a string in all uppercase.
// Trim()
//     Removes leading and trailing characters from a string based on the given input.

//struct in go this is a collection of different data types in a variable
// assigning value to struct using the . notation
// package main

// import "fmt"

// 	type Car struct {
// 		brand string
// 		year int
// 		mileage float64
// 	}
//  func main() {
// 	myCar := Car{}
// 	myCar.brand = "BMW"
// 	myCar.year = 2023
// 	myCar.mileage = 200021
//  }

// assigning a value using the default value
// package main

// import "fmt"

// 	type Car struct {
// 		brand string
// 		year int
// 		mileage float64
// 	}

// func NewCar(brand string, year int, mileage float64) *Car {
//     return &Car{brand, year, mileage}
// }

// func main(){
//     car1 := NewCar("Toyota", 2023, 0)

//     fmt.Printf("Brand: %s\n", car1.brand)
//     fmt.Printf("Year: %d\n", car1.year)
//     fmt.Printf("Mileage: %f\n", car1.mileage)
// }
//initializing structs using the var keyword
//this is useful when initializing the fields individually

// package main

// import "fmt"

// type car struct {
// 	Name string
// 	Year int
// 	Rate float64
// }

// func main() {
// 	var grandma car

// 	grandma.Name = "Tesla"
// 	grandma.Year = 2003
// 	grandma.Rate = 65.4

// 	fmt.Println(grandma.Name)
// 	fmt.Println(grandma.Year)
// 	fmt.Println(grandma.Rate)

// }
// multiple variable initializing
// package main

// import "fmt"

//	func main() {
//		name := "aondongu"
//	    for i, j := range name {
//		//x := []byte(name)
//		fmt.Println(i, j)
//		}
//	}
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//filePath := "sample.txt"

	// Open the file located at the given file path
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize a new buffered reader to improve read efficiency
	// reader := bufio.NewReader(file)
	// for {
	// 	char, _, err := reader.ReadRune()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Print(string(char))
	text, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(text))
}

//}
