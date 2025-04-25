package validators

import "testing"

func TestPage(t *testing.T) {
	testCases := []struct {
		name    string
		pageN   int32
		wantErr bool
	}{
		{
			name:    "negative",
			pageN:   -1,
			wantErr: true,
		},
		{
			name:    "zero",
			pageN:   0,
			wantErr: true,
		},
		{
			name:    "correct",
			pageN:   10,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				err := Page(tc.pageN)
				if (err != nil) != tc.wantErr {
					t.Errorf("Page() error = %v, wantErr %v", err, tc.wantErr)
				}
			},
		)
	}
}
