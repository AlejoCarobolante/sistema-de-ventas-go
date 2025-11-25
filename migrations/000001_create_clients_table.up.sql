CREATE TABLE clients (
    id VARCHAR(191) not null,
    name longtext DEFAULT NULL,
    email longtext DEFAULT NULL,
    balance FLOAT DEFAULT NULL,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
    PRIMARY KEY (id)
);