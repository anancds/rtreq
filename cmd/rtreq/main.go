package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bbengfort/rtreq"
	"github.com/joho/godotenv"
	zmq "github.com/pebbe/zmq4"
	"github.com/urfave/cli"
)

//===========================================================================
// Main Method
//===========================================================================

func main() {

	// Load the .env file if it exists
	godotenv.Load()

	// Instantiate the command line application
	app := cli.NewApp()
	app.Name = "rtreq"
	app.Version = "0.1"
	app.Usage = "run async zmq server or client with REQ/ROUTER pattern"

	// Define commands available to the application
	app.Commands = []cli.Command{
		{
			Name:     "serve",
			Usage:    "run the rtreq server",
			Category: "server",
			Action:   serve,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address to bind the server to",
					Value: "*:4157",
				},
				cli.StringFlag{
					Name:  "n, name",
					Usage: "name to identify the server (default is hostname)",
				},
				cli.StringFlag{
					Name:  "u, uptime",
					Usage: "pass a parsable duration to shut the server down after",
				},
				cli.UintFlag{
					Name:  "verbosity",
					Usage: "set log level from 0-4, lower is more verbose",
					Value: 2,
				},
			},
		},
		{
			Name:     "send",
			Usage:    "send a message to the server",
			Category: "client",
			Action:   send,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address to connect to the server on",
					Value: "localhost:4157",
				},
				cli.StringFlag{
					Name:  "n, name",
					Usage: "name to identify the client (default is hostname)",
				},
				cli.StringFlag{
					Name:  "t, timeout",
					Usage: "recv timeout for each message",
					Value: "5s",
				},
				cli.IntFlag{
					Name:  "r, retries",
					Usage: "number of retries before quitting",
					Value: 3,
				},
			},
		},
		{
			Name:     "bench",
			Usage:    "run throughput benchmarks",
			Category: "client",
			Action:   bench,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address to connect to the server on",
					Value: "localhost:4157",
				},
				cli.StringFlag{
					Name:  "n, name",
					Usage: "name to identify the server (default is hostname)",
				},
				cli.StringFlag{
					Name:  "d, duration",
					Usage: "parsable duration of the benchmark",
					Value: "30s",
				},
				cli.StringFlag{
					Name:  "t, timeout",
					Usage: "recv timeout for each message",
					Value: "5s",
				},
				cli.IntFlag{
					Name:  "r, retries",
					Usage: "number of retries before quitting",
					Value: 3,
				},
				cli.IntFlag{
					Name:  "c, clients",
					Usage: "extra information: number of clients",
				},
				cli.StringFlag{
					Name:  "o, results",
					Usage: "path to write the results to",
					Value: "results.json",
				},
				cli.UintFlag{
					Name:  "verbosity",
					Usage: "set log level from 0-4, lower is more verbose",
					Value: 3,
				},
			},
		},
	}

	// Run the CLI program
	app.Run(os.Args)
}

//===========================================================================
// Server Commands
//===========================================================================

func exit(msg string, err error, a ...interface{}) error {
	if msg != "" {
		msg = fmt.Sprintf(msg, a...)
		msg += ": %s"
	} else {
		msg = "fatal error: %s"
	}
	return cli.NewExitError(fmt.Sprintf(msg, err), 1)
}

func serve(c *cli.Context) error {
	defer zmq.Term()
	context, err := zmq.NewContext()
	if err != nil {
		return exit("could not create zmq context", err)
	}

	// Set the debug log level
	verbose := c.Uint("verbosity")
	rtreq.SetLogLevel(uint8(verbose))

	// Create the server
	server := new(rtreq.Server)
	server.Init(c.String("addr"), c.String("name"), context)

	// If uptime is specified, set a fixed duration for the server to run.
	if uptime := c.String("uptime"); uptime != "" {
		d, err := time.ParseDuration(uptime)
		if err != nil {
			return exit("could not parse uptime", err)
		}

		time.AfterFunc(d, func() {
			zmq.Term()
			os.Exit(0)
		})
	}

	// Run the network server and broadcast clients
	if err := server.Run(); err != nil {
		return exit("could not run server", err)
	}
	return nil
}

//===========================================================================
// Client Commands
//===========================================================================

func send(c *cli.Context) error {
	defer zmq.Term()
	context, err := zmq.NewContext()
	if err != nil {
		return exit("", err)
	}

	client := new(rtreq.Client)
	client.Init(c.String("addr"), c.String("name"), context)

	if err = client.Connect(); err != nil {
		return exit("", err)
	}

	var timeout time.Duration
	if timeout, err = time.ParseDuration(c.String("timeout")); err != nil {
		return exit("", err)
	}

	for _, msg := range c.Args() {
		if err := client.Send(msg, c.Int("retries"), timeout); err != nil {
			exit("", err)
		}
	}

	return client.Close()
}

func bench(c *cli.Context) error {

	// Set the debug log level
	verbose := c.Uint("verbosity")
	rtreq.SetLogLevel(uint8(verbose))

	defer zmq.Term()
	context, err := zmq.NewContext()
	if err != nil {
		exit("", err)
	}

	client := new(rtreq.Client)
	client.Init(c.String("addr"), c.String("name"), context)

	if err = client.Connect(); err != nil {
		return exit("", err)
	}
	defer client.Close()

	var duration time.Duration
	if duration, err = time.ParseDuration(c.String("duration")); err != nil {
		return exit("", err)
	}

	var timeout time.Duration
	if timeout, err = time.ParseDuration(c.String("timeout")); err != nil {
		return exit("", err)
	}

	nClients := c.Int("clients")
	retries := c.Int("retries")
	results := c.String("results")

	return client.Benchmark(duration, results, retries, timeout, nClients)
}