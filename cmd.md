#get
grpcurl -plaintext -import-path proto/user  -proto user.proto localhost:9999 user.UserService/GetUser

#save
grpcurl -d '{"id":"600","name":"hiroshi","address":"東京"}' -plaintext -import-path proto/user  -proto user.proto localhost:9999  user.UserService/SaveUser
