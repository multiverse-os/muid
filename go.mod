module github.com/multiverse-os/muid

go 1.19

require (
	github.com/multiverse-os/cli v0.1.0
	golang.org/x/crypto v0.14.0
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

require (
	github.com/multiverse-os/banner v0.0.0-20231006133835-80f8c892b073 // indirect
	github.com/multiverse-os/cli/data v0.0.0-00010101000000-000000000000 // indirect
	github.com/multiverse-os/cli/terminal/ansi v0.0.0-00010101000000-000000000000 // indirect
	github.com/multiverse-os/cli/terminal/loading v0.0.0-00010101000000-000000000000 // indirect
	github.com/multiverse-os/cli/terminal/text/banner v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
