package services

import (
	"errors"
	"github.com/nikhilryan/go-featuristic/internal/models"
	"hash/fnv"
)

func determineABTestVariant(flag *models.FeatureFlag, userID, targetGroup string) (string, error) {
	if flag.ABTestType != "A/B" {
		return "", errors.New("feature flag is not an A/B test")
	}

	if flag.TargetGroup != "" && flag.TargetGroup != targetGroup {
		return "", errors.New("user does not belong to the target group for this feature flag")
	}

	hash := fnv.New32a()
	_, err := hash.Write([]byte(userID))
	if err != nil {
		return "", err
	}
	userHash := hash.Sum32()

	if flag.TargetGroupB != "" && flag.TargetGroupB == targetGroup {
		return flag.ABTestValue, nil
	}

	if userHash%2 == 0 {
		return flag.Value, nil
	} else {
		return flag.ABTestValue, nil
	}
}
