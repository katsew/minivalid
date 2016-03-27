package gotest

import (
  "testing"
  "github.com/katsew/minivalid/mail"
)

var stub = "sample@example.com"
func TestMail(t *testing.T) {
  if !mail.Validate(stub) {
    t.Error("failed")
  } else {
    t.Log("ok")
  }
}
