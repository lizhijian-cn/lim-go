package main

var (
	HttpUrl = "http://localhost:8080/"
	Socket  = "127.0.0.1:8999"

	UserId int64

	FriendIdMap = make(map[string]int64)
	IdFriendMap = make(map[int64]string)
)
