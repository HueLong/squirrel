package validator

type AdShowPositionGet struct {
	ShowPosition string `form:"show_position" binding:"required" label:"展示标识"`
}

type AdListGet struct {
	ShowPosition string `form:"show_position"`
}
