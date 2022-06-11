package thing

import tea "github.com/charmbracelet/bubbletea"

type ThingModel struct {
  Val string
}

func NewThingModel() ThingModel {
  m := ThingModel{ Val: "Thing" }
  return m
}

func (m ThingModel) Init() tea.Cmd {
  return nil
}

func (m ThingModel) Update(tea.Msg) (ThingModel, tea.Cmd) {
  return m, nil
}

func (m ThingModel) View() string {
  s := ""
  s += m.Val + "\n"
  return s
}
