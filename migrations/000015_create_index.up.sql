CREATE INDEX idx_shops_event_id ON shops(event_id);
CREATE INDEX idx_shops_location_shop_id ON shops_location(shop_id);
CREATE INDEX idx_shops_time_composite ON shops_time(shop_id, week_number, day_of_week, is_holiday);
CREATE INDEX idx_stamps_composite ON stamps(shop_id, user_id, deleted_at);
CREATE INDEX idx_shops_event_year ON events(year);
CREATE INDEX idx_shops_no_name ON shops(no, shop_name);
CREATE INDEX idx_shops_location_spatial ON shops_location USING GIST(location);
CREATE INDEX idx_shops_time_temporal ON shops_time(start_time, end_time);
CREATE INDEX idx_stamps_user_id ON stamps(user_id);
