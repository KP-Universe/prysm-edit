package validator_client_factory

import (
	"github.com/KP-Universe/prysm/v4/config/features"
	beaconApi "github.com/KP-Universe/prysm/v4/validator/client/beacon-api"
	grpcApi "github.com/KP-Universe/prysm/v4/validator/client/grpc-api"
	"github.com/KP-Universe/prysm/v4/validator/client/iface"
	validatorHelpers "github.com/KP-Universe/prysm/v4/validator/helpers"
)

func NewNodeClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.NodeClient {
	grpcClient := grpcApi.NewNodeClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewNodeClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}
