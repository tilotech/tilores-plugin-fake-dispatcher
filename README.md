# Fake Dispatcher Plugin

The fake dispatcher plugin implements the
[dispatcher plugin interface](https://pkg.go.dev/gitlab.com/tilotech/tilores-plugin-api/dispatcher#Dispatcher).

In contrast to the real (proprietary) TiloRes Core Dispatcher it takes a lot of
shortcuts and is only intended for testing the GraphQL API functionality.

* `Submit` adds at maximum 10 records, new records override existing ones
* `Entity` always returns an entity with the last up to 10 submitted records
* `Search` returns at max one entity with all records that exactly match parts of the request parameters