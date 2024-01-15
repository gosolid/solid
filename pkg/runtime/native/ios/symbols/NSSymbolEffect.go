//js:package native/ios/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Symbols.NSSymbolEffect : objc.NSObject
*/

type NSSymbolEffect struct {
  *objc.NSObject

}

func (r *NSSymbolEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

