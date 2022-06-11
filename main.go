package main

import (
  "github.com/demarcot/btea-example/app"

  tea "github.com/charmbracelet/bubbletea"
)

func main() {
  appModel := app.NewAppModel()
  appProg := tea.NewProgram(appModel, tea.WithAltScreen())
  appProg.Start()
}
