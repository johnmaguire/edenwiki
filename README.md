# edenwiki

A Git-backed Markdown wiki with a web frontend.

![Demo Screenshot](../assets/screenshot.png)

**NOTE**: EdenWiki is incomplete software. It can read and write pages but does not offer user authentication, page editing, a UI for edit history, etc. No guarantees are made on the consistency or durability of data.

## Running

1. To start the API server, run `make bin && ./edenwiki`
2. To start the web UI, run `cd web && npm start`
3. Navigate to https://localhost:3000
