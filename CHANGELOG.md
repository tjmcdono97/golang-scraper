# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2023-06-21

### Added

- Initial release of the web scraper.
- `main.go`: Main function that sets up logging and runs the scraper.
- `pkg/db.go`: Handles interactions with the SQLite database.
- `pkg/log.go`: Manages the opening and logging to log files.
- `pkg/sms.go`: Manages sending SMS alerts using the Twilio API.
- `pkg/scrape.go`: Contains the scraping functionality.
- SQLite database for storing scraped data.
- Environment variable support for sensitive information.
- Twilio API integration for sending SMS alerts.
- Added a Dockerfile for convience on needing a gcc+ compiler.
- Added Dockerignore
- Updated README.md