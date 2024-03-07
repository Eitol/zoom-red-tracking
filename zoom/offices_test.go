package zoom

import (
	"reflect"
	"testing"
)

func Test_getOffices(t *testing.T) {
	got, err := getOffices()
	if err != nil {
		t.Errorf("getOffices() error = %v, wantErr %v", err, nil)
		return
	}
	if reflect.TypeOf(got).String() != "[]zoom.CourierOffice" {
		t.Errorf("getOffices() = %v, want %v", reflect.TypeOf(got).String(), "[]zoom.CourierOffice")
	}
}
