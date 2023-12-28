-- few notes here: inserting in this table will always be a upsert operation due to PK is only ID
CREATE TABLE IF NOT EXISTS heartrate_v1 (
   pet_chip_id uuid,
   time timestamp,
   heart_rate int,
   PRIMARY KEY (pet_chip_id)
);
