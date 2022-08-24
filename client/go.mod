module github.com/logto-io/go/client

go 1.19

require (
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	gopkg.in/square/go-jose.v2 v2.6.0
	github.com/logto-io/go/core v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/logto-io/go/core => ../core/
