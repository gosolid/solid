//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextCheckingController : objc.NSObject
*/

type NSTextCheckingController struct {
  *objc.NSObject

}

func (r *NSTextCheckingController) MenuAtIndex() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) SetSpellCheckerDocumentTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) InitWithClient() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) InsertedTextInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) DidChangeSelectedRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) OrderFrontSubstitutionsPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) ShowGuessPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) SpellCheckerDocumentTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) DidChangeTextInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) ChangeSpelling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) IgnoreSpelling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) ValidAnnotations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) Client() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) ConsiderTextCheckingForRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) CheckTextInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) CheckTextInDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) UpdateCandidates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) CheckTextInSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingController) CheckSpelling() error {
  return fmt.Errorf("unimplemented")
}

