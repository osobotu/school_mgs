CREATE TABLE "teachers" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "subject_id" int NOT NULL,
  "classes" int[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "subjects" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "classes" int[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "classes" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "form_master_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "students" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "class_id" int NOT NULL,
  "subjects" int[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "scores" (
  "student_id" int NOT NULL,
  "subject_id" int NOT NULL,
  "first_term_assessment" int,
  "first_term_exam" int,
  "second_term_assessment" int,
  "second_term_exam" int,
  "third_term_assessment" int,
  "third_term_exam" int,
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
