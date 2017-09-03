# --- !Ups
CREATE SEQUENCE wallet_seq;
CREATE TABLE wallet (
    id bigint NOT NULL DEFAULT NEXTVAL('wallet_seq'),
    amount float8,
    updated timestamp,
    PRIMARY KEY (id)
);


# --- !Downs
DROP TABLE wallet;
DROP SEQUENCE wallet_seq;