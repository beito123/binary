package binary

/*
 * Binary
 *
 * Copyright (c) 2018 beito
 *
 * This software is released under the MIT License.
 * http://opensource.org/licenses/mit-license.php
 */

import (
	"bytes"
	"encoding/hex"
	"net"
	"strconv"
	"strings"
)

// NewStream returns new Stream
func NewStream() *Stream {
	return &Stream{
		Buffer: bytes.NewBuffer([]byte{}),
	}
}

// NewStreamBytes returns new Stream from bytes
func NewStreamBytes(b []byte) *Stream {
	return &Stream{
		Buffer: bytes.NewBuffer(b),
	}
}

// NewStreamString returns new Stream from string
func NewStreamString(s string) *Stream {
	return &Stream{
		Buffer: bytes.NewBufferString(s),
	}
}

// Stream is basic binary stream.
type Stream struct {
	Buffer *bytes.Buffer
}

// Reset resets Buffer
func (bs *Stream) Reset() error {
	bs.Buffer.Reset()
	return nil
}

// Get returns n bytes from Buffer with []byte
func (bs *Stream) Get(n int) []byte {
	return bs.Buffer.Next(n)
}

// Put puts value to buffer
func (bs *Stream) Put(value []byte) error {
	return Write(bs.Buffer, BigEndian, value)
}

// Bytes returns the bytes left from Buffer.
func (bs *Stream) Bytes() []byte {
	return bs.Buffer.Bytes()
}

// Len returns len the bytes left
func (bs *Stream) Len() int {
	return bs.Buffer.Len()
}

// Skip skips n bytes on buffer
func (bs *Stream) Skip(n int) {
	_ = bs.Buffer.Next(n)
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
 *  Triad   3bytes  Little                     0 - 16777215
 *  Int     4bytes  Big              -2147483648 - 2147483647
 *  Long    8bytes  Big     -9223372036854775808 - 9223372036854775807
 *  String  ?bytes  Big                        ? - ?
 */

/*
 * Byte
 * SignedByte
 * Short
 * SignedShort
 * LShort
 * SignedLShort
 * Triad
 * LTriad
 * Int
 * Float
 * LFloat
 * Double
 * LDouble
 * Long
 * /////////////
 * Bool
 * String
 * HexString
 * Address
 * NBT
 * Item
 * UUID
 * Position
 * BlockPosition
 * EntityMetadata
 */

// Byte sets byte(unsign) got from buffer to value
func (bs *Stream) Byte(value *byte) error {
	return Read(bs.Buffer, BigEndian, value)
}

// SByte sets byte(sign) got from buffer to value
func (bs *Stream) SByte(value *int8) error {
	return Read(bs.Buffer, BigEndian, value)
}

// PutByte puts byte(unsign) from value to buffer
func (bs *Stream) PutByte(value byte) error {
	return Write(bs.Buffer, BigEndian, value)
}

// PutSByte puts byte(sign) from value to buffer
func (bs *Stream) PutSByte(value int8) error {
	return Write(bs.Buffer, BigEndian, value)
}

// Short sets short(unsign) got from buffer to value
func (bs *Stream) Short(value *uint16) error {
	return Read(bs.Buffer, BigEndian, value)
}

// SShort sets short(sign) got from buffer to value
func (bs *Stream) SShort(value *int16) error {
	return Read(bs.Buffer, BigEndian, value)
}

// LShort sets short(unsign) got from buffer as LittleEndian to value
func (bs *Stream) LShort(value *uint16) error {
	return Read(bs.Buffer, LittleEndian, value)
}

// LSShort sets short(sign) got from buffer as LittleEndian to value
func (bs *Stream) LSShort(value *int16) error {
	return Read(bs.Buffer, LittleEndian, value)
}

// PutShort puts short(unsign) from value to buffer
func (bs *Stream) PutShort(value uint16) error {
	return Write(bs.Buffer, BigEndian, value)
}

// PutSShort puts short(sign) from value to buffer
func (bs *Stream) PutSShort(value int16) error {
	return Write(bs.Buffer, BigEndian, value)
}

// PutLShort puts short(unsign) from value to buffer as LittleEndian
func (bs *Stream) PutLShort(value uint16) error {
	return Write(bs.Buffer, LittleEndian, value)
}

// PutLSShort puts short(sign) from value to buffer as LittleEndian
func (bs *Stream) PutLSShort(value int16) error {
	return Write(bs.Buffer, LittleEndian, value)
}

// Triad sets triad got from buffer to value
func (bs *Stream) Triad(value *Triad) error {
	return Read(bs.Buffer, BigEndian, value)
}

// PutTriad puts triad from value to buffer
func (bs *Stream) PutTriad(value Triad) error {
	return Write(bs.Buffer, BigEndian, value)
}

// LTriad sets triad got from buffer as LittleEndian to value
func (bs *Stream) LTriad(value *Triad) error {
	return Read(bs.Buffer, LittleEndian, value)
}

// PutLTriad puts triad from value to buffer as LittleEndian
func (bs *Stream) PutLTriad(value Triad) error {
	return Write(bs.Buffer, LittleEndian, value)
}

// Int sets int got from buffer to value
func (bs *Stream) Int(value *int32) error {
	return Read(bs.Buffer, BigEndian, value)
}

// PutInt puts int from value to buffer
func (bs *Stream) PutInt(value int32) error {
	return Write(bs.Buffer, BigEndian, value)
}

// Long sets long got from buffer to value
func (bs *Stream) Long(value *int64) error {
	return Read(bs.Buffer, BigEndian, value)
}

// PutLong puts long from value to buffer
func (bs *Stream) PutLong(value int64) error {
	return Write(bs.Buffer, BigEndian, value)
}

// Bool sets byte got from buffer as bool to value
func (bs *Stream) Bool(value *bool) error {
	var val byte

	err := bs.Byte(&val)
	if err != nil {
		return err
	}

	*value = (val != 0)

	return nil
}

//PutBool puts bool as byte from value to buffer
func (bs *Stream) PutBool(value bool) error {
	var val byte
	if value {
		val = 1 // true
	}

	return bs.PutByte(val)
}

// String sets string(len short, str string) got from buffer to value
func (bs *Stream) String(value *string) error {
	var n uint16
	err := bs.Short(&n)
	if err != nil {
		return err
	}

	*value = string(bs.Get(int(n)))
	return nil
}

// PutString puts string(len short, str string) to Buffer
func (bs *Stream) PutString(value string) error {
	n := uint16(len(value))
	err := bs.PutShort(n)
	if err != nil {
		return err
	}
	return bs.Put([]byte(value))
}

// HexString gets hex string from Buffer (for Magic)
func (bs *Stream) HexString(n int, value *string) {
	*value = hex.EncodeToString(bs.Buffer.Bytes())
}

// PutHexString puts hex string to Buffer
func (bs *Stream) PutHexString(value string) error {
	bytes, err := hex.DecodeString(value)
	if err != nil {
		return err
	}
	return bs.Put(bytes)
}

// Address sets address got from Buffer to addr and port
// address(version byte, address byte x4, port ushort)
func (bs *Stream) Address(addr *string, port *uint16) error {
	var version byte
	err := bs.Byte(&version)
	if err != nil {
		return err
	}

	var address string

	if version == 4 {
		var bytes byte
		for i := 0; i < 4; i++ {
			err = bs.Byte(&bytes)
			if err != nil {
				return err
			}

			address = address + strconv.Itoa(int(^bytes&0xff))
			if i < 3 {
				address = address + "."
			}
		}
		addr = &address

		err = bs.Short(port)
		if err != nil {
			return err
		}
	} else {
		//IPv6
	}

	return nil
}

// PutAddress puts address to Buffer
// address(version byte, address byte x4, port ushort)
func (bs *Stream) PutAddress(addr string, port, version uint16) error {
	err := bs.PutByte(byte(version))
	if err != nil {
		return err
	}

	if version == 4 {
		for _, str := range strings.Split(addr, ".") {
			i, _ := strconv.Atoi(str)
			err = bs.PutByte(^byte(i) & 0xff)
			if err != nil {
				return err
			}
		}
		err = bs.PutShort(port)
		if err != nil {
			return err
		}
	} else {
		//ipv6
	}

	return nil
}


// AddressUDPAddr sets address got from Buffer to UDPAddr
func (bs *Stream) AddressUDPAddr(addr *net.UDPAddr) error {
	var add string
	var port uint16

	err := bs.Address(&add, &port)
	if err != nil {
		return err
	}

	naddr := net.UDPAddr{
		IP:   net.ParseIP(add),
		Port: int(port),
		//Zone: "", // TODO: support ipv6
	}

	*addr = naddr

	return nil
}

// PutAddressUDPAddr puts address from UDPAddr to Buffer
func (bs *Stream) PutAddressUDPAddr(addr net.UDPAddr) error {
	var ver uint16
	switch len(addr.IP) {
	case net.IPv6len:
		ver = 6
	case net.IPv4len:
		fallthrough
	default:
		ver = 4
	}

	return bs.PutAddress(addr.IP.String(), uint16(addr.Port), ver)
}

// AddressTCPAddr sets address got from Buffer to TCPAddr
func (bs *Stream) AddressTCPAddr(addr *net.TCPAddr) error {
	var add string
	var port uint16

	err := bs.Address(&add, &port)
	if err != nil {
		return err
	}

	naddr := net.TCPAddr{
		IP:   net.ParseIP(add),
		Port: int(port),
		//Zone: "", // TODO: support ipv6
	}

	*addr = naddr

	return nil
}

// PutAddressTCPAddr puts address from TCPAddr to Buffer
func (bs *Stream) PutAddressTCPAddr(addr net.TCPAddr) error {
	var ver uint16
	switch len(addr.IP) {
	case net.IPv6len:
		ver = 6
	case net.IPv4len:
		fallthrough
	default:
		ver = 4
	}

	return bs.PutAddress(addr.IP.String(), uint16(addr.Port), ver)
}