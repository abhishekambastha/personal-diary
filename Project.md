# A personal diary app

## Functional Requirement

1.  Add entries with a date and text
2.  Entries have a date
3.  Show more recent first
4.  Attach pictures
5.  Store them somewhere
6.  User account creation and authentication


## API Spec

1. POST /entry

```json 
    {
        "date":"timestamp",
        "text":"some text",
    } 
```

2. GET /entry
3. GET entry/{id}
4. Add authentication mechanim - JWT, tenantId management etc.



## DB Schema

### Entries


|  id  |       date      |       text       | picture | tid |
|------|-----------------|------------------|---------|-----|
| uuid | Epoch timestamp | some text or ref | ref url | t1  |



