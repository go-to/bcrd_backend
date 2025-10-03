CREATE TABLE IF NOT EXISTS bcrd.stamps
(
    id              SERIAL
        PRIMARY KEY,
    user_id         VARCHAR(255) NOT NULL,
    shop_id         INTEGER      NOT NULL,
    number_of_times INTEGER      NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP,
    CONSTRAINT stamps_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES bcrd.shops (id)
);
COMMENT ON TABLE bcrd.stamps IS 'スタンプ';
COMMENT ON COLUMN bcrd.stamps.id IS 'ID';
COMMENT ON COLUMN bcrd.stamps.user_id IS 'ユーザーID';
COMMENT ON COLUMN bcrd.stamps.shop_id IS '店舗ID';
COMMENT ON COLUMN bcrd.stamps.number_of_times IS '回数';
