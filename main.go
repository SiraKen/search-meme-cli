package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func main() {

	cErr := color.New(color.FgWhite, color.BgRed, color.Bold)

	if len(os.Args) < 2 {
		cErr.Println("Usage: meme <url>")
		os.Exit(1)
	}

	queryArray := os.Args[1:]

	// Make search query string
	query := ""
	for _, q := range queryArray {
		query += q + " "
	}

	fmt.Printf("Search %s on Know Your Meme...\n", query)

	// Search on Know Your Meme
	// https://knowyourmeme.com/search?q=query
	openBrowser("https://knowyourmeme.com/search?q=" + query)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		os.Exit(1)
	}

}
