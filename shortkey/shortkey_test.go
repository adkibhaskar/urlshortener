package shortkey

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortKey(t *testing.T){

	t.Run("Matching Short key",func(t *testing.T) {

		shortkey:=GenerateShortKey()
		
		typeOfGeneratedShortkey:=fmt.Sprintf("%T",shortkey)
		typeOfShortKey:=fmt.Sprintf("%T","")

		assert.Equal(t,typeOfShortKey,typeOfGeneratedShortkey,"Generted Short Key")
	})
}