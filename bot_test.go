package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunJobs(t *testing.T) {
	t.Run("affected tasks are calculated correctly", func(t *testing.T) {
		gitLog := `* Yanis: Align filters
		* Yanis: SMAR-2523: Add currency support for all charts
		* Yanis: SMAR-2517: Unable to login if previous attempt w..
		* Yanis: SMAR-2522: Fix counter
		* Yanis: Refactor home
		* Yanis: SMAR-2511: Add error boundary in my pnrs
		* Yanis: SMAR-2511 Add notification error boundary
		* Yanis: SMAR-2511: Put error boundaries in dashboard
		* Yanis: Refactor`
		expected := []string{"SMAR-2511", "SMAR-2517", "SMAR-2522", "SMAR-2523"}
		assert.Equal(t, expected , GetAffectedTasks(gitLog))
	})
}




