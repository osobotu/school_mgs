CREATE TABLE "teachers" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "subject_id" integer,
  "department_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "teacher_teaches_class" (
  "teacher_id" integer NOT NULL,
  "class_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("teacher_id", "class_id")
);

CREATE TABLE "subjects" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "classes" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "form_masters" (
  "id" serial PRIMARY KEY,
  "teacher_id" integer UNIQUE,
  "class_id" integer,
  "arm_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  UNIQUE (class_id, arm_id)
);

CREATE TABLE "arms" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "class_has_arms" (
  "class_id" integer NOT NULL,
  "arm_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("class_id", "arm_id")
);

CREATE TABLE "departments" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "department_has_subjects" (
  "subject_id" integer NOT NULL,
  "department_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("subject_id", "department_id")
);

CREATE TABLE "students" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "middle_name" varchar,
  "class_id" integer,
  "department_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "student_offers_subject" (
  "student_id" integer NOT NULL,
  "subject_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("student_id", "subject_id")
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
  "assessment" float NOT NULL,
  "exam" float NOT NULL,
  "subject_id" integer NOT NULL,
  "term_id" integer NOT NULL,
  "session_id" integer NOT NULL,
  "class_id" integer NOT NULL,
  "arm_id" integer NOT NULL,
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

ALTER TABLE "teachers" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id") ON DELETE SET NULL;

ALTER TABLE "teachers" ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("id") ON DELETE SET NULL;

ALTER TABLE "teacher_teaches_class" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id") ON DELETE CASCADE;

ALTER TABLE "teacher_teaches_class" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE CASCADE;

ALTER TABLE "form_masters" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id") ON DELETE CASCADE;

ALTER TABLE "form_masters" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE CASCADE;

ALTER TABLE "class_has_arms" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE CASCADE;

ALTER TABLE "class_has_arms" ADD FOREIGN KEY ("arm_id") REFERENCES "arms" ("id") ON DELETE CASCADE;

ALTER TABLE "department_has_subjects" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id") ON DELETE CASCADE;

ALTER TABLE "department_has_subjects" ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("id") ON DELETE CASCADE;

ALTER TABLE "students" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE SET NULL;

ALTER TABLE "students" ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("id") ON DELETE SET NULL;

ALTER TABLE "student_offers_subject" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id") ON DELETE CASCADE;

ALTER TABLE "student_offers_subject" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id") ON DELETE CASCADE;

ALTER TABLE "term_scores" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id") ON DELETE CASCADE;

ALTER TABLE "term_scores" ADD FOREIGN KEY ("term_id") REFERENCES "terms" ("id") ON DELETE CASCADE;

ALTER TABLE "term_scores" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id") ON DELETE CASCADE;

ALTER TABLE "term_scores" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE CASCADE;

ALTER TABLE "term_scores" ADD FOREIGN KEY ("arm_id") REFERENCES "arms" ("id") ON DELETE CASCADE;

ALTER TABLE "scores" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id") ON DELETE CASCADE;

ALTER TABLE "scores" ADD FOREIGN KEY ("term_scores_id") REFERENCES "term_scores" ("id") ON DELETE CASCADE;
