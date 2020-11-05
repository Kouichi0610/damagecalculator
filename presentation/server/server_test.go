package server

import (
	//"net/http"
	_ "net/http/httptest"
	_ "strings"
	"testing"

	"encoding/json"
)

/*
type speedResult struct {
	Info  string `json:"info"`
	Speed uint   `json:"speed"`
}
*/

// interface は無理？
type testSampler interface {
	Name() string
	Mail() string
}

type testImpl struct {
	name string `json:"name"`
	mail string `json:"mail"`
}

func (i *testImpl) Name() string {
	return i.name
}
func (i *testImpl) Mail() string {
	return i.mail
}

func newSampler(name, mail string) testSampler {
	return &testImpl{
		name: name,
		mail: mail,
	}
}

func Test_Json(t *testing.T) {
	s := &speedResult{
		Info:  "Speed Sample.",
		Speed: 158,
	}
	str, _ := json.Marshal(s)

	t.Errorf("%s", str)

	smpl := newSampler("Kou", "test@mail.com")
	str, _ = json.Marshal(smpl)
	t.Errorf("%s", str)

}

/*
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
*/
