package api

import "encoding/gob"

func init() {
	gob.Register([]*Record{})
	gob.Register(map[string]interface{}{})
}