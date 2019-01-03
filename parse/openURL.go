package parse

import (
    "../lib"
    "../settings"
    "time"
    "net/url"
    "net/http"
    "golang.org/x/net/proxy"
)

// Function read html content from URL page.
func OpenURL(urlpage string) (html_page string) {
    var client *http.Client = &http.Client{}

    // Turn on tors proxy.
    if settings.TorsProxy {
        torProxyUrl, err := url.Parse(settings.TORS_PROXY)
        lib.CheckError(err)

        torDialer, err := proxy.FromURL(torProxyUrl, proxy.Direct)
        lib.CheckError(err)

        torTransport := &http.Transport {
            Dial: torDialer.Dial,
        }

        client = &http.Client {
            Transport: torTransport, 
            Timeout: time.Second * 5,
        }
    }

    req, err := http.NewRequest("GET", urlpage, nil)
    lib.CheckError(err)

    // Turn on random user-agent from array.
    req.Header.Add("Accept", "text/html")
    if settings.UserAgent {
        req.Header.Add(
            "User-Agent", 
            settings.USER_AGENT[lib.RandInt(0, len(settings.USER_AGENT))],
        )
    } 

    resp, err := client.Do(req)
    lib.CheckError(err)
    defer resp.Body.Close()

    var buffer []byte = make([]byte, settings.BUFF_SIZE)
    for {
        length, err := resp.Body.Read(buffer)
        html_page += string(buffer[:length])
        if length == 0 || err != nil { break }
    }

    return html_page
}
