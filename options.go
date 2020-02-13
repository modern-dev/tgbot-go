// tgbot-go
// https://github.com/modern-dev/tgbot-go
// Copyright (c) 2020 Bohdan Shtepan
// Licensed under the MIT license.

package tgbot

import "strconv"

type SendMessageOptions struct {
	DisableWebPagePreview bool
	DisableNotification   bool
	ReplyToMessageId      int
	ReplyMarkup           string
}

func (smo *SendMessageOptions) addOptions(params map[string]string) {
	if smo.DisableWebPagePreview {
		params["disable_web_page_preview"] = "true"
	}

	if smo.DisableNotification {
		params["disable_notification"] = "true"
	}

	if smo.ReplyToMessageId != 0 {
		params["reply_to_message_id"] = strconv.Itoa(smo.ReplyToMessageId)
	}

	if smo.ReplyMarkup != "" {
		params["reply_markup"] = smo.ReplyMarkup
	}
}
