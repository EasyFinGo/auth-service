CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    first_name      TEXT NOT NULL CHECK (length(trim(first_name)) >= 1),
    middle_name     TEXT,
    last_name       TEXT NOT NULL CHECK (length(trim(last_name)) >= 1),
    date_of_birth   DATE NOT NULL,
    pesel           CHAR(11) UNIQUE NOT NULL CHECK (pesel ~ '^[0-9]{11}$'),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE user_addresses (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_current      BOOLEAN NOT NULL DEFAULT true,
    postcode        TEXT NOT NULL,
    street          TEXT NOT NULL,
    building_number TEXT NOT NULL,
    apartment_number TEXT,
    city            TEXT NOT NULL,
    country         TEXT NOT NULL DEFAULT 'Poland',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE user_consents (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    consent_type    TEXT NOT NULL CHECK (consent_type IN ('privacy_policy', 'terms_of_use')),
    version         TEXT NOT NULL,
    accepted_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    ip_address      INET,
    UNIQUE (user_id, consent_type)
);

CREATE INDEX idx_users_pesel ON users(pesel);
CREATE INDEX idx_addresses_user ON user_addresses(user_id);