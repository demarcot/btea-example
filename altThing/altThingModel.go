package altThing 

import tea "github.com/charmbracelet/bubbletea"

type AltThingModel struct {
  Val string
}

func NewAltThingModel() AltThingModel {
  m := AltThingModel{ Val: "AltThing" }
  return m
}

func (m AltThingModel) Init() tea.Cmd {
  return nil
}

func (m AltThingModel) Update(tea.Msg) (AltThingModel, tea.Cmd) {
  return m, nil
}

func (m AltThingModel) View() string {
  s := ""
  s += m.Val + "\n"
  return s
}
