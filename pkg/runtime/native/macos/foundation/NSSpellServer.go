//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSSpellServer : objc.NSObject
*/

type NSSpellServer struct {
  *objc.NSObject

}

func (r *NSSpellServer) IsWordInUserDictionaries() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSpellServer) Run() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellServer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpellServer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpellServer) RegisterLanguage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

