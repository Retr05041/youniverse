package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

// Runner for UI
func Run() {
	app := app.New()
	mainWindow := app.NewWindow("Youniverse")

	// Chat messages display
	chatDisplay := widget.NewMultiLineEntry()
	chatDisplay.SetPlaceHolder("Chat messages will appear here...")
	chatDisplay.Disable() // Make it read-only

	// User input field
	userInput := widget.NewEntry()
	userInput.SetPlaceHolder("Type your message...")

	// Send button
	sendButton := widget.NewButton("Send", func() {
		message := userInput.Text
		if message != "" {
			chatDisplay.SetText(chatDisplay.Text + "\n" + "You: " + message)
			userInput.SetText("") // Clear the input field
		}
	})

	// Create a horizontal box for input and button, making them fill the width
	inputRow := container.New(layout.NewGridLayout(2), userInput, sendButton)

	// Combine chat display and input row into a border layout
	content := container.NewBorder(
		nil,         // Top (no widget)
		inputRow,    // Bottom
		nil, nil,    // Left, Right
		chatDisplay, // Center
	)

	// Set up the window content and show
	mainWindow.SetContent(content)
	mainWindow.Resize(fyne.NewSize(400, 300))
	mainWindow.ShowAndRun()
}

