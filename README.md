# Go Web Scraper

This project is a web scraper built using Go. It scrapes data from multiple URLs and sends alerts through SMS using the Twilio API whenever it finds new data. The scraped data is stored in a SQLite database.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.14 or higher)
- SQLite
- Twilio account (for SMS alerts)

### Environment Variables

The project uses the following environment variables:

- `TWILIO_ACCOUNT_SID`: Your Twilio Account SID
- `TWILIO_SID`: Your Twilio SID
- `TWILIO_SECRET`: Your Twilio Secret
- `RECIPIENT_PHONE_NUMBER`: The phone number to which alerts are sent
- `TWILIO_PHONE_NUMBER`: The phone number from which alerts are sent
- `SQLITE_DB`: The path to your SQLite database
- `SELECT_QUERY`: Query used to scan previously viewed urls


You can set these in your shell, or directly in your code for testing purposes. For production use, it's recommended to use a secure method for setting your environment variables.

### Running the Scraper

Navigate to the directory containing the project and run:

```bash
go run main.go
```

### Project Structure

- main.go: The main file that runs the scraper and sets up logs.
- pkg/: This directory contains the utility packages used by the main file.
    - db.go: This file handles interaction with the SQLite database.
    - log.go: Provides functionality for opening and managing log files.
    - sms.go: This file manages sending SMS alerts via the Twilio API.
    - scrape.go: This file contains functions to scrape data from URLs.
