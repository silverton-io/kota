
# [private]

version := `cat .VERSION`
kota_dir := "./cmd/kota/"
test_profile := "testprofile.out"

# List all commands
default:
    @just --list

# Build a Kota binary
build:
    go build -ldflags="-X main.VERSION={{version}}" -o kota {{kota_dir}}

# Run Kota locally
run:
	go run -ldflags="-X 'main.VERSION=x.x.dev'" {{kota_dir}}

# Run Kota locally in debug mode
debug:
	DEBUG=1 go run -ldflags="-X 'main.VERSION=x.x.dev' -gcflags='all=-N -l'" {{kota_dir}}

# Lint Kota
lint:
	@golangci-lint run --config .golangci.yml

# Run Kota tests
test:
	@go test ./pkg/...

# Run tests against Kota pkg, output test profile, and open profile in browser
test-cover-pkg:
	go test ./pkg/... -v -coverprofile=$(TEST_PROFILE) || true
	go tool cover -html=$(TEST_PROFILE) || true
