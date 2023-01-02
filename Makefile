# Set the name of the compiled binary
BINARY_NAME=sondar

# Set the installation prefix (e.g., /usr/local)
PREFIX=/usr/local

# Set the output directory for the compiled binary
OUTPUT_DIR=./bin

# Set the build type to "development" by default
BUILD_TYPE=development

# Set the build flags to debug flags and symbols by default
GO_FLAGS=-gcflags='all=-N -l'
GOFLAGS=-ldflags='-N -l'

# Detect and set the target operating system and target architecture
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

build:
	# Create the output directory if it doesn't exist
	mkdir -p $(OUTPUT_DIR)

	# Print a message indicating the binary being built
	printf "%s\n" "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH) ($(BUILD_TYPE) build)"

	# Build the binary with the specified flags
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_FLAGS) $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)

production:
	# Set the build type to "production"
	$(eval BUILD_TYPE=production)

	# Set the build flags to remove symbol and debugging information to reduce the file size
	$(eval GO_FLAGS=-ldflags=-s -ldflags=-w)

	# Set the build flags to enable aggressive optimization, optimize the code for the host machine's microarchitecture,
	# and use pipes instead of temporary files when invoking the assembler and linker
	$(eval GOFLAGS=-gcflags=-O3 -march=native -pipe)

	# Invoke the build target to compile the binary with the updated flags
	$(MAKE) build

development:
	# Invoke the build target to compile the binary with the default development flags
	$(MAKE) build

test:
	# Run the Go tests
	go test

clean:
	# Remove the compiled binary and any intermediate build files
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)

format:
	# Format the Go code using gofmt
	gofmt -s -w .

lint:
	# Run golint to check for common coding mistakes and style issues
	golint -set_exit_status

run:
	# Execute the compiled binary
	$(OUTPUT_DIR)/$(BINARY_NAME)

install:
	# Install the compiled binary to /usr/local/bin
	cp $(OUTPUT_DIR)/$(BINARY_NAME) $(PREFIX)/bin
