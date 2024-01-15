//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSTermOfAddress : objc.NSObject
*/

type NSTermOfAddress struct {
  *objc.NSObject

}

func (r *NSTermOfAddress) Feminine() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Masculine() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) LocalizedForLanguageIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) LanguageIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Pronouns() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Neutral() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

