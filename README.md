Status: in progress

```
package main

import (
	i "github.com/tockins/realize-examples/interact"
	"github.com/fatih/color"
	"fmt"
)

func main() {

	type A struct {
		Q1 bool
		Q2 string
		Q3 int
	}

	a := A{}
	b := color.New(color.FgHiWhite).Add(color.BgRed).SprintfFunc()
	//h := color.New(color.FgHiWhite).Add(color.BgWhite).SprintfFunc()
	y := color.New(color.FgYellow).SprintFunc()
	g := color.New(color.FgGreen).SprintFunc()

	i.Run(&i.Interact{
		Prefix: i.Prefix{W:color.Output,T:y("[")+"REALIZE"+y("]")},
		Questions: []*i.Quest{
			{
				Q: &i.Q{
					Text:  "Would you like some coffee? "+g("[yes/no]"),
					Err: b("INVALID"),
					Response: a.Q1,
					Default: i.D{Text:y("(yes)"),Value:bool(true)},
				},
				Validate:func(c *i.Context) interface{}{
					return nil
					//return h("INVALID INPUT")
				},
			},
		},
	})

	fmt.Println(a.Q1)
}
