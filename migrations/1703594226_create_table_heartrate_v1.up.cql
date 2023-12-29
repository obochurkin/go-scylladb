-- few notes here: inserting in this table will always be a upsert operation due to PK is only ID
-- row based on ID (PK) will be updated on each insert
CREATE TABLE IF NOT EXISTS heartrate_v1 (
   pet_chip_id uuid,
   time timestamp,
   heart_rate int,
   PRIMARY KEY (pet_chip_id)
);
