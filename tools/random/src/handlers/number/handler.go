package handler_number

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"random/src/platform/colors"
	"strings"
)

type args struct {
	length int
}

func Handler(argList []string) {
	arguments := parseArgs(argList)

	ret := make([]string, arguments.length)
	for i := 0; i < arguments.length; i++ {
		ret[i] = getRandomNumber()
	}

	for ret[0] == "0" {
		ret[0] = getRandomNumber()
	}

	fmt.Println(strings.Join(ret, ""))
}

func getRandomNumber() string {
	const letters = "1234567890"

	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		fmt.Println(colors.Error("Unable to generate number"))
		os.Exit(1)
	}
	return string(letters[num.Int64()])
}
