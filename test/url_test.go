package gotest

import (
  "testing"
  "github.com/katsew/minivalid/url"
)

var stub = "http://exmaple.com/a/b/c.html"
func TestUrl(t *testing.T) {
  if !url.Validate(stub) {
    t.Error("failed")
  } else {
    t.Log("ok")
  }
}
