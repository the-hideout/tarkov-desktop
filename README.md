# queue-scanner 🔍

Escape from Tarkov queue scanner utility

## About 💡

This utility runs alongside your Tarkov game. It scans and reads log files emitted from the game to collect data about queue times. These queue times are then submitted to our API for aggregation and analysis

## Usage 💻

This is a community tool that is run by anyone and everyone who wants to help crowdsource EFT data for queue times. By running this tool, you are helping to make queue data available to the community

## How it Works 🔨

This tool works by doing the following:

- Scans your local tarkov files to detect where the game is installed
- Grab the latest log file folder from your installation directory
- "Watches" the `application.log` for newly emmited events (like `tail -f` in bash)
- When the scanner sees relevant events for queue times, it parses the event and submits the data to our API
