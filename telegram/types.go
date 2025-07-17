/* telegram.go : Telegram Bot API Wrapper
 *
 * Copyright (c) 2025 Paolo Giordano
 * Licensed undert the MIT License. See the LICENSE file for more details.
 *
 * (See the Golang conventions on documentation)
 * Goal of this project:
 * - learn Golang
 * - learn Telegram API
 * - create a self-documented library with extensive explanations of the Telegram API
 *   (maybe clarify the obscure or bad documentation)
 *   (people that write documentations hate humanity)
 */

package telegram

// The Bot API sends an Update struct, which contains various nested structs.
// We are building these complex structs, like Update and Message, from their fundamental components such as User and Chat.

// This struct can represent both a user and a bot
type User struct {
	// Unique identifier for this user or bot
	ID int64

	// True if the user is a bot
	IsBot bool

	// User's or bot's first name
	FirstName string

	// [Optional] User's or bot's last name
	LastName string

	// [Optional] User's or bot's username
	Username string

	// [Optional] IETF language tag of the user's language
	LanguageCode string

	// [Optional] True if this user is a Telegram Premium user
	IsPremium bool

	// [Optional] True if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool

	// [Optional] True if the bot can be invited to groups. Returned only in GetMe

	CanJoinGroups bool

	// [Optional] True if privacy mode is disabled for the bot. Returned only in GetMe
	CanReadAllGroupMessages bool

	// [Optional] True if the bot supports inline queries. Returned only in GetMe
	SupportsInlineQueries bool

	// [Optional] True if the bot can be connected to a telegram business account to receive its messages. Returned only in GetMe
	CanConnectToBusiness bool

	// [Optional] True if the bot has a main Web App. Returned only in GetMe
	HasMainWebApp bool     
}

// This struct represents a chat
type Chat struct {
	// Unique identifier for this chat
	ID int64

	// Type of the chat. Can be either "private", "group", "supergroup" or "channel"
	Type string

	// [Optional] Title; for supergroups, channels and group chats
	Title string

	// [Optional] Username, for private chats, supergroups and channels if available
	Username string

	// [Optional] First name of the other party in a private chat
	FirstName string

	// [Optional] Last name of the other party in a private chat
	LastName string

	// [Optional] True if the supergroup chat is a forum (has topics enabled)
	IsForum bool
}

// This type represents a unique message identifier
type MessageID struct {
	// In specific instances (e.g., message containing a video sento to a big chat),
	// the server might automatically schedule a message instead of sending it immediately.
	// In such cases, this field will be 0 and the relevant message will be unusable until
	// it is actually sent
	MessageID int64
}

// This struct descrives a message that was deleted or it is otherwise inaccessible to the bot
type InaccessibleMessage struct {
	// Chat the message belonged to
	Chat Chat

	// Unique message identifier inside the chat
	MessageID int64

	// Always zero. The field can be used to differentiate regular and inaccessible messages
	Date int64
}

// MaybeInaccessibleMessage
// This is a Union of the types Message e InaccessibleMessage. Does golang have unions?

// This struct represents one special entity in a text message.
// For examples, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Currently, can be "mention" (@username),
	// "hashtag" (#hashtag or #hashtag@chatusername),
	// "cashtag" ($USD or $USD@chatusername),
	// "bot_command" (/start@jobs_bot),
	// "url" (https://telegram.org),
	// "email" (do-not-reply@telegram.org),
	// "phone_number" (+1-212-555-0123),
	// "bold", "italic", "underline", "strikethrough", "spoiler", "blockquote",
	// "expandable_blockquote" (collapsed-by-default block quotation), "code",
	// "pre" (monowidth block), "text_link" (for clickable text URLs),
	// "text_mention" (for users withous username),
	// "custom_emoji" (for inline custom emoji stickers)
	Type string

	// Offset in UTF-16 code units to the start of the entity
	Offset int64

	// Length of the entity in UTF-16 code units
	Length int64

	// [Optional] For "text_link" only, URL that will be opened after user taps on the text
	URL string

	// [Optional] For "text_mention" only, the mentioned user
	User User

	// [Optional] For "pre" only, the programming language of the entity text
	Language string

	// [Optional] For "custom_emoji" only, unique identifier of the custom emoji. Use GetCustomEmojiStickers to get full information about the sticker
	CustomEmojiID string
}

// This struct contains information about the quoted part of a message that is replied to by the given message
type TextQuote struct {
	// Text of the quoted part of a message that is replied to by the given message
	Text string

	// [Optional] Special entities that appear in the quote. Currently, only
	// "bold", "italic", "underline", "strikethrough", "spoiler", and "custom_emoji"
	// entities are kept in quotes
	Entities []MessageEntity

	// Approximate quote position in the original message in UTF-16 code units as specified by the sender
	Position int64

	// [Optional] True if the quote was chisen manually by the message sender.
	// Otherwise, the quote was added automatically by the server
	IsManual bool
}

// This struct contains parameters for the message that is being sent
type ReplyParameters struct {
	// Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	MessageID int64

	// [Optional] If the message to be replied to is from a different chat
	// unique identifier for the chat or username of the channel (in the format
	// @channelusername). Not supported for messages sent on behalf of a business
	// account
	// Should be Int or String, but Golang doesn't have union
	// For now we use string, but I have to think carefully to this
	ChatID string

	// [Optional] Pass True if the message should be sent
	// even if the specified message to be replied to is not found.
	// Always False for replies in another chat or forum topic.
	// Always True for messages sent on behalf of a business account
	AllowSendingWithoutReply bool

	// [Optional] Quoted part of the messages to be replied to; 0 - 1024
	// characters after entities parsing. The quote must be an exact substring of
	// the message to be replied to, including "bold", "italic", "underline",
	// "strikethrough", "spoiler", and "custom_emoji" entities. The message
	// will fail to send if the quote isn't found in the original message
	Quote string

	// [Optional] Mode for parsing entities in the quote. See formatting
	// options for more details
	QuoteParseMode string

	// [Optional] A JSON-serialized list of special entities that appear in
	// the quote. It can be specified instead of quote_parse_mode
	QuoteEntities []MessageEntity

	// [Optional] Position of the quote in the original message in UTF-16
	// code units
	QuotePosition Integer
}

// MessageOrigin, another "union".
// - MessageOriginUser
// - MessageOriginHiddenUser
// - MessageOriginChat
// - MessageOriginChannel

// This struct contains information in the case the message
// was originally sent by a known user
type MessageOriginUser struct {
	// Type of the message origin, always "user"
	Type string

	// Date the message was sent originally in Unix time
	Date int64

	// User that sent the message originally
	SenderUser User
}

// This struct contains information in the case the message
// was originally sent by an unknown user
type MessageOriginHiddenUser struct {
	// Type of the message origin, always "hidde_user"
	Type string

	// Date the messahe was sent originally in Unix time
	Date int64

	// Name of the user that sent the message originally
	SendUserName string
}