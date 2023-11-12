package graphql

import (
	"net/http"
	"os"

	graphql "github.com/hasura/go-graphql-client"
)

type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Set(k, v[0])
	}
	return t.base.RoundTrip(req)
}
func Client() *graphql.Client {
	headers := http.Header{}
	headers.Add("X-Hasura-Admin-Secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))
	httpClient := &http.Client{Transport: &headersTransport{headers, http.DefaultTransport}}
	newClient := graphql.NewClient(os.Getenv("HASURA_GRAPHQL_URL"), httpClient)
	return newClient
}
