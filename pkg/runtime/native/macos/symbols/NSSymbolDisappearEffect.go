//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolDisappearEffect : Symbols.NSSymbolEffect
*/

type NSSymbolDisappearEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolDisappearEffect) EffectWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolDisappearEffect) Effect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolDisappearEffect) DisappearUpEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolDisappearEffect) DisappearDownEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolDisappearEffect) EffectWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

