//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSFilePresenter
*/

type NSFilePresenter struct {

}

func (r *NSFilePresenter) PresentedItemDidGainVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedSubitemDidAppearAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) SavePresentedItemChangesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemDidMoveToURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemDidLoseVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemOperationQueue() (*NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) ObservedPresentedItemUbiquityAttributes() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) RelinquishPresentedItemToReader() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) RelinquishPresentedItemToWriter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) AccommodatePresentedItemDeletionWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemDidChangeUbiquityAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedItemDidResolveConflictVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) AccommodatePresentedSubitemDeletionAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedSubitemAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PresentedSubitemDidChangeAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePresenter) PrimaryPresentedItemURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

