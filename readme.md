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

## Debts
### Feature Debts
---
1. Negative case for API soccer
2. Multiple DBMS support using adapters

### Technical Debts
---
1. Unit test for soccer handler & controller. Currently removed due to breaking test
2. Singleton pattern for passing config and DB needs to refactored
3. Usage of channel
4. Usage of proper context

