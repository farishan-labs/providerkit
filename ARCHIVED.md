# ARCHIVED

This repository is superseded by [farishan-labs/providers](https://github.com/farishan-labs/providers),
which consolidates provider catalog, health, config generation, web UI, and MCP
into a single Go binary (powered by gidbams).

Functionality moved:
- provider-intel -> providers (internal/web, internal/mcp)
- provider-dashboard -> providers (internal/config, internal/web)
- forager -> providers (internal/health, internal/registry import)
- providerkit -> providers (internal/shared)

This repo is kept read-only for history. New work happens in providers.
