package services

import (
	"reflect"
	"shopping-cart/cnm-carts/entities"
	re "shopping-cart/cnm-carts/entities"
	"shopping-cart/cnm-carts/models"
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

func Test_service_Add(t *testing.T) {
	type args struct {
		request models.AddToCartRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.CartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.Add(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteCartItems(t *testing.T) {
	type args struct {
		request models.DeleteCartItemRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.CartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.DeleteCartItems(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCartItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteCartItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	type args struct {
		request models.GetCartRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.CartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_calculateTotalAmount(t *testing.T) {
	type args struct {
		items []*re.CartItems
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			if got := s.calculateTotalAmount(tt.args.items); got != tt.want {
				t.Errorf("calculateTotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_parePareCartResponse(t *testing.T) {
	type args struct {
		cart entities.Cart
		err  error
	}
	tests := []struct {
		name    string
		args    args
		want    models.CartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.parePareCartResponse(tt.args.cart, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("parePareCartResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parePareCartResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}