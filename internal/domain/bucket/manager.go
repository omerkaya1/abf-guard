package bucket

import "github.com/omerkaya1/abf-guard/internal/domain/interfaces"

// Manager .
// One ring to rule em all!
type Manager struct {
	store map[string]interfaces.Bucket
}
