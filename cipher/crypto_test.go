package cipher

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"reflect"
	"strings"
	"testing"
)

func TestMessageEncryptDecrypt(t *testing.T) {

	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Error("Generate key error:", err)
	}

	msg := []byte("test message")
	ciphertext, err := MessageEncrypt(privateKey.PubKey(), &msg)
	if err != nil {
		t.Error("MessageEncrypt error:", err)
	}

	plaintext, err := MessageDecrypt(privateKey, ciphertext)
	if err != nil {
		t.Error("MessageDecrypt error:", err)
	}

	if !reflect.DeepEqual(msg, *plaintext) {
		t.Error("MessageDecrypt failed:", "\n", "original: ", fmt.Sprint(msg), "\n", "decrypted:", fmt.Sprint(*plaintext))
	}
}

func TestOutputAddress(t *testing.T) {
	{
		//2-of-3 multisig test
		testM := 2
		testN := 3
		testPublicKeys := "04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd,046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187,0411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e83"
		testAddress := "347N1Thc213QqfYCz3PZkjoJpNv5b14kBd"
		testRedeemScriptHex := "524104a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd41046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187410411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e8353ae"

		P2SHAddress, redeemScriptHex, _ := OutputAddress(testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			t.Error(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			t.Error(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
	{
		//7-of-7 multisig test
		testM := 7
		testN := 7
		testPublicKeys := "0446f1c8de232a065da428bf76e44b41f59a46620dec0aedfc9b5ab651e91f2051d610fddc78b8eba38a634bfe9a74bb015a88c52b9b844c74997035e08a695ce9,04704e19d4fc234a42d707d41053c87011f990b564949532d72cab009e136bd60d7d0602f925fce79da77c0dfef4a49c6f44bd0540faef548e37557d74b36da124,04b75a8cb10fd3f1785addbafdb41b409ecd6ffd50d5ad71d8a3cdc5503bcb35d3d13cdf23f6d0eb6ab88446276e2ba5b92d8786da7e5c0fb63aafb62f87443d28,04033a82ccb1291bbc27cf541c6c487c213f25db85c620ecb9cbb76ca461ef13db5a80b90c3ae7d2a5e47623cdf520a2586cac7e41f779103a71a1fe177189781e,045e3b4030be5fd9c4c40e7076bd49f022118d90ae9182de61f3a1adb2ff511c97e8a6a82a9292b01878a18c08b7cd658ebdf80e6ed3f26783b25ba1a52fa9e52d,04c93ceb8f4482e131addc58d3efa0b4967bb7c574de15786d55379cc4a43a61571518abe0f05ebf188bcce9580aa70b3f5b1024ca579819c8810ff79967de3f23,04a66f63d2941f0befcfba4b73495a7b99fc7ed28cb41e7934e1de82d852628766dc96ee1e196387a68e7fd8898862c2260f1f2557ac2147af07900695f15abd3f"
		testAddress := "3ErDPiDD7AsJDqKkayMA39iLJevTjDCjUa"
		testRedeemScriptHex := "57410446f1c8de232a065da428bf76e44b41f59a46620dec0aedfc9b5ab651e91f2051d610fddc78b8eba38a634bfe9a74bb015a88c52b9b844c74997035e08a695ce94104704e19d4fc234a42d707d41053c87011f990b564949532d72cab009e136bd60d7d0602f925fce79da77c0dfef4a49c6f44bd0540faef548e37557d74b36da1244104b75a8cb10fd3f1785addbafdb41b409ecd6ffd50d5ad71d8a3cdc5503bcb35d3d13cdf23f6d0eb6ab88446276e2ba5b92d8786da7e5c0fb63aafb62f87443d284104033a82ccb1291bbc27cf541c6c487c213f25db85c620ecb9cbb76ca461ef13db5a80b90c3ae7d2a5e47623cdf520a2586cac7e41f779103a71a1fe177189781e41045e3b4030be5fd9c4c40e7076bd49f022118d90ae9182de61f3a1adb2ff511c97e8a6a82a9292b01878a18c08b7cd658ebdf80e6ed3f26783b25ba1a52fa9e52d4104c93ceb8f4482e131addc58d3efa0b4967bb7c574de15786d55379cc4a43a61571518abe0f05ebf188bcce9580aa70b3f5b1024ca579819c8810ff79967de3f234104a66f63d2941f0befcfba4b73495a7b99fc7ed28cb41e7934e1de82d852628766dc96ee1e196387a68e7fd8898862c2260f1f2557ac2147af07900695f15abd3f57ae"

		P2SHAddress, redeemScriptHex, err := OutputAddress(testM, testN, testPublicKeys)

		if testAddress != P2SHAddress {
			t.Error(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			t.Error(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}

		if err != nil {
			errString := err.Error()
			if !strings.Contains(errString, "WARNING:") {
				t.Error(t, "Should has a warning in this test case")
			}
		} else {
			t.Error("Test case failed due to no WARNING message returns")
		}
	}
	{
		//5-of-7 multisig test
		testM := 5
		testN := 7
		testPublicKeys := "04c22e4293d1d462eef905e592ad4aff332aa52c3415b824cd85cf594258d92c836fe797187bc2459261e0597c4ef351c5d0c26f7a60165221e221a38e448ad08c,04bb28684dfe23852a7c276827dd448c955007e7ccbfacbf536e13f1097b30430ebec5af0bc001e50d3f0e796d52ba43e3c07337bfed2a842659d51632f2b21d28,048f8551173f8e7414ff0e144899b3f70accd957e6913f5cf877bd576f6c16f0aa67fb9b96e0df10562b4f7ba4060acd22f142329ff83f1d96e27f4e4394adeda2,04aa81def7dda6a4f40be2f3287ee3423f255b07965104a7888df075217c9ee5b3e9e2e70115d43bfecbff8062f8289f5cab3d0ebd96c9f55c85f6147ff3a5e949,04493aa5f89ec34184a235b2c9f608eade1634636f94f64b59419875e15cb86a6d8c708a9d5eda3304cb983b2325a57af881ed75f28179f5f263d7758039b68d89,04dc284f749208d7fec57937bc5e72187b064df7d29b7aa82cae273e9a1c91beae9c510e0fd632a3db272c67db04061ea761d1ed91fdb8ab07e354047c64ce405d,042fc7796f54dd482db20f1bcce584f930ae74d5f27fc8336e2701bd0243d681281810c57e079947ebdfdfc8860ed34b0ba32db82a85249adc7c64ab547d48af64"
		testAddress := "34wgSuG9qtaNEV4MGye9UJcffcFTxnmXSC"
		testRedeemScriptHex := "554104c22e4293d1d462eef905e592ad4aff332aa52c3415b824cd85cf594258d92c836fe797187bc2459261e0597c4ef351c5d0c26f7a60165221e221a38e448ad08c4104bb28684dfe23852a7c276827dd448c955007e7ccbfacbf536e13f1097b30430ebec5af0bc001e50d3f0e796d52ba43e3c07337bfed2a842659d51632f2b21d2841048f8551173f8e7414ff0e144899b3f70accd957e6913f5cf877bd576f6c16f0aa67fb9b96e0df10562b4f7ba4060acd22f142329ff83f1d96e27f4e4394adeda24104aa81def7dda6a4f40be2f3287ee3423f255b07965104a7888df075217c9ee5b3e9e2e70115d43bfecbff8062f8289f5cab3d0ebd96c9f55c85f6147ff3a5e9494104493aa5f89ec34184a235b2c9f608eade1634636f94f64b59419875e15cb86a6d8c708a9d5eda3304cb983b2325a57af881ed75f28179f5f263d7758039b68d894104dc284f749208d7fec57937bc5e72187b064df7d29b7aa82cae273e9a1c91beae9c510e0fd632a3db272c67db04061ea761d1ed91fdb8ab07e354047c64ce405d41042fc7796f54dd482db20f1bcce584f930ae74d5f27fc8336e2701bd0243d681281810c57e079947ebdfdfc8860ed34b0ba32db82a85249adc7c64ab547d48af6457ae"

		P2SHAddress, redeemScriptHex, _ := OutputAddress(testM, testN, testPublicKeys)
		if testAddress != P2SHAddress {
			t.Error(t, "Generated P2SH address different from expected address.", testAddress, P2SHAddress)
		}
		if testRedeemScriptHex != redeemScriptHex {
			t.Error(t, "Generated P2SH address different from expected address.", testRedeemScriptHex, redeemScriptHex)
		}
	}
}