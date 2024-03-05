# baseproject

This project contains the very simple bootstrapping that I do for most of my Go work, namely:
- installing [`lo`](https://github.com/samber/lo), [`zerolog`](https://github.com/rs/zerolog), and [`envconfig`](https://github.com/kelseyhightower/envconfig), three straightforward libraries that simplify working with collections, logging, and configuration 
- creating a base `Config` struct and code to instantiate it as a singleton using `sync.Once`
- configuring `.gitignore` to ignore `.idea`.

The code in `cmd/baseproject/main.go` is just to prevent `go mod tidy` from removing the dependencies listed above.
