package service

import (
	"fmt"
	"strings"
)

func ValidateRegister(username string, password string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return fmt.Errorf("username wajib diisi")
	}

	if len(username) < 3 {
		return fmt.Errorf("username minimal 3 karakter")
	}

	if len(username) > 50 {
		return fmt.Errorf("username maksimal 50 karakter")
	}

	if password == "" {
		return fmt.Errorf("password wajib diisi")
	}

	if len(password) < 6 {
		return fmt.Errorf("password minimal 6 karakter")
	}

	return nil
}

func ValidateLogin(username string, password string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return fmt.Errorf("username wajib diisi")
	}

	if password == "" {
		return fmt.Errorf("password wajib diisi")
	}

	return nil
}

func ValidateRole(role string) error {
	role = strings.TrimSpace(role)

	switch role {
	case "admin", "staff", "customer":
		return nil
	default:
		return fmt.Errorf("role tidak valid")
	}
}