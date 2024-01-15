//js:package native/ios/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Symbols.NSSymbolContentTransition : objc.NSObject
*/

type NSSymbolContentTransition struct {
  *objc.NSObject

}

func (r *NSSymbolContentTransition) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolContentTransition) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

