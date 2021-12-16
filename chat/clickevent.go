package chat

import "strconv"

// ClickEvent defines an event that occurs when this component is clicked.
type ClickEvent struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

// OpenURL create a ClickEvent opens the given URL in the default web browser.
// Ignored if the player has opted to disable links in chat;
// may open a GUI prompting the user if the setting for that is enabled.
// The link's protocol must be set and must be http or https, for security reasons.
func OpenURL(url string) *ClickEvent {
	return &ClickEvent{
		Action: "open_url",
		Value:  url,
	}
}

// RunCommand create a ClickEvent runs the given command. Not required to be a command -
// clicking this only causes the client to send the given content as a chat message,
// so if not prefixed with /, they will say the given text instead.
// If used in a book GUI, the GUI is closed after clicking.
func RunCommand(cmd string) *ClickEvent {
	return &ClickEvent{
		Action: "run_command",
		Value:  cmd,
	}
}

// SuggestCommand create a ClickEvent replaces the content of the chat box with the given text -
// usually a command, but it is not required to be a command
// (commands should be prefixed with /).
// This is only usable for messages in chat.
func SuggestCommand(cmd string) *ClickEvent {
	return &ClickEvent{
		Action: "suggest_command",
		Value:  cmd,
	}
}

// ChangePage create a ClickEvent usable within written books.
// Changes the page of the book to the given page, starting at 1.
// For instance, "value":1 switches the book to the first page.
// If the page is less than one or beyond the number of pages in the book, the event is ignored.
func ChangePage(page int) *ClickEvent {
	return &ClickEvent{
		Action: "change_page",
		Value:  strconv.Itoa(page),
	}
}

// CopyToClipboard create a ClickEvent copies the given text to the client's clipboard when clicked.
func CopyToClipboard(text string) *ClickEvent {
	return &ClickEvent{
		Action: "copy_to_clipboard",
		Value:  text,
	}
}
