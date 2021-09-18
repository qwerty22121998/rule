package rule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequire(t *testing.T) {
	var emptyArr []int
	assert.Error(t, Require("nil", nil).Validate())
	assert.Error(t, Require("number", 0).Validate())
	assert.Error(t, Require("slice", emptyArr).Validate())
	assert.Error(t, Require("string", "").Validate())

	assert.Nil(t, Require("number", 1).Validate())
	assert.Nil(t, Require("slice", []int{1, 2, 3}).Validate())
	assert.Nil(t, Require("string", "hello").Validate())

}
