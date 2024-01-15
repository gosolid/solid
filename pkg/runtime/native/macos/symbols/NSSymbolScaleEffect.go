//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolScaleEffect : Symbols.NSSymbolEffect
*/

type NSSymbolScaleEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolScaleEffect) Effect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolScaleEffect) ScaleUpEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolScaleEffect) ScaleDownEffect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolScaleEffect) EffectWithByLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolScaleEffect) EffectWithWholeSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

