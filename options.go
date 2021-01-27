package main

import (
	"fmt"
	"github.com/shumybest/ragnaros"
	"github.com/shumybest/ragnaros/config"
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "profile,P",
		Usage:       "profile to specify runtime environment (dev, sit, prod)",
		EnvVars:      []string{"SPRING_PROFILES_ACTIVE"},
		Destination: &config.Context.Profiles,
	},
	&cli.StringFlag{
		Name:        "port,p",
		Usage:       "http service port",
		EnvVars:      []string{"SERVER_PORT"},
		Destination: &config.Context.Port,
	},
	&cli.StringFlag{
		Name:        "config,conf",
		Usage:       "specify the directory contains configuration files",
		EnvVars:      []string{"CONF_DIR"},
		Destination: &config.Context.ConfDir,
	},
}

var Commands = []*cli.Command{
	{
		Name:    "version",
		Aliases: []string{"ver"},
		Usage:   "framework version",
		Action: func(c *cli.Context) error {
			fmt.Println(config.GetVersion())
			return nil
		},
	},
	{
		Name:   "start",
		Usage:  "start micro service",
		Action: ragnaros.StartOpts,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "conn,c",
				Value:       "",
				Usage:       "connection string for eureka server",
				EnvVars:     []string{"EUREKA_CLIENT_SERVICE_URL_DEFAULTZONE"},
			},
		},
	},
}
