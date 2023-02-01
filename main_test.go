package main

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_randomUserAgent(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomUserAgent(); got != tt.want {
				t.Errorf("randomUserAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildGoogleUrls(t *testing.T) {
	type args struct {
		searchTerm   string
		countryCode  string
		languageCode string
		pages        int
		count        int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildGoogleUrls(tt.args.searchTerm, tt.args.countryCode, tt.args.languageCode, tt.args.pages, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildGoogleUrls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildGoogleUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_googleResultParsing(t *testing.T) {
	type args struct {
		response *http.Response
		rank     int
	}
	tests := []struct {
		name    string
		args    args
		want    []SearchResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := googleResultParsing(tt.args.response, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("googleResultParsing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("googleResultParsing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScrapeClient(t *testing.T) {
	type args struct {
		proxyString interface{}
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScrapeClient(tt.args.proxyString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getScrapeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleScrape(t *testing.T) {
	type args struct {
		searchTerm   string
		countryCode  string
		languageCode string
		proxyString  interface{}
		pages        int
		count        int
		backoff      int
	}
	tests := []struct {
		name    string
		args    args
		want    []SearchResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GoogleScrape(tt.args.searchTerm, tt.args.countryCode, tt.args.languageCode, tt.args.proxyString, tt.args.pages, tt.args.count, tt.args.backoff)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleScrape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleScrape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scrapeClientRequest(t *testing.T) {
	type args struct {
		searchURL   string
		proxyString interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scrapeClientRequest(tt.args.searchURL, tt.args.proxyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("scrapeClientRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrapeClientRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
