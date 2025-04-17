package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

/**
*	@return executable app from command line
* */
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca Ip's e Nomes de Servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
