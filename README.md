# DBconnection

A small Go program that connects to MySQL and creates a `PRODUCTS` table if it
does not already exist.

## Run

Set `MYSQL_DSN` to your MySQL connection string, then run the program:

```powershell
$env:MYSQL_DSN = "username:password@tcp(127.0.0.1:3306)/database"
go run .
```

Do not commit real credentials.
