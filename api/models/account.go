package models

import (
	"errors"
	"time"

	"appengine/datastore"
)

type (
	AccountItem struct {
		Owner       string           `json:"owner" datastore:"owner"`
		DateTime    time.Time        `json:"dateTime" datastore:"datetime"`
		Amount      int64            `json:"amount" datastore:"amount,noindex"`
		Tags        []Tag            `json:"tags,emitempty" datastore:"tags"`
		Description string           `json:"description,emitempty" datastore:"description,noindex"`
		SubItems    []*AccountInItem `json:"subItems,emitempty" datastore:"-"`
		SubItemKeys []string         `json:"-" datastore:"subitemkeys,noindex"`
		Links       []*AccountInItem `json:"links,emitempty" datastore:"-"`
		LinkKeys    []string         `json:"-" datastore:"linkkeys,noindex"`
		IsSubItem   bool             `json:"isSubItem" datastore:"issubitem,noindex"`
	}

	AccountInItem struct {
		AccountItem
	}

	AccountOutItem struct {
		AccountItem
	}
)

func (ai *AccountInItem) Save(c chan<- datastore.Property) error {
	defer close(c)

	if ai.IsSubItem {
		return datastore.SaveStruct(ai, c)
	}

	subItemLen := len(ai.SubItems)

	if subItemLen == 0 {
		return datastore.SaveStruct(ai, c)
	}

	checkSum(&ai.AccountItem)

	return datastore.SaveStruct(ai, c)
}

func (ai *AccountOutItem) Save(c chan<- datastore.Property) error {
	defer close(c)

	if ai.IsSubItem {
		return datastore.SaveStruct(ai, c)
	}

	subItemLen := len(ai.SubItems)

	if subItemLen == 0 {
		return datastore.SaveStruct(ai, c)
	}

	checkSum(&ai.AccountItem)

	return datastore.SaveStruct(ai, c)
}

func checkSum(ai *AccountItem) error {
	var (
		total   int64 = 0
		itemLen int   = len(ai.SubItems)
	)

	for i := 0; i < itemLen; i++ {
		item := ai.SubItems[i]
		total += item.Amount
	}

	if total != ai.Amount {
		return errors.New("Item amount is not equal to total amount of sub item.")
	}

	return nil
}
