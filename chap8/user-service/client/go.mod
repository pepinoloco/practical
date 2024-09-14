module practical/chap8/user-sevice/client

go 1.16

require (
	practical/chap8/user-service/service v0.0.0
	google.golang.org/grpc v1.37.0 // indirect
)

replace practical/chap8/user-service/service => ../service
