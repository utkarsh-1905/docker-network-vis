package main

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
	util "github.com/utkarsh-1905/docker-network-vis/utils"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	util.HandleError(err)

	ctx := context.Background()

	app := fiber.New()

	app.Get("/networks", func(c *fiber.Ctx) error {
		networks := util.GetNetworks(ctx, cli)
		return c.SendString(networks)
	})

	app.Listen(":3001")

}

// f40291c5c6f30f44251fa288632922ba3ca75774b1b721bf50f21bc14dff8840 - custom-test
