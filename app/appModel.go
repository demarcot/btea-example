package app

import (
  "github.com/demarcot/btea-example/thing"
  "github.com/demarcot/btea-example/altThing"

  tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
  State string
  thing thing.ThingModel
  altThing altThing.AltThingModel
}

func NewAppModel() tea.Model {
  m := AppModel{ State: "thing", thing: thing.NewThingModel(), altThing: altThing.NewAltThingModel() }
  return m
}

func (m AppModel) Init() tea.Cmd {
  return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd

  switch msg := msg.(type) {
    default:
      break

    case tea.KeyMsg:
      switch msg.String() {
        case "ctrl+c", "q":
          return m, tea.Quit
        case "t":
          if m.State == "thing" {
            m.State = "altThing"
            _, cmd := m.thing.Update(msg)
            cmds = append(cmds, cmd)
          } else {
            m.State = "thing"
            _, cmd := m.altThing.Update(msg)
            cmds = append(cmds, cmd)
          }
          break
      }
  }
  return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
  s := "--- App (t: toggle view, q: quit) ---\n\n"
  if m.State == "thing" {
    s += m.thing.View()
  } else {
    s += m.altThing.View()
  }
  return s
}
