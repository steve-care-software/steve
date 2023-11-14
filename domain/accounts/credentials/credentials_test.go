package credentials

import (
	"bytes"
	"testing"
)

func TestCreate_withUsername_withPassword_Success(t *testing.T) {
	username := "myUsername"
	password := []byte("myPassword")
	ins, err := NewBuilder().Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}

	retUsername := ins.Username()
	if username != retUsername {
		t.Errorf("the username was expected to be '%s', '%s' returned", username, retUsername)
		return
	}

	retPassword := ins.Password()
	if !bytes.Equal(password, retPassword) {
		t.Errorf("the password was expected to be '%s', '%s' returned", password, retPassword)
		return
	}
}

func TestCreate_withoutUsername_ReturnsError(t *testing.T) {
	password := []byte("myPassword")
	_, err := NewBuilder().Create().
		WithPassword(password).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCreate_withUsername_withEmptyPassword_ReturnsError(t *testing.T) {
	username := "myUsername"
	password := []byte("")
	_, err := NewBuilder().Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCreate_withUsername_withoutPassword_ReturnsError(t *testing.T) {
	username := "myUsername"
	_, err := NewBuilder().Create().
		WithUsername(username).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
