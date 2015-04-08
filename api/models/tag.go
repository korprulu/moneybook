package models

type Tag struct {
	Owner string `json:"owner" datastore:"owner"`
	Name  string `json:"name" datastore:"name"`
}
