package main

import "fmt"

const (
	TO_HOUR   = 3600
	TO_MINUTE = 60
)

func main() {
	var seconds int

	fmt.Print("Enter seconds: ")
	fmt.Scan(&seconds)

	fmt.Print(seconds/TO_HOUR, "h ", (seconds%TO_HOUR)/TO_MINUTE, "m ", (seconds%TO_HOUR)%TO_MINUTE, "s")
}
