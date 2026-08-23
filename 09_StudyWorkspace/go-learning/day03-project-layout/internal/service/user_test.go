package service

import "testing"

func TestGetUserName(t *testing.T) {
	name := GetUserName()

	if name != "UESTCHh" {
		t.Errorf("expected UESTCHh but got %s", name)
	}
}
