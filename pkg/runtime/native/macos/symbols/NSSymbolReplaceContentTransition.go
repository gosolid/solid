//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolReplaceContentTransition : Symbols.NSSymbolContentTransition
*/

type NSSymbolReplaceContentTransition struct {
  *NSSymbolContentTransition

}

func (r *NSSymbolReplaceContentTransition) TransitionWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolReplaceContentTransition) TransitionWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolReplaceContentTransition) Transition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolReplaceContentTransition) ReplaceDownUpTransition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolReplaceContentTransition) ReplaceUpUpTransition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolReplaceContentTransition) ReplaceOffUpTransition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

