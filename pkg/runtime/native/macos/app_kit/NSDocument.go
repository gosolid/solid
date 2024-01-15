//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDocument : objc.NSObject
*/

type NSDocument struct {
  *objc.NSObject
  *NSEditorRegistration
  *foundation.NSFilePresenter
  *NSMenuItemValidation
  *NSUserInterfaceValidations
}

func (r *NSDocument) AutosavingFileType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileNameExtensionWasHiddenInLastRunSavePanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ShouldCloseWindowController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) DuplicateAndReturnError() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ChangeCountTokenForSaveOperation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) WindowNibName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) CanCloseDocumentWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveDocumentAs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) LockDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RunPageLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidGainVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsInViewingMode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ReadFromURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) UpdateChangeCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidLoseVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) UnlockDocumentWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) ScheduleAutosaving() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) MoveDocumentToUbiquityContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetFileURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetAutosavedContentsFileURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) HasUndoManager() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ReadFromData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrintDocumentWithSettings() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) MakeWindowControllers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetFileType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) KeepBackupFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) HasUnautosavedChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ReadFromFileWrapper() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) WriteToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) DuplicateDocumentWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrepareSharingServicePicker() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) ShowWindows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PreparePageLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileNameExtensionForType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosavesInPlace() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetHasUndoManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveDocumentTo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrepareSavePanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) LockDocumentWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileAttributesToWriteToURL() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) CanAsynchronouslyWriteToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) MoveDocumentWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) UnlockWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PerformActivityWithSynchronousWaiting() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) ShouldChangePrintInfo() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ValidateUserInterfaceItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileModificationDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosavingIsImplicitlyCancellable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ObservedPresentedItemUbiquityAttributes() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) MoveDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileWrapperOfType() (*foundation.NSFileWrapper, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) UnblockUserInteraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveToURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosaveDocumentWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RevertToContentsOfURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PerformSynchronousFileAccessUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WindowControllers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) InitForURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) RunModalSavePanelForSaveOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RelinquishPresentedItemToWriter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveDocumentWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) DuplicateDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RunModalPrintOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) UpdateChangeCountWithToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) DefaultDraftName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) BackupFileURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PreservesVersions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosavedContentsFileURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) InitWithType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) AllowsDocumentSharing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) WritableTypesForSaveOperation() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) RelinquishPresentedItemToReader() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidChangeUbiquityAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsLocked() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsDocumentEdited() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrintDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RunModalPageLayoutWithPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetDraft() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) UndoManager() (*foundation.NSUndoManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) LockWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) AccommodatePresentedItemDeletionWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrintInfo() (*NSPrintInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SavePresentedItemChangesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentError() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetUndoManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WindowForSheet() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) DataOfType() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) UsesUbiquitousStorage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) MoveToURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WriteSafelyToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidMoveToURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetPreviewRepresentableActivityItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WindowControllerWillLoadNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WindowControllerDidLoadNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsNativeType() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PDFPrintOperation() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) WritableTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveDocumentToPDF() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) UnlockDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsEntireFileLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) DisplayName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) RenameDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RevertDocumentToSaved() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetDisplayName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) CanConcurrentlyReadDocumentsOfType() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PreviewRepresentableActivityItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ShareDocumentWithSharingService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) AddWindowController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosavesDrafts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PerformAsynchronousFileAccessUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) ContinueAsynchronousWorkOnMainThreadUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) PrintOperationWithSettings() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsDraft() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ContinueActivityUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) BrowseDocumentVersions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) StopBrowsingVersionsWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WillPresentError() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) ShouldRunSavePanelWithAccessoryView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) FileTypeFromLastRunSavePanel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) IsBrowsingVersions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) SaveDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) AutosaveWithImplicitCancellability() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) RemoveWindowController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) SetFileModificationDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) ReadableTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocument) CheckAutosavingSafetyAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocument) PresentedItemDidResolveConflictVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocument) WillNotPresentError() error {
  return fmt.Errorf("unimplemented")
}

