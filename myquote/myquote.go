//laget ny directory fordi funker ikke i den samme

package myquote

import (
"rsc.io/quote"
)

func Go() string {
goQuote := quote.Go()
return goQuote
}

func Glass() string {
glassQuote := quote.Glass()
return glassQuote
}

func Opt() string {
optQuote := quote.Opt()
return optQuote
}

func Hello() string {
helloQuote := quote.Hello()
return helloQuote
}