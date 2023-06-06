
API_DIR =
PROTO_OUT_DIR = pkg/contracts
RESTAURANT_MENU_PROTO_OUT_DIR = $(PROTO_OUT_DIR)/restaurant_menu
CUSTOMER_OFFICE_PROTO_OUT_DIR = $(PROTO_OUT_DIR)/customer_office
STATISTICS_STATISTICS_PROTO_OUT_DIR = $(PROTO_OUT_DIR)/statistics_statistics


.PHONY: gen-proto
gen-proto:
	mkdir -p pkg/contracts/restaurant_menu
	protoc \
		-I api/contracts/restaurant_menu \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./$(RESTAURANT_MENU_PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=./$(RESTAURANT_MENU_PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        --validate_out="lang=go:./$(RESTAURANT_MENU_PROTO_OUT_DIR)" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./$(RESTAURANT_MENU_PROTO_OUT_DIR) --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./$(RESTAURANT_MENU_PROTO_OUT_DIR) \
        --experimental_allow_proto3_optional \
		./api/contracts/restaurant_menu/*.proto

	mkdir -p pkg/contracts/customer_office
	protoc \
		-I api/contracts/customer_office \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./$(CUSTOMER_OFFICE_PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=./$(CUSTOMER_OFFICE_PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        --validate_out="lang=go:./$(CUSTOMER_OFFICE_PROTO_OUT_DIR)" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./$(CUSTOMER_OFFICE_PROTO_OUT_DIR) --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./$(CUSTOMER_OFFICE_PROTO_OUT_DIR) \
        --experimental_allow_proto3_optional \
		./api/contracts/customer_office/*.proto

	mkdir -p pkg/contracts/statistics_statistics
	protoc \
		-I api/contracts/statistics_statistics \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./$(STATISTICS_STATISTICS_PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=./$(STATISTICS_STATISTICS_PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        --validate_out="lang=go:./$(STATISTICS_STATISTICS_PROTO_OUT_DIR)" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./$(STATISTICS_STATISTICS_PROTO_OUT_DIR) --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./$(STATISTICS_STATISTICS_PROTO_OUT_DIR) \
        --experimental_allow_proto3_optional \
		./api/contracts/statistics_statistics/*.proto
