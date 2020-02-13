// tgbot-go
// https://github.com/modern-dev/tgbot-go
// Copyright (c) 2020 Bohdan Shtepan
// Licensed under the MIT license.

package tgbot

type Bot struct {
	Token string
	Me *User
}

func NewBot(token string) (*Bot, error) {
	bot := &Bot{
		Token: token,
	}

	user, err := bot.GetMe()

	if err != nil {
		return nil, err
	}

	bot.Me = user

	return bot, nil
}
