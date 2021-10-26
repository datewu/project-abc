-- name: GetRecordByID :one
SELECT * FROM records 
WHERE videos.id = $1;

