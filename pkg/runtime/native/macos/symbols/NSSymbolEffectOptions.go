//js:package native/macos/symbols
package symbols

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Symbols.NSSymbolEffectOptions : objc.NSObject
*/

type NSSymbolEffectOptions struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSSymbolEffectOptions) OptionsWithSpeed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) Options() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) OptionsWithRepeating() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) OptionsWithNonRepeating() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSymbolEffectOptions) OptionsWithRepeatCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}
