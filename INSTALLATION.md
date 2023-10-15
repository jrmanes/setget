# Installation

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
# then run the following command, it will build the container and push it to your registry
make docker_all
```

*Remember to update image reference in the Helm values in `infra/setget/values.yaml`*

---


Jose Ramon Mañes