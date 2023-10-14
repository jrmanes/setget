# Installation

---

## Description

The resources in this repo deploys a service to Kubernetes using Helm.

There are two componentes:
- MYSQL: We are going to use this service as database.
- SETGET: An API service that has two endpoints:
  - `/get`: Returns one of the previously persisted string values (choose a random one).
  - `/set`: Accepts a body value and persists it, the body looks like:
    - ```json
      {
        "item": "value to insert"
      }
      ```
      
---

## How to install it

Run the command:

```make
make helm_install
```

---

## How to uninstall it

Run the command:

```make
make helm_uninstall
```

---

## If you want to fill the DB

```bash
# port forward to your computer
kubectl -n setget port-forward service/setget 8080:8080

# run the loop to make some requests
for i in {1..100}; do
    item="random-$i"
    data="{\"item\":\"$item\"}"
    curl -d "$data" -H "Content-Type: application/json" -X POST http://localhost:8080/set
    sleep 0.01 # make some delay
done

```


## Build the container

If you want to build the container and push it to your own registry, you have to update the vlaue in the `Makefile`, the variable:
```
REGISTRY_NAME=<your-registry>
```

---


Jose Ramon Mañes