.PHONY: generate schema build test lint clean

# Pull latest schema from npm
schema:
	@echo "Pulling latest schema from @caido/schema-proxy..."
	@curl -sL "$$(curl -sL 'https://registry.npmjs.org/@caido/schema-proxy/latest' | python3 -c 'import sys,json; print(json.load(sys.stdin)["dist"]["tarball"])')" \
		| tar -xz --strip-components=1 -C /tmp/caido-schema-update 2>/dev/null || \
		(mkdir -p /tmp/caido-schema-update && curl -sL "$$(curl -sL 'https://registry.npmjs.org/@caido/schema-proxy/latest' | python3 -c 'import sys,json; print(json.load(sys.stdin)["dist"]["tarball"])')" | tar -xz --strip-components=1 -C /tmp/caido-schema-update)
	@cp /tmp/caido-schema-update/schema.graphql graphql/schema.graphql
	@rm -rf /tmp/caido-schema-update
	@echo "Schema updated: graphql/schema.graphql"

# Run genqlient code generation
generate:
	go run github.com/Khan/genqlient

# Build all packages
build:
	go build ./...

# Run tests
test:
	go test -v -race ./...

# Run linter
lint:
	golangci-lint run ./...

# Run go vet
vet:
	go vet ./...

# Full check: generate, build, vet
check: generate build vet

# Clean generated files
clean:
	rm -f graphql/generated.go
