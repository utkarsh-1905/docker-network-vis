package util

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetNetworks(ctx context.Context, cli *client.Client) []types.NetworkResource {
	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	HandleError(err)
	return networks
}

func GetNetworkInfo(ctx context.Context, cli *client.Client, id string) types.NetworkResource {
	network, err := cli.NetworkInspect(ctx, id, types.NetworkInspectOptions{})
	HandleError(err)
	return network
}
