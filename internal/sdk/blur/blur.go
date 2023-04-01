package blur

// Used to encapsulate some of blur's functional apis

type Client struct {
}

type ClientWs struct {
}

func New() (*Client, error) {
	return &Client{}, nil
}

func NewWs() (*ClientWs, error) {
	return &ClientWs{}, nil
}
