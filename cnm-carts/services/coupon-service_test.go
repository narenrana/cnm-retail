package services

import (
	"reflect"
	ce "shopping-cart/cnm-carts/entities"
	"shopping-cart/cnm-coupons/entities"
	"testing"


	e "shopping-cart/cnm-coupons/entities"

)

func TestNewCouponServiceService(t *testing.T) {
	tests := []struct {
		name string
		want CouponService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCouponServiceService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCouponServiceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_couponService_couponDiscount(t *testing.T) {
	type args struct {
		cart ce.Cart
	}
	tests := []struct {
		name    string
		args    args
		want    *e.DiscountCoupons
		want1   float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &couponService{}
			got, got1, err := s.couponDiscount(tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("couponDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("couponDiscount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("couponDiscount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_couponService_findDiscount(t *testing.T) {
	type args struct {
		items          []*ce.CartItems
		discountCoupon entities.DiscountCoupons
	}
	tests := []struct {
		name    string
		args    args
		want    *e.DiscountCoupons
		want1   float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &couponService{}
			got, got1, err := s.findDiscount(tt.args.items, tt.args.discountCoupon)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiscount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findDiscount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_couponService_productDiscountRuleMap(t *testing.T) {
	type args struct {
		coupon  entities.DiscountCoupons
		itemMap map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &couponService{}
		})
	}
}

func Test_couponService_productsAndQuantityMap(t *testing.T) {
	type args struct {
		items                 []*ce.CartItems
		productAndQuantityMap map[string]int
		productPriceMapping   map[string]float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &couponService{}
		})
	}
}
