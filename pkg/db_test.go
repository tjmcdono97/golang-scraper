package pkg

import (
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRepository()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_FetchData(t *testing.T) {
	tests := []struct {
		name    string
		r       *Repository
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FetchData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FetchData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FetchData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_FetchIDs(t *testing.T) {
	tests := []struct {
		name    string
		r       *Repository
		want    map[string]bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FetchIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FetchIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FetchIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_InsertAsset(t *testing.T) {
	type args struct {
		id  string
		url string
	}
	tests := []struct {
		name    string
		r       *Repository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.InsertAsset(tt.args.id, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Repository.InsertAsset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
