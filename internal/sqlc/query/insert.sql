-- name: CreateRecords :one
INSERT INTO records 
  (name, status)
    values ($1, $2)
  ON CONFLICT (name) 
  DO 
    UPDATE SET status = EXCLUDED.status,
               updated_at = EXCLUDED.updated_at,
  RETURNING id;

