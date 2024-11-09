.PHONY: bench

bench:
	go test -v -run=none -benchtime 5s -benchmem -bench=. -count=5