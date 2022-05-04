CREATE SCHEMA IF NOT EXISTS account;
SET client_min_messages = error;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA pg_catalog VERSION "1.1";

-- auto update timestamp
CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.date_modified = now();
    RETURN NEW;
END;
$$ language 'plpgsql';


DROP TABLE IF EXISTS account.user;
CREATE TABLE account.user (
    user_id         bigint not null ,
	username        varchar(32) not null  ,
	password        varchar(128) not null  ,
    date_added               timestamptz DEFAULT current_timestamp,
    date_modified            timestamptz ,
    CONSTRAINT pk_account PRIMARY KEY ( user_id ),
    UNIQUE (username)
);
DROP INDEX IF EXISTS idx_user_0;
CREATE INDEX idx_user_0  ON account.user ( username );
DROP INDEX IF EXISTS idx_user_1;
CREATE INDEX idx_user_1  ON account.user ( password );

DROP TRIGGER IF EXISTS update_user ON account.user;
CREATE TRIGGER update_user
    BEFORE UPDATE ON account.user
    FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();


CREATE TABLE account.apikey (
                              user_id         bigint  NOT NULL  ,
                              api_key         bigint   NOT NULL ,
                              date_added               timestamptz DEFAULT current_timestamp,
                              date_modified            timestamptz ,
                              CONSTRAINT pk_apikey PRIMARY KEY ( user_id )
);

CREATE INDEX idx_apikey_0  ON account.apikey ( api_key );
CREATE TRIGGER update_apikey
    BEFORE UPDATE ON account.apikey
    FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();




CREATE TABLE account.transactions (
                                user_id         bigint  NOT NULL  ,
                                action          varchar(32)   ,
                                date_added               timestamptz DEFAULT current_timestamp,
                                date_modified            timestamptz ,
                                CONSTRAINT pk_transaction PRIMARY KEY ( user_id )
);

CREATE INDEX idx_transaction_0  ON account.transactions ( action );
CREATE TRIGGER update_transactions
    BEFORE UPDATE ON account.transactions
    FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
