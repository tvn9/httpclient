package gohttp

type httpClient struct {
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Patch()
	Put()
	Delete()
}

func (c *httpClient) Get()    {}
func (c *httpClient) Post()   {}
func (c *httpClient) Patch()  {}
func (c *httpClient) Put()    {}
func (c *httpClient) Delete() {}
