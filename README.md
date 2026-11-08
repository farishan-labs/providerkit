# providerkit

Small, dependency-free Go helpers shared by `forager` and `provider-intel`
(and any other labs tool dealing with LLM/API providers).

## Functions

- `Slugify(s)` — stable lowercase slug.
- `Prettify(slug)` — slug → display name (with known-provider map).
- `HasTag(tags, tag)` — case-insensitive membership.
- `SplitModelID(id)` — split "provider/model".
- `BuildURL(base, path)` — join base + path with single slash.
- `Redact(s, secret)` — replace secret with `***`.
- `JoinTags` / `SplitTags` — comma (de)serialization.

## Usage

Both consumers reference it via a local `replace` in their `go.mod`:

```
require github.com/farishan/providerkit v0.0.0
replace github.com/farishan/providerkit => ../providerkit
```

A `go.work` at `~/dev/labs` ties the modules together for local development.