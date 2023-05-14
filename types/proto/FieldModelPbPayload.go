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

// Fast Binary Encoding PbPayload field model
type FieldModelPbPayload struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    StringEntry *fbe.FieldModelString
    SmallInteger *fbe.FieldModelInt32
    NormalInteger *fbe.FieldModelInt64
    SerializationMethod *fbe.FieldModelString
    Boolean *fbe.FieldModelBool
    SomeFloat *fbe.FieldModelFloat
    IntArray *FieldModelVectorInt32
    Chart *FieldModelMapStringInt32
    SubShop *FieldModelSubStruct
}

// Create a new PbPayload field model
func NewFieldModelPbPayload(buffer *fbe.Buffer, offset int) *FieldModelPbPayload {
    fbeResult := FieldModelPbPayload{buffer: buffer, offset: offset}
    fbeResult.StringEntry = fbe.NewFieldModelString(buffer, 4 + 4)
    fbeResult.SmallInteger = fbe.NewFieldModelInt32(buffer, fbeResult.StringEntry.FBEOffset() + fbeResult.StringEntry.FBESize())
    fbeResult.NormalInteger = fbe.NewFieldModelInt64(buffer, fbeResult.SmallInteger.FBEOffset() + fbeResult.SmallInteger.FBESize())
    fbeResult.SerializationMethod = fbe.NewFieldModelString(buffer, fbeResult.NormalInteger.FBEOffset() + fbeResult.NormalInteger.FBESize())
    fbeResult.Boolean = fbe.NewFieldModelBool(buffer, fbeResult.SerializationMethod.FBEOffset() + fbeResult.SerializationMethod.FBESize())
    fbeResult.SomeFloat = fbe.NewFieldModelFloat(buffer, fbeResult.Boolean.FBEOffset() + fbeResult.Boolean.FBESize())
    fbeResult.IntArray = NewFieldModelVectorInt32(buffer, fbeResult.SomeFloat.FBEOffset() + fbeResult.SomeFloat.FBESize())
    fbeResult.Chart = NewFieldModelMapStringInt32(buffer, fbeResult.IntArray.FBEOffset() + fbeResult.IntArray.FBESize())
    fbeResult.SubShop = NewFieldModelSubStruct(buffer, fbeResult.Chart.FBEOffset() + fbeResult.Chart.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelPbPayload) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelPbPayload) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.StringEntry.FBESize() +
        fm.SmallInteger.FBESize() +
        fm.NormalInteger.FBESize() +
        fm.SerializationMethod.FBESize() +
        fm.Boolean.FBESize() +
        fm.SomeFloat.FBESize() +
        fm.IntArray.FBESize() +
        fm.Chart.FBESize() +
        fm.SubShop.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelPbPayload) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.StringEntry.FBEExtra() +
        fm.SmallInteger.FBEExtra() +
        fm.NormalInteger.FBEExtra() +
        fm.SerializationMethod.FBEExtra() +
        fm.Boolean.FBEExtra() +
        fm.SomeFloat.FBEExtra() +
        fm.IntArray.FBEExtra() +
        fm.Chart.FBEExtra() +
        fm.SubShop.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelPbPayload) FBEType() int { return 2 }

// Get the field offset
func (fm *FieldModelPbPayload) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelPbPayload) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelPbPayload) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelPbPayload) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelPbPayload) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelPbPayload) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelPbPayload) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.StringEntry.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.StringEntry.Verify() {
        return false
    }
    fbeCurrentSize += fm.StringEntry.FBESize()

    if (fbeCurrentSize + fm.SmallInteger.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.SmallInteger.Verify() {
        return false
    }
    fbeCurrentSize += fm.SmallInteger.FBESize()

    if (fbeCurrentSize + fm.NormalInteger.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.NormalInteger.Verify() {
        return false
    }
    fbeCurrentSize += fm.NormalInteger.FBESize()

    if (fbeCurrentSize + fm.SerializationMethod.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.SerializationMethod.Verify() {
        return false
    }
    fbeCurrentSize += fm.SerializationMethod.FBESize()

    if (fbeCurrentSize + fm.Boolean.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Boolean.Verify() {
        return false
    }
    fbeCurrentSize += fm.Boolean.FBESize()

    if (fbeCurrentSize + fm.SomeFloat.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.SomeFloat.Verify() {
        return false
    }
    fbeCurrentSize += fm.SomeFloat.FBESize()

    if (fbeCurrentSize + fm.IntArray.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.IntArray.Verify() {
        return false
    }
    fbeCurrentSize += fm.IntArray.FBESize()

    if (fbeCurrentSize + fm.Chart.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Chart.Verify() {
        return false
    }
    fbeCurrentSize += fm.Chart.FBESize()

    if (fbeCurrentSize + fm.SubShop.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.SubShop.Verify() {
        return false
    }
    fbeCurrentSize += fm.SubShop.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelPbPayload) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelPbPayload) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelPbPayload) Get() (*PbPayload, error) {
    fbeResult := NewPbPayload()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelPbPayload) GetValue(fbeValue *PbPayload) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelPbPayload) GetFields(fbeValue *PbPayload, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.StringEntry.FBESize()) <= fbeStructSize {
        fbeValue.StringEntry, _ = fm.StringEntry.Get()
    } else {
        fbeValue.StringEntry = ""
    }
    fbeCurrentSize += fm.StringEntry.FBESize()

    if (fbeCurrentSize + fm.SmallInteger.FBESize()) <= fbeStructSize {
        fbeValue.SmallInteger, _ = fm.SmallInteger.Get()
    } else {
        fbeValue.SmallInteger = 0
    }
    fbeCurrentSize += fm.SmallInteger.FBESize()

    if (fbeCurrentSize + fm.NormalInteger.FBESize()) <= fbeStructSize {
        fbeValue.NormalInteger, _ = fm.NormalInteger.Get()
    } else {
        fbeValue.NormalInteger = 0
    }
    fbeCurrentSize += fm.NormalInteger.FBESize()

    if (fbeCurrentSize + fm.SerializationMethod.FBESize()) <= fbeStructSize {
        fbeValue.SerializationMethod, _ = fm.SerializationMethod.Get()
    } else {
        fbeValue.SerializationMethod = ""
    }
    fbeCurrentSize += fm.SerializationMethod.FBESize()

    if (fbeCurrentSize + fm.Boolean.FBESize()) <= fbeStructSize {
        fbeValue.Boolean, _ = fm.Boolean.Get()
    } else {
        fbeValue.Boolean = false
    }
    fbeCurrentSize += fm.Boolean.FBESize()

    if (fbeCurrentSize + fm.SomeFloat.FBESize()) <= fbeStructSize {
        fbeValue.SomeFloat, _ = fm.SomeFloat.Get()
    } else {
        fbeValue.SomeFloat = 0.0
    }
    fbeCurrentSize += fm.SomeFloat.FBESize()

    if (fbeCurrentSize + fm.IntArray.FBESize()) <= fbeStructSize {
        fbeValue.IntArray, _ = fm.IntArray.Get()
    } else {
        fbeValue.IntArray = make([]int32, 0)
    }
    fbeCurrentSize += fm.IntArray.FBESize()

    if (fbeCurrentSize + fm.Chart.FBESize()) <= fbeStructSize {
        fbeValue.Chart, _ = fm.Chart.Get()
    } else {
        fbeValue.Chart = make(map[string]int32)
    }
    fbeCurrentSize += fm.Chart.FBESize()

    if (fbeCurrentSize + fm.SubShop.FBESize()) <= fbeStructSize {
        _ = fm.SubShop.GetValue(&fbeValue.SubShop)
    } else {
        fbeValue.SubShop = *NewSubStruct()
    }
    fbeCurrentSize += fm.SubShop.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelPbPayload) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelPbPayload) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelPbPayload) Set(fbeValue *PbPayload) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelPbPayload) SetFields(fbeValue *PbPayload) error {
    var err error = nil

    if err = fm.StringEntry.Set(fbeValue.StringEntry); err != nil {
        return err
    }
    if err = fm.SmallInteger.Set(fbeValue.SmallInteger); err != nil {
        return err
    }
    if err = fm.NormalInteger.Set(fbeValue.NormalInteger); err != nil {
        return err
    }
    if err = fm.SerializationMethod.Set(fbeValue.SerializationMethod); err != nil {
        return err
    }
    if err = fm.Boolean.Set(fbeValue.Boolean); err != nil {
        return err
    }
    if err = fm.SomeFloat.Set(fbeValue.SomeFloat); err != nil {
        return err
    }
    if err = fm.IntArray.Set(fbeValue.IntArray); err != nil {
        return err
    }
    if err = fm.Chart.Set(fbeValue.Chart); err != nil {
        return err
    }
    if err = fm.SubShop.Set(&fbeValue.SubShop); err != nil {
        return err
    }
    return err
}