package settings

import "../lib"

// Function checking arguments and turn on modules.
func CheckArgs(args []string) {
    var (
        flag_tag bool
        flag_atr bool
    )

    switch len(args) {
        case 1: lib.PrintError("no url specified")
        case 2:
            switch args[1] {
                case "-h", "--help": PrintInfo()
                default: return
            }
        default:
            for _, value := range args[2:] {
                switch value {
                    case "-tp", "--tor-proxy":  TorsProxy = true; continue
                    case "-ua", "--user-agent": UserAgent = true; continue
                    case "-t", "--tag":  flag_tag = true; continue
                    case "-a", "--attr": flag_atr = true; continue
                }

                switch {
                    case flag_tag:
                        TagName = value
                        flag_tag = false
                    case flag_atr:
                        AttrName = value
                        flag_atr = false
                }
            }
    }
}
