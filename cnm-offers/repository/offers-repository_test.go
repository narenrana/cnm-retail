package repository

import (
	"reflect"
	"shopping-cart/cnm-offers/entities"
	"testing"
	e "shopping-cart/cnm-offers/entities"
)

func TestOffersRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OffersRepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OffersRepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Add(t *testing.T) {
	type args struct {
		cart entities.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Offers
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

func Test_repository_Delete(t *testing.T) {
	type args struct {
		offer entities.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Offers
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
		user entities.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Offers
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

func Test_repository_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []*e.Offers
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
