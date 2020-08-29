package nk

type Theme int

const (
	THEME_WHITE Theme = iota
	THEME_DARK
	THEME_RED
	THEME_BLUE
)

func (ctx *Context) SetStyle(theme Theme) {
	table := make([]Color, ColorCount)

	switch theme {
	case THEME_WHITE:
		table[ColorText] = NkRgba(70, 70, 70, 255)
		table[ColorWindow] = NkRgba(175, 175, 175, 255)
		table[ColorHeader] = NkRgba(175, 175, 175, 255)
		table[ColorBorder] = NkRgba(0, 0, 0, 255)
		table[ColorButton] = NkRgba(185, 185, 185, 255)
		table[ColorButtonHover] = NkRgba(170, 170, 170, 255)
		table[ColorButtonActive] = NkRgba(160, 160, 160, 255)
		table[ColorToggle] = NkRgba(150, 150, 150, 255)
		table[ColorToggleHover] = NkRgba(120, 120, 120, 255)
		table[ColorToggleCursor] = NkRgba(175, 175, 175, 255)
		table[ColorSelect] = NkRgba(190, 190, 190, 255)
		table[ColorSelectActive] = NkRgba(175, 175, 175, 255)
		table[ColorSlider] = NkRgba(190, 190, 190, 255)
		table[ColorSliderCursor] = NkRgba(80, 80, 80, 255)
		table[ColorSliderCursorHover] = NkRgba(70, 70, 70, 255)
		table[ColorSliderCursorActive] = NkRgba(60, 60, 60, 255)
		table[ColorProperty] = NkRgba(175, 175, 175, 255)
		table[ColorEdit] = NkRgba(150, 150, 150, 255)
		table[ColorEditCursor] = NkRgba(0, 0, 0, 255)
		table[ColorCombo] = NkRgba(175, 175, 175, 255)
		table[ColorChart] = NkRgba(160, 160, 160, 255)
		table[ColorChartColor] = NkRgba(45, 45, 45, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(180, 180, 180, 255)
		table[ColorScrollbarCursor] = NkRgba(140, 140, 140, 255)
		table[ColorScrollbarCursorHover] = NkRgba(150, 150, 150, 255)
		table[ColorScrollbarCursorActive] = NkRgba(160, 160, 160, 255)
		table[ColorTabHeader] = NkRgba(180, 180, 180, 255)
		NkStyleFromTable(ctx, table)
	case THEME_DARK:
		table[ColorText] = NkRgba(210, 210, 210, 255)
		table[ColorWindow] = NkRgba(57, 67, 71, 215)
		table[ColorHeader] = NkRgba(51, 51, 56, 220)
		table[ColorBorder] = NkRgba(46, 46, 46, 255)
		table[ColorButton] = NkRgba(48, 83, 111, 255)
		table[ColorButtonHover] = NkRgba(58, 93, 121, 255)
		table[ColorButtonActive] = NkRgba(63, 98, 126, 255)
		table[ColorToggle] = NkRgba(50, 58, 61, 255)
		table[ColorToggleHover] = NkRgba(45, 53, 56, 255)
		table[ColorToggleCursor] = NkRgba(48, 83, 111, 255)
		table[ColorSelect] = NkRgba(57, 67, 61, 255)
		table[ColorSelectActive] = NkRgba(48, 83, 111, 255)
		table[ColorSlider] = NkRgba(50, 58, 61, 255)
		table[ColorSliderCursor] = NkRgba(48, 83, 111, 245)
		table[ColorSliderCursorHover] = NkRgba(53, 88, 116, 255)
		table[ColorSliderCursorActive] = NkRgba(58, 93, 121, 255)
		table[ColorProperty] = NkRgba(50, 58, 61, 255)
		table[ColorEdit] = NkRgba(50, 58, 61, 225)
		table[ColorEditCursor] = NkRgba(210, 210, 210, 255)
		table[ColorCombo] = NkRgba(50, 58, 61, 255)
		table[ColorChart] = NkRgba(50, 58, 61, 255)
		table[ColorChartColor] = NkRgba(48, 83, 111, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(50, 58, 61, 255)
		table[ColorScrollbarCursor] = NkRgba(48, 83, 111, 255)
		table[ColorScrollbarCursorHover] = NkRgba(53, 88, 116, 255)
		table[ColorScrollbarCursorActive] = NkRgba(58, 93, 121, 255)
		table[ColorTabHeader] = NkRgba(48, 83, 111, 255)
		NkStyleFromTable(ctx, table)
	case THEME_RED:
		table[ColorText] = NkRgba(190, 190, 190, 255)
		table[ColorWindow] = NkRgba(30, 33, 40, 215)
		table[ColorHeader] = NkRgba(181, 45, 69, 220)
		table[ColorBorder] = NkRgba(51, 55, 67, 255)
		table[ColorButton] = NkRgba(181, 45, 69, 255)
		table[ColorButtonHover] = NkRgba(190, 50, 70, 255)
		table[ColorButtonActive] = NkRgba(195, 55, 75, 255)
		table[ColorToggle] = NkRgba(51, 55, 67, 255)
		table[ColorToggleHover] = NkRgba(45, 60, 60, 255)
		table[ColorToggleCursor] = NkRgba(181, 45, 69, 255)
		table[ColorSelect] = NkRgba(51, 55, 67, 255)
		table[ColorSelectActive] = NkRgba(181, 45, 69, 255)
		table[ColorSlider] = NkRgba(51, 55, 67, 255)
		table[ColorSliderCursor] = NkRgba(181, 45, 69, 255)
		table[ColorSliderCursorHover] = NkRgba(186, 50, 74, 255)
		table[ColorSliderCursorActive] = NkRgba(191, 55, 79, 255)
		table[ColorProperty] = NkRgba(51, 55, 67, 255)
		table[ColorEdit] = NkRgba(51, 55, 67, 225)
		table[ColorEditCursor] = NkRgba(190, 190, 190, 255)
		table[ColorCombo] = NkRgba(51, 55, 67, 255)
		table[ColorChart] = NkRgba(51, 55, 67, 255)
		table[ColorChartColor] = NkRgba(170, 40, 60, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(30, 33, 40, 255)
		table[ColorScrollbarCursor] = NkRgba(64, 84, 95, 255)
		table[ColorScrollbarCursorHover] = NkRgba(70, 90, 100, 255)
		table[ColorScrollbarCursorActive] = NkRgba(75, 95, 105, 255)
		table[ColorTabHeader] = NkRgba(181, 45, 69, 220)
		NkStyleFromTable(ctx, table)
	case THEME_BLUE:
		table[ColorText] = NkRgba(20, 20, 20, 255)
		table[ColorWindow] = NkRgba(202, 212, 214, 215)
		table[ColorHeader] = NkRgba(137, 182, 224, 220)
		table[ColorBorder] = NkRgba(140, 159, 173, 255)
		table[ColorButton] = NkRgba(137, 182, 224, 255)
		table[ColorButtonHover] = NkRgba(142, 187, 229, 255)
		table[ColorButtonActive] = NkRgba(147, 192, 234, 255)
		table[ColorToggle] = NkRgba(177, 210, 210, 255)
		table[ColorToggleHover] = NkRgba(182, 215, 215, 255)
		table[ColorToggleCursor] = NkRgba(137, 182, 224, 255)
		table[ColorSelect] = NkRgba(177, 210, 210, 255)
		table[ColorSelectActive] = NkRgba(137, 182, 224, 255)
		table[ColorSlider] = NkRgba(177, 210, 210, 255)
		table[ColorSliderCursor] = NkRgba(137, 182, 224, 245)
		table[ColorSliderCursorHover] = NkRgba(142, 188, 229, 255)
		table[ColorSliderCursorActive] = NkRgba(147, 193, 234, 255)
		table[ColorProperty] = NkRgba(210, 210, 210, 255)
		table[ColorEdit] = NkRgba(210, 210, 210, 225)
		table[ColorEditCursor] = NkRgba(20, 20, 20, 255)
		table[ColorCombo] = NkRgba(210, 210, 210, 255)
		table[ColorChart] = NkRgba(210, 210, 210, 255)
		table[ColorChartColor] = NkRgba(137, 182, 224, 255)
		table[ColorChartColorHighlight] = NkRgba(255, 0, 0, 255)
		table[ColorScrollbar] = NkRgba(190, 200, 200, 255)
		table[ColorScrollbarCursor] = NkRgba(64, 84, 95, 255)
		table[ColorScrollbarCursorHover] = NkRgba(70, 90, 100, 255)
		table[ColorScrollbarCursorActive] = NkRgba(75, 95, 105, 255)
		table[ColorTabHeader] = NkRgba(156, 193, 220, 255)
		NkStyleFromTable(ctx, table)
	}
}
