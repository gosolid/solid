//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSDocumentController : objc.NSObject
*/

type NSDocumentController struct {
  *objc.NSObject
  *foundation.NSCoding
  *NSMenuItemValidation
  *NSUserInterfaceValidations
}

func (r *NSDocumentController) MakeUntitledDocumentOfType() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) CloseAllDocumentsWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) NoteNewRecentDocumentURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DisplayNameForType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) SetAutosavingDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) HasEditedDocuments() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) OpenUntitledDocumentAndDisplay() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) MakeDocumentWithContentsOfURL() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) AllowsAutomaticShareMenu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DocumentForWindow() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) OpenDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) OpenDocumentWithContentsOfURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) NoteNewRecentDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) ValidateUserInterfaceItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) AutosavingDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DocumentForURL() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) NewDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DocumentClassForType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) Documents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) CurrentDirectory() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) RecentDocumentURLs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) AddDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DuplicateDocumentWithContentsOfURL() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) StandardShareMenuItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) PresentError() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) MaximumRecentDocumentCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DefaultType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) URLsFromRunningOpenPanel() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) WillPresentError() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) ClearRecentDocuments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) CurrentDocument() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) DocumentClassNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) RemoveDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) ReviewUnsavedDocumentsWithAlertTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) TypeForContentsOfURL() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) SaveAllDocuments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) SharedDocumentController() (*NSDocumentController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) RunModalOpenPanel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) BeginOpenPanelWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) BeginOpenPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) ReopenDocumentForURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDocumentController) MakeDocumentForURL() (*NSDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

