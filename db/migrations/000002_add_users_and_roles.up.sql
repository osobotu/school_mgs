CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "role_id" integer NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "role" varchar UNIQUE NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "teachers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");