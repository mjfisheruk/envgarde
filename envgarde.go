package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func loadFile() []string {
	content, err := ioutil.ReadFile(".envgarde")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {

	var docMode = flag.Bool("d", false, "Run in documentation mode")

	flag.Parse()
	var lines = loadFile()

	if *docMode {
		for _, envVarName := range lines {
			fmt.Println(envVarName, "(Required)")
		}
	} else {
		var errorCount = 0
		for _, envVarName := range lines {
			fmt.Println(envVarName, "(Required)")
			_, isSet := os.LookupEnv(envVarName)
			status := "(OK)"
			if !isSet {
				status = "(ERROR: Not set)"
				errorCount++
			}
			fmt.Println(envVarName, status)
		}

		if errorCount == 0 {
			fmt.Println("All envrionment variables set. OK.")
			os.Exit(0)
		} else {
			fmt.Printf("Error: %v envrionment variables were not set.\n", errorCount)
			os.Exit(1)
		}
	}
}
