# Define variables for the location of the protoc binary and the output directory
PROTOC_BIN = protoc
OUT_DIR = .

# Define a function to generate Go code from a proto file
define generate_go_code
	$(PROTOC_BIN) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--proto_path=$(1) \
		$(2)
endef

# Define targets to generate Go code for each proto file
.PHONY: gen

gen:
	$(call generate_go_code,./proto/machine,./proto/machine/machine.proto)

mock:
	mockgen github.com/rht6226/grpc-machine/machine MachineClient,Machine_ExecuteServer,Machine_ExecuteClient > mock_machine/machine_mock.go

run-server:
	go run ./server/cmd

run-client:
	go run ./client/cmd

# Test coverage flags
COVERPKG=./...
COVER_DIR=cov
COVERFILE=coverage.out
HTML_FILE=index.html

test:
	rm -rf $(COVER_DIR)
	mkdir -p $(COVER_DIR)
	go test -v -race -coverprofile=$(COVER_DIR)/$(COVERFILE) $(COVERPKG)
	go tool cover -html=$(COVER_DIR)/$(COVERFILE) -o $(COVER_DIR)/$(HTML_FILE)
		open $(COVER_DIR)/$(HTML_FILE)

