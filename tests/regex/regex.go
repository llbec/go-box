package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	// Denominations can be 3 ~ 16 characters long.
	reDnm     = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt     = `[[:digit:]]+`
	reDecAmt  = `[[:digit:]]*\.[[:digit:]]+`
	reSpc     = `[[:space:]]*`
	reCoin    = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
	reDecCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reDecAmt, reSpc, reDnm))
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("arg is the text to match the regex")
		return
	}

	coinStr := strings.TrimSpace(os.Args[1])
	matches := reCoin.FindStringSubmatch(coinStr)
	fmt.Print(matches)
}
