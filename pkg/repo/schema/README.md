# Database schema

Run these files against the same PostgreSQL database in this order:

1. `05_users.sql`
2. `02_crop.sql`
3. `01_mandi.sql`
4. `03_fertilizer.sql`
5. `04_soil.sql`
6. `06_schemes.sql`
7. `07_weather.sql`
8. `08_equipment.sql`

The files create tables in the `kisansathi` schema. The application must use
`search_path=kisansathi,public` or fully qualified table names when querying
these tables.

For Neon, run the files with `psql` using the Neon connection string. Do not
commit the database password; provide it through `PGPASSWORD` or a secret
manager.
