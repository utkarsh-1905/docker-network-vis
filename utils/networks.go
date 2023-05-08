package util

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetNetworks(ctx context.Context, cli *client.Client) string {
	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	HandleError(err)
	networkJson, err := json.MarshalIndent(networks, "", "  ")
	HandleError(err)
	return string(networkJson)
}

func GetNetworkInfo(ctx context.Context, cli *client.Client, id string) string {
	network, err := cli.NetworkInspect(ctx, id, types.NetworkInspectOptions{})
	HandleError(err)
	networkJson, err := json.MarshalIndent(network, "", "  ")
	HandleError(err)
	return string(networkJson)
}
