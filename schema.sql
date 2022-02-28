CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS customers(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL,
  email CHAR(100) UNIQUE NOT NULL,
  password TEXT,
  address VARCHAR,
  balance DECIMAL(12, 2) DEFAULT 0
);