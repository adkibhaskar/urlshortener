package handleform

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleForm(t *testing.T) {

	req,err:=http.NewRequest("GET","/",nil)

	if err != nil{

		t.Fatalf("Could not created request : %v",err)
	}

	rec:=httptest.NewRecorder()

	HandleForm(rec,req)

	res:=rec.Result()

	if res.StatusCode != http.StatusOK{
		
		t.Errorf("Expected status OK;got %v",res.StatusCode)

	}


	
}