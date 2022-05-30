package metrics

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	ts := httptest.NewServer(Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})))
	defer ts.Close()

	tsTLS := httptest.NewTLSServer(Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})))
	defer tsTLS.Close()

	var cases = []struct {
		in     string
		server *httptest.Server
	}{
		{
			"standard http",
			ts,
		},
		{
			"tls",
			tsTLS,
		},
	}

	for _, tt := range cases {
		t.Run(tt.in, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.server.URL, nil)
			assert.NoError(t, err)

			resp, err := tt.server.Client().Do(req)
			assert.NoError(t, err)
			assert.EqualValues(t, http.StatusOK, resp.StatusCode)
		})
	}
}
