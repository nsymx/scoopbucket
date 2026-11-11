## scoopbucket
Reusable code and assets for [Go](https://go.dev) projects.

> [!NOTE]
> - Type `make help` to see usage for `Makefile`
> - In `.golangci.yml`, `gosec` is temporarily disabled

### Requirements
- Operating system : Linux 
- Architecture     : amd64 / x86_64

### Installation
From you project's root directory, run the following:
```bash
go get github.com/nsymx/scoopbucket
```

### Usage
At the momment, the only available package is `termenv`. \
To use it inside your project, add it to your imports:
```go
package main

import (
       "github.com/nsymx/scoopbucket/termenv"
)
```

> This section needs expansion...
