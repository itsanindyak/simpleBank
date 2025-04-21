CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "accounts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_accounts_id" bigint NOT NULL,
  "to_accounts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("accounts_id");

CREATE INDEX ON "transfers" ("from_accounts_id");

CREATE INDEX ON "transfers" ("to_accounts_id");

CREATE INDEX ON "transfers" ("from_accounts_id", "to_accounts_id");

COMMENT ON COLUMN "entries"."amount" IS 'Can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'Must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("accounts_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_accounts_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_accounts_id") REFERENCES "accounts" ("id");