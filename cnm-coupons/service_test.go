package coupons

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
		request discountCouponsAddRequest
	}
	tests := []struct {
		name    string
		args    args
		want    discountCouponsResponse
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

func Test_service_find(t *testing.T) {
	type args struct {
		request findCouponRequest
	}
	tests := []struct {
		name    string
		args    args
		want    findCouponResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.find(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_list(t *testing.T) {
	tests := []struct {
		name    string
		want    discountCouponsListResponse
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
