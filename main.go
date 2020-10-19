package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	name    = "dotenv"
	version = "0.0.0"
	date    = "0001-01-01T00:00:00Z"
	commit  = "0000000"
)

func ver() string {
	return fmt.Sprintf("%s v%s (%s) %s", name, version, commit[:7], date)
}

func usage() {
	fmt.Fprintf(os.Stderr, "%s\n", ver())

	fmt.Fprintf(os.Stderr, `
Usage:
    %s [-f .env.alternate] -- <command> [arguments...]"

`, name)

	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, `
Example:
    %s -f .env -- caddy run --config Caddyfile

`, name)
	//See '%s --help' for details."
}

func main() {
	if len(os.Args) > 1 {
		if "version" == strings.TrimLeft(os.Args[1], "-") || "-V" == os.Args[1] {
			fmt.Println(ver())
			os.Exit(0)
			return
		}
	}

	_ = flag.Bool("version", false, "print version and exit")
	help := flag.Bool("help", false, "print usage and exit")
	envList := flag.String("f", ".env", "path to .env file")
	flag.Parse()

	if *help {
		usage()
		os.Exit(0)
		return
	}

	envs := strings.Fields(
		strings.ReplaceAll(*envList, ",", " "),
	)

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(1)
		return
	}

	cmd := args[0]
	cmdArgs := args[1:]
	if err := godotenv.Exec(envs, cmd, cmdArgs); nil != err {
		fmt.Fprintf(os.Stderr, "error executing %q: %v\n", strings.Join(args, " "), err)
		os.Exit(1)
		return
	}
}
