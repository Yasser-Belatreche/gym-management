-- Create "users" table
CREATE TABLE "users" (
  "id" text NOT NULL,
  "password" text NOT NULL,
  "role" text NOT NULL,
  "profile" jsonb NOT NULL,
  "restricted" boolean NOT NULL,
  "last_login" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "usernames" table
CREATE TABLE "usernames" (
  "user_id" text NOT NULL,
  "username" text NOT NULL,
  PRIMARY KEY ("user_id", "username"),
  CONSTRAINT "uni_usernames_username" UNIQUE ("username"),
  CONSTRAINT "fk_users_usernames" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
