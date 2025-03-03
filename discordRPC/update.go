package discordRPC

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func Update(mapName string, serverName string, startTime time.Time) {
	err := client.SetActivity(client.Activity{
		State:      "Map: " + mapName,
		Details:    "Server: " + serverName,
		LargeImage: "cssicon",
		LargeText:  "Counter Strike: Source",
		Timestamps: &client.Timestamps{
			Start: &startTime,
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}
