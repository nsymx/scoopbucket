## scoopbucket

### Usage
> [!TIP]
> - Type `make help` to see usage for `Makefile`
> - In `.golangci.yml`, `gosec` is temporarily disabled

Build the binary:
```bash
make build
```

Run the binary:
```bash
./bin/scoopbucket
```

Show version:
```bash
./bin/scoopbucket --version
```

### Misc
Project structure:
```go
./scoopbucket/
├── cmd
│   ├── build
│   │   ├── scoopbucket-v0.0.0-0-linux-amd64
│   │   └── sha512sum.txt
│   └── scoopbucket
│       └── main.go
├── .github
│   ├── dependabot.yml
│   └── workflows
│       ├── main.yml
│       └── release.yml
├── .gitignore
├── .golangci.yml
├── go.mod
├── go.sum
├── internal
│   ├── logger
│   │   └── custom.go
│   └── styles
│       └── lipgloss.go
├── LICENSE
├── linuxenv
│   ├── detect_system.go
│   ├── detect_system_test.go
│   ├── terminal.go
│   └── terminal_test.go
├── Makefile
├── README.md
├── scripts
│   └── go-validate.sh
└── .version
    └── latest
```
