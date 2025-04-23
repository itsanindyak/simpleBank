ALTER TABLE accounts 
ALTER COLUMN created_at SET DEFAULT (now() AT TIME ZONE 'Asia/Kolkata');

ALTER TABLE entries 
ALTER COLUMN created_at SET DEFAULT (now() AT TIME ZONE 'Asia/Kolkata');

ALTER TABLE transfers 
ALTER COLUMN created_at SET DEFAULT (now() AT TIME ZONE 'Asia/Kolkata');
