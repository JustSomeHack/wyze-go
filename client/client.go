package client

type WyzeClient interface {
}

type wyzeClient struct {
	AuthURL  string
	BaseURL  string
	Email    string
	Password string
	KeyID    string
	APIKey   string
}

func NewWyzeClient(Email string, Password string, KeyID string, APIKey string) WyzeClient {
	return &wyzeClient{
		AuthURL:  AuthURL,
		BaseURL:  BaseURL,
		Email:    Email,
		Password: Password,
		KeyID:    KeyID,
		APIKey:   APIKey,
	}
}
