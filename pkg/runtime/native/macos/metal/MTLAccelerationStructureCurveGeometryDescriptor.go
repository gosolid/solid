//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureCurveGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureCurveGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) ControlPointBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) ControlPointStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetIndexBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) CurveEndCaps() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) RadiusBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetSegmentCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) ControlPointCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetIndexType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) CurveType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetCurveType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) ControlPointFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) IndexBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SegmentCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) RadiusFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetIndexBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetCurveBasis() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) RadiusBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) IndexBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) IndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetControlPointStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) SetRadiusBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) RadiusStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCurveGeometryDescriptor) CurveBasis() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

