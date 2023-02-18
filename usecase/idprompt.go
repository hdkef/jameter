package usecase

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prompt struct{}

func (p *Prompt) GetReqIDSlice() (idSlice []int, msg string, valid bool) {
	fmt.Print("\nInput request ids (separated by space) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idStr := scanner.Text()

	for _, v := range strings.Split(idStr, " ") {
		id, err := strconv.Atoi(v)
		if err != nil {
			msg = "Invalid input"
			return
		}
		idSlice = append(idSlice, id)
	}

	return idSlice, "", true
}
