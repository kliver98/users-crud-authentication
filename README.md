# users-crud-authentication
Service for authenticate users by their username and password
This application allows:

1. Valid user login by username and password

## How To Download this version and execute?

Steps (assuming you already have git and yarn):

1. Go to root folder where you want the project it's located and run
```
git clone https://github.com/kliver98/users-crud-authentication
```
2. Open folder downloaded in desired code editor and run on the terminal:
```
go mod download
```
3. Lastly to run the project type on the terminal:
```
go run main.go
```
Note 1: When run start command, if nothing it's wrong, the console will NOT show information about that loads. You must go to open web browser and access [http://localhost:3001/api/latest/users/](http://localhost:3001/api/latest/users/) if 3001 port wasn't in use.
Note 2: You must have **_.env.development.local_** file in root project to can load url and connect mongodb. This file will be passed if you collaborate on this project. Otherwise, you can create that file and set a new variable called **MONGODB_URL** and set to your own project on mongodb. Example: 
```
MONGODB_URL = mongodb+srv://user:password@cluster0...
```
Note 3: Delete unused dependencies: 
```
go mod tidy
```
## How To Update your fork repository

Run the following commands:

```bash
  git remote add upstream git@github.com:kliver98/users-crud-authentication.git
  git pull upstream main
```

If you have altered it, you then need to rebase it.

```bash
  git rebase upstream/main
````

## How to run Tests ?

1. Contract Test:
```
go test -v ./test/contract/api
```
or if you have installed make
```
make provider
```
*Note: You need _**.env.development.local**_ file where are setted variables to correct run. These are PACT_BROKER_BASE_URL, PACT_BROKER_TOKEN, API (localhost where client will start with port) and MONGODB_URL for connection with database. This file will be send to you if collaborate on this project (if not sent yet, [message me](mailto:kliver1998@gmail.com)), other wise you must create and set your own variables.
