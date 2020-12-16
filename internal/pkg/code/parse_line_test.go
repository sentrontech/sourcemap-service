package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const LongLine = "This is a really long example of a line that will need to be shortened if this whole content parsing thing is going to work??"
const ShortLine = "This is a short line"

func Test_ParseLine_ShortLine(t *testing.T) {
	content, truncated := ParseLine(ShortLine, 1)
	expContent := "This is a short line"
	expTruncated := false
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}

func Test_ParseLine_LongLineColumnNumberNearStart(t *testing.T) {
	content, truncated := ParseLine(LongLine, 10)
	expContent := "This is a really long example of a line that will need to be shortened if this w"
	expTruncated := true
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}

func Test_ParseLine_LongLineColumnNumberNearEnd(t *testing.T) {
	content, truncated := ParseLine(LongLine, 110)
	expContent := "will need to be shortened if this whole content parsing thing is going to work??"
	expTruncated := true
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}

func Test_ParseLine_LongLineColumnNumberInMiddle(t *testing.T) {
	content, truncated := ParseLine(LongLine, 73)
	expContent := "a line that will need to be shortened if this whole content parsing thing is goi"
	expTruncated := true
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}

func Test_ParseLine_ColumnRefNotFoundShort(t *testing.T) {
	content, truncated := ParseLine(ShortLine, 30)
	expContent := "This is a short line"
	expTruncated := false
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}

func Test_ParseLine_ColumnRefNotFoundLong(t *testing.T) {
	content, truncated := ParseLine(LongLine, 9999)
	expContent := "This is a really long example of a line that will need to be shortened if this w"
	expTruncated := true
	assert.Equal(t, expContent, content)
	assert.Equal(t, expTruncated, truncated)
}
