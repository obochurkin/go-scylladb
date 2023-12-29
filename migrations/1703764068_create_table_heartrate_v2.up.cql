-- this migration introduces a PK as a compound key => Partition Key as uuid and Clustering Key as time
-- due to compound key now each insert operation will create a new record (row)
CREATE TABLE IF NOT EXISTS heartrate_v2 (
   pet_chip_id uuid,
   time timestamp,
   heart_rate int,
   PRIMARY KEY (pet_chip_id, time)
);