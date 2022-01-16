package regex

import (
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func TestFindGroups(t *testing.T) {
	source := `(?P<one>\d+) (?P<two>.*)`

	groups := FindGroups(regexp.MustCompile(source), "123 foo")
	assert.Equal(t, "123", groups.MustGet("one"))
	assert.Equal(t, "foo", groups.MustGet("two"))
	assert.Equal(t, 123, groups.MustGetAsInt("one"))
}
