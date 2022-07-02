package main

import (
	"context"

	"github.com/ably/ably-go/ably"
)

func main() {

	ctx := context.Background()


	realtimeClient, err := ably.NewRealtime(
		ably.WithKey("insert your key here"),
	)
	if err != nil {
		panic(err)
	}



	newChannel := realtimeClient.Channels.Get("test")

	newChannel.Publish(ctx, "message", "Hello from wasm!")

	realtimeClient.Close()
}
