CREATE TABLE IF NOT EXISTS bcrd.config
(
    id         SERIAL PRIMARY KEY,
    conf_name  VARCHAR(255)                        NOT NULL,
    conf_value VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE bcrd.config IS '設定値';
COMMENT ON COLUMN bcrd.config.id IS 'ID';
COMMENT ON COLUMN bcrd.config.conf_name IS '設定項目名';
COMMENT ON COLUMN bcrd.config.conf_value IS '設定値';
