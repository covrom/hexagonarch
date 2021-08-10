# hexagonarch
Example of hexagonal (clean) architecture template with Golang.

Packages in `core` contain the core business logic and entities.

Packages in `app` contain the application layer (ports) for core business logic and control its behavior.

The `frame` folder contains packages (adapters) for processing incoming requests (see` handler`) and processing requests to the database (see `dbuser` and` dbspace`). SQL queries are currently not implemented.