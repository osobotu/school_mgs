CREATE TABLE "teachers" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "subject_id" integer NOT NULL,
  "classes" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subjects" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "classes" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "classes" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "form_master_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "students" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "class_id" integer[] NOT NULL,
  "subjects" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "terms" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "number" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "term_scores" (
  "id" serial PRIMARY KEY,
  "assessment" float,
  "exam" float,
  "subject_id" integer NOT NULL,
  "term_id" integer NOT NULL,
  "session_id" integer NOT NULL,
  "class_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" serial PRIMARY KEY,
  "session" varchar NOT NULL,
  "start_date" timestamptz,
  "end_date" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "scores" (
  "student_id" integer NOT NULL,
  "term_scores_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("student_id", "term_scores_id")
);

CREATE INDEX ON "teachers" ("first_name");

CREATE INDEX ON "teachers" ("last_name");

CREATE INDEX ON "subjects" ("name");

CREATE INDEX ON "students" ("first_name");

CREATE INDEX ON "students" ("last_name");

CREATE UNIQUE INDEX ON "scores" ("student_id", "term_scores_id");

COMMENT ON COLUMN "subjects"."classes" IS 'These are the classes that can take this subject';

COMMENT ON COLUMN "students"."class_id" IS 'A student can only belong to one class, ensure this array has a length of one';

ALTER TABLE "teachers" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("form_master_id") REFERENCES "teachers" ("id");

ALTER TABLE "term_scores" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "term_scores" ADD FOREIGN KEY ("term_id") REFERENCES "terms" ("id");

ALTER TABLE "term_scores" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id");

ALTER TABLE "term_scores" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("term_scores_id") REFERENCES "term_scores" ("id");
