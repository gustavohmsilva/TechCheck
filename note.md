# notes

Package `tech` contains the domain stuff such as the service and related models
for the domain

Package `mysql` contains the persistence layer and can be injected into the
domain services

Package `api` contains the http transport layer and will use tech services to
retrieve and store

cmd/techcheck/Main is binding everything together
