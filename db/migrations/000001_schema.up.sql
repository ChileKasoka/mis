-- Step 1: Create the users table
CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,             -- Unique identifier for each user
    "name" varchar NOT NULL,           -- Name of the user
    "email" varchar UNIQUE NOT NULL,   -- Unique email for each user
    "password" varchar NOT NULL,       -- Password (hashed) for authentication
    "user_type" varchar NOT NULL,  -- Type of user (e.g., admin, vendor, customer)
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Timestamp for when the user was created
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Timestamp for when the user was last updated
);


CREATE TABLE "contacts" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "vendors" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar NOT NULL,
  "location" varchar ,
  "website" varchar ,
  "biography" varchar ,
  "profile_picture" varchar ,
  "business_type" varchar NOT NULL,
  "experience" varchar ,
  "certification" varchar ,
  "active" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE days_available (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "vendor_id" UUID REFERENCES vendors(id) ON DELETE CASCADE,
    "day" TEXT NOT NULL CHECK (day IN ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'))
);

CREATE TABLE "hours_available" (
  "id" uuid PRIMARY KEY,
  "hour" varchar,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "customers" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid
);

CREATE TABLE "appointments" (
  "id" uuid PRIMARY KEY,
  "customer_id" uuid NOT NULL,
  "vendor_id" uuid NOT NULL,
  "date" date NOT NULL,
  "time_slot_id" uuid NOT NULL,
  "service_id" uuid NOT NULL,
  "status" varchar NOT NULL, -- Assuming AppointmentStatus is an enum
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "time_slots" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid NOT NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "is_booked" boolean NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "feedback" (
  "id" uuid PRIMARY KEY,
  "appointment_id" uuid,
  "rating" int,
  "comment" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "services" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" numeric NOT NULL,
  "duration" int NOT NULL, -- Duration in minutes
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "payments" (
  "id" uuid PRIMARY KEY,
  "appointment_id" uuid,
  "customer_id" uuid,
  "vendor_id" uuid,
  "amount" decimal(10,2),
  "payment_method" payment_method_enum,
  "status" payment_status_enum,
  "payment_date" timestamp,
  "transaction_id" varchar,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

-- Step 2: Create foreign keys
ALTER TABLE "contacts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "vendors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("time_slot_id") REFERENCES "time_slots" ("id");

ALTER TABLE "time_slots" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

-- ALTER TABLE "vendor_availability" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "feedback" ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

-- ALTER TABLE "payments" ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id");

-- ALTER TABLE "payments" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

-- ALTER TABLE "payments" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");
