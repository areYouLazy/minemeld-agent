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

GET /api/v1/check-fqdn/contoso.ltd
```

## Configuration

Endpoint Must be written in the urls.json (or in a custom JSON file) in the format of

```json
[
    {
        "type": "ipv4",
        "endpoint": "https://minemeld.contoso.ltd/feeds/office365_IPv4s",
        "description": "minemeld ipv4 feed for Office365"
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
  -insecure
        Set to true to ignore certificate errors
  -log-colors
        Set to false to turn off colored log output (default true)
  -log-debug
        Set true to print debug message
  -log-output string
        Set the output interface for log
  -url-file string
        PATH of the JSON file containing urls. (default "urls.json")
```
