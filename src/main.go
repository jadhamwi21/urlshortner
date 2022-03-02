package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(MAINMENU)
		scanner.Scan()
		url := scanner.Text()
		if shortenedUrl, err := shortenURL(url); err == nil {
			if shortenedUrl != "" {
				fmt.Printf("Your Shortened URL : %v\n", shortenedUrl)
			} else {
				fmt.Printf("Invalid URL\n")
			}
		} else {
			fmt.Printf("An Error Occurred : %v\n", err)
		}
		fmt.Println(SEPERATOR)

	}
}
