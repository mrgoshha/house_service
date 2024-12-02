CREATE TABLE IF NOT EXISTS houses(
    house_number SERIAL PRIMARY KEY,
    address VARCHAR NOT NULL,
    year_construction INTEGER NOT NULL,
    developer VARCHAR,
    created_at TIMESTAMP NOT NULL,
    last_update TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS flats(
    flat_number SERIAL NOT NULL,
    house_number INTEGER REFERENCES houses ON DELETE CASCADE,
    price INTEGER NOT NULL,
    number_of_rooms INTEGER NOT NULL,
    status VARCHAR NOT NULL,
    PRIMARY KEY (flat_number, house_number)
);


CREATE TABLE IF NOT EXISTS users(
    id VARCHAR PRIMARY KEY,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    user_type VARCHAR NOT NULL
)