package repository

import (
	"reflect"
	"testing"
)

func TestNewDatabaseManager(t *testing.T) {
	tests := []struct {
		name string
		want Product
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDatabaseManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_findBy(t *testing.T) {
	type fields struct {
		ProductId    uint
		ProductName  string
		ProductPrice uint
		BaseCurrency string
		ProductTitle string
		ProductDesc  *string
		ImageUrl     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &Product{
				ProductId:    tt.fields.ProductId,
				ProductName:  tt.fields.ProductName,
				ProductPrice: tt.fields.ProductPrice,
				BaseCurrency: tt.fields.BaseCurrency,
				ProductTitle: tt.fields.ProductTitle,
				ProductDesc:  tt.fields.ProductDesc,
				ImageUrl:     tt.fields.ImageUrl,
			}
			got, err := pr.findBy()
			if (err != nil) != tt.wantErr {
				t.Errorf("findBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
