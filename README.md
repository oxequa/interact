### Interact

An easy and fast Go library, without external imports, to handle questions and answers by command line

##### Features

- [Single question](#single-question)
- [Questions list](#questions-list)
- [Multiple choices](#multiple-choice)
- [Sub questions](#sub-questions)
- [Question prefix](#question-prefix)
- [Default values](#default-values)
- [Custom errors](#custom-errors)
- After/Before listeners
- Colors support (fatih/color)
- Windows support
- Run Single/Sequence/One by one 

##### Installation

To install interact:
```
$ go get github.com/tockins/interact
```

##### Single question

Run a simple question and manage the response. 
The response field is used to get the answer as a specific type.
``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
    i.Run(&i.Question{
    		Quest: i.Quest{
    			Msg:      "Would you like some coffee?",
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

Question struct is only for single question whereas **Interact struct** supports multiple questions
``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
	i.Run(&i.Interact{
		Questions: []*i.Question{
			{
				Quest: i.Quest{
					Msg:     "Would you like some coffee?",
				},
				Action: func(c i.Context) interface{} {
					fmt.Println(c.Answer().Bool())
					return nil
				},
			},
			{
                Quest: i.Quest{
                    Msg:     "What's 2+2?",
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

``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
	i.Run(&i.Interact{
		Questions: []*i.Question{
			{
				Quest: i.Quest{
                    Msg:     "how much for a teacup?",
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

The sub questions list is managed by the **"Resolve"** func.
Each sub question can access to the parent answer by the **"Parent"** method

``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
    i.Run(&i.Question{
        Quest: i.Quest{
            Msg:     "Would you like some coffee?",
            Resolve: func(c i.Context) bool{
                return c.Answer().Bool()
            },
        },
        Subs: []*i.Question{
            {
                Quest: i.Quest{
                    Msg:     "What Kind of Coffee?",
                    Choices: i.Choices{
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
                Action: func(c i.Context) interface{} {
                    fmt.Println(c.Answer().String())
                    fmt.Println(c.Parent().Answer().String())
                    return nil
                },
            },
        },
        Action: func(c i.Context) interface{} {
            fmt.Println(c.Answer().String())
            return nil
        },
    })
}
```

##### Question prefix

Interact support a custom prefix for each question

You can define a **global prefix** for all questions but you can **overwrite it** in each question with ease

With the first param you can pass a custom **io.writer** instance

``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
	i.Run(&i.Interact{
		Before: func(c i.Context) error{
			c.Prefix(nil,"GLOBAL PREFIX")
			return nil
		},
		Questions: []*i.Question{
			{
				Before: func(c i.Context) error{
					c.Prefix(nil,"OVERWRITTEN PREFIX")
					return nil
				},
				Quest: i.Quest{
					Msg:     "Would you like some coffee?",
				},
				Action: func(c i.Context) interface{} {
					return nil
				},
			},
			{
				Quest: i.Quest{
					Msg:     "What's 2+2?",
				},
				Action: func(c i.Context) interface{} {
					return nil
				},
			},
		},
	})
}
```

##### Default values

You can define a default value for each question and get it in the action func as an answer

``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {
    i.Run(&i.Question{
        Quest: i.Quest{
            Msg:     "Would you like some coffee?",
            Default: i.Default{Value:"test val",Preview:false},
        },
        Action: func(c i.Context) interface{} {
            fmt.Println(c.Answer().String())
            return nil
        },
    })
}
``` 

##### Custom errors

You can define a default error for each questions, or you can return a custom error in specific cases


``` go
package main

import (
	i "github.com/tockins/interact"
)

func main() {

	i.Run(&i.Question{
		Quest: i.Quest{
			Msg: "Would you like some coffee?",
			Err: "Default error",
		},
		Action: func(c i.Context) interface{} {
			if(c.Answer().Bool() == false){
				return "Invalid answer"
			}
			return nil
		},
	})
}
```