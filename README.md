# ably-wasm-demo
Compiling ably-go to wasm

The target go code to be compiled to web assembly can be found in the `/src` folder.

Two files need from your Go installation are required, they can be found at `$GOROOT/misc/wasm`
* `wasm_exec.html`
* `wasm_exec.js`

After copying these files to this project in `wasm_exec.html` I renamed `test.wasm` to `lib.wasm`
I also renamed the file itself from `wasm_exec.html` > `index.html`.

The Go code in `/src` is compiled into a Web Assembly file called `lib.wasm`.
To compile to wasm, the `GOARCH` env var needs to be changed to `wasm` and the `GOOS` env var needs to be changed to `js`.

Can compile `main.go` to Web Assembly with this command which sets these go env vars.
run this command from `cd src && GOARCH=wasm GOOS=js go build -o lib.wasm main.go` 

After compiling `main.go` to `lib.wasm`, move the `lib.wasm` to the project root folder. Then start the local server with `go run server.go`. 
Then navigate to `localhost:8080`
A html page with a run button should be served. 
Clicking on the run button will execute the wasm file in the browser.

# Observations
## ably-go v1.2.7 
This version of the ably-go SDK does compile to wasm.
Was able to compile some Go code that attempts to publish to an Ably Realtime channel.
An invalid key error was encountered when the Go code attempted to use a key stored in an env var.
After hardcoding the key and attempting to execute the compiled wasm in a browser the following error was seen in every few seconds in the browser console..
```
2022/07/02 11:40:55 [ERROR] Received recoverable error websocket.Dial wss://realtime.ably.io:443?echo=true&format=msgpack&heartbeats=true&key=<redacted>&timestamp=1656758455288&v=1.2: dial tcp: lookup realtime.ably.io: Protocol not available
```
Tried setting a client option to use json protocol instead of message pack protocol.
all that did was change the error message to 
```
2022/07/02 11:49:00 [ERROR] Received recoverable error websocket.Dial wss://realtime.ably.io:443?echo=true&format=json&heartbeats=true&key=<redacted>&timestamp=1656758940832&v=1.2: dial tcp: lookup realtime.ably.io: Protocol not available
```

## ably-go v1.2.8
In this newly released version of the SDK, the deprecated websocket package https://pkg.go.dev/golang.org/x/net/websocket was replaced with package https://pkg.go.dev/nhooyr.io/websocket I wanted to see if this changed allowed ably-go to compile to wasm and execute in a browser.

When attempting to compile the same Go code to wasm. The code failed to compile with this error.
```
C:\dev\go\pkg\mod\github.com\ably\ably-go@v1.2.8\ably\websocket.go:105:6: ops.HTTPHeader undefined (type websocket.DialOptions has no field or method HTTPHeader)
C:\dev\go\pkg\mod\github.com\ably\ably-go@v1.2.8\ably\websocket.go:106:6: ops.HTTPHeader undefined (type websocket.DialOptions has no field or method HTTPHeader)
```



