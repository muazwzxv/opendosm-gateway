-- +goose Up
-- +goose StatementBegin
CREATE TABLE item_lookups (
    id BIGSERIAL PRIMARY KEY,
    item_code VARCHAR(255) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    item_group VARCHAR(255) NOT NULL,
    item_category VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(255) NOT NULL
);

CREATE TABLE premise_lookups (
    id BIGSERIAL PRIMARY KEY,
    premise_code BIGINT NOT NULL,
    premise_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    type VARCHAR(100) NOT NULL,
    district VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item_lookups;
DROP TABLE premise_lookups;
-- +goose StatementEnd
