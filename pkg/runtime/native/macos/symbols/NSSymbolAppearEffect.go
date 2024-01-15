//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolAppearEffect : Symbols.NSSymbolEffect
*/

type NSSymbolAppearEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolAppearEffect) Effect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolAppearEffect) AppearUpEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolAppearEffect) AppearDownEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolAppearEffect) EffectWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolAppearEffect) EffectWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

