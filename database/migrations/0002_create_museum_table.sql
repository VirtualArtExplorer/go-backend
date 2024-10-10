IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[museums]') AND type in (N'U'))
BEGIN
    CREATE TABLE museums (
        id INT IDENTITY(1,1) PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        image VARBINARY(MAX),
        category1 VARCHAR(100),
        category2 VARCHAR(100),
        link VARCHAR(255),
        address VARCHAR(255),
        cep VARCHAR(20),
        city VARCHAR(100),
        state VARCHAR(100),
        information TEXT,
        manager_id INT NOT NULL,
        active BIT DEFAULT 1,
        CONSTRAINT fk_manager FOREIGN KEY (manager_id) REFERENCES managers(id)
    );
END
