package coupons

import (
	"reflect"
	"shopping-cart/cnm-coupons/entities"
	"testing"
	e "shopping-cart/cnm-coupons/entities"
)

func TestRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Add(t *testing.T) {
	type args struct {
		d entities.DiscountCoupons
	}
	tests := []struct {
		name    string
		args    args
		want    entities.DiscountCoupons
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.Add(tt.args.d)
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

func Test_repository_Delete(t *testing.T) {
	type args struct {
		offer entities.DiscountCoupons
	}
	tests := []struct {
		name    string
		args    args
		want    entities.DiscountCoupons
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.Delete(tt.args.offer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_FindBy(t *testing.T) {
	type args struct {
		user entities.DiscountCoupons
	}
	tests := []struct {
		name    string
		args    args
		want    entities.DiscountCoupons
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.FindBy(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_FindByDiscountCoupon(t *testing.T) {
	type args struct {
		coupon string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.DiscountCoupons
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &repository{}
			got, err := re.FindByDiscountCoupon(tt.args.coupon)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByDiscountCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByDiscountCoupon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []e.DiscountCoupons
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
