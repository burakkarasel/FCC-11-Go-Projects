CREATE TABLE "leads" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "company_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" bigint NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE INDEX ON "leads" ("id");