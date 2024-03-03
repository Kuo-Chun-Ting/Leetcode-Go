package a2_binary_search

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
	left := 1
	right := n

	for {
		if left > right {
			break
		}

		mid := (left + right) / 2

		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
