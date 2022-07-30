package test

import (
	"Raspi-TV-Control/pkg/handler/apihandler"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer(t *testing.T) {
	t.Run("Test Home Route", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		apihandler.HomeHandler(response, request)

		got := map[string]string{}
		json.Unmarshal(response.Body.Bytes(), &got)

		if _, ok := got["OS"]; !ok {
			t.Error("Not all system data has been returned")
			return
		}

	})

}
