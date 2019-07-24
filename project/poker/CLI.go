package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore PlayerStore
	Input       io.Reader
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		Input:       bufio.NewReader(in),
	}
}
func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.Input)
	reader.Scan()
	//input, _ := ioutil.ReadAll(c.in)
	winner, err := extractWinner(reader.Text())
	if err != nil {
		fmt.Printf("Error extracting winner. %v", err.Error())
		return
	}

	c.PlayerStore.RecordWin(winner)
}

func extractWinner(in string) (winner string, err error) {

	tokens := strings.Split(string(in), " ")

	if len(tokens) != 2 || tokens[1] != "wins" {
		errMessage := fmt.Sprintf("Wrong format. got %q want 'playername wins'\n", in)
		return "", errors.New(errMessage)
	}

	tokens[1] = strings.TrimSpace(tokens[1])
	tokens[1] = strings.ToLower(tokens[1])

	if tokens[1] != "wins" || len(tokens) != 2 {
		errMessage := fmt.Sprintf("Wrong format. got %q want 'playername wins'\n", in)
		return "", errors.New(errMessage)
	}

	return tokens[0], nil
}
