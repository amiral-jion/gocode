package main

import "fmt"

/*const (
	first = iota + 10
	second
)
 */


type user struct {
	ID int
	FirstName string
	LastName string
}
func main()  {


	u1 := user {
		ID:        1,
		FirstName: "Amir",
		LastName:  "ben",
	}
	fmt.Println(u1)
	/*
	m := map[string]int{"foo": 34}

	fmt.Println(m["foo"])

	m["foo"] = 55
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)
	 */
	/*
	slice := []int{1, 2, 3}
		fmt.Println(slice)

		slice = append(slice, 4)
		fmt.Println(slice)

		s1 := slice[1:]
		s2 := slice[:2]
		s3 := slice[1:3]
		fmt.Println(s1, s2, s3)
	*/
	/*
	arr := [3]int {1, 2, 3}
	fmt.Println(arr)
 	*/
	/*var arr [3]int
	arr[0] = 1
	fmt.Println(arr)
	 */
	// fmt.Println(first, second)
	/*
	firstName := "Amir"
	ptr := &firstName
	fmt.Println(ptr, *ptr)
	 */
	/*
	var firstName *string = new(string)
	fmt.Println(firstName)
	*firstName = "Amir"
	fmt.Println(*firstName)

	*/

	/*var i int
	i = 45
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	fistName := "Amir"
	fmt.Println(fistName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)

	fmt.Println(r, im)*/
}
