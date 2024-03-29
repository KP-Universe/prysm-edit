package beacon_api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KP-Universe/go-kpu/common/hexutil"
	"github.com/pkg/errors"
	"github.com/KP-Universe/prysm/v4/consensus-types/primitives"
	ethpb "github.com/KP-Universe/prysm/v4/proto/prysm/v1alpha1"
)

// IndexNotFoundError represents an error scenario where no validator index matches a pubkey.
type IndexNotFoundError struct {
	message string
}

// NewIndexNotFoundError creates a new error instance.
func NewIndexNotFoundError(pubkey string) IndexNotFoundError {
	return IndexNotFoundError{
		message: fmt.Sprintf("could not find validator index for public key `%s`", pubkey),
	}
}

// Error returns the underlying error message.
func (e *IndexNotFoundError) Error() string {
	return e.message
}

func (c beaconApiValidatorClient) validatorIndex(ctx context.Context, in *ethpb.ValidatorIndexRequest) (*ethpb.ValidatorIndexResponse, error) {
	stringPubKey := hexutil.Encode(in.PublicKey)

	stateValidator, err := c.stateValidatorsProvider.GetStateValidators(ctx, []string{stringPubKey}, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get state validator")
	}

	if len(stateValidator.Data) == 0 {
		e := NewIndexNotFoundError(stringPubKey)
		return nil, &e
	}

	stringValidatorIndex := stateValidator.Data[0].Index

	index, err := strconv.ParseUint(stringValidatorIndex, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse validator index")
	}

	return &ethpb.ValidatorIndexResponse{Index: primitives.ValidatorIndex(index)}, nil
}
