package main

import (
	"github.com/kolo/xmlrpc"
)

const (
	apiURI string = "https://secure.gravatar.com/xmlrpc?user="
)

type request struct {
	Password  string   `xmlrpc:"password"`
	Hashes    []string `xmlrpc:"hashes"`
	Data      string   `xmlrpc:"data"`
	Rating    int      `xmlrpc:"rating"`
	UserImage string   `xmlrpc:"userimage"`
	Addresses []string `xmlrpc:"addresses"`
}

func Exists(hash string, password string, hashes []string) map[string]bool {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string]int
	client.Call("grav.exists", Request{Password: password, Hashes: hashes}, &response)

	result := map[string]bool{}
	for hash, value := range response {
		result[hash] = value == 1
	}

	return result
}
