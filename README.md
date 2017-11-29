# Perandus

Real-time item alerts for PoE item listings, without indexing.

## What is this?

This is a tool for finding specific items at a specific price in Path of Exile.
When an item is listed which matches your filter criteria, you'll receive an
alert and will be able to message the seller.

## Why use this instead of poe.trade?

During peak times such as league starts, poe.trade can have significant delays
for real-time alerts. This tool will give you notifications several seconds
before poe.trade, allowing you to purchase the item before poe.trade users see
it. Even outside of peak player counts, this tool is consistently 2 or 3
seconds faster than poe.trade.

## How do I use it?

1. Copy `settings.toml.example` to `settings.toml` and fill in the fields.
2. Build and run the docker image with `make build && make run`.
3. Matching items will show on the command-line.

## Code Organization

Code is grouped in the following packages:

* `client`: Rate-limited client for the PoE API
* `filter`: Stream processor for items
* `models`: Models for stashes and items
* `util`: Helpers for retrieving data from other sources

## To Do

* Create an API package for the unfiltered stream
  * Send items over a WebSocket
* Implement filtering in JavaScript
* Expand the power and flexibility of filters
  * Minimum and Maximum Price
  * Sockets and Links
