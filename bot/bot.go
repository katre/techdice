package bot

import (
	"strings"

	"github.com/katre/techdice/dice"
	"github.com/katre/techdice/parser"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	parser  *parser.Parser
	botId   string
	session *discordgo.Session
}

func New(seed int64, botId string) *Bot {
	return &Bot{
		parser:  parser.New(dice.New(seed)),
		botId:   botId,
		session: nil,
	}
}

func (b *Bot) createRouter() *exrouter.Route {
	router := exrouter.New()

	router.On("techdice", b.handleTechDice).Desc("roll dice for Technoir")

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
	ctx.Reply("input: " + input)
}
