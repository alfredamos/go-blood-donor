package requests

type BloodStatCreateRequest struct {
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string 				`json:"userId" binding:"required"`
}

type BloodStatUpdateRequest struct {
	ID 				 string         `json:"id"`
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string 				`json:"userId" binding:"required"`
}