//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSpeechRecognizer : objc.NSObject
*/

type NSSpeechRecognizer struct {
  *objc.NSObject

}

func (r *NSSpeechRecognizer) StartListening() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) ListensInForegroundOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) SetBlocksOtherRecognizers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) Commands() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) SetCommands() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) DisplayedCommandsTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) StopListening() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) BlocksOtherRecognizers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) SetDisplayedCommandsTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpeechRecognizer) SetListensInForegroundOnly() error {
  return fmt.Errorf("unimplemented")
}

