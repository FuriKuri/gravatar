package gravatar

import (
	"encoding/base64"
	"fmt"
	"github.com/kolo/xmlrpc"
)

const (
	apiURI string = "https://secure.gravatar.com/xmlrpc?user="
)

type request struct {
	Password  string   `xmlrpc:"password"`
	Hashes    []string `xmlrpc:"hashes"`
	Data      string   `xmlrpc:"data"`
	URL       string   `xmlrpc:"url"`
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

func CallSaveURL(hash string, password, url string, rating int) string {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response string
	client.Call("grav.saveUrl", request{Password: password, URL: url, Rating: rating}, &response)

	return response
}

func CallSaveData(hash string, password string, data []byte, rating int) string {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)
	imageData := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("base 64 %s", imageData)
	var response string
	client.Call("grav.saveData", request{Password: password, Data: imageData, Rating: rating}, &response)

	return response
}

func CallRemoveImage(hash string, password string, addresses []string) map[string]bool {
	client, _ := xmlrpc.NewClient(apiURI+hash, nil)

	var response map[string]bool
	client.Call("grav.removeImage", request{Password: password, Addresses: addresses}, &response)

	return response
}
