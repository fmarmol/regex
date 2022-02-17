package regex

import (
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func TestFindGroups(t *testing.T) {
	source := `(?P<one>\d+) (?P<two>.*) (?P<three>[0-9]+.[0-9]+)`

	groups := FindGroups(regexp.MustCompile(source), "123 foo 3.4")
	assert.Equal(t, "123", groups.MustGet("one"))
	assert.Equal(t, "foo", groups.MustGet("two"))
	assert.Equal(t, "3.4", groups.MustGet("three"))
	assert.Equal(t, 123, groups.MustGetAsInt("one"))
	assert.Equal(t, 3.4, groups.MustGetAsFloat("three"))
}
