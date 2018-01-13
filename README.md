# DummyServer #

DummyServer is a lightweight tool for running a dummy debug HTTP server with no dependencies.

This is useful in cases where you want to ensure connectivity, without installing a fully fledged web server (Nginx, Apache etc.) or if you want to debug connectivity issues.

The server will print information about each request it receives and also provide debug info in the HTTP response.

## Instalation ##
If you have go installed on you system you can use it to download and install:
```bash
go get github.com/diarmuidie/dummyserver
```

Alternatively download a precompiled binary file from the [releases tab](https://github.com/diarmuidie/assistme/releases) (replace the link below with the binary suitable for your environment):
```bash
curl -o dummyserver https://github.com/diarmuidie/assistme/releases/download/1.0.0/dummyserver-1.0.0-darwin-386
```

## Usage ##

To serve dummy requests on port 80:
```bash
dummyserver -address :80
```
(You might need to run the command as root when serving requests on privileged ports, < 1024).

Running the server on a specified IP:
```bash
dummyserver -address 127.0.0.1:8080
```

You can also set the response format:
```bash
dummyserver -format json
```
Available formatters:
- `json`
- `text`
- `html`
- `auto` (Default. Chooses the correct response format based on the request headers)
