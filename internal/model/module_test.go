package model

import "testing"

func TestModuleStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status ModuleStatus
		want   string
	}{
		{
			name:   "active status",
			status: StatusActive,
			want:   "Active",
		},
		{
			name:   "historical status",
			status: StatusHistorical,
			want:   "Historical",
		},
		{
			name:   "in process status",
			status: StatusInProcess,
			want:   "In Process",
		},
		{
			name:   "unknown status",
			status: ModuleStatus(99),
			want:   "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("ModuleStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
