Dealerships CLI
===============

A small command-line tool written in Go to manage dealership records stored as CSV files per region.

This repository provides simple commands to list, view, create, edit, and show status for dealerships in a given region.

Quick plan
----------
- Provide build & run instructions
- Document available commands and interactive behavior
- Describe the CSV storage format and sample row
- Add notes about regions and troubleshooting

Checklist
---------
- [x] Build instructions
- [x] Usage examples for each command
- [x] CSV format description
- [x] Notes and possible improvements

Project layout
--------------

- `main.go` — CLI entrypoint that parses flags and dispatches commands
- `dealership.go` — Dealership data model and interactive command implementations
- `storage.go` — CSV load/save helpers (per-region files under `data/`)
- `data/` — CSV files for regions (e.g. `tehran.csv`, `fars.csv`)

Building
--------
Make sure you have Go installed (tested with Go 1.16+). From the project root:

```bash
go build -o dealership-cli .
```

Or run directly without building:

```bash
go run ./...
```

Running (usage)
---------------
The CLI accepts two flags:

- `-command` — which command to run (one of `list`, `get`, `create`, `edit`, `status`)
- `-region` — the region name (maps to a CSV file `data/<region>.csv`)

Basic usage example:

```bash
./dealership-cli -command=list -region=tehran
# or with go run
go run main.go -command=list -region=tehran
```

Available commands
------------------

- `list` — prints a concise list of dealerships in the region.
- `get` — interactively prompts for a dealership ID and then prints detailed info.
- `create` — interactively prompts for name, phone, address, date, and employee count; saves a new record.
- `edit` — interactively prompts for an ID then allows editing each field (press Enter to keep current value).
- `status` — prints summary stats: number of dealerships and total employees.

Notes about interactive commands
--------------------------------
`get`, `create`, and `edit` read additional input from stdin. For example:

```bash
./dealership-cli -command=get -region=tehran
# then type the ID and press Enter when prompted
```

CSV storage format
------------------
Each region is stored as `data/<region>.csv`. Each row has the following columns (in order):

1. ID (int)
2. Name (string)
3. Address (string)
4. Phone (string)
5. MembershipDate (string)
6. EmployeeCount (int)

Example row:

```
1,AutoPlus,123 Main St,+98-21-12345678,2020-01-15,12
```

Important details and behavior
------------------------------
- Regions correspond to CSV filenames. Use e.g. `-region=tehran` to read/write `data/tehran.csv`.
- If a region file does not exist, listing returns empty results and creating a new dealership will create the file when saved.
- IDs are assigned automatically when creating new records (next-highest integer).
- The project currently stores data in plain CSV files and does not perform concurrency control.

Troubleshooting
---------------
- If you see "Usage: ./dealership-cli -command=<cmd> -region=<region>", make sure both flags are provided.
- If CSV files are not tracked in Git, check `.gitignore` (this repo ignores `*.csv` by default).
- Input validation is minimal; non-numeric ID/employee inputs will yield an error.

Extending the project
---------------------
Ideas for future work:

- Add automated tests for CSV parsing and command logic.
- Add non-interactive flags to pass values for `create`/`edit` to support scripting.
- Use a proper storage backend (SQLite, JSON, or a small REST API) instead of CSV for robustness.
- Improve validation and user-friendly error messages.

Contributing
------------
Feel free to open issues or submit PRs. Keep changes small and focused; include tests where appropriate.

License
-------
This project is provided without a license file. Add a LICENSE if you want to set terms.

