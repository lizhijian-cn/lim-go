package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type HttpApi struct{}

func (api *HttpApi) Login(username, password string) (map[string]string, error) {
	loginForm := map[string]string{
		"username": username,
		"password": password,
	}

	jsonValue, _ := json.Marshal(loginForm)
	res, err := http.Post(fmt.Sprintf("%slogin", HttpUrl), "application/json",
		bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var info map[string]string
	err = json.NewDecoder(res.Body).Decode(&info)
	fmt.Println(info)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(info["err"])
	}

	userId, err := strconv.ParseInt(info["userId"], 10, 64)
	if err != nil {
		return nil, err
	}
	UserId = userId
	return info, nil
}

func (api *HttpApi) SyncMsg() error {
	res, err := http.Get(fmt.Sprintf("%srelation/%d", HttpUrl, UserId))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var info struct {
		FriendIds   []int64
		FriendNames []string
	}
	err = json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		return err
	}

	fmt.Println("here are you friends now:")
	for i, name := range info.FriendNames {
		FriendIdMap[name] = info.FriendIds[i]
		IdFriendMap[info.FriendIds[i]] = name
		fmt.Printf("%d: %s", info.FriendIds[i], name)
	}
	fmt.Println()
	return nil
}

func (api *HttpApi) GetUser(userId int64) {

}

func (api *HttpApi) AddFriend(userId int64) {

}

func (api *HttpApi) ReplyAddFriend() {

}
