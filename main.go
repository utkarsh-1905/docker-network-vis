package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	util "github.com/utkarsh-1905/docker-network-vis/utils"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	util.HandleError(err)

	ctx := context.Background()

	engine := html.New("./views", ".html") // engine for rendering html files

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	type FullNetwork struct {
		Name        string
		NetworkInfo types.NetworkResource
	}

	app.Get("/networks", func(c *fiber.Ctx) error {
		networks := util.GetNetworks(ctx, cli)
		var fullNetwork []FullNetwork
		for _, network := range networks {
			networkInfo := util.GetNetworkInfo(ctx, cli, network.ID)
			var container FullNetwork = FullNetwork{
				Name:        network.Name,
				NetworkInfo: networkInfo,
			}
			fullNetwork = append(fullNetwork, container)
		}
		networkJson, err := json.MarshalIndent(fullNetwork, "", "  ")
		util.HandleError(err)
		return c.SendString(string(networkJson))
	})

	app.Get("/networks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		network := util.GetNetworkInfo(ctx, cli, id)
		networkJson, err := json.MarshalIndent(network, "", "  ")
		util.HandleError(err)
		return c.SendString(string(networkJson))
	})

	app.Get("/", func(c *fiber.Ctx) error {
		networks := util.GetNetworks(ctx, cli)
		var fullNetwork []FullNetwork
		for _, network := range networks {
			networkInfo := util.GetNetworkInfo(ctx, cli, network.ID)
			var container FullNetwork = FullNetwork{
				Name:        network.Name,
				NetworkInfo: networkInfo,
			}
			fullNetwork = append(fullNetwork, container)
		}
		return c.Render("index", fiber.Map{
			"networks": fullNetwork,
		})
	})

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))

}
