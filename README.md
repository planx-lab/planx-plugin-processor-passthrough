# planx-plugin-processor-passthrough

A standalone Planx Connector plugin shipping a single **passthrough** processor
(one self-describing binary; the component returns each batch unchanged, 1:1).

## Positioning — why a separate repo?

- [`planx-plugin-processors`](../planx-plugin-processors) is the bundled built-in
  processors connector: eight components (`passthrough`, `filter`, `field-mapper`,
  `json-transform`, `json-validate`, `json-redact`, `regex-replace`, `text-template`)
  in one binary.
- This repo is a **standalone, minimal, single-purpose** connector shipping ONLY a
  passthrough processor — useful as a deployable reference, a minimal container
  image, or a template for a new standalone processor connector.
- Both are permitted by `planx-repo-layout.lock`; they are **not** duplicates by
  mistake.

See [`AGENTS.md`](./AGENTS.md) for plugin authoring rules and
[`planx-sdk-go`](../planx-sdk-go) for the SPI (`sdk.ProcessorSPI`).
