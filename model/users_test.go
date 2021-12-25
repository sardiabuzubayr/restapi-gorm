package model

import (
	"fmt"
	"restapi/config"
	"testing"
)

func TestCreate(t *testing.T) {
	config.ConnectDB()
	user := new(Users)
	user.Email = "private.sardi@gmail.com"
	user.Nama = "Sardi"
	user.Alamat = "Jalan - "
	user.NoHandphone = "08xxxx"
	user.Ktp = "ktp.jpg"

	if user.CreateUser() != nil {
		t.Errorf("Gagal create user")
	}
	fmt.Println(user)
}

func TestUpdate(t *testing.T) {
	config.ConnectDB()
	user := new(Users)
	user.Email = "sardi@gmail.com"
	user.Nama = "Sardi Bastian"
	user.Alamat = "Jalan Sultan Sulaiman"
	user.NoHandphone = "021xxxx"
	user.Ktp = "ktps.jpg"

	if user.UpdateUser("private.sardi@gmail.com") != nil {
		t.Errorf("Gagal update user")
	}
	fmt.Println(user)
}

func TestDelete(t *testing.T) {
	config.ConnectDB()
	user, err := GetOneByEmail("sardi@gmail.com")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(user)
	if user.DeleteUser() != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	config.ConnectDB()
	// user, err := GetOneByEmail("private.sardi@gmail.com")

	// if err != nil {
	// 	t.Errorf("Error get one")
	// }

	// fmt.Println(user)
	users, err := GetAll("celik")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(users)
}
