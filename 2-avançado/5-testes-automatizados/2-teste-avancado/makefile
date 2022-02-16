test: 
	go test ./...

test-verbose: 
	go test -v ./...

test-cover: 
	go test ./... --cover

test-cover-profile: 
	go test ./... --coverprofile cobertura.txt

test-cover-tool-func: 
	go tool cover --func=cobertura.txt

test-cover-tool-html: 
	go tool cover --html=cobertura.txt
