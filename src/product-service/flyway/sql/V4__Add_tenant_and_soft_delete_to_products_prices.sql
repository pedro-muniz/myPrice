ALTER TABLE prices ADD COLUMN company_id UUID;
ALTER TABLE prices ADD COLUMN branch_id UUID;
ALTER TABLE prices ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

ALTER TABLE products ADD COLUMN company_id UUID;
ALTER TABLE products ADD COLUMN branch_id UUID;
ALTER TABLE products ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;

-- Add foreign keys (optional but recommended for data integrity in local cache)
ALTER TABLE prices ADD CONSTRAINT fk_prices_company FOREIGN KEY (company_id) REFERENCES companies(id);
ALTER TABLE prices ADD CONSTRAINT fk_prices_branch FOREIGN KEY (branch_id) REFERENCES branches(id);

ALTER TABLE products ADD CONSTRAINT fk_products_company FOREIGN KEY (company_id) REFERENCES companies(id);
ALTER TABLE products ADD CONSTRAINT fk_products_branch FOREIGN KEY (branch_id) REFERENCES branches(id);

CREATE INDEX idx_prices_tenant ON prices(company_id, branch_id);
CREATE INDEX idx_products_tenant ON products(company_id, branch_id);
