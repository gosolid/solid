//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDictationPhrase : objc.NSObject
*/

type UIDictationPhrase struct {
  *objc.NSObject

}

func (r *UIDictationPhrase) Text() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDictationPhrase) AlternativeInterpretations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

