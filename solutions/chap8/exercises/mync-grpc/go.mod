module practical/chap8/exercises/mync-grpc

require (
	practical/chap8/user-service/service v0.0.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace practical/chap8/user-service/service => ../../user-service/service

go 1.16
