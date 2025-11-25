CREATE TABLE compatibility_rules(
    id VARCHAR(191) NOT NULL,
    isCompatible BOOLEAN    DEFAULT NULL,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
)