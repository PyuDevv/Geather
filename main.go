package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Print("Enter a city: ")
	var city string
	fmt.Scanln(&city)

	if runtime.GOOS == "windows" {
		println("Incompatible with windows.")
	} else {
		out, err := exec.Command("curl", "wttr.in/"+city).Output()

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			output := string(out[:])

			fmt.Println(output)
		}
	}
}
