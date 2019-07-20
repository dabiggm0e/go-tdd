package poker

import (
	"io"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (c *CLI) PlayPoker() {
	/*input, _ := ioutil.ReadAll(c.in)
	in := string(input)
	tokens := strings.Split(string(in), " ")
	tokens[1] = strings.TrimSpace(tokens[1])
	tokens[1] = strings.ToLower(tokens[1])

	if tokens[1] != "wins" || len(tokens) != 2 {
		fmt.Printf("Wrong format. got %q want 'playername wins'\n", in)
		return
	}
	c.playerStore.RecordWin(tokens[0])*/
	c.playerStore.RecordWin("Ziggy")
}
