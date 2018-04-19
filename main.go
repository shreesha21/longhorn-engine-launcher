package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-units"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/yasker/longhorn-engine-launcher/rpc"
)

const (
	UpgradeTimeout = 10 * time.Second
)

func StartCmd() cli.Command {
	return cli.Command{
		Name: "start",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "longhorn-binary",
				Value: "/usr/local/bin/longhorn",
			},
			cli.StringFlag{
				Name:  "launcher-listen",
				Value: "localhost:9510",
			},
			cli.StringFlag{
				Name: "size",
			},
			cli.StringFlag{
				Name:  "listen",
				Value: "localhost:9501",
			},
			cli.StringFlag{
				Name:  "frontend",
				Value: "tgt-blockdev",
				Usage: "Supports tgt-blockdev",
			},
			cli.StringSliceFlag{
				Name:  "enable-backend",
				Value: (*cli.StringSlice)(&[]string{"tcp"}),
			},
			cli.StringSliceFlag{
				Name: "replica",
			},
		},
		Action: func(c *cli.Context) {
			if err := start(c); err != nil {
				logrus.Fatalf("Error running start command: %v.", err)
			}
		},
	}
}

func UpgradeCmd() cli.Command {
	return cli.Command{
		Name: "upgrade",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "longhorn-binary",
				Value: "/usr/local/bin/longhorn",
			},
			cli.StringFlag{
				Name:  "listen",
				Value: "localhost:9501",
			},
			cli.StringFlag{
				Name:  "frontend",
				Value: "tgt-blockdev",
				Usage: "Supports tgt-blockdev",
			},
			cli.StringSliceFlag{
				Name:  "enable-backend",
				Value: (*cli.StringSlice)(&[]string{"tcp"}),
			},
			cli.StringSliceFlag{
				Name: "replica",
			},
		},
		Action: func(c *cli.Context) {
			if err := upgrade(c); err != nil {
				logrus.Fatalf("Error running start command: %v.", err)
			}
		},
	}
}

func start(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("volume name is required")
	}
	name := c.Args()[0]

	launcherListen := c.String("launcher-listen")
	longhornBinary := c.String("longhorn-binary")

	listen := c.String("listen")
	backends := c.StringSlice("enable-backend")
	replicas := c.StringSlice("replica")
	frontend := c.String("frontend")

	sizeString := c.String("size")
	if sizeString == "" {
		return fmt.Errorf("Invalid empty size")
	}
	size, err := units.RAMInBytes(sizeString)
	if err != nil {
		return err
	}

	l, err := NewLauncher(launcherListen, longhornBinary, frontend, name, size)
	if err != nil {
		return err
	}
	controller := NewController(longhornBinary, name, listen, backends, replicas)
	if err := l.StartController(controller); err != nil {
		return err
	}
	if err := l.StartRPCServer(); err != nil {
		return err
	}
	return l.WaitForShutdown()
}

func upgrade(c *cli.Context) error {
	url := c.GlobalString("url")
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to %v: %v", url, err)
	}
	defer conn.Close()

	longhornBinary := c.String("longhorn-binary")
	listen := c.String("listen")
	backends := c.StringSlice("enable-backend")
	replicas := c.StringSlice("replica")

	client := rpc.NewLonghornLauncherServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), UpgradeTimeout)
	defer cancel()

	if _, err := client.UpgradeEngine(ctx, &rpc.Engine{
		Binary:         longhornBinary,
		Listen:         listen,
		Replicas:       replicas,
		EnableBackends: backends,
	}); err != nil {
		return fmt.Errorf("failed to upgrade: %v", err)
	}
	return nil
}

func main() {
	a := cli.NewApp()
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Value: "http://localhost:9510",
		},
		cli.BoolFlag{
			Name: "debug",
		},
	}
	a.Commands = []cli.Command{
		StartCmd(),
	}
	if err := a.Run(os.Args); err != nil {
		logrus.Fatal("Error when executing command: ", err)
	}
}