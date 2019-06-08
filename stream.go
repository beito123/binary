package binary

/*
	Binary

	Copyright (c) 2018 beito

	This software is released under the MIT License.
	http://opensource.org/licenses/mit-license.php
*/

import (
	"errors"
)

var errNoEnough = errors.New("no enough buffer")

// NewStream returns new Stream
func NewStream() *Stream {
	return NewStreamBytes([]byte{})
}

// NewStreamBytes returns new Stream with bytes
func NewStreamBytes(b []byte) *Stream {
	return &Stream{
		buf:     b,
		correct: true,
	}
}

// Stream is basic binary stream.
type Stream struct {
	buf     []byte
	off     int
	correct bool
}

// Reset resets Buffer
func (bs *Stream) Reset() {
	bs.correct = true
	bs.off = 0
	bs.buf = []byte{}
}

// Off returns offset
func (bs *Stream) Off() int {
	return bs.off
}

// Get gets n bytes from the buffer
func (bs *Stream) Get(n int) []byte {
	off := bs.off
	if n > bs.Len() {
		n = bs.Len()
	}

	bs.off += n

	return bs.buf[off : off+n]
}

// Put puts value to buffer
func (bs *Stream) Put(value []byte) error {
	_, err := bs.Write(value)

	return err
}

// Bytes returns the bytes left from the buffer
func (bs *Stream) Bytes() []byte {
	return bs.buf[bs.off:]
}

// AllBytes return all bytes
func (bs *Stream) AllBytes() []byte {
	return bs.buf
}

// SetBytes sets bytes
func (bs *Stream) SetBytes(b []byte) {
	bs.Reset()

	bs.buf = b
}

// Len returns len the bytes left
func (bs *Stream) Len() int {
	return len(bs.buf[bs.off:])
}

// Skip skips n bytes on buffer
func (bs *Stream) Skip(n int) {
	if n > bs.Len() {
		n = bs.Len()
	}

	bs.off += n
}

// Pad puts empty bytes (0x00) of le (len).
func (bs *Stream) Pad(le int) error {
	return bs.Put(make([]byte, le))
}

// Read reads and sets p
func (bs *Stream) Read(p []byte) (n int, err error) {
	return copy(p, bs.Get(len(p))), nil
}

// Write writes p
func (bs *Stream) Write(p []byte) (n int, err error) {
	bs.buf = append(bs.buf, p...)

	return len(p), nil
}

/*
 * Data types
 * | name  | size | encode |                   range                   |
 *  Byte    1byte   Big                        0 - 255
 *  SByte   1byte   Big                     -128 - 127
 *  Short   2bytes  Big                        0 - 65535
 *  SShort  2bytes  Big                   -32768 - 32767
 *  LShort  2bytes  Little                     0 - 65535
 *  LSShort 2bytes  Little                -32768 - 32767
 *  Int     4bytes  Big              -2147483648 - 2147483647
 *  Long    8bytes  Big     -9223372036854775808 - 9223372036854775807
 *  String  ?bytes  Big                        ? - ?
 *  Float   4bytes  Big        IEEE-754 32bits floating-point numbers
 *  LFloat  4bytes  Little     IEEE-754 32bits floating-point numbers
 *  Double  8bytes  Big        IEEE-754 64bits floating-point numbers
 *  LDouble 8bytes  Little     IEEE-754 64bits floating-point numbers
 */

// Byte gets an unsigned byte
func (bs *Stream) Byte() (byte, error) {
	return ReadEByte(bs.Get(ByteSize))
}

// SByte gets a signed byte
func (bs *Stream) SByte() (int8, error) {
	return ReadESByte(bs.Get(ByteSize))
}

// PutByte puts an unsigned byte
func (bs *Stream) PutByte(value byte) error {
	return bs.Put(WriteByte(value))
}

// PutSByte puts a signed byte
func (bs *Stream) PutSByte(value int8) error {
	return bs.Put(WriteSByte(value))
}

// Short gets an unsigned short
func (bs *Stream) Short() (uint16, error) {
	return ReadEUShort(bs.Get(ShortSize))
}

// SShort gets a signed short
func (bs *Stream) SShort() (int16, error) {
	return ReadEShort(bs.Get(ShortSize))
}

// LShort gets an unsigned short with LittleEndian
func (bs *Stream) LShort() (uint16, error) {
	return ReadELUShort(bs.Get(ShortSize))
}

// LSShort gets a signed short with LittleEndian
func (bs *Stream) LSShort() (int16, error) {
	return ReadELShort(bs.Get(ShortSize))
}

// PutShort puts an unsigned short
func (bs *Stream) PutShort(value uint16) error {
	return bs.Put(WriteUShort(value))
}

// PutSShort puts a signed short
func (bs *Stream) PutSShort(value int16) error {
	return bs.Put(WriteShort(value))
}

// PutLShort puts an unsigned short with LittleEndian
func (bs *Stream) PutLShort(value uint16) error {
	return bs.Put(WriteLUShort(value))
}

// PutLSShort puts a signed short with LittleEndian
func (bs *Stream) PutLSShort(value int16) error {
	return bs.Put(WriteLShort(value))
}

// Int gets a signed int
func (bs *Stream) Int() (int32, error) {
	return ReadEInt(bs.Get(IntSize))
}

// PutInt puts a signed int
func (bs *Stream) PutInt(value int32) error {
	return bs.Put(WriteInt(value))
}

// UInt gets an unsigned int
func (bs *Stream) UInt() (uint32, error) {
	return ReadEUInt(bs.Get(IntSize))
}

// PutUInt puts an unsigned int
func (bs *Stream) PutUInt(value uint32) error {
	return bs.Put(WriteUInt(value))
}

// LInt gets a signed int with LittleEndian
func (bs *Stream) LInt() (int32, error) {
	return ReadELInt(bs.Get(IntSize))
}

// PutLInt puts a signed int with LittleEndian
func (bs *Stream) PutLInt(value int32) error {
	return bs.Put(WriteLInt(value))
}

// LUInt gets an unsigned int with LittleEndian
func (bs *Stream) LUInt() (uint32, error) {
	return ReadELUInt(bs.Get(IntSize))
}

// PutLUInt puts an unsigned int with LittleEndian
func (bs *Stream) PutLUInt(value uint32) error {
	return bs.Put(WriteLUInt(value))
}

// Long gets a signed long
func (bs *Stream) Long() (int64, error) {
	return ReadELong(bs.Get(LongSize))
}

// PutLong puts a signed long
func (bs *Stream) PutLong(value int64) error {
	return bs.Put(WriteLong(value))
}

// LLong gets a signed long with LittleEndian
func (bs *Stream) LLong() (int64, error) {
	return ReadELLong(bs.Get(LongSize))
}

// PutLLong puts a signed long with LittleEndian
func (bs *Stream) PutLLong(value int64) error {
	return bs.Put(WriteLLong(value))
}

// ULong gets an unsigned long
func (bs *Stream) ULong() (uint64, error) {
	return ReadEULong(bs.Get(LongSize))
}

// PutULong puts an unsigned long
func (bs *Stream) PutULong(value uint64) error {
	return bs.Put(WriteULong(value))
}

// LULong gets an unsigned long with LittleEndian
func (bs *Stream) LULong() (uint64, error) {
	return ReadELULong(bs.Get(LongSize))
}

// PutLULong puts an unsigned long with LittleEndian
func (bs *Stream) PutLULong(value uint64) error {
	return bs.Put(WriteLULong(value))
}

// Float gets a float
func (bs *Stream) Float() (float32, error) {
	return ReadEFloat(bs.Get(FloatSize))
}

// PutFloat puts a float
func (bs *Stream) PutFloat(value float32) error {
	return bs.Put(WriteFloat(value))
}

// LFloat gtes a float with LittleEndian
func (bs *Stream) LFloat() (float32, error) {
	return ReadELFloat(bs.Get(FloatSize))
}

// PutLFloat puts a float
func (bs *Stream) PutLFloat(value float32) error {
	return bs.Put(WriteLFloat(value))
}

// Double gets a double
func (bs *Stream) Double() (float64, error) {
	return ReadEDouble(bs.Get(DoubleSize))
}

// PutDouble puts a double
func (bs *Stream) PutDouble(value float64) error {
	return bs.Put(WriteDouble(value))
}

// LDouble gets a double with LittleEndian
func (bs *Stream) LDouble() (float64, error) {
	return ReadELDouble(bs.Get(DoubleSize))
}

// PutLDouble puts a double with LittleEndian
func (bs *Stream) PutLDouble(value float64) error {
	return bs.Put(WriteLDouble(value))
}

// Bool gets a byte and returns as bool
func (bs *Stream) Bool() (bool, error) {
	val, err := bs.Byte()
	if err != nil {
		return false, err
	}

	return val != 0, nil
}

// PutBool puts a byte as bool
func (bs *Stream) PutBool(value bool) error {
	var val byte
	if value {
		val = 1 // true
	}

	return bs.PutByte(val)
}

// NewOrderStream returns new Stream
func NewOrderStream(order Order) *OrderStream {
	return NewOrderStreamBytes(order, []byte{})
}

// NewOrderStreamBytes returns new Stream with bytes
func NewOrderStreamBytes(order Order, b []byte) *OrderStream {
	return &OrderStream{
		Stream: NewStreamBytes(b),
		Order:  order,
	}
}

// OrderStream is a binary stream with a order
type OrderStream struct {
	*Stream
	Order Order
}

func (bs *OrderStream) get(size int) ([]byte, error) {
	b := bs.Get(size)
	if len(b) != size {
		return nil, ErrNotEnought
	}

	return b, nil
}

// Short get an unsigned short with the order
func (bs *OrderStream) Short() (value uint16, err error) {
	b, err := bs.get(ShortSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.UShort(b), nil
}

// SShort get a signed short with the order
func (bs *OrderStream) SShort() (value int16, err error) {
	b, err := bs.get(ShortSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.Short(b), nil
}

// PutShort puts an unsigned short with the order
func (bs *OrderStream) PutShort(value uint16) error {
	return bs.Put(bs.Order.PutUShort(value))
}

// PutSShort puts a signed short with the order
func (bs *OrderStream) PutSShort(value int16) error {
	return bs.Put(bs.Order.PutShort(value))
}

// Int get a signed int with the order
func (bs *OrderStream) Int() (value int32, err error) {
	b, err := bs.get(IntSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.Int(b), nil
}

// PutInt puts a signed int with the order
func (bs *OrderStream) PutInt(value int32) error {
	return bs.Put(bs.Order.PutInt(value))
}

// UInt get an unsigned int with the order
func (bs *OrderStream) UInt() (value uint32, err error) {
	b, err := bs.get(IntSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.UInt(b), nil
}

// PutUInt puts an unsigned int with the order
func (bs *OrderStream) PutUInt(value uint32) error {
	return bs.Put(bs.Order.PutUInt(value))
}

// Long gets a signed long with the order
func (bs *OrderStream) Long() (value int64, err error) {
	b, err := bs.get(LongSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.Long(b), nil
}

// PutLong puts a signed long with the order
func (bs *OrderStream) PutLong(value int64) error {
	return bs.Put(bs.Order.PutLong(value))
}

// ULong gets an unsigned long with the order
func (bs *OrderStream) ULong() (value uint64, err error) {
	b, err := bs.get(LongSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.ULong(b), nil
}

// PutULong puts an unsigned long with the order
func (bs *OrderStream) PutULong(value uint64) error {
	return bs.Put(bs.Order.PutULong(value))
}

// Float gets a float with the order
func (bs *OrderStream) Float() (value float32, err error) {
	b, err := bs.get(FloatSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.Float(b), nil
}

// PutFloat puts a float with the order
func (bs *OrderStream) PutFloat(value float32) error {
	return bs.Put(bs.Order.PutFloat(value))
}

// Double gets a double with the order
func (bs *OrderStream) Double() (value float64, err error) {
	b, err := bs.get(DoubleSize)
	if err != nil {
		return 0, err
	}

	return bs.Order.Double(b), nil
}

// PutDouble puts a double with the order
func (bs *OrderStream) PutDouble(value float64) error {
	return bs.Put(bs.Order.PutDouble(value))
}
