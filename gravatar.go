package gravatar

type Gravatar struct {
	Email    string
	Password string
}

type Image struct {
	ID       string
	Rating   int
	ImageURL string
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

func (g *Gravatar) AddressInfo() []Address {
	return hashToImageArray(Addresses(Hash(g.Email), g.Password))
}

func hashToImageArray(addresses map[string]ImageResponse) []Address {
	result := make([]Address, len(addresses))

	index := 0
	for key, value := range addresses {
		result[index] = Address{
			ID: key,
			Image: Image{
				Rating:   value.Rating,
				ID:       value.UserImage,
				ImageURL: value.UserImageURL,
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
