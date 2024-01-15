//js:package native/ios/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Symbols.NSSymbolVariableColorEffect : Symbols.NSSymbolEffect
*/

type NSSymbolVariableColorEffect struct {
  *NSSymbolEffect

}

func (r *NSSymbolVariableColorEffect) EffectWithDimInactiveLayers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) EffectWithIterative() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) EffectWithCumulative() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) EffectWithReversing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) EffectWithNonReversing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolVariableColorEffect) EffectWithHideInactiveLayers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

