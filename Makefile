build:
	go build -v -o ./bin/in-my-dream-i-hit-the-bay.5s.cgo github.com/davidhsingyuchen/xbar-plugin-in-my-dream-i-hit-the-bay
    # Make it easier to test in a macOS guest machine via a shared folder with the host non-macOS machine.
    # In other words, I can develop on the host machine,
    # and the only thing I need to do in the guest machine is to
    # copy the updated binary to the xbar plugins directory.
	GOOS=darwin go build -v -o ./bin/darwin-in-my-dream-i-hit-the-bay.5s.cgo github.com/davidhsingyuchen/xbar-plugin-in-my-dream-i-hit-the-bay

run: build
	./bin/in-my-dream-i-hit-the-bay.5s.cgo

.PHONY: run
