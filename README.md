# go-responses
Some structs to have standard responses everywhere.

## Unified responses in services :
Examples :

```go
    ... // sending not found error
        return ctx.Status(fiber.StatusNotFound).JSON(responses.NotFoundError("url"))
    ...
```

response :

```json
{
    "messages": {
        "url": ["url not found!"]
    },
    "status_code": 404,
    "result": "ERROR",
    "data": null
}
```

