package responses

type VitalResponse struct{
	ID					string         `json:"id"`
	PressureUp  float64        `json:"pressureUp" binding:"required"`
	PressureLow float64        `json:"pressureLow" binding:"required"`
	Temperature float64        `json:"temperature" binding:"required"`
	Height      float64        `json:"height" binding:"required"`
	Weight      float64        `json:"weight" binding:"required"`
	BMI         float64        `json:"bmi"`
	UserID      string         `json:"userId"`
}
