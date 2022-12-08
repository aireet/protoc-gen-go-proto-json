PB_DIR=./proto
GEN_DIR=./proto/gen

run:
	go install .

	mkdir -p $(GEN_DIR)

	protoc 		\
	-I $(PB_DIR) \
	--go_out=$(GEN_DIR) \
	--go_opt paths=source_relative \
	--go-proto-json_out=$(GEN_DIR) \
	--go-proto-json_opt paths=source_relative \
	$(PB_DIR)/enums.proto $(PB_DIR)/timestamp.proto