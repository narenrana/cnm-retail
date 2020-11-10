package repository

import (
	"reflect"
	"shopping-cart/cnm-auth/entities"
	"testing"
)

func TestNewUsersRepository(t *testing.T) {
	tests := []struct {
		name string
		want UsersAuthRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersAuthRepository_AddOrUpdate(t *testing.T) {
	type args struct {
		user entities.Users
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersAuthRepository{}
			got, err := us.AddOrUpdate(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOrUpdate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersAuthRepository_Delete(t *testing.T) {
	type args struct {
		user entities.Users
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersAuthRepository{}
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

func Test_usersAuthRepository_FindById(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &usersAuthRepository{}
			got, err := us.FindById(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
