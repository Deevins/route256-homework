-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            BIGSERIAL PRIMARY KEY   NOT NULL,
    username      text                    NOT NULL,
    email         text      DEFAULT NULL UNIQUE,
    register_date TIMESTAMP DEFAULT CURRENT_DATE,
    created_at    TIMESTAMP DEFAULT now() NOT NULL,
    updated_at    TIMESTAMP DEFAULT NULL
);

CREATE TABLE orders
(
    id           BIGSERIAL PRIMARY KEY       NOT NULL,
    user_id      int references users (id),
    product_name text                        not null,
    quantity     int       default 0,
    status       text      DEFAULT 'created' NOT NULL,
    order_date   TIMESTAMP DEFAULT CURRENT_DATE,
    created_at   TIMESTAMP DEFAULT now()     NOT NULL,
    updated_at   TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
DROP TABLE users;
-- +goose StatementEnd
