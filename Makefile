run:
	nodemon --exec "go run" main.go --signal SIGTERM

gen-go-blog:
	@protoc --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                grpc/blog/blog.proto

gen-go-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                grpc/user/user.proto

gen-go-all: gen-go-blog gen-go-user