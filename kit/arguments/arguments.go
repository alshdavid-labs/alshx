package arguments

import (
	"strconv"
	"strings"
)

type Arguments struct {
	values  map[string][]string
	command []string
}

func newArguments(argsList []string) *Arguments {
	a := &Arguments{
		values:  map[string][]string{},
		command: []string{},
	}

	for i := 0; i < len(argsList); i++ {
		completeSymbol := argsList[i]
		strippedSymbol := strings.Replace(completeSymbol, "-", "", 2)

		if strings.HasPrefix(completeSymbol, "-") {
			if a.values[strippedSymbol] == nil {
				a.values[strippedSymbol] = []string{}
			}

			if len(argsList) <= i+1 || strings.HasPrefix(argsList[i+1], "-") {
				a.values[strippedSymbol] = append(a.values[strippedSymbol], "true")
				continue
			}

			if len(argsList) >= i+1 {
				a.values[strippedSymbol] = append(a.values[strippedSymbol], argsList[i+1])
				i += 1
				continue
			}
		}

		a.command = append(a.command, strippedSymbol)
	}

	return a
}

func (a *Arguments) GetStrings(names ...string) []string {
	l := []string{}

	for _, name := range names {
		f := a.values[name]
		if f != nil {
			l = append(l, f...)
		}
	}

	if l == nil {
		return []string{}
	}
	return l
}

func (a *Arguments) GetString(name string) string {
	l := a.GetStrings(name)
	if len(l) < 1 {
		return ""
	}
	return l[0]
}

func (a *Arguments) GetNumbers(names ...string) []int64 {
	ls := []string{}

	for _, name := range names {
		f := a.values[name]
		if f != nil {
			ls = append(ls, f...)
		}
	}

	if len(ls) == 0 {
		return []int64{}
	}

	o := []int64{}

	for _, v := range ls {
		v, _ := strconv.ParseInt(v, 10, 64)
		o = append(o, v)
	}

	return o
}

func (a *Arguments) GetNumber(name string) int64 {
	l := a.GetNumbers(name)
	if len(l) < 1 {
		return 0
	}
	return l[0]
}

func (a *Arguments) GetBoolean(names ...string) bool {
	l := []string{}

	for _, name := range names {
		f := a.values[name]
		if f != nil {
			l = append(l, f...)
		}
	}

	if len(l) == 0 {
		return false
	}

	return l[0] == "true"
}
