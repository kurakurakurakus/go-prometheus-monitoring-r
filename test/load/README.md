# Load Test

Load testing is performed using [Vegeta](https://github.com/tsenart/vegeta) running in Docker.

## Prerequisites

Make sure Docker is running and start the application stack:

```bash
docker compose up -d
```

Verify that the services are running:

```bash
docker compose ps
```

The `-d` option runs the containers in detached mode, allowing the terminal to be used for other commands.

## Configuration

Load-test configuration is defined in:

```text
test/load/run.sh
```

Example:

```bash
RATE="20/s"
DURATION="1m"
```

### Request Rate

`RATE` controls the total number of requests sent per second.

For example:

```bash
RATE="20/s"
```

runs at approximately 20 requests per second.

To increase the load:

```bash
RATE="50/s"
```

or:

```bash
RATE="100/s"
```

The rate is the **total load across all configured targets**, not per target.

### Duration

`DURATION` controls how long the load test runs.

For example:

```bash
DURATION="1m"
```

runs the test for 1 minute.

Other examples:

```bash
DURATION="30s"
DURATION="5m"
DURATION="10m"
```

## Targets

HTTP targets are configured in:

```text
test/load/targets.txt
```

Example:

```text
GET http://api:8080/v1/random
```

The API hostname should use the Docker Compose service name because Vegeta runs inside Docker.

For example, if the API service is named `api`:

```yaml
services:
  api:
    ...
```

use:

```text
http://api:8080/v1/random
```

instead of:

```text
http://localhost:8080/v1/random
```

## Run Load Test

Run the load test from the project root:

```bash
./test/load/run.sh
```

The script will:

1. Send requests to the configured targets.
2. Generate the configured request rate.
3. Run for the configured duration.
4. Save the raw Vegeta results.
5. Generate a summary report.

For example, with:

```bash
RATE="20/s"
DURATION="1m"
```

the test generates approximately:

```text
20 requests/second × 60 seconds = ~1,200 requests
```

## Results

Raw Vegeta results are saved to:

```text
test/load/results/latest.bin
```

The results directory is ignored by Git and should not be committed.

To generate a report from an existing result:

```bash
docker compose run --rm vegeta \
  report /load/results/latest.bin
```

The report includes information such as:

* Total requests
* Actual request rate
* Throughput
* Success rate
* Latency
* P50
* P90
* P95
* P99
* HTTP status codes


```

## Stopping the Application

When finished, stop the application stack:

```bash
docker compose down
```

This stops and removes the containers but **does not remove named Docker volumes**.

To start the stack again:

```bash
docker compose up -d
```
