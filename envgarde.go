package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var docMode = flag.Bool("d", false, "Run in documentation mode")

	flag.Parse()

	file, err := os.Open(".envgarde")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorCount := 0

	for scanner.Scan() {
		envVarName := scanner.Text()
		if *docMode {
			fmt.Println(envVarName, "(Required)")
		} else {
			_, isSet := os.LookupEnv(envVarName)
			status := "(OK)"
			if !isSet {
				status = "(ERROR: Not set)"
				errorCount++
			}
			fmt.Println(envVarName, status)
		}
	}

	if !*docMode {
		if errorCount == 0 {
			fmt.Println("All envrionment variables set. OK.")
			os.Exit(0)
		} else {
			fmt.Printf("Error: %v envrionment variables were not set.\n", errorCount)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
