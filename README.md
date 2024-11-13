# Ver

**Ver** is a versioning tool that helps you manage version numbers for your projects. It provides simple commands to initialize, increment, and retrieve version numbers.

## Description

This tool maintains a version file in the format `year.dayOfYear.micro`. The `year` is the current UTC year minus 2000, `dayOfYear` is the UTC day of the year, and `micro` is an incrementing number for changes made on the same day.

For example, the version number for the date `2024-05-13 14:30:00 +0000 UTC` would be `24.134.1`.

## Installation

You can install the **ver** using the command:

```sh
go install github.com/av1ppp/ver/cmd/ver@latest
```

Ensure you have Go installed and properly configured in your environment.

## Usage

Run the `ver` command with one of the following actions:

- `init` - Initialize **ver** in the current directory by creating a `.version` file.
- `incr` - Increment the version number in the `.version` file.
- `get` - Get the current version number from the `.version` file.
- `help` - Print the help message.

## License

This project is licensed under the MIT License.
