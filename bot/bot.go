package bot

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/katre/techdice/dice"
	"github.com/katre/techdice/parser"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	prefix  string
	parser  *parser.Parser
	botId   string
	session *discordgo.Session
}

func New(seed int64, botId string, prefix string) *Bot {
	return &Bot{
		prefix:  prefix,
		parser:  parser.New(dice.New(seed), prefix),
		botId:   botId,
		session: nil,
	}
}

func (b *Bot) createRouter() *exrouter.Route {
	router := exrouter.New()

	router.On(b.prefix, b.handleTechDice).Desc("roll dice for Technoir")

	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var text = ""
		for _, v := range router.Routes {
			text += v.Name + " : \t" + v.Description + "\n"
		}
		ctx.Reply("```" + text + "```")
	}).Desc("prints this help menu")

	return router
}

func (b *Bot) Start() error {
	// Start the connection and create a router
	session, err := discordgo.New("Bot " + b.botId)
	if err != nil {
		return err
	}
	b.session = session

	router := b.createRouter()

	b.session.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		router.FindAndExecute(b.session, "!", b.session.State.User.ID, m.Message)
	})

	err = b.session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Close() {
	b.session.Close()
}

func (b *Bot) handleTechDice(ctx *exrouter.Context) {
	input := strings.Join(ctx.Args, " ")
	//ctx.Reply("input: " + input)
	log.Printf("Received input: %q", input)
	result, err := b.parser.Roll(input)
	if err != nil {
		ctx.Reply("You said: " + input + ", which I didn't understand: " + err.Error())
		log.Printf("Parse error: %v", err)
		return
	}
	ctx.Reply("Result: " + describe(result))
	log.Printf("Response: %s", describe(result))
}

func describe(result dice.Result) string {
	var b strings.Builder

	// Print the score.
	fmt.Fprintf(&b, "Score: %s", result.Score)

	// Print the dice.
	fmt.Fprint(&b, " [")
	invalid := make(map[int]bool)
	for _, val := range result.HurtDice {
		invalid[val] = true
	}

	describeDice(&b, result.VerbDice, invalid)
	if len(result.PushDice) != 0 {
		fmt.Fprint(&b, ", push: ")
		describeDice(&b, result.PushDice, invalid)
	}
	if len(result.ManaDice) != 0 {
		fmt.Fprint(&b, ", mana: ")
		describeDice(&b, result.ManaDice, invalid)
	}
	// hurt
	if len(result.HurtDice) != 0 {
		fmt.Fprint(&b, ", hurt: ")
		describeDice(&b, result.HurtDice, map[int]bool{})
	}
	fmt.Fprint(&b, "]")

	return b.String()
}

func describeDice(w io.Writer, dice []int, invalid map[int]bool) {
	for i, die := range dice {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		template := "%d"
		if invalid[die] {
			template = "~~%d~~"
		}
		fmt.Fprintf(w, template, die)
	}
}
