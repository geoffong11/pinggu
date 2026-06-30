package sender

type User interface {
	Action() ([]byte, error)
}

type Post struct {
	Endpoint string
	Body     []byte
}

type Get struct {
	Endpoint string
}
