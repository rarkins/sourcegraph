package graphqlbackend

import (
	_ "embed"
)

// MainSchema is the main raw graqhql schema.
//go:embed schema.graphql
var MainSchema string

// BatchesSchema is the Batch Changes raw graqhql schema.
//go:embed batches.graphql
var BatchesSchema string

// DotcomSchema is the Dotcom schema extension raw graqhql schema.
//go:embed dotcom.graphql
var DotcomSchema string
