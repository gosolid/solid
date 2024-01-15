//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolBounceEffect : Symbols.NSSymbolEffect
*/

type NSSymbolBounceEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolBounceEffect) Effect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolBounceEffect) BounceUpEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolBounceEffect) BounceDownEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolBounceEffect) EffectWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolBounceEffect) EffectWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

