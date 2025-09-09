package responses

type BloodStatResponse struct {
	ID 				 string         `json:"id"`
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string 				`json:"userId" binding:"required"`
}