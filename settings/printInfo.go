package settings

import (
	"os"
	"fmt"
)

// Function prints modules and example of use.
func PrintInfo() {
    fmt.Println(
`Modules:
    -tp || --tor-proxy  -> Turn on tor proxy
    -ua || --user-agent -> Turn on user-agent
    -t  || --tag        -> Tag name
    -a  || --attr       -> Attribute name
Example:
    $ ./parse google.com --tag a --attr href -tp -ua`)
    os.Exit(0)
}
