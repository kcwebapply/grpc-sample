#get
grpcurl -plaintext -import-path proto/user  -proto user.proto localhost:9999 user.UserService/GetUsers

#save
grpcurl -d '{"id":"600","name":"hiroshi","address":"東京"}' -plaintext -import-path proto/user  -proto user.proto localhost:9999  user.UserService/SaveUser


# progobuf genenate
protoc proto/user/user.proto --go_out=protogen


# doc generate
protoc --doc_out=html,index.html:./ proto/*.proto
