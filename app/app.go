package main

import (
	"log"
	"os"

	hotdaily "github.com/TaceyWong/HotDaily"
	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate
}

var ServerCMD = &cli.Command{
	Name:            "server",
	Usage:           "Web-API Server相关命令",
	Action:          serverAction,
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		{
			Name: "start", Usage: "启动API Server",
			Action: func(ctx *cli.Context) error {
				host := ctx.String("host")
				port := ctx.Int("port")
				hotdaily.Server(host, port)
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "host", Aliases: []string{"H"}, Usage: "设置服务`主机IP`", Value: "127.0.0.1"},
				&cli.IntFlag{Name: "port", Aliases: []string{"P"}, Usage: "设置服务`端口`", Value: 8080},
				&cli.BoolFlag{Name: "debug", Aliases: []string{"D"}, Usage: "启动调试模式", Value: false},
			},
		},
	},
}

func serverAction(ctx *cli.Context) error {
	return nil
}

func main() {
	app := &cli.App{
		Name:  "hd",
		Usage: "HotDaily - 热门新闻数据聚合API服务",
		Authors: []*cli.Author{
			{Name: "Tacey Wong", Email: "xinyong.wang@qq.com"}},
		Copyright: "© 2024 Tacey Wong All Rights Reserved",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{ServerCMD},
	}
	app.HideHelpCommand = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
