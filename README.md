## Arg Parser
Create a function ```ParseArguments``` that get's a string and parses its content into a configuration object.

## Description

## Examples:
``` 
go run arg_parser.go --help -> {help: true}
go run arg_parser.go --foo bar -> {foo: "bar"}
go run arg_parser.go --foo 1 -> {foo: 1}
go run arg_parser.go --foo bar --foo rab -> {foo: ["bar", "rab"]}
```