package redirectapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoToOriginalPage(t *testing.T) {
	t.Run("key Not Found", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/short/qRpf", nil)

		if err != nil {

			t.Fatalf("Could not created request : %v", err)
		}

		rec := httptest.NewRecorder()
		GoToOriginalPage(rec, req)

		// res:=rec.Result()

		// assert.Equal(t,http.StatusNotFound,res.StatusCode,"key Not Found")

	})

}
