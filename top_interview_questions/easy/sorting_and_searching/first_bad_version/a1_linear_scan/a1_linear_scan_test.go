package a1_linear_scan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenOneVersion(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{BAD}
	n := 1

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 1
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenTwoGoodVersions(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, GOOD, BAD}
	n := 3

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 3
	assert.Equal(t, expectedVersion, firstBadVersion)
}
