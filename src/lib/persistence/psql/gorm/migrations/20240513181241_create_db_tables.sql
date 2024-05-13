-- Create "users" table
CREATE TABLE "users" (
  "id" text NOT NULL,
  "password" text NOT NULL,
  "restricted" boolean NOT NULL,
  "last_login" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "customers" table
CREATE TABLE "customers" (
  "id" text NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "phone_number" text NOT NULL,
  "email" text NOT NULL,
  "birth_year" bigint NOT NULL,
  "gender" text NOT NULL,
  "restricted" boolean NOT NULL,
  "created_by" text NOT NULL,
  "updated_by" text NOT NULL,
  "deleted_by" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_customers_created_by_user" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_customers_deleted_by_user" FOREIGN KEY ("deleted_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_customers_updated_by_user" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create index "idx_customers_deleted_at" to table: "customers"
CREATE INDEX "idx_customers_deleted_at" ON "customers" ("deleted_at");
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
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_gym_owners_created_by_user" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_gym_owners_deleted_by_user" FOREIGN KEY ("deleted_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_gym_owners_updated_by_user" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
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
  CONSTRAINT "fk_gyms_created_by_user" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_gyms_deleted_by_user" FOREIGN KEY ("deleted_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_gyms_owner" FOREIGN KEY ("owner_id") REFERENCES "gym_owners" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_gyms_updated_by_user" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create index "idx_gyms_deleted_at" to table: "gyms"
CREATE INDEX "idx_gyms_deleted_at" ON "gyms" ("deleted_at");
-- Create "plans" table
CREATE TABLE "plans" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "featured" boolean NOT NULL,
  "sessions_per_week" bigint NOT NULL,
  "with_coach" boolean NOT NULL,
  "monthly_price" numeric NOT NULL,
  "gym_id" text NOT NULL,
  "created_by" text NOT NULL,
  "updated_by" text NOT NULL,
  "deleted_by" text NULL,
  "deleted_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_plans_created_by_user" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_plans_deleted_by_user" FOREIGN KEY ("deleted_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_plans_gym" FOREIGN KEY ("gym_id") REFERENCES "gyms" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_plans_updated_by_user" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create index "idx_plans_deleted_at" to table: "plans"
CREATE INDEX "idx_plans_deleted_at" ON "plans" ("deleted_at");
-- Create "memberships" table
CREATE TABLE "memberships" (
  "id" text NOT NULL,
  "code" text NOT NULL,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NULL,
  "enabled" boolean NOT NULL,
  "disabled_for" text NULL,
  "sessions_per_week" bigint NOT NULL,
  "with_coach" boolean NOT NULL,
  "monthly_price" numeric NOT NULL,
  "plan_id" text NOT NULL,
  "customer_id" text NOT NULL,
  "renewed_at" timestamptz NULL,
  "created_by" text NOT NULL,
  "updated_by" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_memberships_code" UNIQUE ("code"),
  CONSTRAINT "fk_memberships_created_by_user" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_memberships_customer" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_memberships_plan" FOREIGN KEY ("plan_id") REFERENCES "plans" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "fk_memberships_updated_by_user" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create "bills" table
CREATE TABLE "bills" (
  "id" text NOT NULL,
  "amount" numeric NOT NULL,
  "paid" boolean NOT NULL,
  "paid_at" timestamptz NULL,
  "due_to" timestamptz NOT NULL,
  "membership_id" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_bills_membership" FOREIGN KEY ("membership_id") REFERENCES "memberships" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create "training_sessions" table
CREATE TABLE "training_sessions" (
  "id" text NOT NULL,
  "membership_id" text NOT NULL,
  "started_at" timestamptz NOT NULL,
  "ended_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_training_sessions_membership" FOREIGN KEY ("membership_id") REFERENCES "memberships" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create "usernames" table
CREATE TABLE "usernames" (
  "user_id" text NOT NULL,
  "username" text NOT NULL,
  PRIMARY KEY ("user_id", "username"),
  CONSTRAINT "uni_usernames_username" UNIQUE ("username"),
  CONSTRAINT "fk_users_usernames" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
