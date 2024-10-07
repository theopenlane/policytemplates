# Policy Templates

## Description

This repository contains audit compliance templates for multiple frameworks
including:

- SOC2
- NIST CSF
- NIST 800-53
- ... more to come ...

## Usage

There is a `cli` included to generate and validate standards based on a csv
input. This will parse the data in the provided format and output to a standard
which must conform to the [jsonschema](schema/jsonschema/frameworks.json)

### Schema

1. Run `task schema` to regenerate the jsonschema(s) based on the framework go
   structs

   ```
   task schema
   task: [schema] go run main.go schema
   12:55PM INF generating schema
   12:55PM INF writing schema to file
   12:55PM INF schema generated successfully file location=schema/jsonschema/frameworks.json
   ```

### Parse

1. Run `task parse` (or you can run `go run main.go parse` directly), which will
   bring up a cli prompt
1. Make a `framework` selection
   ```
   task parse
   task: [parse] go run main.go parse
   Use the arrow keys to navigate: ↓ ↑ → ←  and / toggles search
   Frameworks:
   👉 SOC2
       NIST CSF
       NIST 800-53

   Description:        2017 Trust Services Criteria for Security, Availability, Processing Integrity, Confidentiality, and Privacy (with Revised Points of Focus – 2022)
   ```
1. Make an `output` selection
   ```
   task parse
   task: [parse] go run main.go parse
   👉 NIST CSF
   Use the arrow keys to navigate: ↓ ↑ → ←  and / toggles search
   Output Format:
   👉 Save To File
   Standard Out - JSON
   ```
1. Result will either go to `stdout` or the files in `templates/standards`
   depending on the selection
   ```
   2:23PM INF parsing compliance standards format=file framework=nist-csf
   2:23PM INF validating standards against schema
   2:23PM INF standards saved to file filename=templates/standards/nist-csf-1.1.json
   ```
