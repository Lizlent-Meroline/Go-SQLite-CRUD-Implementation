# Go-SQLite-CRUD-Implementation

**Project Overview**

This project is a backend system written in Go that manages a product inventory using a SQLite database.

It focuses on:

Database design
CRUD operations
Transaction handling
Safe SQL execution using prepared statements
Clean separation of business logic

The system is designed as a library-style backend module, not tied to any UI or HTTP layer.

**Objectives**
Implement CRUD operations using Go and SQLite
Ensure safe database interactions using prepared statements
Handle errors properly for all operations
Support filtering and querying of records
Implement transaction support for multi-step updates
Maintain clean separation between database and business logic

 **System architecture**

       Application Layer (Store Methods)
              ↓
      Database Layer (SQLite)
              ↓
      Persistent Storage (store.db file)


Author

LIZLENT MEROLINE
