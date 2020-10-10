# API Server for test
This project is to serve as a result from a test

## Get Started

### Installation
---
1. Adjust `config.yaml` as you see fit. Don't forget to configure your DB
2. Let's initiate the DB first by doing
>```
>make migrate
>```
3. Tidying up the dependencies
>```
>go mod tidy
>```
4. That's it :D

### Testing
---
Test the app by doing
>```
>make test
>```

### Run locally
---
Run the app locally by doing
>```
>make run-dev
>```