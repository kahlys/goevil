package goevil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDependencies(t *testing.T) {
	tests := map[string]struct {
		fileIn  string
		want    []string
		wantErr bool
	}{
		"ok": {
			fileIn:  "./testdata/gomod.golden",
			want:    []string{"github.com/go-kit/kit", "github.com/pkg/errors", "golang.org/x/time"},
			wantErr: false,
		},
		"no-file": {
			fileIn:  "",
			want:    []string{},
			wantErr: true,
		},
		"corrupt": {
			fileIn:  "./testdata/corrupt.golden",
			want:    []string{},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getDependencies(tt.fileIn)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
