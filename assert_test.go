package assert

import "testing"

func TestAssert(t *testing.T) {
	Equal(t, 1, 1)

	NotEqual(t, "abc", "def")

	Nil(t, nil)

	NotNil(t, "")

	Contains(t, []int{1, 2, 3}, 1)

	NotContains(t, []int{1, 2, 3}, 4)

	True(t, true)

	False(t, false)
}
