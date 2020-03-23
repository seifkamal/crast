build-dev:
		go build -o .build/crast -gcflags="-N -l" ./cmd
build:
		go build -o crast ./cmd
