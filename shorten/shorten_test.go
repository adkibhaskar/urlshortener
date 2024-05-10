package shorten

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleShorten(t *testing.T){

	t.Run("Request Successful",func(t *testing.T) {

		req,err:=http.NewRequest("GET","/shorten",nil)

		if err != nil{
			
			t.Fatalf("Could not created request : %v",err)
		}

		rec:=httptest.NewRecorder()
		
		HandleShorten(rec,req)
		
		res:=rec.Result()

	    assert.Equal(t,http.StatusMethodNotAllowed,res.StatusCode,"Error Occurred while making request")
	})

	t.Run("Url Paramter is Missing",func(t *testing.T) {

		req,err:=http.NewRequest("POST","",nil)

		if err != nil{

			t.Fatalf("Error Occured %v ",err)
		}

		rec:=httptest.NewRecorder()

		HandleShorten(rec,req)

		assert.Equal(t,http.StatusBadRequest,rec.Code,"Url Paramter is Missing")
	})
}

