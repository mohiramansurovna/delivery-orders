-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id      BIGINT       NOT NULL,
    currency         CHAR(3)      NOT NULL,
    items            JSONB        NOT NULL DEFAULT '[]'::jsonb,
    status           TEXT         NOT NULL,
    total_amount     BIGINT       NOT NULL CHECK (total_amount >= 0),
    delivery_address TEXT         NOT NULL,
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ  NOT NULL DEFAULT now(),
)
    CONSTRAINT orders_items_is_array CHECK (jsonb_typeof(items) = 'array'),
    CONSTRAINT orders_status_valid CHECK (status IN
    ('pending', 'confirmed', 'preparing', 'ready', 'delivering', 'delivered', 'cancelled'))

CREATE INDEX idx_orders_customer_id ON orders (customer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd