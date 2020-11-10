package repository

import (
	"reflect"
	"shopping-cart/cnm-carts/entities"
	"testing"
	ce "shopping-cart/cnm-carts/entities"
)

func TestCartsRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CartsRepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartsRepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Add(t *testing.T) {
	type args struct {
		cart entities.Cart
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.Add(tt.args.cart)
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

func Test_repository_DeleteCartItem(t *testing.T) {
	type args struct {
		cartItemIds []int
	}
	tests := []struct {
		name    string
		args    args
		want    []*ce.CartItems
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.DeleteCartItem(tt.args.cartItemIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCartItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteCartItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_FirstOrCreate(t *testing.T) {
	type args struct {
		userId *int
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.FirstOrCreate(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []*ce.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
