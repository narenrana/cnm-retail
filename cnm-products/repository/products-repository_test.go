package repository

import (
	"reflect"
	"shopping-cart/cnm-products/entities"
	e "shopping-cart/cnm-products/entities"
	"testing"
)

func TestProductRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductRepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Add(t *testing.T) {
	type args struct {
		user entities.Product
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repository{}
			got, err := p.Add(tt.args.user)
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
		user entities.Product
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repository{}
			got, err := p.Delete(tt.args.user)
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
		user entities.Product
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repository{}
			got, err := p.FindBy(tt.args.user)
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

func Test_repository_FindByIds(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name    string
		args    args
		want    []e.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repository{}
			got, err := p.FindByIds(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []e.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repository{}
			got, err := p.List()
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
