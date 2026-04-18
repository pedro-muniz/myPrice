CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    barcode VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    discount DOUBLE PRECISION NOT NULL,
    price_id UUID NOT NULL,
    profit_percent DOUBLE PRECISION NOT NULL,
    origin SMALLINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_price FOREIGN KEY (price_id) REFERENCES prices(id) ON DELETE CASCADE
);

CREATE INDEX idx_products_barcode ON products(barcode);
