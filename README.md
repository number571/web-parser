# GoWebParser

#### Example:
```
$ go build -o main main.go
$ sudo systemctl start tor.service
$ ./main github.com/Number571 --tag a --attr href --tor-proxy --user-agent
```

#### Result:
```
#start-of-content
https://chrome.google.com
https://mozilla.org/firefox/
https://help.github.com/articles/supported-browsers
https://github.com/
...
https://developer.github.com
https://training.github.com
https://blog.github.com
https://github.com/about
```
