package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"com.backend/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&models.Tenant{},
		&models.TenantAgreement{},
		&models.TenantAgreementStatus{},
		&models.Unit{},
		&models.User{},
		&models.Landlord{},
		&models.Document{},
		&models.DocumentType{},
		&models.Property{},
		&models.Address{},
		&models.Role{},
		&models.TenantStatus{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
