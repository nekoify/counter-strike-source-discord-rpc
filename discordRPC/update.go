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
		LargeImage: "https://i.ibb.co/MVjZ877/kisspng-counter-strike-source-counter-strike-condition-z-steam-5aba1f45424574-4628449515221471412715.png",
		LargeText:  "Counter Strike: Source",
		Timestamps: &client.Timestamps{
			Start: &startTime,
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}
