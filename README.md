# Perandus

Real-time item alerts for PoE item listings, without indexing.

## Packages

* `client`: Rate-limited client for PoE API
* `filter`: Reads settings and discards non-matching items
* `items`: Models for PoE items
* `util`: Helpers for retrieving data from other sources

## Flow

1. Get latest change ID from poe.ninja
2. Create a rate-limited stash client for pathofexile.com
   Client should embed rate-limited code
3. Begin stash polling, starting with change ID from #1 and automatically updating
	 Client should embed current change ID.
4. Parse stashes and items
5. Compare each item to the configured filter
6. Alert on match
