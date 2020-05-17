package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/katre/techdice/bot"
)

var seed = flag.Int64("seed", time.Now().UnixNano(), "Seed the RNG")
var botId = flag.String("botId", "", "The Discord bot id")

func main() {
	if *botId == "" {
		fmt.Println("Usage: techdice --botId id")
		os.Exit(1)
	}

	bot := bot.New(*seed, *botId)

	// Start the bot.
	err := bot.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the bot.
	bot.Close()
}
