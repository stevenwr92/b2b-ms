-- Create "admins" table
CREATE TABLE "public"."admins" (
  "id" bigserial NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "admins_email_key" to table: "admins"
CREATE UNIQUE INDEX "admins_email_key" ON "public"."admins" ("email");
-- Create index "admins_name_key" to table: "admins"
CREATE UNIQUE INDEX "admins_name_key" ON "public"."admins" ("name");
-- Create index "admins_password_key" to table: "admins"
CREATE UNIQUE INDEX "admins_password_key" ON "public"."admins" ("password");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
-- Create index "users_password_key" to table: "users"
CREATE UNIQUE INDEX "users_password_key" ON "public"."users" ("password");
