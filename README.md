# tarkov-desktop

Escape from Tarkov desktop utility

## About 💡

This utility runs alongside your Tarkov game. It scans and reads log files emitted from the game to collect data about queue times. These queue times are then submitted to our API for aggregation and analysis

## Usage 💻

> This is an extremely alpha utility

The current primary usage of this application is to run alongside your EFT game to collect data about queue times. It does not interact with the game in anyway and is totally safe to run. It simply looks at the lines written to your `application.log` files for the EFT game and parses the total queue time for a map. This data is then automatically submitted to our open source API. By running this application, you are directly helping to crowdsource queue data 🎉!

You can also link the QR code seen in the application to [the-hideout](https://play.google.com/store/apps/details?id=com.austinhodak.thehideout&hl=en_US&gl=US) (Android app) and get real-time notifications on your phone when you are about to load into a raid!

## How it Works 🔨

This tool works by doing the following:

- Scans your local tarkov files to detect where the game is installed
- Grab the latest log file folder from your installation directory
- "Watches" the `application.log` for newly emmited events (like `tail -f` in bash)
- When the scanner sees relevant events for queue times, it parses the event and submits the data to our API

## Running 🏃

The binary to run this application is located in the `release/` directory for convenience

## Development ⚙️

To build and develop this application locally do the following:

- [Install wails](https://wails.io/docs/gettingstarted/installation)
- Run `wails dev` to start the application
- Run `wails build` to build the application for production
