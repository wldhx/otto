package foundation

import (
	"testing"
)

func TestMock_impl(t *testing.T) {
	var _ Foundation = new(Mock)
}
