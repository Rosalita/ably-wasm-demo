package main

import (
	"context"
	"os"

	"github.com/ably/ably-go/ably"
)

func main() {

	ctx := context.Background()

	key, _ := os.LookupEnv("ABLY_PRIVATE_KEY")

	realtimeClient, err := ably.NewRealtime(
		ably.WithKey(key),
		//ably.WithUseBinaryProtocol(false),
	)
	if err != nil {
		panic(err)
	}

	newChannel := realtimeClient.Channels.Get("test")

	newChannel.Publish(ctx, "message", "Hello from wasm!")
}
