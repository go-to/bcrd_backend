CREATE TABLE IF NOT EXISTS bcrd.shops_location
(
    id         SERIAL PRIMARY KEY,
    shop_id    INT                                 NULL,
    latitude   DOUBLE PRECISION                    NOT NULL,
    longitude  DOUBLE PRECISION                    NOT NULL,
    location   GEOMETRY(POINT)                     NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE bcrd.shops_location IS '店舗位置情報';
COMMENT ON COLUMN bcrd.shops_location.id IS 'ID';
COMMENT ON COLUMN bcrd.shops_location.shop_id IS '店舗ID';
COMMENT ON COLUMN bcrd.shops_location.latitude IS '緯度';
COMMENT ON COLUMN bcrd.shops_location.longitude IS '経度';
COMMENT ON COLUMN bcrd.shops_location.location IS '位置情報';
