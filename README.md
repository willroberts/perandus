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

1. Copy `settings.toml.example` to `settings.toml` and fill in the values for
   the item for which you want alerts.
2. Run the binary (`perandus.exe`).
3. Wait for the alert sound to play.
4. Copy the username from the alert to buy the item in-game.
   Or should the paste buffer automatically be filled?

## Code Organization

Code is grouped in the following packages:

* `alert`: Output mechanisms for notifications
* `client`: Rate-limited client for PoE API
* `filter`: Reads settings and evaluates item listings
* `models`: Models for stashes and items
* `util`: Helpers for retrieving data from other sources

## To Do

* Expand the power and flexibility of filters
  * Minimum and Maximum Price
  * Sockets and Links
* Add sound alerts
