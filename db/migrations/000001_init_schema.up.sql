CREATE TABLE "teachers" (
  "id" integer PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "subject_id" integer NOT NULL,
  "classes" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subjects" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "classes" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "classes" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "form_master_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "students" (
  "id" integer PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "class_id" integer NOT NULL,
  "subjects" integer[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "scores" (
  "student_id" integer NOT NULL,
  "subject_id" integer NOT NULL,
  "first_term_assessment" integer,
  "first_term_exam" integer,
  "second_term_assessment" integer,
  "second_term_exam" integer,
  "third_term_assessment" integer,
  "third_term_exam" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY (student_id, subject_id)
);

CREATE INDEX ON "teachers" ("first_name");

CREATE INDEX ON "teachers" ("last_name");

CREATE INDEX ON "subjects" ("name");

CREATE INDEX ON "students" ("first_name");

CREATE INDEX ON "students" ("last_name");

CREATE INDEX ON "students" ("class_id");

CREATE UNIQUE INDEX ON "scores" ("student_id", "subject_id");

COMMENT ON COLUMN "subjects"."classes" IS 'These are the classes that can take this subject';

ALTER TABLE "teachers" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("form_master_id") REFERENCES "teachers" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");
