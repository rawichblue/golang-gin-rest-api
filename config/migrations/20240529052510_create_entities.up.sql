SET statement_timeout = 0;

--bun:split

SELECT 1

--bun:split

SELECT 2

--bun:split

CREATE TABLE "title_prefixs" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(255),
    "short_name" varchar(255),
    "description" varchar(255),
    "is_actived" boolean,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz
);
