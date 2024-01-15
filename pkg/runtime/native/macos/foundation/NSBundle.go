//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSBundle : objc.NSObject
*/

type NSBundle struct {
  *objc.NSObject

}

func (r *NSBundle) BundleIdentifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) DevelopmentLocalization() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InitWithPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Unload() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBundle) MainBundle() (*NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathsForResourcesOfType() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedFrameworksURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ResourcePath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Load() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLForAuxiliaryExecutable() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLForResource() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreflightAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBundle) IsLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedFrameworksPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedSupportURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BuiltInPlugInsPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InfoDictionary() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreferredLocalizations() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) URLsForResourcesWithExtension() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ClassNamed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AllFrameworks() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutableArchitectures() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AllBundles() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ResourceURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutableURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrivateFrameworksPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleWithIdentifier() (*NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LoadAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathForResource() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedStringForKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrincipalClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) SharedSupportPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PathForAuxiliaryExecutable() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedAttributedStringForKey() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ObjectForInfoDictionaryKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BuiltInPlugInsURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) LocalizedInfoDictionary() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PrivateFrameworksURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) AppStoreReceiptURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundlePath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) ExecutablePath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) BundleForClass() (*NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) PreferredLocalizationsFromArray() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundle) Localizations() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

