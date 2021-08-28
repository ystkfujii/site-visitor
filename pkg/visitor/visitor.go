package visitor

import (
	//"fmt"
	"context"
	"os"
	"log"


	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ystkfujii/site-visitor/pkg/gate"
)

type Visitor interface {
	Run(ctx context.Context) error
}

type SiteVisitor struct {
}

func NewSiteVisitor() Visitor {
	return &SiteVisitor{}
}


func (v *SiteVisitor)Run(ctx context.Context) error {

	httpGate := gate.NewHttpGate()

	url := "https://books.rakuten.co.jp/rb/16462859/?bkts=1"
	str := "new_addToCart" 

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	ok, _ := httpGate.DoesItContainStr(ctx, url, str)

	if ok {
		message := linebot.NewTextMessage("【報告】ps5が現在販売中です\n" + url)
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Fatal(err)
		}
	}


	return nil
}
