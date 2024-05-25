package contracts

type MembershipsUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
	GymId   string `uri:"gymId" binding:"required"`
}
