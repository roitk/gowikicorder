package main

import (
	"github.com/thoj/go-ircevent"
	"fmt"
)

var roomName = "#en.wikipedia"

func main() {
	con := irc.IRC("gowikicorder", "gowikicorder")
	err := con.Connect("irc.wikimedia.org:6667")
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}
	con.AddCallback("001", func (e *irc.Event) {
		con.Join(roomName)
	})
	con.AddCallback("PRIVMSG", func (e *irc.Event) {
		fmt.Println(e.Message())
	})
	con.Loop()
}
