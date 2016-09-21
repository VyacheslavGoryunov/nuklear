// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 22 Sep 2016 01:34:53 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package nk

/*
#cgo CFLAGS: -DNK_INCLUDE_FIXED_TYPES=1
#include "nuklear.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Char type as declared in nk/nuklear.h:370
type Char byte

// Uchar type as declared in nk/nuklear.h:371
type Uchar byte

// Byte type as declared in nk/nuklear.h:372
type Byte byte

// Short type as declared in nk/nuklear.h:373
type Short int16

// Ushort type as declared in nk/nuklear.h:374
type Ushort uint16

// Int type as declared in nk/nuklear.h:375
type Int int32

// Uint type as declared in nk/nuklear.h:376
type Uint uint32

// Size type as declared in nk/nuklear.h:377
type Size uint

// Ptr type as declared in nk/nuklear.h:378
type Ptr uint

// Hash type as declared in nk/nuklear.h:380
type Hash uint32

// Flags type as declared in nk/nuklear.h:381
type Flags uint32

// Rune type as declared in nk/nuklear.h:382
type Rune uint32

// Glyph type as declared in nk/nuklear.h:422
type Glyph [4]byte

// Handle as declared in nk/nuklear.h:423
const sizeofHandle = unsafe.Sizeof(C.nk_handle{})

type Handle [sizeofHandle]byte

// PluginAlloc type as declared in nk/nuklear.h:442
type PluginAlloc func(arg0 Handle, old unsafe.Pointer, arg2 Size) unsafe.Pointer

// PluginFree type as declared in nk/nuklear.h:443
type PluginFree func(arg0 Handle, old unsafe.Pointer)

// PluginFilter type as declared in nk/nuklear.h:444
type PluginFilter func(arg0 []TextEdit, unicode Rune) int32

// PluginPaste type as declared in nk/nuklear.h:445
type PluginPaste func(arg0 Handle, arg1 []TextEdit)

// PluginCopy type as declared in nk/nuklear.h:446
type PluginCopy func(arg0 Handle, arg1 string, len int32)

// TextWidthF type as declared in nk/nuklear.h:1380
type TextWidthF func(arg0 Handle, h float32, arg2 string, len int32) float32

// QueryFontGlyphF type as declared in nk/nuklear.h:1381
type QueryFontGlyphF func(handle Handle, fontHeight float32, glyph []UserFontGlyph, codepoint Rune, nextCodepoint Rune)

// ConfigStackFloat as declared in nk/nuklear.h:2631
type ConfigStackFloat struct {
	Head           int32
	Elements       [32]ConfigStackFloatElement
	ref8c8d0ef5    *C.struct_nk_config_stack_float
	allocs8c8d0ef5 interface{}
}

// Clipboard as declared in nk/nuklear.h:1259
type Clipboard struct {
	Userdata       Handle
	Paste          PluginPaste
	Copy           PluginCopy
	ref564266a3    *C.struct_nk_clipboard
	allocs564266a3 interface{}
}

// StyleSlider as declared in nk/nuklear.h:2085
type StyleSlider struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	BorderColor    Color
	BarNormal      Color
	BarHover       Color
	BarActive      Color
	BarFilled      Color
	CursorNormal   StyleItem
	CursorHover    StyleItem
	CursorActive   StyleItem
	Border         float32
	Rounding       float32
	BarHeight      float32
	Padding        Vec2
	Spacing        Vec2
	CursorSize     Vec2
	ShowButtons    int32
	IncButton      StyleButton
	DecButton      StyleButton
	IncSymbol      SymbolType
	DecSymbol      SymbolType
	Userdata       Handle
	DrawBegin      *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd        *func(arg0 []CommandBuffer, arg1 Handle)
	refd2203cf4    *C.struct_nk_style_slider
	allocsd2203cf4 interface{}
}

// Context as declared in nk/nuklear.h:412
type Context struct {
	Input            Input
	Style            Style
	Memory           Buffer
	Clip             Clipboard
	LastWidgetState  Flags
	DeltaTimeSeconds float32
	ButtonBehavior   ButtonBehavior
	Stacks           ConfigurationStacks
	TextEdit         TextEdit
	Overlay          CommandBuffer
	Build            int32
	UsePool          int32
	Pool             Pool
	Begin            []Window
	End              []Window
	Active           []Window
	Current          []Window
	Freelist         []PageElement
	Count            uint32
	Seq              uint32
	refc5463edd      *C.struct_nk_context
	allocsc5463edd   interface{}
}

// ConfigStackFlags as declared in nk/nuklear.h:2633
type ConfigStackFlags struct {
	Head           int32
	Elements       [32]ConfigStackFlagsElement
	ref4e2d11da    *C.struct_nk_config_stack_flags
	allocs4e2d11da interface{}
}

// StyleWindow as declared in nk/nuklear.h:2344
type StyleWindow struct {
	Header                StyleWindowHeader
	FixedBackground       StyleItem
	Background            Color
	BorderColor           Color
	PopupBorderColor      Color
	ComboBorderColor      Color
	ContextualBorderColor Color
	MenuBorderColor       Color
	GroupBorderColor      Color
	TooltipBorderColor    Color
	Scaler                StyleItem
	Border                float32
	ComboBorder           float32
	ContextualBorder      float32
	MenuBorder            float32
	GroupBorder           float32
	TooltipBorder         float32
	PopupBorder           float32
	Rounding              float32
	Spacing               Vec2
	ScrollbarSize         Vec2
	MinSize               Vec2
	Padding               Vec2
	GroupPadding          Vec2
	PopupPadding          Vec2
	ComboPadding          Vec2
	ContextualPadding     Vec2
	MenuPadding           Vec2
	TooltipPadding        Vec2
	ref9603d52e           *C.struct_nk_style_window
	allocs9603d52e        interface{}
}

// StyleChart as declared in nk/nuklear.h:2248
type StyleChart struct {
	Background     StyleItem
	BorderColor    Color
	SelectedColor  Color
	Color          Color
	Border         float32
	Rounding       float32
	Padding        Vec2
	refb56abd3a    *C.struct_nk_style_chart
	allocsb56abd3a interface{}
}

// PropertyState as declared in nk/nuklear.h:2541
type PropertyState struct {
	Active         int32
	Prev           int32
	Buffer         [64]byte
	Length         int32
	Cursor         int32
	Name           Hash
	Seq            uint32
	Old            uint32
	State          int32
	ref491b08a7    *C.struct_nk_property_state
	allocs491b08a7 interface{}
}

// DrawVertexLayoutElement as declared in nk/nuklear.h:413
type DrawVertexLayoutElement C.struct_nk_draw_vertex_layout_element

// Buffer as declared in nk/nuklear.h:402
type Buffer struct {
	Marker         [2]BufferMarker
	Pool           Allocator
	Type           AllocationType
	Memory         Memory
	GrowFactor     float32
	Allocated      Size
	Needed         Size
	Calls          Size
	Size           Size
	ref28a0f983    *C.struct_nk_buffer
	allocs28a0f983 interface{}
}

// DrawNullTexture as declared in nk/nuklear.h:454
type DrawNullTexture struct {
	Texture        Handle
	Uv             Vec2
	refc45a53be    *C.struct_nk_draw_null_texture
	allocsc45a53be interface{}
}

// CommandPolygonFilled as declared in nk/nuklear.h:1696
type CommandPolygonFilled struct {
	Header         Command
	Color          Color
	PointCount     uint16
	Points         [1]Vec2i
	ref5893ac22    *C.struct_nk_command_polygon_filled
	allocs5893ac22 interface{}
}

// Input as declared in nk/nuklear.h:1802
type Input struct {
	Keyboard       Keyboard
	Mouse          Mouse
	ref4e173241    *C.struct_nk_input
	allocs4e173241 interface{}
}

// ConfigStackVec2 as declared in nk/nuklear.h:2632
type ConfigStackVec2 struct {
	Head           int32
	Elements       [16]ConfigStackVec2Element
	ref5007555e    *C.struct_nk_config_stack_vec2
	allocs5007555e interface{}
}

// ConfigStackStyleItemElement as declared in nk/nuklear.h:2622
type ConfigStackStyleItemElement struct {
	Address        []StyleItem
	OldValue       StyleItem
	ref3797a13d    *C.struct_nk_config_stack_style_item_element
	allocs3797a13d interface{}
}

// CommandRect as declared in nk/nuklear.h:1612
type CommandRect struct {
	Header         Command
	Rounding       uint16
	LineThickness  uint16
	X              int16
	Y              int16
	W              uint16
	H              uint16
	Color          Color
	ref213e4f25    *C.struct_nk_command_rect
	allocs213e4f25 interface{}
}

// ConfigStackFloatElement as declared in nk/nuklear.h:2623
type ConfigStackFloatElement struct {
	Address       []float32
	OldValue      float32
	ref48484a2    *C.struct_nk_config_stack_float_element
	allocs48484a2 interface{}
}

// Str as declared in nk/nuklear.h:1176
type Str struct {
	Buffer         Buffer
	Len            int32
	reff7fc1f53    *C.struct_nk_str
	allocsf7fc1f53 interface{}
}

// Key as declared in nk/nuklear.h:1791
type Key struct {
	Down           int32
	Clicked        uint32
	ref21c21703    *C.struct_nk_key
	allocs21c21703 interface{}
}

// CommandCurve as declared in nk/nuklear.h:1603
type CommandCurve struct {
	Header         Command
	LineThickness  uint16
	Begin          Vec2i
	End            Vec2i
	Ctrl           [2]Vec2i
	Color          Color
	ref86b4d23a    *C.struct_nk_command_curve
	allocs86b4d23a interface{}
}

// ConfigStackColorElement as declared in nk/nuklear.h:2626
type ConfigStackColorElement struct {
	Address        []Color
	OldValue       Color
	ref9a04ca79    *C.struct_nk_config_stack_color_element
	allocs9a04ca79 interface{}
}

// Style as declared in nk/nuklear.h:2380
type Style struct {
	Font             []UserFont
	Cursors          [7]*Cursor
	CursorActive     []Cursor
	CursorLast       []Cursor
	CursorVisible    int32
	Text             StyleText
	Button           StyleButton
	ContextualButton StyleButton
	MenuButton       StyleButton
	Option           StyleToggle
	Checkbox         StyleToggle
	Selectable       StyleSelectable
	Slider           StyleSlider
	Progress         StyleProgress
	Property         StyleProperty
	Edit             StyleEdit
	Chart            StyleChart
	Scrollh          StyleScrollbar
	Scrollv          StyleScrollbar
	Tab              StyleTab
	Combo            StyleCombo
	Window           StyleWindow
	refa582b8fc      *C.struct_nk_style
	allocsa582b8fc   interface{}
}

// StyleItem as declared in nk/nuklear.h:407
type StyleItem struct {
	Type           StyleItemType
	Data           StyleItemData
	ref26fd096e    *C.struct_nk_style_item
	allocs26fd096e interface{}
}

// Page as declared in nk/nuklear.h:2671
type Page struct {
	Size           uint32
	Next           []Page
	Win            [1]PageElement
	ref22a2ae6a    *C.struct_nk_page
	allocs22a2ae6a interface{}
}

// StyleEdit as declared in nk/nuklear.h:2183
type StyleEdit struct {
	Normal             StyleItem
	Hover              StyleItem
	Active             StyleItem
	BorderColor        Color
	Scrollbar          StyleScrollbar
	CursorNormal       Color
	CursorHover        Color
	CursorTextNormal   Color
	CursorTextHover    Color
	TextNormal         Color
	TextHover          Color
	TextActive         Color
	SelectedNormal     Color
	SelectedHover      Color
	SelectedTextNormal Color
	SelectedTextHover  Color
	Border             float32
	Rounding           float32
	CursorSize         float32
	ScrollbarSize      Vec2
	Padding            Vec2
	RowPadding         float32
	refb8d3f26a        *C.struct_nk_style_edit
	allocsb8d3f26a     interface{}
}

// StyleProperty as declared in nk/nuklear.h:2217
type StyleProperty struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	BorderColor    Color
	LabelNormal    Color
	LabelHover     Color
	LabelActive    Color
	SymLeft        SymbolType
	SymRight       SymbolType
	Border         float32
	Rounding       float32
	Padding        Vec2
	Edit           StyleEdit
	IncButton      StyleButton
	DecButton      StyleButton
	Userdata       Handle
	DrawBegin      *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd        *func(arg0 []CommandBuffer, arg1 Handle)
	ref86bbdfa4    *C.struct_nk_style_property
	allocs86bbdfa4 interface{}
}

// Mouse as declared in nk/nuklear.h:1780
type Mouse struct {
	Buttons        [3]MouseButton
	Pos            Vec2
	Prev           Vec2
	Delta          Vec2
	ScrollDelta    float32
	Grab           byte
	Grabbed        byte
	Ungrab         byte
	ref390ab67b    *C.struct_nk_mouse
	allocs390ab67b interface{}
}

// StyleButton as declared in nk/nuklear.h:1993
type StyleButton struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	BorderColor    Color
	TextBackground Color
	TextNormal     Color
	TextHover      Color
	TextActive     Color
	TextAlignment  Flags
	Border         float32
	Rounding       float32
	Padding        Vec2
	ImagePadding   Vec2
	TouchPadding   Vec2
	Userdata       Handle
	DrawBegin      *func(arg0 []CommandBuffer, userdata Handle)
	DrawEnd        *func(arg0 []CommandBuffer, userdata Handle)
	ref27e180ce    *C.struct_nk_style_button
	allocs27e180ce interface{}
}

// Allocator as declared in nk/nuklear.h:403
type Allocator C.struct_nk_allocator

// StyleCombo as declared in nk/nuklear.h:2261
type StyleCombo struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	BorderColor    Color
	LabelNormal    Color
	LabelHover     Color
	LabelActive    Color
	SymbolNormal   Color
	SymbolHover    Color
	SymbolActive   Color
	Button         StyleButton
	SymNormal      SymbolType
	SymHover       SymbolType
	SymActive      SymbolType
	Border         float32
	Rounding       float32
	ContentPadding Vec2
	ButtonPadding  Vec2
	Spacing        Vec2
	refe100a75a    *C.struct_nk_style_combo
	allocse100a75a interface{}
}

// TextUndoRecord as declared in nk/nuklear.h:1265
type TextUndoRecord struct {
	Where          int32
	InsertLength   int16
	DeleteLength   int16
	CharStorage    int16
	refe570b77c    *C.struct_nk_text_undo_record
	allocse570b77c interface{}
}

// ConfigStackUserFont as declared in nk/nuklear.h:2635
type ConfigStackUserFont struct {
	Head           int32
	Elements       [8]ConfigStackUserFontElement
	refa664861d    *C.struct_nk_config_stack_user_font
	allocsa664861d interface{}
}

// PageElement as declared in nk/nuklear.h:2665
type PageElement struct {
	Data           PageData
	Next           []PageElement
	Prev           []PageElement
	ref8c626785    *C.struct_nk_page_element
	allocs8c626785 interface{}
}

// CommandPolyline as declared in nk/nuklear.h:1703
type CommandPolyline struct {
	Header         Command
	Color          Color
	LineThickness  uint16
	PointCount     uint16
	Points         [1]Vec2i
	refdf010473    *C.struct_nk_command_polyline
	allocsdf010473 interface{}
}

// CommandTriangleFilled as declared in nk/nuklear.h:1648
type CommandTriangleFilled struct {
	Header         Command
	A              Vec2i
	B              Vec2i
	C              Vec2i
	Color          Color
	ref410c5769    *C.struct_nk_command_triangle_filled
	allocs410c5769 interface{}
}

// CommandTriangle as declared in nk/nuklear.h:1639
type CommandTriangle struct {
	Header         Command
	LineThickness  uint16
	A              Vec2i
	B              Vec2i
	C              Vec2i
	Color          Color
	refcfbe3d74    *C.struct_nk_command_triangle
	allocscfbe3d74 interface{}
}

// Pool as declared in nk/nuklear.h:2677
type Pool struct {
	Alloc          Allocator
	Type           AllocationType
	PageCount      uint32
	Pages          []Page
	Freelist       []PageElement
	Capacity       uint32
	Size           Size
	Cap            Size
	ref9939b1cc    *C.struct_nk_pool
	allocs9939b1cc interface{}
}

// ConfigStackUserFontElement as declared in nk/nuklear.h:2627
type ConfigStackUserFontElement struct {
	Address        [][]UserFont
	OldValue       []UserFont
	ref5572630c    *C.struct_nk_config_stack_user_font_element
	allocs5572630c interface{}
}

// Vec2 as declared in nk/nuklear.h:418
type Vec2 struct {
	X              float32
	Y              float32
	ref91a35839    *C.struct_nk_vec2
	allocs91a35839 interface{}
}

// Window as declared in nk/nuklear.h:2552
type Window struct {
	Seq                  uint32
	Name                 Hash
	NameString           [64]byte
	Flags                Flags
	Bounds               Rect
	Scrollbar            Scroll
	Buffer               CommandBuffer
	Layout               []Panel
	ScrollbarHidingTimer float32
	Property             PropertyState
	Popup                PopupState
	Edit                 EditState
	Scrolled             uint32
	Tables               []Table
	TableCount           uint16
	TableSize            uint16
	Next                 []Window
	Prev                 []Window
	Parent               []Window
	ref921ef0ac          *C.struct_nk_window
	allocs921ef0ac       interface{}
}

// UserFontGlyph as declared in nk/nuklear.h:1379
type UserFontGlyph C.struct_nk_user_font_glyph

// CommandImage as declared in nk/nuklear.h:1711
type CommandImage struct {
	Header         Command
	X              int16
	Y              int16
	W              uint16
	H              uint16
	Img            Image
	Col            Color
	ref14108bd2    *C.struct_nk_command_image
	allocs14108bd2 interface{}
}

// StyleSelectable as declared in nk/nuklear.h:2050
type StyleSelectable struct {
	Normal            StyleItem
	Hover             StyleItem
	Pressed           StyleItem
	NormalActive      StyleItem
	HoverActive       StyleItem
	PressedActive     StyleItem
	TextNormal        Color
	TextHover         Color
	TextPressed       Color
	TextNormalActive  Color
	TextHoverActive   Color
	TextPressedActive Color
	TextBackground    Color
	TextAlignment     Flags
	Rounding          float32
	Padding           Vec2
	TouchPadding      Vec2
	ImagePadding      Vec2
	Userdata          Handle
	DrawBegin         *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd           *func(arg0 []CommandBuffer, arg1 Handle)
	refdf1c4e20       *C.struct_nk_style_selectable
	allocsdf1c4e20    interface{}
}

// MenuState as declared in nk/nuklear.h:2469
type MenuState struct {
	X              float32
	Y              float32
	W              float32
	H              float32
	Offset         Scroll
	ref6455ac1d    *C.struct_nk_menu_state
	allocs6455ac1d interface{}
}

// ConfigStackFlagsElement as declared in nk/nuklear.h:2625
type ConfigStackFlagsElement struct {
	Address        []Flags
	OldValue       Flags
	ref3f8db22a    *C.struct_nk_config_stack_flags_element
	allocs3f8db22a interface{}
}

// StyleToggle as declared in nk/nuklear.h:2020
type StyleToggle struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	BorderColor    Color
	CursorNormal   StyleItem
	CursorHover    StyleItem
	TextNormal     Color
	TextHover      Color
	TextActive     Color
	TextBackground Color
	TextAlignment  Flags
	Padding        Vec2
	TouchPadding   Vec2
	Spacing        float32
	Border         float32
	Userdata       Handle
	DrawBegin      *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd        *func(arg0 []CommandBuffer, arg1 Handle)
	ref380d8178    *C.struct_nk_style_toggle
	allocs380d8178 interface{}
}

// ConfigStackButtonBehaviorElement as declared in nk/nuklear.h:2628
type ConfigStackButtonBehaviorElement struct {
	Address       []ButtonBehavior
	OldValue      ButtonBehavior
	ref485b597    *C.struct_nk_config_stack_button_behavior_element
	allocs485b597 interface{}
}

// BufferMarker as declared in nk/nuklear.h:1126
type BufferMarker struct {
	Active         int32
	Offset         Size
	refe53e6784    *C.struct_nk_buffer_marker
	allocse53e6784 interface{}
}

// TextUndoState as declared in nk/nuklear.h:1272
type TextUndoState struct {
	UndoRec        [99]TextUndoRecord
	UndoChar       [999]Rune
	UndoPoint      int16
	RedoPoint      int16
	UndoCharPoint  int16
	RedoCharPoint  int16
	ref56c782f6    *C.struct_nk_text_undo_state
	allocs56c782f6 interface{}
}

// ChartSlot as declared in nk/nuklear.h:2432
type ChartSlot struct {
	Type           ChartType
	Color          Color
	Highlight      Color
	Min            float32
	Max            float32
	Range          float32
	Count          int32
	Last           Vec2
	Index          int32
	refa88720a9    *C.struct_nk_chart_slot
	allocsa88720a9 interface{}
}

// CommandPolygon as declared in nk/nuklear.h:1688
type CommandPolygon struct {
	Header         Command
	Color          Color
	LineThickness  uint16
	PointCount     uint16
	Points         [1]Vec2i
	ref348e8888    *C.struct_nk_command_polygon
	allocs348e8888 interface{}
}

// Panel as declared in nk/nuklear.h:411
type Panel struct {
	Type           PanelType
	Flags          Flags
	Bounds         Rect
	Offset         []Scroll
	AtX            float32
	AtY            float32
	MaxX           float32
	FooterHeight   float32
	HeaderHeight   float32
	Border         float32
	HasScrolling   uint32
	Clip           Rect
	Menu           MenuState
	Row            RowLayout
	Chart          Chart
	PopupBuffer    PopupBuffer
	Buffer         []CommandBuffer
	Parent         []Panel
	ref3492d399    *C.struct_nk_panel
	allocs3492d399 interface{}
}

// Color as declared in nk/nuklear.h:416
type Color struct {
	R              Byte
	G              Byte
	B              Byte
	A              Byte
	reff069487f    *C.struct_nk_color
	allocsf069487f interface{}
}

// CommandText as declared in nk/nuklear.h:1719
type CommandText struct {
	Header         Command
	Font           []UserFont
	Background     Color
	Foreground     Color
	X              int16
	Y              int16
	W              uint16
	H              uint16
	Height         float32
	Length         int32
	String         [1]byte
	refad63db63    *C.struct_nk_command_text
	allocsad63db63 interface{}
}

// CommandCircle as declared in nk/nuklear.h:1656
type CommandCircle struct {
	Header         Command
	X              int16
	Y              int16
	LineThickness  uint16
	W              uint16
	H              uint16
	Color          Color
	ref476fb76b    *C.struct_nk_command_circle
	allocs476fb76b interface{}
}

// Rect as declared in nk/nuklear.h:420
type Rect struct {
	X              float32
	Y              float32
	W              float32
	H              float32
	ref817e2bcb    *C.struct_nk_rect
	allocs817e2bcb interface{}
}

// StyleText as declared in nk/nuklear.h:1988
type StyleText struct {
	Color         Color
	Padding       Vec2
	ref26d8bb7    *C.struct_nk_style_text
	allocs26d8bb7 interface{}
}

// Chart as declared in nk/nuklear.h:2442
type Chart struct {
	Slots          [4]ChartSlot
	Slot           int32
	X              float32
	Y              float32
	W              float32
	H              float32
	ref73692abc    *C.struct_nk_chart
	allocs73692abc interface{}
}

// PageData as declared in nk/nuklear.h:2660
const sizeofPageData = unsafe.Sizeof(C.union_nk_page_data{})

type PageData [sizeofPageData]byte

// DrawList as declared in nk/nuklear.h:409
type DrawList C.struct_nk_draw_list

// PopupState as declared in nk/nuklear.h:2517
type PopupState struct {
	Win            []Window
	Type           PanelType
	Name           Hash
	Active         int32
	ComboCount     uint32
	ConCount       uint32
	ConOld         uint32
	ActiveCon      uint32
	Header         Rect
	ref4fb276c7    *C.struct_nk_popup_state
	allocs4fb276c7 interface{}
}

// ConfigStackStyleItem as declared in nk/nuklear.h:2630
type ConfigStackStyleItem struct {
	Head           int32
	Elements       [16]ConfigStackStyleItemElement
	ref99f88018    *C.struct_nk_config_stack_style_item
	allocs99f88018 interface{}
}

// StyleTab as declared in nk/nuklear.h:2292
type StyleTab struct {
	Background         StyleItem
	BorderColor        Color
	Text               Color
	TabMaximizeButton  StyleButton
	TabMinimizeButton  StyleButton
	NodeMaximizeButton StyleButton
	NodeMinimizeButton StyleButton
	SymMinimize        SymbolType
	SymMaximize        SymbolType
	Border             float32
	Rounding           float32
	Indent             float32
	Padding            Vec2
	Spacing            Vec2
	ref967cf1ef        *C.struct_nk_style_tab
	allocs967cf1ef     interface{}
}

// CommandLine as declared in nk/nuklear.h:1595
type CommandLine struct {
	Header         Command
	LineThickness  uint16
	Begin          Vec2i
	End            Vec2i
	Color          Color
	ref47fcc852    *C.struct_nk_command_line
	allocs47fcc852 interface{}
}

// CommandScissor as declared in nk/nuklear.h:1589
type CommandScissor struct {
	Header         Command
	X              int16
	Y              int16
	W              uint16
	H              uint16
	refdad76a99    *C.struct_nk_command_scissor
	allocsdad76a99 interface{}
}

// Colorf as declared in nk/nuklear.h:417
type Colorf struct {
	R              float32
	G              float32
	B              float32
	A              float32
	refb6992e05    *C.struct_nk_colorf
	allocsb6992e05 interface{}
}

// StyleWindowHeader as declared in nk/nuklear.h:2319
type StyleWindowHeader struct {
	Normal         StyleItem
	Hover          StyleItem
	Active         StyleItem
	CloseButton    StyleButton
	MinimizeButton StyleButton
	CloseSymbol    SymbolType
	MinimizeSymbol SymbolType
	MaximizeSymbol SymbolType
	LabelNormal    Color
	LabelHover     Color
	LabelActive    Color
	Align          StyleHeaderAlign
	Padding        Vec2
	LabelPadding   Vec2
	Spacing        Vec2
	ref95bb44b7    *C.struct_nk_style_window_header
	allocs95bb44b7 interface{}
}

// MemoryStatus as declared in nk/nuklear.h:1106
type MemoryStatus struct {
	Memory         unsafe.Pointer
	Type           uint32
	Size           Size
	Allocated      Size
	Needed         Size
	Calls          Size
	ref7d6685cf    *C.struct_nk_memory_status
	allocs7d6685cf interface{}
}

// CommandArcFilled as declared in nk/nuklear.h:1680
type CommandArcFilled struct {
	Header         Command
	Cx             int16
	Cy             int16
	R              uint16
	A              [2]float32
	Color          Color
	reff41e98e9    *C.struct_nk_command_arc_filled
	allocsf41e98e9 interface{}
}

// RowLayout as declared in nk/nuklear.h:2448
type RowLayout struct {
	Type           int32
	Index          int32
	Height         float32
	Columns        int32
	Ratio          []float32
	ItemWidth      float32
	ItemHeight     float32
	ItemOffset     float32
	Filled         float32
	Item           Rect
	TreeDepth      int32
	ref3c76e5b1    *C.struct_nk_row_layout
	allocs3c76e5b1 interface{}
}

// CommandArc as declared in nk/nuklear.h:1671
type CommandArc struct {
	Header         Command
	Cx             int16
	Cy             int16
	R              uint16
	LineThickness  uint16
	A              [2]float32
	Color          Color
	ref9e3279d9    *C.struct_nk_command_arc
	allocs9e3279d9 interface{}
}

// ConfigStackVec2Element as declared in nk/nuklear.h:2624
type ConfigStackVec2Element struct {
	Address        []Vec2
	OldValue       Vec2
	ref215f67d3    *C.struct_nk_config_stack_vec2_element
	allocs215f67d3 interface{}
}

// Keyboard as declared in nk/nuklear.h:1796
type Keyboard struct {
	Keys           [29]Key
	Text           [16]byte
	TextLen        int32
	refcc8c697c    *C.struct_nk_keyboard
	allocscc8c697c interface{}
}

// Command as declared in nk/nuklear.h:1581
type Command struct {
	Type           CommandType
	Next           Size
	refa9d15177    *C.struct_nk_command
	allocsa9d15177 interface{}
}

// ConfigurationStacks as declared in nk/nuklear.h:2638
type ConfigurationStacks struct {
	StyleItems      ConfigStackStyleItem
	Floats          ConfigStackFloat
	Vectors         ConfigStackVec2
	Flags           ConfigStackFlags
	Colors          ConfigStackColor
	Fonts           ConfigStackUserFont
	ButtonBehaviors ConfigStackButtonBehavior
	refcbfb6c4b     *C.struct_nk_configuration_stacks
	allocscbfb6c4b  interface{}
}

// EditState as declared in nk/nuklear.h:2528
type EditState struct {
	Name           Hash
	Seq            uint32
	Old            uint32
	Active         int32
	Prev           int32
	Cursor         int32
	SelStart       int32
	SelEnd         int32
	Scrollbar      Scroll
	Mode           byte
	SingleLine     byte
	refcc4bfa07    *C.struct_nk_edit_state
	allocscc4bfa07 interface{}
}

// MouseButton as declared in nk/nuklear.h:1774
type MouseButton struct {
	Down           int32
	Clicked        uint32
	ClickedPos     Vec2
	ref265a908d    *C.struct_nk_mouse_button
	allocs265a908d interface{}
}

// StyleScrollbar as declared in nk/nuklear.h:2150
type StyleScrollbar struct {
	Normal            StyleItem
	Hover             StyleItem
	Active            StyleItem
	BorderColor       Color
	CursorNormal      StyleItem
	CursorHover       StyleItem
	CursorActive      StyleItem
	CursorBorderColor Color
	Border            float32
	Rounding          float32
	BorderCursor      float32
	RoundingCursor    float32
	Padding           Vec2
	ShowButtons       int32
	IncButton         StyleButton
	DecButton         StyleButton
	IncSymbol         SymbolType
	DecSymbol         SymbolType
	Userdata          Handle
	DrawBegin         *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd           *func(arg0 []CommandBuffer, arg1 Handle)
	refe74766c5       *C.struct_nk_style_scrollbar
	allocse74766c5    interface{}
}

// CommandRectMultiColor as declared in nk/nuklear.h:1629
type CommandRectMultiColor struct {
	Header         Command
	X              int16
	Y              int16
	W              uint16
	H              uint16
	Left           Color
	Top            Color
	Bottom         Color
	Right          Color
	ref49471261    *C.struct_nk_command_rect_multi_color
	allocs49471261 interface{}
}

// Memory as declared in nk/nuklear.h:1131
type Memory struct {
	Ptr            unsafe.Pointer
	Size           Size
	reff3973d44    *C.struct_nk_memory
	allocsf3973d44 interface{}
}

// ConfigStackButtonBehavior as declared in nk/nuklear.h:2636
type ConfigStackButtonBehavior struct {
	Head           int32
	Elements       [8]ConfigStackButtonBehaviorElement
	ref6ed3475c    *C.struct_nk_config_stack_button_behavior
	allocs6ed3475c interface{}
}

// Recti as declared in nk/nuklear.h:421
type Recti struct {
	X              int16
	Y              int16
	W              int16
	H              int16
	refea5b5362    *C.struct_nk_recti
	allocsea5b5362 interface{}
}

// Cursor as declared in nk/nuklear.h:425
type Cursor struct {
	Img            Image
	Size           Vec2
	Offset         Vec2
	refdc8563ff    *C.struct_nk_cursor
	allocsdc8563ff interface{}
}

// StyleProgress as declared in nk/nuklear.h:2124
type StyleProgress struct {
	Normal            StyleItem
	Hover             StyleItem
	Active            StyleItem
	BorderColor       Color
	CursorNormal      StyleItem
	CursorHover       StyleItem
	CursorActive      StyleItem
	CursorBorderColor Color
	Rounding          float32
	Border            float32
	CursorBorder      float32
	CursorRounding    float32
	Padding           Vec2
	Userdata          Handle
	DrawBegin         *func(arg0 []CommandBuffer, arg1 Handle)
	DrawEnd           *func(arg0 []CommandBuffer, arg1 Handle)
	ref2f48313c       *C.struct_nk_style_progress
	allocs2f48313c    interface{}
}

// DrawCommand as declared in nk/nuklear.h:405
type DrawCommand C.struct_nk_draw_command

// Image as declared in nk/nuklear.h:424
type Image struct {
	Handle         Handle
	W              uint16
	H              uint16
	Region         [4]uint16
	ref530204c9    *C.struct_nk_image
	allocs530204c9 interface{}
}

// CommandCircleFilled as declared in nk/nuklear.h:1664
type CommandCircleFilled struct {
	Header         Command
	X              int16
	Y              int16
	W              uint16
	H              uint16
	Color          Color
	reff3c944cd    *C.struct_nk_command_circle_filled
	allocsf3c944cd interface{}
}

// Table as declared in nk/nuklear.h:2500
type Table struct {
	Seq            uint32
	Keys           [52]Hash
	Values         [52]Uint
	Next           []Table
	Prev           []Table
	ref60168fd0    *C.struct_nk_table
	allocs60168fd0 interface{}
}

// StyleItemData as declared in nk/nuklear.h:1978
const sizeofStyleItemData = unsafe.Sizeof(C.union_nk_style_item_data{})

type StyleItemData [sizeofStyleItemData]byte

// UserFont as declared in nk/nuklear.h:410
type UserFont struct {
	Userdata       Handle
	Height         float32
	Width          TextWidthF
	ref738ce62e    *C.struct_nk_user_font
	allocs738ce62e interface{}
}

// CommandRectFilled as declared in nk/nuklear.h:1621
type CommandRectFilled struct {
	Header        Command
	Rounding      uint16
	X             int16
	Y             int16
	W             uint16
	H             uint16
	Color         Color
	ref5029f39    *C.struct_nk_command_rect_filled
	allocs5029f39 interface{}
}

// CommandBuffer as declared in nk/nuklear.h:404
type CommandBuffer struct {
	Base           []Buffer
	Clip           Rect
	UseClipping    int32
	Userdata       Handle
	Begin          Size
	End            Size
	Last           Size
	refa28222e0    *C.struct_nk_command_buffer
	allocsa28222e0 interface{}
}

// PopupBuffer as declared in nk/nuklear.h:2461
type PopupBuffer struct {
	Begin          Size
	Parent         Size
	Last           Size
	End            Size
	Active         int32
	ref1ed9add1    *C.struct_nk_popup_buffer
	allocs1ed9add1 interface{}
}

// ConvertConfig as declared in nk/nuklear.h:406
type ConvertConfig struct {
	GlobalAlpha        float32
	LineAa             AntiAliasing
	ShapeAa            AntiAliasing
	CircleSegmentCount uint32
	ArcSegmentCount    uint32
	CurveSegmentCount  uint32
	Null               DrawNullTexture
	VertexLayout       []DrawVertexLayoutElement
	VertexSize         Size
	VertexAlignment    Size
	ref82bf4c25        *C.struct_nk_convert_config
	allocs82bf4c25     interface{}
}

// TextEdit as declared in nk/nuklear.h:408
type TextEdit struct {
	Clip              Clipboard
	String            Str
	Filter            PluginFilter
	Scrollbar         Vec2
	Cursor            int32
	SelectStart       int32
	SelectEnd         int32
	Mode              byte
	CursorAtEndOfLine byte
	Initialized       byte
	HasPreferredX     byte
	SingleLine        byte
	Active            byte
	Padding1          byte
	PreferredX        float32
	Undo              TextUndoState
	ref5fd2021a       *C.struct_nk_text_edit
	allocs5fd2021a    interface{}
}

// Vec2i as declared in nk/nuklear.h:419
type Vec2i struct {
	X              int16
	Y              int16
	refb9f81d21    *C.struct_nk_vec2i
	allocsb9f81d21 interface{}
}

// Scroll as declared in nk/nuklear.h:426
type Scroll struct {
	X              uint16
	Y              uint16
	reff4ed37f2    *C.struct_nk_scroll
	allocsf4ed37f2 interface{}
}

// ConfigStackColor as declared in nk/nuklear.h:2634
type ConfigStackColor struct {
	Head           int32
	Elements       [32]ConfigStackColorElement
	ref237e1889    *C.struct_nk_config_stack_color
	allocs237e1889 interface{}
}
