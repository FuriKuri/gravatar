package gravatar

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

type ImageResponse struct {
	Rating       int    `xmlrpc:"rating"`
	UserImage    string `xmlrpc:"userimage"`
	UserImageURL string `xmlrpc:"userimage_url"`
}

func Exists(hash string, password string, hashes []string) map[string]bool {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string]int
	client.Call("grav.exists", request{Password: password, Hashes: hashes}, &response)

	result := map[string]bool{}
	for hash, value := range response {
		result[hash] = value == 1
	}

	return result
}

func UserAddresses(hash string, password string) map[string]ImageResponse {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string]ImageResponse
	client.Call("grav.addresses", request{Password: password}, &response)

	return response
}

func UserImages(hash string, password string) map[string][]string {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string][]string
	client.Call("grav.userimages", request{Password: password}, &response)

	return response
}

func UseUserImage(hash string, password string, userImage string, addresses []string) map[string]bool {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string]bool
	client.Call("grav.useUserimage", request{Password: password, Addresses: addresses, UserImage: userImage}, &response)

	return response
}
