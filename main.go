package main

import . "github.com/visionmedia/go-gracefully"
import . "github.com/segmentio/loggly-cat/pkg"
import "github.com/segmentio/go-loggly"
import "github.com/segmentio/go-log"
import "github.com/docopt/docopt-go"
import "time"
import "os"

// Version
const Version = "1.0.0"

// Usage
const Usage = `

  Usage:
    loggly-cat --token t [--tag t]...
    loggly-cat -h | --help
    loggly-cat --version

  Options:
    -t, --token t     loggly api token
    -T, --tag t       loggly tag(s)
    -h, --help        output help information
    -v, --version     output version

`

func check(err error) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	// options
	tags := args["--tag"].([]string)
	token := args["--token"].(string)

	// implicit tags
	host, _ := os.Hostname()
	tags = append(tags, "host-"+host)

	// configure
	log.SetPrefix("loggly-cat")

	l := loggly.New(token, tags...)
	t := NewTailer(os.Stdin, l)

	l.BufferSize = 2000
	l.FlushInterval = 10 * time.Second

	// start
	log.Info("starting loggly-cat %s", Version)
	t.Start()
	Shutdown()
	t.Stop()
}
