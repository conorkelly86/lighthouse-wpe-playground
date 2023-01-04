package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
		outputFlag := "--output-path=./lighthouseReport.html"
		input := flag.String("input", "default", "a string")
		flag.Parse()

		fmt.Println(*input, " is the input")
		url := *input
		// headers := `--extra-headers "{\"X-WPE-No-Cache\":\"no-cache\"}"`
        cmd := exec.Command("lighthouse", url,  outputFlag)
        output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Printf("Error running Lighthouse: %s\n", err)
                return
        }
		file, err := os.Create("lighthouseOutput.html")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = fmt.Fprintf(file, string(output))
	if err != nil {
		panic(err)
	}	
    fmt.Printf("Lighthouse output:\n%s\n", output)
}
