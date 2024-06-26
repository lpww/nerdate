CREATE TABLE IF NOT EXISTS swipes (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  liked boolean,
  swiped_user_id uuid NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  deleted_at timestamp(0) with time zone,
  version integer NOT NULL DEFAULT 1
);
