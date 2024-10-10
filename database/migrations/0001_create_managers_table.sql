IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[managers]') AND type in (N'U'))
BEGIN
    CREATE TABLE managers (
        id INT IDENTITY(1,1) PRIMARY KEY,
        email VARCHAR(100) NOT NULL UNIQUE,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL,
        active BIT DEFAULT 1
    );
END
