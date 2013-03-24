package main

import "github.com/codysoyland/go-talkfilters"
import "fmt"

func main() {
    fmt.Println(talkfilters.ListFilters())
    fmt.Println(talkfilters.ListFilters()["chef"]("This should look like a chef wrote it."))
    fmt.Println(talkfilters.RunFilter("jethro", "This should look like Jethro wrote it."))
    fmt.Println(talkfilters.GetFilter("pirate")("This should look like a pirate wrote it."))
}
