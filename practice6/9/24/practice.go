package main

import "fmt"

func main() {

	// 1) Create a new array that contains three hobbies
	// I have Output (print) that in the command line.
	// Done it great
	hobbies := [3]string{"Basketball", "VideoEditing", "Coding"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	// - The first element (standalone)
	// - The 2nd and 3rd element combined a new list
	// Done it great
	ex2the1stElement := hobbies[0]
	fmt.Println(ex2the1stElement)
	ex2the2ndAnd3rdElement := hobbies[1:3]
	fmt.Println(ex2the2ndAnd3rdElement)

	// 3) Create a slice based on the first element that
	// contains the 1st and 2nd elements.
	// Create that slice in two different ways (i.e. create two slices in the end)
	// Done it Great
	ex3 := hobbies[0:2]
	fmt.Println(ex3)

	// 4) Re-slice the slice from (3) and change it to
	// contains the second and last element of the origin array.
	// Done it Great
	ex3 = ex3[1:3]
	fmt.Println(ex3)

	exersice5()
	exersice7()

}

func exersice5() {
	// 5) Create a "dynamic array" that contains your course
	// goals (at least 2 goals)
	// Done it Great
	goals := []string{"SolveEveryProblemThatIMeet", "SolveTinkoffCourseAfterThisOne"}
	fmt.Println(goals)

	// 6) Set the 2nd goal to a different one AND then
	// add a 3rd goal to that existing array
	// Done it Great
	goals[1] = "PassAnInterview"
	fmt.Println(goals)

	updatedGoals := append(goals, "SolveTinkoffCourseAfterThisOne")
	fmt.Println(updatedGoals)
}

// 7) Create a "Product" struct with title, id, price and create a new
// dynamic list of products (at least 2 products)
// Then add a 3rd product to the existing list of products

type Product struct {
	title string
	id    string
	price float64
}

func exersice7() {
	products := []Product{
		{
			"The First Produt",
			"Coke",
			13.13,
		},
		{
			"The Second Product",
			"Chips",
			14.14,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"The Third product",
		"Bread",
		5.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
