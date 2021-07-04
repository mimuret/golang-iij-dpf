package api_test

import (
	"testing"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestIsAuthError(t *testing.T) {
	testcases := []struct {
		b *api.BadResponse
		r bool
	}{
		{
			&api.BadResponse{},
			false,
		},
		{
			&api.BadResponse{
				ErrorDetails: api.ErrorDetails{
					api.ErrorDetail{
						Code: "invalid",
					},
				},
			},
			false,
		},
		{
			&api.BadResponse{
				ErrorDetails: api.ErrorDetails{
					api.ErrorDetail{
						Code:      "invalid",
						Attribute: "hoge",
					},
				},
			},
			false,
		},
		{
			&api.BadResponse{
				ErrorDetails: api.ErrorDetails{
					api.ErrorDetail{
						Code:      "invalid",
						Attribute: "access_token",
					},
				},
			},
			true,
		},
		{
			&api.BadResponse{
				ErrorDetails: api.ErrorDetails{
					api.ErrorDetail{
						Code:      "invalid",
						Attribute: "hoge",
					},
					api.ErrorDetail{
						Code:      "invalid",
						Attribute: "access_token",
					},
				},
			},
			true,
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, api.IsAuthError(tc.b), tc.r)
	}
}
