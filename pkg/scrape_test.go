package pkg

import (
	"reflect"
	"testing"
)

func TestScrapeURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScrapeURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScrapeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrapeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostListings(t *testing.T) {
	type args struct {
		search    string
		assetList map[string]bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostListings(tt.args.search, tt.args.assetList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostListings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomSleep(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randomSleep()
		})
	}
}

func TestIsNewAsset(t *testing.T) {
	type args struct {
		id        string
		assetList map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNewAsset(tt.args.id, tt.args.assetList); got != tt.want {
				t.Errorf("IsNewAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}
