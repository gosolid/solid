//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLCompileOptions : objc.NSObject
*/

type MTLCompileOptions struct {
  *objc.NSObject

}

func (r *MTLCompileOptions) OptimizationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetOptimizationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) CompileSymbolVisibility() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetCompileSymbolVisibility() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) Libraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) PreserveInvariance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetAllowReferencingUndefinedSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetMaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLibraryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetInstallName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetFastMathEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) LanguageVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) LibraryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) PreprocessorMacros() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetPreprocessorMacros() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) InstallName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetPreserveInvariance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) AllowReferencingUndefinedSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) MaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) FastMathEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLanguageVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

