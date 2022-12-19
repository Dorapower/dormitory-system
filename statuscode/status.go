package statuscode

const (
	StatusSuccess = 200

	// Public Module

	StatusInvalidRequest = 510001 // Request is invalid
	StatusNoPermission   = 510002 // No permission
	StatusTooManyRequest = 510003 // Too many request
	StatusServerError    = 510004 // Server error

	// Authorize Module

	StatusInvalidToken  = 511001 // Token is invalid
	StatusExpiredToken  = 511002 // Token is expired
	StatusNoToken       = 511003 // Token is not found
	StatusWrongPassword = 511004 // Wrong password
	// User Module

	// Room Module
	StatusNoRoom = 512001 // Room is not found

	// Team Module
	StatusCreateTeamFailed  = 513001 // Create team failed
	StatusJoinTeamFailed    = 513002 // Join team failed
	StatusInvalidInviteCode = 513003 // Invite code is invalid
	StatusMisMatchedGender  = 513004 // Gender Mismatches
	StatusTeamIsFull        = 513005 // Team is full
	// Order Module
	StatusCreateOrderFailed = 514001 // Create order failed
	// Sys Module
)
