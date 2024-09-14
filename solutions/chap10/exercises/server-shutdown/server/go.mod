module practical/chap10/server-shutdown/server

go 1.16

require google.golang.org/grpc v1.37.0

require practical/chap10/server-shutdown/service v0.0.0

replace practical/chap10/server-shutdown/service => ../service
