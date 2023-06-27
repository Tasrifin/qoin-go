package helpers

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func CheckInput() (result int, err error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	result, err = strconv.Atoi(input)
	if err != nil {
		return
	}

	if result <= 0 {
		err = errors.New("input must be greather than 0")
		return
	}

	return
}

func GenerateRandomNumber() (result int) {
	randomInt := rand.Intn(6)
	result = randomInt + 1

	return
}
