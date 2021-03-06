package platforms

import (
	"reflect"
	"runtime"
	"testing"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func TestDefault(t *testing.T) {
	expected := specs.Platform{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
	}
	p := Default()
	if !reflect.DeepEqual(p, expected) {
		t.Fatalf("default platform not as expected: %#v != %#v", p, expected)
	}
}
