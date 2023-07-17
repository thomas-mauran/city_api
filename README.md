<div align="center">
<h1>City api</h1>
<img src="https://media.tenor.com/GyXuxAaiWjYAAAAC/studio-ghibli-city.gif" alt="cobol mystery number game" style="width: 500px"/>
</div>

### Description

City api is a simple go api that returns a list of cities and their respective infos. The main goal here was to learn ci/cd with github actions and learn the good practices of go.

### How to run

```bash
# Clone this repository
git clone git@github.com:thomas-mauran/city_api.git

# Go into the repository
cd city-api

# Setup the .env file
cp .env.local .env

# Run the app
cd ..
docker compose up -d

```

### The endpoints

| Method | Endpoint  | Description                                |
| ------ | --------- | ------------------------------------------ |
| GET    | /cities   | Returns a list of cities                   |
| POST   | /cities   | Creates a new city                         |
| GET    | /\_health | Return the state of the database connexion |

### Monitor the app

```bash
# Run the monitoring docker compose
cd server_monitoring
docker compose up -d
```

You can access the grafana dashboard here : http://localhost:3000

Username: admin
Password: admin

You will have to set a new password. You can now to see the metrics of the app in the dashboard/node-exporter section

### The CI/CD

The CI containes multiple jobs doing the following things:

- The lint-city job lints the Go code in the city-api directory and runs unit tests.
- The build-and-push-images job builds a Docker image from the code in the city-api directory and pushes it to the ghcr.io registry.
- The prod-deploy-test job deploys the new Docker image to a production-like environment using Kind.
- The commit-new-image job updates the tag of the image in the Helm chart and commits the changes to the repository.
