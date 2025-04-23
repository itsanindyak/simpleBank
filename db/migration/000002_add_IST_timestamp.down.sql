ALTER TABLE accounts 
ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE entries 
ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE transfers 
ALTER COLUMN created_at SET DEFAULT now();
