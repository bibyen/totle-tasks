# Health & Reflection Guide

The server implements the Connect RPC Health and Reflection protocols to support health monitoring and dynamic API discovery.

This guide assumes the server is running locally.

## Prerequisites

- **[Go](https://go.dev/doc/install)** – Language runtime and compiler.
- **[grpcurl](https://github.com/fullstorydev/grpcurl#installation)** – CLI for interacting with gRPC/Connect.

---

## Service Discovery (Reflection)

Use these to verify which services are registered and see their available methods.

**List all services:**

```bash
grpcurl -plaintext localhost:8080 list
```

---

## Health Checks

Verify if the server or specific internal services are "SERVING".

```bash
# Check overall server health:
grpcurl -plaintext -d '{"service": ""}' localhost:8080 grpc.health.v1.Health/Check

# Check GoalService:
grpcurl -plaintext -d '{"service": "totle_tasks.v1.GoalService"}' localhost:8080 grpc.health.v1.Health/Check

## Check BingoService:
grpcurl -plaintext -d '{"service": "totle_tasks.v1.BingoService"}' localhost:8080 grpc.health.v1.Health/Check
```
