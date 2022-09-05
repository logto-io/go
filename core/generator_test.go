package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCodeVerifierShouldGenerateRandomString(t *testing.T) {
	codeVerifier1 := GenerateCodeVerifier()
	codeVerifier2 := GenerateCodeVerifier()
	assert.NotEqual(t, codeVerifier1, codeVerifier2)
}

func TestGenerateCodeVerifierShouldGenerateStringsLessThan128Characters(t *testing.T) {
	assert.Greater(t, 128, len(GenerateCodeVerifier()))
}

func TestGenerateCodeChallengeShouldGenerateDifferentStringByDifferentCodeVerifier(t *testing.T) {
	codeVerifier1 := GenerateCodeVerifier()
	codeChallenge1 := GenerateCodeChallenge(codeVerifier1)

	codeVerifier2 := GenerateCodeVerifier()
	codeChallenge2 := GenerateCodeChallenge(codeVerifier2)

	assert.NotEqual(t, codeChallenge1, codeChallenge2)
}

func TestGenerateCodeChallengeShouldGenerateCorrectString(t *testing.T) {
	codeVerifierAndChallengeMap := map[string]string{
		"tO6MabnMFRAatnlMa1DdSstypzzkgalL1-k8Hr_GdfTj-VXGiEACqAkSkDhFuAuD8FOU8lMishaXjt29Xt2Oww": "0K3SLeGlNNzFswYJjcVzcN4C76m_8NZORxFJLBJWGwg",
		"ipK7uh7F41nJyYY4RZQzEwBwBTd-BlXSO4W8q0tK5VA":                                            "C51JGVPSnuLTTumLt6X5w2JAL_kBaeqHON3KPIviYaU",
		"Á": "p3yvZiKYauPicLIDZ0W1peDz4Z9KFC-9uxtDfoO1KOQ",
		"🚀": "67wLKHDrMj8rbP-lxJPO74GufrNq_HPU4DZzAWMdrsU",
	}
	for codeVerifier, codeChallenge := range codeVerifierAndChallengeMap {
		assert.Equal(t, codeChallenge, GenerateCodeChallenge(codeVerifier))
	}
}

func TestGenerateCodeChallengeShouldGenerateTheSameStringWithTheSameCodeVerifier(t *testing.T) {
	codeVerifier := GenerateCodeVerifier()
	codeChallenge1 := GenerateCodeChallenge(codeVerifier)
	codeChallenge2 := GenerateCodeChallenge(codeVerifier)
	assert.Equal(t, codeChallenge1, codeChallenge2)
}

func TestGenerateStateShouldGenerateRandomString(t *testing.T) {
	state1 := GenerateState()
	state2 := GenerateState()
	assert.NotEqual(t, state1, state2)
}
