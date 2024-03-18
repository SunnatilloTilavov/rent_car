/*sudo -u postgres psql -d sqldatabase*/
CREATE TABLE IF NOT EXISTS cars (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name Varchar(50) NOT NULL,
    brand Varchar(20) NOT NULL,
    model Varchar(30) NOT NULL,
    hourse_power INTEGER DEFAULT 0,
    colour VARCHAR(20) NOT NULL DEFAULT 'black',
    engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS customers (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    gmail VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
)

CREATE UNIQUE INDEX index_phone
ON Customers(phone,deleted_at);

CREATE TABLE orders (
  "id" uuid PRIMARY KEY,
  "customer_id" uuid REFERENCES customers(id),
  "car_id" uuid REFERENCES cars(id),
  "from_date" TIMESTAMP,
  "to_date" TIMESTAMP,
  "status" varchar(50),
  "paid" BOOLEAN DEFAULT true,
  "amount" DECIMAL(10,2),
  "created_ad"  TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP
);

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("car_id") REFERENCES "cars" ("id");