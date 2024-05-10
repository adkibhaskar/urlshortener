package summary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCheckIfUrlssArraySorted(t *testing.T){

	t.Run("Urls Array is Sorted",func(t *testing.T) {

		urlsArray:=[]int{5,4,3,2,1}

		_,result:=CheckIfUrlssArraySorted(urlsArray)

		assert.Equal(t,"Urls Array is  Sorted",result,"Urls Array is Sorted")
	})

	t.Run("Urls Array is Not Sorted",func(t *testing.T) {

		urlsArray:=[]int{1,2,3,4,5}

		_,result:=CheckIfUrlssArraySorted(urlsArray)

		assert.Equal(t,"Urls Array is Not Sorted",result,"Urls Array is Not Sorted")
	
	})
}