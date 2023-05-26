//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: payload.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package proto

import "errors"
import "github/lwileczek/goBenchmarkSerialization/types/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding SubStruct model
type SubStructModel struct {
    // Model buffer
    buffer *fbe.Buffer

    // Field model
    model *FieldModelSubStruct
}

// Create a new SubStruct model
func NewSubStructModel(buffer *fbe.Buffer) *SubStructModel {
    return &SubStructModel{buffer: buffer, model: NewFieldModelSubStruct(buffer, 4)}
}

// Get the model buffer
func (m *SubStructModel) Buffer() *fbe.Buffer { return m.buffer }
// Get the field model
func (m *SubStructModel) Model() *FieldModelSubStruct { return m.model }

// Get the model size
func (m *SubStructModel) FBESize() int { return m.model.FBESize() + m.model.FBEExtra() }
// // Get the model type
func (m *SubStructModel) FBEType() int { return m.model.FBEType() }

// Check if the struct value is valid
func (m *SubStructModel) Verify() bool {
    if (m.buffer.Offset() + m.model.FBEOffset() - 4) > m.buffer.Size() {
        return false
    }

    fbeFullSize := int(fbe.ReadUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4))
    if fbeFullSize < m.model.FBESize() {
        return false
    }

    return m.model.Verify()
}

// Create a new model (begin phase)
func (m *SubStructModel) CreateBegin() int {
    fbeBegin := m.buffer.Allocate(4 + m.model.FBESize())
    return fbeBegin
}

// Create a new model (end phase)
func (m *SubStructModel) CreateEnd(fbeBegin int) int {
    fbeEnd := m.buffer.Size()
    fbeFullSize := fbeEnd - fbeBegin
    fbe.WriteUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4, uint32(fbeFullSize))
    return fbeFullSize
}

// Serialize the struct value
func (m *SubStructModel) Serialize(value *SubStruct) (int, error) {
    fbeBegin := m.CreateBegin()
    err := m.model.Set(value)
    fbeFullSize := m.CreateEnd(fbeBegin)
    return fbeFullSize, err
}

// Deserialize the struct value
func (m *SubStructModel) Deserialize() (*SubStruct, int, error) {
    value := NewSubStruct()
    fbeFullSize, err := m.DeserializeValue(value)
    return value, fbeFullSize, err
}

// Deserialize the struct value by the given pointer
func (m *SubStructModel) DeserializeValue(value *SubStruct) (int, error) {
    if (m.buffer.Offset() + m.model.FBEOffset() - 4) > m.buffer.Size() {
        value = NewSubStruct()
        return 0, nil
    }

    fbeFullSize := int(fbe.ReadUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4))
    if fbeFullSize < m.model.FBESize() {
        value = NewSubStruct()
        return 0, errors.New("model is broken")
    }

    err := m.model.GetValue(value)
    return fbeFullSize, err
}

// Move to the next struct value
func (m *SubStructModel) Next(prev int) {
    m.model.FBEShift(prev)
}
