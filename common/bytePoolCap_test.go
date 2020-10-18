package common

import "testing"

// Tests - bytePool functionality.
func TestBytePool(t *testing.T) {
	var size = 4
	var width = 10
	var capWidth = 16

	// Check the width
	if BytePool.Width() != width {
		t.Fatalf("bytepool width invalid: got %v want %v", BytePool.Width(), width)
	}

	// Check with width cap
	if BytePool.WidthCap() != capWidth {
		t.Fatalf("bytepool capWidth invalid: got %v want %v", BytePool.WidthCap(), capWidth)
	}

	// Check that retrieved buffer are of the expected width
	b := BytePool.Get()
	if len(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", len(b), width)
	}
	if cap(b) != capWidth {
		t.Fatalf("bytepool length invalid: got %v want %v", cap(b), capWidth)
	}

	BytePool.Put(b)

	// Fill the pool beyond the capped pool size.
	for i := 0; i < size*2; i++ {
		BytePool.Put(make([]byte, BytePool.w))
	}

	b = BytePool.Get()
	if len(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", len(b), width)
	}
	if cap(b) != capWidth {
		t.Fatalf("bytepool length invalid: got %v want %v", cap(b), capWidth)
	}

	BytePool.Put(b)

	// Close the channel so we can iterate over it.
	close(BytePool.c)

	// Check the size of the pool.
	if len(BytePool.c) != size {
		t.Fatalf("bytepool size invalid: got %v want %v", len(BytePool.c), size)
	}

	// bufPoolNoCap := NewBytePoolCap(size, width, 0)
	// Check the width
	if BytePool.Width() != width {
		t.Fatalf("bytepool width invalid: got %v want %v", BytePool.Width(), width)
	}

	// Check with width cap
	if BytePool.WidthCap() != 0 {
		t.Fatalf("bytepool capWidth invalid: got %v want %v", BytePool.WidthCap(), 0)
	}
	b = BytePool.Get()
	if len(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", len(b), width)
	}
	if cap(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", cap(b), width)
	}
}
