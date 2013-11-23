// parsetumbr project parsetumbr_test.go
package parsetumblr

import (
	"net/http"
	"testing"
)

func TestNewFeed(t *testing.T) {
	url := "http://wilnregina.tumblr.com"
	newFeed := NewFeed(url)
	if newFeed.Url != url {
		t.Errorf("newFeed not getting initialized correctly.")
	}
}

func TestGetFeed(t *testing.T) {
	url := "http://wilnregina.tumblr.com"
	newFeed := NewFeed(url)

	newFeed.Limit = 50

	var client http.Client

	err := newFeed.GetFeed(&client)
	if err != nil {
		t.Errorf("%v", err)
	}

	newFeed = NewFeed(url)
	newFeed.StartIndex = 2
	newFeed.Limit = 3

	err = newFeed.GetFeed(&client)
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(newFeed.Entries.Entries) != 3 {
		t.Errorf("The Limit is not being set.")
	}

	if newFeed.StartIndex != 2 {
		t.Errorf("The StartIndex is not being set.")
	}

	if newFeed.Entries.Entries[0].Published().IsZero() {
		t.Errorf("The Published time is not being set on the Post.")
	}
}
