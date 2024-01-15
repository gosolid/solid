//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSUserActivity : objc.NSObject
*/

type NSUserActivity struct {
  *objc.NSObject

}

func (r *NSUserActivity) ActivityType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ExpirationDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetTargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForHandoff() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ResignCurrent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetReferrerURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Keywords() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetSupportsContinuationStreams() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForHandoff() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetRequiredUserInfoKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetWebpageURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetExpirationDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SupportsContinuationStreams() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForPublicIndexing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) PersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetPersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) AddUserInfoEntriesFromDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) BecomeCurrent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) ReferrerURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetKeywords() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForSearch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) DeleteSavedUserActivitiesWithPersistentIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetNeedsSave() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) TargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForPublicIndexing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) InitWithActivityType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) IsEligibleForPrediction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) GetContinuationStreamsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForPrediction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) DeleteAllSavedUserActivitiesWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) RequiredUserInfoKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) NeedsSave() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) WebpageURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserActivity) SetEligibleForSearch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

