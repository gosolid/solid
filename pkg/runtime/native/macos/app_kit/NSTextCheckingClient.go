//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTextCheckingClient
*/

type NSTextCheckingClient struct {

}

func (r *NSTextCheckingClient) AddAnnotations() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) RemoveAnnotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) ReplaceCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) SelectAndShowRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) ViewForRange() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) CandidateListTouchBarItem() (*NSCandidateListTouchBarItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) AnnotatedSubstringForProposedRange() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingClient) SetAnnotations() error {
  return fmt.Errorf("unimplemented")
}

