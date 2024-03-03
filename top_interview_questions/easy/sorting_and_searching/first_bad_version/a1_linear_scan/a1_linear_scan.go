package a1_linear_scan

type VersionQuality string

const (
	GOOD VersionQuality = "Good"
	BAD  VersionQuality = "Bad"
)

var versionQualityList = []VersionQuality{}

func isBadVersion(version int) bool {
	versionQuality := versionQualityList[version-1]
	return versionQuality == BAD
}

func firstBadVersion(n int) int {
	for i := 1; i < n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}
