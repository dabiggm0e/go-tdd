package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	//input, _ := ioutil.ReadAll(c.in)
	winner, err := extractWinner(reader.Text())
	if err != nil {
		fmt.Errorf("Error extracting winner. %v", err.Error())
		return
	}

	c.playerStore.RecordWin(winner)
}

func extractWinner(in string) (winner string, err error) {

	tokens := strings.Split(string(in), " ")
	tokens[1] = strings.TrimSpace(tokens[1])
	tokens[1] = strings.ToLower(tokens[1])

	if tokens[1] != "wins" || len(tokens) != 2 {
		errMessage := fmt.Sprintf("Wrong format. got %q want 'playername wins'\n", in)
		return "", errors.New(errMessage)
	}

	return tokens[0], nil
}
