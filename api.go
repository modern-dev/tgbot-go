// tgbot-go -
// https://github.com/modern-dev/tgbot-go
// Copyright (c) 2020 Bohdan Shtepan
// Licensed under the MIT license.

package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ApiUrl = "https://api.telegram.org/bot%s/%s"
)

func (bot *Bot) makeRequest(method string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf(ApiUrl, bot.Token, method)

	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(payload); err != nil {
		return []byte{}, err
	}

	resp, err := http.Post(url, "application/json", &b)

	if err != nil {
		return []byte{}, err
	}

	resp.Close = true

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return respBytes, nil
}

func (bot *Bot) makeFileRequest(method, token, name, path string, params map[string]string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return []byte{}, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(name, filepath.Base(path))

	if err != nil {
		return []byte{}, err
	}

	if _, err = io.Copy(part, file); err != nil {
		return []byte{}, err
	}

	for field, value := range params {
		writer.WriteField(field, value)
	}

	if err = writer.Close(); err != nil {
		return []byte{}, err
	}

	url := fmt.Sprintf(ApiUrl, token, method)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return []byte{}, fmt.Errorf("internal server error")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return respBytes, nil
}

// GetMe is a simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	jsonResp, err := bot.makeRequest("getMe", nil)

	if err != nil {
		return nil, err
	}

	var resp struct {
		Ok          bool   `json:"ok"`
		Result      *User  `json:"result"`
		Description string `json:"description"`
	}

	if err = json.Unmarshal(jsonResp, &resp); err != nil {
		return nil, err
	}

	if resp.Ok {
		return resp.Result, nil
	}

	return nil, fmt.Errorf("tgbot: %s", resp.Description)
}

// SendMessage is to send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(chatId, text string, opts *SendMessageOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": chatId,
		"text":    text,
	}

	if opts != nil {
		opts.addOptions(params)
	}

	jsonResp, err := bot.makeRequest("sendMessage", params)

	if err != nil {
		return nil, err
	}

	var resp struct {
		Ok          bool     `json:"ok"`
		Result      *Message `json:"result"`
		Description string   `json:"description"`
	}

	if err = json.Unmarshal(jsonResp, &resp); err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, fmt.Errorf("tgbot: %s", resp.Description)
	}

	return resp.Result, nil
}

//func (bot *Bot) SendPhoto(chatID string, photo InputFile)
