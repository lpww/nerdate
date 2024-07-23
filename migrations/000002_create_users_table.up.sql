CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  deleted_at timestamp(0) with time zone,
  name text NOT NULL,
  gender text NOT NULL,
  dob date NOT NULL,
  ascii_art text NOT NULL,
  description text NOT NULL,
  email citext UNIQUE NOT NULL,
  password_hash bytea NOT NULL,
  activated bool NOT NULL,
  version integer NOT NULL DEFAULT 1
);
