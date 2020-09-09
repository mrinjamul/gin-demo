# gin-demo

Go Gin Demo GoLang

## Step to RUN projects

**1. Fetch Dependency:**

```sh
go mod tidy
```

**2. Environment Settings:**

Copy `docker/.env.sample` to `docker/.env` and write environment details. Then export using this command.

```sh
cp docker/.env{.sample,}
export $(cat docker/.env | xargs)
```

**3. Build Project:**

```sh
go build
```

**4. Run Application:**

```sh
./gin-demo
```

Or run with docker:
```sh
cd docker
docker-compose up --build
```

Now your api is running on 3000 port (if you use docker then port is 3001).

