package second

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsEmptyWhenNoAnswer(t *testing.T) {
	assert.Equal(t, Rearrange("aaa"), "")
	assert.Equal(t, Rearrange("aaaaaaaaaaaaaaaaaab"), "")
	assert.Equal(t, Rearrange("aaaaaabbcc"), "")
	assert.Equal(t, Rearrange(""), "")
}

func TestReturnsValidStringWhenHasAnswer(t *testing.T) {
	assert.Equal(t, Rearrange("aab"), "aba")
	assert.Equal(t, Rearrange("h"), "h")
	assert.Equal(t, Rearrange("abababababababaccdcdcdcddddddddzzzzzzxxxxxxxxx"), "dadadadadadbdbdbdbdbdbxbxzxzxzxzxzxzxcxcacacac")
	assert.Equal(t, Rearrange("aaacccbbb"), "abacacbcb")
}
