package middleware

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestUntypedFormPost(t *testing.T) {
	params := parametersForFormUpload()
	binder := newUntypedRequestBinder(params, nil, strfmt.Default)

	urlStr := "http://localhost:8002/hello"
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(`name=the-name&age=32`))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := make(map[string]interface{})
	assert.NoError(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))
	assert.Equal(t, "the-name", data["name"])
	assert.EqualValues(t, 32, data["age"])

	req, _ = http.NewRequest("POST", urlStr, bytes.NewBufferString(`name=%3&age=32`))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	data = make(map[string]interface{})
	assert.Error(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))
}

func TestUntypedFileUpload(t *testing.T) {
	binder := paramsForFileUpload()

	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "plain-jane.txt")
	assert.NoError(t, err)

	part.Write([]byte("the file contents"))
	writer.WriteField("name", "the-name")
	assert.NoError(t, writer.Close())

	urlStr := "http://localhost:8002/hello"
	req, _ := http.NewRequest("POST", urlStr, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	data := make(map[string]interface{})
	assert.NoError(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))
	assert.Equal(t, "the-name", data["name"])
	assert.NotNil(t, data["file"])
	assert.IsType(t, httpkit.File{}, data["file"])
	file := data["file"].(httpkit.File)
	assert.NotNil(t, file.Header)
	assert.Equal(t, "plain-jane.txt", file.Header.Filename)

	bb, err := ioutil.ReadAll(file.Data)
	assert.NoError(t, err)
	assert.Equal(t, []byte("the file contents"), bb)

	req, _ = http.NewRequest("POST", urlStr, body)
	req.Header.Set("Content-Type", "application/json")
	data = make(map[string]interface{})
	assert.Error(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))

	req, _ = http.NewRequest("POST", urlStr, body)
	req.Header.Set("Content-Type", "application(")
	data = make(map[string]interface{})
	assert.Error(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))

	body = bytes.NewBuffer(nil)
	writer = multipart.NewWriter(body)
	part, err = writer.CreateFormFile("bad-name", "plain-jane.txt")
	assert.NoError(t, err)

	part.Write([]byte("the file contents"))
	writer.WriteField("name", "the-name")
	assert.NoError(t, writer.Close())
	req, _ = http.NewRequest("POST", urlStr, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	data = make(map[string]interface{})
	assert.Error(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))

	req, _ = http.NewRequest("POST", urlStr, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.MultipartReader()

	data = make(map[string]interface{})
	assert.Error(t, binder.Bind(req, nil, httpkit.JSONConsumer(), &data))
}

func TestUntypedBindingTypesForValid(t *testing.T) {

	op2 := parametersForAllTypes("")
	binder := newUntypedRequestBinder(op2, nil, strfmt.Default)

	confirmed := true
	name := "thomas"
	friend := map[string]interface{}{"name": "toby", "age": float64(32)}
	id, age, score, factor := int64(7575), int32(348), float32(5.309), float64(37.403)
	requestID := 19394858
	tags := []string{"one", "two", "three"}
	dt1 := time.Date(2014, 8, 9, 0, 0, 0, 0, time.UTC)
	planned := strfmt.Date{Time: dt1}
	dt2 := time.Date(2014, 10, 12, 8, 5, 5, 0, time.UTC)
	delivered := strfmt.DateTime{Time: dt2}
	picture := base64.URLEncoding.EncodeToString([]byte("hello"))
	uri, _ := url.Parse("http://localhost:8002/hello/7575")
	qs := uri.Query()
	qs.Add("name", name)
	qs.Add("confirmed", "true")
	qs.Add("age", "348")
	qs.Add("score", "5.309")
	qs.Add("factor", "37.403")
	qs.Add("tags", strings.Join(tags, ","))
	qs.Add("planned", planned.String())
	qs.Add("delivered", delivered.String())
	qs.Add("picture", picture)

	req, _ := http.NewRequest("POST", uri.String()+"?"+qs.Encode(), bytes.NewBuffer([]byte(`{"name":"toby","age":32}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-Id", "19394858")

	data := make(map[string]interface{})
	err := binder.Bind(req, RouteParams([]RouteParam{{"id", "7575"}}), httpkit.JSONConsumer(), &data)
	assert.NoError(t, err)
	assert.Equal(t, id, data["id"])
	assert.Equal(t, name, data["name"])
	assert.Equal(t, friend, data["friend"])
	assert.EqualValues(t, requestID, data["X-Request-Id"])
	assert.Equal(t, tags, data["tags"])
	assert.Equal(t, planned, data["planned"])
	assert.Equal(t, delivered, data["delivered"])
	assert.Equal(t, confirmed, data["confirmed"])
	assert.Equal(t, age, data["age"])
	assert.Equal(t, factor, data["factor"])
	assert.Equal(t, score, data["score"])
	pb, _ := base64.URLEncoding.DecodeString(picture)
	assert.EqualValues(t, pb, data["picture"].(strfmt.Base64))

}
