// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "fmt"

const _FeatureFlag_name = "unusedIDNASupportAllowAccountDeactivationAllowKeyRolloverResubmitMissingSCTsOnlyGoogleSafeBrowsingV4UseAIAIssuerURLAllowTLS02ChallengesGenerateOCSPEarlyIPv6FirstCountCertificatesExact"

var _FeatureFlag_index = [...]uint8{0, 6, 17, 41, 57, 80, 100, 115, 135, 152, 161, 183}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return fmt.Sprintf("FeatureFlag(%d)", i)
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
