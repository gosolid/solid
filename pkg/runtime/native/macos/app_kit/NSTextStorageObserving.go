//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextStorageObserving
*/

type NSTextStorageObserving struct {

}

func (r *NSTextStorageObserving) TextStorage() (*NSTextStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorageObserving) SetTextStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorageObserving) ProcessEditingForTextStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorageObserving) PerformEditingTransactionForTextStorage() error {
  return fmt.Errorf("unimplemented")
}

