module github.com/helloWorld

go 1.18

require (
	github.com/GreetingPackage v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/example v0.0.0-20220412213650-2e68773dfca0 // indirect
	golang.org/x/text v0.3.3 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/GreetingPackage => ../GreetingPackage
