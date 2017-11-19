# Perandus

Real-time item alerts for PoE item listings, without indexing.

## Packages

* `client`: Rate-limited client for PoE API
* `filter`: Reads settings and discards non-matching items
* `items`: Models for PoE items
* `util`: Helpers for retrieving data from other sources

## To Do

1. Parse stashes and items
2. Compare each item to the configured filter
3. Alert on match
