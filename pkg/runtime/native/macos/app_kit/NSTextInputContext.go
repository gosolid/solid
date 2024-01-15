//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextInputContext : objc.NSObject
*/

type NSTextInputContext struct {
  *objc.NSObject

}

func (r *NSTextInputContext) CurrentInputContext() (*NSTextInputContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) AcceptsGlyphInfo() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) AllowedInputSourceLocales() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) SetSelectedKeyboardInputSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) TextInputClientDidEndScrollingOrZooming() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) TextInputClientWillStartScrollingOrZooming() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) SetAcceptsGlyphInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) SelectedKeyboardInputSource() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) InvalidateCharacterCoordinates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) HandleEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) DiscardMarkedText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) LocalizedNameForInputSource() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) Client() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) SetAllowedInputSourceLocales() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) Activate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) Deactivate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) KeyboardInputSources() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputContext) InitWithClient() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

