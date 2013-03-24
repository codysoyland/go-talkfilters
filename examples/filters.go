package main

import "github.com/codysoyland/go-talkfilters"
import "fmt"

func main() {
    phrase := "Can anybody think of a use for this ridiculous library?"
    fmt.Println(talkfilters.Austro(phrase))
    fmt.Println(talkfilters.B1ff(phrase))
    fmt.Println(talkfilters.Brooklyn(phrase))
    fmt.Println(talkfilters.Chef(phrase))
    fmt.Println(talkfilters.Cockney(phrase))
    fmt.Println(talkfilters.Drawl(phrase))
    fmt.Println(talkfilters.Dubya(phrase))
    fmt.Println(talkfilters.Fudd(phrase))
    fmt.Println(talkfilters.Funetak(phrase))
    fmt.Println(talkfilters.Jethro(phrase))
    fmt.Println(talkfilters.Jive(phrase))
    fmt.Println(talkfilters.Kraut(phrase))
    fmt.Println(talkfilters.Pansy(phrase))
    fmt.Println(talkfilters.Pirate(phrase))
    fmt.Println(talkfilters.Postmodern(phrase))
    fmt.Println(talkfilters.Redneck(phrase))
    fmt.Println(talkfilters.Valspeak(phrase))
    fmt.Println(talkfilters.Warez(phrase))
}
