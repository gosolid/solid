//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOrthography : objc.NSObject
*/

type NSOrthography struct {
  *objc.NSObject

}

func (r *NSOrthography) DominantScript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) LanguageMap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) InitWithDominantScript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

