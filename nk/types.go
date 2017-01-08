// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 08 Jan 2017 04:33:47 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package nk

/*
#include "nuklear.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Char type as declared in nk/nuklear.h:390
type Char byte

// Uchar type as declared in nk/nuklear.h:391
type Uchar byte

// Byte type as declared in nk/nuklear.h:392
type Byte byte

// Short type as declared in nk/nuklear.h:393
type Short int16

// Ushort type as declared in nk/nuklear.h:394
type Ushort uint16

// Int type as declared in nk/nuklear.h:395
type Int int32

// Uint type as declared in nk/nuklear.h:396
type Uint uint32

// Size type as declared in nk/nuklear.h:397
type Size uint

// Ptr type as declared in nk/nuklear.h:398
type Ptr uint

// Hash type as declared in nk/nuklear.h:400
type Hash uint32

// Flags type as declared in nk/nuklear.h:401
type Flags uint32

// Rune type as declared in nk/nuklear.h:402
type Rune uint32

// Glyph type as declared in nk/nuklear.h:455
type Glyph [4]byte

// Handle as declared in nk/nuklear.h:456
const sizeofHandle = unsafe.Sizeof(C.nk_handle{})

type Handle [sizeofHandle]byte

// PluginAlloc type as declared in nk/nuklear.h:475
type PluginAlloc func(arg0 Handle, old unsafe.Pointer, arg2 Size) unsafe.Pointer

// PluginFree type as declared in nk/nuklear.h:476
type PluginFree func(arg0 Handle, old unsafe.Pointer)

// PluginFilter type as declared in nk/nuklear.h:477
type PluginFilter func(arg0 *TextEdit, unicode Rune) int32

// PluginPaste type as declared in nk/nuklear.h:478
type PluginPaste func(arg0 Handle, arg1 *TextEdit)

// PluginCopy type as declared in nk/nuklear.h:479
type PluginCopy func(arg0 Handle, arg1 string, len int32)

// TextWidthF type as declared in nk/nuklear.h:1285
type TextWidthF func(arg0 Handle, h float32, arg2 string, len int32) float32

// QueryFontGlyphF type as declared in nk/nuklear.h:1286
type QueryFontGlyphF func(handle Handle, fontHeight float32, glyph *UserFontGlyph, codepoint Rune, nextCodepoint Rune)

// DrawIndex type as declared in nk/nuklear.h:2041
type DrawIndex uint16

// TextUndoState as declared in nk/nuklear.h:1646
type TextUndoState C.struct_nk_text_undo_state

// DrawList as declared in nk/nuklear.h:429
type DrawList C.struct_nk_draw_list

// StyleToggle as declared in nk/nuklear.h:435
type StyleToggle C.struct_nk_style_toggle

// Mouse as declared in nk/nuklear.h:1980
type Mouse C.struct_nk_mouse

// PageData as declared in nk/nuklear.h:2887
const sizeofPageData = unsafe.Sizeof(C.union_nk_page_data{})

type PageData [sizeofPageData]byte

// StyleWindowHeader as declared in nk/nuklear.h:445
type StyleWindowHeader C.struct_nk_style_window_header

// Key as declared in nk/nuklear.h:1991
type Key C.struct_nk_key

// ConfigStackUserFont as declared in nk/nuklear.h:2861
type ConfigStackUserFont struct {
	Head           int32
	Elements       [8]ConfigStackUserFontElement
	refa664861d    *C.struct_nk_config_stack_user_font
	allocsa664861d interface{}
}

// StyleCombo as declared in nk/nuklear.h:443
type StyleCombo C.struct_nk_style_combo

// ConfigStackButtonBehaviorElement as declared in nk/nuklear.h:2854
type ConfigStackButtonBehaviorElement C.struct_nk_config_stack_button_behavior_element

// ListView as declared in nk/nuklear.h:504
type ListView C.struct_nk_list_view

// StyleItemData as declared in nk/nuklear.h:2180
const sizeofStyleItemData = unsafe.Sizeof(C.union_nk_style_item_data{})

type StyleItemData [sizeofStyleItemData]byte

// Page as declared in nk/nuklear.h:2899
type Page C.struct_nk_page

// StyleTab as declared in nk/nuklear.h:444
type StyleTab C.struct_nk_style_tab

// PopupBuffer as declared in nk/nuklear.h:2663
type PopupBuffer C.struct_nk_popup_buffer

// Panel as declared in nk/nuklear.h:431
type Panel C.struct_nk_panel

// StyleWindow as declared in nk/nuklear.h:446
type StyleWindow C.struct_nk_style_window

// CommandTriangle as declared in nk/nuklear.h:1840
type CommandTriangle C.struct_nk_command_triangle

// Memory as declared in nk/nuklear.h:1503
type Memory C.struct_nk_memory

// CommandCurve as declared in nk/nuklear.h:1804
type CommandCurve C.struct_nk_command_curve

// MemoryStatus as declared in nk/nuklear.h:1478
type MemoryStatus C.struct_nk_memory_status

// StyleEdit as declared in nk/nuklear.h:440
type StyleEdit C.struct_nk_style_edit

// ConfigurationStacks as declared in nk/nuklear.h:2864
type ConfigurationStacks C.struct_nk_configuration_stacks

// Buffer as declared in nk/nuklear.h:422
type Buffer C.struct_nk_buffer

// Vec2i as declared in nk/nuklear.h:452
type Vec2i C.struct_nk_vec2i

// Recti as declared in nk/nuklear.h:454
type Recti C.struct_nk_recti

// StyleScrollbar as declared in nk/nuklear.h:439
type StyleScrollbar C.struct_nk_style_scrollbar

// PropertyState as declared in nk/nuklear.h:2744
type PropertyState C.struct_nk_property_state

// CommandCircle as declared in nk/nuklear.h:1857
type CommandCircle C.struct_nk_command_circle

// ConfigStackVec2Element as declared in nk/nuklear.h:2850
type ConfigStackVec2Element C.struct_nk_config_stack_vec2_element

// RowLayout as declared in nk/nuklear.h:2649
type RowLayout C.struct_nk_row_layout

// Color as declared in nk/nuklear.h:449
type Color C.struct_nk_color

// Input as declared in nk/nuklear.h:2001
type Input C.struct_nk_input

// Pool as declared in nk/nuklear.h:2905
type Pool C.struct_nk_pool

// MenuState as declared in nk/nuklear.h:2671
type MenuState C.struct_nk_menu_state

// CommandLine as declared in nk/nuklear.h:1796
type CommandLine C.struct_nk_command_line

// ConfigStackUserFontElement as declared in nk/nuklear.h:2853
type ConfigStackUserFontElement struct {
	Address        [][]UserFont
	OldValue       []UserFont
	ref5572630c    *C.struct_nk_config_stack_user_font_element
	allocs5572630c interface{}
}

// FontAtlas as declared in nk/nuklear.h:1395
type FontAtlas C.struct_nk_font_atlas

// ConfigStackFlagsElement as declared in nk/nuklear.h:2851
type ConfigStackFlagsElement C.struct_nk_config_stack_flags_element

// StyleSlider as declared in nk/nuklear.h:2287
type StyleSlider C.struct_nk_style_slider

// CommandImage as declared in nk/nuklear.h:1912
type CommandImage C.struct_nk_command_image

// ConfigStackVec2 as declared in nk/nuklear.h:2858
type ConfigStackVec2 C.struct_nk_config_stack_vec2

// PopupState as declared in nk/nuklear.h:2720
type PopupState C.struct_nk_popup_state

// ConfigStackStyleItemElement as declared in nk/nuklear.h:2848
type ConfigStackStyleItemElement C.struct_nk_config_stack_style_item_element

// PageElement as declared in nk/nuklear.h:2893
type PageElement C.struct_nk_page_element

// CommandRect as declared in nk/nuklear.h:1813
type CommandRect C.struct_nk_command_rect

// ConfigStackColorElement as declared in nk/nuklear.h:2852
type ConfigStackColorElement C.struct_nk_config_stack_color_element

// ConfigStackFloat as declared in nk/nuklear.h:2857
type ConfigStackFloat C.struct_nk_config_stack_float

// CommandArcFilled as declared in nk/nuklear.h:1881
type CommandArcFilled C.struct_nk_command_arc_filled

// Cursor as declared in nk/nuklear.h:458
type Cursor C.struct_nk_cursor

// Table as declared in nk/nuklear.h:2703
type Table C.struct_nk_table

// Colorf as declared in nk/nuklear.h:450
type Colorf C.struct_nk_colorf

// StyleText as declared in nk/nuklear.h:2190
type StyleText C.struct_nk_style_text

// StyleChart as declared in nk/nuklear.h:442
type StyleChart C.struct_nk_style_chart

// TextUndoRecord as declared in nk/nuklear.h:1639
type TextUndoRecord C.struct_nk_text_undo_record

// Rect as declared in nk/nuklear.h:453
type Rect C.struct_nk_rect

// CommandTriangleFilled as declared in nk/nuklear.h:1849
type CommandTriangleFilled C.struct_nk_command_triangle_filled

// CommandText as declared in nk/nuklear.h:1920
type CommandText C.struct_nk_command_text

// UserFontGlyph as declared in nk/nuklear.h:1284
type UserFontGlyph struct {
	Uv             [2]Vec2
	Offset         Vec2
	Width          float32
	Height         float32
	Xadvance       float32
	ref4a84b297    *C.struct_nk_user_font_glyph
	allocs4a84b297 interface{}
}

// CommandScissor as declared in nk/nuklear.h:1790
type CommandScissor C.struct_nk_command_scissor

// Context as declared in nk/nuklear.h:432
type Context C.struct_nk_context

// StyleButton as declared in nk/nuklear.h:434
type StyleButton C.struct_nk_style_button

// Vec2 as declared in nk/nuklear.h:451
type Vec2 C.struct_nk_vec2

// Font as declared in nk/nuklear.h:1378
type Font C.struct_nk_font

// CommandArc as declared in nk/nuklear.h:1872
type CommandArc C.struct_nk_command_arc

// Keyboard as declared in nk/nuklear.h:1995
type Keyboard C.struct_nk_keyboard

// Clipboard as declared in nk/nuklear.h:1633
type Clipboard C.struct_nk_clipboard

// TextEdit as declared in nk/nuklear.h:428
type TextEdit C.struct_nk_text_edit

// CommandRectFilled as declared in nk/nuklear.h:1822
type CommandRectFilled C.struct_nk_command_rect_filled

// Image as declared in nk/nuklear.h:457
type Image C.struct_nk_image

// ConvertConfig as declared in nk/nuklear.h:426
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

// Command as declared in nk/nuklear.h:1782
type Command C.struct_nk_command

// MouseButton as declared in nk/nuklear.h:1975
type MouseButton C.struct_nk_mouse_button

// StyleItem as declared in nk/nuklear.h:427
type StyleItem C.struct_nk_style_item

// EditState as declared in nk/nuklear.h:2731
type EditState C.struct_nk_edit_state

// Scroll as declared in nk/nuklear.h:459
type Scroll C.struct_nk_scroll

// UserFont as declared in nk/nuklear.h:430
type UserFont struct {
	Userdata       Handle
	Height         float32
	Width          TextWidthF
	Query          QueryFontGlyphF
	Texture        Handle
	ref738ce62e    *C.struct_nk_user_font
	allocs738ce62e interface{}
}

// FontConfig as declared in nk/nuklear.h:1337
type FontConfig C.struct_nk_font_config

// StyleProperty as declared in nk/nuklear.h:441
type StyleProperty C.struct_nk_style_property

// DrawNullTexture as declared in nk/nuklear.h:487
type DrawNullTexture C.struct_nk_draw_null_texture

// Style as declared in nk/nuklear.h:2581
type Style C.struct_nk_style

// FontGlyph as declared in nk/nuklear.h:1371
type FontGlyph C.struct_nk_font_glyph

// CommandRectMultiColor as declared in nk/nuklear.h:1830
type CommandRectMultiColor C.struct_nk_command_rect_multi_color

// CommandPolygon as declared in nk/nuklear.h:1889
type CommandPolygon C.struct_nk_command_polygon

// Chart as declared in nk/nuklear.h:2643
type Chart C.struct_nk_chart

// ConfigStackFlags as declared in nk/nuklear.h:2859
type ConfigStackFlags C.struct_nk_config_stack_flags

// StyleSelectable as declared in nk/nuklear.h:436
type StyleSelectable C.struct_nk_style_selectable

// ConfigStackStyleItem as declared in nk/nuklear.h:2856
type ConfigStackStyleItem C.struct_nk_config_stack_style_item

// CommandBuffer as declared in nk/nuklear.h:424
type CommandBuffer C.struct_nk_command_buffer

// StyleProgress as declared in nk/nuklear.h:438
type StyleProgress C.struct_nk_style_progress

// Window as declared in nk/nuklear.h:2755
type Window C.struct_nk_window

// Allocator as declared in nk/nuklear.h:423
type Allocator C.struct_nk_allocator

// StyleSlide as declared in nk/nuklear.h:437
type StyleSlide C.struct_nk_style_slide

// ConfigStackButtonBehavior as declared in nk/nuklear.h:2862
type ConfigStackButtonBehavior C.struct_nk_config_stack_button_behavior

// Str as declared in nk/nuklear.h:1550
type Str C.struct_nk_str

// CommandCircleFilled as declared in nk/nuklear.h:1865
type CommandCircleFilled C.struct_nk_command_circle_filled

// ConfigStackColor as declared in nk/nuklear.h:2860
type ConfigStackColor C.struct_nk_config_stack_color

// ChartSlot as declared in nk/nuklear.h:2633
type ChartSlot C.struct_nk_chart_slot

// DrawVertexLayoutElement as declared in nk/nuklear.h:433
type DrawVertexLayoutElement struct {
	Attribute      DrawVertexLayoutAttribute
	Format         DrawVertexLayoutFormat
	Offset         Size
	refeb0614d6    *C.struct_nk_draw_vertex_layout_element
	allocseb0614d6 interface{}
}

// CommandPolyline as declared in nk/nuklear.h:1904
type CommandPolyline C.struct_nk_command_polyline

// ConfigStackFloatElement as declared in nk/nuklear.h:2849
type ConfigStackFloatElement C.struct_nk_config_stack_float_element

// CommandPolygonFilled as declared in nk/nuklear.h:1897
type CommandPolygonFilled C.struct_nk_command_polygon_filled

// DrawCommand as declared in nk/nuklear.h:425
type DrawCommand C.struct_nk_draw_command

// BakedFont as declared in nk/nuklear.h:1324
type BakedFont C.struct_nk_baked_font

// BufferMarker as declared in nk/nuklear.h:1498
type BufferMarker C.struct_nk_buffer_marker
