# cswsh-scanner

A command-line tool for Cross-Site WebSocket Hijacking (CSWSH)

## Install

```bash
$ go get -v -u github.com/ambalabanov/websocket/cswsh-scanner/...
```

## Basic Usage

cswsh-scanner accepts line-delimited URLs on `stdin`, output `csv`:

```bash
$ cat test.txt
ws://echo.websocket.org
wss://echo.websocket.org
$ cat test.txt | cswsh-scanner
true,ws://echo.websocket.org
true,wss://echo.websocket.org

```

## Extra parameters

You can use custom Origin header, socket.io support, verbose output and multithreading

```bash
$ cswsh-scanner -h
Usage of cswsh-scanner:
  -o string
    	Origin (default "http://hacker.com")
  -s	Socket.IO
  -v	Verbose output
  -w int
    	Number of workers (default 1)
```

Example

```bash
echo "wss://juice-shop.herokuapp.com/socket.io/" | cswsh-scanner -o http://example.com -s -v -w 10
GET /socket.io/?EIO=3&sid=UGv7cfNFrvOiAezJAFa4&transport=websocket HTTP/1.1
Host: juice-shop.herokuapp.com
Connection: Upgrade
Origin: http://example.com
Sec-Websocket-Key: hpydpogayYbH54j8WWatHg==
Sec-Websocket-Version: 13
Upgrade: websocket


HTTP/1.1 101 Switching Protocols
Connection: Upgrade
Sec-Websocket-Accept: ZD8iR647ozR65gsrZpA30Mvcw/U=
Upgrade: websocket
Via: 1.1 vegur


true,wss://juice-shop.herokuapp.com/socket.io/
```

