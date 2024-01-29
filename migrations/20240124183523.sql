-- Modify "documents" table
ALTER TABLE "documents" DROP CONSTRAINT "fk_tenant_agreements_lease";
-- Modify "tenant_agreement_statuses" table
ALTER TABLE "tenant_agreement_statuses" DROP CONSTRAINT "fk_tenant_agreements_status";
-- Modify "tenant_agreements" table
ALTER TABLE "tenant_agreements" ADD COLUMN "lease_id" bigint NULL, ADD COLUMN "status_id" bigint NULL, ADD
 CONSTRAINT "fk_tenant_agreements_lease" FOREIGN KEY ("lease_id") REFERENCES "documents" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_tenant_agreements_status" FOREIGN KEY ("status_id") REFERENCES "tenant_agreement_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
