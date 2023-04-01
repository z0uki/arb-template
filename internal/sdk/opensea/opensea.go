package opensea

// Used to encapsulate some of opensea's functional apis

type Client struct {
}

func New() (*Client, error) {
	return &Client{}, nil
}
