protoc -I=. --go_out=. proto.proto
protoc-go-inject-tag -input=proto.pb.go
