dev:
	go install github.com/giladsher/proto-zod/cmd/protoc-gen-proto-zod && \
		protoc --go_out=. --go_opt=paths=source_relative \
		--proto-zod_out=. --proto-zod_opt=paths=source_relative \
		test.proto
