CREATE DATABASE IF NOT EXISTS sampledb_staging;
CREATE DATABASE IF NOT EXISTS sampledb_development;

GRANT ALL PRIVILEGES ON *.* TO 'sample'@'%';