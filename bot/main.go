package main

import (
	"flag"
	"time"

	"github.com/katre/techdice/bot"
)

var seed = flag.Int64("seed", time.Now().UnixNano(), "Seed the RNG")
var botId = flag.String("botId", "", "The Discord bot id")

func main() {
	//bot := bot.New(*seed, *botId)
	bot.New(*seed, *botId)

	// Start the bot.
}
