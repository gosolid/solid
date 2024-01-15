//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextElementProvider
*/

type NSTextElementProvider struct {

}

func (r *NSTextElementProvider) SynchronizeToBackingStore() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) LocationFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) OffsetFromLocation() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) AdjustedRangeFromRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) DocumentRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) EnumerateTextElementsFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElementProvider) ReplaceContentsInRange() error {
  return fmt.Errorf("unimplemented")
}

