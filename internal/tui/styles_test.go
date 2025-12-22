package tui

import (
	"strings"
	"testing"

	"github.com/ethanolivertroy/cmvp-tui/internal/model"
)

func TestStatusBadge(t *testing.T) {
	tests := []struct {
		name     string
		status   model.ModuleStatus
		contains string
		empty    bool
	}{
		{
			name:     "active status",
			status:   model.StatusActive,
			contains: "ACTIVE",
			empty:    false,
		},
		{
			name:     "historical status",
			status:   model.StatusHistorical,
			contains: "HISTORICAL",
			empty:    false,
		},
		{
			name:     "in process status",
			status:   model.StatusInProcess,
			contains: "IN PROCESS",
			empty:    false,
		},
		{
			name:   "unknown status",
			status: model.ModuleStatus(99),
			empty:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StatusBadge(tt.status)
			if tt.empty {
				if result != "" {
					t.Errorf("expected empty string for unknown status, got %q", result)
				}
			} else {
				if !strings.Contains(result, tt.contains) {
					t.Errorf("expected badge to contain %q, got %q", tt.contains, result)
				}
			}
		})
	}
}

func TestLevelBadge(t *testing.T) {
	tests := []struct {
		name     string
		level    int
		contains string
		empty    bool
	}{
		{
			name:  "level 0",
			level: 0,
			empty: true,
		},
		{
			name:     "level 1",
			level:    1,
			contains: "Level 1",
			empty:    false,
		},
		{
			name:     "level 2",
			level:    2,
			contains: "Level 2",
			empty:    false,
		},
		{
			name:     "level 3",
			level:    3,
			contains: "Level 3",
			empty:    false,
		},
		{
			name:     "level 4",
			level:    4,
			contains: "Level 4",
			empty:    false,
		},
		{
			name:  "level 5 (invalid)",
			level: 5,
			empty: true,
		},
		{
			name:  "negative level",
			level: -1,
			empty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LevelBadge(tt.level)
			if tt.empty {
				if result != "" {
					t.Errorf("expected empty string for level %d, got %q", tt.level, result)
				}
			} else {
				if !strings.Contains(result, tt.contains) {
					t.Errorf("expected badge to contain %q, got %q", tt.contains, result)
				}
			}
		})
	}
}

func TestColorConstants(t *testing.T) {
	// Verify color constants are defined (non-empty)
	colors := []struct {
		name  string
		color string
	}{
		{"PrimaryColor", string(PrimaryColor)},
		{"SecondaryColor", string(SecondaryColor)},
		{"WarningColor", string(WarningColor)},
		{"ErrorColor", string(ErrorColor)},
		{"SubtleColor", string(SubtleColor)},
		{"ActiveColor", string(ActiveColor)},
		{"HistoricalColor", string(HistoricalColor)},
		{"InProcessColor", string(InProcessColor)},
	}

	for _, c := range colors {
		t.Run(c.name, func(t *testing.T) {
			if c.color == "" {
				t.Errorf("%s should not be empty", c.name)
			}
		})
	}
}
