package parse

import (
    "../settings"
    "regexp"
)

// Function takes objects from tag/attribute.
func GetObjects(html string) (slice_result []string) {
    var (
        index int
        regular *regexp.Regexp 
    )

    switch settings.AttrName {
        case "": regular = regexp.MustCompile(`<` + settings.TagName + `[^>]*>.+</` + settings.TagName + `>`)
        default:
            regular = regexp.MustCompile(`<` + settings.TagName + `.*?` + settings.AttrName + `\s*=\s*['"]([^\s'"]+)[\s'"]`)
            index = 1
    }

    result := regular.FindAllStringSubmatch(html, -1)
    for _, slice := range result {
        slice_result = append(slice_result, slice[index])
    }

    return slice_result
}
