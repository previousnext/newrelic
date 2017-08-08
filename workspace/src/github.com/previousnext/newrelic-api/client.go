package newrelic_api

// Client is used to interacting with New Relics API.
type Client struct {
	key string
}

// New returns a new New Relic client.
func NewClient(key string) Client {
	return Client{
		key: key,
	}
}
