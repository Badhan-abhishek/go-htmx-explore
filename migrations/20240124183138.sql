-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "first_name" text NULL,
  "last_name" text NULL,
  "email" text NULL,
  "password" bytea NULL,
  "phone_number" text NULL,
  "profile_pic" text NULL,
  "active" boolean NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_number_key" ON "users" ("phone_number");
-- Create "landlords" table
CREATE TABLE "landlords" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_landlords_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_landlords_deleted_at" to table: "landlords"
CREATE INDEX "idx_landlords_deleted_at" ON "landlords" ("deleted_at");
-- Create "properties" table
CREATE TABLE "properties" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "label" text NULL,
  "landlord_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_landlords_properties" FOREIGN KEY ("landlord_id") REFERENCES "landlords" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_properties_deleted_at" to table: "properties"
CREATE INDEX "idx_properties_deleted_at" ON "properties" ("deleted_at");
-- Create "addresses" table
CREATE TABLE "addresses" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "address1" text NULL,
  "address2" text NULL,
  "city" text NULL,
  "country" text NULL,
  "postal_code" text NULL,
  "state" text NULL,
  "property_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_properties_address" FOREIGN KEY ("property_id") REFERENCES "properties" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_addresses_deleted_at" to table: "addresses"
CREATE INDEX "idx_addresses_deleted_at" ON "addresses" ("deleted_at");
-- Create "document_types" table
CREATE TABLE "document_types" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "sequence" bigint NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_document_types_deleted_at" to table: "document_types"
CREATE INDEX "idx_document_types_deleted_at" ON "document_types" ("deleted_at");
-- Create "tenant_statuses" table
CREATE TABLE "tenant_statuses" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "sequence" bigint NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_tenant_statuses_deleted_at" to table: "tenant_statuses"
CREATE INDEX "idx_tenant_statuses_deleted_at" ON "tenant_statuses" ("deleted_at");
-- Create "tenants" table
CREATE TABLE "tenants" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NULL,
  "tenant_status_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tenants_tenant_status" FOREIGN KEY ("tenant_status_id") REFERENCES "tenant_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tenants_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_tenants_deleted_at" to table: "tenants"
CREATE INDEX "idx_tenants_deleted_at" ON "tenants" ("deleted_at");
-- Create "tenant_agreements" table
CREATE TABLE "tenant_agreements" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "label" text NULL,
  "start_date" timestamptz NULL,
  "end_date" timestamptz NULL,
  "rent" numeric NULL,
  "total_rent_amount" numeric NULL,
  "due_date" text NULL,
  "tenant_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tenants_tenant_agreement" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_tenant_agreements_deleted_at" to table: "tenant_agreements"
CREATE INDEX "idx_tenant_agreements_deleted_at" ON "tenant_agreements" ("deleted_at");
-- Create "documents" table
CREATE TABLE "documents" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "url" text NULL,
  "type_id" bigint NULL,
  "name" text NULL,
  "external_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_documents_type" FOREIGN KEY ("type_id") REFERENCES "document_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tenant_agreements_lease" FOREIGN KEY ("id") REFERENCES "tenant_agreements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_documents_deleted_at" to table: "documents"
CREATE INDEX "idx_documents_deleted_at" ON "documents" ("deleted_at");
-- Create "tenant_agreement_other_document" table
CREATE TABLE "tenant_agreement_other_document" (
  "tenant_agreement_id" bigint NOT NULL,
  "document_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_agreement_id", "document_id"),
  CONSTRAINT "fk_tenant_agreement_other_document_document" FOREIGN KEY ("document_id") REFERENCES "documents" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_tenant_agreement_other_document_tenant_agreement" FOREIGN KEY ("tenant_agreement_id") REFERENCES "tenant_agreements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "tenant_agreement_statuses" table
CREATE TABLE "tenant_agreement_statuses" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "sequence" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tenant_agreements_status" FOREIGN KEY ("id") REFERENCES "tenant_agreements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_tenant_agreement_statuses_deleted_at" to table: "tenant_agreement_statuses"
CREATE INDEX "idx_tenant_agreement_statuses_deleted_at" ON "tenant_agreement_statuses" ("deleted_at");
-- Create "units" table
CREATE TABLE "units" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "property_id" bigint NULL,
  "max_capacity" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_properties_units" FOREIGN KEY ("property_id") REFERENCES "properties" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_units_deleted_at" to table: "units"
CREATE INDEX "idx_units_deleted_at" ON "units" ("deleted_at");
-- Create "unit_image" table
CREATE TABLE "unit_image" (
  "unit_id" bigint NOT NULL,
  "document_id" bigint NOT NULL,
  PRIMARY KEY ("unit_id", "document_id"),
  CONSTRAINT "fk_unit_image_document" FOREIGN KEY ("document_id") REFERENCES "documents" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_unit_image_unit" FOREIGN KEY ("unit_id") REFERENCES "units" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "roles" table
CREATE TABLE "roles" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "sequence" bigint NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "roles" ("deleted_at");
-- Create "user_roles" table
CREATE TABLE "user_roles" (
  "user_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
