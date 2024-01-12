-- this migration introduces a PK as a compound key => Partition Key as uuid and Clustering Key as time
-- due to compound primary key now each insert operation with unique key will create a new record (row), but all records related to this ID will create a  big Partition
-- to avoid this as example can be used compound Partition key: PRIMARY KEY ((pet_chip_id,date), time)
-- this will resolve that each day will be created a new record and partition size will be limited by the day
--
-- Best Practices for selecting a partition and clustering key:
-- 1.Partitions are not too small.
-- 2.No large partitions
-- 3.Adequate per query (query according to partition key)
-- 4.The tradeoff of all expected queries
CREATE TABLE IF NOT EXISTS heartrate_v2 (
   pet_chip_id uuid,
   time timestamp,
   heart_rate int,
   PRIMARY KEY (pet_chip_id, time)
);