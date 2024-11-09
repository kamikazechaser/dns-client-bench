.PHONY: bench

bench:
	go test -v -run=none -benchmem -bench=.