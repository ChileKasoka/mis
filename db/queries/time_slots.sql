-- name: CreateTimeSlot :one
INSERT INTO time_slots (id, vendor_id, start_time, end_time, is_booked, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, vendor_id, start_time, end_time, is_booked, created_at, updated_at;

-- name: GetTimeSlot :one
SELECT id, vendor_id, start_time, end_time, is_booked, created_at, updated_at
FROM time_slots
WHERE id = $1
LIMIT 1;

-- name: UpdateTimeSlot :one
UPDATE time_slots
SET vendor_id = $2,
    start_time = $3,
    end_time = $4,
    is_booked = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, vendor_id, start_time, end_time, is_booked, updated_at;

-- name: DeleteTimeSlot :one
DELETE FROM time_slots
WHERE id = $1
RETURNING id;
