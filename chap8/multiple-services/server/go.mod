module practical/chap8/multiple-sevices/server

go 1.16

require google.golang.org/grpc v1.37.1

require practical/chap8/multiple-services/service v0.0.0

replace practical/chap8/multiple-services/service => ../service
