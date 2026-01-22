CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    first_name      TEXT NOT NULL CHECK (length(trim(first_name)) >= 1),
    middle_name     TEXT,
    last_name       TEXT NOT NULL CHECK (length(trim(last_name)) >= 1),
    date_of_birth   DATE NOT NULL,
    pesel           CHAR(11) UNIQUE NOT NULL CHECK (pesel ~ '^[0-9]{11}$'),
    is_accepted     BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE user_addresses (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    postcode        TEXT NOT NULL,
    street          TEXT NOT NULL,
    building_number TEXT NOT NULL,
    apartment_number TEXT,
    city            TEXT NOT NULL,
    country         TEXT NOT NULL DEFAULT 'Poland',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);
