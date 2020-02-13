// tgbot-go -
// https://github.com/modern-dev/tgbot-go
// Copyright (c) 2020 Bohdan Shtepan
// Licensed under the MIT license.

package tgbot

// User object represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot
	Id int `json:"id"`
	// True, if this user is a bot
	IsBot bool `json:"is_bot"`
	// User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name,omitempty"`
	// Optional. User‘s or bot’s username
	Username string `json:"username,omitempty"`
	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`
	// Optional. True, if the bot can be invited to groups.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// Optional. True, if privacy mode is disabled for the bot.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// Optional. True, if the bot supports inline queries.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// Message represents a message.
type Message struct {
	// Unique message identifier inside this chat
	MessageId int `json:"message_id"`
	// Optional. Sender, empty for messages sent to channels
	From *User `json:"from,omitempty"`
	// Date the message was sent in Unix time
	Date int `json:"date"`
	// Conversation the message belongs to
	Chat *Chat `json:"chat"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For messages forwarded from channels, information about the original channel
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardFromMessageId int `json:"forward_from_message_id,omitempty"`
	// Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSignature string `json:"forward_signature,omitempty"`
	// Optional. Sender's name for messages forwarded from users who disallow adding a link
	// to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date,omitempty"`
	// Optional. For replies, the original message. Note that the Message object in this field
	// will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`
	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupId string `json:"media_group_id,omitempty"`
	// Optional. Signature of the post author for messages in channels
	AuthorSignature string `json:"author_signature,omitempty"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames, URLs, bot commands,
	// etc. that appear in the text
	Entities *[]MessageEntity `json:"entities,omitempty"`
	// Optional. For messages with a caption, special entities like usernames, URLs, bot commands,
	// etc. that appear in the caption
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio"`
	// Optional. Message is a general file, information about the file
	Document *Document `json:"document"`
	// Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set
	Animation *Animation `json:"animation"`
	// Optional. Message is a game, information about the game.
	Game *Game `json:"game"`
	// Optional. Message is a photo, available sizes of the photo
	Photo *[]PhotoSize `json:"photo"`
	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker"`
	// Optional. Message is a video, information about the video
	Video *Video `json:"video"`
	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice"`
	// Optional. Message is a video note, information about the video message
	VideoNoe *VideoNote `json:"video_noe"`
	// Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Caption string `json:"caption"`
	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact"`
	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`
	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue"`
	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll"`
	// Optional. New members that were added to the group or supergroup and information about them
	// (the bot itself may be one of these members)
	NewChatMembers *[]User `json:"new_chat_members"`
	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member"`
	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`
	// Optional. A chat photo was change to this value
	NewChatPhoto *[]PhotoSize `json:"new_chat_photo"`
	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`
	// Optional. Service message: the supergroup has been created. This field can‘t be received in a message
	// coming through updates, because bot can’t be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message
	// in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	// Optional. Service message: the channel has been created. This field can‘t be received in a message
	// coming through updates, because bot can’t be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created"`
	// Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may be greater than 32 bits and some programming languages may have difficulty/silent defects
	// in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	MigrateToChatId int64 `json:"migrate_to_chat_id"`
	// Optional. The supergroup has been migrated from a group with the specified identifier.
	// This number may be greater than 32 bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	MigrateFromChatId int64 `json:"migrate_from_chat_id"`
	// Optional. Specified message was pinned. Note that the Message object in this field
	// will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message"`
	// Optional. Message is an invoice for a payment, information about the invoice.
	Invoice *Invoice `json:"invoice"`
	// Optional. Message is a service message about a successful payment, information about the payment.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`
	// Optional. The domain name of the website on which the user has logged in.
	ConnectedWebsite string `json:"connected_website"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// Optional. rue, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes,
	// implies CanSendMessages.
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// Optional. True, if the user is allowed to send polls, implies CanSendMessages.
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// Optional. True, if the user is allowed to send animations, games, stickers and use inline bots,
	//implies CanSendMediaMessages.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// Optional. True, if the user is allowed to add web page previews to their messages, implies CanSendMediaMessages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo.
	// This FileID can be used only for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo.
	// This FileId can be used only for photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

// Chat represents a chat.
type Chat struct {
	// Unique identifier for this chat.
	// This number may be greater than 32 bits and some programming languages may have difficulty/silent defects
	// in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	Id int64 `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”.
	Type string `json:"type"`
	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`
	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`
	// Optional. Chat photo. Returned only in GetChat.
	Photo *ChatPhoto `json:"photo,omitempty"`
	// Optional. Description, for groups, supergroups and channel chats. Returned only in GetChat.
	Description string `json:"description,omitempty"`
	// Optional. Chat invite link, for groups, supergroups and channel chats. Each administrator in a chat
	// generates their own invite links, so the bot must first generate the link using ExportChatInviteLink.
	// Returned only in GetChat.
	InviteLink string `json:"invite_link,omitempty"`
	// Optional. Pinned message, for groups, supergroups and channels. Returned only in GetChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Default chat member permissions, for groups and supergroups. Returned only in GetChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	// Optional. For supergroups, the minimum allowed delay between consecutive messages
	// sent by each unpriviledged user. Returned only in GetChat.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	// Optional. For supergroups, name of group sticker set. Returned only in GetChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// Optional. True, if the bot can change the group sticker set. Returned only in GetChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	// “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text),
	// “strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block),
	// “text_link” (for clickable text URLs), “text_mention” (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units.
	Length int `json:"length"`
	// Optional. For “text_link” only, url that will be opened after user taps on the text.
	Url string `json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`
	// Optional. For “pre” only, the programming language of the entity text.
	Language string `json:"language,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Photo width.
	Width int `json:"width"`
	// Photo height.
	Height int `json:"height"`
	// Optional. File size.
	FileSize int `json:"file_size,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
	// Video width as defined by sender.
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width as defined by sender.
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description int `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo *[]PhotoSize `json:"photo"`
	// Optional. Brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game when the bot calls setGameScore,
	// or manually edited using editMessageText. 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
	// Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
	Animation *Animation `json:"animation,omitempty"`
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Contact's user identifier in Telegram
	UserId int `json:"user_id,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float32 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float32 `json:"latitude"`
}

// Venue represents a venue.
type Venue struct {
	// Venue location
	Location *Location `json:"location"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// Unique poll identifier
	PollId string `json:"poll_id"`
	// The user, who changed the answer to the poll
	User *User `json:"user"`
	// 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
	OptionIds []int `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	// Unique poll identifier
	Id string `json:"id"`
	// Poll question, 1-255 characters
	Question string `json:"question"`
	// List of poll options
	Options *[]PollAnswer `json:"options"`
	// Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`
	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`
	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode,
	// which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionId int `json:"correct_option_id,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”,
	// “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float32 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}

// Sticker represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
	// Sticker width
	Width int `json:"width"`
	// Sticker height
	Height int `json:"height"`
	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated"`
	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs
	SetMame string `json:"set_mame,omitempty"`
	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	// Product name
	Title string `json:"title"`
	// Product description
	Description string `json:"description"`
	// Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// State, if applicable
	State string `json:"state"`
	// City
	City string `json:"city"`
	// First line for the address
	StreetLine1 string `json:"street_line_1"`
	// Second line for the address
	StreetLine2 string `json:"street_line_2"`
	// Address post code
	PostCode string `json:"post_code"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	// Optional. User name
	Name string `json:"name,omitempty"`
	// Optional. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User email
	Email string `json:"email,omitempty"`
	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// SuccessfulPayment
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	// Optional. Order info provided by the user
	OrderInfo OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}
