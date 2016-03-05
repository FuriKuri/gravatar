package gravatar

type Gravatar struct {
	Email    string
	Password string
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

func hashToEmailMap(emails []string, hashExisting map[string]bool) map[string]bool {
	result := map[string]bool{}

	for _, email := range emails {
		hashKey := Hash(email)
		result[email] = hashExisting[hashKey]
	}

	return result
}
