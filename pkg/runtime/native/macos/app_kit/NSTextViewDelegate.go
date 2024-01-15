//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTextViewDelegate
*/

type NSTextViewDelegate struct {

}

func (r *NSTextViewDelegate) UndoManagerForTextView() (*foundation.NSUndoManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewDelegate) TextView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextViewDelegate) TextViewDidChangeSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewDelegate) TextViewDidChangeTypingAttributes() error {
  return fmt.Errorf("unimplemented")
}

