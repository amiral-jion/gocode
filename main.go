package main

import (
	"errors"
	"fmt"

	//"github.com/amiraljion/gocode/models"
)
func main()  {
	/*u := models.User{
		ID:        2,
		FirstName: "Amir",
		LastName:  "bengh",
	}
	fmt.Println(u)
	 */

	port := 3000
	_ , isStarted := startWebServer(port, 3)
	fmt.Println(isStarted)
}

func startWebServer(port, numberRetry int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started", port)
	fmt.Println("Number of retry", numberRetry)

	return port, errors.New("something went wrong")
}