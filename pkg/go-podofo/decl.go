package podofo

const (
	// BufferSize is used for internal buffers.
	BufferSize = 4096

	// DefaultPDFVersion is the default PDF version
	// used by new PDF documents.
	DefaultPDFVersion = PDFVersion14
)

// Matrix2D.
//
// NOTE: This may change if the future.
type Matrix2D [6]float64

// CIDToGIDMap is a backing storage for a CID to GID map.
//
// It must preserve ordering.
type CIDToGIDMap []struct {
	CID, GID uint
}

type PDFALevel uint8

const (
	PDFALevelUnknown PDFALevel = iota
	PDFALevel1B
	PDFALevel1A
	PDFALevel2B
	PDFALevel2A
	PDFALevel2U
	PDFALevel3B
	PDFALevel3A
	PDFALevel3U
	PDFALevel4E
	PDFALevel4F
)

type PDFStringState uint8

const (
	// PDFStringStateRawBuffer is for an unvaluated
	// raw buffer string.
	PDFStringStateRawBuffer PDFStringState = iota
	// PDFStringStateASCII is for both an ASCII
	// and PDFDocEncoding charsets.
	PDFStringStateASCII
	// PDFStringStateDocEncoding is for strings
	// that use the whole PDFDocEncoding charset.
	PDFStringStateDocEncoding
	// PDFStringStateUnicode is for strings that
	// use the whole Unicode charset.
	PDFStringStateUnicode
)

type EncodingMapType uint8

const (
	// PDFEncodingMapTypeIndeterminate is for
	// indeterminate map type, such as identity
	// encodings.
	EncodingMapTypeIndeterminate EncodingMapType = iota
	// PDFEncodingMapTypeSimple is for a legacy
	// encoding, such as built-in or difference.
	EncodingMapTypeSimple
	// PDFEncodingMapTypeCMap is for proper
	// CMap encoding or pre-defined CMap names.
	EncodingMapTypeCMap
)

type WriteFlags uint16

const (
	WriteFlagsNone WriteFlags = 0
	// WriteFlagsClean is used to created a PDF
	// that is readable in a text editor, i.e.
	// isert spaces and linebreaks between tokens.
	WriteFlagsClean WriteFlags = 1
	// WriteFlagsNoInlineLiteral is used to prevent writing spaces before literal types
	// (numerical, references, null).
	WriteFlagsNoInlineLiteral WriteFlags = 2
	// WriteFlagsNoFlateCompress is used to write
	// PDF with Flate compression.
	WriteFlagsNoFlateCompress WriteFlags = 4
	// WriteFlagsNoPDFAPreserve is used to write
	// compact (WriteFlagsClean is unsed) code,
	// preserving PDF/A compliance is not required.
	WriteFlagsNoPDFAPreserve WriteFlags = 256
)

type TextExtractFlags uint8

const (
	TextExtractFlagNone       TextExtractFlags = 0
	TextExtractFlagIgnoreCase TextExtractFlags = 1 << (iota - 1)
	TextExtractFlagKeepWhiteTokens
	TextExtractFlagTokenizeWords
	TextExtractFlagMatchWholeWord
	TextExtractFlagRegexPattern
	TextExtractFlagComputeBoundingBox
	TextExtractFlagRawCoordinates
	TextExtractFlagExtractSubstring
)

type XObjectType uint8

const (
	XObjectTypeUnknown XObjectType = iota
	XObjectTypeForm
	XObjectTypeImage
	XObjectTypePostScript
)

type FilterType uint8

const (
	FilterTypeNone FilterType = iota
	FilterTypeASCIIHexDecode
	FilterTypeASCII85Decode
	FilterTypeLZWDecode
	FilterTypeFlateDecode
	FilterTypeRunLengthDecode
	FilterTypeCCITTFaxDecode
	FilterTypeJBIG2Decode
	FilterTypeDCTDecode
	FilterTypeJPXDecode
	FilterTypeCrypt
)

type ExportFormat uint8

const (
	ExportFormatPNG  ExportFormat = 1
	ExportFormatJPEG ExportFormat = 2
)

type FontDescriptorFlag uint32

const (
	FontDescriptorFlagNone       FontDescriptorFlag = 0
	FontDescriptorFlagFixedPitch FontDescriptorFlag = 1 << (iota - 1)
	FontDescriptorFlagSerif
	FontDescriptorFlagSymbolic
	FontDescriptorFlagScript
	FontDescriptorFlagNonSymbolic FontDescriptorFlag = 1 << iota
	FontDescriptorFlagItalic
	FontDescriptorFlagAllCap FontDescriptorFlag = 1 << (iota + 9)
	FontDescriptorFlagSmallCap
	FontDescriptorFlagForceBold
)

type FontStretch uint8

const (
	FontStretchUnknown FontStretch = iota
	FontStretchUltraCondensed
	FontStretchExtraCondensed
	FontStretchCondensed
	FontStretchSemiCondensed
	FontStretchNormal
	FontStretchSemiExpanded
	FontStretchExpanded
	FontStretchExtraExpanded
	FontStretchUltraExpanded
)

type FontType uint8

const (
	FontTypeUnknown FontType = iota
	FontTypeType1
	FontTypeType3
	FontTypeTrueType
	FontTypeCIDType1
	FontTypeCIDTrueType
)

type FontFileType uint8

const (
	FontFileTypeUnknown FontFileType = iota
	FontFileTypeType1
	FontFileTypeType1CCF
	FontFileTypeCIDType1
	FontFileTypeType3
	FontFileTypeTrueType
	FontFileTypeOpenType
)

type FontStyle uint8

const (
	FontStyleRegular FontStyle = iota
	FontStyleItalic
	FontStyleBold
)

type GlyphAccess uint8

const (
	GlyphAccessWidth       GlyphAccess = 1
	GlyphAccessFontProgram GlyphAccess = 2
)

type FontAutoSelectBehavior uint8

const (
	FontAutoSelectBehaviorNone FontAutoSelectBehavior = iota
	FontAutoSelectBehaviorStandard14
	FontAutoSelectBehaviorStandard14Alt
)

type FontCreateFlag uint8

const (
	FontCreateFlagNone      FontCreateFlag = 0
	FontCreateFlagDontEmbed FontCreateFlag = 1 << (iota - 1)
	FontCreateFlagDontSubset
	FontCreateFlagPreferNonCID
)

type FontMatchBehaviorFlag uint8

const (
	FontMatchBehaviorFlagNone                FontMatchBehaviorFlag = 0
	FontMatchBehaviorFlagNormalizePattern    FontMatchBehaviorFlag = 1
	FontMatchBehaviorFlagMatchPostScriptName FontMatchBehaviorFlag = 2
)

type ColorSpace uint8

const (
	ColorSpaceUnknown ColorSpace = iota
	ColorSpaceDeviceGray
	ColorSpaceDeviceRGB
	ColorSpaceDeviceCMYK
	ColorSpaceCalGray
	ColorSpaceCalRGB
	ColorSpaceLab
	ColorSpaceICCBased
	ColorSpaceIndexed
	ColorSpacePattern
	ColorSpaceSeparation
	ColorSpaceDeviceN
)

type PixelFormat uint8

const (
	PixelFormatUnknown PixelFormat = 0
	PixelFormatGrayscale
	PixelFormatRGB24
	PixelFormatBGR24
	PixelFormatRGBA
	PixelFormatBGRA
	PixelFormatARGB
	PixelFormatABGR
)

type TextRenderingMode uint8

const (
	TextRenderingModeFill TextRenderingMode = 0
	TextRenderingModeStroke
	TextRenderingModeFillStroke
	TextRenderingModeInvisible
	TextRenderingModeFillAddToClipPath
	TextRenderingModeStrokeAddToClipPath
	TextRenderingModeFillStrokeAddToClipPath
	TextRenderingModeAddToClipPath
)

type StrokeStyle uint8

const (
	StrokeStyleSolid StrokeStyle = iota
	StrokeStyleDash
	StrokeStyleDot
	StrokeStyleDashDot
	StrokeStyleDashDotDot
)

type InfoInitial uint8

const (
	InfoInitialNone InfoInitial = iota
	InfoInitialWriteCreationTime
	InfoInitialWriteModificationTime
	InfoInitialWriteProducer
)

type LineCapStyle uint8

const (
	LineCapStyleButt LineCapStyle = iota
	LineCapStyleRound
	LineCapStyleSquare
)

type LineJoinStyle uint8

const (
	LineJoinStyleMiter LineJoinStyle = iota
	LineJoinStyleRound
	LineJoinStyleBevel
)

type VerticalAlignment uint8

const (
	VerticalAligmentTop VerticalAlignment = iota
	VerticalAligmentCenter
	VerticalAligmentBottom
)

type HorizontalAlignment uint8

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

type SaveOption uint8

const (
	SaveOptionNone      SaveOption = 0
	SaveOptionReserved1 SaveOption = 1 << (iota - 1)
	SaveOptionReserved2
	SaveOptionNoFlateCompress
	SaveOptionNoCollectGarbage
	SaveOptionNoModifyDateUpdate
	SaveOptionClean
)

type PageSize uint8

const (
	PageSizeUnknown PageSize = iota
	PageSizeA0
	PageSizeA1
	PageSizeA2
	PageSizeA3
	PageSizeA4
	PageSizeA5
	PageSizeA6
	PageSizeLetter
	PageSizeLegal
	PageSizeTabloid
)

type PageMode uint8

const (
	PageModeDontCare PageMode = iota
	PageModeUseNone
	PageModeUseThumbs
	PageModeUseBookmarks
	PageModeFullScreen
	PageModeUseOC
	PageModeUseAttachments
)

type PageLayout uint8

const (
	PageLayoutIgnore PageLayout = iota
	PageLayoutDefault
	PageLayoutSinglePage
	PageLayoutOneColumn
	PageLayoutTwoColumnLeft
	PageLayoutTwoColumnRight
	PageLayoutTwoPageLeft
	PageLayoutTwoPageRight
)

type Standard14FontType uint8

const (
	Standart14FontTypeUnknown Standard14FontType = iota
	Standart14FontTypeTimesRoman
	Standart14FontTypeTimesItalic
	Standart14FontTypeTimesBold
	Standart14FontTypeTimesBoldItalic
	Standart14FontTypeHelvetica
	Standart14FontTypeHelveticaOblique
	Standart14FontTypeHelveticaBold
	Standart14FontTypeHelveticaBoldOblique
	Standart14FontTypeCourier
	Standart14FontTypeCourierOblique
	Standart14FontTypeCourierBold
	Standart14FontTypeCourierBoldOblique
	Standart14FontTypeSymbol
	Standart14FontTypeZapfDingbats
)

type AnnotationType uint8

const (
	AnnotationTypeUnknown AnnotationType = iota
	AnnotationTypeText
	AnnotationTypeLink
	AnnotationTypeFreeText
	AnnotationTypeLine
	AnnotationTypeSquare
	AnnotationTypeCircle
	AnnotationTypePolygon
	AnnotationTypePolyLine
	AnnotationTypeHighlight
	AnnotationTypeUnderline
	AnnotationTypeSquiggly
	AnnotationTypeStrikeOut
	AnnotationTypeStamp
	AnnotationTypeCaret
	AnnotationTypeInk
	AnnotationTypePopup
	AnnotationTypeFileAttachement
	AnnotationTypeSound
	AnnotationTypeMovie
	AnnotationTypeWidget
	AnnotationTypeScreen
	AnnotationTypePrinterMark
	AnnotationTypeTrapNet
	AnnotationTypeWatermark
	AnnotationTypeModel3D
	AnnotationTypeRichMedia
	AnnotationTypeWebMedia
	AnnotationTypeRedact
	AnnotationTypeProjection
)

type AnnotationFlag uint32

const (
	AnnotationFlagNone      AnnotationFlag = 0x0000
	AnnotationFlagInvisible AnnotationFlag = 1 << (iota - 1)
	AnnotationFlagHidden
	AnnotationFlagPrint
	AnnotationFlagNoZoom
	AnnotationFlagNoRotate
	AnnotationFlagNoView
	AnnotationFlagReadOnly
	AnnotationFlagLocked
	AnnotationFlagToggleNoView
	AnnotationFlagLockedContents
)

type FieldType uint32

const (
	FieldTypeUnknown FieldType = iota
	FieldTypePushButton
	FieldTypeCheckBox
	FieldTypeRadioButton
	FieldTypeTextBox
	FieldTypeComboBox
	FieldTypeListBox
	FieldTypeSignature
)

type HighlightingMode uint8

const (
	HighlightingModeUnknown HighlightingMode = iota
	HighlightingModeNone
	HighlightingModeInvert
	HighlightingModeInvertOutline
	HighlightingModePush
)

type FieldFlag uint8

const (
	FieldFlagReadOnly FieldFlag = 1
	FieldFlagRequired FieldFlag = 2
	FieldFlagNoExport FieldFlag = 4
)

type AppearanceType uint8

const (
	AppearenceTypeNormal AppearanceType = iota
	AppearenceTypeRollover
	AppearenceTypeDown
)

type Operator uint8

const (
	OperatorUnknown Operator = iota
	// ISO 32008-1:2008 Table 51 – Operator Categories
	// General graphics state
	Operatorw
	OperatorJ
	Operatorj
	OperatorM
	Operatord
	Operatorri
	Operatori
	Operatorgs
	// Special graphics state
	Operatorq
	OperatorQ
	Operatorcm
	// Path construction
	Operatorm
	Operatorl
	Operatorc
	Operatorv
	Operatory
	Operatorh
	Operatorre
	// Path painting
	OperatorS
	Operators
	Operatorf
	OperatorF
	OperatorfStar
	OperatorB
	OperatorBStar
	Operatorb
	OperatorbStar
	Operatorn
	// Clipping paths
	OperatorW
	OperatorWStar
	// Text objects
	OperatorBT
	OperatorET
	// Text state
	OperatorTc
	OperatorTw
	OperatorTz
	OperatorTL
	OperatorTf
	OperatorTr
	OperatorTs
	// Text positioning
	OperatorTd
	OperatorTD
	OperatorTm
	OperatorTStar
	// Text showing
	OperatorTj
	OperatorTJ
	OperatorQuote
	OperatorDoubleQuote
	// Type 3 fonts
	Operatord0
	Operatord1
	// Color
	OperatorCS
	Operatorcs
	OperatorSC
	OperatorSCN
	Operatorsc
	Operatorscn
	OperatorG
	Operatorg
	OperatorRG
	Operatorrg
	OperatorK
	Operatork
	// Shading patterns
	Operatorsh
	// Inline images
	OperatorBI
	OperatorID
	OperatorEI
	// XObjects
	OperatorDo
	// Marked content
	OperatorMP
	OperatorDP
	OperatorBMC
	OperatorBDC
	OperatorEMC
	// Compatibility
	OperatorBX
	OperatorEX
)

type RenderingIntent uint8

const (
	RenderingIntentAbsoluteColorimetric RenderingIntent = iota
	RenderingIntentRelativeColorimetric
	RenderingIntentPerceptual
	RenderingIntentSaturation
)

type BlendMode uint8

const (
	BlendModeNormal BlendMode = iota
	BlendModeMultiply
	BlendModeScreen
	BlendModeOverlay
	BlendModeDarken
	BlendModeLighten
	BlendModeColorDodge
	BlendModeColorBurn
	BlendModeHardLight
	BlendModeSoftLight
	BlendModeDifference
	BlendModeExclusion
	BlendModeHue
	BlendModeSaturation
	BlendModeColor
	BlendModeLuminosity
)
