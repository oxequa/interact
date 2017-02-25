### Interact

An easy and fast Go library, without external imports, to handle questions and answers by command line

##### Features

- Single question 
- Questions list
- Multiple choices 
- Sub questions
- Questions prefix
- Questions default values
- Custom errors 
- After/Before listeners
- Colors support (fatih/color)
- Windows support
- Run Single/Sequence/One by one 

##### Installation

To install interact:
```
$ go get github.com/tockins/interact
```

##### Getting Started - Single question

Run a simple question and manage the response. 
The response field is used to get the answer as a specific type.
```
package main

func main() {
    i.Run(&i.Question{
    		Quest: i.Quest{
    			Msg:      "Would you like some coffee?",
    			Err:      "INVALID INPUT",
    		},
    		Action: func(c i.Context) interface{} {
    			fmt.Println(c.Input().Bool())
    			return nil
    		}
    	})
}
``` 

##### Questions list

Define a list of questions to be run in sequence.
The Action func can be used for validate the answer and can return a custom error.

```
package main

func main() {
	i.Run(&i.Interact{
		Questions: []*i.Question{
			{
				Quest: i.Quest{
					Msg:     "Would you like some coffee?",
					Err:      "INVALID INPUT",
				},
				Action: func(c i.Context) interface{} {
					fmt.Println(c.Answer().Bool())
					return nil
				},
			},
			{
                Quest: i.Quest{
                    Msg:     "What's 2+2?",
                    Err:      "INVALID INPUT",
                },
                Action: func(c i.Context) interface{} {
                    // get the answer as integer
                    if c.Answer().Int() < 4 {
                        // return a custom error and rerun the question
                        return "INCREASE"
                    }else if c.Answer().Int() > 4 {
                        return "DECREASE"
                    }
                    return nil
                },
            },
		},
	})
}
```

##### Multiple choice

Define a multiple choice question


```
package main

func main() {
	i.Run(&i.Interact{
		Questions: []*i.Question{
			{
				Quest: i.Quest{
                    Msg:     "how much for a teacup?",
                    Err:     "INVALID INPUT",
                    Choices: i.Choices{
                        Alternatives: []i.Choice{
                            {
                                Text: "Gyokuro teapcup",
                                Response: "20",
                            },
                            {
                                Text: "Sencha teacup",
                                Response: -10,
                            },
                            {
                                Text: "Matcha teacup",
                                Response: 15.50,
                            },
                        },
                    },
                },
                Action: func(c i.Context) interface{} {
                    fmt.Println(c.Answer().Int())
                    return nil
                },
			},
		},
	})
}
```


##### Sub questions
