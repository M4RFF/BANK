package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":  "http://google.com",
		"Youtube": "http://youtube.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Youtube"])

	// adding the new key in the map
	websites["Facebook"] = "http://facebook.com"
	fmt.Println(websites)

	// removing the key
	delete(websites, "Google")
	fmt.Println(websites)

}
