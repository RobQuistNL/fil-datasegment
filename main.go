package main

import (
	"context"
	"fmt"
	"github.com/ribasushi/fil-datasegment/pkg/dlass"
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/ribasushi/go-toolbox/ufcli"
)

var log = logging.Logger(fmt.Sprintf("%s(%d)", "fil-datasegment", os.Getpid()))

func main() {

	logging.SetLogLevel("*", "INFO")

	(&ufcli.UFcli{
		Logger:              log,
		AllowConcurrentRuns: true,
		AppConfig: ufcli.App{
			Name:  "fil-datasegment",
			Usage: "Basic CLI app for downloading FRC58-formatted aggregates",
			Commands: []*ufcli.Command{
				dlass.DownloadAndAssemble,
			},
		},
	}).RunAndExit(context.Background())
}
