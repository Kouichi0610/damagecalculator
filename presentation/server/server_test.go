package server

import (
	//"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Server(t *testing.T) {
	sv := NewServer(nil, nil, nil, nil, nil)
	router, err := sv.Serve()

	if err != nil {
		t.Error(err)
	}
	// TODO:テスト方法

	bodyReader := strings.NewReader(`{"x": "15" "y": "30}`)
	req := httptest.NewRequest("POST", "/post_sample", bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	t.Errorf("%d\n", rec.Code)

}
