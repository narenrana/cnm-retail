package payments

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_add(t *testing.T) {
	type args struct {
		request paymentsAddRequest
	}
	tests := []struct {
		name    string
		args    args
		want    paymentsAddResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.add(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_list(t *testing.T) {
	tests := []struct {
		name    string
		want    paymentsListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.list()
			if (err != nil) != tt.wantErr {
				t.Errorf("list() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("list() got = %v, want %v", got, tt.want)
			}
		})
	}
}
