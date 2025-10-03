CREATE TABLE IF NOT EXISTS bcrd.shops_time
(
    id          SERIAL PRIMARY KEY,
    shop_id     INT                                 NOT NULL,
    week_number INT                                 NOT NULL,
    day_of_week INT                                 NOT NULL,
    start_time  TIME                                NOT NULL,
    end_time    TIME                                NOT NULL,
    is_holiday  BOOLEAN                             NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_time_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES bcrd.shops (id)
);
COMMENT ON TABLE bcrd.shops_time IS '店舗営業時間';
COMMENT ON COLUMN bcrd.shops_time.id IS 'ID';
COMMENT ON COLUMN bcrd.shops_time.shop_id IS '店舗ID';
COMMENT ON COLUMN bcrd.shops_time.week_number IS '週番号';
COMMENT ON COLUMN bcrd.shops_time.day_of_week IS '曜日';
COMMENT ON COLUMN bcrd.shops_time.start_time IS '営業開始時間';
COMMENT ON COLUMN bcrd.shops_time.end_time IS '営業終了時間';
COMMENT ON COLUMN bcrd.shops_time.is_holiday IS '定休日フラグ';
