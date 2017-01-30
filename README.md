Status: in progress

```
package main

import (
	"github.com/fatih/color"
	i "github.com/tockins/realize-examples/interact"
	"fmt"
)

func main() {

	b := color.New(color.FgHiWhite).Add(color.BgRed).SprintfFunc()
	y := color.New(color.FgYellow).SprintFunc()
	g := color.New(color.FgGreen).SprintFunc()
	prefix := y("[") + "REALIZE" + y("]")

	i.Run(&i.Interact{
		Before: func(c *i.Context) error{
			c.Prefix(color.Output, prefix)
			return nil
		},
		Questions: []*i.Question{
			{
				Before: func(c *i.Context) error{
					return nil
				},
				Quest: &i.Quest{
					Text:     "Would you like some coffee?",
					Options:  g("[yes/no]"),
					Err:      b("INVALID"),
					Response: bool(false),
					Default:  i.Default{Text: y("(yes)"), Status: true},
				},
				Action: func(c *i.Context) interface{} {
					//fmt.Println(c.Response().(bool))
					return nil
					//return h("INVALID INPUT")
				},
				After: func(c *i.Context) error{
					return nil
				},
			},
		},
		After: func(c *i.Context) error{
			return nil
		},
	})

	i.Run(&i.Question{
		Before: func(c *i.Context) error{
			return nil
		},
		Quest: &i.Quest{
			Text:     "What Kind of Coffee?",
			Err:      b("INVALID"),
			Response: string("none"),
			Default:  i.Default{Text: y("(none)"), Status: true},
			Choices: i.Choices{
				Prefix: "\t",
				Color: g,
				Alternatives: []i.Choice{
					{
						Text: "Black coffee",
						Response: "black",
					},
					{
						Text: "With milk",
						Response: "milk",
					},
				},
			},
		},
		Action: func(c *i.Context) interface{} {
			fmt.Println(c.Response().(string))
			return nil
			//return h("INVALID INPUT")
		},
		After: func(c *i.Context) error{
			return nil
		},
	})
}

context diventa un interfaccia e tutte le volte durante il ciclo di quest viene cambiata passa da interact a quest a choice etc
tutte le funzioni di context diventano quindi metodi non collegati ad una struct ma aventi come argomento l'interfaccia stessa. 
metodi comuni a tutti e tre i tipi di elementi
prefix
response
etc
