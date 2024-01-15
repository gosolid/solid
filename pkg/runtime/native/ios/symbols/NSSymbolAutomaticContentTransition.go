//js:package native/ios/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolAutomaticContentTransition : Symbols.NSSymbolContentTransition
*/

type NSSymbolAutomaticContentTransition struct {
  *NSSymbolContentTransition

}

func (r *NSSymbolAutomaticContentTransition) Transition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

