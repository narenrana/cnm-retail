package services

import (
	"reflect"
	ce "shopping-cart/cnm-carts/entities"
	"testing"
)

func TestNewProductPriceService(t *testing.T) {
	tests := []struct {
		name string
		want ProductPriceService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductPriceService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductPriceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productPriceService_findCartItemProduct(t *testing.T) {
	type args struct {
		items []*ce.CartItems
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
			p := productPriceService{}
			got, err := p.findCartItemProduct(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("findCartItemProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCartItemProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
