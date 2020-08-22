package tracercli

import (
    "errors"
    "time"

    "github.com/urfave/cli"
)

const DEFAULT_PORT = 33424
const DEFAULT_TIMEOUT_MS = 5000
const DEFAULT_RETRIES = 3
const DEFAULT_MAX_HOPS = 24
const DEFAULT_INITIAL_TTL = 1
const DEFAULT_PACKET_SIZE = 64
const DEFAULT_DESTINATION = "google.com"

func Run(args []string) error {
    app := cli.NewApp()
    app.Name = "tracer"
    app.Compiled = time.Now()
    app.Usage = "a simple implementation of traceroute in go."
    app.UsageText = "trace a destination: tracer --hop=3 --timeout=4000 google.com"
    app.Flags = []cli.Flag{
        cli.IntFlag{Name: "hop", Value: DEFAULT_INITIAL_TTL, Usage: "start from a custom hop number"},
		cli.IntFlag{Name: "maxhops", Value: DEFAULT_MAX_HOPS, Usage: "custom max hops"},
		cli.IntFlag{Name: "port", Value: DEFAULT_PORT, Usage: "custom port number"},
		cli.IntFlag{Name: "timeout", Value: DEFAULT_TIMEOUT_MS, Usage: "custom timeout in ms"},
		cli.IntFlag{Name: "retries", Value: DEFAULT_RETRIES, Usage: "custom retries"},
		cli.IntFlag{Name: "packetsize", Value: DEFAULT_PACKET_SIZE, Usage: "custom packet size"},
    }
    app.Action = func(c *cli.Context) error {
        if c.NArg() >= 1 {
            println("Running tracer.")
            return nil
        } else {
            return errors.New("Improper usage, type in --help to see usage details.")
        }
    }
    return app.Run(args)
}
