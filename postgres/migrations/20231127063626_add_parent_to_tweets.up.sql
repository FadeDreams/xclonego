
ALTER TABLE tweets ADD COLUMN parent_id UUID REFERENCES tweets (id) ON DELETE CASCADE;
