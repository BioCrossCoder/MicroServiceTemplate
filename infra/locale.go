package infra

import (
	"encoding/json"
	"sync"
)

type I18nResource interface {
	Get(key, language string, args ...any) (text string, err error)
}

var (
	i18n     I18nResource
	i18nOnce sync.Once
)

type i18nResource struct {
	client HttpClient
}

func NewI18nResource() I18nResource {
	i18nOnce.Do(func() {
		i18n = &i18nResource{
			client: NewHttpClient(),
		}
	})
	return i18n
}

func (r *i18nResource) Get(key, language string, args ...any) (text string, err error) {
	reqBody := &i18nReq{
		Key:      key,
		Language: language,
		Args:     args,
	}
	resBody, err := r.client.Post("http://localhost:8888/i18n", reqBody)
	if err != nil {
		return
	}
	var res i18nRes
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return
	}
	text = res.Text
	return
}

type i18nReq struct {
	Key      string `json:"key"`
	Language string `json:"language"`
	Args     []any  `json:"args"`
}

type i18nRes struct {
	Text string `json:"text"`
}
