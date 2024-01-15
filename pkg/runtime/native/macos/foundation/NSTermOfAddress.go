//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSTermOfAddress : objc.NSObject
*/

type NSTermOfAddress struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSTermOfAddress) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) LanguageIdentifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Pronouns() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Neutral() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Feminine() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) Masculine() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTermOfAddress) LocalizedForLanguageIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

