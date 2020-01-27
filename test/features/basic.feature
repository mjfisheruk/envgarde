Feature: basics

  Scenario: running with the required environment variables set
    Given there is a basic config file
    And the HELLO environment variable is set to WORLD
    When envgarde is run without arguments
    Then it's exit code is 0

  Scenario: running without the required environment variable set
    Given there is a basic config file
    And the HELLO environment variable is not set
    When envgarde is run without arguments
    Then it's exit code is 1

  Scenario: running in describe mode
    Given there is a basic config file
    When envgarde is run with the arguments '-d'
    Then it's exit code is 0
    And it's output includes 'HELLO'
  
  Scenario: running the program with no config file
    Given there is no .envgarde file
    When envgarde is run without arguments
    Then it's exit code is 1

  Scenario: running the program with a yaml config file
    Given there is a basic.yaml config file
    And the HELLO environment variable is set to WORLD
    When envgarde is run without arguments
    Then it's exit code is 0

  Scenario: running in describe mode with a yaml config file
    Given there is a basic.yaml config file
    When envgarde is run with the arguments '-d'
    Then it's exit code is 0
    And it's output includes 'You always have to say hello'
