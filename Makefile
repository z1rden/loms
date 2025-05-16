# Используется bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

PHONY: .proto-generate
.proto-generate: .bin-proto .vendor-proto  .order-api-generate .stock-api-generate

# https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file
# https://grpc.io/docs/languages/go/quickstart/
.PHONY: .bin-proto
.bin-proto:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6 && \
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.3 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.26.3 && \
	GOBIN=$(LOCAL_BIN) go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31

.vendor-proto: .vendor-rm  vendor-proto/google/protobuf vendor-proto/buf/validate vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options
	go mod tidy

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

# Устанавливается proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания buf/validate для protovalidate
vendor-proto/buf/validate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone proto/protovalidate/buf/validate &&\
		git checkout
		mkdir -p vendor-proto/buf
		mv vendor-proto/tmp/proto/protovalidate/buf vendor-proto/
		rm -rf vendor-proto/tmp

# Устанавливается proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливается proto описания protoc-gen-openapiv2/options
vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

ORDER_API_PROTO_PATH:=api/order
PHONY: .order-api-generate
.order-api-generate:
	rm -rf pkg/${ORDER_API_PROTO_PATH}
	mkdir -p pkg/${ORDER_API_PROTO_PATH}
	protoc \
	-I ${ORDER_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${ORDER_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${ORDER_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${ORDER_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	${ORDER_API_PROTO_PATH}/*.proto
	#--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
    #--openapiv2_out pkg/${ORDER_API_PROTO_PATH} \
    #--openapiv2_opt logtostderr=true \
	


STOCK_API_PROTO_PATH:=api/stock
PHONY: .stock-api-generate
.stock-api-generate:
	rm -rf pkg/${STOCK_API_PROTO_PATH}
	mkdir -p pkg/${STOCK_API_PROTO_PATH}
	protoc \
	-I ${STOCK_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${STOCK_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${STOCK_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${STOCK_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	${STOCK_API_PROTO_PATH}/*.proto
	#--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
    #--openapiv2_out pkg/${STOCK_API_PROTO_PATH} \
    #--openapiv2_opt logtostderr=true \

.PHONY: generate-apis
generate-apis: .stock-api-generate .order-api-generate