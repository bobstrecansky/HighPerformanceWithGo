Generate userinfo.pb.go with:
protoc -I ../userinfo --go_out=plugins=grpc:../userinfo ../userinfo/userinfo.proto
