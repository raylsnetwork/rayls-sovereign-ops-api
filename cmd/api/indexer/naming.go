package indexer

import "strings"

// subjectNamespacePrefix is the top-level NATS namespace for every subject
// published or consumed by ops-api.
const subjectNamespacePrefix = "ops."

// InstanceSubject inserts the configured instance name as the second segment
// of an "ops.*" subject when instance is non-empty:
//
//	("a", "ops.access_manager.events") -> "ops.a.access_manager.events"
//	("a", "ops.tokens.>")               -> "ops.a.tokens.>"
//	("",  "ops.tokens.>")               -> "ops.tokens.>"             (unchanged)
//
// Subjects that do not start with the "ops." namespace are returned unchanged
// so callers can safely wrap any string without checking first.
func InstanceSubject(instance, subject string) string {
	if instance == "" || !strings.HasPrefix(subject, subjectNamespacePrefix) {
		return subject
	}
	return subjectNamespacePrefix + instance + "." + subject[len(subjectNamespacePrefix):]
}

// InstanceJSName suffixes a JetStream stream or durable consumer name with the
// configured instance when non-empty:
//
//	("a", "AM_EVENTS")           -> "AM_EVENTS_a"
//	("a", "am_events_processor") -> "am_events_processor_a"
//	("",  "AM_EVENTS")           -> "AM_EVENTS"                       (unchanged)
//
// This keeps the bare base name in single-instance deployments (no migration
// of existing streams) and yields a deterministic, unique name when multiple
// instances share a NATS server.
func InstanceJSName(instance, base string) string {
	if instance == "" {
		return base
	}
	return base + "_" + instance
}
