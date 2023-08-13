package cmd

import (
	"testing"
)

func TestIncrementVersion(t *testing.T) {
	testCases := []struct {
		name            string
		currentVersion  string
		incrementValue  int
		versionType     string
		expectedVersion string
		expectedError   bool
	}{
		{
			name:            "major delta",
			currentVersion:  "1.0.0",
			incrementValue:  1,
			versionType:     "major",
			expectedVersion: "2.0.0",
		},
		{
			name:            "minor delta",
			currentVersion:  "1.1.0",
			incrementValue:  1,
			versionType:     "minor",
			expectedVersion: "1.2.0",
		},
		{
			name:            "patch delta",
			currentVersion:  "1.0.1",
			incrementValue:  1,
			versionType:     "patch",
			expectedVersion: "1.0.2",
		},
		{
			name:            "major 2 delta",
			currentVersion:  "1.2.0",
			incrementValue:  2,
			versionType:     "major",
			expectedVersion: "3.0.0",
		},
		{
			name:            "minor 2 delta",
			currentVersion:  "1.2.0",
			incrementValue:  2,
			versionType:     "minor",
			expectedVersion: "1.4.0",
		},
		{
			name:            "patch 2 delta",
			currentVersion:  "1.2.0",
			incrementValue:  2,
			versionType:     "patch",
			expectedVersion: "1.2.2",
		},
		{
			name:            "invalid version type",
			currentVersion:  "1.2",
			incrementValue:  2,
			versionType:     "dummy",
			expectedVersion: "",
			expectedError:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := incrementVersion(tc.currentVersion, tc.incrementValue, tc.versionType)
			if got != tc.expectedVersion {
				t.Errorf("expected %v: got %v", tc.expectedVersion, got)
			}
			if (err != nil) != tc.expectedError {
				if tc.expectedError {
					t.Errorf("expected an error but got no error")
				} else {
					t.Errorf("expected no error but got an error: %v", err)
				}
			}
		})
	}
}
