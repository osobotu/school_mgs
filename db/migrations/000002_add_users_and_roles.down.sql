ALTER TABLE IF EXISTS "teachers" DROP CONSTRAINT IF EXISTS "teachers_user_id_fkey";
ALTER TABLE IF EXISTS "students" DROP CONSTRAINT IF EXISTS "students_user_id_fkey";
ALTER TABLE IF EXISTS "users" DROP CONSTRAINT IF EXISTS "users_role_id_fkey";

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;