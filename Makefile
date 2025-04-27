.PHONY: benchmark test

test:
	go test ./as ./assert ./measuring ./must ./recovers ./types

benchmark:
	go test -bench=. -benchmem ./as