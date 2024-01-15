//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSOrthography : objc.NSObject
*/

type NSOrthography struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSOrthography) InitWithDominantScript() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) DominantScript() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrthography) LanguageMap() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

