package a2_binary_search

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

func TestWhenTwoVersionsAreGB(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, BAD}
	n := 2

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 2
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtLeftSideWithOddVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, BAD, BAD, BAD, BAD}
	n := 5

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 2
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtRightSideWithOddVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, GOOD, GOOD, BAD, BAD}
	n := 5

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 4
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtLeftSideWithEvenVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, BAD, BAD, BAD, BAD, BAD}
	n := 6

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 2
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtRightSideWithEvenVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, GOOD, GOOD, GOOD, BAD, BAD}
	n := 6

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 5
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtHeadWithOddVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{BAD, BAD, BAD, BAD, BAD}
	n := 5

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 1
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtTailWithOddVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, GOOD, GOOD, GOOD, BAD}
	n := 5

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 5
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtHeadWithEvenVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{BAD, BAD, BAD, BAD, BAD, BAD}
	n := 6

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 1
	assert.Equal(t, expectedVersion, firstBadVersion)
}

func TestWhenFirstBadVersionAtTailWithEvenVersionList(t *testing.T) {
	// Arrange
	versionQualityList = []VersionQuality{GOOD, GOOD, GOOD, GOOD, GOOD, BAD}
	n := 6

	// Act
	firstBadVersion := firstBadVersion(n)

	// Assert
	expectedVersion := 6
	assert.Equal(t, expectedVersion, firstBadVersion)
}
