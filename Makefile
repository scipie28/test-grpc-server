PHONY: generate
generate: .generate

PHONY: .generate
.generate:
		mkdir -p pkg/store_v1
		protoc -I api/store_v1 \
			--go_out=plugins=grpc:pkg/store_v1 \
			api/store_v1/store.proto
		mv pkg/store_v1/github.com/scipie28/test-grpc-server/api/store_v1/* pkg/store_v1/
		rm -rf pkg/store_v1/github.com