package services

import (
	"reflect"
	"testing"
	ce "shopping-cart/cnm-carts/entities"
	oe "shopping-cart/cnm-offers/entities"
)

func TestNewOffersServiceService(t *testing.T) {
	tests := []struct {
		name string
		want OffersService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOffersServiceService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOffersServiceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_offersService_comboOfferDiscount(t *testing.T) {
	type args struct {
		items []*ce.CartItems
		offer *oe.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    *oe.Offers
		want1   float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &offersService{}
			got, got1, err := s.comboOfferDiscount(tt.args.items, tt.args.offer)
			if (err != nil) != tt.wantErr {
				t.Errorf("comboOfferDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("comboOfferDiscount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("comboOfferDiscount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_offersService_getOffers(t *testing.T) {
	type args struct {
		items []*ce.CartItems
	}
	tests := []struct {
		name    string
		args    args
		want    []*oe.Offers
		want1   float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &offersService{}
			got, got1, err := s.getOffers(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOffers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOffers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getOffers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_offersService_individualOfferDiscount(t *testing.T) {
	type args struct {
		items []*ce.CartItems
		offer *oe.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    *oe.Offers
		want1   float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &offersService{}
			got, got1, err := s.individualOfferDiscount(tt.args.items, tt.args.offer)
			if (err != nil) != tt.wantErr {
				t.Errorf("individualOfferDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("individualOfferDiscount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("individualOfferDiscount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_offersService_individualPossiblePairs(t *testing.T) {
	type args struct {
		productComboMap       map[string]int
		individualCombination map[string]int
		productAndQuantityMap map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &offersService{}
		})
	}
}

func Test_offersService_maxPossibleCombo(t *testing.T) {
	type args struct {
		individualCombination map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &offersService{}
			if got := s.maxPossibleCombo(tt.args.individualCombination); got != tt.want {
				t.Errorf("maxPossibleCombo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_offersService_productAndQuantityMap(t *testing.T) {
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
			//s := &offersService{}
		})
	}
}

func Test_offersService_productRuleMap(t *testing.T) {
	type args struct {
		offer   *oe.Offers
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
			//s := &offersService{}
		})
	}
}
