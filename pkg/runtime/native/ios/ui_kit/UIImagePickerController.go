//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIImagePickerController : UIKit.UINavigationController
*/

type UIImagePickerController struct {
  *UINavigationController

}

func (r *UIImagePickerController) TakePicture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) ShowsCameraControls() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetShowsCameraControls() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetCameraOverlayView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetCameraCaptureMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) IsCameraDeviceAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) AllowsImageEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetAllowsImageEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetCameraViewTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SourceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) MediaTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetMediaTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) VideoMaximumDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetVideoQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) CameraOverlayView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetCameraDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) IsSourceTypeAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetSourceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) VideoQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) CameraViewTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) CameraDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) AvailableMediaTypesForSourceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) IsFlashAvailableForCameraDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) StopVideoCapture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetAllowsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) CameraCaptureMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetCameraFlashMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) AllowsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) VideoExportPreset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetVideoExportPreset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) CameraFlashMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) AvailableCaptureModesForCameraDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) StartVideoCapture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) ImageExportPreset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetImageExportPreset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImagePickerController) SetVideoMaximumDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

