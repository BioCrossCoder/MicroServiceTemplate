package models

import (
	"encoding"
	"encoding/json"

	"github.com/biocrosscoder/flex/typed/collections/set"
)

type cacheModel interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

var (
	_ = cacheModel(&AppCache{})
	_ = cacheModel(&UserAccessorsCache{})
)

type AppCache struct {
	Name        string        `json:"name"`
	Entry       string        `json:"entry"`
	Order       int           `json:"order"`
	Enabled     bool          `json:"enabled"`
	LimitAccess bool          `json:"limit_access"`
	Users       []*UserObject `json:"users"`
}

func (c *AppCache) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *AppCache) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

type UserAccessorsCache struct {
	UserID  string          `json:"user_id"`
	Parents set.Set[string] `json:"parents"`
}

func (c *UserAccessorsCache) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *UserAccessorsCache) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
