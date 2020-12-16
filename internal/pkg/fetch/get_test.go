package fetch

import (
	"testing"

	"github.com/nbio/st"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_Get_OK(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/example.txt").
		MatchHeader("User-Agent", "sentron/1.1").
		Reply(200).
		BodyString("hello!")

	body, err := Get("http://example.com/example.txt")
	assert.Equal(t, []byte("hello!"), body)
	assert.Nil(t, err)

	st.Expect(t, gock.IsDone(), true)
}

func Test_Get_NotOK(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/example.txt").
		MatchHeader("User-Agent", "sentron/1.1").
		Reply(404).
		BodyString("not found")

	body, err := Get("http://example.com/example.txt")
	assert.EqualErrorf(t, err, "Error fetching URL - 404", "Error message mismatch")
	assert.Nil(t, body)

	st.Expect(t, gock.IsDone(), true)
}
