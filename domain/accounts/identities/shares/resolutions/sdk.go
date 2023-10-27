package resolutions

import "time"

// Resolution represents the share's resolution
type Resolution interface {
	VoteDuration() time.Duration
	VoteAmountToPassResolution() uint
	MinimumParticipation() uint
}
