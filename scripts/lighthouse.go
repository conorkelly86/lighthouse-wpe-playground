package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
		// outputFlag := "--output-path=./output/lighthouseReport.html"
		input := flag.String("input", "default", "a string")
		flag.Parse()

		fmt.Println(*input, " is the input")
		url := *input
		flags := []string{url, "--only-categories=performance", "--view", "--chrome-flags=\"--headless\"", "--preset=\"desktop\""}
		// headers := `--extra-headers "{\"X-WPE-No-Cache\":\"no-cache\"}"`
        cmd := exec.Command("lighthouse", flags...)
        output, err := cmd.CombinedOutput()
		fmt.Printf("Lighthouse output:\n%s\n", output)
		if err != nil {
			panic(err)
		}
		
	if err != nil {
		panic(err)
	}	
    fmt.Printf("Lighthouse output:\n%s\n", output)
}
