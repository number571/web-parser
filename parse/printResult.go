package parse

import (
	"fmt"
	"../settings"
)

// Function prints result with separator. 
func PrintResult(html, sep string) {
    switch settings.TagName {
        case "": 
        	fmt.Print(html)
        	fmt.Print(sep)
        default:
            for _, result := range GetObjects(html) {
                fmt.Print(result)
                fmt.Print(sep)
            }
    }
}
