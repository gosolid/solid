//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserActivity : objc.NSObject
*/

type NSUserActivity struct {
  *objc.NSObject

}

func (r *NSUserActivity) IsEligibleForSearch() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) InitWithActivityType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) GetContinuationStreamsWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetRequiredUserInfoKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) WebpageURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetWebpageURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) AddUserInfoEntriesFromDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ResignCurrent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) TargetContentIdentifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetTargetContentIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForPublicIndexing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ExpirationDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) DeleteSavedUserActivitiesWithPersistentIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Title() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetKeywords() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetPersistentIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForPublicIndexing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) PersistentIdentifier() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ActivityType() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) NeedsSave() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetReferrerURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Keywords() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetSupportsContinuationStreams() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForPrediction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) BecomeCurrent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) DeleteAllSavedUserActivitiesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ReferrerURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetExpirationDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForHandoff() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) RequiredUserInfoKeys() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForSearch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetNeedsSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SupportsContinuationStreams() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForHandoff() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForPrediction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

