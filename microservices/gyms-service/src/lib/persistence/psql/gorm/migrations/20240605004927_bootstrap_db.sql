-- Create "gym_owners" table
CREATE TABLE "gym_owners" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "phone_number" text NOT NULL,
  "email" text NOT NULL,
  "restricted" boolean NOT NULL,
  "created_by" text NOT NULL,
  "updated_by" text NOT NULL,
  "deleted_by" text NULL,
  "deleted_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_gym_owners_deleted_at" to table: "gym_owners"
CREATE INDEX "idx_gym_owners_deleted_at" ON "gym_owners" ("deleted_at");
-- Create "gyms" table
CREATE TABLE "gyms" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "address" text NOT NULL,
  "enabled" boolean NOT NULL,
  "disabled_for" text NULL,
  "owner_id" text NOT NULL,
  "created_by" text NOT NULL,
  "updated_by" text NOT NULL,
  "deleted_by" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_gym_owners_gyms" FOREIGN KEY ("owner_id") REFERENCES "gym_owners" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create index "idx_gyms_deleted_at" to table: "gyms"
CREATE INDEX "idx_gyms_deleted_at" ON "gyms" ("deleted_at");
