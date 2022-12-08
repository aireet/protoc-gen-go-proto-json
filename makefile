SRC_DIR=./proto


run:
	go install .

	protoc 		\
	-I $(SRC_DIR) \
	--go_out=paths=source_relative:$(SRC_DIR) \
	--go-proto-json_out=source_relative:$(SRC_DIR) \
	$(SRC_DIR)/example.proto