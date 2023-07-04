# Dockerized Monitoring with Grafana and Prometheus

This repository provides a streamlined monitoring solution using **Docker** and **Docker Compose v2**. The setup enables efficient monitoring and visualization of various metrics from your applications and infrastructure. By default the setup include JSON file as template dashboards from [Node Exporter Full](https://grafana.com/grafana/dashboards/1860).

[![wakatime](https://wakatime.com/badge/user/8c51dfaf-cc71-4c33-bb4f-07b1a77dce06/project/a1414466-0d7b-444b-a4fd-92d7e5633550.svg)](https://wakatime.com/badge/user/8c51dfaf-cc71-4c33-bb4f-07b1a77dce06/project/a1414466-0d7b-444b-a4fd-92d7e5633550)

## Table of contents
  * [Requirements](#requirements)
  * [Key Features](#key-features)
  * [Getting Started](#getting-started)
    + [Run the monitoring stack locally](#run-the-monitoring-stack-locally)

## Requirements

The setup is based on the following tools:

- [Docker](https://www.docker.com/)
- [Docker Compose v2](https://docs.docker.com/compose/)
- [Grafana](https://grafana.com/)
- [Prometheus](https://prometheus.io/)
- [Node Exporter](https://prometheus.io/docs/guides/node-exporter/)
- [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/)


## Key Features

- **Dockerized deployment** - Easily spin up the monitoring stack using Docker and Docker Compose v2, ensuring portability and reproducibility across different environments.

- **Grafana:** Utilize Grafana's robust dashboarding capabilities to visualize and analyze your metrics, allowing you to gain valuable insights into the performance and health of your systems.

- **Prometheus:** Leverage Prometheus as the monitoring and alerting toolkit to collect and store time-series data from your applications, enabling efficient metric querying and triggering of alerts.

- **Easy configuration:** Customize the monitoring setup according to your specific requirements through simple configuration files, making it easy to adapt the monitoring stack to your environment.


## Getting Started

### Run the monitoring stack locally

**_Note:_** The following instructions assume that you have Docker and Docker Compose v2 installed on your machine. Please make have all needed ports available on your machine. All ports used by the monitoring stack are listed in the `docker compose` file.

1. Clone the repository:

```bash
https://github.com/Caremitou/Infrastructure.git
```

2. Navigate to the `server monitoring` directory:

```bash
cd server_monitoring
```

3. Run the following command to start the monitoring stack:

```bash
docker compose up 
```

Then you can acces to your Grafana dashboard on [localhost:3000](http://localhost:3000) with the default credentials `admin:admin`.
	local