-- Create "products" table
CREATE TABLE "products" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(100) NULL,
  "price" bigint NULL,
  "stock" bigint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "chk_products_price" CHECK ((price >= 0) AND (price <= 1000000)),
  CONSTRAINT "chk_products_stock" CHECK ((stock >= 0) AND (stock <= 100))
);
-- Create index "idx_products_deleted_at" to table: "products"
CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "image" character varying(255) NULL,
  "role" character varying(20) NOT NULL DEFAULT 'USER',
  "verified_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create "verification_codes" table
CREATE TABLE "verification_codes" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "code" character varying(6) NOT NULL,
  "purpose" character varying(50) NOT NULL,
  "expired_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_verification_codes" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
