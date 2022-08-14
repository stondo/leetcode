package mirrorrefelction

import (
	"testing"
)

// TestMirrorReflection1 calls mirrorrefelction.mirrorReflection with 2 int:
// p is the length of the side of asquare room covered with mirrors and
// q is the distance on East wall where the laser hits, checking
// for the first receptor hit.
func TestMirrorReflection1(t *testing.T) {
	p := 2
	q := 1

	expected := 2
	result := mirrorReflection(p, q)

	if result != expected {
		t.Fatalf(`mirrorReflection(p, q) = %v, want %d`, result, expected)
	}
}

// TestMirrorReflection2 calls mirrorrefelction.mirrorReflection with 2 int:
// p is the length of the side of asquare room covered with mirrors and
// q is the distance on East wall where the laser hits, checking
// for the first receptor hit.
func TestMirrorReflection2(t *testing.T) {
	p := 3
	q := 1

	expected := 1
	result := mirrorReflection(p, q)

	if result != expected {
		t.Fatalf(`mirrorReflection(p, q) = %v, want %d`, result, expected)
	}
}

// TestMirrorReflection3 calls mirrorrefelction.mirrorReflection with 2 int:
// p is the length of the side of asquare room covered with mirrors and
// q is the distance on East wall where the laser hits, checking
// for the first receptor hit.
func TestMirrorReflection3(t *testing.T) {
	p := 4
	q := 3

	expected := 2
	result := mirrorReflection(p, q)

	if result != expected {
		t.Fatalf(`mirrorReflection(p, q) = %v, want %d`, result, expected)
	}
}
