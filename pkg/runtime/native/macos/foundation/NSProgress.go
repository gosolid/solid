//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSProgress : objc.NSObject
*/

type NSProgress struct {
  *objc.NSObject

}

func (r *NSProgress) LocalizedAdditionalDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetEstimatedTimeRemaining() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCancellable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetPausingHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) PerformAsCurrentWithPendingUnitCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetLocalizedAdditionalDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetThroughput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCancellationHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsFinished() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) DiscreteProgressWithTotalUnitCount() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetUserInfoObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) Throughput() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileOperationKind() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsOld() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) RemoveSubscriber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetPausable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) Kind() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) CancellationHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) PausingHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileCompletedCount() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) BecomeCurrentWithPendingUnitCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) CompletedUnitCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetLocalizedDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileCompletedCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) AddSubscriberForFileURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) EstimatedTimeRemaining() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileOperationKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) FractionCompleted() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ProgressWithTotalUnitCount() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetTotalUnitCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsPaused() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Publish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) Unpublish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) TotalUnitCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) CurrentProgress() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) AddChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) Pause() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsPausable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsCancellable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsIndeterminate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Resume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) LocalizedDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsCancelled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ResumingHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetResumingHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileTotalCount() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) InitWithParent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ResignCurrent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCompletedUnitCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileTotalCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgress) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

