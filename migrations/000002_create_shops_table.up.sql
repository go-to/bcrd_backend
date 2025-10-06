CREATE TABLE IF NOT EXISTS bcrd.shops
(
    id                           SERIAL PRIMARY KEY,
    event_id                     INT                                 NOT NULL,
    no                           VARCHAR(255)                        NOT NULL,
    shop_name                    VARCHAR(255)                        NOT NULL,
    image_url                    TEXT                                NOT NULL,
    google_url                   TEXT                                NOT NULL,
    tabelog_url                  TEXT                                NOT NULL,
    official_url                 TEXT                                NOT NULL,
    instagram_url                TEXT                                NOT NULL,
    address                      TEXT                                NOT NULL,
    business_days                TEXT                                NOT NULL,
    regular_holiday              VARCHAR(255)                        NOT NULL,
    is_open_holiday              BOOLEAN                             NOT NULL,
    is_irregular_holiday         BOOLEAN                             NOT NULL,
    created_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_events_id_fk
        FOREIGN KEY (event_id) REFERENCES bcrd.events (id)
);
COMMENT ON TABLE bcrd.shops IS '店舗';
COMMENT ON COLUMN bcrd.shops.id IS 'ID';
COMMENT ON COLUMN bcrd.shops.event_id IS 'イベントID';
COMMENT ON COLUMN bcrd.shops.no IS 'No.';
COMMENT ON COLUMN bcrd.shops.shop_name IS '店舗名';
COMMENT ON COLUMN bcrd.shops.image_url IS '画像URL';
COMMENT ON COLUMN bcrd.shops.google_url IS 'Google店舗URL';
COMMENT ON COLUMN bcrd.shops.tabelog_url IS '食べログURL';
COMMENT ON COLUMN bcrd.shops.official_url IS '公式サイトURL';
COMMENT ON COLUMN bcrd.shops.instagram_url IS 'インスタグラムURL';
COMMENT ON COLUMN bcrd.shops.address IS '住所';
COMMENT ON COLUMN bcrd.shops.business_days IS '営業日';
COMMENT ON COLUMN bcrd.shops.regular_holiday IS '定休日';
COMMENT ON COLUMN bcrd.shops.is_open_holiday IS '祝日営業フラグ';
COMMENT ON COLUMN bcrd.shops.is_irregular_holiday IS '不定休フラグ';
