//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Symbols.NSSymbolEffect : objc.NSObject
*/

type NSSymbolEffect struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSSymbolEffect) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffect) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

