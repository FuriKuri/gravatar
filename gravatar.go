package gravatar

import (
	"strconv"
)

type Gravatar struct {
	Email    string
	Password string
}

type Image struct {
	ID     string
	Rating int
	URL    string
}

type Address struct {
	ID    string
	Image Image
}

func New(email string, password string) (*Gravatar, error) {
	g := &Gravatar{
		Email:    email,
		Password: password,
	}
	return g, nil
}

func (g *Gravatar) IsImageSet(emails []string) map[string]bool {
	hashes := Map(emails, Hash)
	hashExisting := Exists(Hash(g.Email), g.Password, hashes)
	return hashToEmailMap(emails, hashExisting)
}

func (g *Gravatar) Addresses() []Address {
	return hashToAddressArray(UserAddresses(Hash(g.Email), g.Password))
}

func (g *Gravatar) Images() []Image {
	return hashToImageArray(UserImages(Hash(g.Email), g.Password))
}

func (g *Gravatar) SaveURL(url string, rating int) string {
	return CallSaveURL(Hash(g.Email), g.Password, url, rating)
}

func (g *Gravatar) SaveData(data []byte, rating int) string {
	return CallSaveData(Hash(g.Email), g.Password, data, rating)
}

func (g *Gravatar) RemoveImage(addresses []string) map[string]bool {
	return CallRemoveImage(Hash(g.Email), g.Password, addresses)
}

func (g *Gravatar) DeleteImage(image string) bool {
	return DeleteUserImage(Hash(g.Email), g.Password, image)
}

func hashToImageArray(images map[string][]string) []Image {
	result := make([]Image, len(images))

	index := 0
	for key, value := range images {
		rating, _ := strconv.Atoi(value[0])
		result[index] = Image{
			ID:     key,
			Rating: rating,
			URL:    value[1],
		}
		index = index + 1
	}

	return result
}

func hashToAddressArray(addresses map[string]ImageResponse) []Address {
	result := make([]Address, len(addresses))

	index := 0
	for key, value := range addresses {
		result[index] = Address{
			ID: key,
			Image: Image{
				Rating: value.Rating,
				ID:     value.UserImage,
				URL:    value.UserImageURL,
			},
		}
		index = index + 1
	}

	return result
}

func hashToEmailMap(emails []string, hashExisting map[string]bool) map[string]bool {
	result := map[string]bool{}

	for _, email := range emails {
		hashKey := Hash(email)
		result[email] = hashExisting[hashKey]
	}

	return result
}
