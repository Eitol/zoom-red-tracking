package zoom

import (
	"testing"
)

func TestHttpClient_GetTrackingInfo(t *testing.T) {
	client := NewDefaultClient()
	tests := []struct {
		name     string
		tracking int
		want     *GetTrackingResponse
		wantErr  bool
	}{
		{
			name:     "Test 1",
			tracking: 1553824799,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracking, err := client.GetTrackingInfo(tt.tracking)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTrackingInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tracking == nil {
				t.Errorf("GetTrackingInfo() tracking = %v, want %v", tracking, tt.want)
			}
		})
	}
}
