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

// NewStreamBytes returns new Stream from bytes
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

// Get returns n bytes from Buffer with []byte
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

// Bytes returns the bytes left from Buffer.
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

//Pad puts empty bytes (0x00) of le (len).
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

// Byte sets byte(unsign) got from buffer to value
func (bs *Stream) Byte() (byte, error) {
	return ReadEByte(bs.Get(ByteSize))
}

// SByte sets byte(sign) got from buffer to value
func (bs *Stream) SByte() (int8, error) {
	return ReadESByte(bs.Get(ByteSize))
}

// PutByte puts byte(unsign) from value to buffer
func (bs *Stream) PutByte(value byte) error {
	return bs.Put(WriteByte(value))
}

// PutSByte puts byte(sign) from value to buffer
func (bs *Stream) PutSByte(value int8) error {
	return bs.Put(WriteSByte(value))
}

// Short sets short(unsign) got from buffer to value
func (bs *Stream) Short() (uint16, error) {
	return ReadEUShort(bs.Get(ShortSize))
}

// SShort sets short(sign) got from buffer to value
func (bs *Stream) SShort() (int16, error) {
	return ReadEShort(bs.Get(ShortSize))
}

// LShort sets short(unsign) got from buffer as LittleEndian to value
func (bs *Stream) LShort() (uint16, error) {
	return ReadELUShort(bs.Get(ShortSize))
}

// LSShort sets short(sign) got from buffer as LittleEndian to value
func (bs *Stream) LSShort() (int16, error) {
	return ReadELShort(bs.Get(ShortSize))
}

// PutShort puts short(unsign) from value to buffer
func (bs *Stream) PutShort(value uint16) error {
	return bs.Put(WriteUShort(value))
}

// PutSShort puts short(sign) from value to buffer
func (bs *Stream) PutSShort(value int16) error {
	return bs.Put(WriteShort(value))
}

// PutLShort puts short(unsign) from value to buffer as LittleEndian
func (bs *Stream) PutLShort(value uint16) error {
	return bs.Put(WriteLUShort(value))
}

// PutLSShort puts short(sign) from value to buffer as LittleEndian
func (bs *Stream) PutLSShort(value int16) error {
	return bs.Put(WriteLShort(value))
}

// Int sets int got from buffer to value
func (bs *Stream) Int() (int32, error) {
	return ReadEInt(bs.Get(IntSize))
}

// PutInt puts int from value to buffer
func (bs *Stream) PutInt(value int32) error {
	return bs.Put(WriteInt(value))
}

// Int sets int got from buffer to value as LittleEndian
func (bs *Stream) LInt() (int32, error) {
	return ReadELInt(bs.Get(IntSize))
}

// PutInt puts int from value to buffer as LittleEndian
func (bs *Stream) PutLInt(value int32) error {
	return bs.Put(WriteLInt(value))
}

// Long sets long got from buffer to value
func (bs *Stream) Long() (int64, error) {
	return ReadELong(bs.Get(LongSize))
}

// PutLong puts long from value to buffer
func (bs *Stream) PutLong(value int64) error {
	return bs.Put(WriteLong(value))
}

// Long sets long got from buffer to value as LittleEndian
func (bs *Stream) LLong() (int64, error) {
	return ReadELLong(bs.Get(LongSize))
}

// PutLong puts long from value to buffer as LittleEndian
func (bs *Stream) PutLLong(value int64) error {
	return bs.Put(WriteLLong(value))
}

// Float sets float got from buffer to value
func (bs *Stream) Float() (float32, error) {
	return ReadEFloat(bs.Get(FloatSize))
}

// PutFloat puts float from value to buffer
func (bs *Stream) PutFloat(value float32) error {
	return bs.Put(WriteFloat(value))
}

// Float sets float got from buffer to value as LittleEndian
func (bs *Stream) LFloat() (float32, error) {
	return ReadELFloat(bs.Get(FloatSize))
}

// PutFloat puts float from value to buffer as LittleEndian
func (bs *Stream) PutLFloat(value float32) error {
	return bs.Put(WriteLFloat(value))
}

// Double sets double got from buffer to value
func (bs *Stream) Double() (float64, error) {
	return ReadEDouble(bs.Get(DoubleSize))
}

// PutFloat puts double from value to buffer
func (bs *Stream) PutDouble(value float64) error {
	return bs.Put(WriteDouble(value))
}

// Double sets double got from buffer to value as LittleEndian
func (bs *Stream) LDouble() (float64, error) {
	return ReadELDouble(bs.Get(DoubleSize))
}

// PutFloat puts double from value to buffer as LittleEndian
func (bs *Stream) PutLDouble(value float64) error {
	return bs.Put(WriteLDouble(value))
}

// Bool sets byte got from buffer as bool to value
func (bs *Stream) Bool() (bool, error) {
	val, err := bs.Byte()
	if err != nil {
		return false, err
	}

	return val != 0, nil
}

//PutBool puts bool as byte from value to buffer
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

// NewOrderStreamBytes returns new Stream from bytes
func NewOrderStreamBytes(order Order, b []byte) *OrderStream {
	return &OrderStream{
		Stream: NewStreamBytes(b),
		Order:  order,
	}
}

// OrderStream is a binary stream with order
type OrderStream struct {
	*Stream
	Order Order
}

// Short sets short(unsign) got from buffer to value
func (bs *OrderStream) Short() (value uint16, err error) {
	return value, Read(bs, bs.Order, &value)
}

// SShort sets short(sign) got from buffer to value
func (bs *OrderStream) SShort() (value int16, err error) {
	return value, Read(bs, bs.Order, &value)
}

// PutShort puts short(unsign) from value to buffer
func (bs *OrderStream) PutShort(value uint16) error {
	return Write(bs, bs.Order, value)
}

// PutSShort puts short(sign) from value to buffer
func (bs *OrderStream) PutSShort(value int16) error {
	return Write(bs, bs.Order, value)
}

// Int sets int got from buffer to value
func (bs *OrderStream) Int() (value int32, err error) {
	return value, Read(bs, bs.Order, &value)
}

// PutInt puts int from value to buffer
func (bs *OrderStream) PutInt(value int32) error {
	return Write(bs, bs.Order, value)
}

// Long sets long got from buffer to value
func (bs *OrderStream) Long() (value int64, err error) {
	return value, Read(bs, bs.Order, &value)
}

// PutLong puts long from value to buffer
func (bs *OrderStream) PutLong(value int64) error {
	return Write(bs, bs.Order, value)
}

// Float sets float got from buffer to value
func (bs *OrderStream) Float() (value float32, err error) {
	return value, Read(bs, bs.Order, &value)
}

// PutFloat puts float from value to buffer
func (bs *OrderStream) PutFloat(value float32) error {
	return Write(bs, bs.Order, value)
}

// Double sets double got from buffer to value
func (bs *OrderStream) Double() (value float64, err error) {
	return value, Read(bs, bs.Order, &value)
}

// PutFloat puts double from value to buffer
func (bs *OrderStream) PutDouble(value float64) error {
	return Write(bs, bs.Order, value)
}
