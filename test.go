package httpclient

import "github.com/tvn9/httpclient/gohttp"

func exampleUsage() {
	client := gohttp.New()

	client.Get()
}
