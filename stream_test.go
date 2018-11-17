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
	"testing"
)

// Magic is a bytes data for testing
var Magic = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var MagicLen = len(Magic)

func TestStreamReset(t *testing.T) {
	stream := NewStreamBytes(Magic)

	stream.Reset()

	ln := stream.Len()
	if ln != 0 {
		t.Logf("Expected %d for offset, but %d", 0, ln)
	}
}

func TestStreamOff(t *testing.T) {
	stream := NewStreamBytes(Magic)

	stream.Get(MagicLen)

	off := stream.Off()
	if off != MagicLen {
		t.Logf("Expected %d for offset, but %d", MagicLen, off)
	}
}

func TestStreamGet(t *testing.T) {
	stream := NewStreamBytes(Magic)

	// get 5 bytes
	exp := Magic[:5]
	ret := stream.Get(5)
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}

	// get all bytes
	exp = Magic[5:]
	ret = stream.Get(len(exp))
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}

	// get overflow
	exp = []byte{}
	ret = stream.Get(1)
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}

	stream.SetBytes(Magic)

	exp = Magic
	ret = stream.Get(MagicLen + 1)
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}
}

func TestStreamPut(t *testing.T) {
	stream := NewStream()

	exp := Magic
	if err := stream.Put(exp); err != nil {
		t.Fatalf("Failed to put bytes")
	}

	ret := stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}

	exp = append(exp, Magic...)
	if err := stream.Put(Magic); err != nil {
		t.Fatalf("Failed to put bytes Error: %s", err)
	}

	ret = stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for got bytes, but %d", exp, ret)
	}
}

func TestStreamBytes(t *testing.T) {
	stream := NewStreamBytes(Magic)

	exp := Magic
	ret := stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}

	stream.Skip(5)

	exp = Magic[5:]
	ret = stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}
}

func TestStreamAllBytes(t *testing.T) {
	stream := NewStreamBytes(Magic)

	exp := Magic
	ret := stream.AllBytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}
}

func TestStreamSetBytes(t *testing.T) {
	stream := NewStreamBytes([]byte{0xff, 0xff, 0xff, 0xff})

	stream.Skip(2) // change offset

	// check resetting offset, bytes data
	exp := Magic
	stream.SetBytes(exp)

	ret := stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}
}

func TestStreamLen(t *testing.T) {
	stream := NewStreamBytes(Magic)

	stream.Get(MagicLen)

	ln := stream.Len()
	if ln != 0 {
		t.Logf("Expected %d for offset, but %d", 0, ln)
	}
}

func TestStreamSkip(t *testing.T) {
	stream := NewStreamBytes(Magic)

	// check moving offset correctly
	exp := Magic[5:]
	stream.Skip(5)

	ret := stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}

	exp = Magic[5+10:]
	stream.Skip(10)

	ret = stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}

	stream.SetBytes(Magic)

	// check over skipping
	exp = []byte{}
	stream.Skip(stream.Len() + 1)

	ret = stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}
}

func TestStreamPad(t *testing.T) {
	stream := NewStream()

	exp := make([]byte, 8)
	if err := stream.Pad(8); err != nil {
		t.Fatalf("Failed to put bytes Error: %s", err)
	}

	ret := stream.Bytes()
	if !bytes.Equal(ret, exp) {
		t.Fatalf("Expected %d for bytes, but %d", exp, ret)
	}
}
