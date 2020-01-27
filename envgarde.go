package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Rule struct {
	Name        string
	Description string
}

func loadRules() ([]Rule, error) {
	config, err := loadRulesFromTextFile()
	if err != nil {
		config, err = loadRulesFromYamlFile()
	}
	return config, err
}

func loadRulesFromTextFile() ([]Rule, error) {
	content, err := ioutil.ReadFile(".envgarde")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	config := make([]Rule, len(lines))
	for i, line := range lines {
		config[i] = Rule{Name: line, Description: ""}
	}

	return config, nil
}

func loadRulesFromYamlFile() ([]Rule, error) {
	yamlFile, err := ioutil.ReadFile(".envgarde.yaml")
	if err != nil {
		return nil, err
	}

	var config []Rule
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

func printRulesDescription(rules []Rule) {
	for _, Rule := range rules {
		fmt.Println(Rule.Name, Rule.Description, "(Required)")
	}
}

func checkEnvironmentVariables(rules []Rule) {
	var passedRules = make([]Rule, 0)
	var failedRules = make([]Rule, 0)

	for _, Rule := range rules {
		_, isSet := os.LookupEnv(Rule.Name)
		status := "(OK)"
		if !isSet {
			failedRules = append(failedRules, Rule)
			status = "(ERROR: Not set)"
		} else {
			passedRules = append(passedRules, Rule)
		}
		fmt.Println(Rule.Name, status)
	}

	if len(failedRules) == 0 {
		fmt.Println("All environment variables set. OK.")
		os.Exit(0)
	} else {
		fmt.Printf("Error: %v envrionment variables were not set.\n", len(failedRules))
		os.Exit(1)
	}
}

func main() {
	rules, err := loadRules()
	if err != nil {
		fmt.Println(err, "No rules file found.")
		os.Exit(1)
	}

	var describeRules = flag.Bool("d", false, "Describe rules")
	flag.Parse()

	if *describeRules {
		printRulesDescription(rules)
	} else {
		checkEnvironmentVariables(rules)
	}
}
