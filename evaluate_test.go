package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBasic(t *testing.T) {
	res, err := Evaluate(`x eq "abc" and compareVersion(appVersion >= 1.2.3)`, map[string]interface{}{
		"x":          "abc",
		"appVersion": "1.3",
	})
	require.NoError(t, err)
	require.True(t, res)

	res, err = Evaluate(`x eq abc`, map[string]interface{}{
		"x": "abc",
	})
	require.Error(t, err)
	require.False(t, res)
}
