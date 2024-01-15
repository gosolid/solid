//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTokenField : AppKit.NSTextField
*/

type NSTokenField struct {
  *NSTextField

}

func (r *NSTokenField) TokenizingCharacterSet() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) DefaultTokenizingCharacterSet() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenField) SetTokenStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenField) CompletionDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) DefaultCompletionDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) SetTokenizingCharacterSet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenField) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) TokenStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenField) SetCompletionDelay() error {
  return fmt.Errorf("unimplemented")
}

