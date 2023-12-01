# HTTP API documentation

## Routes for consumers API

All routes are under `/api/v1`.

### Offers

Get all the offers.

| Path       | Method | ContentType | QueryString |
| :--------- | :----- | :---------- | :---------- |
| `/offers/` | GET    | JSON        | `c`         |

Response schema:

```json
[]{
  "id": 0,
  "title": "...",
  "role": "...",
  "company": "...",
  "content": "...",
  "contract": "...",
  "location": "...",
  "salary": 0,
  "contact_info": "...",
  "created_at": "yyy-mm-dd"
}
```

If the `c=true` (`c` is for compact representation) query string is provided:

```json
{
  "id": 0,
  "title": "...",
  "role": "...",
  "company": "...",
  "contract": "...",
  "salary": 0,
  "created_at": "yyy-mm-dd"
}
```

### Offer by id

Get an offer by id.

| Path           | Method | ContentType | QueryString |
| :------------- | :----- | :---------- | :---------- |
| `/offers/{id}` | GET    | JSON        | None        |

Response schema:

```json
{
  "id": 0,
  "title": "...",
  "role": "...",
  "company": "...",
  "content": "...",
  "contract": "...",
  "location": "...",
  "salary": 0,
  "contact_info": "...",
  "created_at": "yyy-mm-dd"
}
```

### Learning Resource

Get all the learning resources.

| Path          | Method | ContentType | QueryString |
| :------------ | :----- | :---------- | :---------- |
| `/resources/` | GET    | JSON        | `c`         |

Response schema:

```json
[]{
  "id": 0,
  "title": "...",
  "author": "...",
  "summary": "...",
  "content": "...",
  "created_at": "yyy-mm-dd"
}
```

If the `c=true` query string is provided:

```json
[]{
  "id": 0,
  "title": "...",
  "author": "...",
  "created_at": "yyy-mm-dd"
}
```

### Resource by id

Get a learning resource by id.

| Path              | Method | ContentType | QueryString |
| :---------------- | :----- | :---------- | :---------- |
| `/resources/{id}` | GET    | JSON        | None        |

Response schema:

```json
{
  "id": 0,
  "title": "...",
  "author": "...",
  "summary": "...",
  "content": "...",
  "created_at": "yyy-mm-dd"
}
```

### Companies

Get all the companies registered in the system.

| Path          | Method | ContentType | QueryString |
| :------------ | :----- | :---------- | :---------- |
| `/companies/` | GET    | JSON        | None        |

Response schema:

```json
[]{
  "id": 0,
  "name": "...",
  "logo_url": "..."
}
```
