//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSProgress : objc.NSObject
*/

type NSProgress struct {
  *objc.NSObject

}

func (r *NSProgress) BecomeCurrentWithPendingUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Pause() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetTotalUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) EstimatedTimeRemaining() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) CompletedUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) InitWithParent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) LocalizedAdditionalDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCancellationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Kind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetUserInfoObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) CancellationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ProgressWithTotalUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetEstimatedTimeRemaining() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Throughput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) AddSubscriberForFileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) PausingHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsCancellable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileTotalCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) RemoveSubscriber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetResumingHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCancellable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsPaused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ResumingHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetThroughput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Publish() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetCompletedUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetPausable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileOperationKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileTotalCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) PerformAsCurrentWithPendingUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Resume() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsPausable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsOld() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) CurrentProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) DiscreteProgressWithTotalUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) ResignCurrent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetPausingHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FileCompletedCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) Unpublish() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) TotalUnitCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) FractionCompleted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileOperationKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetFileCompletedCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) AddChild() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) IsIndeterminate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) LocalizedDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetLocalizedDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) SetLocalizedAdditionalDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProgress) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

