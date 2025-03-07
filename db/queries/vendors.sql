-- name: CreateVendor :one
INSERT INTO vendors (id, first_name, last_name, email, password, phone, address, location, website, biography, profile_picture, business_type, experience, certification, active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
RETURNING *;

-- -- name: GetVendor :one
-- SELECT *
-- FROM vendors
-- WHERE id = $1
-- LIMIT 1;

-- name: GetVendor :one
SELECT 
    v.id, v.first_name, v.last_name, v.email, v.biography, v.profile_picture, 
    v.active, v.created_at, v.updated_at, 
    COALESCE(json_agg(d.day) FILTER (WHERE d.day IS NOT NULL), '[]') AS days_available
FROM vendors v
LEFT JOIN days_available d ON v.id = d.vendor_id
WHERE v.id = $1
GROUP BY v.id;


-- name: UpdateVendor :one
UPDATE vendors
SET first_name = $2,
    last_name = $3,
    email = $4,
    password = $5,
    phone = $6,
    address = $7,
    location = $8,
    website = $9,
    biography = $10,
    profile_picture = $11,
    business_type = $12,
    experience = $13,
    certification = $14,
    active = $15,
    updated_at = $16
WHERE id = $1
RETURNING *;

-- name: DeleteVendor :one
DELETE FROM vendors
WHERE id = $1
RETURNING id;
