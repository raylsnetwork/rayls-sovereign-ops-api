package testutil

// StubLogger is a no-op logger for unit tests.
type StubLogger struct{}

func (l *StubLogger) Debug(_ string, _ ...any) {}
func (l *StubLogger) Info(_ string, _ ...any)  {}
func (l *StubLogger) Warn(_ string, _ ...any)  {}
func (l *StubLogger) Error(_ string, _ ...any) {}

// RecordingLogger captures Warn/Error messages so a test can assert that a diagnostic actually
// fired. StubLogger discards everything, which cannot distinguish "warned" from "stayed silent".
type RecordingLogger struct {
	Warns  []string
	Errors []string
}

func (l *RecordingLogger) Debug(_ string, _ ...any) {}
func (l *RecordingLogger) Info(_ string, _ ...any)  {}

func (l *RecordingLogger) Warn(msg string, _ ...any) {
	l.Warns = append(l.Warns, msg)
}

func (l *RecordingLogger) Error(msg string, _ ...any) {
	l.Errors = append(l.Errors, msg)
}
