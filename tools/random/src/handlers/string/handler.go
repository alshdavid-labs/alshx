package handler_string

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
		ret[i] = getRandomString()
	}

	fmt.Println(strings.Join(ret, ""))
}

func getRandomString() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		fmt.Println(colors.Error("Unable to generate number"))
		os.Exit(1)
	}
	return string(letters[num.Int64()])
}
