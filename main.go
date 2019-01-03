/*
 * [Turn on Tor, for use proxy]:
 *     $ sudo systemctl start tor.service
 * [Example]:
 *     $ go build main.go 
 *     $ ./main google.com -t a -a href -tp -ua
 *     $ ./main google.com --tag a --attr href --tor-proxy --user-agent
 */
package main

import (
    "os"
    "./parse"
    "./settings"
)

func main() {
    settings.CheckArgs(os.Args)
    parse.PrintResult(parse.OpenURL("http://" + os.Args[1]), "\n")
}
