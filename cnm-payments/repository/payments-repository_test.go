package payments

import (
	"reflect"
	"shopping-cart/cnm-payments/entities"
	"testing"
	e "shopping-cart/cnm-payments/entities"
)

func TestPaymentsRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaymentsRepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentsRepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Add(t *testing.T) {
	type args struct {
		item entities.Payments
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repository{}
			got, err := re.Add(tt.args.item)
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
		item entities.Payments
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repository{}
			got, err := re.Delete(tt.args.item)
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
		item entities.Payments
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repository{}
			got, err := re.FindBy(tt.args.item)
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
		want    []e.Payments
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repository{}
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
