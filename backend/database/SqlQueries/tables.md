# Tables

---

## Overpass Business Scraper

- Industry (Lookup Table)

```sql
CREATE TABLE IndustryMaster (
    IndustryMasterID INT PRIMARY KEY AUTO_INCREMENT
    IndustryName TEXT,
    OverpassTags TEXT,
)
```

- Overpass (Raw JSON data)

```sql
CREATE TABLE Overpass (
    OverpassID PRIMARY KEY,
    DateRecordCreated DATETIME DEFAULT TIMESTAMP(),
    DateRecordUpdated DATETIME,
    OverpassJSON JSON,
)
```

- Business Register table<br>
  Contains information about businesses including:
  - Industry
  - Location

```sql
CREATE TABLE BusinessRegistar (
    RegisterID INT PRIMARY KEY AUTO_INCREMENT,
    DateRecordCreated DATETIME DEFAULT TIMESTAMP(),
    BusinessName TEXT,
    LocationLat FLOAT,
    LocationLon FLOAT,
    AddressText TEXT,
    IndustryMasterID INT,
    OverpassID INT,
)
```

---

## Forum Scraper

- Forums (Lookup Table)

```sql
CREATE TABLE Forums (
    ForumID INT PRIMARY KEY AUTO_INCREMENT,
    DateRecordCreated DATETIME DEFAULT TIMESTAMP(),
    ForumName TEXT,
    ForumURL TEXT,
)
```

- ForumData (Raw JSON Data)

```sql
CREATE TABLE ForumData (
    ForumDataID INT PRIMARY KEY AUTO_INCREMENT,
    DateRecordCreated DATETIME DEFAULT TIMESTAMP(),
    ForumJSON JSON,
)
```

- ForumTransactions

Currated data from the ForumData Table

```sql
CREATE TABLE ForumTransactions (
    TransactionID INT PRIMARY KEY AUTO_INCREMENT,
    ForumDataID INT,
    TransactionType TEXT, -- Question or Answer
    TransactionText TEXT,
)
```
