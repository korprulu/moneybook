package models

import (
	"errors"
	"time"

	"appengine/datastore"
)

type AccountItem struct {
	Owner       string        `json:"owner" datastore:"owner"`
	DateTime    time.Time     `json:"dateTime" datastore:"datetime"`
	Amount      int64         `json:"amount" datastore:"amount"`
	Tags        []Tag         `json:"tags,emitempty" datastore:"tags"`
	Description string        `json:"description" datastore:"description,noindex"`
	SubItems    []AccountItem `json:"subItem,emitempty"`
}

type AccountSubItem struct {
	Description string `json:"description,emitempty" datastore:"description,noindex"`
	Amount      int64  `json:"amount" datastore:"amount"`
}

func (ai *AccountItem) Save(c chan<- datastore.Property) error {
	defer close(c)

	subItemLen := len(ai.SubItems)

	if subItemLen == 0 {
		return datastore.SaveStruct(ai, c)
	}

	// Check total amount of sub item is equal to item's amount.
	var total int64 = 0

	for i := 0; i < subItemLen; i++ {
		item := ai.SubItems[i]
		total += item.Amount
	}

	if total != ai.Amount {
		return errors.New("Item amount is not equal to total amount of sub item.")
	}

	return datastore.SaveStruct(ai, c)
}
