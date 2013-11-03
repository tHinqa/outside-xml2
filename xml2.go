// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package xml2 provides API definitions for accessing the libxml2 dll.
package xml2

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllData(dll, false, dataList)
}

//TODO(t):check for *Char->string and arithmetic candidates

type (
	fix uintptr

	FILE   fix
	SOCKET fix

	Char          int8 // joined with XmlChar uint8; hope no arithmetic
	Double        float64
	Enum          int
	Long          int32 // TODO(t): Size?
	SizeT         uintptr
	UnsignedChar  uint8
	UnsignedInt   uint
	UnsignedLong  uint32 //TODO(t): check  size
	UnsignedShort uint16

	Void struct{}

	DocbDoc           Doc
	DocbParserCtxt    ParserCtxt
	DocbParserInput   ParserInput
	DocbSAXHandler    SAXHandler
	HtmlDoc           Doc
	HtmlNode          Node
	HtmlParserCtxt    ParserCtxt
	HtmlSAXHandler    SAXHandler
	XlinkHRef         *Char
	XlinkRole         *Char
	XlinkTitle        *Char
	AttributeTable    HashTable
	ElementTable      HashTable
	EntitiesTable     HashTable
	IDTable           HashTable
	NotationTable     HashTable
	NsType            ElementType
	RefTable          HashTable
	TextReaderLocator Void
	IsPubidCharTab    [256]uint8

	Automata             struct{}
	AutomataState        struct{}
	Buf                  struct{}
	Buffer               struct{}
	Catalog              struct{}
	Dict                 struct{}
	DOMWrapCtxt          struct{}
	ExpCtxt              struct{}
	ExpNode              struct{}
	HashTable            struct{}
	Link                 struct{}
	List                 struct{}
	Module               struct{}
	Mutex                struct{}
	Pattern              struct{}
	RegExecCtxt          struct{}
	Regexp               struct{}
	RelaxNG              struct{}
	RelaxNGParserCtxt    struct{}
	RelaxNGValidCtxt     struct{}
	RMutex               struct{}
	SaveCtxt             struct{}
	Schema               struct{}
	SchemaFacet          struct{}
	SchemaParserCtxt     struct{}
	SchemaSAXPlugStruct  struct{}
	Schematron           struct{}
	SchematronParserCtxt struct{}
	SchematronValidCtxt  struct{}
	SchemaType           struct{}
	SchemaVal            struct{}
	SchemaValidCtxt      struct{}
	StreamCtxt           struct{}
	TextReader           struct{}
	TextWriter           struct{}
	ValidState           struct{}
	XIncludeCtxt         struct{}
	XPathCompExpr        struct{}
	XPathParserContext   struct{}

	InternalSubsetSAXFunc func(
		ctx *Void, name, ExternalID, systemID string)

	IsStandaloneSAXFunc func(ctx *Void) int

	HasInternalSubsetSAXFunc func(ctx *Void) int

	HasExternalSubsetSAXFunc func(ctx *Void) int

	StartElementNsSAX2Func func(ctx *Void,
		localname, prefix, URI string,
		nbNamespaces int, namespaces *string,
		nbAttributes, nbDefaulted int, attributes *string)

	EndElementNsSAX2Func func(
		ctx *Void, localname, prefix, URI string)

	ResolveEntitySAXFunc func(
		ctx *Void,
		publicId string,
		systemId string) *ParserInput

	ExternalSubsetSAXFunc func(
		ctx *Void,
		name string,
		externalID string,
		systemID string)

	GetEntitySAXFunc func(
		ctx *Void, name string) *Entity

	GetParameterEntitySAXFunc func(
		ctx *Void, name string) *Entity

	EntityDeclSAXFunc func(ctx *Void, name string,
		typ int, publicId, systemId, content string)

	NotationDeclSAXFunc func(
		ctx *Void, name, publicId, systemId string)

	AttributeDeclSAXFunc func(ctx *Void, elem,
		fullname string, typ, def int,
		defaultValue string, tree *Enumeration)

	ElementDeclSAXFunc func(ctx *Void, name string,
		typ int, content *ElementContent)

	UnparsedEntityDeclSAXFunc func(ctx *Void,
		name, publicId, systemId, notationName string)

	SetDocumentLocatorSAXFunc func(
		ctx *Void, loc *SAXLocator)

	StartDocumentSAXFunc func(ctx *Void)

	EndDocumentSAXFunc func(ctx *Void)

	StartElementSAXFunc func(
		ctx *Void, name string, atts *string)

	EndElementSAXFunc func(ctx *Void, name string)

	AttributeSAXFunc func(ctx *Void, name, value string)

	ReferenceSAXFunc func(ctx *Void, name string)

	CharactersSAXFunc func(ctx *Void, ch string, leng int)

	IgnorableWhitespaceSAXFunc func(
		ctx *Void, ch string, leng int)

	ProcessingInstructionSAXFunc func(
		ctx *Void, target, data string)

	CommentSAXFunc func(ctx *Void, value string)

	CdataBlockSAXFunc func(ctx *Void, value string, leng int)

	ParserInputDeallocate func(str string)

	InputReadCallback func(
		context *Void, buffer string, leng int) int

	InputCloseCallback func(context *Void) int

	WarningSAXFunc func(ctx *Void, msg string, v ...VArg)

	ErrorSAXFunc func(ctx *Void, msg string, v ...VArg)

	FatalErrorSAXFunc func(ctx *Void, msg string, v ...VArg)

	DeregisterNodeFunc func(node *Node)

	GenericErrorFunc func(ctx *Void, msg string, v ...VArg)

	CharEncodingInputFunc func(
		out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int) int

	CharEncodingOutputFunc func(
		out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int) int

	ParserInputBufferCreateFilenameFunc func(
		URI string, enc CharEncoding) *ParserInputBuffer

	RegisterNodeFunc func(node *Node)

	StructuredErrorFunc func(
		userData *Void, Error *Error)

	OutputWriteCallback func(
		context *Void, buffer string, leng int) int

	OutputCloseCallback func(context *Void) int

	ValidityErrorFunc func(ctx *Void, msg string, v ...VArg)

	ValidityWarningFunc func(ctx *Void, msg string, v ...VArg)

	OutputBufferCreateFilenameFunc func(
		URI string,
		encoder *CharEncodingHandler,
		compression int) *OutputBuffer

	XlinkNodeDetectFunc func(ctx *Void, node *Node)

	XlinkSimpleLinkFunk func(ctx *Void, node *Node,
		href XlinkHRef, role XlinkRole, title XlinkTitle)

	XlinkExtendedLinkFunk func(ctx *Void, node *Node,
		nbLocators int, hrefs *XlinkHRef, roles *XlinkRole,
		nbArcs int, from *XlinkRole, to *XlinkRole,
		show *XlinkShow, actuate *XlinkActuate,
		nbTitles int, titles *XlinkTitle, langs *string)

	XlinkExtendedLinkSetFunk func(ctx *Void, node *Node,
		nbLocators int, hrefs *XlinkHRef, roles *XlinkRole,
		nbTitles int, titles *XlinkTitle, langs *string)

	FreeFunc func(mem *Void)

	MallocFunc func(size SizeT) *Void

	ReallocFunc func(mem *Void, size SizeT) *Void

	StrdupFunc func(str string) string

	HashDeallocator func(payload *Void, name string)

	HashCopier func(payload *Void, name string) *Void

	HashScanner func(payload, data *Void, name string)

	HashScannerFull func(payload, data *Void,
		name, name2, name3 string)

	ExternalEntityLoader func(URL, ID string,
		context *ParserCtxt) *ParserInput

	RegExecCallbacks func(exec *RegExecCtxt,
		token string, transdata, inputdata *Void)

	ListDeallocator func(lk *Link)

	ListDataCompare func(data0, data1 *Void) int

	ListWalker func(data, user *Void) int

	InputMatchCallback func(filename string) int

	InputOpenCallback func(filename string) *Void

	OutputMatchCallback func(filename string) int

	OutputOpenCallback func(filename string) *Void

	XPathAxisFunc func(ctxt *XPathParserContext,
		cur *XPathObject) *XPathObject

	XPathVariableLookupFunc func(
		ctxt *Void, name, nsUri string) *XPathObject

	XPathFuncLookupFunc func(
		ctxt *Void, name, nsUri string) XPathFunction

	XPathFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathConvertFunc func(obj *XPathObject, typ int) int

	SchemaValidityErrorFunc func(
		ctx *Void, msg string, v ...VArg)

	SchemaValidityWarningFunc func(
		ctx *Void, msg string, v ...VArg)

	SchemaValidityLocatorFunc func(
		ctx *Void, file *string, line *UnsignedLong) int

	EntityReferenceFunc func(
		ent *Entity, firstNode, lastNode *Node)

	RelaxNGValidityErrorFunc func(
		ctx *Void, msg string, v ...VArg)

	RelaxNGValidityWarningFunc func(
		ctx *Void, msg string, v ...VArg)

	TextReaderErrorFunc func(arg *Void, msg string,
		severity ParserSeverities,
		locator *TextReaderLocator)

	ShellReadlineFunc func(prompt string) string

	FtpListCallback func(userData *Void,
		filename, attrib, owner, group string,
		size UnsignedLong, links, year int,
		month string, day, hour, minute int)

	FtpDataCallback func(userData *Void, data string, leng int)

	C14NIsVisibleCallback func(
		userData *Void, node, parent *Node) int

	ParserCtxt struct {
		Sax               *SAXHandler
		UserData          *Void
		MyDoc             *Doc
		WellFormed        int
		ReplaceEntities   int
		Version           *Char
		Encoding          *Char
		Standalone        int
		Html              int
		Input             *ParserInput
		InputNr           int
		InputMax          int
		InputTab          **ParserInput
		Node              *Node
		NodeNr            int
		NodeMax           int
		NodeTab           **Node
		RecordInfo        int
		NodeSeq           ParserNodeInfoSeq
		ErrNo             int
		HasExternalSubset int
		HasPErefs         int
		External          int
		Valid             int
		Validate          int
		Vctxt             ValidCtxt
		Instate           ParserInputState
		Token             int
		Directory         *Char
		Name              *Char
		NameNr            int
		NameMax           int
		NameTab           **Char
		NbChars           Long
		CheckIndex        Long
		KeepBlanks        int
		DisableSAX        int
		InSubset          int
		IntSubName        *Char
		ExtSubURI         *Char
		ExtSubSystem      *Char
		Space             *int
		SpaceNr           int
		SpaceMax          int
		SpaceTab          *int
		Depth             int
		Entity            *ParserInput
		Charset           int
		Nodelen           int
		Nodemem           int
		Pedantic          int
		_                 *Void
		Loadsubset        int
		Linenumbers       int
		Catalogs          *Void
		Recovery          int
		Progressive       int
		Dict              *Dict
		Atts              **Char
		Maxatts           int
		Docdict           int
		StrXml            *Char
		StrXmlns          *Char
		StrXmlNs          *Char
		Sax2              int
		NsNr              int
		NsMax             int
		NsTab             **Char
		Attallocs         *int
		PushTab           **Void
		AttsDefault       *HashTable
		AttsSpecial       *HashTable
		NsWellFormed      int
		Options           int
		DictNames         int
		FreeElemsNr       int
		FreeElems         *Node
		FreeAttrsNr       int
		FreeAttrs         *Attr
		LastError         Error
		ParseMode         ParserMode
		Nbentities        UnsignedLong
		Sizeentities      UnsignedLong
		NodeInfo          *ParserNodeInfo
		NodeInfoNr        int
		NodeInfoMax       int
		NodeInfoTab       *ParserNodeInfo
		InputId           int
	}

	SAXHandlerV1 struct {
		InternalSubset        InternalSubsetSAXFunc
		IsStandalone          IsStandaloneSAXFunc
		HasInternalSubset     HasInternalSubsetSAXFunc
		HasExternalSubset     HasExternalSubsetSAXFunc
		ResolveEntity         ResolveEntitySAXFunc
		GetEntity             GetEntitySAXFunc
		EntityDecl            EntityDeclSAXFunc
		NotationDecl          NotationDeclSAXFunc
		AttributeDecl         AttributeDeclSAXFunc
		ElementDecl           ElementDeclSAXFunc
		UnparsedEntityDecl    UnparsedEntityDeclSAXFunc
		SetDocumentLocator    SetDocumentLocatorSAXFunc
		StartDocument         StartDocumentSAXFunc
		EndDocument           EndDocumentSAXFunc
		StartElement          StartElementSAXFunc
		EndElement            EndElementSAXFunc
		Reference             ReferenceSAXFunc
		Characters            CharactersSAXFunc
		IgnorableWhitespace   IgnorableWhitespaceSAXFunc
		ProcessingInstruction ProcessingInstructionSAXFunc
		Comment               CommentSAXFunc
		Warning               WarningSAXFunc
		Error                 ErrorSAXFunc
		FatalError            FatalErrorSAXFunc
		GetParameterEntity    GetParameterEntitySAXFunc
		CdataBlock            CdataBlockSAXFunc
		ExternalSubset        ExternalSubsetSAXFunc
		Initialized           UnsignedInt
	}

	ParserInput struct {
		Buf        *ParserInputBuffer
		Filename   *Char
		Directory  *Char
		Base       *Char
		Cur        *Char
		End        *Char
		Length     int
		Line       int
		Col        int
		Consumed   UnsignedLong
		Free       ParserInputDeallocate
		Encoding   *Char
		Version    *Char
		Standalone int
		Id         int
	}

	CharEncodingHandler struct {
		Name   *Char
		Input  CharEncodingInputFunc
		Output CharEncodingOutputFunc
	}

	Error struct {
		Domain  int
		Code    int
		Message *Char
		Level   ErrorLevel
		File    *Char
		Line    int
		Str1    *Char
		Str2    *Char
		Str3    *Char
		Int1    int
		Int2    int
		Ctxt    *Void
		Node    *Void
	}

	OutputBuffer struct {
		Context       *Void
		Writecallback OutputWriteCallback
		Closecallback OutputCloseCallback
		Encoder       *CharEncodingHandler
		Buffer        *Buf
		Conv          *Buf
		Written       int
		Error         int
	}

	ParserInputBuffer struct {
		Context       *Void
		Readcallback  InputReadCallback
		Closecallback InputCloseCallback
		Encoder       *CharEncodingHandler
		Buffer        *Buf
		Raw           *Buf
		Compressed    int
		Error         int
		Rawconsumed   UnsignedLong
	}

	SAXLocator struct {
		GetPublicId     func(ctx *Void) *Char
		GetSystemId     func(ctx *Void) *Char
		GetLineNumber   func(ctx *Void) int
		GetColumnNumber func(ctx *Void) int
	}

	SAXHandler struct {
		InternalSubset        InternalSubsetSAXFunc
		IsStandalone          IsStandaloneSAXFunc
		HasInternalSubset     HasInternalSubsetSAXFunc
		HasExternalSubset     HasExternalSubsetSAXFunc
		ResolveEntity         ResolveEntitySAXFunc
		GetEntity             GetEntitySAXFunc
		EntityDecl            EntityDeclSAXFunc
		NotationDecl          NotationDeclSAXFunc
		AttributeDecl         AttributeDeclSAXFunc
		ElementDecl           ElementDeclSAXFunc
		UnparsedEntityDecl    UnparsedEntityDeclSAXFunc
		SetDocumentLocator    SetDocumentLocatorSAXFunc
		StartDocument         StartDocumentSAXFunc
		EndDocument           EndDocumentSAXFunc
		StartElement          StartElementSAXFunc
		EndElement            EndElementSAXFunc
		Reference             ReferenceSAXFunc
		Characters            CharactersSAXFunc
		IgnorableWhitespace   IgnorableWhitespaceSAXFunc
		ProcessingInstruction ProcessingInstructionSAXFunc
		Comment               CommentSAXFunc
		Warning               WarningSAXFunc
		Error                 ErrorSAXFunc
		FatalError            FatalErrorSAXFunc
		GetParameterEntity    GetParameterEntitySAXFunc
		CdataBlock            CdataBlockSAXFunc
		ExternalSubset        ExternalSubsetSAXFunc
		Initialized           UnsignedInt
		_                     *Void
		StartElementNs        StartElementNsSAX2Func
		EndElementNs          EndElementNsSAX2Func
		Serror                StructuredErrorFunc
	}

	ParserNodeInfoSeq struct {
		Maximum UnsignedLong
		Length  UnsignedLong
		Buffer  *ParserNodeInfo
	}

	ParserNodeInfo struct {
		Node      *Node
		BeginPos  UnsignedLong
		BeginLine UnsignedLong
		EndPos    UnsignedLong
		EndLine   UnsignedLong
	}

	ValidCtxt struct {
		UserData  *Void
		Error     ValidityErrorFunc
		Warning   ValidityWarningFunc
		Node      *Node
		NodeNr    int
		NodeMax   int
		NodeTab   **Node
		FinishDtd UnsignedInt
		Doc       *Doc
		Valid     int
		Vstate    *ValidState
		VstateNr  int
		VstateMax int
		VstateTab *ValidState
		Am        *Automata
		State     *AutomataState
	}

	Dtd struct {
		_          *Void
		Type       ElementType
		Name       *Char
		Children   *Node
		Last       *Node
		Parent     *Doc
		Next       *Node
		Prev       *Node
		Doc        *Doc
		Notations  *Void
		Elements   *Void
		Attributes *Void
		Entities   *Void
		ExternalID *Char
		SystemID   *Char
		Pentities  *Void
	}

	HtmlElemDesc struct {
		Name          *Char
		StartTag      Char
		EndTag        Char
		SaveEndTag    Char
		Empty         Char
		Depr          Char
		Dtd           Char
		Isinline      Char
		Desc          *Char
		Subelts       **Char
		Defaultsubelt *Char
		AttrsOpt      **Char
		AttrsDepr     **Char
		AttrsReq      **Char
	}

	HtmlEntityDesc struct {
		Value UnsignedInt
		Name  *Char
		Desc  *Char
	}

	XlinkHandler struct {
		Simple   XlinkSimpleLinkFunk
		Extended XlinkExtendedLinkFunk
		Set      XlinkExtendedLinkSetFunk
	}

	AttributeStruct struct {
		_    *Void
		Type ElementType
		Name *Char
	}

	Notation struct {
		Name     *Char
		PublicID *Char
		systemID *Char
	}

	Element struct {
		_          *Void
		Type       ElementType
		Name       *Char
		Children   *Node
		Last       *Node
		Parent     *Dtd
		Next       *Node
		Prev       *Node
		Doc        *Doc
		Etype      ElementTypeVal
		Content    *ElementContent
		Attributes *AttributeStruct
		Prefix     *Char
		ContModel  *Regexp
	}

	ID struct {
		Next   *ID
		Value  *Char
		Attr   *Attr
		Name   *Char
		Lineno int
		Doc    *Doc
	}

	Ref struct {
		Next   *Ref
		Value  *Char
		Attr   *Attr
		Name   *Char
		Lineno int
	}

	GlobalState struct {
		XmlParserVersion                        *Char
		XmlDefaultSAXLocator                    SAXLocator
		XmlDefaultSAXHandler                    SAXHandlerV1
		DocbDefaultSAXHandler                   SAXHandlerV1
		HtmlDefaultSAXHandler                   SAXHandlerV1
		XmlFree                                 FreeFunc
		XmlMalloc                               MallocFunc
		XmlMemStrdup                            StrdupFunc
		XmlRealloc                              ReallocFunc
		XmlGenericError                         GenericErrorFunc
		XmlStructuredError                      StructuredErrorFunc
		XmlGenericErrorContext                  *Void
		OldXMLWDcompatibility                   int
		XmlBufferAllocScheme                    BufferAllocationScheme
		XmlDefaultBufferSize                    int
		XmlSubstituteEntitiesDefaultValue       int
		XmlDoValidityCheckingDefaultValue       int
		XmlGetWarningsDefaultValue              int
		XmlKeepBlanksDefaultValue               int
		XmlLineNumbersDefaultValue              int
		XmlLoadExtDtdDefaultValue               int
		XmlParserDebugEntities                  int
		XmlPedanticParserDefaultValue           int
		XmlSaveNoEmptyTags                      int
		XmlIndentTreeOutput                     int
		XmlTreeIndentString                     *Char
		XmlRegisterNodeDefaultValue             RegisterNodeFunc
		XmlDeregisterNodeDefaultValue           DeregisterNodeFunc
		XmlMallocAtomic                         MallocFunc
		XmlLastError                            Error
		XmlParserInputBufferCreateFilenameValue ParserInputBufferCreateFilenameFunc
		XmlOutputBufferCreateFilenameValue      OutputBufferCreateFilenameFunc
		XmlStructuredErrorContext               *Void
	}

	Enumeration struct {
		Next *Enumeration
		Name *Char
	}

	Node struct {
		_          *Void
		Type       ElementType
		Name       *Char
		Children   *Node
		Last       *Node
		Parent     *Node
		Next       *Node
		Prev       *Node
		Doc        *Doc
		Ns         *Ns
		Content    *Char
		Properties *Attr
		NsDef      *Ns
		Psvi       *Void
		Line       UnsignedShort
		Extra      UnsignedShort
	}

	Doc struct {
		_           *Void
		Type        ElementType
		Name        *Char
		Children    *Node
		Last        *Node
		Parent      *Node
		Next        *Node
		Prev        *Node
		Doc         *Doc
		Compression int
		Standalone  int
		IntSubset   *Dtd
		ExtSubset   *Dtd
		OldNs       *Ns
		Version     *Char
		Encoding    *Char
		Ids         *Void
		Refs        *Void
		URL         *Char
		Charset     int
		Dict        *Dict
		Psvi        *Void
		ParseFlags  int
		Properties  int
	}

	Entity struct {
		_          *Void
		Type       ElementType
		Name       *Char
		Children   *Node
		Last       *Node
		Parent     *Dtd
		Next       *Node
		Prev       *Node
		Doc        *Doc
		Orig       *Char
		Content    *Char
		Length     int
		Etype      EntityType
		ExternalID *Char
		SystemID   *Char
		Nexte      *Entity
		URI        *Char
		Owner      int
		Checked    int
	}

	ElementContent struct {
		Type   ElementContentType
		Ocur   ElementContentOccur
		Name   *Char
		C1     *ElementContent
		C2     *ElementContent
		Parent *ElementContent
		Prefix *Char
	}

	Attr struct {
		_        *Void
		Type     ElementType
		Name     *Char
		Children *Node
		Last     *Node
		Parent   *Node
		Next     *Attr
		Prev     *Attr
		Doc      *Doc
		Ns       *Ns
		Atype    AttributeType
		Psvi     *Void
	}

	Ns struct {
		Next    *Ns
		Type    NsType
		Href    *Char
		Prefix  *Char
		_       *Void
		Context *Doc
	}

	NodeSet struct {
		NodeNr  int
		NodeMax int
		NodeTab **Node
	}

	XPathObject struct {
		Type       XPathObjectType
		Nodesetval *NodeSet
		Boolval    int
		Floatval   Double
		Stringval  *Char
		User       *Void
		Index      int
		User2      *Void
		Index2     int
	}

	XPathContext struct {
		Doc                *Doc
		Node               *Node
		NbVariablesUnused  int
		MaxVariablesUnused int
		VarHash            *HashTable
		NbTypes            int
		MaxTypes           int
		Types              *XPathType
		NbFuncsUnused      int
		MaxFuncsUnused     int
		FuncHash           *HashTable
		NbAxis             int
		MaxAxis            int
		Axis               *XPathAxis
		Namespaces         **Ns
		NsNr               int
		User               *Void
		ContextSize        int
		ProximityPosition  int
		Xptr               int
		Here               *Node
		Origin             *Node
		NsHash             *HashTable
		VarLookupFunc      XPathVariableLookupFunc
		VarLookupData      *Void
		Extra              *Void
		Function           *Char
		FunctionURI        *Char
		FuncLookupFunc     XPathFuncLookupFunc
		FuncLookupData     *Void
		TmpNsList          **Ns
		TmpNsNr            int
		UserData           *Void
		Error              StructuredErrorFunc
		LastError          Error
		DebugNode          *Node
		Dict               *Dict
		Flags              int
		Cache              *Void
	}

	XPathType struct {
		Name *Char
		Func XPathConvertFunc
	}

	XPathAxis struct {
		Name *Char
		Func XPathAxisFunc
	}

	LocationSet struct {
		LocNr  int
		LocMax int
		LocTab **XPathObject
	}

	SchemaWildcard struct {
		Type            SchemaTypeType
		Id              *Char
		Annot           *SchemaAnnot
		Node            *Node
		MinOccurs       int
		MaxOccurs       int
		ProcessContents int
		Any             int
		NsSet           *SchemaWildcardNs
		NegNsSet        *SchemaWildcardNs
		Flags           int
	}

	SchemaAnnot struct {
		Next    *SchemaAnnot
		Content *Node
	}

	SchemaWildcardNs struct {
		Next  *SchemaWildcardNs
		Value *Char
	}

	ChRangeGroup struct {
		NbShortRange int
		NbLongRange  int
		ShortRange   *ChSRange
		LongRange    *ChLRange
	}

	ChSRange struct {
		Low  UnsignedShort
		High UnsignedShort
	}

	ChLRange struct {
		Low  UnsignedInt
		High UnsignedInt
	}

	ShellCtxt struct {
		Filename *Char
		Doc      *Doc
		Node     *Node
		Pctxt    *XPathContext
		Loaded   int
		Output   *FILE
		Input    ShellReadlineFunc
	}

	URI struct {
		Scheme    *Char
		Opaque    *Char
		Authority *Char
		Server    *Char
		User      *Char
		Port      int
		Path      *Char
		Query     *Char
		Fragment  *Char
		Cleanup   int
		QueryRaw  *Char
	}
)

type EntityType Enum

const (
	INTERNAL_GENERAL_ENTITY EntityType = iota + 1
	EXTERNAL_GENERAL_PARSED_ENTITY
	EXTERNAL_GENERAL_UNPARSED_ENTITY
	INTERNAL_PARAMETER_ENTITY
	EXTERNAL_PARAMETER_ENTITY
	INTERNAL_PREDEFINED_ENTITY
)

type ElementType Enum

const (
	ELEMENT_NODE ElementType = iota + 1
	ATTRIBUTE_NODE
	TEXT_NODE
	CDATA_SECTION_NODE
	ENTITY_REF_NODE
	ENTITY_NODE
	PI_NODE
	COMMENT_NODE
	DOCUMENT_NODE
	DOCUMENT_TYPE_NODE
	DOCUMENT_FRAG_NODE
	NOTATION_NODE
	HTML_DOCUMENT_NODE
	DTD_NODE
	ELEMENT_DECL
	ATTRIBUTE_DECL
	ENTITY_DECL
	NAMESPACE_DECL
	XINCLUDE_START
	XINCLUDE_END
	DOCB_DOCUMENT_NODE
)

type BufferAllocationScheme Enum

const (
	BUFFER_ALLOC_DOUBLEIT BufferAllocationScheme = iota
	BUFFER_ALLOC_EXACT
	BUFFER_ALLOC_IMMUTABLE
	BUFFER_ALLOC_IO
	BUFFER_ALLOC_HYBRID
)

type ElementContentType Enum

const (
	ELEMENT_CONTENT_PCDATA ElementContentType = iota + 1
	ELEMENT_CONTENT_ELEMENT
	ELEMENT_CONTENT_SEQ
	ELEMENT_CONTENT_OR
)

type ElementContentOccur Enum

const (
	ELEMENT_CONTENT_ONCE ElementContentOccur = iota + 1
	ELEMENT_CONTENT_OPT
	ELEMENT_CONTENT_MULT
	ELEMENT_CONTENT_PLUS
)

type CharEncoding Enum

const (
	CHAR_ENCODING_ERROR CharEncoding = iota - 1
	CHAR_ENCODING_NONE
	CHAR_ENCODING_UTF8
	CHAR_ENCODING_UTF16LE
	CHAR_ENCODING_UTF16BE
	CHAR_ENCODING_UCS4LE
	CHAR_ENCODING_UCS4BE
	CHAR_ENCODING_EBCDIC
	CHAR_ENCODING_UCS4_2143
	CHAR_ENCODING_UCS4_3412
	CHAR_ENCODING_UCS2
	CHAR_ENCODING_8859_1
	CHAR_ENCODING_8859_2
	CHAR_ENCODING_8859_3
	CHAR_ENCODING_8859_4
	CHAR_ENCODING_8859_5
	CHAR_ENCODING_8859_6
	CHAR_ENCODING_8859_7
	CHAR_ENCODING_8859_8
	CHAR_ENCODING_8859_9
	CHAR_ENCODING_2022_JP
	CHAR_ENCODING_SHIFT_JIS
	CHAR_ENCODING_EUC_JP
	CHAR_ENCODING_ASCII
)

type AttributeType Enum

const (
	ATTRIBUTE_CDATA AttributeType = iota + 1
	ATTRIBUTE_ID
	ATTRIBUTE_IDREF
	ATTRIBUTE_IDREFS
	ATTRIBUTE_ENTITY
	ATTRIBUTE_ENTITIES
	ATTRIBUTE_NMTOKEN
	ATTRIBUTE_NMTOKENS
	ATTRIBUTE_ENUMERATION
	ATTRIBUTE_NOTATION
)

type ErrorLevel Enum

const (
	ERR_NONE ErrorLevel = iota
	ERR_WARNING
	ERR_ERROR
	ERR_FATAL
)

type ParserInputState Enum

const (
	PARSER_EOF ParserInputState = iota - 1
	PARSER_START
	PARSER_MISC
	PARSER_PI
	PARSER_DTD
	PARSER_PROLOG
	PARSER_COMMENT
	PARSER_START_TAG
	PARSER_CONTENT
	PARSER_CDATA_SECTION
	PARSER_END_TAG
	PARSER_ENTITY_DECL
	PARSER_ENTITY_VALUE
	PARSER_ATTRIBUTE_VALUE
	PARSER_SYSTEM_LITERAL
	PARSER_EPILOG
	PARSER_IGNORE
	PARSER_PUBLIC_LITERAL
)

type ParserMode Enum

const (
	PARSE_UNKNOWN ParserMode = iota
	PARSE_DOM
	PARSE_SAX
	PARSE_PUSH_DOM
	PARSE_PUSH_SAX
	PARSE_READER
)

type HtmlStatus Enum

const (
	HTML_INVALID HtmlStatus = 1 << iota
	HTML_DEPRECATED
	HTML_VALID
	htmlStatusAnon
	HTML_REQUIRED            = HTML_VALID | htmlStatusAnon
	HTML_NA       HtmlStatus = 0
)

type XlinkShow Enum

const (
	XLINK_SHOW_NONE XlinkShow = iota
	XLINK_SHOW_NEW
	XLINK_SHOW_EMBED
	XLINK_SHOW_REPLACE
)

type XlinkActuate Enum

const (
	XLINK_ACTUATE_NONE XlinkActuate = iota
	XLINK_ACTUATE_AUTO
	XLINK_ACTUATE_ONREQUEST
)

type XlinkType Enum

const (
	XLINK_TYPE_NONE XlinkType = iota
	XLINK_TYPE_SIMPLE
	XLINK_TYPE_EXTENDED
	XLINK_TYPE_EXTENDED_SET
)

type ElementTypeVal Enum

const (
	ELEMENT_TYPE_UNDEFINED ElementTypeVal = iota
	ELEMENT_TYPE_EMPTY
	ELEMENT_TYPE_ANY
	ELEMENT_TYPE_MIXED
	ELEMENT_TYPE_ELEMENT
)

type AttributeDefault Enum

const (
	ATTRIBUTE_NONE AttributeDefault = iota + 1
	ATTRIBUTE_REQUIRED
	ATTRIBUTE_IMPLIED
	ATTRIBUTE_FIXED
)

type Feature Enum

const (
	WITH_THREAD Feature = iota + 1
	WITH_TREE
	WITH_OUTPUT
	WITH_PUSH
	WITH_READER
	WITH_PATTERN
	WITH_WRITER
	WITH_SAX1
	WITH_FTP
	WITH_HTTP
	WITH_VALID
	WITH_HTML
	WITH_LEGACY
	WITH_C14N
	WITH_CATALOG
	WITH_XPATH
	WITH_XPTR
	WITH_XINCLUDE
	WITH_ICONV
	WITH_ISO8859X
	WITH_UNICODE
	WITH_REGEXP
	WITH_AUTOMATA
	WITH_EXPR
	WITH_SCHEMAS
	WITH_SCHEMATRON
	WITH_MODULES
	WITH_DEBUG
	WITH_DEBUG_MEM
	WITH_DEBUG_RUN
	WITH_ZLIB
	WITH_ICU
	WITH_LZMA
	WITH_NONE Feature = 99999
)

type ParserErrors Enum

const (
	ERR_OK ParserErrors = iota
	ERR_INTERNAL_ERROR
	ERR_NO_MEMORY
	ERR_DOCUMENT_START
	ERR_DOCUMENT_EMPTY
	ERR_DOCUMENT_END
	ERR_INVALID_HEX_CHARREF
	ERR_INVALID_DEC_CHARREF
	ERR_INVALID_CHARREF
	ERR_INVALID_CHAR
	ERR_CHARREF_AT_EOF
	ERR_CHARREF_IN_PROLOG
	ERR_CHARREF_IN_EPILOG
	ERR_CHARREF_IN_DTD
	ERR_ENTITYREF_AT_EOF
	ERR_ENTITYREF_IN_PROLOG
	ERR_ENTITYREF_IN_EPILOG
	ERR_ENTITYREF_IN_DTD
	ERR_PEREF_AT_EOF
	ERR_PEREF_IN_PROLOG
	ERR_PEREF_IN_EPILOG
	ERR_PEREF_IN_INT_SUBSET
	ERR_ENTITYREF_NO_NAME
	ERR_ENTITYREF_SEMICOL_MISSING
	ERR_PEREF_NO_NAME
	ERR_PEREF_SEMICOL_MISSING
	ERR_UNDECLARED_ENTITY
	WAR_UNDECLARED_ENTITY
	ERR_UNPARSED_ENTITY
	ERR_ENTITY_IS_EXTERNAL
	ERR_ENTITY_IS_PARAMETER
	ERR_UNKNOWN_ENCODING
	ERR_UNSUPPORTED_ENCODING
	ERR_STRING_NOT_STARTED
	ERR_STRING_NOT_CLOSED
	ERR_NS_DECL_ERROR
	ERR_ENTITY_NOT_STARTED
	ERR_ENTITY_NOT_FINISHED
	ERR_LT_IN_ATTRIBUTE
	ERR_ATTRIBUTE_NOT_STARTED
	ERR_ATTRIBUTE_NOT_FINISHED
	ERR_ATTRIBUTE_WITHOUT_VALUE
	ERR_ATTRIBUTE_REDEFINED
	ERR_LITERAL_NOT_STARTED
	ERR_LITERAL_NOT_FINISHED
	ERR_COMMENT_NOT_FINISHED
	ERR_PI_NOT_STARTED
	ERR_PI_NOT_FINISHED
	ERR_NOTATION_NOT_STARTED
	ERR_NOTATION_NOT_FINISHED
	ERR_ATTLIST_NOT_STARTED
	ERR_ATTLIST_NOT_FINISHED
	ERR_MIXED_NOT_STARTED
	ERR_MIXED_NOT_FINISHED
	ERR_ELEMCONTENT_NOT_STARTED
	ERR_ELEMCONTENT_NOT_FINISHED
	ERR_XMLDECL_NOT_STARTED
	ERR_XMLDECL_NOT_FINISHED
	ERR_CONDSEC_NOT_STARTED
	ERR_CONDSEC_NOT_FINISHED
	ERR_EXT_SUBSET_NOT_FINISHED
	ERR_DOCTYPE_NOT_FINISHED
	ERR_MISPLACED_CDATA_END
	ERR_CDATA_NOT_FINISHED
	ERR_RESERVED_NAME
	ERR_SPACE_REQUIRED
	ERR_SEPARATOR_REQUIRED
	ERR_NMTOKEN_REQUIRED
	ERR_NAME_REQUIRED
	ERR_PCDATA_REQUIRED
	ERR_URI_REQUIRED
	ERR_PUBID_REQUIRED
	ERR_LT_REQUIRED
	ERR_GT_REQUIRED
	ERR_LTSLASH_REQUIRED
	ERR_EQUAL_REQUIRED
	ERR_TAG_NAME_MISMATCH
	ERR_TAG_NOT_FINISHED
	ERR_STANDALONE_VALUE
	ERR_ENCODING_NAME
	ERR_HYPHEN_IN_COMMENT
	ERR_INVALID_ENCODING
	ERR_EXT_ENTITY_STANDALONE
	ERR_CONDSEC_INVALID
	ERR_VALUE_REQUIRED
	ERR_NOT_WELL_BALANCED
	ERR_EXTRA_CONTENT
	ERR_ENTITY_CHAR_ERROR
	ERR_ENTITY_PE_INTERNAL
	ERR_ENTITY_LOOP
	ERR_ENTITY_BOUNDARY
	ERR_INVALID_URI
	ERR_URI_FRAGMENT
	WAR_CATALOG_PI
	ERR_NO_DTD
	ERR_CONDSEC_INVALID_KEYWORD
	ERR_VERSION_MISSING
	WAR_UNKNOWN_VERSION
	WAR_LANG_VALUE
	WAR_NS_URI
	WAR_NS_URI_RELATIVE
	ERR_MISSING_ENCODING
	WAR_SPACE_VALUE
	ERR_NOT_STANDALONE
	ERR_ENTITY_PROCESSING
	ERR_NOTATION_PROCESSING
	WAR_NS_COLUMN
	WAR_ENTITY_REDEFINED
	ERR_UNKNOWN_VERSION
	ERR_VERSION_MISMATCH
	ERR_NAME_TOO_LONG
)
const (
	NS_ERR_NAMESPACE ParserErrors = iota + 200
	NS_ERR_UNDEFINED_NAMESPACE
	NS_ERR_QNAME
	NS_ERR_ATTRIBUTE_REDEFINED
	NS_ERR_EMPTY
	NS_ERR_COLON
)
const (
	DTD_ATTRIBUTE_DEFAULT ParserErrors = iota + 500
	DTD_ATTRIBUTE_REDEFINED
	DTD_ATTRIBUTE_VALUE
	DTD_CONTENT_ERROR
	DTD_CONTENT_MODEL
	DTD_CONTENT_NOT_DETERMINIST
	DTD_DIFFERENT_PREFIX
	DTD_ELEM_DEFAULT_NAMESPACE
	DTD_ELEM_NAMESPACE
	DTD_ELEM_REDEFINED
	DTD_EMPTY_NOTATION
	DTD_ENTITY_TYPE
	DTD_ID_FIXED
	DTD_ID_REDEFINED
	DTD_ID_SUBSET
	DTD_INVALID_CHILD
	DTD_INVALID_DEFAULT
	DTD_LOAD_ERROR
	DTD_MISSING_ATTRIBUTE
	DTD_MIXED_CORRUPT
	DTD_MULTIPLE_ID
	DTD_NO_DOC
	DTD_NO_DTD
	DTD_NO_ELEM_NAME
	DTD_NO_PREFIX
	DTD_NO_ROOT
	DTD_NOTATION_REDEFINED
	DTD_NOTATION_VALUE
	DTD_NOT_EMPTY
	DTD_NOT_PCDATA
	DTD_NOT_STANDALONE
	DTD_ROOT_NAME
	DTD_STANDALONE_WHITE_SPACE
	DTD_UNKNOWN_ATTRIBUTE
	DTD_UNKNOWN_ELEM
	DTD_UNKNOWN_ENTITY
	DTD_UNKNOWN_ID
	DTD_UNKNOWN_NOTATION
	DTD_STANDALONE_DEFAULTED
	DTD_XMLID_VALUE
	DTD_XMLID_TYPE
	DTD_DUP_TOKEN
)
const (
	HTML_STRUCURE_ERROR ParserErrors = iota + 800
	HTML_UNKNOWN_TAG
)
const (
	RNGP_ANYNAME_ATTR_ANCESTOR ParserErrors = iota + 1000
	RNGP_ATTR_CONFLICT
	RNGP_ATTRIBUTE_CHILDREN
	RNGP_ATTRIBUTE_CONTENT
	RNGP_ATTRIBUTE_EMPTY
	RNGP_ATTRIBUTE_NOOP
	RNGP_CHOICE_CONTENT
	RNGP_CHOICE_EMPTY
	RNGP_CREATE_FAILURE
	RNGP_DATA_CONTENT
	RNGP_DEF_CHOICE_AND_INTERLEAVE
	RNGP_DEFINE_CREATE_FAILED
	RNGP_DEFINE_EMPTY
	RNGP_DEFINE_MISSING
	RNGP_DEFINE_NAME_MISSING
	RNGP_ELEM_CONTENT_EMPTY
	RNGP_ELEM_CONTENT_ERROR
	RNGP_ELEMENT_EMPTY
	RNGP_ELEMENT_CONTENT
	RNGP_ELEMENT_NAME
	RNGP_ELEMENT_NO_CONTENT
	RNGP_ELEM_TEXT_CONFLICT
	RNGP_EMPTY
	RNGP_EMPTY_CONSTRUCT
	RNGP_EMPTY_CONTENT
	RNGP_EMPTY_NOT_EMPTY
	RNGP_ERROR_TYPE_LIB
	RNGP_EXCEPT_EMPTY
	RNGP_EXCEPT_MISSING
	RNGP_EXCEPT_MULTIPLE
	RNGP_EXCEPT_NO_CONTENT
	RNGP_EXTERNALREF_EMTPY
	RNGP_EXTERNAL_REF_FAILURE
	RNGP_EXTERNALREF_RECURSE
	RNGP_FORBIDDEN_ATTRIBUTE
	RNGP_FOREIGN_ELEMENT
	RNGP_GRAMMAR_CONTENT
	RNGP_GRAMMAR_EMPTY
	RNGP_GRAMMAR_MISSING
	RNGP_GRAMMAR_NO_START
	RNGP_GROUP_ATTR_CONFLICT
	RNGP_HREF_ERROR
	RNGP_INCLUDE_EMPTY
	RNGP_INCLUDE_FAILURE
	RNGP_INCLUDE_RECURSE
	RNGP_INTERLEAVE_ADD
	RNGP_INTERLEAVE_CREATE_FAILED
	RNGP_INTERLEAVE_EMPTY
	RNGP_INTERLEAVE_NO_CONTENT
	RNGP_INVALID_DEFINE_NAME
	RNGP_INVALID_URI
	RNGP_INVALID_VALUE
	RNGP_MISSING_HREF
	RNGP_NAME_MISSING
	RNGP_NEED_COMBINE
	RNGP_NOTALLOWED_NOT_EMPTY
	RNGP_NSNAME_ATTR_ANCESTOR
	RNGP_NSNAME_NO_NS
	RNGP_PARAM_FORBIDDEN
	RNGP_PARAM_NAME_MISSING
	RNGP_PARENTREF_CREATE_FAILED
	RNGP_PARENTREF_NAME_INVALID
	RNGP_PARENTREF_NO_NAME
	RNGP_PARENTREF_NO_PARENT
	RNGP_PARENTREF_NOT_EMPTY
	RNGP_PARSE_ERROR
	RNGP_PAT_ANYNAME_EXCEPT_ANYNAME
	RNGP_PAT_ATTR_ATTR
	RNGP_PAT_ATTR_ELEM
	RNGP_PAT_DATA_EXCEPT_ATTR
	RNGP_PAT_DATA_EXCEPT_ELEM
	RNGP_PAT_DATA_EXCEPT_EMPTY
	RNGP_PAT_DATA_EXCEPT_GROUP
	RNGP_PAT_DATA_EXCEPT_INTERLEAVE
	RNGP_PAT_DATA_EXCEPT_LIST
	RNGP_PAT_DATA_EXCEPT_ONEMORE
	RNGP_PAT_DATA_EXCEPT_REF
	RNGP_PAT_DATA_EXCEPT_TEXT
	RNGP_PAT_LIST_ATTR
	RNGP_PAT_LIST_ELEM
	RNGP_PAT_LIST_INTERLEAVE
	RNGP_PAT_LIST_LIST
	RNGP_PAT_LIST_REF
	RNGP_PAT_LIST_TEXT
	RNGP_PAT_NSNAME_EXCEPT_ANYNAME
	RNGP_PAT_NSNAME_EXCEPT_NSNAME
	RNGP_PAT_ONEMORE_GROUP_ATTR
	RNGP_PAT_ONEMORE_INTERLEAVE_ATTR
	RNGP_PAT_START_ATTR
	RNGP_PAT_START_DATA
	RNGP_PAT_START_EMPTY
	RNGP_PAT_START_GROUP
	RNGP_PAT_START_INTERLEAVE
	RNGP_PAT_START_LIST
	RNGP_PAT_START_ONEMORE
	RNGP_PAT_START_TEXT
	RNGP_PAT_START_VALUE
	RNGP_PREFIX_UNDEFINED
	RNGP_REF_CREATE_FAILED
	RNGP_REF_CYCLE
	RNGP_REF_NAME_INVALID
	RNGP_REF_NO_DEF
	RNGP_REF_NO_NAME
	RNGP_REF_NOT_EMPTY
	RNGP_START_CHOICE_AND_INTERLEAVE
	RNGP_START_CONTENT
	RNGP_START_EMPTY
	RNGP_START_MISSING
	RNGP_TEXT_EXPECTED
	RNGP_TEXT_HAS_CHILD
	RNGP_TYPE_MISSING
	RNGP_TYPE_NOT_FOUND
	RNGP_TYPE_VALUE
	RNGP_UNKNOWN_ATTRIBUTE
	RNGP_UNKNOWN_COMBINE
	RNGP_UNKNOWN_CONSTRUCT
	RNGP_UNKNOWN_TYPE_LIB
	RNGP_URI_FRAGMENT
	RNGP_URI_NOT_ABSOLUTE
	RNGP_VALUE_EMPTY
	RNGP_VALUE_NO_CONTENT
	RNGP_XmlNs_NAME
	RNGP_NS
)
const (
	XPATH_EXPRESSION_OK ParserErrors = iota + 1200
	XPATH_NUMBER_ERROR
	XPATH_UNFINISHED_LITERAL_ERROR
	XPATH_START_LITERAL_ERROR
	XPATH_VARIABLE_REF_ERROR
	XPATH_UNDEF_VARIABLE_ERROR
	XPATH_INVALID_PREDICATE_ERROR
	XPATH_EXPR_ERROR
	XPATH_UNCLOSED_ERROR
	XPATH_UNKNOWN_FUNC_ERROR
	XPATH_INVALID_OPERAND
	XPATH_INVALID_TYPE
	XPATH_INVALID_ARITY
	XPATH_INVALID_CTXT_SIZE
	XPATH_INVALID_CTXT_POSITION
	XPATH_MEMORY_ERROR
	XPTR_SYNTAX_ERROR
	XPTR_RESOURCE_ERROR
	XPTR_SUB_RESOURCE_ERROR
	XPATH_UNDEF_PREFIX_ERROR
	XPATH_ENCODING_ERROR
	XPATH_INVALID_CHAR_ERROR
)
const (
	TREE_INVALID_HEX ParserErrors = iota + 1300
	TREE_INVALID_DEC
	TREE_UNTERMINATED_ENTITY
	TREE_NOT_UTF8
)
const (
	SAVE_NOT_UTF8 ParserErrors = iota + 1400
	SAVE_CHAR_INVALID
	SAVE_NO_DOCTYPE
	SAVE_UNKNOWN_ENCODING
	REGEXP_COMPILE_ERROR ParserErrors = 1450
)
const (
	IO_UNKNOWN ParserErrors = iota + 1500
	IO_EACCES
	IO_EAGAIN
	IO_EBADF
	IO_EBADMSG
	IO_EBUSY
	IO_ECANCELED
	IO_ECHILD
	IO_EDEADLK
	IO_EDOM
	IO_EEXIST
	IO_EFAULT
	IO_EFBIG
	IO_EINPROGRESS
	IO_EINTR
	IO_EINVAL
	IO_EIO
	IO_EISDIR
	IO_EMFILE
	IO_EMLINK
	IO_EMSGSIZE
	IO_ENAMETOOLONG
	IO_ENFILE
	IO_ENODEV
	IO_ENOENT
	IO_ENOEXEC
	IO_ENOLCK
	IO_ENOMEM
	IO_ENOSPC
	IO_ENOSYS
	IO_ENOTDIR
	IO_ENOTEMPTY
	IO_ENOTSUP
	IO_ENOTTY
	IO_ENXIO
	IO_EPERM
	IO_EPIPE
	IO_ERANGE
	IO_EROFS
	IO_ESPIPE
	IO_ESRCH
	IO_ETIMEDOUT
	IO_EXDEV
	IO_NETWORK_ATTEMPT
	IO_ENCODER
	IO_FLUSH
	IO_WRITE
	IO_NO_INPUT
	IO_BUFFER_FULL
	IO_LOAD_ERROR
	IO_ENOTSOCK
	IO_EISCONN
	IO_ECONNREFUSED
	IO_ENETUNREACH
	IO_EADDRINUSE
	IO_EALREADY
	IO_EAFNOSUPPORT
)
const (
	XINCLUDE_RECURSION ParserErrors = iota + 1600
	XINCLUDE_PARSE_VALUE
	XINCLUDE_ENTITY_DEF_MISMATCH
	XINCLUDE_NO_HREF
	XINCLUDE_NO_FALLBACK
	XINCLUDE_HREF_URI
	XINCLUDE_TEXT_FRAGMENT
	XINCLUDE_TEXT_DOCUMENT
	XINCLUDE_INVALID_CHAR
	XINCLUDE_BUILD_FAILED
	XINCLUDE_UNKNOWN_ENCODING
	XINCLUDE_MULTIPLE_ROOT
	XINCLUDE_XPTR_FAILED
	XINCLUDE_XPTR_RESULT
	XINCLUDE_INCLUDE_IN_INCLUDE
	XINCLUDE_FALLBACKS_IN_INCLUDE
	XINCLUDE_FALLBACK_NOT_IN_INCLUDE
	XINCLUDE_DEPRECATED_NS
	XINCLUDE_FRAGMENT_ID
)
const (
	CATALOG_MISSING_ATTR ParserErrors = iota + 1650
	CATALOG_ENTRY_BROKEN
	CATALOG_PREFER_VALUE
	CATALOG_NOT_CATALOG
	CATALOG_RECURSION
)
const (
	SCHEMAP_PREFIX_UNDEFINED ParserErrors = iota + 1700
	SCHEMAP_ATTRFORMDEFAULT_VALUE
	SCHEMAP_ATTRGRP_NONAME_NOREF
	SCHEMAP_ATTR_NONAME_NOREF
	SCHEMAP_COMPLEXTYPE_NONAME_NOREF
	SCHEMAP_ELEMFORMDEFAULT_VALUE
	SCHEMAP_ELEM_NONAME_NOREF
	SCHEMAP_EXTENSION_NO_BASE
	SCHEMAP_FACET_NO_VALUE
	SCHEMAP_FAILED_BUILD_IMPORT
	SCHEMAP_GROUP_NONAME_NOREF
	SCHEMAP_IMPORT_NAMESPACE_NOT_URI
	SCHEMAP_IMPORT_REDEFINE_NSNAME
	SCHEMAP_IMPORT_SCHEMA_NOT_URI
	SCHEMAP_INVALID_BOOLEAN
	SCHEMAP_INVALID_ENUM
	SCHEMAP_INVALID_FACET
	SCHEMAP_INVALID_FACET_VALUE
	SCHEMAP_INVALID_MAXOCCURS
	SCHEMAP_INVALID_MINOCCURS
	SCHEMAP_INVALID_REF_AND_SUBTYPE
	SCHEMAP_INVALID_WHITE_SPACE
	SCHEMAP_NOATTR_NOREF
	SCHEMAP_NOTATION_NO_NAME
	SCHEMAP_NOTYPE_NOREF
	SCHEMAP_REF_AND_SUBTYPE
	SCHEMAP_RESTRICTION_NONAME_NOREF
	SCHEMAP_SIMPLETYPE_NONAME
	SCHEMAP_TYPE_AND_SUBTYPE
	SCHEMAP_UNKNOWN_ALL_CHILD
	SCHEMAP_UNKNOWN_ANYATTRIBUTE_CHILD
	SCHEMAP_UNKNOWN_ATTR_CHILD
	SCHEMAP_UNKNOWN_ATTRGRP_CHILD
	SCHEMAP_UNKNOWN_ATTRIBUTE_GROUP
	SCHEMAP_UNKNOWN_BASE_TYPE
	SCHEMAP_UNKNOWN_CHOICE_CHILD
	SCHEMAP_UNKNOWN_COMPLEXCONTENT_CHILD
	SCHEMAP_UNKNOWN_COMPLEXTYPE_CHILD
	SCHEMAP_UNKNOWN_ELEM_CHILD
	SCHEMAP_UNKNOWN_EXTENSION_CHILD
	SCHEMAP_UNKNOWN_FACET_CHILD
	SCHEMAP_UNKNOWN_FACET_TYPE
	SCHEMAP_UNKNOWN_GROUP_CHILD
	SCHEMAP_UNKNOWN_IMPORT_CHILD
	SCHEMAP_UNKNOWN_LIST_CHILD
	SCHEMAP_UNKNOWN_NOTATION_CHILD
	SCHEMAP_UNKNOWN_PROCESSCONTENT_CHILD
	SCHEMAP_UNKNOWN_REF
	SCHEMAP_UNKNOWN_RESTRICTION_CHILD
	SCHEMAP_UNKNOWN_SCHEMAS_CHILD
	SCHEMAP_UNKNOWN_SEQUENCE_CHILD
	SCHEMAP_UNKNOWN_SIMPLECONTENT_CHILD
	SCHEMAP_UNKNOWN_SIMPLETYPE_CHILD
	SCHEMAP_UNKNOWN_TYPE
	SCHEMAP_UNKNOWN_UNION_CHILD
	SCHEMAP_ELEM_DEFAULT_FIXED
	SCHEMAP_REGEXP_INVALID
	SCHEMAP_FAILED_LOAD
	SCHEMAP_NOTHING_TO_PARSE
	SCHEMAP_NOROOT
	SCHEMAP_REDEFINED_GROUP
	SCHEMAP_REDEFINED_TYPE
	SCHEMAP_REDEFINED_ELEMENT
	SCHEMAP_REDEFINED_ATTRGROUP
	SCHEMAP_REDEFINED_ATTR
	SCHEMAP_REDEFINED_NOTATION
	SCHEMAP_FAILED_PARSE
	SCHEMAP_UNKNOWN_PREFIX
	SCHEMAP_DEF_AND_PREFIX
	SCHEMAP_UNKNOWN_INCLUDE_CHILD
	SCHEMAP_INCLUDE_SCHEMA_NOT_URI
	SCHEMAP_INCLUDE_SCHEMA_NO_URI
	SCHEMAP_NOT_SCHEMA
	SCHEMAP_UNKNOWN_MEMBER_TYPE
	SCHEMAP_INVALID_ATTR_USE
	SCHEMAP_RECURSIVE
	SCHEMAP_SUPERNUMEROUS_LIST_ITEM_TYPE
	SCHEMAP_INVALID_ATTR_COMBINATION
	SCHEMAP_INVALID_ATTR_INLINE_COMBINATION
	SCHEMAP_MISSING_SIMPLETYPE_CHILD
	SCHEMAP_INVALID_ATTR_NAME
	SCHEMAP_REF_AND_CONTENT
	SCHEMAP_CT_PROPS_CORRECT_1
	SCHEMAP_CT_PROPS_CORRECT_2
	SCHEMAP_CT_PROPS_CORRECT_3
	SCHEMAP_CT_PROPS_CORRECT_4
	SCHEMAP_CT_PROPS_CORRECT_5
	SCHEMAP_DERIVATION_OK_RESTRICTION_1
	SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_1
	SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_2
	SCHEMAP_DERIVATION_OK_RESTRICTION_2_2
	SCHEMAP_DERIVATION_OK_RESTRICTION_3
	SCHEMAP_WILDCARD_INVALID_NS_MEMBER
	SCHEMAP_INTERSECTION_NOT_EXPRESSIBLE
	SCHEMAP_UNION_NOT_EXPRESSIBLE
	SCHEMAP_SRC_IMPORT_3_1
	SCHEMAP_SRC_IMPORT_3_2
	SCHEMAP_DERIVATION_OK_RESTRICTION_4_1
	SCHEMAP_DERIVATION_OK_RESTRICTION_4_2
	SCHEMAP_DERIVATION_OK_RESTRICTION_4_3
	SCHEMAP_COS_CT_EXTENDS_1_3
)
const (
	SCHEMAV_NOROOT ParserErrors = iota + 1801
	SCHEMAV_UNDECLAREDELEM
	SCHEMAV_NOTTOPLEVEL
	SCHEMAV_MISSING
	SCHEMAV_WRONGELEM
	SCHEMAV_NOTYPE
	SCHEMAV_NOROLLBACK
	SCHEMAV_ISABSTRACT
	SCHEMAV_NOTEMPTY
	SCHEMAV_ELEMCONT
	SCHEMAV_HAVEDEFAULT
	SCHEMAV_NOTNILLABLE
	SCHEMAV_EXTRACONTENT
	SCHEMAV_INVALIDATTR
	SCHEMAV_INVALIDELEM
	SCHEMAV_NOTDETERMINIST
	SCHEMAV_CONSTRUCT
	SCHEMAV_INTERNAL
	SCHEMAV_NOTSIMPLE
	SCHEMAV_ATTRUNKNOWN
	SCHEMAV_ATTRINVALID
	SCHEMAV_VALUE
	SCHEMAV_FACET
	SCHEMAV_CVC_DATATYPE_VALID_1_2_1
	SCHEMAV_CVC_DATATYPE_VALID_1_2_2
	SCHEMAV_CVC_DATATYPE_VALID_1_2_3
	SCHEMAV_CVC_TYPE_3_1_1
	SCHEMAV_CVC_TYPE_3_1_2
	SCHEMAV_CVC_FACET_VALID
	SCHEMAV_CVC_LENGTH_VALID
	SCHEMAV_CVC_MINLENGTH_VALID
	SCHEMAV_CVC_MAXLENGTH_VALID
	SCHEMAV_CVC_MININCLUSIVE_VALID
	SCHEMAV_CVC_MAXINCLUSIVE_VALID
	SCHEMAV_CVC_MINEXCLUSIVE_VALID
	SCHEMAV_CVC_MAXEXCLUSIVE_VALID
	SCHEMAV_CVC_TOTALDIGITS_VALID
	SCHEMAV_CVC_FRACTIONDIGITS_VALID
	SCHEMAV_CVC_PATTERN_VALID
	SCHEMAV_CVC_ENUMERATION_VALID
	SCHEMAV_CVC_COMPLEX_TYPE_2_1
	SCHEMAV_CVC_COMPLEX_TYPE_2_2
	SCHEMAV_CVC_COMPLEX_TYPE_2_3
	SCHEMAV_CVC_COMPLEX_TYPE_2_4
	SCHEMAV_CVC_ELT_1
	SCHEMAV_CVC_ELT_2
	SCHEMAV_CVC_ELT_3_1
	SCHEMAV_CVC_ELT_3_2_1
	SCHEMAV_CVC_ELT_3_2_2
	SCHEMAV_CVC_ELT_4_1
	SCHEMAV_CVC_ELT_4_2
	SCHEMAV_CVC_ELT_4_3
	SCHEMAV_CVC_ELT_5_1_1
	SCHEMAV_CVC_ELT_5_1_2
	SCHEMAV_CVC_ELT_5_2_1
	SCHEMAV_CVC_ELT_5_2_2_1
	SCHEMAV_CVC_ELT_5_2_2_2_1
	SCHEMAV_CVC_ELT_5_2_2_2_2
	SCHEMAV_CVC_ELT_6
	SCHEMAV_CVC_ELT_7
	SCHEMAV_CVC_ATTRIBUTE_1
	SCHEMAV_CVC_ATTRIBUTE_2
	SCHEMAV_CVC_ATTRIBUTE_3
	SCHEMAV_CVC_ATTRIBUTE_4
	SCHEMAV_CVC_COMPLEX_TYPE_3_1
	SCHEMAV_CVC_COMPLEX_TYPE_3_2_1
	SCHEMAV_CVC_COMPLEX_TYPE_3_2_2
	SCHEMAV_CVC_COMPLEX_TYPE_4
	SCHEMAV_CVC_COMPLEX_TYPE_5_1
	SCHEMAV_CVC_COMPLEX_TYPE_5_2
	SCHEMAV_ELEMENT_CONTENT
	SCHEMAV_DOCUMENT_ELEMENT_MISSING
	SCHEMAV_CVC_COMPLEX_TYPE_1
	SCHEMAV_CVC_AU
	SCHEMAV_CVC_TYPE_1
	SCHEMAV_CVC_TYPE_2
	SCHEMAV_CVC_IDC
	SCHEMAV_CVC_WILDCARD
	SCHEMAV_MISC
)
const (
	XPTR_UNKNOWN_SCHEME ParserErrors = iota + 1900
	XPTR_CHILDSEQ_START
	XPTR_EVAL_FAILED
	XPTR_EXTRA_OBJECTS
)
const (
	C14N_CREATE_CTXT ParserErrors = iota + 1950
	C14N_REQUIRES_UTF8
	C14N_CREATE_STACK
	C14N_INVALID_NODE
	C14N_UNKNOW_NODE
	C14N_RELATIVE_NAMESPACE
)
const (
	FTP_PASV_ANSWER ParserErrors = iota + 2000
	FTP_EPSV_ANSWER
	FTP_ACCNT
	FTP_URL_SYNTAX
)
const (
	HTTP_URL_SYNTAX ParserErrors = iota + 2020
	HTTP_USE_IP
	HTTP_UNKNOWN_HOST
)
const (
	SCHEMAP_SRC_SIMPLE_TYPE_1 ParserErrors = iota + 3000
	SCHEMAP_SRC_SIMPLE_TYPE_2
	SCHEMAP_SRC_SIMPLE_TYPE_3
	SCHEMAP_SRC_SIMPLE_TYPE_4
	SCHEMAP_SRC_RESOLVE
	SCHEMAP_SRC_RESTRICTION_BASE_OR_SIMPLETYPE
	SCHEMAP_SRC_LIST_ITEMTYPE_OR_SIMPLETYPE
	SCHEMAP_SRC_UNION_MEMBERTYPES_OR_SIMPLETYPES
	SCHEMAP_ST_PROPS_CORRECT_1
	SCHEMAP_ST_PROPS_CORRECT_2
	SCHEMAP_ST_PROPS_CORRECT_3
	SCHEMAP_COS_ST_RESTRICTS_1_1
	SCHEMAP_COS_ST_RESTRICTS_1_2
	SCHEMAP_COS_ST_RESTRICTS_1_3_1
	SCHEMAP_COS_ST_RESTRICTS_1_3_2
	SCHEMAP_COS_ST_RESTRICTS_2_1
	SCHEMAP_COS_ST_RESTRICTS_2_3_1_1
	SCHEMAP_COS_ST_RESTRICTS_2_3_1_2
	SCHEMAP_COS_ST_RESTRICTS_2_3_2_1
	SCHEMAP_COS_ST_RESTRICTS_2_3_2_2
	SCHEMAP_COS_ST_RESTRICTS_2_3_2_3
	SCHEMAP_COS_ST_RESTRICTS_2_3_2_4
	SCHEMAP_COS_ST_RESTRICTS_2_3_2_5
	SCHEMAP_COS_ST_RESTRICTS_3_1
	SCHEMAP_COS_ST_RESTRICTS_3_3_1
	SCHEMAP_COS_ST_RESTRICTS_3_3_1_2
	SCHEMAP_COS_ST_RESTRICTS_3_3_2_2
	SCHEMAP_COS_ST_RESTRICTS_3_3_2_1
	SCHEMAP_COS_ST_RESTRICTS_3_3_2_3
	SCHEMAP_COS_ST_RESTRICTS_3_3_2_4
	SCHEMAP_COS_ST_RESTRICTS_3_3_2_5
	SCHEMAP_COS_ST_DERIVED_OK_2_1
	SCHEMAP_COS_ST_DERIVED_OK_2_2
	SCHEMAP_S4S_ELEM_NOT_ALLOWED
	SCHEMAP_S4S_ELEM_MISSING
	SCHEMAP_S4S_ATTR_NOT_ALLOWED
	SCHEMAP_S4S_ATTR_MISSING
	SCHEMAP_S4S_ATTR_INVALID_VALUE
	SCHEMAP_SRC_ELEMENT_1
	SCHEMAP_SRC_ELEMENT_2_1
	SCHEMAP_SRC_ELEMENT_2_2
	SCHEMAP_SRC_ELEMENT_3
	SCHEMAP_P_PROPS_CORRECT_1
	SCHEMAP_P_PROPS_CORRECT_2_1
	SCHEMAP_P_PROPS_CORRECT_2_2
	SCHEMAP_E_PROPS_CORRECT_2
	SCHEMAP_E_PROPS_CORRECT_3
	SCHEMAP_E_PROPS_CORRECT_4
	SCHEMAP_E_PROPS_CORRECT_5
	SCHEMAP_E_PROPS_CORRECT_6
	SCHEMAP_SRC_INCLUDE
	SCHEMAP_SRC_ATTRIBUTE_1
	SCHEMAP_SRC_ATTRIBUTE_2
	SCHEMAP_SRC_ATTRIBUTE_3_1
	SCHEMAP_SRC_ATTRIBUTE_3_2
	SCHEMAP_SRC_ATTRIBUTE_4
	SCHEMAP_NO_XmlNs
	SCHEMAP_NO_XSI
	SCHEMAP_COS_VALID_DEFAULT_1
	SCHEMAP_COS_VALID_DEFAULT_2_1
	SCHEMAP_COS_VALID_DEFAULT_2_2_1
	SCHEMAP_COS_VALID_DEFAULT_2_2_2
	SCHEMAP_CVC_SIMPLE_TYPE
	SCHEMAP_COS_CT_EXTENDS_1_1
	SCHEMAP_SRC_IMPORT_1_1
	SCHEMAP_SRC_IMPORT_1_2
	SCHEMAP_SRC_IMPORT_2
	SCHEMAP_SRC_IMPORT_2_1
	SCHEMAP_SRC_IMPORT_2_2
	SCHEMAP_INTERNAL
	SCHEMAP_NOT_DETERMINISTIC
	SCHEMAP_SRC_ATTRIBUTE_GROUP_1
	SCHEMAP_SRC_ATTRIBUTE_GROUP_2
	SCHEMAP_SRC_ATTRIBUTE_GROUP_3
	SCHEMAP_MG_PROPS_CORRECT_1
	SCHEMAP_MG_PROPS_CORRECT_2
	SCHEMAP_SRC_CT_1
	SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_3
	SCHEMAP_AU_PROPS_CORRECT_2
	SCHEMAP_A_PROPS_CORRECT_2
	SCHEMAP_C_PROPS_CORRECT
	SCHEMAP_SRC_REDEFINE
	SCHEMAP_SRC_IMPORT
	SCHEMAP_WARN_SKIP_SCHEMA
	SCHEMAP_WARN_UNLOCATED_SCHEMA
	SCHEMAP_WARN_ATTR_REDECL_PROH
	SCHEMAP_WARN_ATTR_POINTLESS_PROH
	SCHEMAP_AG_PROPS_CORRECT
	SCHEMAP_COS_CT_EXTENDS_1_2
	SCHEMAP_AU_PROPS_CORRECT
	SCHEMAP_A_PROPS_CORRECT_3
	SCHEMAP_COS_ALL_LIMITED
)
const (
	SCHEMATRONV_ASSERT ParserErrors = iota + 4000
	SCHEMATRONV_REPORT
)
const (
	MODULE_OPEN ParserErrors = iota + 4900
	MODULE_CLOSE
)
const (
	CHECK_FOUND_ELEMENT ParserErrors = iota + 5000
	CHECK_FOUND_ATTRIBUTE
	CHECK_FOUND_TEXT
	CHECK_FOUND_CDATA
	CHECK_FOUND_ENTITYREF
	CHECK_FOUND_ENTITY
	CHECK_FOUND_PI
	CHECK_FOUND_COMMENT
	CHECK_FOUND_DOCTYPE
	CHECK_FOUND_FRAGMENT
	CHECK_FOUND_NOTATION
	CHECK_UNKNOWN_NODE
	CHECK_ENTITY_TYPE
	CHECK_NO_PARENT
	CHECK_NO_DOC
	CHECK_NO_NAME
	CHECK_NO_ELEM
	CHECK_WRONG_DOC
	CHECK_NO_PREV
	CHECK_WRONG_PREV
	CHECK_NO_NEXT
	CHECK_WRONG_NEXT
	CHECK_NOT_DTD
	CHECK_NOT_ATTR
	CHECK_NOT_ATTR_DECL
	CHECK_NOT_ELEM_DECL
	CHECK_NOT_ENTITY_DECL
	CHECK_NOT_NS_DECL
	CHECK_NO_HREF
	CHECK_WRONG_PARENT
	CHECK_NS_SCOPE
	CHECK_NS_ANCESTOR
	CHECK_NOT_UTF8
	CHECK_NO_DICT
	CHECK_NOT_NCNAME
	CHECK_OUTSIDE_DICT
	CHECK_WRONG_NAME
	CHECK_NAME_NOT_NULL
)
const (
	I18N_NO_NAME ParserErrors = iota + 6000
	I18N_NO_HANDLER
	I18N_EXCESS_HANDLER
	I18N_CONV_FAILED
	I18N_NO_OUTPUT
	BUF_OVERFLOW ParserErrors = 7000
)

type XPathObjectType Enum

const (
	XPATH_UNDEFINED XPathObjectType = iota
	XPATH_NODESET
	XPATH_BOOLEAN
	XPATH_NUMBER
	XPATH_STRING
	XPATH_POINT
	XPATH_RANGE
	XPATH_LOCATIONSET
	XPATH_USERS
	XPATH_XSLT_TREE
)

type SchemaTypeType Enum

const (
	SCHEMA_TYPE_BASIC SchemaTypeType = iota + 1
	SCHEMA_TYPE_ANY
	SCHEMA_TYPE_FACET
	SCHEMA_TYPE_SIMPLE
	SCHEMA_TYPE_COMPLEX
	SCHEMA_TYPE_SEQUENCE
	SCHEMA_TYPE_CHOICE
	SCHEMA_TYPE_ALL
	SCHEMA_TYPE_SIMPLE_CONTENT
	SCHEMA_TYPE_COMPLEX_CONTENT
	SCHEMA_TYPE_UR
	SCHEMA_TYPE_RESTRICTION
	SCHEMA_TYPE_EXTENSION
	SCHEMA_TYPE_ELEMENT
	SCHEMA_TYPE_ATTRIBUTE
	SCHEMA_TYPE_ATTRIBUTEGROUP
	SCHEMA_TYPE_GROUP
	SCHEMA_TYPE_NOTATION
	SCHEMA_TYPE_LIST
	SCHEMA_TYPE_UNION
	SCHEMA_TYPE_ANY_ATTRIBUTE
	SCHEMA_TYPE_IDC_UNIQUE
	SCHEMA_TYPE_IDC_KEY
	SCHEMA_TYPE_IDC_KEYREF
	SCHEMA_TYPE_PARTICLE
	SCHEMA_TYPE_ATTRIBUTE_USE
)
const (
	SCHEMA_FACET_MININCLUSIVE SchemaTypeType = 1000
	SCHEMA_FACET_MINEXCLUSIVE
	SCHEMA_FACET_MAXINCLUSIVE
	SCHEMA_FACET_MAXEXCLUSIVE
	SCHEMA_FACET_TOTALDIGITS
	SCHEMA_FACET_FRACTIONDIGITS
	SCHEMA_FACET_PATTERN
	SCHEMA_FACET_ENUMERATION
	SCHEMA_FACET_WHITESPACE
	SCHEMA_FACET_LENGTH
	SCHEMA_FACET_MAXLENGTH
	SCHEMA_FACET_MINLENGTH
)
const (
	SCHEMA_EXTRA_QNAMEREF SchemaTypeType = 2000
	SCHEMA_EXTRA_ATTR_USE_PROHIB
)

type SchemaWhitespaceValueType Enum

const (
	SCHEMA_WHITESPACE_UNKNOWN SchemaWhitespaceValueType = iota
	SCHEMA_WHITESPACE_PRESERVE
	SCHEMA_WHITESPACE_REPLACE
	SCHEMA_WHITESPACE_COLLAPSE
)

type SchemaValType Enum

const (
	SCHEMAS_UNKNOWN SchemaValType = iota
	SCHEMAS_STRING
	SCHEMAS_NORMSTRING
	SCHEMAS_DECIMAL
	SCHEMAS_TIME
	SCHEMAS_GDAY
	SCHEMAS_GMONTH
	SCHEMAS_GMONTHDAY
	SCHEMAS_GYEAR
	SCHEMAS_GYEARMONTH
	SCHEMAS_DATE
	SCHEMAS_DATETIME
	SCHEMAS_DURATION
	SCHEMAS_FLOAT
	SCHEMAS_DOUBLE
	SCHEMAS_BOOLEAN
	SCHEMAS_TOKEN
	SCHEMAS_LANGUAGE
	SCHEMAS_NMTOKEN
	SCHEMAS_NMTOKENS
	SCHEMAS_NAME
	SCHEMAS_QNAME
	SCHEMAS_NCNAME
	SCHEMAS_ID
	SCHEMAS_IDREF
	SCHEMAS_IDREFS
	SCHEMAS_ENTITY
	SCHEMAS_ENTITIES
	SCHEMAS_NOTATION
	SCHEMAS_ANYURI
	SCHEMAS_INTEGER
	SCHEMAS_NPINTEGER
	SCHEMAS_NINTEGER
	SCHEMAS_NNINTEGER
	SCHEMAS_PINTEGER
	SCHEMAS_INT
	SCHEMAS_UINT
	SCHEMAS_LONG
	SCHEMAS_ULONG
	SCHEMAS_SHORT
	SCHEMAS_USHORT
	SCHEMAS_BYTE
	SCHEMAS_UBYTE
	SCHEMAS_HEXBINARY
	SCHEMAS_BASE64BINARY
	SCHEMAS_ANYTYPE
	SCHEMAS_ANYSIMPLETYPE
)

type ParserSeverities Enum

const (
	PARSER_SEVERITY_VALIDITY_WARNING ParserSeverities = iota + 1
	PARSER_SEVERITY_VALIDITY_ERROR
	PARSER_SEVERITY_WARNING
	PARSER_SEVERITY_ERROR
)

type CatalogPrefer Enum

const (
	CATA_PREFER_NONE CatalogPrefer = iota
	CATA_PREFER_PUBLIC
	CATA_PREFER_SYSTEM
)

type CatalogAllow Enum

const (
	CATA_ALLOW_NONE CatalogAllow = iota
	CATA_ALLOW_GLOBAL
	CATA_ALLOW_DOCUMENT
	CATA_ALLOW_ALL
)

type ModuleOption Enum

const (
	MODULE_LAZY ModuleOption = iota + 1
	MODULE_LOCAL
)

type ParserOption Enum

const (
	PARSE_RECOVER ParserOption = 1 << iota
	PARSE_NOENT
	PARSE_DTDLOAD
	PARSE_DTDATTR
	PARSE_DTDVALID
	PARSE_NOERROR
	PARSE_NOWARNING
	PARSE_PEDANTIC
	PARSE_NOBLANKS
	PARSE_SAX1
	PARSE_XINCLUDE
	PARSE_NONET
	PARSE_NODICT
	PARSE_NSCLEAN
	PARSE_NOCDATA
	PARSE_NOXINCNODE
	PARSE_COMPACT
	PARSE_OLD10
	PARSE_NOBASEFIX
	PARSE_HUGE
	PARSE_OLDSAX
	PARSE_IGNORE_ENC
	PARSE_BIG_LINES
)

var (
	CheckVersion func(version int)

	Strdup func(cur string) string

	Strndup func(cur string, leng int) string

	CharStrndup func(cur string, leng int) string

	CharStrdup func(cur string) string

	Strsub func(str string, start int, leng int) string

	Strchr func(str string, val Char) string

	Strstr func(str, val string) string

	Strcasestr func(str, val string) string

	Strcmp func(str1, str2 string) int

	Strncmp func(str1, str2 string, leng int) int

	Strcasecmp func(str1, str2 string) int

	Strncasecmp func(str1, str2 string, leng int) int

	StrEqual func(str1, str2 string) int

	StrQEqual func(pref, name, str string) int

	Strlen func(str string) int

	Strcat func(cur, add string) string

	Strncat func(cur, add string, leng int) string

	StrncatNew func(str1, str2 string, leng int) string

	StrPrintf func(
		buf string, leng int, msg string, v ...VArg) int

	StrVPrintf func(
		buf string, leng int, msg string, ap VAList) int

	GetUTF8Char func(utf *UnsignedChar, leng *int) int

	CheckUTF8 func(utf *UnsignedChar) int

	UTF8Strsize func(utf string, leng int) int

	UTF8Strndup func(utf string, leng int) string

	UTF8Strpos func(utf string, pos int) string

	UTF8Strloc func(utf string, utfchar string) int

	UTF8Strsub func(utf string, start, leng int) string

	UTF8Strlen func(utf string) int

	UTF8Size func(utf string) int

	UTF8Charcmp func(utf1, utf2 string) int

	BufContent func(buf *Buf) string

	BufEnd func(buf *Buf) string

	BufUse func(buf *Buf) SizeT

	BufShrink func(buf *Buf, leng SizeT) SizeT

	InitializeDict func() int

	DictCreate func() *Dict

	DictSetLimit func(dict *Dict, limit SizeT) SizeT

	DictGetUsage func(dict *Dict) SizeT

	DictCreateSub func(sub *Dict) *Dict

	DictReference func(dict *Dict) int

	DictFree func(dict *Dict)

	DictLookup func(
		dict *Dict, name string, leng int) string

	DictExists func(
		dict *Dict, name string, leng int) string

	DictQLookup func(
		dict *Dict, prefix, name string) string

	DictOwns func(dict *Dict, str string) int

	DictSize func(dict *Dict) int

	DictCleanup func()

	RegexpCompile func(regexp string) *Regexp

	RegFreeRegexp func(regexp *Regexp)

	RegexpExec func(comp *Regexp, value string) int

	RegexpPrint func(output *FILE, regexp *Regexp)

	RegexpIsDeterminist func(comp *Regexp) int

	RegNewExecCtxt func(
		comp *Regexp, callback RegExecCallbacks,
		data *Void) *RegExecCtxt

	RegFreeExecCtxt func(exec *RegExecCtxt)

	RegExecPushString func(
		exec *RegExecCtxt, value string, data *Void) int

	RegExecPushString2 func(exec *RegExecCtxt,
		value, value2 string, data *Void) int

	RegExecNextValues func(exec *RegExecCtxt,
		nbval, nbneg *int, values *string, terminal *int) int

	RegExecErrInfo func(exec *RegExecCtxt, string *string,
		nbval, nbneg *int, values *string, terminal *int) int

	ExpFreeCtxt func(ctxt *ExpCtxt)

	ExpNewCtxt func(maxNodes int, dict *Dict) *ExpCtxt

	ExpCtxtNbNodes func(ctxt *ExpCtxt) int

	ExpCtxtNbCons func(ctxt *ExpCtxt) int

	ExpFree func(ctxt *ExpCtxt, expr *ExpNode)

	ExpRef func(expr *ExpNode)

	ExpParse func(ctxt *ExpCtxt, expr string) *ExpNode

	ExpNewAtom func(ctxt *ExpCtxt,
		name string, leng int) *ExpNode

	ExpNewOr func(
		ctxt *ExpCtxt, left, right *ExpNode) *ExpNode

	ExpNewSeq func(
		ctxt *ExpCtxt, left, right *ExpNode) *ExpNode

	ExpNewRange func(ctxt *ExpCtxt, subset *ExpNode,
		min, max int) *ExpNode

	ExpIsNillable func(expr *ExpNode) int

	ExpMaxToken func(expr *ExpNode) int

	ExpGetLanguage func(ctxt *ExpCtxt, expr *ExpNode,
		langList *string, leng int) int

	ExpGetStart func(ctxt *ExpCtxt, expr *ExpNode,
		tokList *string, leng int) int

	ExpStringDerive func(ctxt *ExpCtxt, expr *ExpNode,
		str string, leng int) *ExpNode

	ExpExpDerive func(
		ctxt *ExpCtxt, expr, sub *ExpNode) *ExpNode

	ExpSubsume func(
		ctxt *ExpCtxt, expr, sub *ExpNode) int

	ExpDump func(buf *Buffer, expr *ExpNode)

	ValidateNCName func(value string, space int) int

	ValidateQName func(value string, space int) int

	ValidateName func(value string, space int) int

	ValidateNMToken func(value string, space int) int

	BuildQName func(
		ncname, prefix, memory string, leng int) string

	SplitQName2 func(name string, prefix *string) string

	SplitQName3 func(name string, leng *int) string

	SetBufferAllocationScheme func(
		scheme BufferAllocationScheme)

	GetBufferAllocationScheme func() BufferAllocationScheme

	BufferCreate func() *Buffer

	BufferCreateSize func(size SizeT) *Buffer

	BufferCreateStatic func(
		mem *Void, size SizeT) *Buffer

	BufferResize func(buf *Buffer, size UnsignedInt) int

	BufferFree func(buf *Buffer)

	BufferDump func(file *FILE, buf *Buffer) int

	BufferAdd func(buf *Buffer, str string, leng int) int

	BufferAddHead func(buf *Buffer, str string, leng int) int

	BufferCat func(buf *Buffer, str string) int

	BufferCCat func(buf *Buffer, str string) int

	BufferShrink func(buf *Buffer, leng UnsignedInt) int

	BufferGrow func(buf *Buffer, leng UnsignedInt) int

	BufferEmpty func(buf *Buffer)

	BufferContent func(buf *Buffer) string

	BufferDetach func(buf *Buffer) string

	BufferSetAllocationScheme func(
		buf *Buffer, scheme BufferAllocationScheme)

	BufferLength func(buf *Buffer) int

	CreateIntSubset func(
		doc *Doc, name, externalID, systemID string) *Dtd

	NewDtd func(
		doc *Doc, name, externalID, systemID string) *Dtd

	GetIntSubset func(doc *Doc) *Dtd

	FreeDtd func(cur *Dtd)

	NewGlobalNs func(doc *Doc, href, prefix string) *Ns

	NewNs func(node *Node, href, prefix string) *Ns

	FreeNs func(cur *Ns)

	FreeNsList func(
		cur *Ns)

	NewDoc func(version string) *Doc

	FreeDoc func(cur *Doc)

	NewDocProp func(doc *Doc, name, value string) *Attr

	NewProp func(node *Node, name, value string) *Attr

	NewNsProp func(
		node *Node, ns *Ns, name, value string) *Attr

	NewNsPropEatName func(
		node *Node, ns *Ns, name, value string) *Attr

	FreePropList func(cur *Attr)

	FreeProp func(cur *Attr)

	CopyProp func(target *Node, cur *Attr) *Attr

	CopyPropList func(target *Node, cur *Attr) *Attr

	CopyDtd func(dtd *Dtd) *Dtd

	CopyDoc func(doc *Doc, recursive int) *Doc

	NewDocNode func(
		doc *Doc, ns *Ns, name, content string) *Node

	NewDocNodeEatName func(
		doc *Doc, ns *Ns, name, content string) *Node

	NewNode func(ns *Ns, name string) *Node

	NewNodeEatName func(ns *Ns, name string) *Node

	NewChild func(parent *Node,
		ns *Ns, name, content string) *Node

	NewDocText func(doc *Doc, content string) *Node

	NewText func(content string) *Node

	NewDocPI func(doc *Doc, name, content string) *Node

	NewPI func(name, content string) *Node

	NewDocTextLen func(
		doc *Doc, content string, leng int) *Node

	NewTextLen func(content string, leng int) *Node

	NewDocComment func(doc *Doc, content string) *Node

	NewComment func(content string) *Node

	NewCDataBlock func(
		doc *Doc, content string, leng int) *Node

	NewCharRef func(doc *Doc, name string) *Node

	NewReference func(doc *Doc, name string) *Node

	CopyNode func(node *Node, recursive int) *Node

	DocCopyNode func(
		node *Node, doc *Doc, recursive int) *Node

	DocCopyNodeList func(doc *Doc, node *Node) *Node

	CopyNodeList func(node *Node) *Node

	NewTextChild func(parent *Node,
		ns *Ns, name, content string) *Node

	NewDocRawNode func(
		doc *Doc, ns *Ns, name, content string) *Node

	NewDocFragment func(doc *Doc) *Node

	GetLineNo func(node *Node) Long

	GetNodePath func(node *Node) string

	DocGetRootElement func(doc *Doc) *Node

	GetLastChild func(parent *Node) *Node

	NodeIsText func(node *Node) int

	IsBlankNode func(node *Node) int

	DocSetRootElement func(
		doc *Doc, root *Node) *Node

	NodeSetName func(cur *Node, name string)

	AddChild func(parent, cur *Node) *Node

	AddChildList func(parent, cur *Node) *Node

	ReplaceNode func(old, cur *Node) *Node

	AddPrevSibling func(cur, elem *Node) *Node

	AddSibling func(cur, elem *Node) *Node

	AddNextSibling func(cur, elem *Node) *Node

	UnlinkNode func(cur *Node)

	TextMerge func(first, second *Node) *Node

	TextConcat func(
		node *Node, content string, leng int) int

	FreeNodeList func(cur *Node)

	FreeNode func(cur *Node)

	SetTreeDoc func(tree *Node, doc *Doc)

	SetListDoc func(list *Node, doc *Doc)

	SearchNs func(
		doc *Doc, node *Node, nameSpace string) *Ns

	SearchNsByHref func(
		doc *Doc, node *Node, href string) *Ns

	GetNsList func(doc *Doc, node *Node) **Ns

	SetNs func(node *Node, ns *Ns)

	CopyNamespace func(cur *Ns) *Ns

	CopyNamespaceList func(cur *Ns) *Ns

	SetProp func(node *Node, name, value string) *Attr

	SetNsProp func(
		node *Node, ns *Ns, name, value string) *Attr

	GetNoNsProp func(node *Node, name string) string

	GetProp func(node *Node, name string) string

	HasProp func(node *Node, name string) *Attr

	HasNsProp func(
		node *Node, name, nameSpace string) *Attr

	GetNsProp func(
		node *Node, name, nameSpace string) string

	StringGetNodeList func(doc *Doc, value string) *Node

	StringLenGetNodeList func(
		doc *Doc, value string, leng int) *Node

	NodeListGetString func(
		doc *Doc, list *Node, inLine int) string

	NodeListGetRawString func(
		doc *Doc, list *Node, inLine int) string

	NodeSetContent func(cur *Node, content string)

	NodeSetContentLen func(
		cur *Node, content string, leng int)

	NodeAddContent func(cur *Node, content string)

	NodeAddContentLen func(
		cur *Node, content string, leng int)

	NodeGetContent func(cur *Node) string

	NodeBufGetContent func(
		buffer *Buffer, cur *Node) int

	BufGetNodeContent func(buf *Buf, cur *Node) int

	NodeGetLang func(cur *Node) string

	NodeGetSpacePreserve func(cur *Node) int

	NodeSetLang func(cur *Node, lang string)

	NodeSetSpacePreserve func(cur *Node, val int)

	NodeGetBase func(doc *Doc, cur *Node) string

	NodeSetBase func(cur *Node, uri string)

	RemoveProp func(cur *Attr) int

	UnsetNsProp func(
		node *Node, ns *Ns, name string) int

	UnsetProp func(node *Node, name string) int

	BufferWriteCHAR func(buf *Buffer, str string)

	BufferWriteChar func(buf *Buffer, str string)

	BufferWriteQuotedString func(
		buf *Buffer, str string)

	AttrSerializeTxtContent func(buf *Buffer,
		doc *Doc, attr *Attr, str string)

	ReconciliateNs func(doc *Doc, tree *Node) int

	DocDumpFormatMemory func(
		cur *Doc, mem *string, size *int, format int)

	DocDumpMemory func(cur *Doc, mem *string, size *int)

	DocDumpMemoryEnc func(outDoc *Doc, docTxtPtr *string,
		docTxtLen *int, txt_encoding string)

	DocDumpFormatMemoryEnc func(outDoc *Doc, docTxtPtr *string,
		docTxtLen *int, txt_encoding string, format int)

	DocFormatDump func(f *FILE, cur *Doc, format int) int

	DocDump func(f *FILE, cur *Doc) int

	ElemDump func(f *FILE, doc *Doc, cur *Node)

	SaveFile func(filename string, cur *Doc) int

	SaveFormatFile func(
		filename string, cur *Doc, format int) int

	BufNodeDump func(buf *Buf, doc *Doc,
		cur *Node, level, format int) SizeT

	NodeDump func(buf *Buffer, doc *Doc,
		cur *Node, level, format int) int

	SaveFileTo func(
		buf *OutputBuffer, cur *Doc, encoding string) int

	SaveFormatFileTo func(buf *OutputBuffer,
		cur *Doc, encoding string, format int) int

	NodeDumpOutput func(buf *OutputBuffer, doc *Doc,
		cur *Node, level, format int, encoding string)

	SaveFormatFileEnc func(filename string,
		cur *Doc, encoding string, format int) int

	SaveFileEnc func(filename string, cur *Doc, encoding string) int

	IsXHTML func(systemID, publicID string) int

	GetDocCompressMode func(doc *Doc) int

	SetDocCompressMode func(doc *Doc, mode int)

	GetCompressMode func() int

	SetCompressMode func(mode int)

	DOMWrapNewCtxt func() *DOMWrapCtxt

	DOMWrapFreeCtxt func(ctxt *DOMWrapCtxt)

	DOMWrapReconcileNamespaces func(
		ctxt *DOMWrapCtxt, elem *Node, options int) int

	DOMWrapAdoptNode func(ctxt *DOMWrapCtxt, sourceDoc *Doc,
		node *Node, destDoc *Doc,
		destParent *Node, options int) int

	DOMWrapRemoveNode func(ctxt *DOMWrapCtxt, doc *Doc,
		node *Node, options int) int

	DOMWrapCloneNode func(ctxt *DOMWrapCtxt, sourceDoc *Doc,
		node *Node, clonedNode **Node, destDoc *Doc,
		destParent *Node, deep, options int) int

	ChildElementCount func(parent *Node) UnsignedLong

	NextElementSibling func(node *Node) *Node

	FirstElementChild func(parent *Node) *Node

	LastElementChild func(parent *Node) *Node

	PreviousElementSibling func(node *Node) *Node

	MemSetup func(freeFunc FreeFunc, mallocFunc MallocFunc,
		reallocFunc ReallocFunc, strdupFunc StrdupFunc) int

	MemGet func(freeFunc *FreeFunc, mallocFunc *MallocFunc,
		reallocFunc *ReallocFunc, strdupFunc *StrdupFunc) int

	GcMemSetup func(freeFunc FreeFunc, mallocFunc MallocFunc,
		mallocAtomicFunc MallocFunc, reallocFunc ReallocFunc,
		strdupFunc StrdupFunc) int

	GcMemGet func(freeFunc *FreeFunc, mallocFunc *MallocFunc,
		mallocAtomicFunc *MallocFunc, reallocFunc *ReallocFunc,
		strdupFunc *StrdupFunc) int

	InitMemory func() int

	CleanupMemory func()

	MemUsed func() int

	MemBlocks func() int

	MemDisplay func(fp *FILE)

	MemDisplayLast func(fp *FILE, nbBytes Long)

	MemShow func(fp *FILE, nr int)

	MemoryDump func()

	MemMalloc func(size SizeT) *Void

	MemRealloc func(ptr *Void, size SizeT) *Void

	MemFree func(ptr *Void)

	MemoryStrdup func(str string) string

	MallocLoc func(size SizeT, file string, line int) *Void

	ReallocLoc func(
		ptr *Void, size SizeT, file string, line int) *Void

	MallocAtomicLoc func(
		size SizeT, file string, line int) *Void

	MemStrdupLoc func(str, file string, line int) string

	HashCreate func(size int) *HashTable

	HashCreateDict func(size int, dict *Dict) *HashTable

	HashFree func(table *HashTable, f HashDeallocator)

	HashAddEntry func(
		table *HashTable, name string, userdata *Void) int

	HashUpdateEntry func(table *HashTable,
		name string, userdata *Void, f HashDeallocator) int

	HashAddEntry2 func(
		table *HashTable, name, name2 string, userdata *Void) int

	HashUpdateEntry2 func(table *HashTable, name, name2 string,
		userdata *Void, f HashDeallocator) int

	HashAddEntry3 func(table *HashTable,
		name, name2, name3 string, userdata *Void) int

	HashUpdateEntry3 func(table *HashTable,
		name, name2, name3 string, userdata *Void,
		f HashDeallocator) int

	HashRemoveEntry func(
		table *HashTable, name string, f HashDeallocator) int

	HashRemoveEntry2 func(table *HashTable,
		name, name2 string, f HashDeallocator) int

	HashRemoveEntry3 func(table *HashTable,
		name, name2, name3 string, f HashDeallocator) int

	HashLookup func(table *HashTable, name string) *Void

	HashLookup2 func(table *HashTable, name, name2 string) *Void

	HashLookup3 func(table *HashTable,
		name, name2, name3 string) *Void

	HashQLookup func(table *HashTable, name, prefix string) *Void

	HashQLookup2 func(table *HashTable,
		name, prefix, name2, prefix2 string) *Void

	HashQLookup3 func(table *HashTable,
		name, prefix, name2, prefix2, name3,
		prefix3 string) *Void

	HashCopy func(table *HashTable, f HashCopier) *HashTable

	HashSize func(table *HashTable) int

	HashScan func(table *HashTable, f HashScanner, data *Void)

	HashScan3 func(table *HashTable,
		name, name2, name3 string, f HashScanner, data *Void)

	HashScanFull func(
		table *HashTable, f HashScannerFull, data *Void)

	HashScanFull3 func(table *HashTable,
		name, name2, name3 string, f HashScannerFull, data *Void)

	SetGenericErrorFunc func(ctx *Void, handler GenericErrorFunc)

	InitGenericErrorDefaultFunc func(handler *GenericErrorFunc)

	SetStructuredErrorFunc func(
		ctx *Void, handler StructuredErrorFunc)

	ParserError func(ctx *Void, msg string, v ...VArg)

	ParserWarning func(ctx *Void, msg string, v ...VArg)

	ParserValidityError func(ctx *Void, msg string, v ...VArg)

	ParserValidityWarning func(ctx *Void, msg string, v ...VArg)

	ParserPrintFileInfo func(input *ParserInput)

	ParserPrintFileContext func(input *ParserInput)

	GetLastError func() *Error

	ResetLastError func()

	CtxtGetLastError func(ctx *Void) *Error

	CtxtResetLastError func(ctx *Void)

	ResetError func(err *Error)

	CopyError func(from, to *Error) int

	ListCreate func(deallocator ListDeallocator,
		compare ListDataCompare) *List

	ListDelete func(l *List)

	ListSearch func(l *List, data *Void) *Void

	ListReverseSearch func(l *List, data *Void) *Void

	ListInsert func(l *List, data *Void) int

	ListAppend func(l *List, data *Void) int

	ListRemoveFirst func(l *List, data *Void) int

	ListRemoveLast func(l *List, data *Void) int

	ListRemoveAll func(l *List, data *Void) int

	ListClear func(l *List)

	ListEmpty func(l *List) int

	ListFront func(l *List) *Link

	ListEnd func(l *List) *Link

	ListSize func(l *List) int

	ListPopFront func(l *List)

	ListPopBack func(l *List)

	ListPushFront func(l *List, data *Void) int

	ListPushBack func(l *List, data *Void) int

	ListReverse func(l *List)

	ListSort func(l *List)

	ListWalk func(l *List, walker ListWalker, user *Void)

	ListReverseWalk func(l *List, walker ListWalker, user *Void)

	ListMerge func(l1, l2 *List)

	ListDup func(old *List) *List

	ListCopy func(cur, old *List) int

	LinkGetData func(lk *Link) *Void

	NewAutomata func() *Automata

	FreeAutomata func(am *Automata)

	AutomataGetInitState func(am *Automata) *AutomataState

	AutomataSetFinalState func(
		am *Automata, state *AutomataState) int

	AutomataNewState func(am *Automata) *AutomataState

	AutomataNewTransition func(am *Automata,
		from, to *AutomataState, token string,
		data *Void) *AutomataState

	AutomataNewTransition2 func(am *Automata,
		from, to *AutomataState, token, token2 string,
		data *Void) *AutomataState

	AutomataNewNegTrans func(am *Automata,
		from, to *AutomataState, token, token2 string,
		data *Void) *AutomataState

	AutomataNewCountTrans func(am *Automata,
		from, to *AutomataState, token string,
		min, max int, data *Void) *AutomataState

	AutomataNewCountTrans2 func(am *Automata,
		from, to *AutomataState, token, token2 string,
		min, max int, data *Void) *AutomataState

	AutomataNewOnceTrans func(am *Automata,
		from, to *AutomataState, token string,
		min, max int, data *Void) *AutomataState

	AutomataNewOnceTrans2 func(am *Automata,
		from, to *AutomataState, token, token2 string,
		min, max int, data *Void) *AutomataState

	AutomataNewAllTrans func(am *Automata,
		from, to *AutomataState, lax int) *AutomataState

	AutomataNewEpsilon func(am *Automata,
		from, to *AutomataState) *AutomataState

	AutomataNewCountedTrans func(am *Automata,
		from, to *AutomataState, counter int) *AutomataState

	AutomataNewCounterTrans func(am *Automata,
		from, to *AutomataState, counter int) *AutomataState

	AutomataNewCounter func(am *Automata, min, max int) int

	AutomataCompile func(am *Automata) *Regexp

	AutomataIsDeterminist func(am *Automata) int

	AddNotationDecl func(ctxt *ValidCtxt, dtd *Dtd,
		name, oublicID, systemID string) *Notation

	CopyNotationTable func(
		table *NotationTable) *NotationTable

	FreeNotationTable func(table *NotationTable)

	DumpNotationDecl func(buf *Buffer, nota *Notation)

	DumpNotationTable func(buf *Buffer, table *NotationTable)

	NewElementContent func(name string,
		t ElementContentType) *ElementContent

	CopyElementContent func(
		content *ElementContent) *ElementContent

	FreeElementContent func(cur *ElementContent)

	NewDocElementContent func(doc *Doc,
		name string, t ElementContentType) *ElementContent

	CopyDocElementContent func(doc *Doc,
		content *ElementContent) *ElementContent

	FreeDocElementContent func(doc *Doc, cur *ElementContent)

	SnprintfElementContent func(buf string, size int,
		content *ElementContent, englob int)

	SprintfElementContent func(
		buf string, content *ElementContent, englob int)

	AddElementDecl func(ctxt *ValidCtxt, dtd *Dtd, name string,
		t ElementTypeVal, content *ElementContent) *Element

	CopyElementTable func(table *ElementTable) *ElementTable

	FreeElementTable func(table *ElementTable)

	DumpElementTable func(buf *Buffer, table *ElementTable)

	DumpElementDecl func(buf *Buffer, elem *Element)

	CreateEnumeration func(name string) *Enumeration

	FreeEnumeration func(cur *Enumeration)

	CopyEnumeration func(cur *Enumeration) *Enumeration

	AddAttributeDecl func(ctxt *ValidCtxt, dtd *Dtd,
		elem, name, ns string, t AttributeType,
		def AttributeDefault, defaultValue string,
		tree *Enumeration) *AttributeStruct

	CopyAttributeTable func(
		table *AttributeTable) *AttributeTable

	FreeAttributeTable func(table *AttributeTable)

	DumpAttributeTable func(buf *Buffer, table *AttributeTable)

	DumpAttributeDecl func(buf *Buffer, attr *AttributeStruct)

	AddID func(ctxt *ValidCtxt, doc *Doc,
		value string, attr *Attr) *ID

	FreeIDTable func(table *IDTable)

	GetID func(doc *Doc, ID string) *Attr

	IsID func(doc *Doc, elem *Node, attr *Attr) int

	RemoveID func(doc *Doc, attr *Attr) int

	AddRef func(ctxt *ValidCtxt, doc *Doc,
		value string, attr *Attr) *Ref

	FreeRefTable func(table *RefTable)

	IsRef func(doc *Doc, elem *Node, attr *Attr) int

	RemoveRef func(doc *Doc, attr *Attr) int

	GetRefs func(doc *Doc, ID string) *List

	NewValidCtxt func() *ValidCtxt

	FreeValidCtxt func(*ValidCtxt)

	ValidateRoot func(ctxt *ValidCtxt, doc *Doc) int

	ValidateElementDecl func(ctxt *ValidCtxt,
		doc *Doc, elem *Element) int

	ValidNormalizeAttributeValue func(doc *Doc,
		elem *Node, name string, value string) string

	ValidCtxtNormalizeAttributeValue func(ctxt *ValidCtxt,
		doc *Doc, elem *Node, name, value string) string

	ValidateAttributeDecl func(ctxt *ValidCtxt,
		doc *Doc, attr *AttributeStruct) int

	ValidateAttributeValue func(
		t AttributeType, value string) int

	ValidateNotationDecl func(ctxt *ValidCtxt,
		doc *Doc, nota *Notation) int

	ValidateDtd func(ctxt *ValidCtxt, doc *Doc, dtd *Dtd) int

	ValidateDtdFinal func(ctxt *ValidCtxt, doc *Doc) int

	ValidateDocument func(ctxt *ValidCtxt, doc *Doc) int

	ValidateElement func(
		ctxt *ValidCtxt, doc *Doc, elem *Node) int

	ValidateOneElement func(
		ctxt *ValidCtxt, doc *Doc, elem *Node) int

	ValidateOneAttribute func(ctxt *ValidCtxt, doc *Doc,
		elem *Node, attr *Attr, value string) int

	ValidateOneNamespace func(ctxt *ValidCtxt, doc *Doc,
		elem *Node, prefix string, ns *Ns, value string) int

	ValidateDocumentFinal func(ctxt *ValidCtxt, doc *Doc) int

	ValidateNotationUse func(
		ctxt *ValidCtxt, doc *Doc, notationName string) int

	IsMixedElement func(doc *Doc, name string) int

	GetDtdAttrDesc func(
		dtd *Dtd, elem, name string) *AttributeStruct

	GetDtdQAttrDesc func(
		dtd *Dtd, elem, name, prefix string) *AttributeStruct

	GetDtdNotationDesc func(dtd *Dtd, name string) *Notation

	GetDtdQElementDesc func(
		dtd *Dtd, name, prefix string) *Element

	GetDtdElementDesc func(dtd *Dtd, name string) *Element

	ValidGetPotentialChildren func(ctree *ElementContent,
		names *string, leng *int, max int) int

	ValidGetValidElements func(
		prev, next *Node, names *string, max int) int

	ValidateNameValue func(value string) int

	ValidateNamesValue func(value string) int

	ValidateNmtokenValue func(value string) int

	ValidateNmtokensValue func(value string) int

	ValidBuildContentModel func(
		ctxt *ValidCtxt, elem *Element) int

	ValidatePushElement func(
		ctxt *ValidCtxt, doc *Doc, elem *Node, qname string) int

	ValidatePushCData func(
		ctxt *ValidCtxt, data string, leng int) int

	ValidatePopElement func(
		ctxt *ValidCtxt, doc *Doc, elem *Node, qname string) int

	InitializePredefinedEntities func()

	NewEntity func(doc *Doc, name string, typ int,
		externalID, systemID, content string) *Entity

	AddDocEntity func(doc *Doc, name string, typ int,
		externalID, systemID, content string) *Entity

	AddDtdEntity func(doc *Doc, name string, typ int,
		externalID, systemID, content string) *Entity

	GetPredefinedEntity func(name string) *Entity

	GetDocEntity func(doc *Doc, name string) *Entity

	GetDtdEntity func(doc *Doc, name string) *Entity

	XmlGetParameterEntity func(doc *Doc, name string) *Entity

	EncodeEntities func(doc *Doc, input string) string

	EncodeEntitiesReentrant func(doc *Doc, input string) string

	EncodeSpecialChars func(doc *Doc, input string) string

	CreateEntitiesTable func() *EntitiesTable

	CopyEntitiesTable func(table *EntitiesTable) *EntitiesTable

	FreeEntitiesTable func(table *EntitiesTable)

	DumpEntitiesTable func(buf *Buffer, table *EntitiesTable)

	DumpEntityDecl func(buf *Buffer, ent *Entity)

	CleanupPredefinedEntities func()

	InitCharEncodingHandlers func()

	CleanupCharEncodingHandlers func()

	RegisterCharEncodingHandler func(
		handler *CharEncodingHandler)

	GetCharEncodingHandler func(
		enc CharEncoding) *CharEncodingHandler

	FindCharEncodingHandler func(
		name string) *CharEncodingHandler

	NewCharEncodingHandler func(name string,
		input CharEncodingInputFunc,
		output CharEncodingOutputFunc) *CharEncodingHandler

	AddEncodingAlias func(name, alias string) int

	DelEncodingAlias func(alias string) int

	GetEncodingAlias func(alias string) string

	CleanupEncodingAliases func()

	ParseCharEncoding func(name string) CharEncoding

	GetCharEncodingName func(enc CharEncoding) string

	DetectCharEncoding func(
		in *UnsignedChar, leng int) CharEncoding

	CharEncOutFunc func(
		handler *CharEncodingHandler, out, in *Buffer) int

	CharEncInFunc func(
		handler *CharEncodingHandler, out, in *Buffer) int

	CharEncFirstLine func(
		handler *CharEncodingHandler, out, in *Buffer) int

	CharEncCloseFunc func(handler CharEncodingHandler) int

	UTF8Toisolat1 func(out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int) int

	Isolat1ToUTF8 func(out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int) int

	CleanupInputCallbacks func()

	PopInputCallbacks func() int

	RegisterDefaultInputCallbacks func()

	AllocParserInputBuffer func(
		enc CharEncoding) *ParserInputBuffer

	ParserInputBufferCreateFile func(
		file *FILE, enc CharEncoding) *ParserInputBuffer

	ParserInputBufferCreateFd func(
		fd int, enc CharEncoding) *ParserInputBuffer

	ParserInputBufferCreateMem func(mem string,
		size int, enc CharEncoding) *ParserInputBuffer

	ParserInputBufferCreateStatic func(mem string,
		size int, enc CharEncoding) *ParserInputBuffer

	ParserInputBufferCreateIO func(
		ioread InputReadCallback, ioclose InputCloseCallback,
		ioctx *Void, enc CharEncoding) *ParserInputBuffer

	ParserInputBufferRead func(
		in *ParserInputBuffer, leng int) int

	ParserInputBufferGrow func(
		in *ParserInputBuffer, leng int) int

	ParserInputBufferPush func(
		in *ParserInputBuffer, leng int, buf string) int

	FreeParserInputBuffer func(in *ParserInputBuffer)

	ParserGetDirectory func(filename string) string

	RegisterInputCallbacks func(matchFunc InputMatchCallback,
		openFunc InputOpenCallback, readFunc InputReadCallback,
		closeFunc InputCloseCallback) int

	ParserInputBufferCreateFilename func(URI string,
		enc CharEncoding) *ParserInputBuffer

	CleanupOutputCallbacks func()

	RegisterDefaultOutputCallbacks func()

	AllocOutputBuffer func(
		encoder *CharEncodingHandler) *OutputBuffer

	OutputBufferCreateFilename func(
		URI string, encoder *CharEncodingHandler,
		compression int) *OutputBuffer

	OutputBufferCreateFile func(
		file *FILE, encoder *CharEncodingHandler) *OutputBuffer

	OutputBufferCreateBuffer func(buffer *Buffer,
		encoder *CharEncodingHandler) *OutputBuffer

	OutputBufferCreateFd func(
		fd int, encoder *CharEncodingHandler) *OutputBuffer

	OutputBufferCreateIO func(iowrite OutputWriteCallback,
		ioclose OutputCloseCallback, ioctx *Void,
		encoder *CharEncodingHandler) *OutputBuffer

	OutputBufferGetContent func(out *OutputBuffer) string

	OutputBufferGetSize func(out *OutputBuffer) SizeT

	OutputBufferWrite func(
		out *OutputBuffer, leng int, buf string) int

	OutputBufferWriteString func(
		out *OutputBuffer, str string) int

	OutputBufferWriteEscape func(out *OutputBuffer,
		str string, escaping CharEncodingOutputFunc) int

	OutputBufferFlush func(out *OutputBuffer) int

	OutputBufferClose func(out *OutputBuffer) int

	RegisterOutputCallbacks func(matchFunc OutputMatchCallback,
		openFunc OutputOpenCallback,
		writeFunc OutputWriteCallback,
		closeFunc OutputCloseCallback) int

	RegisterHTTPPostCallbacks func()

	CheckHTTPInput func(
		ctxt *ParserCtxt, ret *ParserInput) *ParserInput

	NoNetExternalEntityLoader func(
		URL, ID string, ctxt *ParserCtxt) *ParserInput

	NormalizeWindowsPath func(path string) string

	CheckFilename func(path string) int

	FileMatch func(filename string) int

	FileOpen func(filename string) *Void

	FileRead func(context *Void, buffer string, leng int) int

	FileClose func(context *Void) int

	IOHTTPMatch func(filename string) int

	IOHTTPOpen func(filename string) *Void

	IOHTTPOpenW func(postUri string, compression int) *Void

	IOHTTPRead func(context *Void, buffer string, leng int) int

	IOHTTPClose func(context *Void) int

	IOFTPMatch func(filename string) int

	IOFTPOpen func(filename string) *Void

	IOFTPRead func(context *Void, buffer string, leng int) int

	IOFTPClose func(context *Void) int

	InitParser func()

	CleanupParser func()

	ParserInputRead func(in *ParserInput, leng int) int

	ParserInputGrow func(in *ParserInput, leng int) int

	ParseDoc func(cur string) *Doc

	ParseFile func(filename string) *Doc

	ParseMemory func(buffer string, size int) *Doc

	SubstituteEntitiesDefault func(val int) int

	KeepBlanksDefault func(val int) int

	StopParser func(ctxt *ParserCtxt)

	PedanticParserDefault func(val int) int

	LineNumbersDefault func(val int) int

	RecoverDoc func(cur string) *Doc

	RecoverMemory func(buffer string, size int) *Doc

	RecoverFile func(filename string) *Doc

	ParseDocument func(ctxt *ParserCtxt) int

	ParseExtParsedEnt func(ctxt *ParserCtxt) int

	SAXUserParseFile func(
		sax *SAXHandler, userData *Void, filename string) int

	SAXUserParseMemory func(sax *SAXHandler,
		userData *Void, buffer string, size int) int

	SAXParseDoc func(
		sax *SAXHandler, cur string, recovery int) *Doc

	SAXParseMemory func(sax *SAXHandler,
		buffer string, size int, recovery int) *Doc

	SAXParseMemoryWithData func(sax *SAXHandler,
		buffer string, size, recovery int, data *Void) *Doc

	SAXParseFile func(
		sax *SAXHandler, filename string, recovery int) *Doc

	SAXParseFileWithData func(sax *SAXHandler,
		filename string, recovery int, data *Void) *Doc

	SAXParseEntity func(sax *SAXHandler, filename string) *Doc

	ParseEntity func(filename string) *Doc

	SAXParseDTD func(
		sax *SAXHandler, externalID, systemID string) *Dtd

	ParseDTD func(externalID, systemID string) *Dtd

	IOParseDTD func(sax *SAXHandler,
		input *ParserInputBuffer, enc CharEncoding) *Dtd

	ParseBalancedChunkMemory func(doc *Doc, sax *SAXHandler,
		userData *Void, depth int, str string, lst **Node) int

	ParseInNodeContext func(node *Node, data string,
		datalen, options int, lst **Node) ParserErrors

	ParseBalancedChunkMemoryRecover func(doc *Doc,
		sax *SAXHandler, userData *Void, depth int,
		str string, lst **Node, recover int) int

	ParseExternalEntity func(doc *Doc,
		sax *SAXHandler, userData *Void, depth int,
		URL, ID string, lst **Node) int

	ParseCtxtExternalEntity func(
		ctx *ParserCtxt, URL, ID string, lst **Node) int

	NewParserCtxt func() *ParserCtxt

	InitParserCtxt func(ctxt *ParserCtxt) int

	ClearParserCtxt func(ctxt *ParserCtxt)

	FreeParserCtxt func(ctxt *ParserCtxt)

	SetupParserForBuffer func(
		ctxt *ParserCtxt, buffer, filename string)

	CreateDocParserCtxt func(cur string) *ParserCtxt

	GetFeaturesList func(leng *int, result *string) int

	GetFeature func(
		ctxt *ParserCtxt, name string, result *Void) int

	SetFeature func(
		ctxt *ParserCtxt, name string, value *Void) int

	CreatePushParserCtxt func(sax *SAXHandler, userData *Void,
		chunk string, size int, filename string) *ParserCtxt

	ParseChunk func(ctxt *ParserCtxt,
		chunk string, size int, terminate int) int

	CreateIOParserCtxt func(sax *SAXHandler, userData *Void,
		ioread InputReadCallback, ioclose InputCloseCallback,
		ioctx *Void, enc CharEncoding) *ParserCtxt

	NewIOInputStream func(ctxt *ParserCtxt,
		input *ParserInputBuffer, enc CharEncoding) *ParserInput

	ParserFindNodeInfo func(ctxt *ParserCtxt,
		node *Node) *ParserNodeInfo

	InitNodeInfoSeq func(seq *ParserNodeInfoSeq)

	ClearNodeInfoSeq func(seq *ParserNodeInfoSeq)

	ParserFindNodeInfoIndex func(
		seq *ParserNodeInfoSeq, node *Node) UnsignedLong

	ParserAddNodeInfo func(
		ctxt *ParserCtxt, info *ParserNodeInfo)

	SetExternalEntityLoader func(f ExternalEntityLoader)

	GetExternalEntityLoader func() ExternalEntityLoader

	LoadExternalEntity func(
		URL, ID string, ctxt *ParserCtxt) *ParserInput

	ByteConsumed func(ctxt *ParserCtxt) Long

	CtxtReset func(ctxt *ParserCtxt)

	CtxtResetPush func(ctxt *ParserCtxt, chunk string,
		size int, filename, encoding string) int

	CtxtUseOptions func(ctxt *ParserCtxt, options int) int

	ReadDoc func(cur, URL, encoding string, options int) *Doc

	ReadFile func(URL, encoding string, options int) *Doc

	ReadMemory func(buffer string,
		size int, URL, encoding string, options int) *Doc

	ReadFd func(fd int, URL, encoding string, options int) *Doc

	ReadIO func(ioread InputReadCallback,
		ioclose InputCloseCallback, ioctx *Void,
		URL, encoding string, options int) *Doc

	CtxtReadDoc func(ctxt *ParserCtxt,
		cur, URL, encoding string, options int) *Doc

	CtxtReadFile func(ctxt *ParserCtxt,
		filename, encoding string, options ParserOption) *Doc

	CtxtReadMemory func(ctxt *ParserCtxt, buffer string,
		size int, URL, encoding string, options int) *Doc

	CtxtReadFd func(ctxt *ParserCtxt, fd int,
		URL, encoding string, options int) *Doc

	CtxtReadIO func(ctxt *ParserCtxt, ioread InputReadCallback,
		ioclose InputCloseCallback, ioctx *Void,
		URL, encoding string, options int) *Doc

	HasFeature func(feature Feature) int

	XlinkGetDefaultDetect func() XlinkNodeDetectFunc

	XlinkSetDefaultDetect func(f XlinkNodeDetectFunc)

	XlinkGetDefaultHandler func() *XlinkHandler

	XlinkSetDefaultHandler func(handler *XlinkHandler)

	XlinkIsLink func(doc, node *Node) XlinkType

	GetPublicId func(ctx *Void) string

	GetSystemId func(ctx *Void) string

	SetDocumentLocator func(ctx *Void, loc *SAXLocator)

	GetLineNumber func(ctx *Void) int

	GetColumnNumber func(ctx *Void) int

	IsStandalone func(ctx *Void) int

	HasInternalSubset func(ctx *Void) int

	HasExternalSubset func(ctx *Void) int

	InternalSubset func(
		ctx *Void, name, externalID, systemID string)

	ExternalSubset func(
		ctx *Void, name, externalID, systemID string)

	GetEntity func(ctx *Void, name string) *Entity

	GetParameterEntity func(
		ctx *Void, name string) *Entity

	ResolveEntity func(ctx *Void,
		publicId, systemId string) *ParserInput

	EntityDecl func(ctx *Void, name string, typ int,
		publicId, systemId, content string)

	AttributeDecl func(ctx *Void, elem, fullname string,
		typ, def int, defaultValue string,
		tree *Enumeration)

	ElementDecl func(ctx *Void, name string,
		typ int, content *ElementContent)

	NotationDecl func(
		ctx *Void, name, publicId, systemId string)

	UnparsedEntityDecl func(ctx *Void,
		name, publicId, systemId, notationName string)

	StartDocument func(ctx *Void)

	EndDocument func(ctx *Void)

	Attribute func(ctx *Void, fullname, value string)

	StartElement func(ctx *Void, fullname, atts *string)

	EndElement func(ctx *Void, name string)

	Reference func(ctx *Void, name string)

	Characters func(ctx *Void, ch string, leng int)

	IgnorableWhitespace func(ctx *Void, ch string, leng int)

	ProcessingInstruction func(ctx *Void, target, data string)

	GlobalNamespace func(ctx *Void, href, prefix string)

	SetNamespace func(ctx *Void, name string)

	GetNamespace func(ctx *Void) *Ns

	CheckNamespace func(ctx *Void, nameSpace string) int

	NamespaceDecl func(ctx *Void, href, prefix string)

	Comment func(ctx *Void, value string)

	CdataBlock func(ctx *Void, value string, leng int)

	InitXmlDefaultSAXHandler func(
		hdlr *SAXHandlerV1, warning int)

	InithtmlDefaultSAXHandler func(hdlr *SAXHandlerV1)

	InitdocbDefaultSAXHandler func(hdlr *SAXHandlerV1)

	SAX2GetPublicId func(ctx *Void) string

	SAX2GetSystemId func(ctx *Void) string

	SAX2SetDocumentLocator func(ctx *Void, loc *SAXLocator)

	SAX2GetLineNumber func(ctx *Void) int

	SAX2GetColumnNumber func(ctx *Void) int

	SAX2IsStandalone func(ctx *Void) int

	SAX2HasInternalSubset func(ctx *Void) int

	SAX2HasExternalSubset func(ctx *Void) int

	SAX2InternalSubset func(
		ctx *Void, name, externalID, systemID string)

	SAX2ExternalSubset func(
		ctx *Void, name, externalID, systemID string)

	SAX2GetEntity func(ctx *Void, name string) *Entity

	SAX2GetParameterEntity func(ctx *Void, name string) *Entity

	SAX2ResolveEntity func(
		ctx *Void, publicId, systemId string) *ParserInput

	SAX2EntityDecl func(ctx *Void, name string,
		typ int, publicId, systemId, content string)

	SAX2AttributeDecl func(ctx *Void,
		elem, fullname string, typ, def int,
		defaultValue string, tree *Enumeration)

	SAX2ElementDecl func(
		ctx *Void, name string, typ int, content *ElementContent)

	SAX2NotationDecl func(
		ctx *Void, name, publicId, systemId string)

	SAX2UnparsedEntityDecl func(ctx *Void,
		name, publicId, systemId, notationName string)

	SAX2StartDocument func(ctx *Void)

	SAX2EndDocument func(ctx *Void)

	SAX2StartElement func(
		ctx *Void, fullname string, atts *string)

	SAX2EndElement func(ctx *Void, name string)

	SAX2StartElementNs func(ctx *Void,
		localname, prefix, URI string,
		nbNamespaces int, namespaces *string,
		nbAttributes, nbDefaulted int, attributes *string)

	SAX2EndElementNs func(
		ctx *Void, localname, prefix, URI string)

	SAX2Reference func(ctx *Void, name string)

	SAX2Characters func(ctx *Void, ch string, leng int)

	SAX2IgnorableWhitespace func(ctx *Void, ch string, leng int)

	SAX2ProcessingInstruction func(
		ctx *Void, target, data string)

	SAX2Comment func(ctx *Void, value string)

	SAX2CDataBlock func(ctx *Void, value string, leng int)

	SAXDefaultVersion func(version int) int

	SAXVersion func(hdlr *SAXHandler, version int) int

	SAX2InitDefaultSAXHandler func(hdlr *SAXHandler, warning int)

	SAX2InitHtmlDefaultSAXHandler func(hdlr *SAXHandler)

	HtmlDefaultSAXHandlerInit func()

	SAX2InitDocbDefaultSAXHandler func(hdlr *SAXHandler)

	DocbDefaultSAXHandlerInit func()

	DefaultSAXHandlerInit func()

	InitGlobals func()

	CleanupGlobals func()

	InitializeGlobalState func(gs *GlobalState)

	ThrDefSetGenericErrorFunc func(
		ctx *Void, handler GenericErrorFunc)

	ThrDefSetStructuredErrorFunc func(
		ctx *Void, handler StructuredErrorFunc)

	RegisterNodeDefault func(
		f RegisterNodeFunc) RegisterNodeFunc

	ThrDefRegisterNodeDefault func(
		f RegisterNodeFunc) RegisterNodeFunc

	DeregisterNodeDefault func(
		f DeregisterNodeFunc) DeregisterNodeFunc

	ThrDefDeregisterNodeDefault func(
		f DeregisterNodeFunc) DeregisterNodeFunc

	ThrDefOutputBufferCreateFilenameDefault func(
		f OutputBufferCreateFilenameFunc) OutputBufferCreateFilenameFunc

	ThrDefParserInputBufferCreateFilenameDefault func(
		f ParserInputBufferCreateFilenameFunc) ParserInputBufferCreateFilenameFunc

	DocbDefaultSAXHandler func() *SAXHandlerV1

	HtmlDefaultSAXHandler func() *SAXHandlerV1

	LastError func() *Error

	OldXMLWDcompatibility func() *int

	BufferAllocScheme func() *BufferAllocationScheme

	ThrDefBufferAllocScheme func(
		v BufferAllocationScheme) BufferAllocationScheme

	DefaultBufferSize func() *int

	ThrDefDefaultBufferSize func(v int) int

	DefaultSAXHandler func() *SAXHandlerV1

	DefaultSAXLocator func() *SAXLocator

	DoValidityCheckingDefaultValue func() *int

	ThrDefDoValidityCheckingDefaultValue func(v int) int

	GenericError func() *GenericErrorFunc

	StructuredError func() *StructuredErrorFunc

	GenericErrorContext func() **Void

	StructuredErrorContext func() **Void

	GetWarningsDefaultValue func() *int

	ThrDefGetWarningsDefaultValue func(v int) int

	IndentTreeOutput func() *int

	ThrDefIndentTreeOutput func(v int) int

	TreeIndentString func() *string

	ThrDefTreeIndentString func(v string) string

	KeepBlanksDefaultValue func() *int

	ThrDefKeepBlanksDefaultValue func(v int) int

	LineNumbersDefaultValue func() *int

	ThrDefLineNumbersDefaultValue func(v int) int

	LoadExtDtdDefaultValue func() *int

	ThrDefLoadExtDtdDefaultValue func(v int) int

	ParserDebugEntities func() *int

	ThrDefParserDebugEntities func(v int) int

	ParserVersion func() *string

	PedanticParserDefaultValue func() *int

	ThrDefPedanticParserDefaultValue func(v int) int

	SaveNoEmptyTags func() *int

	ThrDefSaveNoEmptyTags func(v int) int

	SubstituteEntitiesDefaultValue func() *int

	ThrDefSubstituteEntitiesDefaultValue func(v int) int

	RegisterNodeDefaultValue func() *RegisterNodeFunc

	DeregisterNodeDefaultValue func() *DeregisterNodeFunc

	ParserInputBufferCreateFilenameValue func() *ParserInputBufferCreateFilenameFunc

	OutputBufferCreateFilenameValue func() *OutputBufferCreateFilenameFunc

	NewMutex func() *Mutex

	MutexLock func(tok *Mutex)

	MutexUnlock func(tok *Mutex)

	FreeMutex func(tok *Mutex)

	NewRMutex func() *RMutex

	RMutexLock func(tok *RMutex)

	RMutexUnlock func(tok *RMutex)

	FreeRMutex func(tok *RMutex)

	InitThreads func()

	LockLibrary func()

	UnlockLibrary func()

	GetThreadId func() int

	IsMainThread func() int

	CleanupThreads func()

	GetGlobalState func() *GlobalState

	HtmlTagLookup func(tag string) *HtmlElemDesc

	HtmlEntityLookup func(name string) *HtmlEntityDesc

	HtmlEntityValueLookup func(
		value UnsignedInt) *HtmlEntityDesc

	HtmlIsAutoClosed func(doc *HtmlDoc, elem *HtmlNode) int

	HtmlAutoCloseTag func(
		doc *HtmlDoc, name string, elem *HtmlNode) int

	HtmlParseEntityRef func(
		ctxt *HtmlParserCtxt, str *string) *HtmlEntityDesc

	HtmlParseCharRef func(ctxt *HtmlParserCtxt) int

	HtmlParseElement func(ctxt *HtmlParserCtxt)

	HtmlNewParserCtxt func() *HtmlParserCtxt

	HtmlCreateMemoryParserCtxt func(
		buffer string, size int) *HtmlParserCtxt

	HtmlParseDocument func(ctxt *HtmlParserCtxt) int

	HtmlSAXParseDoc func(cur, encoding string,
		sax *HtmlSAXHandler, userData *Void) *HtmlDoc

	HtmlParseDoc func(cur string, encoding string) *HtmlDoc

	HtmlSAXParseFile func(filename, encoding string,
		sax *HtmlSAXHandler, userData *Void) *HtmlDoc

	HtmlParseFile func(filename, encoding string) *HtmlDoc
	UTF8ToHtml    func(out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int) int

	HtmlEncodeEntities func(out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int, quoteChar int) int

	HtmlIsScriptAttribute func(name string) int

	HtmlHandleOmittedElem func(val int) int

	HtmlCreatePushParserCtxt func(sax *HtmlSAXHandler,
		userData *Void, chunk string, size int,
		filename string, enc CharEncoding) *HtmlParserCtxt

	HtmlParseChunk func(ctxt *HtmlParserCtxt,
		chunk string, size, terminate int) int

	HtmlFreeParserCtxt func(ctxt *HtmlParserCtxt)

	HtmlCtxtReset func(ctxt *HtmlParserCtxt)

	HtmlCtxtUseOptions func(
		ctxt *HtmlParserCtxt, options int) int

	HtmlReadDoc func(
		cur, URL, encoding string, options int) *HtmlDoc

	HtmlReadFile func(
		URL, encoding string, options int) *HtmlDoc

	HtmlReadMemory func(buffer string, size int,
		URL, encoding string, options int) *HtmlDoc

	HtmlReadFd func(
		fd int, URL, encoding string, options int) *HtmlDoc

	HtmlReadIO func(ioread InputReadCallback,
		ioclose InputCloseCallback,
		ioctx *Void, URL, encoding string, options int) *HtmlDoc

	HtmlCtxtReadDoc func(ctxt *ParserCtxt,
		cur, URL, encoding string, options int) *HtmlDoc

	HtmlCtxtReadFile func(ctxt *ParserCtxt,
		filename, encoding string, options int) *HtmlDoc

	HtmlCtxtReadMemory func(ctxt *ParserCtxt, buffer string,
		size int, URL, encoding string, options int) *HtmlDoc

	HtmlCtxtReadFd func(ctxt *ParserCtxt, fd int,
		URL, encoding string, options int) *HtmlDoc

	HtmlCtxtReadIO func(ctxt *ParserCtxt,
		ioread InputReadCallback, ioclose InputCloseCallback,
		ioctx *Void, URL, encoding string, options int) *HtmlDoc

	HtmlAttrAllowed func(*HtmlElemDesc, string, int) HtmlStatus

	HtmlElementAllowedHere func(*HtmlElemDesc, string) int

	HtmlElementStatusHere func(
		*HtmlElemDesc, *HtmlElemDesc) HtmlStatus

	HtmlNodeStatus func(*HtmlNode, int) HtmlStatus

	HtmlNewDoc func(URI string, externalID string) *HtmlDoc

	HtmlNewDocNoDtD func(URI string, externalID string) *HtmlDoc

	HtmlGetMetaEncoding func(doc *HtmlDoc) string

	HtmlSetMetaEncoding func(doc *HtmlDoc, encoding string) int

	HtmlDocDumpMemory func(cur *Doc, mem *string, size *int)

	HtmlDocDumpMemoryFormat func(cur *Doc, mem *string, size *int, format int)

	HtmlDocDump func(f *FILE, cur *Doc) int

	HtmlSaveFile func(filename string, cur *Doc) int

	HtmlNodeDump func(buf *Buffer, doc *Doc, cur *Node) int

	HtmlNodeDumpFile func(out *FILE, doc *Doc, cur *Node)

	HtmlNodeDumpFileFormat func(out *FILE, doc *Doc,
		cur *Node, encoding string, format int) int

	HtmlSaveFileEnc func(
		filename string, cur *Doc, encoding string) int

	HtmlSaveFileFormat func(filename string,
		cur *Doc, encoding string, format int) int

	HtmlNodeDumpFormatOutput func(buf *OutputBuffer,
		doc *Doc, cur *Node, encoding string, format int)

	HtmlDocContentDumpOutput func(buf *OutputBuffer,
		cur *Doc, encoding string)

	HtmlDocContentDumpFormatOutput func(buf *OutputBuffer,
		cur *Doc, encoding string, format int)

	HtmlNodeDumpOutput func(buf *OutputBuffer,
		doc *Doc, cur *Node, encoding string)

	HtmlIsBooleanAttr func(name string) int

	XPathFreeObject func(obj *XPathObject)

	XPathNodeSetCreate func(val *Node) *NodeSet

	XPathFreeNodeSetList func(obj *XPathObject)

	XPathFreeNodeSet func(obj *NodeSet)

	XPathObjectCopy func(val *XPathObject) *XPathObject

	XPathCmpNodes func(node1, node2 *Node) int

	XPathCastNumberToBoolean func(val Double) int

	XPathCastStringToBoolean func(val string) int

	XPathCastNodeSetToBoolean func(ns *NodeSet) int

	XPathCastToBoolean func(val *XPathObject) int

	XPathCastBooleanToNumber func(val int) Double

	XPathCastStringToNumber func(val string) Double

	XPathCastNodeToNumber func(node *Node) Double

	XPathCastNodeSetToNumber func(ns *NodeSet) Double

	XPathCastToNumber func(val *XPathObject) Double

	XPathCastBooleanToString func(val int) string

	XPathCastNumberToString func(val Double) string

	XPathCastNodeToString func(node *Node) string

	XPathCastNodeSetToString func(ns *NodeSet) string

	XPathCastToString func(val *XPathObject) string

	XPathConvertBoolean func(val *XPathObject) *XPathObject

	XPathConvertNumber func(val *XPathObject) *XPathObject

	XPathConvertString func(val *XPathObject) *XPathObject

	XPathNewContext func(doc *Doc) *XPathContext

	XPathFreeContext func(ctxt *XPathContext)

	XPathContextSetCache func(
		ctxt *XPathContext, active, value, options int) int

	XPathOrderDocElems func(doc *Doc) Long

	XPathEval func(str string, ctx *XPathContext) *XPathObject

	XPathEvalExpression func(
		str string, ctxt *XPathContext) *XPathObject

	XPathEvalPredicate func(
		ctxt *XPathContext, res *XPathObject) int

	XPathCompile func(str string) *XPathCompExpr

	XPathCtxtCompile func(
		ctxt *XPathContext, str string) *XPathCompExpr

	XPathCompiledEval func(
		comp *XPathCompExpr, ctx *XPathContext) *XPathObject

	XPathCompiledEvalToBoolean func(
		comp *XPathCompExpr, ctxt *XPathContext) int

	XPathFreeCompExpr func(comp *XPathCompExpr)

	XPathInit func()

	XPathIsNaN func(val Double) int

	XPathIsInf func(val Double) int

	XPathPopBoolean func(ctxt *XPathParserContext) int

	XPathPopNumber func(ctxt *XPathParserContext) Double

	XPathPopString func(ctxt *XPathParserContext) string

	XPathPopNodeSet func(ctxt *XPathParserContext) *NodeSet

	XPathPopExternal func(ctxt *XPathParserContext) *Void

	XPathRegisterVariableLookup func(ctxt *XPathContext,
		f XPathVariableLookupFunc, data *Void)

	XPathRegisterFuncLookup func(ctxt *XPathContext,
		f XPathFuncLookupFunc, funcCtxt *Void)

	XPatherror func(
		ctxt *XPathParserContext, file string, line, no int)

	XPathErr func(ctxt *XPathParserContext, error int)

	XPathDebugDumpObject func(
		output *FILE, cur *XPathObject, depth int)

	XPathDebugDumpCompExpr func(
		output *FILE, comp *XPathCompExpr, depth int)

	XPathNodeSetContains func(cur *NodeSet, val *Node) int

	XPathDifference func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathIntersection func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathDistinctSorted func(nodes *NodeSet) *NodeSet

	XPathDistinct func(nodes *NodeSet) *NodeSet

	XPathHasSameNodes func(nodes1, nodes2 *NodeSet) int

	XPathNodeLeadingSorted func(
		nodes *NodeSet, node *Node) *NodeSet

	XPathLeadingSorted func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathNodeLeading func(nodes *NodeSet, node *Node) *NodeSet

	XPathLeading func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathNodeTrailingSorted func(
		nodes *NodeSet, node *Node) *NodeSet

	XPathTrailingSorted func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathNodeTrailing func(nodes *NodeSet, node *Node) *NodeSet

	XPathTrailing func(nodes1, nodes2 *NodeSet) *NodeSet

	XPathRegisterNs func(
		ctxt *XPathContext, prefix, nsUri string) int

	XPathNsLookup func(
		ctxt *XPathContext, prefix string) string

	XPathRegisteredNsCleanup func(ctxt *XPathContext)

	XPathRegisterFunc func(
		ctxt *XPathContext, name string, f XPathFunction) int

	XPathRegisterFuncNS func(ctxt *XPathContext,
		name, nsUri string, f XPathFunction) int

	XPathRegisterVariable func(ctxt *XPathContext,
		name string, value *XPathObject) int

	XPathRegisterVariableNS func(ctxt *XPathContext,
		name, nsUri string, value *XPathObject) int

	XPathFunctionLookup func(
		ctxt *XPathContext, name string) XPathFunction

	XPathFunctionLookupNS func(
		ctxt *XPathContext, name, nsUri string) XPathFunction

	XPathRegisteredFuncsCleanup func(ctxt *XPathContext)

	XPathVariableLookup func(
		ctxt *XPathContext, name string) *XPathObject

	XPathVariableLookupNS func(
		ctxt *XPathContext, name, nsUri string) *XPathObject

	XPathRegisteredVariablesCleanup func(ctxt *XPathContext)

	XPathNewParserContext func(
		str string, ctxt *XPathContext) *XPathParserContext

	XPathFreeParserContext func(ctxt *XPathParserContext)

	ValuePop func(ctxt *XPathParserContext) *XPathObject

	ValuePush func(
		ctxt *XPathParserContext, value *XPathObject) int

	XPathNewString func(val string) *XPathObject

	XPathNewCString func(val string) *XPathObject

	XPathWrapString func(val string) *XPathObject

	XPathWrapCString func(val string) *XPathObject

	XPathNewFloat func(val Double) *XPathObject

	XPathNewBoolean func(val int) *XPathObject

	XPathNewNodeSet func(val *Node) *XPathObject

	XPathNewValueTree func(val *Node) *XPathObject

	XPathNodeSetAdd func(cur *NodeSet, val *Node) int

	XPathNodeSetAddUnique func(cur *NodeSet, val *Node) int

	XPathNodeSetAddNs func(cur *NodeSet, node *Node, ns *Ns) int

	XPathNodeSetSort func(set *NodeSet)

	XPathRoot func(ctxt *XPathParserContext)

	XPathEvalExpr func(ctxt *XPathParserContext)

	XPathParseName func(ctxt *XPathParserContext) string

	XPathParseNCName func(ctxt *XPathParserContext) string

	XPathStringEvalNumber func(str string) Double

	XPathEvaluatePredicateResult func(
		ctxt *XPathParserContext, res *XPathObject) int

	XPathRegisterAllFunctions func(ctxt *XPathContext)

	XPathNodeSetMerge func(val1, val2 *NodeSet) *NodeSet

	XPathNodeSetDel func(cur *NodeSet, val *Node)

	XPathNodeSetRemove func(cur *NodeSet, val int)

	XPathNewNodeSetList func(val *NodeSet) *XPathObject

	XPathWrapNodeSet func(val *NodeSet) *XPathObject

	XPathWrapExternal func(val *Void) *XPathObject

	XPathEqualValues func(ctxt *XPathParserContext) int

	XPathNotEqualValues func(ctxt *XPathParserContext) int

	XPathCompareValues func(
		ctxt *XPathParserContext, inf, strict int) int

	XPathValueFlipSign func(ctxt *XPathParserContext)

	XPathAddValues func(ctxt *XPathParserContext)

	XPathSubValues func(ctxt *XPathParserContext)

	XPathMultValues func(ctxt *XPathParserContext)

	XPathDivValues func(ctxt *XPathParserContext)

	XPathModValues func(ctxt *XPathParserContext)

	XPathIsNodeType func(name string) int

	XPathNextSelf func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextChild func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextDescendant func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextDescendantOrSelf func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextParent func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextAncestorOrSelf func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextFollowingSibling func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextFollowing func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextNamespace func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextAttribute func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextPreceding func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextAncestor func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathNextPrecedingSibling func(
		ctxt *XPathParserContext, cur *Node) *Node

	XPathLastFunction func(ctxt *XPathParserContext, nargs int)

	XPathPositionFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathCountFunction func(ctxt *XPathParserContext, nargs int)

	XPathIdFunction func(ctxt *XPathParserContext, nargs int)

	XPathLocalNameFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathNamespaceURIFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathStringFunction func(ctxt *XPathParserContext, nargs int)

	XPathStringLengthFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathConcatFunction func(ctxt *XPathParserContext, nargs int)

	XPathContainsFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathStartsWithFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathSubstringFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathSubstringBeforeFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathSubstringAfterFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathNormalizeFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathTranslateFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathNotFunction func(ctxt *XPathParserContext, nargs int)

	XPathTrueFunction func(ctxt *XPathParserContext, nargs int)

	XPathFalseFunction func(ctxt *XPathParserContext, nargs int)

	XPathLangFunction func(ctxt *XPathParserContext, nargs int)

	XPathNumberFunction func(ctxt *XPathParserContext, nargs int)

	XPathSumFunction func(ctxt *XPathParserContext, nargs int)

	XPathFloorFunction func(ctxt *XPathParserContext, nargs int)

	XPathCeilingFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathRoundFunction func(ctxt *XPathParserContext, nargs int)

	XPathBooleanFunction func(
		ctxt *XPathParserContext, nargs int)

	XPathNodeSetFreeNs func(ns *Ns)

	XIncludeProcess func(doc *Doc) int

	XIncludeProcessFlags func(doc *Doc, flags int) int

	XIncludeProcessFlagsData func(
		doc *Doc, flags int, data *Void) int

	XIncludeProcessTreeFlagsData func(
		tree *Node, flags int, data *Void) int

	XIncludeProcessTree func(tree *Node) int

	XIncludeProcessTreeFlags func(tree *Node, flags int) int

	XIncludeNewContext func(doc *Doc) *XIncludeCtxt

	XIncludeSetFlags func(ctxt *XIncludeCtxt, flags int) int

	XIncludeFreeContext func(ctxt *XIncludeCtxt)

	XIncludeProcessNode func(ctxt *XIncludeCtxt, tree *Node) int

	XPtrLocationSetCreate func(val *XPathObject) *LocationSet

	XPtrFreeLocationSet func(obj *LocationSet)

	XPtrLocationSetMerge func(
		val1, val2 *LocationSet) *LocationSet

	XPtrNewRange func(start *Node, startindex int,
		end *Node, endindex int) *XPathObject

	XPtrNewRangePoints func(start, end *XPathObject) *XPathObject

	XPtrNewRangeNodePoint func(
		start *Node, end *XPathObject) *XPathObject

	XPtrNewRangePointNode func(
		start *XPathObject, end *Node) *XPathObject

	XPtrNewRangeNodes func(start, end *Node) *XPathObject

	XPtrNewLocationSetNodes func(start, end *Node) *XPathObject

	XPtrNewLocationSetNodeSet func(set *NodeSet) *XPathObject

	XPtrNewRangeNodeObject func(
		start *Node, end *XPathObject) *XPathObject

	XPtrNewCollapsedRange func(start *Node) *XPathObject

	XPtrLocationSetAdd func(cur *LocationSet, val *XPathObject)

	XPtrWrapLocationSet func(val *LocationSet) *XPathObject

	XPtrLocationSetDel func(cur *LocationSet, val *XPathObject)

	XPtrLocationSetRemove func(cur *LocationSet, val int)

	XPtrNewContext func(
		doc *Doc, here, origin *Node) *XPathContext

	XPtrEval func(str string, ctx *XPathContext) *XPathObject

	XPtrRangeToFunction func(
		ctxt *XPathParserContext, nargs int)

	XPtrBuildNodeList func(obj *XPathObject) *Node

	XPtrEvalRangePredicate func(ctxt *XPathParserContext)

	UCSIsAegeanNumbers func(code int) int

	UCSIsAlphabeticPresentationForms func(code int) int

	UCSIsArabic func(code int) int

	UCSIsArabicPresentationFormsA func(code int) int

	UCSIsArabicPresentationFormsB func(code int) int

	UCSIsArmenian func(code int) int

	UCSIsArrows func(code int) int

	UCSIsBasicLatin func(code int) int

	UCSIsBengali func(code int) int

	UCSIsBlockElements func(code int) int

	UCSIsBopomofo func(code int) int

	UCSIsBopomofoExtended func(code int) int

	UCSIsBoxDrawing func(code int) int

	UCSIsBraillePatterns func(code int) int

	UCSIsBuhid func(code int) int

	UCSIsByzantineMusicalSymbols func(code int) int

	UCSIsCJKCompatibility func(code int) int

	UCSIsCJKCompatibilityForms func(code int) int

	UCSIsCJKCompatibilityIdeographs func(code int) int

	UCSIsCJKCompatibilityIdeographsSupplement func(code int) int

	UCSIsCJKRadicalsSupplement func(code int) int

	UCSIsCJKSymbolsandPunctuation func(code int) int

	UCSIsCJKUnifiedIdeographs func(code int) int

	UCSIsCJKUnifiedIdeographsExtensionA func(code int) int

	UCSIsCJKUnifiedIdeographsExtensionB func(code int) int

	UCSIsCherokee func(code int) int

	UCSIsCombiningDiacriticalMarks func(code int) int

	UCSIsCombiningDiacriticalMarksforSymbols func(code int) int

	UCSIsCombiningHalfMarks func(code int) int

	UCSIsCombiningMarksforSymbols func(code int) int

	UCSIsControlPictures func(code int) int

	UCSIsCurrencySymbols func(code int) int

	UCSIsCypriotSyllabary func(code int) int

	UCSIsCyrillic func(code int) int

	UCSIsCyrillicSupplement func(code int) int

	UCSIsDeseret func(code int) int

	UCSIsDevanagari func(code int) int

	UCSIsDingbats func(code int) int

	UCSIsEnclosedAlphanumerics func(code int) int

	UCSIsEnclosedCJKLettersandMonths func(code int) int

	UCSIsEthiopic func(code int) int

	UCSIsGeneralPunctuation func(code int) int

	UCSIsGeometricShapes func(code int) int

	UCSIsGeorgian func(code int) int

	UCSIsGothic func(code int) int

	UCSIsGreek func(code int) int

	UCSIsGreekExtended func(code int) int

	UCSIsGreekandCoptic func(code int) int

	UCSIsGujarati func(code int) int

	UCSIsGurmukhi func(code int) int

	UCSIsHalfwidthandFullwidthForms func(code int) int

	UCSIsHangulCompatibilityJamo func(code int) int

	UCSIsHangulJamo func(code int) int

	UCSIsHangulSyllables func(code int) int

	UCSIsHanunoo func(code int) int

	UCSIsHebrew func(code int) int

	UCSIsHighPrivateUseSurrogates func(code int) int

	UCSIsHighSurrogates func(code int) int

	UCSIsHiragana func(code int) int

	UCSIsIPAExtensions func(code int) int

	UCSIsIdeographicDescriptionCharacters func(code int) int

	UCSIsKanbun func(code int) int

	UCSIsKangxiRadicals func(code int) int

	UCSIsKannada func(code int) int

	UCSIsKatakana func(code int) int

	UCSIsKatakanaPhoneticExtensions func(code int) int

	UCSIsKhmer func(code int) int

	UCSIsKhmerSymbols func(code int) int

	UCSIsLao func(code int) int

	UCSIsLatin1Supplement func(code int) int

	UCSIsLatinExtendedA func(code int) int

	UCSIsLatinExtendedB func(code int) int

	UCSIsLatinExtendedAdditional func(code int) int

	UCSIsLetterlikeSymbols func(code int) int

	UCSIsLimbu func(code int) int

	UCSIsLinearBIdeograms func(code int) int

	UCSIsLinearBSyllabary func(code int) int

	UCSIsLowSurrogates func(code int) int

	UCSIsMalayalam func(code int) int

	UCSIsMathematicalAlphanumericSymbols func(code int) int

	UCSIsMathematicalOperators func(code int) int

	UCSIsMiscellaneousMathematicalSymbolsA func(code int) int

	UCSIsMiscellaneousMathematicalSymbolsB func(code int) int

	UCSIsMiscellaneousSymbols func(code int) int

	UCSIsMiscellaneousSymbolsandArrows func(code int) int

	UCSIsMiscellaneousTechnical func(code int) int

	UCSIsMongolian func(code int) int

	UCSIsMusicalSymbols func(code int) int

	UCSIsMyanmar func(code int) int

	UCSIsNumberForms func(code int) int

	UCSIsOgham func(code int) int

	UCSIsOldItalic func(code int) int

	UCSIsOpticalCharacterRecognition func(code int) int

	UCSIsOriya func(code int) int

	UCSIsOsmanya func(code int) int

	UCSIsPhoneticExtensions func(code int) int

	UCSIsPrivateUse func(code int) int

	UCSIsPrivateUseArea func(code int) int

	UCSIsRunic func(code int) int

	UCSIsShavian func(code int) int

	UCSIsSinhala func(code int) int

	UCSIsSmallFormVariants func(code int) int

	UCSIsSpacingModifierLetters func(code int) int

	UCSIsSpecials func(code int) int

	UCSIsSuperscriptsandSubscripts func(code int) int

	UCSIsSupplementalArrowsA func(code int) int

	UCSIsSupplementalArrowsB func(code int) int

	UCSIsSupplementalMathematicalOperators func(code int) int

	UCSIsSupplementaryPrivateUseAreaA func(code int) int

	UCSIsSupplementaryPrivateUseAreaB func(code int) int

	UCSIsSyriac func(code int) int

	UCSIsTagalog func(code int) int

	UCSIsTagbanwa func(code int) int

	UCSIsTags func(code int) int

	UCSIsTaiLe func(code int) int

	UCSIsTaiXuanJingSymbols func(code int) int

	UCSIsTamil func(code int) int

	UCSIsTelugu func(code int) int

	UCSIsThaana func(code int) int

	UCSIsThai func(code int) int

	UCSIsTibetan func(code int) int

	UCSIsUgaritic func(code int) int

	UCSIsUnifiedCanadianAboriginalSyllabics func(code int) int

	UCSIsVariationSelectors func(code int) int

	UCSIsVariationSelectorsSupplement func(code int) int

	UCSIsYiRadicals func(code int) int

	UCSIsYiSyllables func(code int) int

	UCSIsYijingHexagramSymbols func(code int) int

	UCSIsBlock func(code int, block string) int

	UCSIsCatC func(code int) int

	UCSIsCatCc func(code int) int

	UCSIsCatCf func(code int) int

	UCSIsCatCo func(code int) int

	UCSIsCatCs func(code int) int

	UCSIsCatL func(code int) int

	UCSIsCatLl func(code int) int

	UCSIsCatLm func(code int) int

	UCSIsCatLo func(code int) int

	UCSIsCatLt func(code int) int

	UCSIsCatLu func(code int) int

	UCSIsCatM func(code int) int

	UCSIsCatMc func(code int) int

	UCSIsCatMe func(code int) int

	UCSIsCatMn func(code int) int

	UCSIsCatN func(code int) int

	UCSIsCatNd func(code int) int

	UCSIsCatNl func(code int) int

	UCSIsCatNo func(code int) int

	UCSIsCatP func(code int) int

	UCSIsCatPc func(code int) int

	UCSIsCatPd func(code int) int

	UCSIsCatPe func(code int) int

	UCSIsCatPf func(code int) int

	UCSIsCatPi func(code int) int

	UCSIsCatPo func(code int) int

	UCSIsCatPs func(code int) int

	UCSIsCatS func(code int) int

	UCSIsCatSc func(code int) int

	UCSIsCatSk func(code int) int

	UCSIsCatSm func(code int) int

	UCSIsCatSo func(code int) int

	UCSIsCatZ func(code int) int

	UCSIsCatZl func(code int) int

	UCSIsCatZp func(code int) int

	UCSIsCatZs func(code int) int

	UCSIsCat func(code int, cat string) int

	SchemaNewParserCtxt func(URL string) *SchemaParserCtxt

	SchemaNewMemParserCtxt func(
		buffer string, size int) *SchemaParserCtxt

	SchemaNewDocParserCtxt func(doc *Doc) *SchemaParserCtxt

	SchemaFreeParserCtxt func(ctxt *SchemaParserCtxt)

	SchemaSetParserErrors func(
		ctxt *SchemaParserCtxt, err SchemaValidityErrorFunc,
		warn SchemaValidityWarningFunc, ctx *Void)

	SchemaSetParserStructuredErrors func(
		ctxt *SchemaParserCtxt, serror StructuredErrorFunc,
		ctx *Void)

	SchemaGetParserErrors func(
		ctxt *SchemaParserCtxt, err *SchemaValidityErrorFunc,
		warn *SchemaValidityWarningFunc, ctx **Void) int

	SchemaIsValid func(ctxt *SchemaValidCtxt) int

	SchemaParse func(ctxt *SchemaParserCtxt) *Schema

	SchemaFree func(schema *Schema)

	SchemaDump func(output *FILE, schema *Schema)

	SchemaSetValidErrors func(
		ctxt *SchemaValidCtxt, err SchemaValidityErrorFunc,
		warn SchemaValidityWarningFunc, ctx *Void)

	SchemaSetValidStructuredErrors func(ctxt *SchemaValidCtxt,
		serror StructuredErrorFunc, ctx *Void)

	SchemaGetValidErrors func(
		ctxt *SchemaValidCtxt, err *SchemaValidityErrorFunc,
		warn *SchemaValidityWarningFunc, ctx **Void) int

	SchemaSetValidOptions func(
		ctxt *SchemaValidCtxt, options int) int

	SchemaValidateSetFilename func(
		vctxt *SchemaValidCtxt, filename string)

	SchemaValidCtxtGetOptions func(ctxt *SchemaValidCtxt) int

	SchemaNewValidCtxt func(schema *Schema) *SchemaValidCtxt

	SchemaFreeValidCtxt func(ctxt *SchemaValidCtxt)

	SchemaValidateDoc func(
		ctxt *SchemaValidCtxt, instance *Doc) int

	SchemaValidateOneElement func(
		ctxt *SchemaValidCtxt, elem *Node) int

	SchemaValidateStream func(
		ctxt *SchemaValidCtxt, input *ParserInputBuffer,
		enc CharEncoding, sax *SAXHandler, userData *Void) int

	SchemaValidateFile func(
		ctxt *SchemaValidCtxt, filename string, options int) int

	SchemaValidCtxtGetParserCtxt func(
		ctxt *SchemaValidCtxt) *ParserCtxt

	SchemaSAXUnplug func(plug *SchemaSAXPlugStruct) int

	SchemaValidateSetLocator func(vctxt *SchemaValidCtxt,
		f SchemaValidityLocatorFunc, ctxt *Void)

	SchemaSAXPlug func(ctxt *SchemaValidCtxt,
		sax **SAXHandler, userData **Void) *SchemaSAXPlugStruct

	SchemaFreeType func(typ *SchemaType)

	SchemaFreeWildcard func(wildcard *SchemaWildcard)

	SchemaInitTypes func()

	SchemaCleanupTypes func()

	SchemaGetPredefinedType func(name, ns string) *SchemaType

	SchemaValidatePredefinedType func(typ *SchemaType,
		value string, val **SchemaVal) int

	SchemaValPredefTypeNode func(typ *SchemaType,
		value string, val **SchemaVal, node *Node) int

	SchemaValidateFacet func(base *SchemaType,
		facet *SchemaFacet, value string, val *SchemaVal) int

	SchemaValidateFacetWhtsp func(facet *SchemaFacet,
		fws SchemaWhitespaceValueType, valType SchemaValType,
		value string, val *SchemaVal,
		ws SchemaWhitespaceValueType) int

	SchemaFreeValue func(val *SchemaVal)

	SchemaNewFacet func() *SchemaFacet

	SchemaCheckFacet func(facet *SchemaFacet,
		typeDecl *SchemaType, ctxt *SchemaParserCtxt,
		name string) int

	SchemaFreeFacet func(facet *SchemaFacet)

	SchemaCompareValues func(x, y *SchemaVal) int

	SchemaGetBuiltInListSimpleTypeItemType func(
		typ *SchemaType) *SchemaType

	SchemaValidateListSimpleTypeFacet func(
		facet *SchemaFacet, value string,
		actualLen UnsignedLong, expectedLen *UnsignedLong) int

	SchemaGetBuiltInType func(typ SchemaValType) *SchemaType

	SchemaIsBuiltInTypeFacet func(
		typ *SchemaType, facetType int) int

	SchemaCollapseString func(value string) string

	SchemaWhiteSpaceReplace func(value string) string

	SchemaGetFacetValueAsULong func(
		facet *SchemaFacet) UnsignedLong

	SchemaValidateLengthFacet func(typ *SchemaType,
		facet *SchemaFacet, value string,
		val *SchemaVal, length *UnsignedLong) int

	SchemaValidateLengthFacetWhtsp func(
		facet *SchemaFacet, valType SchemaValType,
		value string, val *SchemaVal, length *UnsignedLong,
		ws SchemaWhitespaceValueType) int

	SchemaValPredefTypeNodeNoNorm func(
		typ *SchemaType, value string,
		val **SchemaVal, node *Node) int

	SchemaGetCanonValue func(
		val *SchemaVal, retValue *string) int

	SchemaGetCanonValueWhtsp func(val *SchemaVal,
		retValue *string, ws SchemaWhitespaceValueType) int

	SchemaValueAppend func(prev, cur *SchemaVal) int

	SchemaValueGetNext func(cur *SchemaVal) *SchemaVal

	SchemaValueGetAsString func(val *SchemaVal) string

	SchemaValueGetAsBoolean func(val *SchemaVal) int

	SchemaNewStringValue func(
		typ SchemaValType, value string) *SchemaVal

	SchemaNewNOTATIONValue func(name, ns string) *SchemaVal

	SchemaNewQNameValue func(
		namespaceName, localName string) *SchemaVal

	SchemaCompareValuesWhtsp func(
		x *SchemaVal, xws SchemaWhitespaceValueType,
		y *SchemaVal, yws SchemaWhitespaceValueType) int

	SchemaCopyValue func(val *SchemaVal) *SchemaVal

	SchemaGetValType func(val *SchemaVal) SchemaValType

	SchematronNewParserCtxt func(
		URL string) *SchematronParserCtxt

	SchematronNewMemParserCtxt func(
		buffer string, size int) *SchematronParserCtxt

	SchematronNewDocParserCtxt func(
		doc *Doc) *SchematronParserCtxt

	SchematronFreeParserCtxt func(ctxt *SchematronParserCtxt)

	SchematronParse func(ctxt *SchematronParserCtxt) *Schematron

	SchematronFree func(schema *Schematron)

	SchematronSetValidStructuredErrors func(
		ctxt *SchematronValidCtxt,
		serror StructuredErrorFunc, ctx *Void)

	SchematronNewValidCtxt func(
		schema *Schematron, options int) *SchematronValidCtxt

	SchematronFreeValidCtxt func(ctxt *SchematronValidCtxt)

	SchematronValidateDoc func(
		ctxt *SchematronValidCtxt, instance *Doc) int

	CharInRange func(
		val UnsignedInt, group *ChRangeGroup) int

	IsLetter func(c int) int

	CreateFileParserCtxt func(filename string) *ParserCtxt

	CreateURLParserCtxt func(
		filename string, options int) *ParserCtxt

	CreateMemoryParserCtxt func(
		buffer string, size int) *ParserCtxt

	CreateEntityParserCtxt func(URL, ID, base string) *ParserCtxt

	SwitchEncoding func(ctxt *ParserCtxt, enc CharEncoding) int

	SwitchToEncoding func(
		ctxt *ParserCtxt, handler *CharEncodingHandler) int

	SwitchInputEncoding func(ctxt *ParserCtxt,
		input *ParserInput, handler *CharEncodingHandler) int

	NewStringInputStream func(
		ctxt *ParserCtxt, buffer string) *ParserInput

	NewEntityInputStream func(
		ctxt *ParserCtxt, entity *Entity) *ParserInput

	PushInput func(ctxt *ParserCtxt, input *ParserInput) int

	PopInput func(ctxt *ParserCtxt) Char

	FreeInputStream func(input *ParserInput)

	NewInputFromFile func(
		ctxt *ParserCtxt, filename string) *ParserInput

	NewInputStream func(ctxt *ParserCtxt) *ParserInput

	SplitQName func(
		ctxt *ParserCtxt, name, prefix *string) string

	ParseName func(ctxt *ParserCtxt) string

	ParseNmtoken func(ctxt *ParserCtxt) string

	ParseEntityValue func(ctxt *ParserCtxt, orig *string) string

	ParseAttValue func(ctxt *ParserCtxt) string

	ParseSystemLiteral func(ctxt *ParserCtxt) string

	ParsePubidLiteral func(ctxt *ParserCtxt) string

	ParseCharData func(ctxt *ParserCtxt, cdata int)

	ParseExternalID func(
		ctxt *ParserCtxt, publicID *string, strict int) string

	ParseComment func(ctxt *ParserCtxt)

	ParsePITarget func(ctxt *ParserCtxt) string

	ParsePI func(ctxt *ParserCtxt)

	ParseNotationDecl func(ctxt *ParserCtxt)

	ParseEntityDecl func(ctxt *ParserCtxt)

	ParseDefaultDecl func(ctxt *ParserCtxt, value *string) int

	ParseNotationType func(ctxt *ParserCtxt) *Enumeration

	ParseEnumerationType func(ctxt *ParserCtxt) *Enumeration

	ParseEnumeratedType func(
		ctxt *ParserCtxt, tree **Enumeration) int

	ParseAttributeType func(
		ctxt *ParserCtxt, tree **Enumeration) int

	ParseAttributeListDecl func(ctxt *ParserCtxt)

	ParseElementMixedContentDecl func(
		ctxt *ParserCtxt, inputchk int) *ElementContent

	ParseElementChildrenContentDecl func(
		ctxt *ParserCtxt, inputchk int) *ElementContent

	ParseElementContentDecl func(ctxt *ParserCtxt,
		name string, result **ElementContent) int

	ParseElementDecl func(ctxt *ParserCtxt) int

	ParseMarkupDecl func(ctxt *ParserCtxt)

	ParseCharRef func(ctxt *ParserCtxt) int

	ParseEntityRef func(ctxt *ParserCtxt) *Entity

	ParseReference func(ctxt *ParserCtxt)

	ParsePEReference func(ctxt *ParserCtxt)

	ParseDocTypeDecl func(ctxt *ParserCtxt)

	ParseAttribute func(ctxt *ParserCtxt, value *string) string

	ParseStartTag func(ctxt *ParserCtxt) string

	ParseEndTag func(ctxt *ParserCtxt)

	ParseCDSect func(ctxt *ParserCtxt)

	ParseContent func(ctxt *ParserCtxt)

	ParseElement func(ctxt *ParserCtxt)

	ParseVersionNum func(ctxt *ParserCtxt) string

	ParseVersionInfo func(ctxt *ParserCtxt) string

	ParseEncName func(ctxt *ParserCtxt) string

	ParseEncodingDecl func(ctxt *ParserCtxt) string

	ParseSDDecl func(ctxt *ParserCtxt) int

	ParseXMLDecl func(ctxt *ParserCtxt)

	ParseTextDecl func(ctxt *ParserCtxt)

	ParseMisc func(ctxt *ParserCtxt)

	ParseExternalSubset func(
		ctxt *ParserCtxt, externalID string, systemID string)

	StringDecodeEntities func(ctxt *ParserCtxt, str string,
		what int, end, end2, end3 Char) string

	StringLenDecodeEntities func(ctxt *ParserCtxt, str string,
		leng int, what int, end, end2, end3 Char) string

	NodePush func(ctxt *ParserCtxt, value *Node) int

	NodePop func(ctxt *ParserCtxt) *Node

	InputPush func(ctxt *ParserCtxt, value *ParserInput) int

	InputPop func(ctxt *ParserCtxt) *ParserInput

	NamePop func(ctxt *ParserCtxt) string

	NamePush func(ctxt *ParserCtxt, value string) int

	SkipBlankChars func(ctxt *ParserCtxt) int

	StringCurrentChar func(
		ctxt *ParserCtxt, cur string, leng *int) int

	ParserHandlePEReference func(ctxt *ParserCtxt)

	CheckLanguageID func(lang string) int

	CurrentChar func(ctxt *ParserCtxt, len *int) int

	CopyCharMultiByte func(out string, val int) int

	CopyChar func(len int, out string, val int) int

	NextChar func(ctxt *ParserCtxt)

	ParserInputShrink func(in *ParserInput)

	HtmlInitAutoClose func()

	HtmlCreateFileParserCtxt func(
		filename, encoding string) *HtmlParserCtxt

	SetEntityReferenceFunc func(f EntityReferenceFunc)

	ParseQuotedString func(ctxt *ParserCtxt) string

	ParseNamespace func(ctxt *ParserCtxt)

	NamespaceParseNSDef func(ctxt *ParserCtxt) string

	ScanName func(ctxt *ParserCtxt) string

	NamespaceParseNCName func(ctxt *ParserCtxt) string

	ParserHandleReference func(ctxt *ParserCtxt)

	NamespaceParseQName func(
		ctxt *ParserCtxt, prefix *string) string

	DecodeEntities func(ctxt *ParserCtxt, leng, what int,
		end, end2, end3 Char) string

	HandleEntity func(ctxt *ParserCtxt, entity *Entity)

	RelaxNGInitTypes func() int

	RelaxNGCleanupTypes func()

	RelaxNGNewParserCtxt func(URL string) *RelaxNGParserCtxt

	RelaxNGNewMemParserCtxt func(
		buffer string, size int) *RelaxNGParserCtxt

	RelaxNGNewDocParserCtxt func(doc *Doc) *RelaxNGParserCtxt

	RelaxParserSetFlag func(
		ctxt *RelaxNGParserCtxt, flag int) int

	RelaxNGFreeParserCtxt func(ctxt *RelaxNGParserCtxt)

	RelaxNGSetParserErrors func(ctxt *RelaxNGParserCtxt,
		err RelaxNGValidityErrorFunc,
		warn RelaxNGValidityWarningFunc, ctx *Void)

	RelaxNGGetParserErrors func(ctxt *RelaxNGParserCtxt,
		err *RelaxNGValidityErrorFunc,
		warn *RelaxNGValidityWarningFunc, ctx **Void) int

	RelaxNGSetParserStructuredErrors func(
		ctxt *RelaxNGParserCtxt,
		serror StructuredErrorFunc, ctx *Void)

	RelaxNGParse func(ctxt *RelaxNGParserCtxt) *RelaxNG

	RelaxNGFree func(schema *RelaxNG)

	RelaxNGDump func(output *FILE, schema *RelaxNG)

	RelaxNGDumpTree func(output *FILE, schema *RelaxNG)

	RelaxNGSetValidErrors func(ctxt *RelaxNGValidCtxt,
		err RelaxNGValidityErrorFunc,
		warn RelaxNGValidityWarningFunc, ctx *Void)

	RelaxNGGetValidErrors func(ctxt *RelaxNGValidCtxt,
		err *RelaxNGValidityErrorFunc,
		warn *RelaxNGValidityWarningFunc, ctx **Void) int

	RelaxNGSetValidStructuredErrors func(ctxt *RelaxNGValidCtxt,
		serror StructuredErrorFunc, ctx *Void)

	RelaxNGNewValidCtxt func(schema *RelaxNG) *RelaxNGValidCtxt

	RelaxNGFreeValidCtxt func(ctxt *RelaxNGValidCtxt)

	RelaxNGValidateDoc func(ctxt *RelaxNGValidCtxt, doc *Doc) int

	RelaxNGValidatePushElement func(
		ctxt *RelaxNGValidCtxt, doc *Doc, elem *Node) int

	RelaxNGValidatePushCData func(
		ctxt *RelaxNGValidCtxt, data string, len int) int

	RelaxNGValidatePopElement func(
		ctxt *RelaxNGValidCtxt, doc *Doc, elem *Node) int

	RelaxNGValidateFullElement func(
		ctxt *RelaxNGValidCtxt, doc *Doc, elem *Node) int

	NewTextReader func(
		input *ParserInputBuffer, URI string) *TextReader

	NewTextReaderFilename func(URI string) *TextReader

	FreeTextReader func(reader *TextReader)

	TextReaderSetup func(reader *TextReader,
		input *ParserInputBuffer, URL, encoding string,
		options int) int

	TextReaderRead func(reader *TextReader) int

	TextReaderReadInnerXml func(reader *TextReader) string

	TextReaderReadOuterXml func(reader *TextReader) string

	TextReaderReadString func(reader *TextReader) string

	TextReaderReadAttributeValue func(reader *TextReader) int

	TextReaderAttributeCount func(reader *TextReader) int

	TextReaderDepth func(reader *TextReader) int

	TextReaderHasAttributes func(reader *TextReader) int

	TextReaderHasValue func(reader *TextReader) int

	TextReaderIsDefault func(reader *TextReader) int

	TextReaderIsEmptyElement func(reader *TextReader) int

	TextReaderNodeType func(reader *TextReader) int

	TextReaderQuoteChar func(reader *TextReader) int

	TextReaderReadState func(reader *TextReader) int

	TextReaderIsNamespaceDecl func(reader *TextReader) int

	TextReaderConstBaseUri func(reader *TextReader) string

	TextReaderConstLocalName func(reader *TextReader) string

	TextReaderConstName func(reader *TextReader) string

	TextReaderConstNamespaceUri func(reader *TextReader) string

	TextReaderConstPrefix func(reader *TextReader) string

	TextReaderConstXmlLang func(reader *TextReader) string

	TextReaderConstString func(
		reader *TextReader, str string) string

	TextReaderConstValue func(reader *TextReader) string

	TextReaderBaseUri func(reader *TextReader) string

	TextReaderLocalName func(reader *TextReader) string

	TextReaderName func(reader *TextReader) string

	TextReaderNamespaceUri func(reader *TextReader) string

	TextReaderPrefix func(reader *TextReader) string

	TextReaderXmlLang func(reader *TextReader) string

	TextReaderValue func(reader *TextReader) string

	TextReaderClose func(reader *TextReader) int

	TextReaderGetAttributeNo func(
		reader *TextReader, no int) string

	TextReaderGetAttribute func(
		reader *TextReader, name string) string

	TextReaderGetAttributeNs func(reader *TextReader,
		localName, namespaceURI string) string

	TextReaderGetRemainder func(
		reader *TextReader) *ParserInputBuffer

	TextReaderLookupNamespace func(
		reader *TextReader, prefix string) string

	TextReaderMoveToAttributeNo func(
		reader *TextReader, no int) int

	TextReaderMoveToAttribute func(
		reader *TextReader, name string) int

	TextReaderMoveToAttributeNs func(reader *TextReader,
		localName, namespaceURI string) int

	TextReaderMoveToFirstAttribute func(reader *TextReader) int

	TextReaderMoveToNextAttribute func(reader *TextReader) int

	TextReaderMoveToElement func(reader *TextReader) int

	TextReaderNormalization func(reader *TextReader) int

	TextReaderConstEncoding func(reader *TextReader) string

	TextReaderSetParserProp func(
		reader *TextReader, prop, value int) int

	TextReaderGetParserProp func(
		reader *TextReader, prop int) int

	TextReaderCurrentNode func(reader *TextReader) *Node

	TextReaderGetParserLineNumber func(reader *TextReader) int

	TextReaderGetParserColumnNumber func(reader *TextReader) int

	TextReaderPreserve func(reader *TextReader) *Node

	TextReaderPreservePattern func(reader *TextReader,
		pattern string, namespaces *string) int

	TextReaderCurrentDoc func(reader *TextReader) *Doc

	TextReaderExpand func(reader *TextReader) *Node

	TextReaderNext func(reader *TextReader) int

	TextReaderNextSibling func(reader *TextReader) int

	TextReaderIsValid func(reader *TextReader) int

	TextReaderRelaxNGValidate func(
		reader *TextReader, rng string) int

	TextReaderRelaxNGValidateCtxt func(reader *TextReader,
		ctxt *RelaxNGValidCtxt, options int) int

	TextReaderRelaxNGSetSchema func(
		reader *TextReader, schema *RelaxNG) int

	TextReaderSchemaValidate func(
		reader *TextReader, xsd string) int

	TextReaderSchemaValidateCtxt func(reader *TextReader,
		ctxt *SchemaValidCtxt, options int) int

	TextReaderSetSchema func(
		reader *TextReader, schema *Schema) int

	TextReaderConstXmlVersion func(reader *TextReader) string

	TextReaderStandalone func(reader *TextReader) int

	TextReaderByteConsumed func(reader *TextReader) Long

	ReaderWalker func(doc *Doc) *TextReader

	ReaderForDoc func(
		cur, URL, encoding string, options int) *TextReader

	ReaderForFile func(
		filename, encoding string, options int) *TextReader

	ReaderForMemory func(buffer string, size int,
		URL, encoding string, options int) *TextReader

	ReaderForFd func(
		fd int, URL, encoding string, options int) *TextReader

	ReaderForIO func(ioread InputReadCallback,
		ioclose InputCloseCallback, ioctx *Void,
		URL, encoding string, options int) *TextReader

	ReaderNewWalker func(reader *TextReader, doc *Doc) int

	ReaderNewDoc func(reader *TextReader,
		cur, URL, encoding string, options int) int

	ReaderNewFile func(reader *TextReader,
		filename, encoding string, options int) int

	ReaderNewMemory func(reader *TextReader, buffer string,
		size int, URL, encoding string, options int) int

	ReaderNewFd func(reader *TextReader, fd int,
		URL, encoding string, options int) int

	ReaderNewIO func(reader *TextReader,
		ioread InputReadCallback, ioclose InputCloseCallback,
		ioctx *Void, URL, encoding string, options int) int

	TextReaderLocatorLineNumber func(
		locator *TextReaderLocator) int

	TextReaderLocatorBaseURI func(
		locator *TextReaderLocator) string

	TextReaderSetErrorHandler func(
		reader *TextReader, f TextReaderErrorFunc, arg *Void)

	TextReaderSetStructuredErrorHandler func(
		reader *TextReader, f StructuredErrorFunc, arg *Void)

	TextReaderGetErrorHandler func(
		reader *TextReader, f *TextReaderErrorFunc, arg **Void)

	NewCatalog func(sgml int) *Catalog

	LoadACatalog func(filename string) *Catalog

	LoadSGMLSuperCatalog func(filename string) *Catalog

	ConvertSGMLCatalog func(catal *Catalog) int

	ACatalogAdd func(
		catal *Catalog, typ, orig, replace string) int

	ACatalogRemove func(catal *Catalog, value string) int

	ACatalogResolve func(
		catal *Catalog, pubID, sysID string) string

	ACatalogResolveSystem func(
		catal *Catalog, sysID string) string

	ACatalogResolvePublic func(
		catal *Catalog, pubID string) string

	ACatalogResolveURI func(catal *Catalog, URI string) string

	ACatalogDump func(catal *Catalog, out *FILE)

	FreeCatalog func(catal *Catalog)

	CatalogIsEmpty func(catal *Catalog) int

	InitializeCatalog func()

	LoadCatalog func(filename string) int

	LoadCatalogs func(paths string)

	CatalogCleanup func()

	CatalogDump func(out *FILE)

	CatalogResolve func(pubID string, sysID string) string

	CatalogResolveSystem func(sysID string) string

	CatalogResolvePublic func(pubID string) string

	CatalogResolveURI func(URI string) string

	CatalogAdd func(typ string, orig string, replace string) int

	CatalogRemove func(value string) int

	ParseCatalogFile func(filename string) *Doc

	CatalogConvert func() int

	CatalogFreeLocal func(catalogs *Void)

	CatalogAddLocal func(catalogs *Void, URL string) *Void

	CatalogLocalResolve func(
		catalogs *Void, pubID, sysID string) string

	CatalogLocalResolveURI func(
		catalogs *Void, URI string) string

	CatalogSetDebug func(level int) int

	CatalogSetDefaultPrefer func(
		prefer CatalogPrefer) CatalogPrefer

	CatalogSetDefaults func(allow CatalogAllow)

	CatalogGetDefaults func() CatalogAllow

	CatalogGetSystem func(sysID string) string

	CatalogGetPublic func(pubID string) string

	DebugDumpString func(output *FILE, str string)

	DebugDumpAttr func(output *FILE, attr *Attr, depth int)

	DebugDumpAttrList func(output *FILE, attr *Attr, depth int)

	DebugDumpOneNode func(output *FILE, node *Node, depth int)

	DebugDumpNode func(output *FILE, node *Node, depth int)

	DebugDumpNodeList func(output *FILE, node *Node, depth int)

	DebugDumpDocumentHead func(output *FILE, doc *Doc)

	DebugDumpDocument func(output *FILE, doc *Doc)

	DebugDumpDTD func(output *FILE, dtd *Dtd)

	DebugDumpEntities func(output *FILE, doc *Doc)

	DebugCheckDocument func(output *FILE, doc *Doc) int

	LsOneNode func(output *FILE, node *Node)

	LsCountNode func(node *Node) int

	BoolToText func(boolval int) string

	ShellPrintXPathError func(errorType int, arg string)

	ShellPrintXPathResult func(list *XPathObject)

	ShellList func(
		ctxt *ShellCtxt, arg string, node, node2 *Node) int

	ShellBase func(
		ctxt *ShellCtxt, arg string, node, node2 *Node) int

	ShellDir func(
		ctxt *ShellCtxt, arg string, node, node2 *Node) int

	ShellLoad func(
		ctxt *ShellCtxt, filename string, node, node2 *Node) int

	ShellPrintNode func(node *Node)

	ShellCat func(
		ctxt *ShellCtxt, arg string, node, node2 *Node) int

	ShellWrite func(
		ctxt *ShellCtxt, filename string, node, node2 *Node) int

	ShellSave func(
		ctxt *ShellCtxt, filename string, node, node2 *Node) int

	ShellValidate func(
		ctxt *ShellCtxt, dtd string, node, node2 *Node) int

	ShellDu func(
		ctxt *ShellCtxt, arg string, tree, node2 *Node) int

	ShellPwd func(
		ctxt *ShellCtxt, buffer string, node, node2 *Node) int

	Shell func(doc *Doc, filename string,
		input ShellReadlineFunc, output *FILE)

	NewTextWriter func(out *OutputBuffer) *TextWriter

	NewTextWriterFilename func(
		uri string, compression int) *TextWriter

	NewTextWriterMemory func(
		buf *Buffer, compression int) *TextWriter

	NewTextWriterPushParser func(
		ctxt *ParserCtxt, compression int) *TextWriter

	NewTextWriterDoc func(
		doc **Doc, compression int) *TextWriter

	NewTextWriterTree func(
		doc *Doc, node *Node, compression int) *TextWriter

	FreeTextWriter func(writer *TextWriter)

	TextWriterStartDocument func(writer *TextWriter,
		version, encoding, standalone string) int

	TextWriterEndDocument func(writer *TextWriter) int

	TextWriterStartComment func(writer *TextWriter) int

	TextWriterEndComment func(writer *TextWriter) int

	TextWriterWriteFormatComment func(
		writer *TextWriter, format string, v ...VArg) int

	TextWriterWriteVFormatComment func(
		writer *TextWriter, format string, argptr VAList) int

	TextWriterWriteComment func(
		writer *TextWriter, Content string) int

	TextWriterStartElement func(
		writer *TextWriter, name string) int

	TextWriterStartElementNS func(writer *TextWriter,
		prefix, name, namespaceURI string) int

	TextWriterEndElement func(writer *TextWriter) int

	TextWriterFullEndElement func(writer *TextWriter) int

	TextWriterWriteFormatElement func(
		writer *TextWriter, name, format string, v ...VArg) int

	TextWriterWriteVFormatElement func(writer *TextWriter,
		name, format string, argptr VAList) int

	TextWriterWriteElement func(writer *TextWriter,
		name, Content string) int

	TextWriterWriteFormatElementNS func(writer *TextWriter,
		prefix, name, namespaceURI, format string, v ...VArg) int

	TextWriterWriteVFormatElementNS func(writer *TextWriter,
		prefix, name, namespaceURI, format string,
		argptr VAList) int

	TextWriterWriteElementNS func(writer *TextWriter,
		prefix, name, namespaceURI, Content string) int

	TextWriterWriteFormatRaw func(
		writer *TextWriter, format string, v ...VArg) int

	TextWriterWriteVFormatRaw func(
		writer *TextWriter, format string, argptr VAList) int

	TextWriterWriteRawLen func(
		writer *TextWriter, content string, len int) int

	TextWriterWriteRaw func(
		writer *TextWriter, content string) int

	TextWriterWriteFormatString func(
		writer *TextWriter, format string, v ...VArg) int

	TextWriterWriteVFormatString func(
		writer *TextWriter, format string, argptr VAList) int

	TextWriterWriteString func(
		writer *TextWriter, Content string) int

	TextWriterWriteBase64 func(
		writer *TextWriter, data string, start int, len int) int

	TextWriterWriteBinHex func(
		writer *TextWriter, data string, start int, len int) int

	TextWriterStartAttribute func(
		writer *TextWriter, name string) int

	TextWriterStartAttributeNS func(writer *TextWriter,
		prefix, name, namespaceURI string) int

	TextWriterEndAttribute func(writer *TextWriter) int

	TextWriterWriteFormatAttribute func(
		writer *TextWriter, name, format string, v ...VArg) int

	TextWriterWriteVFormatAttribute func(writer *TextWriter,
		name, format string, argptr VAList) int

	TextWriterWriteAttribute func(
		writer *TextWriter, name, Content string) int

	TextWriterWriteFormatAttributeNS func(writer *TextWriter,
		prefix, name, namespaceURI, format string, v ...VArg) int

	TextWriterWriteVFormatAttributeNS func(writer *TextWriter,
		prefix, name, namespaceURI, format string,
		argptr VAList) int

	TextWriterWriteAttributeNS func(writer *TextWriter,
		prefix, name, namespaceURI, Content string) int

	TextWriterStartPI func(writer *TextWriter, target string) int

	TextWriterEndPI func(writer *TextWriter) int

	TextWriterWriteFormatPI func(writer *TextWriter,
		target, format string, v ...VArg) int

	TextWriterWriteVFormatPI func(writer *TextWriter,
		target, format string, argptr VAList) int

	TextWriterWritePI func(
		writer *TextWriter, target, content string) int

	TextWriterStartCDATA func(writer *TextWriter) int

	TextWriterEndCDATA func(writer *TextWriter) int

	TextWriterWriteFormatCDATA func(
		writer *TextWriter, format string, v ...VArg) int

	TextWriterWriteVFormatCDATA func(
		writer *TextWriter, format string, argptr VAList) int

	TextWriterWriteCDATA func(
		writer *TextWriter, content string) int

	TextWriterStartDTD func(
		writer *TextWriter, name, pubid, sysid string) int

	TextWriterEndDTD func(writer *TextWriter) int

	TextWriterWriteFormatDTD func(writer *TextWriter,
		name, pubid, sysid, format string, v ...VArg) int

	TextWriterWriteVFormatDTD func(writer *TextWriter,
		name, pubid, sysid, format string, argptr VAList) int

	TextWriterWriteDTD func(writer *TextWriter,
		name, pubid, sysid, subset string) int

	TextWriterStartDTDElement func(
		writer *TextWriter, name string) int

	TextWriterEndDTDElement func(writer *TextWriter) int

	TextWriterWriteFormatDTDElement func(
		writer *TextWriter, name, format string, v ...VArg) int

	TextWriterWriteVFormatDTDElement func(writer *TextWriter,
		name, format string, argptr VAList) int

	TextWriterWriteDTDElement func(
		writer *TextWriter, name, Content string) int

	TextWriterStartDTDAttlist func(
		writer *TextWriter, name string) int

	TextWriterEndDTDAttlist func(writer *TextWriter) int

	TextWriterWriteFormatDTDAttlist func(writer *TextWriter,
		name, format string, v ...VArg) int

	TextWriterWriteVFormatDTDAttlist func(writer *TextWriter,
		name, format string, argptr VAList) int

	TextWriterWriteDTDAttlist func(
		writer *TextWriter, name, Content string) int

	TextWriterStartDTDEntity func(
		writer *TextWriter, pe int, name string) int

	TextWriterEndDTDEntity func(writer *TextWriter) int

	TextWriterWriteFormatDTDInternalEntity func(
		writer *TextWriter, pe int,
		name, format string, v ...VArg) int

	TextWriterWriteVFormatDTDInternalEntity func(
		writer *TextWriter, pe int,
		name, format string, argptr VAList) int

	TextWriterWriteDTDInternalEntity func(
		writer *TextWriter, pe int, name, content string) int

	TextWriterWriteDTDExternalEntity func(
		writer *TextWriter, pe int,
		name, pubid, sysid, ndataid string) int

	TextWriterWriteDTDExternalEntityContents func(
		writer *TextWriter, pubid, sysid, ndataid string) int

	TextWriterWriteDTDEntity func(writer *TextWriter,
		pe int, name, pubid, sysid, ndataid, content string) int

	TextWriterWriteDTDNotation func(
		writer *TextWriter, name, pubid, sysid string) int

	TextWriterSetIndent func(writer *TextWriter, indent int) int

	TextWriterSetIndentString func(
		writer *TextWriter, str string) int

	TextWriterSetQuoteChar func(
		writer *TextWriter, quotechar Char) int

	TextWriterFlush func(writer *TextWriter) int

	NanoFTPInit func()

	NanoFTPCleanup func()

	NanoFTPNewCtxt func(URL string) *Void

	NanoFTPFreeCtxt func(ctx *Void)

	NanoFTPConnectTo func(server string, port int) *Void

	NanoFTPOpen func(URL string) *Void

	NanoFTPConnect func(ctx *Void) int

	NanoFTPClose func(ctx *Void) int

	NanoFTPQuit func(ctx *Void) int

	NanoFTPScanProxy func(URL string)

	NanoFTPProxy func(
		host string, port int, user, passwd string, typ int)

	NanoFTPUpdateURL func(ctx *Void, URL string) int

	NanoFTPGetResponse func(ctx *Void) int

	NanoFTPCheckResponse func(ctx *Void) int

	NanoFTPCwd func(ctx *Void, directory string) int

	NanoFTPDele func(ctx *Void, file string) int

	NanoFTPGetConnection func(ctx *Void) SOCKET

	NanoFTPCloseConnection func(ctx *Void) int

	NanoFTPList func(ctx *Void, callback FtpListCallback,
		userData *Void, filename string) int

	NanoFTPGetSocket func(ctx *Void, filename string) SOCKET

	NanoFTPGet func(ctx *Void, callback FtpDataCallback,
		userData *Void, filename string) int

	NanoFTPRead func(ctx, dest *Void, len int) int

	NanoHTTPInit func()

	NanoHTTPCleanup func()

	NanoHTTPScanProxy func(URL string)

	NanoHTTPFetch func(
		URL, filename string, contentType *string) int

	NanoHTTPMethod func(URL, method, input string,
		contentType *string, headers string, ilen int) *Void

	NanoHTTPMethodRedir func(
		URL, method, input string, contentType, redir *string,
		headers string, ilen int) *Void

	NanoHTTPOpen func(URL string, contentType *string) *Void

	NanoHTTPOpenRedir func(
		URL string, contentType *string, redir *string) *Void

	NanoHTTPReturnCode func(ctx *Void) int

	NanoHTTPAuthHeader func(ctx *Void) string

	NanoHTTPRedir func(ctx *Void) string

	NanoHTTPContentLength func(ctx *Void) int

	NanoHTTPEncoding func(ctx *Void) string

	NanoHTTPMimeType func(ctx *Void) string

	NanoHTTPRead func(ctx *Void, dest *Void, len int) int

	NanoHTTPSave func(ctxt *Void, filename string) int

	NanoHTTPClose func(ctx *Void)

	FreePattern func(comp *Pattern)

	FreePatternList func(comp *Pattern)

	Patterncompile func(pattern string, dict *Dict,
		flags int, namespaces *string) *Pattern

	PatternMatch func(comp *Pattern, node *Node) int

	PatternStreamable func(comp *Pattern) int

	PatternMaxDepth func(comp *Pattern) int

	PatternMinDepth func(comp *Pattern) int

	PatternFromRoot func(comp *Pattern) int

	PatternGetStreamCtxt func(comp *Pattern) *StreamCtxt

	FreeStreamCtxt func(stream *StreamCtxt)

	StreamPushNode func(
		stream *StreamCtxt, name, ns string, nodeType int) int

	StreamPush func(stream *StreamCtxt, name, ns string) int

	StreamPushAttr func(stream *StreamCtxt, name, ns string) int

	StreamPop func(stream *StreamCtxt) int

	StreamWantsAnyNode func(stream *StreamCtxt) int

	SaveToFd func(
		fd int, encoding string, options int) *SaveCtxt

	SaveToFilename func(
		filename, encoding string, options int) *SaveCtxt

	SaveToBuffer func(
		buffer *Buffer, encoding string, options int) *SaveCtxt

	SaveToIO func(iowrite OutputWriteCallback,
		ioclose OutputCloseCallback,
		ioctx *Void, encoding string, options int) *SaveCtxt

	SaveDoc func(ctxt *SaveCtxt, doc *Doc) Long

	SaveTree func(ctxt *SaveCtxt, node *Node) Long

	SaveFlush func(ctxt *SaveCtxt) int

	SaveClose func(ctxt *SaveCtxt) int

	SaveSetEscape func(
		ctxt *SaveCtxt, escape CharEncodingOutputFunc) int

	SaveSetAttrEscape func(
		ctxt *SaveCtxt, escape CharEncodingOutputFunc) int

	CreateURI func() *URI

	BuildURI func(URI, base string) string

	BuildRelativeURI func(URI, base string) string

	ParseURI func(str string) *URI

	ParseURIRaw func(str string, raw int) *URI

	ParseURIReference func(uri *URI, str string) int

	SaveUri func(uri *URI) string

	PrintURI func(stream *FILE, uri *URI)

	URIEscapeStr func(str, list string) string

	URIUnescapeString func(
		str string, leng int, target string) string

	NormalizeURIPath func(path string) int

	URIEscape func(str string) string

	FreeURI func(uri *URI)

	CanonicPath func(path string) string

	PathToURI func(path string) string

	RaiseError func(schannel StructuredErrorFunc,
		channel GenericErrorFunc, data, ctx, node *Void,
		domain, code int, level ErrorLevel, file string,
		line int, str1, str2, str3 string, int1, col int,
		msg string, v ...VArg)

	SimpleError func(
		domain, code int, node *Node, msg, extra string)

	ErrEncoding func(ctxt *ParserCtxt,
		xmlerr ParserErrors, msg, str1, str2 string)

	ModuleOpen func(
		filename string, XmlModuleOption int) *Module

	ModuleSymbol func(
		module *Module, name string, result **Void) int

	ModuleClose func(module *Module) int

	ModuleFree func(module *Module) int

	DocbEncodeEntities func(out *UnsignedChar, outlen *int,
		in *UnsignedChar, inlen *int, quoteChar int) int

	DocbSAXParseDoc func(cur, encoding *Char,
		sax *DocbSAXHandler, userData *Void) *DocbDoc

	DocbParseDoc func(cur, encoding *Char) *DocbDoc

	DocbSAXParseFile func(filename, encoding *Char,
		sax *DocbSAXHandler, userData *Void) *DocbDoc

	DocbParseFile func(filename *Char, encoding *Char) *DocbDoc

	DocbFreeParserCtxt func(ctxt *DocbParserCtxt)

	DocbCreatePushParserCtxt func(sax *DocbSAXHandler,
		userData *Void, chunk *Char, size int,
		filename *Char, enc CharEncoding) *DocbParserCtxt

	DocbParseChunk func(ctxt *DocbParserCtxt,
		chunk *Char, size int, terminate int) int

	DocbCreateFileParserCtxt func(
		filename, encoding *Char) *DocbParserCtxt

	DocbParseDocument func(ctxt *DocbParserCtxt) int

	InitxmlDefaultSAXHandler func(
		hdlr *SAXHandlerV1, warning int)

	C14NDocSaveTo func(doc *Doc, nodes *NodeSet, mode int,
		inclusiveNsPrefixes **Char, withComments int,
		buf *OutputBuffer) int

	C14NDocDumpMemory func(doc *Doc, nodes *NodeSet, mode int,
		inclusiveNsPrefixes **Char, withComments int,
		docTxtPtr **Char) int

	C14NDocSave func(doc *Doc, nodes *NodeSet, mode int,
		inclusiveNsPrefixes **Char, withComments int,
		filename *Char, compression int) int

	C14NExecute func(doc *Doc,
		is_visibleCallback C14NIsVisibleCallback,
		userData *Void, mode int, inclusiveNsPrefixes **Char,
		withComments int, buf *OutputBuffer) int

	ErrMemory func(ctxt *ParserCtxt, extra *Char)

	ParserInputBufferCreateFilenameDefault func(
		ParserInputBufferCreateFilenameFunc) ParserInputBufferCreateFilenameFunc

	OutputBufferCreateFilenameDefault func(
		OutputBufferCreateFilenameFunc) OutputBufferCreateFilenameFunc

	IsBaseChar func(ch UnsignedInt) int

	IsBlank func(ch UnsignedInt) int

	IsChar func(ch UnsignedInt) int

	IsCombining func(ch UnsignedInt) int

	IsDigit func(ch UnsignedInt) int

	IsExtender func(ch UnsignedInt) int

	IsIdeographic func(ch UnsignedInt) int

	IsPubidChar func(ch UnsignedInt) int
)

var dll = "libxml2-2.dll"

var apiList = outside.Apis{
	{"UTF8ToHtml", &UTF8ToHtml},
	{"UTF8Toisolat1", &UTF8Toisolat1},
	{"__docbDefaultSAXHandler", &DocbDefaultSAXHandler},
	{"__htmlDefaultSAXHandler", &HtmlDefaultSAXHandler},
	{"__oldXMLWDcompatibility", &OldXMLWDcompatibility},
	{"__xmlBufferAllocScheme", &BufferAllocScheme},
	{"__xmlDefaultBufferSize", &DefaultBufferSize},
	{"__xmlDefaultSAXHandler", &DefaultSAXHandler},
	{"__xmlDefaultSAXLocator", &DefaultSAXLocator},
	{"__xmlDeregisterNodeDefaultValue", &DeregisterNodeDefaultValue},
	{"__xmlDoValidityCheckingDefaultValue", &DoValidityCheckingDefaultValue},
	{"__xmlErrEncoding", &ErrEncoding},
	{"__xmlGenericError", &GenericError},
	{"__xmlGenericErrorContext", &GenericErrorContext},
	{"__xmlGetWarningsDefaultValue", &GetWarningsDefaultValue},
	{"__xmlIndentTreeOutput", &IndentTreeOutput},
	{"__xmlKeepBlanksDefaultValue", &KeepBlanksDefaultValue},
	{"__xmlLastError", &LastError},
	{"__xmlLineNumbersDefaultValue", &LineNumbersDefaultValue},
	{"__xmlLoadExtDtdDefaultValue", &LoadExtDtdDefaultValue},
	{"__xmlOutputBufferCreateFilenameValue", &OutputBufferCreateFilenameValue},
	{"__xmlParserDebugEntities", &ParserDebugEntities},
	{"__xmlParserInputBufferCreateFilenameValue", &ParserInputBufferCreateFilenameValue},
	{"__xmlParserVersion", &ParserVersion},
	{"__xmlPedanticParserDefaultValue", &PedanticParserDefaultValue},
	{"__xmlRaiseError", &RaiseError},
	{"__xmlRegisterNodeDefaultValue", &RegisterNodeDefaultValue},
	{"__xmlSaveNoEmptyTags", &SaveNoEmptyTags},
	{"__xmlSimpleError", &SimpleError},
	{"__xmlStructuredError", &StructuredError},
	{"__xmlStructuredErrorContext", &StructuredErrorContext},
	{"__xmlSubstituteEntitiesDefaultValue", &SubstituteEntitiesDefaultValue},
	{"__xmlTreeIndentString", &TreeIndentString},
	{"attribute", &Attribute},
	{"attributeDecl", &AttributeDecl},
	{"cdataBlock", &CdataBlock},
	{"characters", &Characters},
	{"checkNamespace", &CheckNamespace},
	{"comment", &Comment},
	{"docbCreateFileParserCtxt", &DocbCreateFileParserCtxt},
	{"docbCreatePushParserCtxt", &DocbCreatePushParserCtxt},
	{"docbDefaultSAXHandlerInit", &DocbDefaultSAXHandlerInit},
	{"docbEncodeEntities", &DocbEncodeEntities},
	{"docbFreeParserCtxt", &DocbFreeParserCtxt},
	{"docbParseChunk", &DocbParseChunk},
	{"docbParseDoc", &DocbParseDoc},
	{"docbParseDocument", &DocbParseDocument},
	{"docbParseFile", &DocbParseFile},
	{"docbSAXParseDoc", &DocbSAXParseDoc},
	{"docbSAXParseFile", &DocbSAXParseFile},
	{"elementDecl", &ElementDecl},
	{"endDocument", &EndDocument},
	{"endElement", &EndElement},
	{"entityDecl", &EntityDecl},
	{"externalSubset", &ExternalSubset},
	{"getColumnNumber", &GetColumnNumber},
	{"getEntity", &GetEntity},
	{"getLineNumber", &GetLineNumber},
	{"getNamespace", &GetNamespace},
	{"getParameterEntity", &GetParameterEntity},
	{"getPublicId", &GetPublicId},
	{"getSystemId", &GetSystemId},
	{"globalNamespace", &GlobalNamespace},
	{"hasExternalSubset", &HasExternalSubset},
	{"hasInternalSubset", &HasInternalSubset},
	{"htmlAttrAllowed", &HtmlAttrAllowed},
	{"htmlAutoCloseTag", &HtmlAutoCloseTag},
	{"htmlCreateFileParserCtxt", &HtmlCreateFileParserCtxt},
	{"htmlCreateMemoryParserCtxt", &HtmlCreateMemoryParserCtxt},
	{"htmlCreatePushParserCtxt", &HtmlCreatePushParserCtxt},
	{"htmlCtxtReadDoc", &HtmlCtxtReadDoc},
	{"htmlCtxtReadFd", &HtmlCtxtReadFd},
	{"htmlCtxtReadFile", &HtmlCtxtReadFile},
	{"htmlCtxtReadIO", &HtmlCtxtReadIO},
	{"htmlCtxtReadMemory", &HtmlCtxtReadMemory},
	{"htmlCtxtReset", &HtmlCtxtReset},
	{"htmlCtxtUseOptions", &HtmlCtxtUseOptions},
	{"htmlDefaultSAXHandlerInit", &HtmlDefaultSAXHandlerInit},
	{"htmlDocContentDumpFormatOutput", &HtmlDocContentDumpFormatOutput},
	{"htmlDocContentDumpOutput", &HtmlDocContentDumpOutput},
	{"htmlDocDump", &HtmlDocDump},
	{"htmlDocDumpMemory", &HtmlDocDumpMemory},
	{"htmlDocDumpMemoryFormat", &HtmlDocDumpMemoryFormat},
	{"htmlElementAllowedHere", &HtmlElementAllowedHere},
	{"htmlElementStatusHere", &HtmlElementStatusHere},
	{"htmlEncodeEntities", &HtmlEncodeEntities},
	{"htmlEntityLookup", &HtmlEntityLookup},
	{"htmlEntityValueLookup", &HtmlEntityValueLookup},
	{"htmlFreeParserCtxt", &HtmlFreeParserCtxt},
	{"htmlGetMetaEncoding", &HtmlGetMetaEncoding},
	{"htmlHandleOmittedElem", &HtmlHandleOmittedElem},
	{"htmlInitAutoClose", &HtmlInitAutoClose},
	{"htmlIsAutoClosed", &HtmlIsAutoClosed},
	{"htmlIsBooleanAttr", &HtmlIsBooleanAttr},
	{"htmlIsScriptAttribute", &HtmlIsScriptAttribute},
	{"htmlNewDoc", &HtmlNewDoc},
	{"htmlNewDocNoDtD", &HtmlNewDocNoDtD},
	{"htmlNewParserCtxt", &HtmlNewParserCtxt},
	{"htmlNodeDump", &HtmlNodeDump},
	{"htmlNodeDumpFile", &HtmlNodeDumpFile},
	{"htmlNodeDumpFileFormat", &HtmlNodeDumpFileFormat},
	{"htmlNodeDumpFormatOutput", &HtmlNodeDumpFormatOutput},
	{"htmlNodeDumpOutput", &HtmlNodeDumpOutput},
	{"htmlNodeStatus", &HtmlNodeStatus},
	{"htmlParseCharRef", &HtmlParseCharRef},
	{"htmlParseChunk", &HtmlParseChunk},
	{"htmlParseDoc", &HtmlParseDoc},
	{"htmlParseDocument", &HtmlParseDocument},
	{"htmlParseElement", &HtmlParseElement},
	{"htmlParseEntityRef", &HtmlParseEntityRef},
	{"htmlParseFile", &HtmlParseFile},
	{"htmlReadDoc", &HtmlReadDoc},
	{"htmlReadFd", &HtmlReadFd},
	{"htmlReadFile", &HtmlReadFile},
	{"htmlReadIO", &HtmlReadIO},
	{"htmlReadMemory", &HtmlReadMemory},
	{"htmlSAXParseDoc", &HtmlSAXParseDoc},
	{"htmlSAXParseFile", &HtmlSAXParseFile},
	{"htmlSaveFile", &HtmlSaveFile},
	{"htmlSaveFileEnc", &HtmlSaveFileEnc},
	{"htmlSaveFileFormat", &HtmlSaveFileFormat},
	{"htmlSetMetaEncoding", &HtmlSetMetaEncoding},
	{"htmlTagLookup", &HtmlTagLookup},
	{"ignorableWhitespace", &IgnorableWhitespace},
	{"initGenericErrorDefaultFunc", &InitGenericErrorDefaultFunc},
	{"initdocbDefaultSAXHandler", &InitdocbDefaultSAXHandler},
	{"inithtmlDefaultSAXHandler", &InithtmlDefaultSAXHandler},
	{"initxmlDefaultSAXHandler", &InitxmlDefaultSAXHandler},
	{"inputPop", &InputPop},
	{"inputPush", &InputPush},
	{"internalSubset", &InternalSubset},
	{"isStandalone", &IsStandalone},
	{"isolat1ToUTF8", &Isolat1ToUTF8},
	{"namePop", &NamePop},
	{"namePush", &NamePush},
	{"namespaceDecl", &NamespaceDecl},
	{"nodePop", &NodePop},
	{"nodePush", &NodePush},
	{"notationDecl", &NotationDecl},
	{"processingInstruction", &ProcessingInstruction},
	{"reference", &Reference},
	{"resolveEntity", &ResolveEntity},
	{"setDocumentLocator", &SetDocumentLocator},
	{"setNamespace", &SetNamespace},
	{"startDocument", &StartDocument},
	{"startElement", &StartElement},
	{"unparsedEntityDecl", &UnparsedEntityDecl},
	{"valuePop", &ValuePop},
	{"valuePush", &ValuePush},
	{"xlinkGetDefaultDetect", &XlinkGetDefaultDetect},
	{"xlinkGetDefaultHandler", &XlinkGetDefaultHandler},
	{"xlinkIsLink", &XlinkIsLink},
	{"xlinkSetDefaultDetect", &XlinkSetDefaultDetect},
	{"xlinkSetDefaultHandler", &XlinkSetDefaultHandler},
	{"xmlACatalogAdd", &ACatalogAdd},
	{"xmlACatalogDump", &ACatalogDump},
	{"xmlACatalogRemove", &ACatalogRemove},
	{"xmlACatalogResolve", &ACatalogResolve},
	{"xmlACatalogResolvePublic", &ACatalogResolvePublic},
	{"xmlACatalogResolveSystem", &ACatalogResolveSystem},
	{"xmlACatalogResolveURI", &ACatalogResolveURI},
	{"xmlAddAttributeDecl", &AddAttributeDecl},
	{"xmlAddChild", &AddChild},
	{"xmlAddChildList", &AddChildList},
	{"xmlAddDocEntity", &AddDocEntity},
	{"xmlAddDtdEntity", &AddDtdEntity},
	{"xmlAddElementDecl", &AddElementDecl},
	{"xmlAddEncodingAlias", &AddEncodingAlias},
	{"xmlAddID", &AddID},
	{"xmlAddNextSibling", &AddNextSibling},
	{"xmlAddNotationDecl", &AddNotationDecl},
	{"xmlAddPrevSibling", &AddPrevSibling},
	{"xmlAddRef", &AddRef},
	{"xmlAddSibling", &AddSibling},
	{"xmlAllocOutputBuffer", &AllocOutputBuffer},
	{"xmlAllocParserInputBuffer", &AllocParserInputBuffer},
	{"xmlAttrSerializeTxtContent", &AttrSerializeTxtContent},
	{"xmlAutomataCompile", &AutomataCompile},
	{"xmlAutomataGetInitState", &AutomataGetInitState},
	{"xmlAutomataIsDeterminist", &AutomataIsDeterminist},
	{"xmlAutomataNewAllTrans", &AutomataNewAllTrans},
	{"xmlAutomataNewCountTrans", &AutomataNewCountTrans},
	{"xmlAutomataNewCountTrans2", &AutomataNewCountTrans2},
	{"xmlAutomataNewCountedTrans", &AutomataNewCountedTrans},
	{"xmlAutomataNewCounter", &AutomataNewCounter},
	{"xmlAutomataNewCounterTrans", &AutomataNewCounterTrans},
	{"xmlAutomataNewEpsilon", &AutomataNewEpsilon},
	{"xmlAutomataNewNegTrans", &AutomataNewNegTrans},
	{"xmlAutomataNewOnceTrans", &AutomataNewOnceTrans},
	{"xmlAutomataNewOnceTrans2", &AutomataNewOnceTrans2},
	{"xmlAutomataNewState", &AutomataNewState},
	{"xmlAutomataNewTransition", &AutomataNewTransition},
	{"xmlAutomataNewTransition2", &AutomataNewTransition2},
	{"xmlAutomataSetFinalState", &AutomataSetFinalState},
	{"xmlBoolToText", &BoolToText},
	{"xmlBufferAdd", &BufferAdd},
	{"xmlBufferAddHead", &BufferAddHead},
	{"xmlBufferCCat", &BufferCCat},
	{"xmlBufferCat", &BufferCat},
	{"xmlBufferContent", &BufferContent},
	{"xmlBufferCreate", &BufferCreate},
	{"xmlBufferCreateSize", &BufferCreateSize},
	{"xmlBufferCreateStatic", &BufferCreateStatic},
	{"xmlBufferDump", &BufferDump},
	{"xmlBufferEmpty", &BufferEmpty},
	{"xmlBufferFree", &BufferFree},
	{"xmlBufferGrow", &BufferGrow},
	{"xmlBufferLength", &BufferLength},
	{"xmlBufferResize", &BufferResize},
	{"xmlBufferSetAllocationScheme", &BufferSetAllocationScheme},
	{"xmlBufferShrink", &BufferShrink},
	{"xmlBufferWriteCHAR", &BufferWriteCHAR},
	{"xmlBufferWriteChar", &BufferWriteChar},
	{"xmlBufferWriteQuotedString", &BufferWriteQuotedString},
	{"xmlBuildQName", &BuildQName},
	{"xmlBuildRelativeURI", &BuildRelativeURI},
	{"xmlBuildURI", &BuildURI},
	{"xmlByteConsumed", &ByteConsumed},
	{"xmlC14NDocDumpMemory", &C14NDocDumpMemory},
	{"xmlC14NDocSave", &C14NDocSave},
	{"xmlC14NDocSaveTo", &C14NDocSaveTo},
	{"xmlC14NExecute", &C14NExecute},
	{"xmlCanonicPath", &CanonicPath},
	{"xmlCatalogAdd", &CatalogAdd},
	{"xmlCatalogAddLocal", &CatalogAddLocal},
	{"xmlCatalogCleanup", &CatalogCleanup},
	{"xmlCatalogConvert", &CatalogConvert},
	{"xmlCatalogDump", &CatalogDump},
	{"xmlCatalogFreeLocal", &CatalogFreeLocal},
	{"xmlCatalogGetDefaults", &CatalogGetDefaults},
	{"xmlCatalogGetPublic", &CatalogGetPublic},
	{"xmlCatalogGetSystem", &CatalogGetSystem},
	{"xmlCatalogIsEmpty", &CatalogIsEmpty},
	{"xmlCatalogLocalResolve", &CatalogLocalResolve},
	{"xmlCatalogLocalResolveURI", &CatalogLocalResolveURI},
	{"xmlCatalogRemove", &CatalogRemove},
	{"xmlCatalogResolve", &CatalogResolve},
	{"xmlCatalogResolvePublic", &CatalogResolvePublic},
	{"xmlCatalogResolveSystem", &CatalogResolveSystem},
	{"xmlCatalogResolveURI", &CatalogResolveURI},
	{"xmlCatalogSetDebug", &CatalogSetDebug},
	{"xmlCatalogSetDefaultPrefer", &CatalogSetDefaultPrefer},
	{"xmlCatalogSetDefaults", &CatalogSetDefaults},
	{"xmlCharEncCloseFunc", &CharEncCloseFunc},
	{"xmlCharEncFirstLine", &CharEncFirstLine},
	{"xmlCharEncInFunc", &CharEncInFunc},
	{"xmlCharEncOutFunc", &CharEncOutFunc},
	{"xmlCharInRange", &CharInRange},
	{"xmlCharStrdup", &CharStrdup},
	{"xmlCharStrndup", &CharStrndup},
	{"xmlCheckFilename", &CheckFilename},
	{"xmlCheckHTTPInput", &CheckHTTPInput},
	{"xmlCheckLanguageID", &CheckLanguageID},
	{"xmlCheckUTF8", &CheckUTF8},
	{"xmlCheckVersion", &CheckVersion},
	{"xmlChildElementCount", &ChildElementCount},
	{"xmlCleanupCharEncodingHandlers", &CleanupCharEncodingHandlers},
	{"xmlCleanupEncodingAliases", &CleanupEncodingAliases},
	{"xmlCleanupGlobals", &CleanupGlobals},
	{"xmlCleanupInputCallbacks", &CleanupInputCallbacks},
	{"xmlCleanupMemory", &CleanupMemory},
	{"xmlCleanupOutputCallbacks", &CleanupOutputCallbacks},
	{"xmlCleanupParser", &CleanupParser},
	{"xmlCleanupPredefinedEntities", &CleanupPredefinedEntities},
	{"xmlCleanupThreads", &CleanupThreads},
	{"xmlClearNodeInfoSeq", &ClearNodeInfoSeq},
	{"xmlClearParserCtxt", &ClearParserCtxt},
	{"xmlConvertSGMLCatalog", &ConvertSGMLCatalog},
	{"xmlCopyAttributeTable", &CopyAttributeTable},
	{"xmlCopyChar", &CopyChar},
	{"xmlCopyCharMultiByte", &CopyCharMultiByte},
	{"xmlCopyDoc", &CopyDoc},
	{"xmlCopyDocElementContent", &CopyDocElementContent},
	{"xmlCopyDtd", &CopyDtd},
	{"xmlCopyElementContent", &CopyElementContent},
	{"xmlCopyElementTable", &CopyElementTable},
	{"xmlCopyEntitiesTable", &CopyEntitiesTable},
	{"xmlCopyEnumeration", &CopyEnumeration},
	{"xmlCopyError", &CopyError},
	{"xmlCopyNamespace", &CopyNamespace},
	{"xmlCopyNamespaceList", &CopyNamespaceList},
	{"xmlCopyNode", &CopyNode},
	{"xmlCopyNodeList", &CopyNodeList},
	{"xmlCopyNotationTable", &CopyNotationTable},
	{"xmlCopyProp", &CopyProp},
	{"xmlCopyPropList", &CopyPropList},
	{"xmlCreateDocParserCtxt", &CreateDocParserCtxt},
	{"xmlCreateEntitiesTable", &CreateEntitiesTable},
	{"xmlCreateEntityParserCtxt", &CreateEntityParserCtxt},
	{"xmlCreateEnumeration", &CreateEnumeration},
	{"xmlCreateFileParserCtxt", &CreateFileParserCtxt},
	{"xmlCreateIOParserCtxt", &CreateIOParserCtxt},
	{"xmlCreateIntSubset", &CreateIntSubset},
	{"xmlCreateMemoryParserCtxt", &CreateMemoryParserCtxt},
	{"xmlCreatePushParserCtxt", &CreatePushParserCtxt},
	{"xmlCreateURI", &CreateURI},
	{"xmlCreateURLParserCtxt", &CreateURLParserCtxt},
	{"xmlCtxtGetLastError", &CtxtGetLastError},
	{"xmlCtxtReadDoc", &CtxtReadDoc},
	{"xmlCtxtReadFd", &CtxtReadFd},
	{"xmlCtxtReadFile", &CtxtReadFile},
	{"xmlCtxtReadIO", &CtxtReadIO},
	{"xmlCtxtReadMemory", &CtxtReadMemory},
	{"xmlCtxtReset", &CtxtReset},
	{"xmlCtxtResetLastError", &CtxtResetLastError},
	{"xmlCtxtResetPush", &CtxtResetPush},
	{"xmlCtxtUseOptions", &CtxtUseOptions},
	{"xmlCurrentChar", &CurrentChar},
	{"xmlDOMWrapAdoptNode", &DOMWrapAdoptNode},
	{"xmlDOMWrapCloneNode", &DOMWrapCloneNode},
	{"xmlDOMWrapFreeCtxt", &DOMWrapFreeCtxt},
	{"xmlDOMWrapNewCtxt", &DOMWrapNewCtxt},
	{"xmlDOMWrapReconcileNamespaces", &DOMWrapReconcileNamespaces},
	{"xmlDOMWrapRemoveNode", &DOMWrapRemoveNode},
	{"xmlDebugCheckDocument", &DebugCheckDocument},
	{"xmlDebugDumpAttr", &DebugDumpAttr},
	{"xmlDebugDumpAttrList", &DebugDumpAttrList},
	{"xmlDebugDumpDTD", &DebugDumpDTD},
	{"xmlDebugDumpDocument", &DebugDumpDocument},
	{"xmlDebugDumpDocumentHead", &DebugDumpDocumentHead},
	{"xmlDebugDumpEntities", &DebugDumpEntities},
	{"xmlDebugDumpNode", &DebugDumpNode},
	{"xmlDebugDumpNodeList", &DebugDumpNodeList},
	{"xmlDebugDumpOneNode", &DebugDumpOneNode},
	{"xmlDebugDumpString", &DebugDumpString},
	{"xmlDecodeEntities", &DecodeEntities},
	{"xmlDefaultSAXHandlerInit", &DefaultSAXHandlerInit},
	{"xmlDelEncodingAlias", &DelEncodingAlias},
	{"xmlDeregisterNodeDefault", &DeregisterNodeDefault},
	{"xmlDetectCharEncoding", &DetectCharEncoding},
	{"xmlDictCleanup", &DictCleanup},
	{"xmlDictCreate", &DictCreate},
	{"xmlDictCreateSub", &DictCreateSub},
	{"xmlDictExists", &DictExists},
	{"xmlDictFree", &DictFree},
	{"xmlDictLookup", &DictLookup},
	{"xmlDictOwns", &DictOwns},
	{"xmlDictQLookup", &DictQLookup},
	{"xmlDictReference", &DictReference},
	{"xmlDictSize", &DictSize},
	{"xmlDocCopyNode", &DocCopyNode},
	{"xmlDocCopyNodeList", &DocCopyNodeList},
	{"xmlDocDump", &DocDump},
	{"xmlDocDumpFormatMemory", &DocDumpFormatMemory},
	{"xmlDocDumpFormatMemoryEnc", &DocDumpFormatMemoryEnc},
	{"xmlDocDumpMemory", &DocDumpMemory},
	{"xmlDocDumpMemoryEnc", &DocDumpMemoryEnc},
	{"xmlDocFormatDump", &DocFormatDump},
	{"xmlDocGetRootElement", &DocGetRootElement},
	{"xmlDocSetRootElement", &DocSetRootElement},
	{"xmlDumpAttributeDecl", &DumpAttributeDecl},
	{"xmlDumpAttributeTable", &DumpAttributeTable},
	{"xmlDumpElementDecl", &DumpElementDecl},
	{"xmlDumpElementTable", &DumpElementTable},
	{"xmlDumpEntitiesTable", &DumpEntitiesTable},
	{"xmlDumpEntityDecl", &DumpEntityDecl},
	{"xmlDumpNotationDecl", &DumpNotationDecl},
	{"xmlDumpNotationTable", &DumpNotationTable},
	{"xmlElemDump", &ElemDump},
	{"xmlEncodeEntities", &EncodeEntities},
	{"xmlEncodeEntitiesReentrant", &EncodeEntitiesReentrant},
	{"xmlEncodeSpecialChars", &EncodeSpecialChars},
	{"xmlErrMemory", &ErrMemory},
	{"xmlExpCtxtNbCons", &ExpCtxtNbCons},
	{"xmlExpCtxtNbNodes", &ExpCtxtNbNodes},
	{"xmlExpDump", &ExpDump},
	{"xmlExpExpDerive", &ExpExpDerive},
	{"xmlExpFree", &ExpFree},
	{"xmlExpFreeCtxt", &ExpFreeCtxt},
	{"xmlExpGetLanguage", &ExpGetLanguage},
	{"xmlExpGetStart", &ExpGetStart},
	{"xmlExpIsNillable", &ExpIsNillable},
	{"xmlExpMaxToken", &ExpMaxToken},
	{"xmlExpNewAtom", &ExpNewAtom},
	{"xmlExpNewCtxt", &ExpNewCtxt},
	{"xmlExpNewOr", &ExpNewOr},
	{"xmlExpNewRange", &ExpNewRange},
	{"xmlExpNewSeq", &ExpNewSeq},
	{"xmlExpParse", &ExpParse},
	{"xmlExpRef", &ExpRef},
	{"xmlExpStringDerive", &ExpStringDerive},
	{"xmlExpSubsume", &ExpSubsume},
	{"xmlFileClose", &FileClose},
	{"xmlFileMatch", &FileMatch},
	{"xmlFileOpen", &FileOpen},
	{"xmlFileRead", &FileRead},
	{"xmlFindCharEncodingHandler", &FindCharEncodingHandler},
	{"xmlFirstElementChild", &FirstElementChild},
	{"xmlFreeAttributeTable", &FreeAttributeTable},
	{"xmlFreeAutomata", &FreeAutomata},
	{"xmlFreeCatalog", &FreeCatalog},
	{"xmlFreeDoc", &FreeDoc},
	{"xmlFreeDocElementContent", &FreeDocElementContent},
	{"xmlFreeDtd", &FreeDtd},
	{"xmlFreeElementContent", &FreeElementContent},
	{"xmlFreeElementTable", &FreeElementTable},
	{"xmlFreeEntitiesTable", &FreeEntitiesTable},
	{"xmlFreeEnumeration", &FreeEnumeration},
	{"xmlFreeIDTable", &FreeIDTable},
	{"xmlFreeInputStream", &FreeInputStream},
	{"xmlFreeMutex", &FreeMutex},
	{"xmlFreeNode", &FreeNode},
	{"xmlFreeNodeList", &FreeNodeList},
	{"xmlFreeNotationTable", &FreeNotationTable},
	{"xmlFreeNs", &FreeNs},
	{"xmlFreeNsList", &FreeNsList},
	{"xmlFreeParserCtxt", &FreeParserCtxt},
	{"xmlFreeParserInputBuffer", &FreeParserInputBuffer},
	{"xmlFreePattern", &FreePattern},
	{"xmlFreePatternList", &FreePatternList},
	{"xmlFreeProp", &FreeProp},
	{"xmlFreePropList", &FreePropList},
	{"xmlFreeRMutex", &FreeRMutex},
	{"xmlFreeRefTable", &FreeRefTable},
	{"xmlFreeStreamCtxt", &FreeStreamCtxt},
	{"xmlFreeTextReader", &FreeTextReader},
	{"xmlFreeTextWriter", &FreeTextWriter},
	{"xmlFreeURI", &FreeURI},
	{"xmlFreeValidCtxt", &FreeValidCtxt},
	{"xmlGcMemGet", &GcMemGet},
	{"xmlGcMemSetup", &GcMemSetup},
	{"xmlGetBufferAllocationScheme", &GetBufferAllocationScheme},
	{"xmlGetCharEncodingHandler", &GetCharEncodingHandler},
	{"xmlGetCharEncodingName", &GetCharEncodingName},
	{"xmlGetCompressMode", &GetCompressMode},
	{"xmlGetDocCompressMode", &GetDocCompressMode},
	{"xmlGetDocEntity", &GetDocEntity},
	{"xmlGetDtdAttrDesc", &GetDtdAttrDesc},
	{"xmlGetDtdElementDesc", &GetDtdElementDesc},
	{"xmlGetDtdEntity", &GetDtdEntity},
	{"xmlGetDtdNotationDesc", &GetDtdNotationDesc},
	{"xmlGetDtdQAttrDesc", &GetDtdQAttrDesc},
	{"xmlGetDtdQElementDesc", &GetDtdQElementDesc},
	{"xmlGetEncodingAlias", &GetEncodingAlias},
	{"xmlGetExternalEntityLoader", &GetExternalEntityLoader},
	{"xmlGetFeature", &GetFeature},
	{"xmlGetFeaturesList", &GetFeaturesList},
	{"xmlGetGlobalState", &GetGlobalState},
	{"xmlGetID", &GetID},
	{"xmlGetIntSubset", &GetIntSubset},
	{"xmlGetLastChild", &GetLastChild},
	{"xmlGetLastError", &GetLastError},
	{"xmlGetLineNo", &GetLineNo},
	{"xmlGetNoNsProp", &GetNoNsProp},
	{"xmlGetNodePath", &GetNodePath},
	{"xmlGetNsList", &GetNsList},
	{"xmlGetNsProp", &GetNsProp},
	{"xmlGetParameterEntity", &XmlGetParameterEntity},
	{"xmlGetPredefinedEntity", &GetPredefinedEntity},
	{"xmlGetProp", &GetProp},
	{"xmlGetRefs", &GetRefs},
	{"xmlGetThreadId", &GetThreadId},
	{"xmlGetUTF8Char", &GetUTF8Char},
	{"xmlHandleEntity", &HandleEntity},
	{"xmlHasFeature", &HasFeature},
	{"xmlHasNsProp", &HasNsProp},
	{"xmlHasProp", &HasProp},
	{"xmlHashAddEntry", &HashAddEntry},
	{"xmlHashAddEntry2", &HashAddEntry2},
	{"xmlHashAddEntry3", &HashAddEntry3},
	{"xmlHashCopy", &HashCopy},
	{"xmlHashCreate", &HashCreate},
	{"xmlHashCreateDict", &HashCreateDict},
	{"xmlHashFree", &HashFree},
	{"xmlHashLookup", &HashLookup},
	{"xmlHashLookup2", &HashLookup2},
	{"xmlHashLookup3", &HashLookup3},
	{"xmlHashQLookup", &HashQLookup},
	{"xmlHashQLookup2", &HashQLookup2},
	{"xmlHashQLookup3", &HashQLookup3},
	{"xmlHashRemoveEntry", &HashRemoveEntry},
	{"xmlHashRemoveEntry2", &HashRemoveEntry2},
	{"xmlHashRemoveEntry3", &HashRemoveEntry3},
	{"xmlHashScan", &HashScan},
	{"xmlHashScan3", &HashScan3},
	{"xmlHashScanFull", &HashScanFull},
	{"xmlHashScanFull3", &HashScanFull3},
	{"xmlHashSize", &HashSize},
	{"xmlHashUpdateEntry", &HashUpdateEntry},
	{"xmlHashUpdateEntry2", &HashUpdateEntry2},
	{"xmlHashUpdateEntry3", &HashUpdateEntry3},
	{"xmlIOFTPClose", &IOFTPClose},
	{"xmlIOFTPMatch", &IOFTPMatch},
	{"xmlIOFTPOpen", &IOFTPOpen},
	{"xmlIOFTPRead", &IOFTPRead},
	{"xmlIOHTTPClose", &IOHTTPClose},
	{"xmlIOHTTPMatch", &IOHTTPMatch},
	{"xmlIOHTTPOpen", &IOHTTPOpen},
	{"xmlIOHTTPOpenW", &IOHTTPOpenW},
	{"xmlIOHTTPRead", &IOHTTPRead},
	{"xmlIOParseDTD", &IOParseDTD},
	{"xmlInitCharEncodingHandlers", &InitCharEncodingHandlers},
	{"xmlInitGlobals", &InitGlobals},
	{"xmlInitMemory", &InitMemory},
	{"xmlInitNodeInfoSeq", &InitNodeInfoSeq},
	{"xmlInitParser", &InitParser},
	{"xmlInitParserCtxt", &InitParserCtxt},
	{"xmlInitThreads", &InitThreads},
	{"xmlInitializeCatalog", &InitializeCatalog},
	{"xmlInitializeGlobalState", &InitializeGlobalState},
	{"xmlInitializePredefinedEntities", &InitializePredefinedEntities},
	{"xmlIsBaseChar", &IsBaseChar},
	{"xmlIsBlank", &IsBlank},
	{"xmlIsBlankNode", &IsBlankNode},
	{"xmlIsChar", &IsChar},
	{"xmlIsCombining", &IsCombining},
	{"xmlIsDigit", &IsDigit},
	{"xmlIsExtender", &IsExtender},
	{"xmlIsID", &IsID},
	{"xmlIsIdeographic", &IsIdeographic},
	{"xmlIsLetter", &IsLetter},
	{"xmlIsMainThread", &IsMainThread},
	{"xmlIsMixedElement", &IsMixedElement},
	{"xmlIsPubidChar", &IsPubidChar},
	{"xmlIsRef", &IsRef},
	{"xmlIsXHTML", &IsXHTML},
	{"xmlKeepBlanksDefault", &KeepBlanksDefault},
	{"xmlLastElementChild", &LastElementChild},
	{"xmlLineNumbersDefault", &LineNumbersDefault},
	{"xmlLinkGetData", &LinkGetData},
	{"xmlListAppend", &ListAppend},
	{"xmlListClear", &ListClear},
	{"xmlListCopy", &ListCopy},
	{"xmlListCreate", &ListCreate},
	{"xmlListDelete", &ListDelete},
	{"xmlListDup", &ListDup},
	{"xmlListEmpty", &ListEmpty},
	{"xmlListEnd", &ListEnd},
	{"xmlListFront", &ListFront},
	{"xmlListInsert", &ListInsert},
	{"xmlListMerge", &ListMerge},
	{"xmlListPopBack", &ListPopBack},
	{"xmlListPopFront", &ListPopFront},
	{"xmlListPushBack", &ListPushBack},
	{"xmlListPushFront", &ListPushFront},
	{"xmlListRemoveAll", &ListRemoveAll},
	{"xmlListRemoveFirst", &ListRemoveFirst},
	{"xmlListRemoveLast", &ListRemoveLast},
	{"xmlListReverse", &ListReverse},
	{"xmlListReverseSearch", &ListReverseSearch},
	{"xmlListReverseWalk", &ListReverseWalk},
	{"xmlListSearch", &ListSearch},
	{"xmlListSize", &ListSize},
	{"xmlListSort", &ListSort},
	{"xmlListWalk", &ListWalk},
	{"xmlLoadACatalog", &LoadACatalog},
	{"xmlLoadCatalog", &LoadCatalog},
	{"xmlLoadCatalogs", &LoadCatalogs},
	{"xmlLoadExternalEntity", &LoadExternalEntity},
	{"xmlLoadSGMLSuperCatalog", &LoadSGMLSuperCatalog},
	{"xmlLockLibrary", &LockLibrary},
	{"xmlLsCountNode", &LsCountNode},
	{"xmlLsOneNode", &LsOneNode},
	{"xmlMallocAtomicLoc", &MallocAtomicLoc},
	{"xmlMallocLoc", &MallocLoc},
	{"xmlMemBlocks", &MemBlocks},
	{"xmlMemDisplay", &MemDisplay},
	{"xmlMemDisplayLast", &MemDisplayLast},
	{"xmlMemFree", &MemFree},
	{"xmlMemGet", &MemGet},
	{"xmlMemMalloc", &MemMalloc},
	{"xmlMemRealloc", &MemRealloc},
	{"xmlMemSetup", &MemSetup},
	{"xmlMemShow", &MemShow},
	{"xmlMemStrdupLoc", &MemStrdupLoc},
	{"xmlMemUsed", &MemUsed},
	{"xmlMemoryDump", &MemoryDump},
	{"xmlMemoryStrdup", &MemoryStrdup},
	{"xmlModuleClose", &ModuleClose},
	{"xmlModuleFree", &ModuleFree},
	{"xmlModuleOpen", &ModuleOpen},
	{"xmlModuleSymbol", &ModuleSymbol},
	{"xmlMutexLock", &MutexLock},
	{"xmlMutexUnlock", &MutexUnlock},
	{"xmlNamespaceParseNCName", &NamespaceParseNCName},
	{"xmlNamespaceParseNSDef", &NamespaceParseNSDef},
	{"xmlNamespaceParseQName", &NamespaceParseQName},
	{"xmlNanoFTPCheckResponse", &NanoFTPCheckResponse},
	{"xmlNanoFTPCleanup", &NanoFTPCleanup},
	{"xmlNanoFTPClose", &NanoFTPClose},
	{"xmlNanoFTPCloseConnection", &NanoFTPCloseConnection},
	{"xmlNanoFTPConnect", &NanoFTPConnect},
	{"xmlNanoFTPConnectTo", &NanoFTPConnectTo},
	{"xmlNanoFTPCwd", &NanoFTPCwd},
	{"xmlNanoFTPDele", &NanoFTPDele},
	{"xmlNanoFTPFreeCtxt", &NanoFTPFreeCtxt},
	{"xmlNanoFTPGet", &NanoFTPGet},
	{"xmlNanoFTPGetConnection", &NanoFTPGetConnection},
	{"xmlNanoFTPGetResponse", &NanoFTPGetResponse},
	{"xmlNanoFTPGetSocket", &NanoFTPGetSocket},
	{"xmlNanoFTPInit", &NanoFTPInit},
	{"xmlNanoFTPList", &NanoFTPList},
	{"xmlNanoFTPNewCtxt", &NanoFTPNewCtxt},
	{"xmlNanoFTPOpen", &NanoFTPOpen},
	{"xmlNanoFTPProxy", &NanoFTPProxy},
	{"xmlNanoFTPQuit", &NanoFTPQuit},
	{"xmlNanoFTPRead", &NanoFTPRead},
	{"xmlNanoFTPScanProxy", &NanoFTPScanProxy},
	{"xmlNanoFTPUpdateURL", &NanoFTPUpdateURL},
	{"xmlNanoHTTPAuthHeader", &NanoHTTPAuthHeader},
	{"xmlNanoHTTPCleanup", &NanoHTTPCleanup},
	{"xmlNanoHTTPClose", &NanoHTTPClose},
	{"xmlNanoHTTPContentLength", &NanoHTTPContentLength},
	{"xmlNanoHTTPEncoding", &NanoHTTPEncoding},
	{"xmlNanoHTTPFetch", &NanoHTTPFetch},
	{"xmlNanoHTTPInit", &NanoHTTPInit},
	{"xmlNanoHTTPMethod", &NanoHTTPMethod},
	{"xmlNanoHTTPMethodRedir", &NanoHTTPMethodRedir},
	{"xmlNanoHTTPMimeType", &NanoHTTPMimeType},
	{"xmlNanoHTTPOpen", &NanoHTTPOpen},
	{"xmlNanoHTTPOpenRedir", &NanoHTTPOpenRedir},
	{"xmlNanoHTTPRead", &NanoHTTPRead},
	{"xmlNanoHTTPRedir", &NanoHTTPRedir},
	{"xmlNanoHTTPReturnCode", &NanoHTTPReturnCode},
	{"xmlNanoHTTPSave", &NanoHTTPSave},
	{"xmlNanoHTTPScanProxy", &NanoHTTPScanProxy},
	{"xmlNewAutomata", &NewAutomata},
	{"xmlNewCDataBlock", &NewCDataBlock},
	{"xmlNewCatalog", &NewCatalog},
	{"xmlNewCharEncodingHandler", &NewCharEncodingHandler},
	{"xmlNewCharRef", &NewCharRef},
	{"xmlNewChild", &NewChild},
	{"xmlNewComment", &NewComment},
	{"xmlNewDoc", &NewDoc},
	{"xmlNewDocComment", &NewDocComment},
	{"xmlNewDocElementContent", &NewDocElementContent},
	{"xmlNewDocFragment", &NewDocFragment},
	{"xmlNewDocNode", &NewDocNode},
	{"xmlNewDocNodeEatName", &NewDocNodeEatName},
	{"xmlNewDocPI", &NewDocPI},
	{"xmlNewDocProp", &NewDocProp},
	{"xmlNewDocRawNode", &NewDocRawNode},
	{"xmlNewDocText", &NewDocText},
	{"xmlNewDocTextLen", &NewDocTextLen},
	{"xmlNewDtd", &NewDtd},
	{"xmlNewElementContent", &NewElementContent},
	{"xmlNewEntity", &NewEntity},
	{"xmlNewEntityInputStream", &NewEntityInputStream},
	{"xmlNewGlobalNs", &NewGlobalNs},
	{"xmlNewIOInputStream", &NewIOInputStream},
	{"xmlNewInputFromFile", &NewInputFromFile},
	{"xmlNewInputStream", &NewInputStream},
	{"xmlNewMutex", &NewMutex},
	{"xmlNewNode", &NewNode},
	{"xmlNewNodeEatName", &NewNodeEatName},
	{"xmlNewNs", &NewNs},
	{"xmlNewNsProp", &NewNsProp},
	{"xmlNewNsPropEatName", &NewNsPropEatName},
	{"xmlNewPI", &NewPI},
	{"xmlNewParserCtxt", &NewParserCtxt},
	{"xmlNewProp", &NewProp},
	{"xmlNewRMutex", &NewRMutex},
	{"xmlNewReference", &NewReference},
	{"xmlNewStringInputStream", &NewStringInputStream},
	{"xmlNewText", &NewText},
	{"xmlNewTextChild", &NewTextChild},
	{"xmlNewTextLen", &NewTextLen},
	{"xmlNewTextReader", &NewTextReader},
	{"xmlNewTextReaderFilename", &NewTextReaderFilename},
	{"xmlNewTextWriter", &NewTextWriter},
	{"xmlNewTextWriterDoc", &NewTextWriterDoc},
	{"xmlNewTextWriterFilename", &NewTextWriterFilename},
	{"xmlNewTextWriterMemory", &NewTextWriterMemory},
	{"xmlNewTextWriterPushParser", &NewTextWriterPushParser},
	{"xmlNewTextWriterTree", &NewTextWriterTree},
	{"xmlNewValidCtxt", &NewValidCtxt},
	{"xmlNextChar", &NextChar},
	{"xmlNextElementSibling", &NextElementSibling},
	{"xmlNoNetExternalEntityLoader", &NoNetExternalEntityLoader},
	{"xmlNodeAddContent", &NodeAddContent},
	{"xmlNodeAddContentLen", &NodeAddContentLen},
	{"xmlNodeBufGetContent", &NodeBufGetContent},
	{"xmlNodeDump", &NodeDump},
	{"xmlNodeDumpOutput", &NodeDumpOutput},
	{"xmlNodeGetBase", &NodeGetBase},
	{"xmlNodeGetContent", &NodeGetContent},
	{"xmlNodeGetLang", &NodeGetLang},
	{"xmlNodeGetSpacePreserve", &NodeGetSpacePreserve},
	{"xmlNodeIsText", &NodeIsText},
	{"xmlNodeListGetRawString", &NodeListGetRawString},
	{"xmlNodeListGetString", &NodeListGetString},
	{"xmlNodeSetBase", &NodeSetBase},
	{"xmlNodeSetContent", &NodeSetContent},
	{"xmlNodeSetContentLen", &NodeSetContentLen},
	{"xmlNodeSetLang", &NodeSetLang},
	{"xmlNodeSetName", &NodeSetName},
	{"xmlNodeSetSpacePreserve", &NodeSetSpacePreserve},
	{"xmlNormalizeURIPath", &NormalizeURIPath},
	{"xmlNormalizeWindowsPath", &NormalizeWindowsPath},
	{"xmlOutputBufferClose", &OutputBufferClose},
	{"xmlOutputBufferCreateBuffer", &OutputBufferCreateBuffer},
	{"xmlOutputBufferCreateFd", &OutputBufferCreateFd},
	{"xmlOutputBufferCreateFile", &OutputBufferCreateFile},
	{"xmlOutputBufferCreateFilename", &OutputBufferCreateFilename},
	{"xmlOutputBufferCreateFilenameDefault", &OutputBufferCreateFilenameDefault},
	{"xmlOutputBufferCreateIO", &OutputBufferCreateIO},
	{"xmlOutputBufferFlush", &OutputBufferFlush},
	{"xmlOutputBufferWrite", &OutputBufferWrite},
	{"xmlOutputBufferWriteEscape", &OutputBufferWriteEscape},
	{"xmlOutputBufferWriteString", &OutputBufferWriteString},
	{"xmlParseAttValue", &ParseAttValue},
	{"xmlParseAttribute", &ParseAttribute},
	{"xmlParseAttributeListDecl", &ParseAttributeListDecl},
	{"xmlParseAttributeType", &ParseAttributeType},
	{"xmlParseBalancedChunkMemory", &ParseBalancedChunkMemory},
	{"xmlParseBalancedChunkMemoryRecover", &ParseBalancedChunkMemoryRecover},
	{"xmlParseCDSect", &ParseCDSect},
	{"xmlParseCatalogFile", &ParseCatalogFile},
	{"xmlParseCharData", &ParseCharData},
	{"xmlParseCharEncoding", &ParseCharEncoding},
	{"xmlParseCharRef", &ParseCharRef},
	{"xmlParseChunk", &ParseChunk},
	{"xmlParseComment", &ParseComment},
	{"xmlParseContent", &ParseContent},
	{"xmlParseCtxtExternalEntity", &ParseCtxtExternalEntity},
	{"xmlParseDTD", &ParseDTD},
	{"xmlParseDefaultDecl", &ParseDefaultDecl},
	{"xmlParseDoc", &ParseDoc},
	{"xmlParseDocTypeDecl", &ParseDocTypeDecl},
	{"xmlParseDocument", &ParseDocument},
	{"xmlParseElement", &ParseElement},
	{"xmlParseElementChildrenContentDecl", &ParseElementChildrenContentDecl},
	{"xmlParseElementContentDecl", &ParseElementContentDecl},
	{"xmlParseElementDecl", &ParseElementDecl},
	{"xmlParseElementMixedContentDecl", &ParseElementMixedContentDecl},
	{"xmlParseEncName", &ParseEncName},
	{"xmlParseEncodingDecl", &ParseEncodingDecl},
	{"xmlParseEndTag", &ParseEndTag},
	{"xmlParseEntity", &ParseEntity},
	{"xmlParseEntityDecl", &ParseEntityDecl},
	{"xmlParseEntityRef", &ParseEntityRef},
	{"xmlParseEntityValue", &ParseEntityValue},
	{"xmlParseEnumeratedType", &ParseEnumeratedType},
	{"xmlParseEnumerationType", &ParseEnumerationType},
	{"xmlParseExtParsedEnt", &ParseExtParsedEnt},
	{"xmlParseExternalEntity", &ParseExternalEntity},
	{"xmlParseExternalID", &ParseExternalID},
	{"xmlParseExternalSubset", &ParseExternalSubset},
	{"xmlParseFile", &ParseFile},
	{"xmlParseInNodeContext", &ParseInNodeContext},
	{"xmlParseMarkupDecl", &ParseMarkupDecl},
	{"xmlParseMemory", &ParseMemory},
	{"xmlParseMisc", &ParseMisc},
	{"xmlParseName", &ParseName},
	{"xmlParseNamespace", &ParseNamespace},
	{"xmlParseNmtoken", &ParseNmtoken},
	{"xmlParseNotationDecl", &ParseNotationDecl},
	{"xmlParseNotationType", &ParseNotationType},
	{"xmlParsePEReference", &ParsePEReference},
	{"xmlParsePI", &ParsePI},
	{"xmlParsePITarget", &ParsePITarget},
	{"xmlParsePubidLiteral", &ParsePubidLiteral},
	{"xmlParseQuotedString", &ParseQuotedString},
	{"xmlParseReference", &ParseReference},
	{"xmlParseSDDecl", &ParseSDDecl},
	{"xmlParseStartTag", &ParseStartTag},
	{"xmlParseSystemLiteral", &ParseSystemLiteral},
	{"xmlParseTextDecl", &ParseTextDecl},
	{"xmlParseURI", &ParseURI},
	{"xmlParseURIRaw", &ParseURIRaw},
	{"xmlParseURIReference", &ParseURIReference},
	{"xmlParseVersionInfo", &ParseVersionInfo},
	{"xmlParseVersionNum", &ParseVersionNum},
	{"xmlParseXMLDecl", &ParseXMLDecl},
	{"xmlParserAddNodeInfo", &ParserAddNodeInfo},
	{"xmlParserError", &ParserError},
	{"xmlParserFindNodeInfo", &ParserFindNodeInfo},
	{"xmlParserFindNodeInfoIndex", &ParserFindNodeInfoIndex},
	{"xmlParserGetDirectory", &ParserGetDirectory},
	{"xmlParserHandlePEReference", &ParserHandlePEReference},
	{"xmlParserHandleReference", &ParserHandleReference},
	{"xmlParserInputBufferCreateFd", &ParserInputBufferCreateFd},
	{"xmlParserInputBufferCreateFile", &ParserInputBufferCreateFile},
	{"xmlParserInputBufferCreateFilename", &ParserInputBufferCreateFilename},
	{"xmlParserInputBufferCreateFilenameDefault", &ParserInputBufferCreateFilenameDefault},
	{"xmlParserInputBufferCreateIO", &ParserInputBufferCreateIO},
	{"xmlParserInputBufferCreateMem", &ParserInputBufferCreateMem},
	{"xmlParserInputBufferCreateStatic", &ParserInputBufferCreateStatic},
	{"xmlParserInputBufferGrow", &ParserInputBufferGrow},
	{"xmlParserInputBufferPush", &ParserInputBufferPush},
	{"xmlParserInputBufferRead", &ParserInputBufferRead},
	{"xmlParserInputGrow", &ParserInputGrow},
	{"xmlParserInputRead", &ParserInputRead},
	{"xmlParserInputShrink", &ParserInputShrink},
	{"xmlParserPrintFileContext", &ParserPrintFileContext},
	{"xmlParserPrintFileInfo", &ParserPrintFileInfo},
	{"xmlParserValidityError", &ParserValidityError},
	{"xmlParserValidityWarning", &ParserValidityWarning},
	{"xmlParserWarning", &ParserWarning},
	{"xmlPathToURI", &PathToURI},
	{"xmlPatternFromRoot", &PatternFromRoot},
	{"xmlPatternGetStreamCtxt", &PatternGetStreamCtxt},
	{"xmlPatternMatch", &PatternMatch},
	{"xmlPatternMaxDepth", &PatternMaxDepth},
	{"xmlPatternMinDepth", &PatternMinDepth},
	{"xmlPatternStreamable", &PatternStreamable},
	{"xmlPatterncompile", &Patterncompile},
	{"xmlPedanticParserDefault", &PedanticParserDefault},
	{"xmlPopInput", &PopInput},
	{"xmlPopInputCallbacks", &PopInputCallbacks},
	{"xmlPreviousElementSibling", &PreviousElementSibling},
	{"xmlPrintURI", &PrintURI},
	{"xmlPushInput", &PushInput},
	{"xmlRMutexLock", &RMutexLock},
	{"xmlRMutexUnlock", &RMutexUnlock},
	{"xmlReadDoc", &ReadDoc},
	{"xmlReadFd", &ReadFd},
	{"xmlReadFile", &ReadFile},
	{"xmlReadIO", &ReadIO},
	{"xmlReadMemory", &ReadMemory},
	{"xmlReaderForDoc", &ReaderForDoc},
	{"xmlReaderForFd", &ReaderForFd},
	{"xmlReaderForFile", &ReaderForFile},
	{"xmlReaderForIO", &ReaderForIO},
	{"xmlReaderForMemory", &ReaderForMemory},
	{"xmlReaderNewDoc", &ReaderNewDoc},
	{"xmlReaderNewFd", &ReaderNewFd},
	{"xmlReaderNewFile", &ReaderNewFile},
	{"xmlReaderNewIO", &ReaderNewIO},
	{"xmlReaderNewMemory", &ReaderNewMemory},
	{"xmlReaderNewWalker", &ReaderNewWalker},
	{"xmlReaderWalker", &ReaderWalker},
	{"xmlReallocLoc", &ReallocLoc},
	{"xmlReconciliateNs", &ReconciliateNs},
	{"xmlRecoverDoc", &RecoverDoc},
	{"xmlRecoverFile", &RecoverFile},
	{"xmlRecoverMemory", &RecoverMemory},
	{"xmlRegExecErrInfo", &RegExecErrInfo},
	{"xmlRegExecNextValues", &RegExecNextValues},
	{"xmlRegExecPushString", &RegExecPushString},
	{"xmlRegExecPushString2", &RegExecPushString2},
	{"xmlRegFreeExecCtxt", &RegFreeExecCtxt},
	{"xmlRegFreeRegexp", &RegFreeRegexp},
	{"xmlRegNewExecCtxt", &RegNewExecCtxt},
	{"xmlRegexpCompile", &RegexpCompile},
	{"xmlRegexpExec", &RegexpExec},
	{"xmlRegexpIsDeterminist", &RegexpIsDeterminist},
	{"xmlRegexpPrint", &RegexpPrint},
	{"xmlRegisterCharEncodingHandler", &RegisterCharEncodingHandler},
	{"xmlRegisterDefaultInputCallbacks", &RegisterDefaultInputCallbacks},
	{"xmlRegisterDefaultOutputCallbacks", &RegisterDefaultOutputCallbacks},
	{"xmlRegisterHTTPPostCallbacks", &RegisterHTTPPostCallbacks},
	{"xmlRegisterInputCallbacks", &RegisterInputCallbacks},
	{"xmlRegisterNodeDefault", &RegisterNodeDefault},
	{"xmlRegisterOutputCallbacks", &RegisterOutputCallbacks},
	{"xmlRelaxNGCleanupTypes", &RelaxNGCleanupTypes},
	{"xmlRelaxNGDump", &RelaxNGDump},
	{"xmlRelaxNGDumpTree", &RelaxNGDumpTree},
	{"xmlRelaxNGFree", &RelaxNGFree},
	{"xmlRelaxNGFreeParserCtxt", &RelaxNGFreeParserCtxt},
	{"xmlRelaxNGFreeValidCtxt", &RelaxNGFreeValidCtxt},
	{"xmlRelaxNGGetParserErrors", &RelaxNGGetParserErrors},
	{"xmlRelaxNGGetValidErrors", &RelaxNGGetValidErrors},
	{"xmlRelaxNGInitTypes", &RelaxNGInitTypes},
	{"xmlRelaxNGNewDocParserCtxt", &RelaxNGNewDocParserCtxt},
	{"xmlRelaxNGNewMemParserCtxt", &RelaxNGNewMemParserCtxt},
	{"xmlRelaxNGNewParserCtxt", &RelaxNGNewParserCtxt},
	{"xmlRelaxNGNewValidCtxt", &RelaxNGNewValidCtxt},
	{"xmlRelaxNGParse", &RelaxNGParse},
	{"xmlRelaxNGSetParserErrors", &RelaxNGSetParserErrors},
	{"xmlRelaxNGSetParserStructuredErrors", &RelaxNGSetParserStructuredErrors},
	{"xmlRelaxNGSetValidErrors", &RelaxNGSetValidErrors},
	{"xmlRelaxNGSetValidStructuredErrors", &RelaxNGSetValidStructuredErrors},
	{"xmlRelaxNGValidateDoc", &RelaxNGValidateDoc},
	{"xmlRelaxNGValidateFullElement", &RelaxNGValidateFullElement},
	{"xmlRelaxNGValidatePopElement", &RelaxNGValidatePopElement},
	{"xmlRelaxNGValidatePushCData", &RelaxNGValidatePushCData},
	{"xmlRelaxNGValidatePushElement", &RelaxNGValidatePushElement},
	{"xmlRelaxParserSetFlag", &RelaxParserSetFlag},
	{"xmlRemoveID", &RemoveID},
	{"xmlRemoveProp", &RemoveProp},
	{"xmlRemoveRef", &RemoveRef},
	{"xmlReplaceNode", &ReplaceNode},
	{"xmlResetError", &ResetError},
	{"xmlResetLastError", &ResetLastError},
	{"xmlSAX2AttributeDecl", &SAX2AttributeDecl},
	{"xmlSAX2CDataBlock", &SAX2CDataBlock},
	{"xmlSAX2Characters", &SAX2Characters},
	{"xmlSAX2Comment", &SAX2Comment},
	{"xmlSAX2ElementDecl", &SAX2ElementDecl},
	{"xmlSAX2EndDocument", &SAX2EndDocument},
	{"xmlSAX2EndElement", &SAX2EndElement},
	{"xmlSAX2EndElementNs", &SAX2EndElementNs},
	{"xmlSAX2EntityDecl", &SAX2EntityDecl},
	{"xmlSAX2ExternalSubset", &SAX2ExternalSubset},
	{"xmlSAX2GetColumnNumber", &SAX2GetColumnNumber},
	{"xmlSAX2GetEntity", &SAX2GetEntity},
	{"xmlSAX2GetLineNumber", &SAX2GetLineNumber},
	{"xmlSAX2GetParameterEntity", &SAX2GetParameterEntity},
	{"xmlSAX2GetPublicId", &SAX2GetPublicId},
	{"xmlSAX2GetSystemId", &SAX2GetSystemId},
	{"xmlSAX2HasExternalSubset", &SAX2HasExternalSubset},
	{"xmlSAX2HasInternalSubset", &SAX2HasInternalSubset},
	{"xmlSAX2IgnorableWhitespace", &SAX2IgnorableWhitespace},
	{"xmlSAX2InitDefaultSAXHandler", &SAX2InitDefaultSAXHandler},
	{"xmlSAX2InitDocbDefaultSAXHandler", &SAX2InitDocbDefaultSAXHandler},
	{"xmlSAX2InitHtmlDefaultSAXHandler", &SAX2InitHtmlDefaultSAXHandler},
	{"xmlSAX2InternalSubset", &SAX2InternalSubset},
	{"xmlSAX2IsStandalone", &SAX2IsStandalone},
	{"xmlSAX2NotationDecl", &SAX2NotationDecl},
	{"xmlSAX2ProcessingInstruction", &SAX2ProcessingInstruction},
	{"xmlSAX2Reference", &SAX2Reference},
	{"xmlSAX2ResolveEntity", &SAX2ResolveEntity},
	{"xmlSAX2SetDocumentLocator", &SAX2SetDocumentLocator},
	{"xmlSAX2StartDocument", &SAX2StartDocument},
	{"xmlSAX2StartElement", &SAX2StartElement},
	{"xmlSAX2StartElementNs", &SAX2StartElementNs},
	{"xmlSAX2UnparsedEntityDecl", &SAX2UnparsedEntityDecl},
	{"xmlSAXDefaultVersion", &SAXDefaultVersion},
	{"xmlSAXParseDTD", &SAXParseDTD},
	{"xmlSAXParseDoc", &SAXParseDoc},
	{"xmlSAXParseEntity", &SAXParseEntity},
	{"xmlSAXParseFile", &SAXParseFile},
	{"xmlSAXParseFileWithData", &SAXParseFileWithData},
	{"xmlSAXParseMemory", &SAXParseMemory},
	{"xmlSAXParseMemoryWithData", &SAXParseMemoryWithData},
	{"xmlSAXUserParseFile", &SAXUserParseFile},
	{"xmlSAXUserParseMemory", &SAXUserParseMemory},
	{"xmlSAXVersion", &SAXVersion},
	{"xmlSaveClose", &SaveClose},
	{"xmlSaveDoc", &SaveDoc},
	{"xmlSaveFile", &SaveFile},
	{"xmlSaveFileEnc", &SaveFileEnc},
	{"xmlSaveFileTo", &SaveFileTo},
	{"xmlSaveFlush", &SaveFlush},
	{"xmlSaveFormatFile", &SaveFormatFile},
	{"xmlSaveFormatFileEnc", &SaveFormatFileEnc},
	{"xmlSaveFormatFileTo", &SaveFormatFileTo},
	{"xmlSaveSetAttrEscape", &SaveSetAttrEscape},
	{"xmlSaveSetEscape", &SaveSetEscape},
	{"xmlSaveToBuffer", &SaveToBuffer},
	{"xmlSaveToFd", &SaveToFd},
	{"xmlSaveToFilename", &SaveToFilename},
	{"xmlSaveToIO", &SaveToIO},
	{"xmlSaveTree", &SaveTree},
	{"xmlSaveUri", &SaveUri},
	{"xmlScanName", &ScanName},
	{"xmlSchemaCheckFacet", &SchemaCheckFacet},
	{"xmlSchemaCleanupTypes", &SchemaCleanupTypes},
	{"xmlSchemaCollapseString", &SchemaCollapseString},
	{"xmlSchemaCompareValues", &SchemaCompareValues},
	{"xmlSchemaCompareValuesWhtsp", &SchemaCompareValuesWhtsp},
	{"xmlSchemaCopyValue", &SchemaCopyValue},
	{"xmlSchemaDump", &SchemaDump},
	{"xmlSchemaFree", &SchemaFree},
	{"xmlSchemaFreeFacet", &SchemaFreeFacet},
	{"xmlSchemaFreeParserCtxt", &SchemaFreeParserCtxt},
	{"xmlSchemaFreeType", &SchemaFreeType},
	{"xmlSchemaFreeValidCtxt", &SchemaFreeValidCtxt},
	{"xmlSchemaFreeValue", &SchemaFreeValue},
	{"xmlSchemaFreeWildcard", &SchemaFreeWildcard},
	{"xmlSchemaGetBuiltInListSimpleTypeItemType", &SchemaGetBuiltInListSimpleTypeItemType},
	{"xmlSchemaGetBuiltInType", &SchemaGetBuiltInType},
	{"xmlSchemaGetCanonValue", &SchemaGetCanonValue},
	{"xmlSchemaGetCanonValueWhtsp", &SchemaGetCanonValueWhtsp},
	{"xmlSchemaGetFacetValueAsULong", &SchemaGetFacetValueAsULong},
	{"xmlSchemaGetParserErrors", &SchemaGetParserErrors},
	{"xmlSchemaGetPredefinedType", &SchemaGetPredefinedType},
	{"xmlSchemaGetValType", &SchemaGetValType},
	{"xmlSchemaGetValidErrors", &SchemaGetValidErrors},
	{"xmlSchemaInitTypes", &SchemaInitTypes},
	{"xmlSchemaIsBuiltInTypeFacet", &SchemaIsBuiltInTypeFacet},
	{"xmlSchemaIsValid", &SchemaIsValid},
	{"xmlSchemaNewDocParserCtxt", &SchemaNewDocParserCtxt},
	{"xmlSchemaNewFacet", &SchemaNewFacet},
	{"xmlSchemaNewMemParserCtxt", &SchemaNewMemParserCtxt},
	{"xmlSchemaNewNOTATIONValue", &SchemaNewNOTATIONValue},
	{"xmlSchemaNewParserCtxt", &SchemaNewParserCtxt},
	{"xmlSchemaNewQNameValue", &SchemaNewQNameValue},
	{"xmlSchemaNewStringValue", &SchemaNewStringValue},
	{"xmlSchemaNewValidCtxt", &SchemaNewValidCtxt},
	{"xmlSchemaParse", &SchemaParse},
	{"xmlSchemaSAXPlug", &SchemaSAXPlug},
	{"xmlSchemaSAXUnplug", &SchemaSAXUnplug},
	{"xmlSchemaSetParserErrors", &SchemaSetParserErrors},
	{"xmlSchemaSetParserStructuredErrors", &SchemaSetParserStructuredErrors},
	{"xmlSchemaSetValidErrors", &SchemaSetValidErrors},
	{"xmlSchemaSetValidOptions", &SchemaSetValidOptions},
	{"xmlSchemaSetValidStructuredErrors", &SchemaSetValidStructuredErrors},
	{"xmlSchemaValPredefTypeNode", &SchemaValPredefTypeNode},
	{"xmlSchemaValPredefTypeNodeNoNorm", &SchemaValPredefTypeNodeNoNorm},
	{"xmlSchemaValidCtxtGetOptions", &SchemaValidCtxtGetOptions},
	{"xmlSchemaValidCtxtGetParserCtxt", &SchemaValidCtxtGetParserCtxt},
	{"xmlSchemaValidateDoc", &SchemaValidateDoc},
	{"xmlSchemaValidateFacet", &SchemaValidateFacet},
	{"xmlSchemaValidateFacetWhtsp", &SchemaValidateFacetWhtsp},
	{"xmlSchemaValidateFile", &SchemaValidateFile},
	{"xmlSchemaValidateLengthFacet", &SchemaValidateLengthFacet},
	{"xmlSchemaValidateLengthFacetWhtsp", &SchemaValidateLengthFacetWhtsp},
	{"xmlSchemaValidateListSimpleTypeFacet", &SchemaValidateListSimpleTypeFacet},
	{"xmlSchemaValidateOneElement", &SchemaValidateOneElement},
	{"xmlSchemaValidatePredefinedType", &SchemaValidatePredefinedType},
	{"xmlSchemaValidateStream", &SchemaValidateStream},
	{"xmlSchemaValueAppend", &SchemaValueAppend},
	{"xmlSchemaValueGetAsBoolean", &SchemaValueGetAsBoolean},
	{"xmlSchemaValueGetAsString", &SchemaValueGetAsString},
	{"xmlSchemaValueGetNext", &SchemaValueGetNext},
	{"xmlSchemaWhiteSpaceReplace", &SchemaWhiteSpaceReplace},
	{"xmlSchematronFree", &SchematronFree},
	{"xmlSchematronFreeParserCtxt", &SchematronFreeParserCtxt},
	{"xmlSchematronFreeValidCtxt", &SchematronFreeValidCtxt},
	{"xmlSchematronNewDocParserCtxt", &SchematronNewDocParserCtxt},
	{"xmlSchematronNewMemParserCtxt", &SchematronNewMemParserCtxt},
	{"xmlSchematronNewParserCtxt", &SchematronNewParserCtxt},
	{"xmlSchematronNewValidCtxt", &SchematronNewValidCtxt},
	{"xmlSchematronParse", &SchematronParse},
	{"xmlSchematronSetValidStructuredErrors", &SchematronSetValidStructuredErrors},
	{"xmlSchematronValidateDoc", &SchematronValidateDoc},
	{"xmlSearchNs", &SearchNs},
	{"xmlSearchNsByHref", &SearchNsByHref},
	{"xmlSetBufferAllocationScheme", &SetBufferAllocationScheme},
	{"xmlSetCompressMode", &SetCompressMode},
	{"xmlSetDocCompressMode", &SetDocCompressMode},
	{"xmlSetEntityReferenceFunc", &SetEntityReferenceFunc},
	{"xmlSetExternalEntityLoader", &SetExternalEntityLoader},
	{"xmlSetFeature", &SetFeature},
	{"xmlSetGenericErrorFunc", &SetGenericErrorFunc},
	{"xmlSetListDoc", &SetListDoc},
	{"xmlSetNs", &SetNs},
	{"xmlSetNsProp", &SetNsProp},
	{"xmlSetProp", &SetProp},
	{"xmlSetStructuredErrorFunc", &SetStructuredErrorFunc},
	{"xmlSetTreeDoc", &SetTreeDoc},
	{"xmlSetupParserForBuffer", &SetupParserForBuffer},
	{"xmlShell", &Shell},
	{"xmlShellBase", &ShellBase},
	{"xmlShellCat", &ShellCat},
	{"xmlShellDir", &ShellDir},
	{"xmlShellDu", &ShellDu},
	{"xmlShellList", &ShellList},
	{"xmlShellLoad", &ShellLoad},
	{"xmlShellPrintNode", &ShellPrintNode},
	{"xmlShellPrintXPathError", &ShellPrintXPathError},
	{"xmlShellPrintXPathResult", &ShellPrintXPathResult},
	{"xmlShellPwd", &ShellPwd},
	{"xmlShellSave", &ShellSave},
	{"xmlShellValidate", &ShellValidate},
	{"xmlShellWrite", &ShellWrite},
	{"xmlSkipBlankChars", &SkipBlankChars},
	{"xmlSnprintfElementContent", &SnprintfElementContent},
	{"xmlSplitQName", &SplitQName},
	{"xmlSplitQName2", &SplitQName2},
	{"xmlSplitQName3", &SplitQName3},
	{"xmlSprintfElementContent", &SprintfElementContent},
	{"xmlStopParser", &StopParser},
	{"xmlStrEqual", &StrEqual},
	{"xmlStrPrintf", &StrPrintf},
	{"xmlStrQEqual", &StrQEqual},
	{"xmlStrVPrintf", &StrVPrintf},
	{"xmlStrcasecmp", &Strcasecmp},
	{"xmlStrcasestr", &Strcasestr},
	{"xmlStrcat", &Strcat},
	{"xmlStrchr", &Strchr},
	{"xmlStrcmp", &Strcmp},
	{"xmlStrdup", &Strdup},
	{"xmlStreamPop", &StreamPop},
	{"xmlStreamPush", &StreamPush},
	{"xmlStreamPushAttr", &StreamPushAttr},
	{"xmlStreamPushNode", &StreamPushNode},
	{"xmlStreamWantsAnyNode", &StreamWantsAnyNode},
	{"xmlStringCurrentChar", &StringCurrentChar},
	{"xmlStringDecodeEntities", &StringDecodeEntities},
	{"xmlStringGetNodeList", &StringGetNodeList},
	{"xmlStringLenDecodeEntities", &StringLenDecodeEntities},
	{"xmlStringLenGetNodeList", &StringLenGetNodeList},
	{"xmlStrlen", &Strlen},
	{"xmlStrncasecmp", &Strncasecmp},
	{"xmlStrncat", &Strncat},
	{"xmlStrncatNew", &StrncatNew},
	{"xmlStrncmp", &Strncmp},
	{"xmlStrndup", &Strndup},
	{"xmlStrstr", &Strstr},
	{"xmlStrsub", &Strsub},
	{"xmlSubstituteEntitiesDefault", &SubstituteEntitiesDefault},
	{"xmlSwitchEncoding", &SwitchEncoding},
	{"xmlSwitchInputEncoding", &SwitchInputEncoding},
	{"xmlSwitchToEncoding", &SwitchToEncoding},
	{"xmlTextConcat", &TextConcat},
	{"xmlTextMerge", &TextMerge},
	{"xmlTextReaderAttributeCount", &TextReaderAttributeCount},
	{"xmlTextReaderBaseUri", &TextReaderBaseUri},
	{"xmlTextReaderByteConsumed", &TextReaderByteConsumed},
	{"xmlTextReaderClose", &TextReaderClose},
	{"xmlTextReaderConstBaseUri", &TextReaderConstBaseUri},
	{"xmlTextReaderConstEncoding", &TextReaderConstEncoding},
	{"xmlTextReaderConstLocalName", &TextReaderConstLocalName},
	{"xmlTextReaderConstName", &TextReaderConstName},
	{"xmlTextReaderConstNamespaceUri", &TextReaderConstNamespaceUri},
	{"xmlTextReaderConstPrefix", &TextReaderConstPrefix},
	{"xmlTextReaderConstString", &TextReaderConstString},
	{"xmlTextReaderConstValue", &TextReaderConstValue},
	{"xmlTextReaderConstXmlLang", &TextReaderConstXmlLang},
	{"xmlTextReaderConstXmlVersion", &TextReaderConstXmlVersion},
	{"xmlTextReaderCurrentDoc", &TextReaderCurrentDoc},
	{"xmlTextReaderCurrentNode", &TextReaderCurrentNode},
	{"xmlTextReaderDepth", &TextReaderDepth},
	{"xmlTextReaderExpand", &TextReaderExpand},
	{"xmlTextReaderGetAttribute", &TextReaderGetAttribute},
	{"xmlTextReaderGetAttributeNo", &TextReaderGetAttributeNo},
	{"xmlTextReaderGetAttributeNs", &TextReaderGetAttributeNs},
	{"xmlTextReaderGetErrorHandler", &TextReaderGetErrorHandler},
	{"xmlTextReaderGetParserColumnNumber", &TextReaderGetParserColumnNumber},
	{"xmlTextReaderGetParserLineNumber", &TextReaderGetParserLineNumber},
	{"xmlTextReaderGetParserProp", &TextReaderGetParserProp},
	{"xmlTextReaderGetRemainder", &TextReaderGetRemainder},
	{"xmlTextReaderHasAttributes", &TextReaderHasAttributes},
	{"xmlTextReaderHasValue", &TextReaderHasValue},
	{"xmlTextReaderIsDefault", &TextReaderIsDefault},
	{"xmlTextReaderIsEmptyElement", &TextReaderIsEmptyElement},
	{"xmlTextReaderIsNamespaceDecl", &TextReaderIsNamespaceDecl},
	{"xmlTextReaderIsValid", &TextReaderIsValid},
	{"xmlTextReaderLocalName", &TextReaderLocalName},
	{"xmlTextReaderLocatorBaseURI", &TextReaderLocatorBaseURI},
	{"xmlTextReaderLocatorLineNumber", &TextReaderLocatorLineNumber},
	{"xmlTextReaderLookupNamespace", &TextReaderLookupNamespace},
	{"xmlTextReaderMoveToAttribute", &TextReaderMoveToAttribute},
	{"xmlTextReaderMoveToAttributeNo", &TextReaderMoveToAttributeNo},
	{"xmlTextReaderMoveToAttributeNs", &TextReaderMoveToAttributeNs},
	{"xmlTextReaderMoveToElement", &TextReaderMoveToElement},
	{"xmlTextReaderMoveToFirstAttribute", &TextReaderMoveToFirstAttribute},
	{"xmlTextReaderMoveToNextAttribute", &TextReaderMoveToNextAttribute},
	{"xmlTextReaderName", &TextReaderName},
	{"xmlTextReaderNamespaceUri", &TextReaderNamespaceUri},
	{"xmlTextReaderNext", &TextReaderNext},
	{"xmlTextReaderNextSibling", &TextReaderNextSibling},
	{"xmlTextReaderNodeType", &TextReaderNodeType},
	{"xmlTextReaderNormalization", &TextReaderNormalization},
	{"xmlTextReaderPrefix", &TextReaderPrefix},
	{"xmlTextReaderPreserve", &TextReaderPreserve},
	{"xmlTextReaderPreservePattern", &TextReaderPreservePattern},
	{"xmlTextReaderQuoteChar", &TextReaderQuoteChar},
	{"xmlTextReaderRead", &TextReaderRead},
	{"xmlTextReaderReadAttributeValue", &TextReaderReadAttributeValue},
	{"xmlTextReaderReadInnerXml", &TextReaderReadInnerXml},
	{"xmlTextReaderReadOuterXml", &TextReaderReadOuterXml},
	{"xmlTextReaderReadState", &TextReaderReadState},
	{"xmlTextReaderReadString", &TextReaderReadString},
	{"xmlTextReaderRelaxNGSetSchema", &TextReaderRelaxNGSetSchema},
	{"xmlTextReaderRelaxNGValidate", &TextReaderRelaxNGValidate},
	{"xmlTextReaderSchemaValidate", &TextReaderSchemaValidate},
	{"xmlTextReaderSchemaValidateCtxt", &TextReaderSchemaValidateCtxt},
	{"xmlTextReaderSetErrorHandler", &TextReaderSetErrorHandler},
	{"xmlTextReaderSetParserProp", &TextReaderSetParserProp},
	{"xmlTextReaderSetSchema", &TextReaderSetSchema},
	{"xmlTextReaderSetStructuredErrorHandler", &TextReaderSetStructuredErrorHandler},
	{"xmlTextReaderSetup", &TextReaderSetup},
	{"xmlTextReaderStandalone", &TextReaderStandalone},
	{"xmlTextReaderValue", &TextReaderValue},
	{"xmlTextReaderXmlLang", &TextReaderXmlLang},
	{"xmlTextWriterEndAttribute", &TextWriterEndAttribute},
	{"xmlTextWriterEndCDATA", &TextWriterEndCDATA},
	{"xmlTextWriterEndComment", &TextWriterEndComment},
	{"xmlTextWriterEndDTD", &TextWriterEndDTD},
	{"xmlTextWriterEndDTDAttlist", &TextWriterEndDTDAttlist},
	{"xmlTextWriterEndDTDElement", &TextWriterEndDTDElement},
	{"xmlTextWriterEndDTDEntity", &TextWriterEndDTDEntity},
	{"xmlTextWriterEndDocument", &TextWriterEndDocument},
	{"xmlTextWriterEndElement", &TextWriterEndElement},
	{"xmlTextWriterEndPI", &TextWriterEndPI},
	{"xmlTextWriterFlush", &TextWriterFlush},
	{"xmlTextWriterFullEndElement", &TextWriterFullEndElement},
	{"xmlTextWriterSetIndent", &TextWriterSetIndent},
	{"xmlTextWriterSetIndentString", &TextWriterSetIndentString},
	{"xmlTextWriterStartAttribute", &TextWriterStartAttribute},
	{"xmlTextWriterStartAttributeNS", &TextWriterStartAttributeNS},
	{"xmlTextWriterStartCDATA", &TextWriterStartCDATA},
	{"xmlTextWriterStartComment", &TextWriterStartComment},
	{"xmlTextWriterStartDTD", &TextWriterStartDTD},
	{"xmlTextWriterStartDTDAttlist", &TextWriterStartDTDAttlist},
	{"xmlTextWriterStartDTDElement", &TextWriterStartDTDElement},
	{"xmlTextWriterStartDTDEntity", &TextWriterStartDTDEntity},
	{"xmlTextWriterStartDocument", &TextWriterStartDocument},
	{"xmlTextWriterStartElement", &TextWriterStartElement},
	{"xmlTextWriterStartElementNS", &TextWriterStartElementNS},
	{"xmlTextWriterStartPI", &TextWriterStartPI},
	{"xmlTextWriterWriteAttribute", &TextWriterWriteAttribute},
	{"xmlTextWriterWriteAttributeNS", &TextWriterWriteAttributeNS},
	{"xmlTextWriterWriteBase64", &TextWriterWriteBase64},
	{"xmlTextWriterWriteBinHex", &TextWriterWriteBinHex},
	{"xmlTextWriterWriteCDATA", &TextWriterWriteCDATA},
	{"xmlTextWriterWriteComment", &TextWriterWriteComment},
	{"xmlTextWriterWriteDTD", &TextWriterWriteDTD},
	{"xmlTextWriterWriteDTDAttlist", &TextWriterWriteDTDAttlist},
	{"xmlTextWriterWriteDTDElement", &TextWriterWriteDTDElement},
	{"xmlTextWriterWriteDTDEntity", &TextWriterWriteDTDEntity},
	{"xmlTextWriterWriteDTDExternalEntity", &TextWriterWriteDTDExternalEntity},
	{"xmlTextWriterWriteDTDExternalEntityContents", &TextWriterWriteDTDExternalEntityContents},
	{"xmlTextWriterWriteDTDInternalEntity", &TextWriterWriteDTDInternalEntity},
	{"xmlTextWriterWriteDTDNotation", &TextWriterWriteDTDNotation},
	{"xmlTextWriterWriteElement", &TextWriterWriteElement},
	{"xmlTextWriterWriteElementNS", &TextWriterWriteElementNS},
	{"xmlTextWriterWriteFormatAttribute", &TextWriterWriteFormatAttribute},
	{"xmlTextWriterWriteFormatAttributeNS", &TextWriterWriteFormatAttributeNS},
	{"xmlTextWriterWriteFormatCDATA", &TextWriterWriteFormatCDATA},
	{"xmlTextWriterWriteFormatComment", &TextWriterWriteFormatComment},
	{"xmlTextWriterWriteFormatDTD", &TextWriterWriteFormatDTD},
	{"xmlTextWriterWriteFormatDTDAttlist", &TextWriterWriteFormatDTDAttlist},
	{"xmlTextWriterWriteFormatDTDElement", &TextWriterWriteFormatDTDElement},
	{"xmlTextWriterWriteFormatDTDInternalEntity", &TextWriterWriteFormatDTDInternalEntity},
	{"xmlTextWriterWriteFormatElement", &TextWriterWriteFormatElement},
	{"xmlTextWriterWriteFormatElementNS", &TextWriterWriteFormatElementNS},
	{"xmlTextWriterWriteFormatPI", &TextWriterWriteFormatPI},
	{"xmlTextWriterWriteFormatRaw", &TextWriterWriteFormatRaw},
	{"xmlTextWriterWriteFormatString", &TextWriterWriteFormatString},
	{"xmlTextWriterWritePI", &TextWriterWritePI},
	{"xmlTextWriterWriteRaw", &TextWriterWriteRaw},
	{"xmlTextWriterWriteRawLen", &TextWriterWriteRawLen},
	{"xmlTextWriterWriteString", &TextWriterWriteString},
	{"xmlTextWriterWriteVFormatAttribute", &TextWriterWriteVFormatAttribute},
	{"xmlTextWriterWriteVFormatAttributeNS", &TextWriterWriteVFormatAttributeNS},
	{"xmlTextWriterWriteVFormatCDATA", &TextWriterWriteVFormatCDATA},
	{"xmlTextWriterWriteVFormatComment", &TextWriterWriteVFormatComment},
	{"xmlTextWriterWriteVFormatDTD", &TextWriterWriteVFormatDTD},
	{"xmlTextWriterWriteVFormatDTDAttlist", &TextWriterWriteVFormatDTDAttlist},
	{"xmlTextWriterWriteVFormatDTDElement", &TextWriterWriteVFormatDTDElement},
	{"xmlTextWriterWriteVFormatDTDInternalEntity", &TextWriterWriteVFormatDTDInternalEntity},
	{"xmlTextWriterWriteVFormatElement", &TextWriterWriteVFormatElement},
	{"xmlTextWriterWriteVFormatElementNS", &TextWriterWriteVFormatElementNS},
	{"xmlTextWriterWriteVFormatPI", &TextWriterWriteVFormatPI},
	{"xmlTextWriterWriteVFormatRaw", &TextWriterWriteVFormatRaw},
	{"xmlTextWriterWriteVFormatString", &TextWriterWriteVFormatString},
	{"xmlThrDefBufferAllocScheme", &ThrDefBufferAllocScheme},
	{"xmlThrDefDefaultBufferSize", &ThrDefDefaultBufferSize},
	{"xmlThrDefDeregisterNodeDefault", &ThrDefDeregisterNodeDefault},
	{"xmlThrDefDoValidityCheckingDefaultValue", &ThrDefDoValidityCheckingDefaultValue},
	{"xmlThrDefGetWarningsDefaultValue", &ThrDefGetWarningsDefaultValue},
	{"xmlThrDefIndentTreeOutput", &ThrDefIndentTreeOutput},
	{"xmlThrDefKeepBlanksDefaultValue", &ThrDefKeepBlanksDefaultValue},
	{"xmlThrDefLineNumbersDefaultValue", &ThrDefLineNumbersDefaultValue},
	{"xmlThrDefLoadExtDtdDefaultValue", &ThrDefLoadExtDtdDefaultValue},
	{"xmlThrDefOutputBufferCreateFilenameDefault", &ThrDefOutputBufferCreateFilenameDefault},
	{"xmlThrDefParserDebugEntities", &ThrDefParserDebugEntities},
	{"xmlThrDefParserInputBufferCreateFilenameDefault", &ThrDefParserInputBufferCreateFilenameDefault},
	{"xmlThrDefPedanticParserDefaultValue", &ThrDefPedanticParserDefaultValue},
	{"xmlThrDefRegisterNodeDefault", &ThrDefRegisterNodeDefault},
	{"xmlThrDefSaveNoEmptyTags", &ThrDefSaveNoEmptyTags},
	{"xmlThrDefSetGenericErrorFunc", &ThrDefSetGenericErrorFunc},
	{"xmlThrDefSetStructuredErrorFunc", &ThrDefSetStructuredErrorFunc},
	{"xmlThrDefSubstituteEntitiesDefaultValue", &ThrDefSubstituteEntitiesDefaultValue},
	{"xmlThrDefTreeIndentString", &ThrDefTreeIndentString},
	{"xmlUCSIsAegeanNumbers", &UCSIsAegeanNumbers},
	{"xmlUCSIsAlphabeticPresentationForms", &UCSIsAlphabeticPresentationForms},
	{"xmlUCSIsArabic", &UCSIsArabic},
	{"xmlUCSIsArabicPresentationFormsA", &UCSIsArabicPresentationFormsA},
	{"xmlUCSIsArabicPresentationFormsB", &UCSIsArabicPresentationFormsB},
	{"xmlUCSIsArmenian", &UCSIsArmenian},
	{"xmlUCSIsArrows", &UCSIsArrows},
	{"xmlUCSIsBasicLatin", &UCSIsBasicLatin},
	{"xmlUCSIsBengali", &UCSIsBengali},
	{"xmlUCSIsBlock", &UCSIsBlock},
	{"xmlUCSIsBlockElements", &UCSIsBlockElements},
	{"xmlUCSIsBopomofo", &UCSIsBopomofo},
	{"xmlUCSIsBopomofoExtended", &UCSIsBopomofoExtended},
	{"xmlUCSIsBoxDrawing", &UCSIsBoxDrawing},
	{"xmlUCSIsBraillePatterns", &UCSIsBraillePatterns},
	{"xmlUCSIsBuhid", &UCSIsBuhid},
	{"xmlUCSIsByzantineMusicalSymbols", &UCSIsByzantineMusicalSymbols},
	{"xmlUCSIsCJKCompatibility", &UCSIsCJKCompatibility},
	{"xmlUCSIsCJKCompatibilityForms", &UCSIsCJKCompatibilityForms},
	{"xmlUCSIsCJKCompatibilityIdeographs", &UCSIsCJKCompatibilityIdeographs},
	{"xmlUCSIsCJKCompatibilityIdeographsSupplement", &UCSIsCJKCompatibilityIdeographsSupplement},
	{"xmlUCSIsCJKRadicalsSupplement", &UCSIsCJKRadicalsSupplement},
	{"xmlUCSIsCJKSymbolsandPunctuation", &UCSIsCJKSymbolsandPunctuation},
	{"xmlUCSIsCJKUnifiedIdeographs", &UCSIsCJKUnifiedIdeographs},
	{"xmlUCSIsCJKUnifiedIdeographsExtensionA", &UCSIsCJKUnifiedIdeographsExtensionA},
	{"xmlUCSIsCJKUnifiedIdeographsExtensionB", &UCSIsCJKUnifiedIdeographsExtensionB},
	{"xmlUCSIsCat", &UCSIsCat},
	{"xmlUCSIsCatC", &UCSIsCatC},
	{"xmlUCSIsCatCc", &UCSIsCatCc},
	{"xmlUCSIsCatCf", &UCSIsCatCf},
	{"xmlUCSIsCatCo", &UCSIsCatCo},
	{"xmlUCSIsCatCs", &UCSIsCatCs},
	{"xmlUCSIsCatL", &UCSIsCatL},
	{"xmlUCSIsCatLl", &UCSIsCatLl},
	{"xmlUCSIsCatLm", &UCSIsCatLm},
	{"xmlUCSIsCatLo", &UCSIsCatLo},
	{"xmlUCSIsCatLt", &UCSIsCatLt},
	{"xmlUCSIsCatLu", &UCSIsCatLu},
	{"xmlUCSIsCatM", &UCSIsCatM},
	{"xmlUCSIsCatMc", &UCSIsCatMc},
	{"xmlUCSIsCatMe", &UCSIsCatMe},
	{"xmlUCSIsCatMn", &UCSIsCatMn},
	{"xmlUCSIsCatN", &UCSIsCatN},
	{"xmlUCSIsCatNd", &UCSIsCatNd},
	{"xmlUCSIsCatNl", &UCSIsCatNl},
	{"xmlUCSIsCatNo", &UCSIsCatNo},
	{"xmlUCSIsCatP", &UCSIsCatP},
	{"xmlUCSIsCatPc", &UCSIsCatPc},
	{"xmlUCSIsCatPd", &UCSIsCatPd},
	{"xmlUCSIsCatPe", &UCSIsCatPe},
	{"xmlUCSIsCatPf", &UCSIsCatPf},
	{"xmlUCSIsCatPi", &UCSIsCatPi},
	{"xmlUCSIsCatPo", &UCSIsCatPo},
	{"xmlUCSIsCatPs", &UCSIsCatPs},
	{"xmlUCSIsCatS", &UCSIsCatS},
	{"xmlUCSIsCatSc", &UCSIsCatSc},
	{"xmlUCSIsCatSk", &UCSIsCatSk},
	{"xmlUCSIsCatSm", &UCSIsCatSm},
	{"xmlUCSIsCatSo", &UCSIsCatSo},
	{"xmlUCSIsCatZ", &UCSIsCatZ},
	{"xmlUCSIsCatZl", &UCSIsCatZl},
	{"xmlUCSIsCatZp", &UCSIsCatZp},
	{"xmlUCSIsCatZs", &UCSIsCatZs},
	{"xmlUCSIsCherokee", &UCSIsCherokee},
	{"xmlUCSIsCombiningDiacriticalMarks", &UCSIsCombiningDiacriticalMarks},
	{"xmlUCSIsCombiningDiacriticalMarksforSymbols", &UCSIsCombiningDiacriticalMarksforSymbols},
	{"xmlUCSIsCombiningHalfMarks", &UCSIsCombiningHalfMarks},
	{"xmlUCSIsCombiningMarksforSymbols", &UCSIsCombiningMarksforSymbols},
	{"xmlUCSIsControlPictures", &UCSIsControlPictures},
	{"xmlUCSIsCurrencySymbols", &UCSIsCurrencySymbols},
	{"xmlUCSIsCypriotSyllabary", &UCSIsCypriotSyllabary},
	{"xmlUCSIsCyrillic", &UCSIsCyrillic},
	{"xmlUCSIsCyrillicSupplement", &UCSIsCyrillicSupplement},
	{"xmlUCSIsDeseret", &UCSIsDeseret},
	{"xmlUCSIsDevanagari", &UCSIsDevanagari},
	{"xmlUCSIsDingbats", &UCSIsDingbats},
	{"xmlUCSIsEnclosedAlphanumerics", &UCSIsEnclosedAlphanumerics},
	{"xmlUCSIsEnclosedCJKLettersandMonths", &UCSIsEnclosedCJKLettersandMonths},
	{"xmlUCSIsEthiopic", &UCSIsEthiopic},
	{"xmlUCSIsGeneralPunctuation", &UCSIsGeneralPunctuation},
	{"xmlUCSIsGeometricShapes", &UCSIsGeometricShapes},
	{"xmlUCSIsGeorgian", &UCSIsGeorgian},
	{"xmlUCSIsGothic", &UCSIsGothic},
	{"xmlUCSIsGreek", &UCSIsGreek},
	{"xmlUCSIsGreekExtended", &UCSIsGreekExtended},
	{"xmlUCSIsGreekandCoptic", &UCSIsGreekandCoptic},
	{"xmlUCSIsGujarati", &UCSIsGujarati},
	{"xmlUCSIsGurmukhi", &UCSIsGurmukhi},
	{"xmlUCSIsHalfwidthandFullwidthForms", &UCSIsHalfwidthandFullwidthForms},
	{"xmlUCSIsHangulCompatibilityJamo", &UCSIsHangulCompatibilityJamo},
	{"xmlUCSIsHangulJamo", &UCSIsHangulJamo},
	{"xmlUCSIsHangulSyllables", &UCSIsHangulSyllables},
	{"xmlUCSIsHanunoo", &UCSIsHanunoo},
	{"xmlUCSIsHebrew", &UCSIsHebrew},
	{"xmlUCSIsHighPrivateUseSurrogates", &UCSIsHighPrivateUseSurrogates},
	{"xmlUCSIsHighSurrogates", &UCSIsHighSurrogates},
	{"xmlUCSIsHiragana", &UCSIsHiragana},
	{"xmlUCSIsIPAExtensions", &UCSIsIPAExtensions},
	{"xmlUCSIsIdeographicDescriptionCharacters", &UCSIsIdeographicDescriptionCharacters},
	{"xmlUCSIsKanbun", &UCSIsKanbun},
	{"xmlUCSIsKangxiRadicals", &UCSIsKangxiRadicals},
	{"xmlUCSIsKannada", &UCSIsKannada},
	{"xmlUCSIsKatakana", &UCSIsKatakana},
	{"xmlUCSIsKatakanaPhoneticExtensions", &UCSIsKatakanaPhoneticExtensions},
	{"xmlUCSIsKhmer", &UCSIsKhmer},
	{"xmlUCSIsKhmerSymbols", &UCSIsKhmerSymbols},
	{"xmlUCSIsLao", &UCSIsLao},
	{"xmlUCSIsLatin1Supplement", &UCSIsLatin1Supplement},
	{"xmlUCSIsLatinExtendedA", &UCSIsLatinExtendedA},
	{"xmlUCSIsLatinExtendedAdditional", &UCSIsLatinExtendedAdditional},
	{"xmlUCSIsLatinExtendedB", &UCSIsLatinExtendedB},
	{"xmlUCSIsLetterlikeSymbols", &UCSIsLetterlikeSymbols},
	{"xmlUCSIsLimbu", &UCSIsLimbu},
	{"xmlUCSIsLinearBIdeograms", &UCSIsLinearBIdeograms},
	{"xmlUCSIsLinearBSyllabary", &UCSIsLinearBSyllabary},
	{"xmlUCSIsLowSurrogates", &UCSIsLowSurrogates},
	{"xmlUCSIsMalayalam", &UCSIsMalayalam},
	{"xmlUCSIsMathematicalAlphanumericSymbols", &UCSIsMathematicalAlphanumericSymbols},
	{"xmlUCSIsMathematicalOperators", &UCSIsMathematicalOperators},
	{"xmlUCSIsMiscellaneousMathematicalSymbolsA", &UCSIsMiscellaneousMathematicalSymbolsA},
	{"xmlUCSIsMiscellaneousMathematicalSymbolsB", &UCSIsMiscellaneousMathematicalSymbolsB},
	{"xmlUCSIsMiscellaneousSymbols", &UCSIsMiscellaneousSymbols},
	{"xmlUCSIsMiscellaneousSymbolsandArrows", &UCSIsMiscellaneousSymbolsandArrows},
	{"xmlUCSIsMiscellaneousTechnical", &UCSIsMiscellaneousTechnical},
	{"xmlUCSIsMongolian", &UCSIsMongolian},
	{"xmlUCSIsMusicalSymbols", &UCSIsMusicalSymbols},
	{"xmlUCSIsMyanmar", &UCSIsMyanmar},
	{"xmlUCSIsNumberForms", &UCSIsNumberForms},
	{"xmlUCSIsOgham", &UCSIsOgham},
	{"xmlUCSIsOldItalic", &UCSIsOldItalic},
	{"xmlUCSIsOpticalCharacterRecognition", &UCSIsOpticalCharacterRecognition},
	{"xmlUCSIsOriya", &UCSIsOriya},
	{"xmlUCSIsOsmanya", &UCSIsOsmanya},
	{"xmlUCSIsPhoneticExtensions", &UCSIsPhoneticExtensions},
	{"xmlUCSIsPrivateUse", &UCSIsPrivateUse},
	{"xmlUCSIsPrivateUseArea", &UCSIsPrivateUseArea},
	{"xmlUCSIsRunic", &UCSIsRunic},
	{"xmlUCSIsShavian", &UCSIsShavian},
	{"xmlUCSIsSinhala", &UCSIsSinhala},
	{"xmlUCSIsSmallFormVariants", &UCSIsSmallFormVariants},
	{"xmlUCSIsSpacingModifierLetters", &UCSIsSpacingModifierLetters},
	{"xmlUCSIsSpecials", &UCSIsSpecials},
	{"xmlUCSIsSuperscriptsandSubscripts", &UCSIsSuperscriptsandSubscripts},
	{"xmlUCSIsSupplementalArrowsA", &UCSIsSupplementalArrowsA},
	{"xmlUCSIsSupplementalArrowsB", &UCSIsSupplementalArrowsB},
	{"xmlUCSIsSupplementalMathematicalOperators", &UCSIsSupplementalMathematicalOperators},
	{"xmlUCSIsSupplementaryPrivateUseAreaA", &UCSIsSupplementaryPrivateUseAreaA},
	{"xmlUCSIsSupplementaryPrivateUseAreaB", &UCSIsSupplementaryPrivateUseAreaB},
	{"xmlUCSIsSyriac", &UCSIsSyriac},
	{"xmlUCSIsTagalog", &UCSIsTagalog},
	{"xmlUCSIsTagbanwa", &UCSIsTagbanwa},
	{"xmlUCSIsTags", &UCSIsTags},
	{"xmlUCSIsTaiLe", &UCSIsTaiLe},
	{"xmlUCSIsTaiXuanJingSymbols", &UCSIsTaiXuanJingSymbols},
	{"xmlUCSIsTamil", &UCSIsTamil},
	{"xmlUCSIsTelugu", &UCSIsTelugu},
	{"xmlUCSIsThaana", &UCSIsThaana},
	{"xmlUCSIsThai", &UCSIsThai},
	{"xmlUCSIsTibetan", &UCSIsTibetan},
	{"xmlUCSIsUgaritic", &UCSIsUgaritic},
	{"xmlUCSIsUnifiedCanadianAboriginalSyllabics", &UCSIsUnifiedCanadianAboriginalSyllabics},
	{"xmlUCSIsVariationSelectors", &UCSIsVariationSelectors},
	{"xmlUCSIsVariationSelectorsSupplement", &UCSIsVariationSelectorsSupplement},
	{"xmlUCSIsYiRadicals", &UCSIsYiRadicals},
	{"xmlUCSIsYiSyllables", &UCSIsYiSyllables},
	{"xmlUCSIsYijingHexagramSymbols", &UCSIsYijingHexagramSymbols},
	{"xmlURIEscape", &URIEscape},
	{"xmlURIEscapeStr", &URIEscapeStr},
	{"xmlURIUnescapeString", &URIUnescapeString},
	{"xmlUTF8Charcmp", &UTF8Charcmp},
	{"xmlUTF8Size", &UTF8Size},
	{"xmlUTF8Strlen", &UTF8Strlen},
	{"xmlUTF8Strloc", &UTF8Strloc},
	{"xmlUTF8Strndup", &UTF8Strndup},
	{"xmlUTF8Strpos", &UTF8Strpos},
	{"xmlUTF8Strsize", &UTF8Strsize},
	{"xmlUTF8Strsub", &UTF8Strsub},
	{"xmlUnlinkNode", &UnlinkNode},
	{"xmlUnlockLibrary", &UnlockLibrary},
	{"xmlUnsetNsProp", &UnsetNsProp},
	{"xmlUnsetProp", &UnsetProp},
	{"xmlValidBuildContentModel", &ValidBuildContentModel},
	{"xmlValidCtxtNormalizeAttributeValue", &ValidCtxtNormalizeAttributeValue},
	{"xmlValidGetPotentialChildren", &ValidGetPotentialChildren},
	{"xmlValidGetValidElements", &ValidGetValidElements},
	{"xmlValidNormalizeAttributeValue", &ValidNormalizeAttributeValue},
	{"xmlValidateAttributeDecl", &ValidateAttributeDecl},
	{"xmlValidateAttributeValue", &ValidateAttributeValue},
	{"xmlValidateDocument", &ValidateDocument},
	{"xmlValidateDocumentFinal", &ValidateDocumentFinal},
	{"xmlValidateDtd", &ValidateDtd},
	{"xmlValidateDtdFinal", &ValidateDtdFinal},
	{"xmlValidateElement", &ValidateElement},
	{"xmlValidateElementDecl", &ValidateElementDecl},
	{"xmlValidateNCName", &ValidateNCName},
	{"xmlValidateNMToken", &ValidateNMToken},
	{"xmlValidateName", &ValidateName},
	{"xmlValidateNameValue", &ValidateNameValue},
	{"xmlValidateNamesValue", &ValidateNamesValue},
	{"xmlValidateNmtokenValue", &ValidateNmtokenValue},
	{"xmlValidateNmtokensValue", &ValidateNmtokensValue},
	{"xmlValidateNotationDecl", &ValidateNotationDecl},
	{"xmlValidateNotationUse", &ValidateNotationUse},
	{"xmlValidateOneAttribute", &ValidateOneAttribute},
	{"xmlValidateOneElement", &ValidateOneElement},
	{"xmlValidateOneNamespace", &ValidateOneNamespace},
	{"xmlValidatePopElement", &ValidatePopElement},
	{"xmlValidatePushCData", &ValidatePushCData},
	{"xmlValidatePushElement", &ValidatePushElement},
	{"xmlValidateQName", &ValidateQName},
	{"xmlValidateRoot", &ValidateRoot},
	{"xmlXIncludeFreeContext", &XIncludeFreeContext},
	{"xmlXIncludeNewContext", &XIncludeNewContext},
	{"xmlXIncludeProcess", &XIncludeProcess},
	{"xmlXIncludeProcessFlags", &XIncludeProcessFlags},
	{"xmlXIncludeProcessFlagsData", &XIncludeProcessFlagsData},
	{"xmlXIncludeProcessNode", &XIncludeProcessNode},
	{"xmlXIncludeProcessTree", &XIncludeProcessTree},
	{"xmlXIncludeProcessTreeFlags", &XIncludeProcessTreeFlags},
	{"xmlXIncludeProcessTreeFlagsData", &XIncludeProcessTreeFlagsData},
	{"xmlXIncludeSetFlags", &XIncludeSetFlags},
	{"xmlXPathAddValues", &XPathAddValues},
	{"xmlXPathBooleanFunction", &XPathBooleanFunction},
	{"xmlXPathCastBooleanToNumber", &XPathCastBooleanToNumber},
	{"xmlXPathCastBooleanToString", &XPathCastBooleanToString},
	{"xmlXPathCastNodeSetToBoolean", &XPathCastNodeSetToBoolean},
	{"xmlXPathCastNodeSetToNumber", &XPathCastNodeSetToNumber},
	{"xmlXPathCastNodeSetToString", &XPathCastNodeSetToString},
	{"xmlXPathCastNodeToNumber", &XPathCastNodeToNumber},
	{"xmlXPathCastNodeToString", &XPathCastNodeToString},
	{"xmlXPathCastNumberToBoolean", &XPathCastNumberToBoolean},
	{"xmlXPathCastNumberToString", &XPathCastNumberToString},
	{"xmlXPathCastStringToBoolean", &XPathCastStringToBoolean},
	{"xmlXPathCastStringToNumber", &XPathCastStringToNumber},
	{"xmlXPathCastToBoolean", &XPathCastToBoolean},
	{"xmlXPathCastToNumber", &XPathCastToNumber},
	{"xmlXPathCastToString", &XPathCastToString},
	{"xmlXPathCeilingFunction", &XPathCeilingFunction},
	{"xmlXPathCmpNodes", &XPathCmpNodes},
	{"xmlXPathCompareValues", &XPathCompareValues},
	{"xmlXPathCompile", &XPathCompile},
	{"xmlXPathCompiledEval", &XPathCompiledEval},
	{"xmlXPathCompiledEvalToBoolean", &XPathCompiledEvalToBoolean},
	{"xmlXPathConcatFunction", &XPathConcatFunction},
	{"xmlXPathContainsFunction", &XPathContainsFunction},
	{"xmlXPathContextSetCache", &XPathContextSetCache},
	{"xmlXPathConvertBoolean", &XPathConvertBoolean},
	{"xmlXPathConvertNumber", &XPathConvertNumber},
	{"xmlXPathConvertString", &XPathConvertString},
	{"xmlXPathCountFunction", &XPathCountFunction},
	{"xmlXPathCtxtCompile", &XPathCtxtCompile},
	{"xmlXPathDebugDumpCompExpr", &XPathDebugDumpCompExpr},
	{"xmlXPathDebugDumpObject", &XPathDebugDumpObject},
	{"xmlXPathDifference", &XPathDifference},
	{"xmlXPathDistinct", &XPathDistinct},
	{"xmlXPathDistinctSorted", &XPathDistinctSorted},
	{"xmlXPathDivValues", &XPathDivValues},
	{"xmlXPathEqualValues", &XPathEqualValues},
	{"xmlXPathErr", &XPathErr},
	{"xmlXPathEval", &XPathEval},
	{"xmlXPathEvalExpr", &XPathEvalExpr},
	{"xmlXPathEvalExpression", &XPathEvalExpression},
	{"xmlXPathEvalPredicate", &XPathEvalPredicate},
	{"xmlXPathEvaluatePredicateResult", &XPathEvaluatePredicateResult},
	{"xmlXPathFalseFunction", &XPathFalseFunction},
	{"xmlXPathFloorFunction", &XPathFloorFunction},
	{"xmlXPathFreeCompExpr", &XPathFreeCompExpr},
	{"xmlXPathFreeContext", &XPathFreeContext},
	{"xmlXPathFreeNodeSet", &XPathFreeNodeSet},
	{"xmlXPathFreeNodeSetList", &XPathFreeNodeSetList},
	{"xmlXPathFreeObject", &XPathFreeObject},
	{"xmlXPathFreeParserContext", &XPathFreeParserContext},
	{"xmlXPathFunctionLookup", &XPathFunctionLookup},
	{"xmlXPathFunctionLookupNS", &XPathFunctionLookupNS},
	{"xmlXPathHasSameNodes", &XPathHasSameNodes},
	{"xmlXPathIdFunction", &XPathIdFunction},
	{"xmlXPathInit", &XPathInit},
	{"xmlXPathIntersection", &XPathIntersection},
	{"xmlXPathIsInf", &XPathIsInf},
	{"xmlXPathIsNaN", &XPathIsNaN},
	{"xmlXPathIsNodeType", &XPathIsNodeType},
	{"xmlXPathLangFunction", &XPathLangFunction},
	{"xmlXPathLastFunction", &XPathLastFunction},
	{"xmlXPathLeading", &XPathLeading},
	{"xmlXPathLeadingSorted", &XPathLeadingSorted},
	{"xmlXPathLocalNameFunction", &XPathLocalNameFunction},
	{"xmlXPathModValues", &XPathModValues},
	{"xmlXPathMultValues", &XPathMultValues},
	{"xmlXPathNamespaceURIFunction", &XPathNamespaceURIFunction},
	{"xmlXPathNewBoolean", &XPathNewBoolean},
	{"xmlXPathNewCString", &XPathNewCString},
	{"xmlXPathNewContext", &XPathNewContext},
	{"xmlXPathNewFloat", &XPathNewFloat},
	{"xmlXPathNewNodeSet", &XPathNewNodeSet},
	{"xmlXPathNewNodeSetList", &XPathNewNodeSetList},
	{"xmlXPathNewParserContext", &XPathNewParserContext},
	{"xmlXPathNewString", &XPathNewString},
	{"xmlXPathNewValueTree", &XPathNewValueTree},
	{"xmlXPathNextAncestor", &XPathNextAncestor},
	{"xmlXPathNextAncestorOrSelf", &XPathNextAncestorOrSelf},
	{"xmlXPathNextAttribute", &XPathNextAttribute},
	{"xmlXPathNextChild", &XPathNextChild},
	{"xmlXPathNextDescendant", &XPathNextDescendant},
	{"xmlXPathNextDescendantOrSelf", &XPathNextDescendantOrSelf},
	{"xmlXPathNextFollowing", &XPathNextFollowing},
	{"xmlXPathNextFollowingSibling", &XPathNextFollowingSibling},
	{"xmlXPathNextNamespace", &XPathNextNamespace},
	{"xmlXPathNextParent", &XPathNextParent},
	{"xmlXPathNextPreceding", &XPathNextPreceding},
	{"xmlXPathNextPrecedingSibling", &XPathNextPrecedingSibling},
	{"xmlXPathNextSelf", &XPathNextSelf},
	{"xmlXPathNodeLeading", &XPathNodeLeading},
	{"xmlXPathNodeLeadingSorted", &XPathNodeLeadingSorted},
	{"xmlXPathNodeSetAdd", &XPathNodeSetAdd},
	{"xmlXPathNodeSetAddNs", &XPathNodeSetAddNs},
	{"xmlXPathNodeSetAddUnique", &XPathNodeSetAddUnique},
	{"xmlXPathNodeSetContains", &XPathNodeSetContains},
	{"xmlXPathNodeSetCreate", &XPathNodeSetCreate},
	{"xmlXPathNodeSetDel", &XPathNodeSetDel},
	{"xmlXPathNodeSetFreeNs", &XPathNodeSetFreeNs},
	{"xmlXPathNodeSetMerge", &XPathNodeSetMerge},
	{"xmlXPathNodeSetRemove", &XPathNodeSetRemove},
	{"xmlXPathNodeSetSort", &XPathNodeSetSort},
	{"xmlXPathNodeTrailing", &XPathNodeTrailing},
	{"xmlXPathNodeTrailingSorted", &XPathNodeTrailingSorted},
	{"xmlXPathNormalizeFunction", &XPathNormalizeFunction},
	{"xmlXPathNotEqualValues", &XPathNotEqualValues},
	{"xmlXPathNotFunction", &XPathNotFunction},
	{"xmlXPathNsLookup", &XPathNsLookup},
	{"xmlXPathNumberFunction", &XPathNumberFunction},
	{"xmlXPathObjectCopy", &XPathObjectCopy},
	{"xmlXPathOrderDocElems", &XPathOrderDocElems},
	{"xmlXPathParseNCName", &XPathParseNCName},
	{"xmlXPathParseName", &XPathParseName},
	{"xmlXPathPopBoolean", &XPathPopBoolean},
	{"xmlXPathPopExternal", &XPathPopExternal},
	{"xmlXPathPopNodeSet", &XPathPopNodeSet},
	{"xmlXPathPopNumber", &XPathPopNumber},
	{"xmlXPathPopString", &XPathPopString},
	{"xmlXPathPositionFunction", &XPathPositionFunction},
	{"xmlXPathRegisterAllFunctions", &XPathRegisterAllFunctions},
	{"xmlXPathRegisterFunc", &XPathRegisterFunc},
	{"xmlXPathRegisterFuncLookup", &XPathRegisterFuncLookup},
	{"xmlXPathRegisterFuncNS", &XPathRegisterFuncNS},
	{"xmlXPathRegisterNs", &XPathRegisterNs},
	{"xmlXPathRegisterVariable", &XPathRegisterVariable},
	{"xmlXPathRegisterVariableLookup", &XPathRegisterVariableLookup},
	{"xmlXPathRegisterVariableNS", &XPathRegisterVariableNS},
	{"xmlXPathRegisteredFuncsCleanup", &XPathRegisteredFuncsCleanup},
	{"xmlXPathRegisteredNsCleanup", &XPathRegisteredNsCleanup},
	{"xmlXPathRegisteredVariablesCleanup", &XPathRegisteredVariablesCleanup},
	{"xmlXPathRoot", &XPathRoot},
	{"xmlXPathRoundFunction", &XPathRoundFunction},
	{"xmlXPathStartsWithFunction", &XPathStartsWithFunction},
	{"xmlXPathStringEvalNumber", &XPathStringEvalNumber},
	{"xmlXPathStringFunction", &XPathStringFunction},
	{"xmlXPathStringLengthFunction", &XPathStringLengthFunction},
	{"xmlXPathSubValues", &XPathSubValues},
	{"xmlXPathSubstringAfterFunction", &XPathSubstringAfterFunction},
	{"xmlXPathSubstringBeforeFunction", &XPathSubstringBeforeFunction},
	{"xmlXPathSubstringFunction", &XPathSubstringFunction},
	{"xmlXPathSumFunction", &XPathSumFunction},
	{"xmlXPathTrailing", &XPathTrailing},
	{"xmlXPathTrailingSorted", &XPathTrailingSorted},
	{"xmlXPathTranslateFunction", &XPathTranslateFunction},
	{"xmlXPathTrueFunction", &XPathTrueFunction},
	{"xmlXPathValueFlipSign", &XPathValueFlipSign},
	{"xmlXPathVariableLookup", &XPathVariableLookup},
	{"xmlXPathVariableLookupNS", &XPathVariableLookupNS},
	{"xmlXPathWrapCString", &XPathWrapCString},
	{"xmlXPathWrapExternal", &XPathWrapExternal},
	{"xmlXPathWrapNodeSet", &XPathWrapNodeSet},
	{"xmlXPathWrapString", &XPathWrapString},
	{"xmlXPatherror", &XPatherror},
	{"xmlXPtrBuildNodeList", &XPtrBuildNodeList},
	{"xmlXPtrEval", &XPtrEval},
	{"xmlXPtrEvalRangePredicate", &XPtrEvalRangePredicate},
	{"xmlXPtrFreeLocationSet", &XPtrFreeLocationSet},
	{"xmlXPtrLocationSetAdd", &XPtrLocationSetAdd},
	{"xmlXPtrLocationSetCreate", &XPtrLocationSetCreate},
	{"xmlXPtrLocationSetDel", &XPtrLocationSetDel},
	{"xmlXPtrLocationSetMerge", &XPtrLocationSetMerge},
	{"xmlXPtrLocationSetRemove", &XPtrLocationSetRemove},
	{"xmlXPtrNewCollapsedRange", &XPtrNewCollapsedRange},
	{"xmlXPtrNewContext", &XPtrNewContext},
	{"xmlXPtrNewLocationSetNodeSet", &XPtrNewLocationSetNodeSet},
	{"xmlXPtrNewLocationSetNodes", &XPtrNewLocationSetNodes},
	{"xmlXPtrNewRange", &XPtrNewRange},
	{"xmlXPtrNewRangeNodeObject", &XPtrNewRangeNodeObject},
	{"xmlXPtrNewRangeNodePoint", &XPtrNewRangeNodePoint},
	{"xmlXPtrNewRangeNodes", &XPtrNewRangeNodes},
	{"xmlXPtrNewRangePointNode", &XPtrNewRangePointNode},
	{"xmlXPtrNewRangePoints", &XPtrNewRangePoints},
	{"xmlXPtrRangeToFunction", &XPtrRangeToFunction},
	{"xmlXPtrWrapLocationSet", &XPtrWrapLocationSet},
}

type (
	Free         fix
	Malloc       fix
	MallocAtomic fix
	MemStrdup    fix
	Realloc      fix
)

var dataList = outside.Data{
	{"emptyExp", (*ExpNode)(nil)},
	{"forbiddenExp", (*ExpNode)(nil)},
	{"xmlFree", (*Free)(nil)},
	{"xmlIsBaseCharGroup", (*ChRangeGroup)(nil)},
	{"xmlIsCharGroup", (*ChRangeGroup)(nil)},
	{"xmlIsCombiningGroup", (*ChRangeGroup)(nil)},
	{"xmlIsDigitGroup", (*ChRangeGroup)(nil)},
	{"xmlIsExtenderGroup", (*ChRangeGroup)(nil)},
	{"xmlIsIdeographicGroup", (*ChRangeGroup)(nil)},
	{"xmlIsPubidChar_tab", (*IsPubidCharTab)(nil)},
	{"xmlMalloc", (*Malloc)(nil)},
	{"xmlMallocAtomic", (*MallocAtomic)(nil)},
	{"xmlMemStrdup", (*MemStrdup)(nil)},
	{"xmlParserMaxDepth", (*uint)(nil)},
	{"xmlRealloc", (*Realloc)(nil)},
	{"xmlStringComment", (*uint8)(nil)},
	{"xmlStringText", (*uint8)(nil)},
	{"xmlStringTextNoenc", (*uint8)(nil)},
	{"xmlXPathNAN", (*float64)(nil)},
	{"xmlXPathNINF", (*float64)(nil)},
	{"xmlXPathPINF", (*float64)(nil)},
}
