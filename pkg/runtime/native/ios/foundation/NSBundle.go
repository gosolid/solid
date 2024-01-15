//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSBundle : objc.NSObject
*/

type NSBundle struct {
  *objc.NSObject

}

func (r *NSBundle) ResourcePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrivateFrameworksPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreferredLocalizations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundlePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedFrameworksPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LoadAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathForAuxiliaryExecutable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreferredLocalizationsFromArray() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutableURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AppStoreReceiptURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedInfoDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Localizations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedSupportPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrincipalClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathsForResourcesOfType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AllBundles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrivateFrameworksURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BuiltInPlugInsURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutablePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InitWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreflightAndReturnError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedSupportURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLsForResourcesWithExtension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ObjectForInfoDictionaryKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutableArchitectures() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InfoDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) DevelopmentLocalization() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Unload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLForResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedStringForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ClassNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) IsLoaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedFrameworksURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLForAuxiliaryExecutable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedAttributedStringForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ResourceURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BuiltInPlugInsPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleForClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Load() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathForResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) MainBundle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AllFrameworks() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

