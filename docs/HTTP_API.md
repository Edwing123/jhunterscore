# HTTP API documentation

## Routes for consumers API

All routes are under `/api/v1`.

### Offers

Get all the offers.

| Path       | Method | ContentType | QueryString |
| :--------- | :----- | :---------- | :---------- |
| `/offers/` | GET    | JSON        | compact     |

Response schema:

```json
[]{
  "id": 0,
  "title": "...",
  "role": "...",
  "company": "...",
  "description": "...",
  "requirements": ["..."],
  "contract": "...",
  "location": "...",
  "salary": "...",
  "benefits": ["..."],
  "contact_info": "...",
  "published_at": "yyy-mm-dd"
}
```

If the `compact` query string is provided:

```json
{
  "id": 0,
  "title": "...",
  "contract": "...",
  "salary": "...",
  "published_at": "yyy-mm-dd"
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
  "description": "...",
  "requirements": ["..."],
  "contract": "...",
  "location": "...",
  "salary": "...",
  "benefits": ["..."],
  "contact_info": "...",
  "published_at": "yyy-mm-dd"
}
```

### Learning Resource

Get all the learning resources.

| Path          | Method | ContentType | QueryString |
| :------------ | :----- | :---------- | :---------- |
| `/resources/` | GET    | JSON        | compact     |

Response schema:

```json
[]{
  "id": 0,
  "title": "...",
  "author": "...",
  "summary": "...",
  "content": "...",
  "published_at": "yyy-mm-dd"
}
```

If the `compact` query string is provided:

```json
[]{
  "id": 0,
  "title": "...",
  "author": "...",
  "published_at": "yyy-mm-dd"
}
```

### resources/id

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
  "published_at": "yyy-mm-dd"
}
```****
