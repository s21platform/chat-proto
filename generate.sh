protoc  --go_out=./  \
        --go-grpc_out=./ \
        chat.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./chat.proto