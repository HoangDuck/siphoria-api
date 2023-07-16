package res

type TotalReviews struct {
	Average       float32 `json:"average"`
	OneStarRate   float32 `json:"one_star_rate"`
	TwoStarRate   float32 `json:"two_star_rate"`
	ThreeStarRate float32 `json:"three_star_rate"`
	FourStarRate  float32 `json:"four_star_rate"`
	FiveStarRate  float32 `json:"five_star_rate"`
}
