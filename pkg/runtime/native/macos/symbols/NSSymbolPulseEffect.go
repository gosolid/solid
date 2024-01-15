//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolPulseEffect : Symbols.NSSymbolEffect
*/

type NSSymbolPulseEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolPulseEffect) Effect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolPulseEffect) EffectWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolPulseEffect) EffectWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

