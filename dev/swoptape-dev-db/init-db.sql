-- Create the test database
-- The main 'swoptape' database is created automatically via POSTGRES_DB env var
CREATE DATABASE swoptape_test;
GRANT ALL PRIVILEGES ON DATABASE swoptape_test TO swoptape;
