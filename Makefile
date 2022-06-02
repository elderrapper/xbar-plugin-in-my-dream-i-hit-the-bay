build:
	go build -v -o ./bin/in-my-dream-i-hit-the-bay.5s.cgo github.com/davidhsingyuchen/xbar-plugin-in-my-dream-i-hit-the-bay

run: build
	./bin/in-my-dream-i-hit-the-bay.5s.cgo

.PHONY: run
