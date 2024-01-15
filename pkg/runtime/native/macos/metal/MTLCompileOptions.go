//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLCompileOptions : objc.NSObject
*/

type MTLCompileOptions struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLCompileOptions) InstallName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetInstallName() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetPreserveInvariance() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) LanguageVersion() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) Libraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) AllowReferencingUndefinedSymbols() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetMaxTotalThreadsPerThreadgroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetFastMathEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetPreprocessorMacros() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) CompileSymbolVisibility() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetAllowReferencingUndefinedSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) MaxTotalThreadsPerThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) PreprocessorMacros() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLanguageVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) LibraryType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLibraryType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) PreserveInvariance() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) OptimizationLevel() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetOptimizationLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) FastMathEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCompileOptions) SetCompileSymbolVisibility() error {
  return fmt.Errorf("unimplemented")
}

