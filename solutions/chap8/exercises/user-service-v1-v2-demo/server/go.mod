module practical/chap8/user-sevice/server

go 1.16

require google.golang.org/grpc v1.37.1

require practical/chap8/user-service/service-v2 v0.0.0

replace practical/chap8/user-service/service-v2 => ../service-v2
