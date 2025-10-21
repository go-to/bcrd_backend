CREATE TABLE IF NOT EXISTS bcrd.shops_image
(
    id         SERIAL PRIMARY KEY,
    shop_id    INT                                 NOT NULL,
    image_url  VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE bcrd.shops_image IS '店舗画像';
COMMENT ON COLUMN bcrd.shops_image.id IS 'ID';
COMMENT ON COLUMN bcrd.shops_image.shop_id IS '店舗ID';
COMMENT ON COLUMN bcrd.shops_image.image_url IS '画像URL';
