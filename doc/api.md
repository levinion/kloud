# api

## fetch
GET https://kloud.botland.top/file/:id/*path

Payload: None

Return:

eg:  
```json
[
    "ZTg4OTZjYjZjYjAxNGQyZGY4NTE4ZmQ2ZTlhMDFmYzZhOTlmNDA4NDQxOTUyMDZlYWRmMjZlNTYxODk0NDhlMw==",
    "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ==",
    "ZWEwYjkyNTI2ZTA2ZTIxNTkxZGY0MmNmZmJlYWNjM2Q2MzA2YTliODMzNjJhMTc5N2I5NTQzNjUxNDJlNTUxZQ==",
    "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ=="
]
```

## post
POST https://kloud.botland.top/file/:id/*path

Payload: 

eg:  
```json
{
    "hashs": [
        "ZTg4OTZjYjZjYjAxNGQyZGY4NTE4ZmQ2ZTlhMDFmYzZhOTlmNDA4NDQxOTUyMDZlYWRmMjZlNTYxODk0NDhlMw==",
        "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ==",
        "ZWEwYjkyNTI2ZTA2ZTIxNTkxZGY0MmNmZmJlYWNjM2Q2MzA2YTliODMzNjJhMTc5N2I5NTQzNjUxNDJlNTUxZQ==",
        "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ=="
    ],
    "diffs": [
        {
            "hash": "ZTg4OTZjYjZjYjAxNGQyZGY4NTE4ZmQ2ZTlhMDFmYzZhOTlmNDA4NDQxOTUyMDZlYWRmMjZlNTYxODk0NDhlMw==",
            "content": "KLUv/QBY0QEA44CA44CA5LuO5LiN6auY5YW05Yir5Lq65Zyo5oiR5LiN6auY5YW055qE5pe25YCZ56yR5Ye65p2lCg=="
        },
        {
            "hash": "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ==",
            "content": "KLUv/QBYCQAACg=="
        },
        {
            "hash": "ZWEwYjkyNTI2ZTA2ZTIxNTkxZGY0MmNmZmJlYWNjM2Q2MzA2YTliODMzNjJhMTc5N2I5NTQzNjUxNDJlNTUxZQ==",
            "content": "KLUv/QBY8QEA44CA44CA5Zyo5LiA5YiH54Ot5oOF5b2T5Lit77yM54ix5oOF5piv5pyA6Ieq56eB6Ieq5Yip55qEIOOAggo="
        },
        {
            "hash": "Mjk1MTkyZWExZWM4NTY2ZDU2M2IxYTc1ODdlNWYwMTk4NTgwY2RiZDA0Mzg0MmY1MDkwYTRjMTk3YzIwYzY3YQ==",
            "content": "KLUv/QBYCQAACg=="
        }
    ]
}
```

Return: None



