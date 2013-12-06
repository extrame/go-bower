package bower

import (
	"reflect"
	"testing"
)

func TestParseBowerJSON(t *testing.T) {
	tests := []struct {
		data          []byte
		wantComponent Component
		wantError     bool
	}{
		{
			data:          []byte(`{"name":"foo"}`),
			wantComponent: Component{Name: "foo"},
		},
		{
			data:      []byte(``),
			wantError: true,
		},
	}
	for _, test := range tests {
		c, err := ParseBowerJSON(test.data)
		if err != nil {
			if !test.wantError {
				t.Errorf("%q: ParseBowerJSON error: %s", test.data, err)
			}
			continue
		}
		if !reflect.DeepEqual(test.wantComponent, *c) {
			t.Errorf("%q: want component == %+v, got %+v", test.data, test.wantComponent, *c)
		}
	}
}
