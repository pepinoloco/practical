module practical/chap10/user-sevice-tls/server

go 1.16

require google.golang.org/grpc v1.37.0

require practical/chap10/user-service-tls/service v0.0.0

replace practical/chap10/user-service-tls/service => ../service
