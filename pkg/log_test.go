package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestOpenLogFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenLogFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenLogFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenLogFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
