bench:
	go test -bench=. -benchmem -run=^$ . -memprofile memprofile.out -cpuprofile cpuprofile.out

mem:
	go tool pprof -web memprofile.out

cpu:
	go tool pprof -web cpuprofile.out
