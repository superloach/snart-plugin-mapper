package mapper

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/route"
)

func Map(ctx *route.Ctx) error {
	_f := "Map"

	_ = ctx.Flags.Parse()

	args := ctx.Flags.Args()
	query := strings.Join(args, " ")
	queries := strings.Split(query, "+")
	nqueries := make([]string, 0)

	for _, q := range queries {
		q = strings.TrimSpace(q)
		if len(q) == 0 {
			continue
		}

		nqueries = append(nqueries, q)
	}

	if len(nqueries) == 0 {
		rep1 := ctx.Reply()
		rep1.Content = "please specify a query.\nex: `" +
			ctx.CleanPrefix + ctx.Route.Name + " name of place`"

		return rep1.Send()
	}

	msg := "Map given for"

	for _, query := range nqueries {
		err := ctx.Session.ChannelTyping(ctx.Message.ChannelID)
		if err != nil {
			err = fmt.Errorf("typing %#v: %w", ctx.Message.ChannelID, err)
			Log.Warn(_f, err)
		}

		rep := ctx.Reply()
		rep.Embed = &dg.MessageEmbed{
			Title: query,
			URL:   mapURL(query),
			Footer: &dg.MessageEmbedFooter{
				Text: fmt.Sprintf(
					"%s %s",
					msg, nick(ctx.Message),
				),
			},
		}

		err = rep.Send()
		if err != nil {
			Log.Warn(_f, err)
		}
	}

	return nil
}
