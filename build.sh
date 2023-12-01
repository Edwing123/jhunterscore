#

mkdir -p ./dist

# Save current value for GOOS and GOARCH.
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

# Where to store build artifacts.
DIST=./dist

# Build for Linux amd64.
go env -w GOOS=linux
go env -w GOARCH=amd64

go build -o "$DIST/jhunterscore-linux-amd64" $1

# Build for Windows  amd64.
go env -w GOOS=windows
go env -w GOARCH=amd64

go build -o "$DIST/jhunterscore-windows-amd64.exe" $1

# Restore values.
go env -w GOOS=$GOOS
go env -w GOARCH=$GOARCH

