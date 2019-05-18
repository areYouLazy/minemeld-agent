# MineMeld-Agent

This tool has been written to address the need to query MineMeld for a specific IP address to know if it matches a MineMeld list

## Composition

- Loader: this loads a urls.json file containing a list of URL to fetch
- Fetcher: this takes URLs from Loader and fetch lists
- Parser: this parse URL List into a golang object
- Checker: this checks if a given IP/FQDN is in a MineMeld list
- WebServer: this exposes routines through API

## WebServer API

```text
GET /api/v1/check-ipv4/1.1.1.1

GET /api/v1/check-ipv6/::1

GET /api/v1/check-fqdn/example.org
```

## Installation

To get the tool you need [golang](https://golang.org/)

```bash
root@localhost:> go version
go version go1.12.5 darwin/amd64
```

You can download MineMeld-Agent with the command

```bash
root@localhost:> go get github.com/areYouLazy/minemeld-agent
```

MineMeld-Agent can run on the MineMeld machine itself, or in a separate linux machine, just make sure URL in the urls.json file are resolvable (if you're using FQDN) and to use the `-fetch-insecure` flag if MineMeld does not provide a valid certificate.

By default MineMeld-Agent logs to stdout so you can check that everything is working file.
You can than redirect logs to your preferred file with the `-log-output` flag

## Configuration

Endpoint Must be written in the urls.json (or in a custom JSON file) in the format of

```json
[
    {
        "type": "ipv4",
        "endpoint": "https://minemeld.example.org/feeds/office365_IPv4s",
        "description": "MineMeld IPv4 feed for Office365"
    },
    {
        "type": "ipv6",
        "endpoint": "https://minemeld.example.org/feeds/office365_IPv6s",
        "description": "MineMeld IPv6 feed for Office365"
    }
]
```

Valid Entpoint types are:

- ipv4
- ipv6
- fqdn

Any other type will throw an error

## URL Fetch

By design fetch is done for every URL, every 10 seconds

## Help

```text
Usage of minemeld-agent:
  -fetch-insecure
        Set to true to ignore certificate errors while fetching MineMeld URLs
  -log-colors
        Set to false to turn off colored log output (default true)
  -log-debug
        Set true to print debug message
  -log-output string
        Set the output interface for log
  -url-file string
        PATH of the JSON file containing urls. (default "urls.json")
  -webserver-port int
        Specify port for WebServer (default 9000)
```
