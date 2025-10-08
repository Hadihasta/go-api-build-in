CREATE TABLE IF NOT EXISTS posts(
id bigserial PRIMARY KEY,
title text NOT NULL,
user_id bigint NOT NULL,
content text NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW())