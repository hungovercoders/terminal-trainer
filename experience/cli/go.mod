module terminal-trainer-cli

go 1.24.5

require (
	github.com/spf13/cobra v1.10.1
	terminal-trainer/behaviour v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

replace terminal-trainer/behaviour => ../../behaviour
