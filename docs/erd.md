## 🗂 Entity Relationship Diagram (ERD)

Below is a high-level representation of the database schema used by this API. It models the relationships between users, posts, comments, and authentication mechanisms.

```mermaid
erDiagram
    USERS ||--o{ POSTS : creates
    USERS ||--o{ COMMENTS : writes
    POSTS ||--o{ COMMENTS : has
    USERS ||--|| AUTH_TOKENS : owns

    USERS {
        UUID id PK
        string username
        string email
        string password_hash
        string first_name
        string last_name
        string phone_number
        datetime created_at
    }

    POSTS {
        UUID id PK
        UUID user_id FK
        string content
        datetime created_at
    }

    COMMENTS {
        UUID id PK
        UUID user_id FK
        UUID post_id FK
        string content
        datetime created_at
    }

    AUTH_TOKENS {
        UUID id PK
        UUID user_id FK
        string token
        datetime expires_at
    }
```