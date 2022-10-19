package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// This function returns the CLI application ready for use
func Generate() (app *cli.App) {
	app = cli.NewApp()

	app.Name = "Go CLI"

	app.Usage = "Search for IP addresses and servers names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "This command searches for IP addresses on the internet",
			Flags:  flags,
			Action: searchIP,
		},
		{
			Name:   "hosts",
			Usage:  "This command search the host names on the internet",
			Flags:  flags,
			Action: searchHosts,
		},
	}

	return app
}

// This function is used as an available command for the Go CLI application
// This function is responsible for getting the IP addresses and search on the internet
func searchIP(ctx *cli.Context) {
	host := ctx.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Searching hosts for:", host)

	for i, ip := range ips {
		fmt.Printf("%d: %s\n", i+1, ip)
	}
}

// This function is used as an available command for the Go CLI application
// This function is responsible for getting the host names on the internet
func searchHosts(ctx *cli.Context) {
	host := ctx.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Searching hosts names for:", host)

	for i, server := range servers {
		fmt.Printf("%d: %s\n", i+1, server.Host)
	}
}
