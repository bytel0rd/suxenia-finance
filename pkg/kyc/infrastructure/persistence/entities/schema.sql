CREATE TABLE banking_kyc (

    id varchar primary key,
    name varchar(144) NOT NULL,

    bank_account_name varchar(144),
    bank_account_number varchar(10),
    bvn varchar(11),
    bank_code varchar,
    
    owner_id varchar NOT NULL,
    verified boolean NOT NULL,
    
    created_by varchar(144),
    updated_by varchar(144),

    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp

);
