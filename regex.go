package regex

import (
	"fmt"
	"regexp"
	"strconv"
)

type Groups map[string]string

func (g Groups) Get(groupName string) (string, bool) {
	ret, ok := g[groupName]
	return ret, ok
}

func (g Groups) GetAsInt(groupName string) (int, bool, error) {
	ret, ok := g.Get(groupName)
	if !ok {
		return 0, ok, nil
	}
	res, err := strconv.Atoi(ret)
	return res, true, err

}

func (g Groups) MustGet(groupName string) string {
	ret, ok := g.Get(groupName)
	if !ok {
		panic(fmt.Errorf("Group name: %v not found.", groupName))
	}
	return ret
}

func (g Groups) MustGetAsInt(groupName string) int {
	ret, ok, err := g.GetAsInt(groupName)
	if !ok {
		panic(fmt.Errorf("Group name: %v not found.", groupName))
	}
	if err != nil {
		panic(fmt.Errorf("Group name: %v could not cast as integer: %v.", groupName, err))
	}
	return ret
}

func FindGroups(pattern *regexp.Regexp, s string) Groups {
	ret := map[string]string{}

	matches := pattern.FindStringSubmatch(s)
	if len(matches) == 0 {
		return ret
	}

	for _, groupName := range pattern.SubexpNames() {
		index := pattern.SubexpIndex(groupName)
		if index == -1 {
			continue
		}
		if matches[index] != "" {
			ret[groupName] = matches[index]
		}
	}
	return ret
}
