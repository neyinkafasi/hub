package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/v1/types"
)

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

func TestNewQueryProviderRequest(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		require.Equal(
			t,
			&QueryProviderRequest{
				Address: hubtypes.ProvAddress(address).String(),
			},
			NewQueryProviderRequest(address),
		)
	}
}

func TestNewQueryProvidersRequest(t *testing.T) {
	var (
		status     hubtypes.Status
		pagination *query.PageRequest
	)

	for i := 0; i < 20; i++ {
		status = hubtypes.Status(i % 4)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryProvidersRequest{
				Status:     status,
				Pagination: pagination,
			},
			NewQueryProvidersRequest(status, pagination),
		)
	}
}
