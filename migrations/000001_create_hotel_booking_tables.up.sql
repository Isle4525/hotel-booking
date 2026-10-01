CREATE TABLE IF NOT EXISTS categories (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR(100) NOT NULL UNIQUE,
    slug        VARCHAR(120) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS hotels (

    id              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_id     BIGINT REFERENCES categories(id) ON DELETE SET NULL,
    name            VARCHAR(200) NOT NULL,
    slug            VARCHAR(220) NOT NULL UNIQUE,
    city            VARCHAR(120) NOT NULL,
    address         TEXT,
    description     TEXT,
    image_url       TEXT,
    starting_price  NUMERIC(12, 2),
    currency        CHAR(3) NOT NULL DEFAULT 'KZT',
    rating          NUMERIC(2, 1),
    reviews_count   INTEGER NOT NULL DEFAULT 0,
    badge           VARCHAR(80),
    is_published    BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE  IF NOT EXISTS room_types (
    id              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    hotel_id        BIGINT NOT NULL REFERENCES hotels(id) ON DELETE CASCADE,
    name            VARCHAR(160) NOT NULL,
    description     TEXT,
    capacity        SMALLINT NOT NULL DEFAULT 2,
    price_per_night NUMERIC(12, 2) NOT NULL,
    currency        CHAR(3) NOT NULL DEFAULT 'KZT',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (capacity > 0),
    CHECK (price_per_night >= 0)
);

CREATE TABLE  IF NOT EXISTS amenities (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR(100) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE  IF NOT EXISTS hotel_amenities (
    hotel_id    BIGINT NOT NULL REFERENCES hotels(id) ON DELETE CASCADE,
    amenity_id  BIGINT NOT NULL REFERENCES amenities(id) ON DELETE CASCADE,
    PRIMARY KEY (hotel_id, amenity_id)
);
