Status: in progress

```
package main

import (
	i "github.com/tockins/realize-examples/interact"
	"github.com/fatih/color"
)

func main() {

	b := color.New(color.FgHiWhite).Add(color.BgRed).SprintfFunc()
	y := color.New(color.FgYellow).SprintFunc()
	g := color.New(color.FgGreen).SprintFunc()

	i.Start(&i.Quest{
		Q: &i.Q{
			Text:  "Would you like some coffee? "+g("[yes/no]"),
			Value: y("(no)"),
			Error: b("INVALID"),
			Response: bool(false),
			Default: true,
		},
	})
}
``` 
    
    
    
```
package main

import (
	i "github.com/tockins/realize-examples/interact"
	"github.com/fatih/color"
)

func main() {

	b := color.New(color.FgHiWhite).Add(color.BgRed).SprintfFunc()
	y := color.New(color.FgYellow).SprintFunc()
	g := color.New(color.FgGreen).SprintFunc()

	
	c := i.Interact{
		Prefix: y("[")+"REALIZE"+y("]"),
		Questions: []*i.Quest{
			{
				Q: &i.Q{
					Text:  "Would you like some coffee? "+g("[yes/no]"),
					Value: y("(no)"),
					Error: b("INVALID"),
					Response: bool(false),
					Default: true,
				},
			},
		},
	}
	i.Start(&c)
}

    ``` 