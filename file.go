// tgbot-go -
// https://github.com/modern-dev/tgbot-go
// Copyright (c) 2020 Bohdan Shtepan
// Licensed under the MIT license.

package tgbot

type InputFile struct {
	FileId string
	FileURL string
	FilePath string
}

func InputFileFromURL(url string) InputFile {
	return InputFile{FileURL:url}
}

func InputFileFromDisk(path string) InputFile {
	return InputFile{FilePath:path}
}

func (f *InputFile) IsOnDisk() bool {
	return f.FilePath != ""
}
