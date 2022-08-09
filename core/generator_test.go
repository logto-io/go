package core

import (
	"fmt"
	"testing"
)

func TestGenerateCodeVerifierShouldGenerateRandomString(t *testing.T) {
	codeVerifier1 := GenerateCodeVerifier()
	codeVerifier2 := GenerateCodeVerifier()
	fmt.Println(codeVerifier1)
	if codeVerifier1 == codeVerifier2 {
		t.Fatalf("GenerateCodeVerifier should generate random string")
	}
}

func TestGenerateCodeVerifierShouldGenerateStringsLessThan128Characters(t *testing.T) {
	codeVerifier := GenerateCodeVerifier()
	if len(codeVerifier) > 128 {
		t.Fatalf("GenerateCodeVerifier should generate string less than 128 characters")
	}
}

func TestGenerateCodeChallengeShouldGenerateDifferentStringByDifferentCodeVerifier(t *testing.T) {
	codeVerifier1 := GenerateCodeVerifier()
	codeChallenge1 := GenerateCodeChallenge(codeVerifier1)

	codeVerifier2 := GenerateCodeVerifier()
	codeChallenge2 := GenerateCodeChallenge(codeVerifier2)

	if codeChallenge1 == codeChallenge2 {
		t.Fatalf("GenerateCodeChallenge should generate different string by different code verifier")
	}
}

func TestGenerateCodeChallengeShouldGenerateCorrectString(t *testing.T) {
	codeVerifierAndChallengeMap := map[string]string{
		"tO6MabnMFRAatnlMa1DdSstypzzkgalL1-k8Hr_GdfTj-VXGiEACqAkSkDhFuAuD8FOU8lMishaXjt29Xt2Oww": "0K3SLeGlNNzFswYJjcVzcN4C76m_8NZORxFJLBJWGwg",
		"ipK7uh7F41nJyYY4RZQzEwBwBTd-BlXSO4W8q0tK5VA":                                            "C51JGVPSnuLTTumLt6X5w2JAL_kBaeqHON3KPIviYaU",
		"Á": "p3yvZiKYauPicLIDZ0W1peDz4Z9KFC-9uxtDfoO1KOQ",
		"🚀": "67wLKHDrMj8rbP-lxJPO74GufrNq_HPU4DZzAWMdrsU",
	}
	for codeVerifier, codeChallenge := range codeVerifierAndChallengeMap {
		if GenerateCodeChallenge(codeVerifier) != codeChallenge {
			t.Fatalf("GenerateCodeChallenge should generate correct string")
		}
	}
}

func TestGenerateCodeChallengeShouldGenerateTheSameStringWithTheSameCodeVerifier(t *testing.T) {
	codeVerifier := GenerateCodeVerifier()
	codeChallenge1 := GenerateCodeChallenge(codeVerifier)
	codeChallenge2 := GenerateCodeChallenge(codeVerifier)
	if codeChallenge1 != codeChallenge2 {
		t.Fatalf("GenerateCodeChallenge should generate the same string with the same code verifier")
	}
}
