package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("Enter a city: ")
	var city string
	fmt.Scanln(&city)

	out, err := exec.Command("curl", "wttr.in/"+city).Output()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		output := string(out[:])

		fmt.Println(output)
	}
}
