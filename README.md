# Installation

---

## Description

The resources in this repo deploys a service to Kubernetes using Helm.

There are two componentes:
- MySQL: We are going to use this service as database.
- SetGet: An API service that has two endpoints:
  - `/get`: Returns one of the previously persisted string values (a random one).
  - `/set`: Accepts a body value and persists it, the body looks like:
    - ```json
      {
        "item": "value to insert"
      }
      ```
      
---



---


Jose Ramon Mañes