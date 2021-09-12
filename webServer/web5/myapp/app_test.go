package myapp

import (
	assert2 "github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert2.New(t)

	ts := httptest.NewServer(NewHandler()) //목업테스트 서버. 실제 서버는 아니지만
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func testUseres(t *testing.T) {
	assert := assert2.New(t)

	ts := httptest.NewServer(NewHandler()) //목업테스트 서버. 실제 서버는 아니지만
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo")
	//assert.Equal("Hello World", string(data))

}
