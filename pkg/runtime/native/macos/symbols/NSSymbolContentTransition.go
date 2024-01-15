//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Symbols.NSSymbolContentTransition : objc.NSObject
*/

type NSSymbolContentTransition struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSSymbolContentTransition) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolContentTransition) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

