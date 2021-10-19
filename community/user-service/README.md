# User Service 

## What's this?

An API that returns fake user data. 


* Setup 
- clone this repo and cd into the user-serive directory 

- run the api `go run main.go --port <port-number>`


## Endpoints


### Endpoint: /api/user

```bash
curl localhost:8080/api/user/
```

### Response

```json
{
  "firstname": "Yvette",
  "lastname": "Bode",
  "email": "myrtiebartoletti@nikolaus.io",
  "phone": "(577)241-5935",
  "gender": "female"
}

```

### Endpoint: /api/users


```bash
curl localhost:8080/api/users
```

### Response

```json
[
  {
    "firstname": "Thalia",
    "lastname": "Sawayn",
    "email": "mosesreichert@hayes.io",
    "phone": "1-689-190-6921",
    "gender": "male"
  },
  {
    "firstname": "Sincere",
    "lastname": "Gaylord",
    "email": "guyrempel@wiza.info",
    "phone": "1-742-905-3752",
    "gender": "male"
  },
  {
    "firstname": "Josefa",
    "lastname": "Yost",
    "email": "stacykoelpin@collier.name",
    "phone": "346.046.8251",
    "gender": "female"
  },
  {
    "firstname": "Carmelo",
    "lastname": "Hudson",
    "email": "tristonfranecki@thompson.org",
    "phone": "(211)778-9105",
    "gender": "female"
  },
  {
    "firstname": "Kiana",
    "lastname": "Fay",
    "email": "rosalynferry@becker.com",
    "phone": "154.918.9844",
    "gender": "male"
  },
  {
    "firstname": "Isadore",
    "lastname": "Jakubowski",
    "email": "alycemetz@waelchi.com",
    "phone": "(555)809-3065",
    "gender": "male"
  },
  {
    "firstname": "Laila",
    "lastname": "Fadel",
    "email": "delaneylockman@jewess.org",
    "phone": "(031)786-4836",
    "gender": "female"
  },
  {
    "firstname": "Yoshiko",
    "lastname": "Corkery",
    "email": "wilburnmcdermott@rodriguez.net",
    "phone": "139.132.4594",
    "gender": "female"
  },
  {
    "firstname": "Delphia",
    "lastname": "Veum",
    "email": "estellboyer@bednar.io",
    "phone": "748-473-5392",
    "gender": "female"
  },
  {
    "firstname": "Mya",
    "lastname": "Franecki",
    "email": "avajakubowski@gorczany.biz",
    "phone": "1-181-108-6352",
    "gender": "male"
  }
]

```
