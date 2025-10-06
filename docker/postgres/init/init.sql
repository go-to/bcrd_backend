-- スキーマ作成
CREATE SCHEMA IF NOT EXISTS bcrd;
COMMENT ON SCHEMA bcrd IS 'バカルディハイボールスタンプラリー';

-- postgis有効化
CREATE EXTENSION postgis WITH SCHEMA bcrd;
