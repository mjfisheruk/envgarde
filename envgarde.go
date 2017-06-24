package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type envVarRule struct {
	Name        string
	Description string
}

func loadTextFile() ([]envVarRule, error) {
	content, err := ioutil.ReadFile(".envgarde")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	config := make([]envVarRule, len(lines))
	for i, line := range lines {
		config[i] = envVarRule{Name: line, Description: ""}
	}
	return config, nil
}

func loadYamlFile() ([]envVarRule, error) {
	yamlFile, err := ioutil.ReadFile(".envgarde.yaml")
	if err != nil {
		return nil, err
	}

	var config []envVarRule
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

func main() {

	var docMode = flag.Bool("d", false, "Run in documentation mode")

	flag.Parse()
	config, err := loadTextFile()
	if err != nil {
		config, err = loadYamlFile()
		if err != nil {
			fmt.Println(err, "No configuration file found.")
			os.Exit(1)
		}
	}

	if *docMode {
		for _, envVarRule := range config {
			fmt.Println(envVarRule.Name, "(Required)")
		}
	} else {
		var errorCount = 0
		for _, envVarRule := range config {
			_, isSet := os.LookupEnv(envVarRule.Name)
			status := "(OK)"
			if !isSet {
				status = "(ERROR: Not set)"
				errorCount++
			}
			fmt.Println(envVarRule.Name, status)
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
