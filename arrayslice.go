package main

import (
	"fmt"
)

/*
Arrays are value types. Changes made inside a function are not
visible to the caller
*/
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

/*
iterating multidimensional arrays
*/
func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

/*
Changes made to a slice inside a function are visible
to the caller
*/
func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

/*
use copy function to copy the contents from one slice to another
helful to optimize memory since the original array can now be garbage collected
*/
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	/*
	array declaration
	*/
	fmt.Println("array declaration")
	var a [3]int //int array with length 3
	fmt.Println(a)

	var b [3]int //int array with length 3
	b[0] = 12    // array index starts at 0
	b[1] = 78
	b[2] = 50
	fmt.Println(b)

	/*
	short hand declaration to create array
	*/
	fmt.Println("\nshort hand declaration")
	c := [3]int{12, 78, 50}
	fmt.Println(c)

	/*
	... syntactic sugar to make the compiler determine the length
	*/
	fmt.Println("\nsyntactic sugar to determine length")
	d := [...]int{12, 78, 50}
	fmt.Println(d)

	/*
		  e := [3]int{5, 78, 8}
		  var f [5]int
			e = f //not possible since [3]int and [5]int are distinct types
	*/

	/*
	arrays are value types
	*/
	fmt.Println("\nArrays are value types")
	g := [...]string{"USA", "China", "India", "Germany", "France"}
	h := g // a copy of a is assigned to b
	g[0] = "Singapore"
	fmt.Println("g is ", g)
	fmt.Println("h is ", h)

	/*
	Arrays are value types. Changes made in a function are not
	visible to the caller
	*/
	fmt.Println("\nChanges made to an array inside a function are not visible to the caller")
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	/*
	length of an array
	*/
	fmt.Println("\nlength of an array")
	i := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(i))

	/*
	iterating an array using for loop
	*/
	fmt.Println("\niterating array using for loop")
	farray := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, farray[i])
	}

	/*
	iterating an array using for range loop
	*/
	fmt.Println("\niterating array using for range loop")
	for i, v := range farray { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
	}

	/*
	2d arrays
	*/
	fmt.Println("\ndeclaring 2d arrays")
	animals := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(animals)
	var company [3][2]string
	company[0][0] = "apple"
	company[0][1] = "samsung"
	company[1][0] = "microsoft"
	company[1][1] = "google"
	company[2][0] = "AT&T"
	company[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(company)

	/*
	slice declaration
	*/
	fmt.Println("\nSlice declaration")
	as := [5]int{76, 77, 78, 79, 80}
	var slice1 []int = as[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(slice1)

	slice2 := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(slice2)

	/*
	modifying a slice, modifies the underlying array
	*/
	fmt.Println("\nmodifying a slice, modifies the underlying array")
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	/*
	When a number of slices share the same underlying array,
	the changes that each one makes will be reflected in the array.
	*/
	fmt.Println("\nWhen slices share the same underlying array, changes each one makes is reflected in the array")
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	/*
	length and capacity of slice
	*/
	fmt.Println("\nlength and capacity of slice")
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

	/*
	a slice can be re-sliced upto its capacity
	*/
	fmt.Println("\n\nreslicing a slice")
	vegarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	vegslice := vegarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(vegslice), cap(vegslice)) //length of is 2 and capacity is 6
	vegslice = vegslice[:cap(fruitslice)]                                    //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(vegslice), "and capacity is", cap(vegslice))

	/*
	declaring a slice using make
	*/
	fmt.Println("\ndeclaring a slice using make")
	mkslice := make([]int, 5, 5)
	fmt.Println(mkslice)

	/*
	appending to slice
	*/
	fmt.Println("\nappending to a slice")
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	/*
	appending to nil slice
	*/
	fmt.Println("\nappending to a nil slice")
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	/*
	append one slice to another
	*/
	fmt.Println("\nappending one slice to another")
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
	
	/*
	changes made to a slice inside a function are visible to the caller
	*/
	fmt.Println("\nchanges made to a slice inside a function are visible to the caller")
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible out

	/*
	multidimensional slices
	*/
	fmt.Println("\nMultidimensional slices")
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	/*
		memory optimization using copy
	*/
	fmt.Println("\nMemory optimization using copy")
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

}
