//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextAlternatives : objc.NSObject
*/

type NSTextAlternatives struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSTextAlternatives) NoteSelectedAlternativeString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAlternatives) PrimaryString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAlternatives) AlternativeStrings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAlternatives) InitWithPrimaryString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

