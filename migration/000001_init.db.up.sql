CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "telegram_id" bigint UNIQUE,
    "name" VARCHAR (200) NULL

);
CREATE TABLE "tasks"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    user_id integer REFERENCES users ON DELETE CASCADE,
    "text" VARCHAR(255)
);