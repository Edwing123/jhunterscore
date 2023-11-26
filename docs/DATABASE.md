# Database design

This document explains the database design for JobsHunters.

## Entity-Relationship Diagram

![Entity-Relationship Diagram](./assets/erd.png)

## Relational Model Diagram

```txt
users(user_id, username, password, first_name, last_name, role, is_active)

locations(location_id, name)

offers(<pk:offer_id>, title, role, company, content, contract, salary, contact_info, created_at, is_published, <fk:user_id>)

resources(<pk:resource_id>, title, summary, content, created_at, is_published, <fk:user_id>)
```

