# Envgarde

Envgarde is a command to assert that environment variables are set. It reads required environment variables from an `.envgarde` configuration file in the current working directory, and returns successfully (exit code zero) if all are set. If any environment variables are missing, it returns a non-zero exit code.

## Motivation

As different deployment strategies proliferate, it is becoming increasingly common to require large sets of environment variables to be provided to a process for it to spin up successfully. Omitting variables that should have been required can lead to non-obvious runtime issues. Ideally these should be validated upfront by the process itself. In practice this often doesn't happen.

Envgarde aims to:

* Allow you to easily create a workflow that causes processes and containers to fail fast with useful feedback if required environment variables are missing
* Explicitly document which environment variables are required in a way that can be committed to your projects in source control


## Usage

### Example usage

Contents of `.envgarde` file
```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
```

Attempting to run a Node JS server protected by envgarde without `AWS_SECRET_ACCESS_KEY` set:

```
> export AWS_ACCESS_KEY_ID=ABACDASDAWEDAIQ
> envgarde && node server.js

AWS_ACCESS_KEY_ID (OK)
AWS_SECRET_ACCESS_KEY (ERROR: Not set)
Error: 1 envrionment variables were not set.
```

But with both required environment variables set:

```
> export AWS_ACCESS_KEY_ID=ABACDASDAWEDAIQ
> export AWS_SECRET_ACCESS_KEY=AWDWHUIHHINKIQ
> envgarde && node server.js

AWS_ACCESS_KEY_ID (OK)
AWS_SECRET_ACCESS_KEY (OK)
All environment variables set. OK.
Server listnening on ":8080"...
```

### Configuration file formats

Two configuration file formats are supported; plaintext and YAML.

The plain text format is simple. It is found in the current working directory as `.envgarde`. It require only the name of the environment variable setting once per line, e.g.

```
LOCALE
ENCODING
```

would require both `LOCALE` and `ENCODING` variables to be set.

The YAML file format provides the same functionality, but also allows a description to be provided per-environment variable. This allows the purpose and any restrictions of your environment variables to be documented. A yaml envgarde config file should be named `.envgarde.yaml`. Example:

```yaml
- name: LOCALE
  description: Service LOCALE, which automatically sets correct i18n config
- name: ENCODING
  description: Supported for legacy reasons; may be set to 'UTF-8' or 'ASCII' 
```

Both files should not be present in the same project - but if they are, `.envgarde` takes priority over `.envgarde.yaml`.

### Running in description mode

By passing the `-d` flag, envgarde will just output a description of the currently required environment variables, e.g.

```
> envgarde -d
LOCALE Service LOCALE, which automatically sets correct i18n config (Required)
ENCODING Supported for legacy reasons; may be set to 'UTF-8' or 'ASCII' (Required)
```

Please note that **when using the description flag envgarde will return successfully irrespective of what environment variables are currently set**.

## Compiling

The command is implemented in Go, so you'll need [the Go tools installing](https://golang.org/doc/install#install) to compile:

Running

```
go build envgarde
```

will result in an `envgarde` executable being built in the project directory

## Running the tests

Because almost all the functionality of this command relies on environmental settings unit testing does not provide a lot of value. Instead, the executable itself is exercised via a set of Gherkin feature files running under Ruby-Cucumber. Please ensure you have >= Ruby 2.4.6 installed before running the tests. To run the tests from a fresh checkout, run the following command:

```
go build envgarde && \
    cd test && \
    bundle && \
    bundle exec cucumber
```

For additional configuration options via environment variables, see the `test/features/step_defintions/basic_steps.rb`.
