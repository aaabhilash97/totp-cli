all: releaseb

# Create release binary for current Platform
releaseb:
	@ echo "Compiling release binary"
	@ CGO_ENABLED=1 go build -ldflags "-s -w" -o dist/totp main.go;