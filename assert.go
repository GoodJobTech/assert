package assert

import "testing"

func equal[T comparable](a, b T) bool {
	return a == b
}

func Equal[T comparable](t *testing.T, expected, actual T) {
	if !equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func NotEqual[T comparable](t *testing.T, expected, actual T) {
	if equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func Nil(t *testing.T, obj any) {
	if obj != nil {
		t.Fatalf("Expected nil, got %v", obj)
	}
}

func NotNil(t *testing.T, obj any) {
	if obj == nil {
		t.Fatalf("Expected not nil, got nil")
	}
}

func contains[T comparable](haystack []T, needle T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

func Contains[T comparable](t *testing.T, haystack []T, needle T) {
	if !contains(haystack, needle) {
		t.Fatalf("Expected %v to contain %v", haystack, needle)
	}
}

func NotContains[T comparable](t *testing.T, haystack []T, needle T) {
	if contains(haystack, needle) {
		t.Fatalf("Expected %v to not contain %v", haystack, needle)
	}
}

func True(t *testing.T, condition bool) {
	if !condition {
		t.Fatalf("Expected true, got false")
	}
}

func False(t *testing.T, condition bool) {
	if condition {
		t.Fatalf("Expected false, got true")
	}
}
