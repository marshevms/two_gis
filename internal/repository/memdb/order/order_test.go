package order

import (
	"testing"
	"time"
)

func TestIsTimeRangesIntersect(t *testing.T) {
	cases := []struct {
		timeRange1     [2]time.Time
		timeRange2     [2]time.Time
		expectedResult bool
	}{
		{
			timeRange1: [2]time.Time{
				time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
			},
			timeRange2: [2]time.Time{
				time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			expectedResult: false,
		},
		{
			timeRange1: [2]time.Time{
				time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			timeRange2: [2]time.Time{
				time.Date(2023, time.January, 1, 11, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 13, 0, 0, 0, time.UTC),
			},
			expectedResult: true,
		},
		{
			timeRange1: [2]time.Time{
				time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			timeRange2: [2]time.Time{
				time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			expectedResult: true,
		},
		{
			timeRange1: [2]time.Time{
				time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			timeRange2: [2]time.Time{
				time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
				time.Date(2023, time.January, 1, 20, 0, 0, 0, time.UTC),
			},
			expectedResult: true,
		},
	}

	for _, c := range cases {
		result := isTimeRangesIntersect(c.timeRange1, c.timeRange2)
		if result != c.expectedResult {
			t.Errorf("Expected %v, got %v", c.expectedResult, result)
		}
	}
}
