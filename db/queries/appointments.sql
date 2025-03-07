-- name: CreateAppointment :one
INSERT INTO appointments (id, customer_id, vendor_id, date, time_slot_id, service_id, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, customer_id, vendor_id, date, time_slot_id, service_id, status, created_at, updated_at;

-- name: GetAppointmentById :one
SELECT id, customer_id, vendor_id, date, time_slot_id, service_id, status, created_at, updated_at
FROM appointments
WHERE id = $1
LIMIT 1;

-- name: GetAllAppointments :many
SELECT id, customer_id, vendor_id, date, time_slot_id, service_id, status, created_at, updated_at
FROM appointments;

-- name: UpdateAppointment :one
UPDATE appointments
SET customer_id = $2,
    vendor_id = $3,
    date = $4,
    time_slot_id = $5,
    service_id = $6,
    status = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, customer_id, vendor_id, date, time_slot_id, service_id, status, created_at, updated_at;

-- name: DeleteAppointment :one
DELETE FROM appointments
WHERE id = $1
RETURNING id;

-- name: CheckConfirmedAppointment :one
SELECT EXISTS (
    SELECT 1 
    FROM appointments 
    WHERE vendor_id = $1 
    AND time_slot_id = $2 
    AND date = $3 
    AND status = 'confirmed'
) AS exists;
