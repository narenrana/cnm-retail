package repository

import (
	"reflect"
	"shopping-cart/cnm-users/entities"
	e "shopping-cart/cnm-users/entities"
	"testing"
)

func TestUsersRepositoryInstance(t *testing.T) {
	tests := []struct {
		name string
		want UsersRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UsersRepositoryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersRepositoryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersRepository_Add(t *testing.T) {
	type args struct {
		user entities.UserDetails
	}
	tests := []struct {
		name    string
		args    args
		want    entities.UserDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersRepository{}
			got, err := us.Add(tt.args.user)
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

func Test_usersRepository_Delete(t *testing.T) {
	type args struct {
		user entities.UserDetails
	}
	tests := []struct {
		name    string
		args    args
		want    entities.UserDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersRepository{}
			got, err := us.Delete(tt.args.user)
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

func Test_usersRepository_FindBy(t *testing.T) {
	type args struct {
		user entities.UserDetails
	}
	tests := []struct {
		name    string
		args    args
		want    entities.UserDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersRepository{}
			got, err := us.FindBy(tt.args.user)
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

func Test_usersRepository_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []e.UserDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersRepository{}
			got, err := us.List()
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
