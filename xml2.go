// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package xml2 provides API definitions for accessing the libxml2 dll.
package xml2

import (
	. "github.com/tHinqa/outside"
)

func init() {
	AddDllApis(dll, false, apiList)
	AddDllData(dll, false, dataList)
}

type (
	fix uintptr

	FILE    fix
	SOCKET  fix
	Va_list fix

	Char           int8
	Double         float64
	Enum           int
	Long           int32 // TODO(t): Size?
	Size_t         uint
	Unsigned_char  uint8
	Unsigned_int   uint
	Unsigned_long  uint32 //TODO(t): check  size
	Unsigned_short uint16

	Void struct{}
)

type (
	DocbDocPtr                 XmlDocPtr
	DocbParserCtxt             XmlParserCtxt
	DocbParserCtxtPtr          XmlParserCtxtPtr
	DocbParserInput            XmlParserInput
	DocbParserInputPtr         XmlParserInputPtr
	DocbSAXHandler             XmlSAXHandler
	DocbSAXHandlerPtr          XmlSAXHandlerPtr
	HtmlDocPtr                 XmlDocPtr
	HtmlNodePtr                XmlNodePtr
	HtmlParserCtxt             XmlParserCtxt
	HtmlParserCtxtPtr          XmlParserCtxtPtr
	HtmlSAXHandler             XmlSAXHandler
	HtmlSAXHandlerPtr          XmlSAXHandlerPtr
	XlinkHandlerPtr            *XlinkHandler
	XlinkHRef                  *XmlChar
	XlinkRole                  *XmlChar
	XlinkTitle                 *XmlChar
	XmlAttributePtr            *XmlAttribute
	XmlAttributeTable          XmlHashTable
	XmlAttributeTablePtr       *XmlAttributeTable
	XmlAttrPtr                 *XmlAttr
	XmlAutomataPtr             *XmlAutomata
	XmlAutomataStatePtr        *XmlAutomataState
	XmlBufferPtr               *XmlBuffer
	XmlBufPtr                  *XmlBuf
	XmlCatalogPtr              *XmlCatalog
	XmlChar                    Unsigned_char
	XmlCharEncodingHandlerPtr  *XmlCharEncodingHandler
	XmlDictPtr                 *XmlDict
	XmlDocPtr                  *XmlDoc
	XmlDOMWrapCtxtPtr          *XmlDOMWrapCtxt
	XmlDtdPtr                  *XmlDtd
	XmlElementContentPtr       *XmlElementContent
	XmlElementPtr              *XmlElement
	XmlElementTable            XmlHashTable
	XmlElementTablePtr         *XmlElementTable
	XmlEntitiesTable           XmlHashTable
	XmlEntitiesTablePtr        *XmlEntitiesTable
	XmlEntityPtr               *XmlEntity
	XmlEnumerationPtr          *XmlEnumeration
	XmlErrorPtr                *XmlError
	XmlExpCtxtPtr              *XmlExpCtxt
	XmlExpNodePtr              *XmlExpNode
	XmlGlobalStatePtr          *XmlGlobalState
	XmlHashTablePtr            *XmlHashTable
	XmlIDPtr                   *XmlID
	XmlIDTable                 XmlHashTable
	XmlIDTablePtr              *XmlIDTable
	XmlLinkPtr                 *XmlLink
	XmlListPtr                 *XmlList
	XmlLocationSetPtr          *XmlLocationSet
	XmlModulePtr               *XmlModule
	XmlMutexPtr                *XmlMutex
	XmlNodePtr                 *XmlNode
	XmlNodeSetPtr              *XmlNodeSet
	XmlNotationPtr             *XmlNotation
	XmlNotationTable           XmlHashTable
	XmlNotationTablePtr        *XmlNotationTable
	XmlNsPtr                   *XmlNs
	XmlNsType                  XmlElementType
	XmlOutputBufferPtr         *XmlOutputBuffer
	XmlParserCtxtPtr           *XmlParserCtxt
	XmlParserInputBufferPtr    *XmlParserInputBuffer
	XmlParserInputPtr          *XmlParserInput
	XmlParserNodeInfoPtr       *XmlParserNodeInfo
	XmlParserNodeInfoSeqPtr    *XmlParserNodeInfoSeq
	XmlPatternPtr              *XmlPattern
	XmlRefPtr                  *XmlRef
	XmlRefTable                XmlHashTable
	XmlRefTablePtr             *XmlRefTable
	XmlRegExecCtxtPtr          *XmlRegExecCtxt
	XmlRegexpPtr               *XmlRegexp
	XmlRelaxNGParserCtxtPtr    *XmlRelaxNGParserCtxt
	XmlRelaxNGPtr              *XmlRelaxNG
	XmlRelaxNGValidCtxtPtr     *XmlRelaxNGValidCtxt
	XmlRMutexPtr               *XmlRMutex
	XmlSaveCtxtPtr             *XmlSaveCtxt
	XmlSAXHandlerPtr           *XmlSAXHandler
	XmlSAXLocatorPtr           *XmlSAXLocator
	XmlSchemaAnnotPtr          *XmlSchemaAnnot
	XmlSchemaFacetPtr          *XmlSchemaFacet
	XmlSchemaParserCtxtPtr     *XmlSchemaParserCtxt
	XmlSchemaPtr               *XmlSchema
	XmlSchemaSAXPlugPtr        *XmlSchemaSAXPlugStruct
	XmlSchematronParserCtxtPtr *XmlSchematronParserCtxt
	XmlSchematronPtr           *XmlSchematron
	XmlSchematronValidCtxtPtr  *XmlSchematronValidCtxt
	XmlSchemaTypePtr           *XmlSchemaType
	XmlSchemaValidCtxtPtr      *XmlSchemaValidCtxt
	XmlSchemaValPtr            *XmlSchemaVal
	XmlSchemaWildcardNsPtr     *XmlSchemaWildcardNs
	XmlSchemaWildcardPtr       *XmlSchemaWildcard
	XmlShellCtxtPtr            *XmlShellCtxt
	XmlStreamCtxtPtr           *XmlStreamCtxt
	XmlTextReaderLocatorPtr    *Void
	XmlTextReaderPtr           *XmlTextReader
	XmlTextWriterPtr           *XmlTextWriter
	XmlURIPtr                  *XmlURI
	XmlValidCtxtPtr            *XmlValidCtxt
	XmlValidStatePtr           *XmlValidState
	XmlXIncludeCtxtPtr         *XmlXIncludeCtxt
	XmlXPathAxisPtr            *XmlXPathAxis
	XmlXPathCompExprPtr        *XmlXPathCompExpr
	XmlXPathContextPtr         *XmlXPathContext
	XmlXPathObjectPtr          *XmlXPathObject
	XmlXPathParserContextPtr   *XmlXPathParserContext
	XmlXPathTypePtr            *XmlXPathType

	XmlAutomata             struct{}
	XmlAutomataState        struct{}
	XmlBuf                  struct{}
	XmlBuffer               struct{}
	XmlCatalog              struct{}
	XmlDict                 struct{}
	XmlDOMWrapCtxt          struct{}
	XmlExpCtxt              struct{}
	XmlExpNode              struct{}
	XmlHashTable            struct{}
	XmlLink                 struct{}
	XmlList                 struct{}
	XmlModule               struct{}
	XmlMutex                struct{}
	XmlPattern              struct{}
	XmlRegExecCtxt          struct{}
	XmlRegexp               struct{}
	XmlRelaxNG              struct{}
	XmlRelaxNGParserCtxt    struct{}
	XmlRelaxNGValidCtxt     struct{}
	XmlRMutex               struct{}
	XmlSaveCtxt             struct{}
	XmlSchema               struct{}
	XmlSchemaFacet          struct{}
	XmlSchemaParserCtxt     struct{}
	XmlSchemaSAXPlugStruct  struct{}
	XmlSchematron           struct{}
	XmlSchematronParserCtxt struct{}
	XmlSchematronValidCtxt  struct{}
	XmlSchemaType           struct{}
	XmlSchemaVal            struct{}
	XmlSchemaValidCtxt      struct{}
	XmlStreamCtxt           struct{}
	XmlTextReader           struct{}
	XmlTextWriter           struct{}
	XmlValidState           struct{}
	XmlXIncludeCtxt         struct{}
	XmlXPathCompExpr        struct{}
	XmlXPathParserContext   struct{}

	InternalSubsetSAXFunc func(
		ctx *Void, name, ExternalID, SystemID string)

	IsStandaloneSAXFunc func(ctx *Void) int

	HasInternalSubsetSAXFunc func(ctx *Void) int

	HasExternalSubsetSAXFunc func(ctx *Void) int

	StartElementNsSAX2Func func(ctx *Void,
		localname, prefix, URI string,
		nb_namespaces int, namespaces *string,
		nb_attributes, nb_defaulted int, attributes *string)

	EndElementNsSAX2Func func(
		ctx *Void, localname, prefix, URI string)

	ResolveEntitySAXFunc func(
		ctx *Void,
		publicId string,
		systemId string) XmlParserInputPtr

	ExternalSubsetSAXFunc func(
		ctx *Void,
		name string,
		ExternalID string,
		SystemID string)

	GetEntitySAXFunc func(
		ctx *Void, name string) XmlEntityPtr

	GetParameterEntitySAXFunc func(
		ctx *Void, name string) XmlEntityPtr

	EntityDeclSAXFunc func(ctx *Void, name string,
		typ int, publicId, systemId, content string)

	NotationDeclSAXFunc func(
		ctx *Void, name, publicId, systemId string)

	AttributeDeclSAXFunc func(ctx *Void, elem,
		fullname string, typ, def int,
		defaultValue string, tree XmlEnumerationPtr)

	ElementDeclSAXFunc func(ctx *Void, name string,
		typ int, content XmlElementContentPtr)

	UnparsedEntityDeclSAXFunc func(ctx *Void,
		name, publicId, systemId, notationName string)

	SetDocumentLocatorSAXFunc func(
		ctx *Void, loc XmlSAXLocatorPtr)

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

	XmlParserInputDeallocate func(str string)

	XmlInputReadCallback func(
		context *Void, buffer string, leng int) int

	XmlInputCloseCallback func(context *Void) int

	WarningSAXFunc func(ctx *Void, msg string, v ...VArg)

	ErrorSAXFunc func(ctx *Void, msg string, v ...VArg)

	FatalErrorSAXFunc func(ctx *Void, msg string, v ...VArg)

	XmlDeregisterNodeFunc func(node XmlNodePtr)

	XmlGenericErrorFunc func(ctx *Void, msg string, v ...VArg)

	XmlCharEncodingInputFunc func(
		out *Unsigned_char, outlen *int,
		in *Unsigned_char, inlen *int) int

	XmlCharEncodingOutputFunc func(
		out *Unsigned_char, outlen *int,
		in *Unsigned_char, inlen *int) int

	XmlParserInputBufferCreateFilenameFunc func(
		URI string, enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlRegisterNodeFunc func(node XmlNodePtr)

	XmlStructuredErrorFunc func(
		userData *Void, Error XmlErrorPtr)

	XmlOutputWriteCallback func(
		context *Void, buffer string, leng int) int

	XmlOutputCloseCallback func(context *Void) int

	XmlValidityErrorFunc func(ctx *Void, msg string, v ...VArg)

	XmlValidityWarningFunc func(ctx *Void, msg string, v ...VArg)

	XmlOutputBufferCreateFilenameFunc func(
		URI string,
		encoder XmlCharEncodingHandlerPtr,
		compression int) XmlOutputBufferPtr

	XlinkNodeDetectFunc func(
		ctx *Void,
		node XmlNodePtr)

	XlinkSimpleLinkFunk func(
		ctx *Void,
		node XmlNodePtr,
		href XlinkHRef,
		role XlinkRole,
		title XlinkTitle)

	XlinkExtendedLinkFunk func(
		ctx *Void,
		node XmlNodePtr,
		nbLocators int,
		hrefs *XlinkHRef,
		roles *XlinkRole,
		nbArcs int,
		from *XlinkRole,
		to *XlinkRole,
		show *XlinkShow,
		actuate *XlinkActuate,
		nbTitles int,
		titles *XlinkTitle,
		langs *string)

	XlinkExtendedLinkSetFunk func(
		ctx *Void,
		node XmlNodePtr,
		nbLocators int,
		hrefs *XlinkHRef,
		roles *XlinkRole,
		nbTitles int,
		titles *XlinkTitle,
		langs *string)

	XmlFreeFunc func(mem *Void)

	XmlMallocFunc func(size Size_t) *Void

	XmlReallocFunc func(mem *Void, size Size_t) *Void

	XmlStrdupFunc func(str string) string

	XmlHashDeallocator func(payload *Void, name string)

	XmlHashCopier func(payload *Void, name string) *Void

	XmlHashScanner func(payload, data *Void, name string)

	XmlHashScannerFull func(payload, data *Void,
		name, name2, name3 string)

	XmlExternalEntityLoader func(URL, ID string,
		context XmlParserCtxtPtr) XmlParserInputPtr

	XmlRegExecCallbacks func(exec XmlRegExecCtxtPtr,
		token string, transdata, inputdata *Void)

	XmlListDeallocator func(lk XmlLinkPtr)

	XmlListDataCompare func(data0, data1 *Void) int

	XmlListWalker func(data, user *Void) int

	XmlInputMatchCallback func(filename string) int

	XmlInputOpenCallback func(filename string) *Void

	XmlOutputMatchCallback func(filename string) int

	XmlOutputOpenCallback func(filename string) *Void

	XmlXPathAxisFunc func(ctxt XmlXPathParserContextPtr,
		cur XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPathVariableLookupFunc func(
		ctxt *Void, name, ns_uri string) XmlXPathObjectPtr

	XmlXPathFuncLookupFunc func(
		ctxt *Void, name, ns_uri string) XmlXPathFunction

	XmlXPathFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathConvertFunc func(obj XmlXPathObjectPtr, typ int) int

	XmlSchemaValidityErrorFunc func(
		ctx *Void, msg string, v ...VArg)

	XmlSchemaValidityWarningFunc func(
		ctx *Void, msg string, v ...VArg)

	XmlSchemaValidityLocatorFunc func(
		ctx *Void, file *string, line *Unsigned_long) int

	XmlEntityReferenceFunc func(
		ent XmlEntityPtr, firstNode, lastNode XmlNodePtr)

	XmlRelaxNGValidityErrorFunc func(
		ctx *Void, msg string, v ...VArg)

	XmlRelaxNGValidityWarningFunc func(
		ctx *Void, msg string, v ...VArg)

	XmlTextReaderErrorFunc func(arg *Void, msg string,
		severity XmlParserSeverities,
		locator XmlTextReaderLocatorPtr)

	XmlShellReadlineFunc func(prompt string) string

	FtpListCallback func(
		userData *Void,
		filename, attrib, owner, group string,
		size Unsigned_long,
		links, year int,
		month string,
		day, hour, minute int)

	FtpDataCallback func(userData *Void, data string, leng int)

	XmlC14NIsVisibleCallback func(
		user_data *Void,
		node XmlNodePtr,
		parent XmlNodePtr) int

	XmlParserCtxt struct {
		Sax               *XmlSAXHandler
		UserData          *Void
		MyDoc             XmlDocPtr
		WellFormed        int
		ReplaceEntities   int
		Version           *XmlChar
		Encoding          *XmlChar
		Standalone        int
		Html              int
		Input             XmlParserInputPtr
		InputNr           int
		InputMax          int
		InputTab          *XmlParserInputPtr
		Node              XmlNodePtr
		NodeNr            int
		NodeMax           int
		NodeTab           *XmlNodePtr
		Record_info       int
		Node_seq          XmlParserNodeInfoSeq
		ErrNo             int
		HasExternalSubset int
		HasPErefs         int
		External          int
		Valid             int
		Validate          int
		Vctxt             XmlValidCtxt
		Instate           XmlParserInputState
		Token             int
		Directory         *Char
		Name              *XmlChar
		NameNr            int
		NameMax           int
		NameTab           **XmlChar
		NbChars           Long
		CheckIndex        Long
		KeepBlanks        int
		DisableSAX        int
		InSubset          int
		IntSubName        *XmlChar
		ExtSubURI         *XmlChar
		ExtSubSystem      *XmlChar
		Space             *int
		SpaceNr           int
		SpaceMax          int
		SpaceTab          *int
		Depth             int
		Entity            XmlParserInputPtr
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
		Dict              XmlDictPtr
		Atts              **XmlChar
		Maxatts           int
		Docdict           int
		Str_xml           *XmlChar
		Str_XmlNs         *XmlChar
		Str_xml_ns        *XmlChar
		Sax2              int
		NsNr              int
		NsMax             int
		NsTab             **XmlChar
		Attallocs         *int
		PushTab           **Void
		AttsDefault       XmlHashTablePtr
		AttsSpecial       XmlHashTablePtr
		NsWellFormed      int
		Options           int
		DictNames         int
		FreeElemsNr       int
		FreeElems         XmlNodePtr
		FreeAttrsNr       int
		FreeAttrs         XmlAttrPtr
		LastError         XmlError
		ParseMode         XmlParserMode
		Nbentities        Unsigned_long
		Sizeentities      Unsigned_long
		NodeInfo          *XmlParserNodeInfo
		NodeInfoNr        int
		NodeInfoMax       int
		NodeInfoTab       *XmlParserNodeInfo
		Input_id          int
	}

	XmlSAXHandlerV1 struct {
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
		Initialized           Unsigned_int
	}

	XmlParserInput struct {
		Buf        XmlParserInputBufferPtr
		Filename   *Char
		Directory  *Char
		Base       *XmlChar
		Cur        *XmlChar
		End        *XmlChar
		Length     int
		Line       int
		Col        int
		Consumed   Unsigned_long
		Free       XmlParserInputDeallocate
		Encoding   *XmlChar
		Version    *XmlChar
		Standalone int
		Id         int
	}

	XmlCharEncodingHandler struct {
		Name   *Char
		Input  XmlCharEncodingInputFunc
		Output XmlCharEncodingOutputFunc
	}

	XmlError struct {
		Domain  int
		Code    int
		Message *Char
		Level   XmlErrorLevel
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

	XmlOutputBuffer struct {
		Context       *Void
		Writecallback XmlOutputWriteCallback
		Closecallback XmlOutputCloseCallback
		Encoder       XmlCharEncodingHandlerPtr
		Buffer        XmlBufPtr
		Conv          XmlBufPtr
		Written       int
		Error         int
	}

	XmlParserInputBuffer struct {
		Context       *Void
		Readcallback  XmlInputReadCallback
		Closecallback XmlInputCloseCallback
		Encoder       XmlCharEncodingHandlerPtr
		Buffer        XmlBufPtr
		Raw           XmlBufPtr
		Compressed    int
		Error         int
		Rawconsumed   Unsigned_long
	}

	XmlSAXLocator struct {
		GetPublicId     func(ctx *Void) *XmlChar
		GetSystemId     func(ctx *Void) *XmlChar
		GetLineNumber   func(ctx *Void) int
		GetColumnNumber func(ctx *Void) int
	}

	XmlSAXHandler struct {
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
		Initialized           Unsigned_int
		_                     *Void
		StartElementNs        StartElementNsSAX2Func
		EndElementNs          EndElementNsSAX2Func
		Serror                XmlStructuredErrorFunc
	}

	XmlParserNodeInfoSeq struct {
		Maximum Unsigned_long
		Length  Unsigned_long
		Buffer  *XmlParserNodeInfo
	}

	XmlParserNodeInfo struct {
		Node       *XmlNode
		Begin_pos  Unsigned_long
		Begin_line Unsigned_long
		End_pos    Unsigned_long
		End_line   Unsigned_long
	}

	XmlValidCtxt struct {
		UserData  *Void
		Error     XmlValidityErrorFunc
		Warning   XmlValidityWarningFunc
		Node      XmlNodePtr
		NodeNr    int
		NodeMax   int
		NodeTab   *XmlNodePtr
		FinishDtd Unsigned_int
		Doc       XmlDocPtr
		Valid     int
		Vstate    *XmlValidState
		VstateNr  int
		VstateMax int
		VstateTab *XmlValidState
		Am        XmlAutomataPtr
		State     XmlAutomataStatePtr
	}

	XmlDtd struct {
		_          *Void
		Type       XmlElementType
		Name       *XmlChar
		Children   *XmlNode
		Last       *XmlNode
		Parent     *XmlDoc
		Next       *XmlNode
		Prev       *XmlNode
		Doc        *XmlDoc
		Notations  *Void
		Elements   *Void
		Attributes *Void
		Entities   *Void
		ExternalID *XmlChar
		SystemID   *XmlChar
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
		Attrs_opt     **Char
		Attrs_depr    **Char
		Attrs_req     **Char
	}

	HtmlEntityDesc struct {
		Value Unsigned_int
		Name  *Char
		Desc  *Char
	}

	XlinkHandler struct {
		Simple   XlinkSimpleLinkFunk
		Extended XlinkExtendedLinkFunk
		Set      XlinkExtendedLinkSetFunk
	}

	XmlAttribute struct {
		_    *Void
		Type XmlElementType
		Name *XmlChar
	}

	XmlNotation struct {
		Name     *XmlChar
		PublicID *XmlChar
		SystemID *XmlChar
	}

	XmlElement struct {
		_          *Void
		Type       XmlElementType
		Name       *XmlChar
		Children   *XmlNode
		Last       *XmlNode
		Parent     *XmlDtd
		Next       *XmlNode
		Prev       *XmlNode
		Doc        *XmlDoc
		Etype      XmlElementTypeVal
		Content    XmlElementContentPtr
		Attributes XmlAttributePtr
		Prefix     *XmlChar
		ContModel  XmlRegexpPtr
	}

	XmlID struct {
		Next   *XmlID
		Value  *XmlChar
		Attr   XmlAttrPtr
		Name   *XmlChar
		Lineno int
		Doc    *XmlDoc
	}

	XmlRef struct {
		Next   *XmlRef
		Value  *XmlChar
		Attr   XmlAttrPtr
		Name   *XmlChar
		Lineno int
	}

	XmlGlobalState struct {
		XmlParserVersion                        *Char
		XmlDefaultSAXLocator                    XmlSAXLocator
		XmlDefaultSAXHandler                    XmlSAXHandlerV1
		DocbDefaultSAXHandler                   XmlSAXHandlerV1
		HtmlDefaultSAXHandler                   XmlSAXHandlerV1
		XmlFree                                 XmlFreeFunc
		XmlMalloc                               XmlMallocFunc
		XmlMemStrdup                            XmlStrdupFunc
		XmlRealloc                              XmlReallocFunc
		XmlGenericError                         XmlGenericErrorFunc
		XmlStructuredError                      XmlStructuredErrorFunc
		XmlGenericErrorContext                  *Void
		OldXMLWDcompatibility                   int
		XmlBufferAllocScheme                    XmlBufferAllocationScheme
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
		XmlRegisterNodeDefaultValue             XmlRegisterNodeFunc
		XmlDeregisterNodeDefaultValue           XmlDeregisterNodeFunc
		XmlMallocAtomic                         XmlMallocFunc
		XmlLastError                            XmlError
		XmlParserInputBufferCreateFilenameValue XmlParserInputBufferCreateFilenameFunc
		XmlOutputBufferCreateFilenameValue      XmlOutputBufferCreateFilenameFunc
		XmlStructuredErrorContext               *Void
	}

	XmlEnumeration struct {
		Next *XmlEnumeration
		Name *XmlChar
	}

	XmlNode struct {
		_          *Void
		Type       XmlElementType
		Name       *XmlChar
		Children   *XmlNode
		Last       *XmlNode
		Parent     *XmlNode
		Next       *XmlNode
		Prev       *XmlNode
		Doc        *XmlDoc
		Ns         *XmlNs
		Content    *XmlChar
		Properties *XmlAttr
		NsDef      *XmlNs
		Psvi       *Void
		Line       Unsigned_short
		Extra      Unsigned_short
	}

	XmlDoc struct {
		_           *Void
		Type        XmlElementType
		Name        *Char
		Children    *XmlNode
		Last        *XmlNode
		Parent      *XmlNode
		Next        *XmlNode
		Prev        *XmlNode
		Doc         *XmlDoc
		Compression int
		Standalone  int
		IntSubset   *XmlDtd
		ExtSubset   *XmlDtd
		OldNs       *XmlNs
		Version     *XmlChar
		Encoding    *XmlChar
		Ids         *Void
		Refs        *Void
		URL         *XmlChar
		Charset     int
		Dict        *XmlDict
		Psvi        *Void
		ParseFlags  int
		Properties  int
	}

	XmlEntity struct {
		_          *Void
		Type       XmlElementType
		Name       *XmlChar
		Children   *XmlNode
		Last       *XmlNode
		Parent     *XmlDtd
		Next       *XmlNode
		Prev       *XmlNode
		Doc        *XmlDoc
		Orig       *XmlChar
		Content    *XmlChar
		Length     int
		Etype      XmlEntityType
		ExternalID *XmlChar
		SystemID   *XmlChar
		Nexte      *XmlEntity
		URI        *XmlChar
		Owner      int
		Checked    int
	}

	XmlElementContent struct {
		Type   XmlElementContentType
		Ocur   XmlElementContentOccur
		Name   *XmlChar
		C1     *XmlElementContent
		C2     *XmlElementContent
		Parent *XmlElementContent
		Prefix *XmlChar
	}

	XmlAttr struct {
		_        *Void
		Type     XmlElementType
		Name     *XmlChar
		Children *XmlNode
		Last     *XmlNode
		Parent   *XmlNode
		Next     *XmlAttr
		Prev     *XmlAttr
		Doc      *XmlDoc
		Ns       *XmlNs
		Atype    XmlAttributeType
		Psvi     *Void
	}

	XmlNs struct {
		Next    *XmlNs
		Type    XmlNsType
		Href    *XmlChar
		Prefix  *XmlChar
		_       *Void
		Context *XmlDoc
	}

	XmlNodeSet struct {
		NodeNr  int
		NodeMax int
		NodeTab *XmlNodePtr
	}

	XmlXPathObject struct {
		Type       XmlXPathObjectType
		Nodesetval XmlNodeSetPtr
		Boolval    int
		Floatval   Double
		Stringval  *XmlChar
		User       *Void
		Index      int
		User2      *Void
		Index2     int
	}

	XmlXPathContext struct {
		Doc                  XmlDocPtr
		Node                 XmlNodePtr
		Nb_variables_unused  int
		Max_variables_unused int
		VarHash              XmlHashTablePtr
		Nb_types             int
		Max_types            int
		Types                XmlXPathTypePtr
		Nb_funcs_unused      int
		Max_funcs_unused     int
		FuncHash             XmlHashTablePtr
		Nb_axis              int
		Max_axis             int
		Axis                 XmlXPathAxisPtr
		Namespaces           *XmlNsPtr
		NsNr                 int
		User                 *Void
		ContextSize          int
		ProximityPosition    int
		Xptr                 int
		Here                 XmlNodePtr
		Origin               XmlNodePtr
		NsHash               XmlHashTablePtr
		VarLookupFunc        XmlXPathVariableLookupFunc
		VarLookupData        *Void
		Extra                *Void
		Function             *XmlChar
		FunctionURI          *XmlChar
		FuncLookupFunc       XmlXPathFuncLookupFunc
		FuncLookupData       *Void
		TmpNsList            *XmlNsPtr
		TmpNsNr              int
		UserData             *Void
		Error                XmlStructuredErrorFunc
		LastError            XmlError
		DebugNode            XmlNodePtr
		Dict                 XmlDictPtr
		Flags                int
		Cache                *Void
	}

	XmlXPathType struct {
		Name *XmlChar
		Func XmlXPathConvertFunc
	}

	XmlXPathAxis struct {
		Name *XmlChar
		Func XmlXPathAxisFunc
	}

	XmlLocationSet struct {
		LocNr  int
		LocMax int
		LocTab *XmlXPathObjectPtr
	}

	XmlSchemaWildcard struct {
		Type            XmlSchemaTypeType
		Id              *XmlChar
		Annot           XmlSchemaAnnotPtr
		Node            XmlNodePtr
		MinOccurs       int
		MaxOccurs       int
		ProcessContents int
		Any             int
		NsSet           XmlSchemaWildcardNsPtr
		NegNsSet        XmlSchemaWildcardNsPtr
		Flags           int
	}

	XmlSchemaAnnot struct {
		Next    *XmlSchemaAnnot
		Content XmlNodePtr
	}

	XmlSchemaWildcardNs struct {
		Next  *XmlSchemaWildcardNs
		Value *XmlChar
	}

	XmlChRangeGroup struct {
		NbShortRange int
		NbLongRange  int
		ShortRange   *XmlChSRange
		LongRange    *XmlChLRange
	}

	XmlChSRange struct {
		Low  Unsigned_short
		High Unsigned_short
	}

	XmlChLRange struct {
		Low  Unsigned_int
		High Unsigned_int
	}

	XmlShellCtxt struct {
		Filename *Char
		Doc      XmlDocPtr
		Node     XmlNodePtr
		Pctxt    XmlXPathContextPtr
		Loaded   int
		Output   *FILE
		Input    XmlShellReadlineFunc
	}

	XmlURI struct {
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
		Query_raw *Char
	}
)

type XmlEntityType Enum

const (
	XML_INTERNAL_GENERAL_ENTITY XmlEntityType = iota + 1
	XML_EXTERNAL_GENERAL_PARSED_ENTITY
	XML_EXTERNAL_GENERAL_UNPARSED_ENTITY
	XML_INTERNAL_PARAMETER_ENTITY
	XML_EXTERNAL_PARAMETER_ENTITY
	XML_INTERNAL_PREDEFINED_ENTITY
)

type XmlElementType Enum

const (
	XML_ELEMENT_NODE XmlElementType = iota + 1
	XML_ATTRIBUTE_NODE
	XML_TEXT_NODE
	XML_CDATA_SECTION_NODE
	XML_ENTITY_REF_NODE
	XML_ENTITY_NODE
	XML_PI_NODE
	XML_COMMENT_NODE
	XML_DOCUMENT_NODE
	XML_DOCUMENT_TYPE_NODE
	XML_DOCUMENT_FRAG_NODE
	XML_NOTATION_NODE
	XML_HTML_DOCUMENT_NODE
	XML_DTD_NODE
	XML_ELEMENT_DECL
	XML_ATTRIBUTE_DECL
	XML_ENTITY_DECL
	XML_NAMESPACE_DECL
	XML_XINCLUDE_START
	XML_XINCLUDE_END
	XML_DOCB_DOCUMENT_NODE
)

type XmlBufferAllocationScheme Enum

const (
	XML_BUFFER_ALLOC_DOUBLEIT XmlBufferAllocationScheme = iota
	XML_BUFFER_ALLOC_EXACT
	XML_BUFFER_ALLOC_IMMUTABLE
	XML_BUFFER_ALLOC_IO
	XML_BUFFER_ALLOC_HYBRID
)

type XmlElementContentType Enum

const (
	XML_ELEMENT_CONTENT_PCDATA XmlElementContentType = iota + 1
	XML_ELEMENT_CONTENT_ELEMENT
	XML_ELEMENT_CONTENT_SEQ
	XML_ELEMENT_CONTENT_OR
)

type XmlElementContentOccur Enum

const (
	XML_ELEMENT_CONTENT_ONCE XmlElementContentOccur = iota + 1
	XML_ELEMENT_CONTENT_OPT
	XML_ELEMENT_CONTENT_MULT
	XML_ELEMENT_CONTENT_PLUS
)

type XmlCharEncoding Enum

const (
	XML_CHAR_ENCODING_ERROR XmlCharEncoding = iota - 1
	XML_CHAR_ENCODING_NONE
	XML_CHAR_ENCODING_UTF8
	XML_CHAR_ENCODING_UTF16LE
	XML_CHAR_ENCODING_UTF16BE
	XML_CHAR_ENCODING_UCS4LE
	XML_CHAR_ENCODING_UCS4BE
	XML_CHAR_ENCODING_EBCDIC
	XML_CHAR_ENCODING_UCS4_2143
	XML_CHAR_ENCODING_UCS4_3412
	XML_CHAR_ENCODING_UCS2
	XML_CHAR_ENCODING_8859_1
	XML_CHAR_ENCODING_8859_2
	XML_CHAR_ENCODING_8859_3
	XML_CHAR_ENCODING_8859_4
	XML_CHAR_ENCODING_8859_5
	XML_CHAR_ENCODING_8859_6
	XML_CHAR_ENCODING_8859_7
	XML_CHAR_ENCODING_8859_8
	XML_CHAR_ENCODING_8859_9
	XML_CHAR_ENCODING_2022_JP
	XML_CHAR_ENCODING_SHIFT_JIS
	XML_CHAR_ENCODING_EUC_JP
	XML_CHAR_ENCODING_ASCII
)

type XmlAttributeType Enum

const (
	XML_ATTRIBUTE_CDATA XmlAttributeType = iota + 1
	XML_ATTRIBUTE_ID
	XML_ATTRIBUTE_IDREF
	XML_ATTRIBUTE_IDREFS
	XML_ATTRIBUTE_ENTITY
	XML_ATTRIBUTE_ENTITIES
	XML_ATTRIBUTE_NMTOKEN
	XML_ATTRIBUTE_NMTOKENS
	XML_ATTRIBUTE_ENUMERATION
	XML_ATTRIBUTE_NOTATION
)

type XmlErrorLevel Enum

const (
	XML_ERR_NONE XmlErrorLevel = iota
	XML_ERR_WARNING
	XML_ERR_ERROR
	XML_ERR_FATAL
)

type XmlParserInputState Enum

const (
	XML_PARSER_EOF XmlParserInputState = iota - 1
	XML_PARSER_START
	XML_PARSER_MISC
	XML_PARSER_PI
	XML_PARSER_DTD
	XML_PARSER_PROLOG
	XML_PARSER_COMMENT
	XML_PARSER_START_TAG
	XML_PARSER_CONTENT
	XML_PARSER_CDATA_SECTION
	XML_PARSER_END_TAG
	XML_PARSER_ENTITY_DECL
	XML_PARSER_ENTITY_VALUE
	XML_PARSER_ATTRIBUTE_VALUE
	XML_PARSER_SYSTEM_LITERAL
	XML_PARSER_EPILOG
	XML_PARSER_IGNORE
	XML_PARSER_PUBLIC_LITERAL
)

type XmlParserMode Enum

const (
	XML_PARSE_UNKNOWN XmlParserMode = iota
	XML_PARSE_DOM
	XML_PARSE_SAX
	XML_PARSE_PUSH_DOM
	XML_PARSE_PUSH_SAX
	XML_PARSE_READER
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

type XmlElementTypeVal Enum

const (
	XML_ELEMENT_TYPE_UNDEFINED XmlElementTypeVal = iota
	XML_ELEMENT_TYPE_EMPTY
	XML_ELEMENT_TYPE_ANY
	XML_ELEMENT_TYPE_MIXED
	XML_ELEMENT_TYPE_ELEMENT
)

type XmlAttributeDefault Enum

const (
	XML_ATTRIBUTE_NONE XmlAttributeDefault = iota + 1
	XML_ATTRIBUTE_REQUIRED
	XML_ATTRIBUTE_IMPLIED
	XML_ATTRIBUTE_FIXED
)

type XmlFeature Enum

const (
	XML_WITH_THREAD XmlFeature = iota + 1
	XML_WITH_TREE
	XML_WITH_OUTPUT
	XML_WITH_PUSH
	XML_WITH_READER
	XML_WITH_PATTERN
	XML_WITH_WRITER
	XML_WITH_SAX1
	XML_WITH_FTP
	XML_WITH_HTTP
	XML_WITH_VALID
	XML_WITH_HTML
	XML_WITH_LEGACY
	XML_WITH_C14N
	XML_WITH_CATALOG
	XML_WITH_XPATH
	XML_WITH_XPTR
	XML_WITH_XINCLUDE
	XML_WITH_ICONV
	XML_WITH_ISO8859X
	XML_WITH_UNICODE
	XML_WITH_REGEXP
	XML_WITH_AUTOMATA
	XML_WITH_EXPR
	XML_WITH_SCHEMAS
	XML_WITH_SCHEMATRON
	XML_WITH_MODULES
	XML_WITH_DEBUG
	XML_WITH_DEBUG_MEM
	XML_WITH_DEBUG_RUN
	XML_WITH_ZLIB
	XML_WITH_ICU
	XML_WITH_LZMA
	XML_WITH_NONE XmlFeature = 99999
)

type XmlParserErrors Enum

const (
	XML_ERR_OK XmlParserErrors = iota
	XML_ERR_INTERNAL_ERROR
	XML_ERR_NO_MEMORY
	XML_ERR_DOCUMENT_START
	XML_ERR_DOCUMENT_EMPTY
	XML_ERR_DOCUMENT_END
	XML_ERR_INVALID_HEX_CHARREF
	XML_ERR_INVALID_DEC_CHARREF
	XML_ERR_INVALID_CHARREF
	XML_ERR_INVALID_CHAR
	XML_ERR_CHARREF_AT_EOF
	XML_ERR_CHARREF_IN_PROLOG
	XML_ERR_CHARREF_IN_EPILOG
	XML_ERR_CHARREF_IN_DTD
	XML_ERR_ENTITYREF_AT_EOF
	XML_ERR_ENTITYREF_IN_PROLOG
	XML_ERR_ENTITYREF_IN_EPILOG
	XML_ERR_ENTITYREF_IN_DTD
	XML_ERR_PEREF_AT_EOF
	XML_ERR_PEREF_IN_PROLOG
	XML_ERR_PEREF_IN_EPILOG
	XML_ERR_PEREF_IN_INT_SUBSET
	XML_ERR_ENTITYREF_NO_NAME
	XML_ERR_ENTITYREF_SEMICOL_MISSING
	XML_ERR_PEREF_NO_NAME
	XML_ERR_PEREF_SEMICOL_MISSING
	XML_ERR_UNDECLARED_ENTITY
	XML_WAR_UNDECLARED_ENTITY
	XML_ERR_UNPARSED_ENTITY
	XML_ERR_ENTITY_IS_EXTERNAL
	XML_ERR_ENTITY_IS_PARAMETER
	XML_ERR_UNKNOWN_ENCODING
	XML_ERR_UNSUPPORTED_ENCODING
	XML_ERR_STRING_NOT_STARTED
	XML_ERR_STRING_NOT_CLOSED
	XML_ERR_NS_DECL_ERROR
	XML_ERR_ENTITY_NOT_STARTED
	XML_ERR_ENTITY_NOT_FINISHED
	XML_ERR_LT_IN_ATTRIBUTE
	XML_ERR_ATTRIBUTE_NOT_STARTED
	XML_ERR_ATTRIBUTE_NOT_FINISHED
	XML_ERR_ATTRIBUTE_WITHOUT_VALUE
	XML_ERR_ATTRIBUTE_REDEFINED
	XML_ERR_LITERAL_NOT_STARTED
	XML_ERR_LITERAL_NOT_FINISHED
	XML_ERR_COMMENT_NOT_FINISHED
	XML_ERR_PI_NOT_STARTED
	XML_ERR_PI_NOT_FINISHED
	XML_ERR_NOTATION_NOT_STARTED
	XML_ERR_NOTATION_NOT_FINISHED
	XML_ERR_ATTLIST_NOT_STARTED
	XML_ERR_ATTLIST_NOT_FINISHED
	XML_ERR_MIXED_NOT_STARTED
	XML_ERR_MIXED_NOT_FINISHED
	XML_ERR_ELEMCONTENT_NOT_STARTED
	XML_ERR_ELEMCONTENT_NOT_FINISHED
	XML_ERR_XMLDECL_NOT_STARTED
	XML_ERR_XMLDECL_NOT_FINISHED
	XML_ERR_CONDSEC_NOT_STARTED
	XML_ERR_CONDSEC_NOT_FINISHED
	XML_ERR_EXT_SUBSET_NOT_FINISHED
	XML_ERR_DOCTYPE_NOT_FINISHED
	XML_ERR_MISPLACED_CDATA_END
	XML_ERR_CDATA_NOT_FINISHED
	XML_ERR_RESERVED_XML_NAME
	XML_ERR_SPACE_REQUIRED
	XML_ERR_SEPARATOR_REQUIRED
	XML_ERR_NMTOKEN_REQUIRED
	XML_ERR_NAME_REQUIRED
	XML_ERR_PCDATA_REQUIRED
	XML_ERR_URI_REQUIRED
	XML_ERR_PUBID_REQUIRED
	XML_ERR_LT_REQUIRED
	XML_ERR_GT_REQUIRED
	XML_ERR_LTSLASH_REQUIRED
	XML_ERR_EQUAL_REQUIRED
	XML_ERR_TAG_NAME_MISMATCH
	XML_ERR_TAG_NOT_FINISHED
	XML_ERR_STANDALONE_VALUE
	XML_ERR_ENCODING_NAME
	XML_ERR_HYPHEN_IN_COMMENT
	XML_ERR_INVALID_ENCODING
	XML_ERR_EXT_ENTITY_STANDALONE
	XML_ERR_CONDSEC_INVALID
	XML_ERR_VALUE_REQUIRED
	XML_ERR_NOT_WELL_BALANCED
	XML_ERR_EXTRA_CONTENT
	XML_ERR_ENTITY_CHAR_ERROR
	XML_ERR_ENTITY_PE_INTERNAL
	XML_ERR_ENTITY_LOOP
	XML_ERR_ENTITY_BOUNDARY
	XML_ERR_INVALID_URI
	XML_ERR_URI_FRAGMENT
	XML_WAR_CATALOG_PI
	XML_ERR_NO_DTD
	XML_ERR_CONDSEC_INVALID_KEYWORD
	XML_ERR_VERSION_MISSING
	XML_WAR_UNKNOWN_VERSION
	XML_WAR_LANG_VALUE
	XML_WAR_NS_URI
	XML_WAR_NS_URI_RELATIVE
	XML_ERR_MISSING_ENCODING
	XML_WAR_SPACE_VALUE
	XML_ERR_NOT_STANDALONE
	XML_ERR_ENTITY_PROCESSING
	XML_ERR_NOTATION_PROCESSING
	XML_WAR_NS_COLUMN
	XML_WAR_ENTITY_REDEFINED
	XML_ERR_UNKNOWN_VERSION
	XML_ERR_VERSION_MISMATCH
	XML_ERR_NAME_TOO_LONG
)
const (
	XML_NS_ERR_XML_NAMESPACE XmlParserErrors = iota + 200
	XML_NS_ERR_UNDEFINED_NAMESPACE
	XML_NS_ERR_QNAME
	XML_NS_ERR_ATTRIBUTE_REDEFINED
	XML_NS_ERR_EMPTY
	XML_NS_ERR_COLON
)
const (
	XML_DTD_ATTRIBUTE_DEFAULT XmlParserErrors = iota + 500
	XML_DTD_ATTRIBUTE_REDEFINED
	XML_DTD_ATTRIBUTE_VALUE
	XML_DTD_CONTENT_ERROR
	XML_DTD_CONTENT_MODEL
	XML_DTD_CONTENT_NOT_DETERMINIST
	XML_DTD_DIFFERENT_PREFIX
	XML_DTD_ELEM_DEFAULT_NAMESPACE
	XML_DTD_ELEM_NAMESPACE
	XML_DTD_ELEM_REDEFINED
	XML_DTD_EMPTY_NOTATION
	XML_DTD_ENTITY_TYPE
	XML_DTD_ID_FIXED
	XML_DTD_ID_REDEFINED
	XML_DTD_ID_SUBSET
	XML_DTD_INVALID_CHILD
	XML_DTD_INVALID_DEFAULT
	XML_DTD_LOAD_ERROR
	XML_DTD_MISSING_ATTRIBUTE
	XML_DTD_MIXED_CORRUPT
	XML_DTD_MULTIPLE_ID
	XML_DTD_NO_DOC
	XML_DTD_NO_DTD
	XML_DTD_NO_ELEM_NAME
	XML_DTD_NO_PREFIX
	XML_DTD_NO_ROOT
	XML_DTD_NOTATION_REDEFINED
	XML_DTD_NOTATION_VALUE
	XML_DTD_NOT_EMPTY
	XML_DTD_NOT_PCDATA
	XML_DTD_NOT_STANDALONE
	XML_DTD_ROOT_NAME
	XML_DTD_STANDALONE_WHITE_SPACE
	XML_DTD_UNKNOWN_ATTRIBUTE
	XML_DTD_UNKNOWN_ELEM
	XML_DTD_UNKNOWN_ENTITY
	XML_DTD_UNKNOWN_ID
	XML_DTD_UNKNOWN_NOTATION
	XML_DTD_STANDALONE_DEFAULTED
	XML_DTD_XMLID_VALUE
	XML_DTD_XMLID_TYPE
	XML_DTD_DUP_TOKEN
)
const (
	XML_HTML_STRUCURE_ERROR XmlParserErrors = iota + 800
	XML_HTML_UNKNOWN_TAG
)
const (
	XML_RNGP_ANYNAME_ATTR_ANCESTOR XmlParserErrors = iota + 1000
	XML_RNGP_ATTR_CONFLICT
	XML_RNGP_ATTRIBUTE_CHILDREN
	XML_RNGP_ATTRIBUTE_CONTENT
	XML_RNGP_ATTRIBUTE_EMPTY
	XML_RNGP_ATTRIBUTE_NOOP
	XML_RNGP_CHOICE_CONTENT
	XML_RNGP_CHOICE_EMPTY
	XML_RNGP_CREATE_FAILURE
	XML_RNGP_DATA_CONTENT
	XML_RNGP_DEF_CHOICE_AND_INTERLEAVE
	XML_RNGP_DEFINE_CREATE_FAILED
	XML_RNGP_DEFINE_EMPTY
	XML_RNGP_DEFINE_MISSING
	XML_RNGP_DEFINE_NAME_MISSING
	XML_RNGP_ELEM_CONTENT_EMPTY
	XML_RNGP_ELEM_CONTENT_ERROR
	XML_RNGP_ELEMENT_EMPTY
	XML_RNGP_ELEMENT_CONTENT
	XML_RNGP_ELEMENT_NAME
	XML_RNGP_ELEMENT_NO_CONTENT
	XML_RNGP_ELEM_TEXT_CONFLICT
	XML_RNGP_EMPTY
	XML_RNGP_EMPTY_CONSTRUCT
	XML_RNGP_EMPTY_CONTENT
	XML_RNGP_EMPTY_NOT_EMPTY
	XML_RNGP_ERROR_TYPE_LIB
	XML_RNGP_EXCEPT_EMPTY
	XML_RNGP_EXCEPT_MISSING
	XML_RNGP_EXCEPT_MULTIPLE
	XML_RNGP_EXCEPT_NO_CONTENT
	XML_RNGP_EXTERNALREF_EMTPY
	XML_RNGP_EXTERNAL_REF_FAILURE
	XML_RNGP_EXTERNALREF_RECURSE
	XML_RNGP_FORBIDDEN_ATTRIBUTE
	XML_RNGP_FOREIGN_ELEMENT
	XML_RNGP_GRAMMAR_CONTENT
	XML_RNGP_GRAMMAR_EMPTY
	XML_RNGP_GRAMMAR_MISSING
	XML_RNGP_GRAMMAR_NO_START
	XML_RNGP_GROUP_ATTR_CONFLICT
	XML_RNGP_HREF_ERROR
	XML_RNGP_INCLUDE_EMPTY
	XML_RNGP_INCLUDE_FAILURE
	XML_RNGP_INCLUDE_RECURSE
	XML_RNGP_INTERLEAVE_ADD
	XML_RNGP_INTERLEAVE_CREATE_FAILED
	XML_RNGP_INTERLEAVE_EMPTY
	XML_RNGP_INTERLEAVE_NO_CONTENT
	XML_RNGP_INVALID_DEFINE_NAME
	XML_RNGP_INVALID_URI
	XML_RNGP_INVALID_VALUE
	XML_RNGP_MISSING_HREF
	XML_RNGP_NAME_MISSING
	XML_RNGP_NEED_COMBINE
	XML_RNGP_NOTALLOWED_NOT_EMPTY
	XML_RNGP_NSNAME_ATTR_ANCESTOR
	XML_RNGP_NSNAME_NO_NS
	XML_RNGP_PARAM_FORBIDDEN
	XML_RNGP_PARAM_NAME_MISSING
	XML_RNGP_PARENTREF_CREATE_FAILED
	XML_RNGP_PARENTREF_NAME_INVALID
	XML_RNGP_PARENTREF_NO_NAME
	XML_RNGP_PARENTREF_NO_PARENT
	XML_RNGP_PARENTREF_NOT_EMPTY
	XML_RNGP_PARSE_ERROR
	XML_RNGP_PAT_ANYNAME_EXCEPT_ANYNAME
	XML_RNGP_PAT_ATTR_ATTR
	XML_RNGP_PAT_ATTR_ELEM
	XML_RNGP_PAT_DATA_EXCEPT_ATTR
	XML_RNGP_PAT_DATA_EXCEPT_ELEM
	XML_RNGP_PAT_DATA_EXCEPT_EMPTY
	XML_RNGP_PAT_DATA_EXCEPT_GROUP
	XML_RNGP_PAT_DATA_EXCEPT_INTERLEAVE
	XML_RNGP_PAT_DATA_EXCEPT_LIST
	XML_RNGP_PAT_DATA_EXCEPT_ONEMORE
	XML_RNGP_PAT_DATA_EXCEPT_REF
	XML_RNGP_PAT_DATA_EXCEPT_TEXT
	XML_RNGP_PAT_LIST_ATTR
	XML_RNGP_PAT_LIST_ELEM
	XML_RNGP_PAT_LIST_INTERLEAVE
	XML_RNGP_PAT_LIST_LIST
	XML_RNGP_PAT_LIST_REF
	XML_RNGP_PAT_LIST_TEXT
	XML_RNGP_PAT_NSNAME_EXCEPT_ANYNAME
	XML_RNGP_PAT_NSNAME_EXCEPT_NSNAME
	XML_RNGP_PAT_ONEMORE_GROUP_ATTR
	XML_RNGP_PAT_ONEMORE_INTERLEAVE_ATTR
	XML_RNGP_PAT_START_ATTR
	XML_RNGP_PAT_START_DATA
	XML_RNGP_PAT_START_EMPTY
	XML_RNGP_PAT_START_GROUP
	XML_RNGP_PAT_START_INTERLEAVE
	XML_RNGP_PAT_START_LIST
	XML_RNGP_PAT_START_ONEMORE
	XML_RNGP_PAT_START_TEXT
	XML_RNGP_PAT_START_VALUE
	XML_RNGP_PREFIX_UNDEFINED
	XML_RNGP_REF_CREATE_FAILED
	XML_RNGP_REF_CYCLE
	XML_RNGP_REF_NAME_INVALID
	XML_RNGP_REF_NO_DEF
	XML_RNGP_REF_NO_NAME
	XML_RNGP_REF_NOT_EMPTY
	XML_RNGP_START_CHOICE_AND_INTERLEAVE
	XML_RNGP_START_CONTENT
	XML_RNGP_START_EMPTY
	XML_RNGP_START_MISSING
	XML_RNGP_TEXT_EXPECTED
	XML_RNGP_TEXT_HAS_CHILD
	XML_RNGP_TYPE_MISSING
	XML_RNGP_TYPE_NOT_FOUND
	XML_RNGP_TYPE_VALUE
	XML_RNGP_UNKNOWN_ATTRIBUTE
	XML_RNGP_UNKNOWN_COMBINE
	XML_RNGP_UNKNOWN_CONSTRUCT
	XML_RNGP_UNKNOWN_TYPE_LIB
	XML_RNGP_URI_FRAGMENT
	XML_RNGP_URI_NOT_ABSOLUTE
	XML_RNGP_VALUE_EMPTY
	XML_RNGP_VALUE_NO_CONTENT
	XML_RNGP_XmlNs_NAME
	XML_RNGP_XML_NS
)
const (
	XML_XPATH_EXPRESSION_OK XmlParserErrors = iota + 1200
	XML_XPATH_NUMBER_ERROR
	XML_XPATH_UNFINISHED_LITERAL_ERROR
	XML_XPATH_START_LITERAL_ERROR
	XML_XPATH_VARIABLE_REF_ERROR
	XML_XPATH_UNDEF_VARIABLE_ERROR
	XML_XPATH_INVALID_PREDICATE_ERROR
	XML_XPATH_EXPR_ERROR
	XML_XPATH_UNCLOSED_ERROR
	XML_XPATH_UNKNOWN_FUNC_ERROR
	XML_XPATH_INVALID_OPERAND
	XML_XPATH_INVALID_TYPE
	XML_XPATH_INVALID_ARITY
	XML_XPATH_INVALID_CTXT_SIZE
	XML_XPATH_INVALID_CTXT_POSITION
	XML_XPATH_MEMORY_ERROR
	XML_XPTR_SYNTAX_ERROR
	XML_XPTR_RESOURCE_ERROR
	XML_XPTR_SUB_RESOURCE_ERROR
	XML_XPATH_UNDEF_PREFIX_ERROR
	XML_XPATH_ENCODING_ERROR
	XML_XPATH_INVALID_CHAR_ERROR
)
const (
	XML_TREE_INVALID_HEX XmlParserErrors = iota + 1300
	XML_TREE_INVALID_DEC
	XML_TREE_UNTERMINATED_ENTITY
	XML_TREE_NOT_UTF8
)
const (
	XML_SAVE_NOT_UTF8 XmlParserErrors = iota + 1400
	XML_SAVE_CHAR_INVALID
	XML_SAVE_NO_DOCTYPE
	XML_SAVE_UNKNOWN_ENCODING
	XML_REGEXP_COMPILE_ERROR XmlParserErrors = 1450
)
const (
	XML_IO_UNKNOWN XmlParserErrors = iota + 1500
	XML_IO_EACCES
	XML_IO_EAGAIN
	XML_IO_EBADF
	XML_IO_EBADMSG
	XML_IO_EBUSY
	XML_IO_ECANCELED
	XML_IO_ECHILD
	XML_IO_EDEADLK
	XML_IO_EDOM
	XML_IO_EEXIST
	XML_IO_EFAULT
	XML_IO_EFBIG
	XML_IO_EINPROGRESS
	XML_IO_EINTR
	XML_IO_EINVAL
	XML_IO_EIO
	XML_IO_EISDIR
	XML_IO_EMFILE
	XML_IO_EMLINK
	XML_IO_EMSGSIZE
	XML_IO_ENAMETOOLONG
	XML_IO_ENFILE
	XML_IO_ENODEV
	XML_IO_ENOENT
	XML_IO_ENOEXEC
	XML_IO_ENOLCK
	XML_IO_ENOMEM
	XML_IO_ENOSPC
	XML_IO_ENOSYS
	XML_IO_ENOTDIR
	XML_IO_ENOTEMPTY
	XML_IO_ENOTSUP
	XML_IO_ENOTTY
	XML_IO_ENXIO
	XML_IO_EPERM
	XML_IO_EPIPE
	XML_IO_ERANGE
	XML_IO_EROFS
	XML_IO_ESPIPE
	XML_IO_ESRCH
	XML_IO_ETIMEDOUT
	XML_IO_EXDEV
	XML_IO_NETWORK_ATTEMPT
	XML_IO_ENCODER
	XML_IO_FLUSH
	XML_IO_WRITE
	XML_IO_NO_INPUT
	XML_IO_BUFFER_FULL
	XML_IO_LOAD_ERROR
	XML_IO_ENOTSOCK
	XML_IO_EISCONN
	XML_IO_ECONNREFUSED
	XML_IO_ENETUNREACH
	XML_IO_EADDRINUSE
	XML_IO_EALREADY
	XML_IO_EAFNOSUPPORT
)
const (
	XML_XINCLUDE_RECURSION XmlParserErrors = iota + 1600
	XML_XINCLUDE_PARSE_VALUE
	XML_XINCLUDE_ENTITY_DEF_MISMATCH
	XML_XINCLUDE_NO_HREF
	XML_XINCLUDE_NO_FALLBACK
	XML_XINCLUDE_HREF_URI
	XML_XINCLUDE_TEXT_FRAGMENT
	XML_XINCLUDE_TEXT_DOCUMENT
	XML_XINCLUDE_INVALID_CHAR
	XML_XINCLUDE_BUILD_FAILED
	XML_XINCLUDE_UNKNOWN_ENCODING
	XML_XINCLUDE_MULTIPLE_ROOT
	XML_XINCLUDE_XPTR_FAILED
	XML_XINCLUDE_XPTR_RESULT
	XML_XINCLUDE_INCLUDE_IN_INCLUDE
	XML_XINCLUDE_FALLBACKS_IN_INCLUDE
	XML_XINCLUDE_FALLBACK_NOT_IN_INCLUDE
	XML_XINCLUDE_DEPRECATED_NS
	XML_XINCLUDE_FRAGMENT_ID
)
const (
	XML_CATALOG_MISSING_ATTR XmlParserErrors = iota + 1650
	XML_CATALOG_ENTRY_BROKEN
	XML_CATALOG_PREFER_VALUE
	XML_CATALOG_NOT_CATALOG
	XML_CATALOG_RECURSION
)
const (
	XML_SCHEMAP_PREFIX_UNDEFINED XmlParserErrors = iota + 1700
	XML_SCHEMAP_ATTRFORMDEFAULT_VALUE
	XML_SCHEMAP_ATTRGRP_NONAME_NOREF
	XML_SCHEMAP_ATTR_NONAME_NOREF
	XML_SCHEMAP_COMPLEXTYPE_NONAME_NOREF
	XML_SCHEMAP_ELEMFORMDEFAULT_VALUE
	XML_SCHEMAP_ELEM_NONAME_NOREF
	XML_SCHEMAP_EXTENSION_NO_BASE
	XML_SCHEMAP_FACET_NO_VALUE
	XML_SCHEMAP_FAILED_BUILD_IMPORT
	XML_SCHEMAP_GROUP_NONAME_NOREF
	XML_SCHEMAP_IMPORT_NAMESPACE_NOT_URI
	XML_SCHEMAP_IMPORT_REDEFINE_NSNAME
	XML_SCHEMAP_IMPORT_SCHEMA_NOT_URI
	XML_SCHEMAP_INVALID_BOOLEAN
	XML_SCHEMAP_INVALID_ENUM
	XML_SCHEMAP_INVALID_FACET
	XML_SCHEMAP_INVALID_FACET_VALUE
	XML_SCHEMAP_INVALID_MAXOCCURS
	XML_SCHEMAP_INVALID_MINOCCURS
	XML_SCHEMAP_INVALID_REF_AND_SUBTYPE
	XML_SCHEMAP_INVALID_WHITE_SPACE
	XML_SCHEMAP_NOATTR_NOREF
	XML_SCHEMAP_NOTATION_NO_NAME
	XML_SCHEMAP_NOTYPE_NOREF
	XML_SCHEMAP_REF_AND_SUBTYPE
	XML_SCHEMAP_RESTRICTION_NONAME_NOREF
	XML_SCHEMAP_SIMPLETYPE_NONAME
	XML_SCHEMAP_TYPE_AND_SUBTYPE
	XML_SCHEMAP_UNKNOWN_ALL_CHILD
	XML_SCHEMAP_UNKNOWN_ANYATTRIBUTE_CHILD
	XML_SCHEMAP_UNKNOWN_ATTR_CHILD
	XML_SCHEMAP_UNKNOWN_ATTRGRP_CHILD
	XML_SCHEMAP_UNKNOWN_ATTRIBUTE_GROUP
	XML_SCHEMAP_UNKNOWN_BASE_TYPE
	XML_SCHEMAP_UNKNOWN_CHOICE_CHILD
	XML_SCHEMAP_UNKNOWN_COMPLEXCONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_COMPLEXTYPE_CHILD
	XML_SCHEMAP_UNKNOWN_ELEM_CHILD
	XML_SCHEMAP_UNKNOWN_EXTENSION_CHILD
	XML_SCHEMAP_UNKNOWN_FACET_CHILD
	XML_SCHEMAP_UNKNOWN_FACET_TYPE
	XML_SCHEMAP_UNKNOWN_GROUP_CHILD
	XML_SCHEMAP_UNKNOWN_IMPORT_CHILD
	XML_SCHEMAP_UNKNOWN_LIST_CHILD
	XML_SCHEMAP_UNKNOWN_NOTATION_CHILD
	XML_SCHEMAP_UNKNOWN_PROCESSCONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_REF
	XML_SCHEMAP_UNKNOWN_RESTRICTION_CHILD
	XML_SCHEMAP_UNKNOWN_SCHEMAS_CHILD
	XML_SCHEMAP_UNKNOWN_SEQUENCE_CHILD
	XML_SCHEMAP_UNKNOWN_SIMPLECONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_SIMPLETYPE_CHILD
	XML_SCHEMAP_UNKNOWN_TYPE
	XML_SCHEMAP_UNKNOWN_UNION_CHILD
	XML_SCHEMAP_ELEM_DEFAULT_FIXED
	XML_SCHEMAP_REGEXP_INVALID
	XML_SCHEMAP_FAILED_LOAD
	XML_SCHEMAP_NOTHING_TO_PARSE
	XML_SCHEMAP_NOROOT
	XML_SCHEMAP_REDEFINED_GROUP
	XML_SCHEMAP_REDEFINED_TYPE
	XML_SCHEMAP_REDEFINED_ELEMENT
	XML_SCHEMAP_REDEFINED_ATTRGROUP
	XML_SCHEMAP_REDEFINED_ATTR
	XML_SCHEMAP_REDEFINED_NOTATION
	XML_SCHEMAP_FAILED_PARSE
	XML_SCHEMAP_UNKNOWN_PREFIX
	XML_SCHEMAP_DEF_AND_PREFIX
	XML_SCHEMAP_UNKNOWN_INCLUDE_CHILD
	XML_SCHEMAP_INCLUDE_SCHEMA_NOT_URI
	XML_SCHEMAP_INCLUDE_SCHEMA_NO_URI
	XML_SCHEMAP_NOT_SCHEMA
	XML_SCHEMAP_UNKNOWN_MEMBER_TYPE
	XML_SCHEMAP_INVALID_ATTR_USE
	XML_SCHEMAP_RECURSIVE
	XML_SCHEMAP_SUPERNUMEROUS_LIST_ITEM_TYPE
	XML_SCHEMAP_INVALID_ATTR_COMBINATION
	XML_SCHEMAP_INVALID_ATTR_INLINE_COMBINATION
	XML_SCHEMAP_MISSING_SIMPLETYPE_CHILD
	XML_SCHEMAP_INVALID_ATTR_NAME
	XML_SCHEMAP_REF_AND_CONTENT
	XML_SCHEMAP_CT_PROPS_CORRECT_1
	XML_SCHEMAP_CT_PROPS_CORRECT_2
	XML_SCHEMAP_CT_PROPS_CORRECT_3
	XML_SCHEMAP_CT_PROPS_CORRECT_4
	XML_SCHEMAP_CT_PROPS_CORRECT_5
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_3
	XML_SCHEMAP_WILDCARD_INVALID_NS_MEMBER
	XML_SCHEMAP_INTERSECTION_NOT_EXPRESSIBLE
	XML_SCHEMAP_UNION_NOT_EXPRESSIBLE
	XML_SCHEMAP_SRC_IMPORT_3_1
	XML_SCHEMAP_SRC_IMPORT_3_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_3
	XML_SCHEMAP_COS_CT_EXTENDS_1_3
)
const (
	XML_SCHEMAV_NOROOT XmlParserErrors = iota + 1801
	XML_SCHEMAV_UNDECLAREDELEM
	XML_SCHEMAV_NOTTOPLEVEL
	XML_SCHEMAV_MISSING
	XML_SCHEMAV_WRONGELEM
	XML_SCHEMAV_NOTYPE
	XML_SCHEMAV_NOROLLBACK
	XML_SCHEMAV_ISABSTRACT
	XML_SCHEMAV_NOTEMPTY
	XML_SCHEMAV_ELEMCONT
	XML_SCHEMAV_HAVEDEFAULT
	XML_SCHEMAV_NOTNILLABLE
	XML_SCHEMAV_EXTRACONTENT
	XML_SCHEMAV_INVALIDATTR
	XML_SCHEMAV_INVALIDELEM
	XML_SCHEMAV_NOTDETERMINIST
	XML_SCHEMAV_CONSTRUCT
	XML_SCHEMAV_INTERNAL
	XML_SCHEMAV_NOTSIMPLE
	XML_SCHEMAV_ATTRUNKNOWN
	XML_SCHEMAV_ATTRINVALID
	XML_SCHEMAV_VALUE
	XML_SCHEMAV_FACET
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_1
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_2
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_3
	XML_SCHEMAV_CVC_TYPE_3_1_1
	XML_SCHEMAV_CVC_TYPE_3_1_2
	XML_SCHEMAV_CVC_FACET_VALID
	XML_SCHEMAV_CVC_LENGTH_VALID
	XML_SCHEMAV_CVC_MINLENGTH_VALID
	XML_SCHEMAV_CVC_MAXLENGTH_VALID
	XML_SCHEMAV_CVC_MININCLUSIVE_VALID
	XML_SCHEMAV_CVC_MAXINCLUSIVE_VALID
	XML_SCHEMAV_CVC_MINEXCLUSIVE_VALID
	XML_SCHEMAV_CVC_MAXEXCLUSIVE_VALID
	XML_SCHEMAV_CVC_TOTALDIGITS_VALID
	XML_SCHEMAV_CVC_FRACTIONDIGITS_VALID
	XML_SCHEMAV_CVC_PATTERN_VALID
	XML_SCHEMAV_CVC_ENUMERATION_VALID
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_2
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_3
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_4
	XML_SCHEMAV_CVC_ELT_1
	XML_SCHEMAV_CVC_ELT_2
	XML_SCHEMAV_CVC_ELT_3_1
	XML_SCHEMAV_CVC_ELT_3_2_1
	XML_SCHEMAV_CVC_ELT_3_2_2
	XML_SCHEMAV_CVC_ELT_4_1
	XML_SCHEMAV_CVC_ELT_4_2
	XML_SCHEMAV_CVC_ELT_4_3
	XML_SCHEMAV_CVC_ELT_5_1_1
	XML_SCHEMAV_CVC_ELT_5_1_2
	XML_SCHEMAV_CVC_ELT_5_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_2_2
	XML_SCHEMAV_CVC_ELT_6
	XML_SCHEMAV_CVC_ELT_7
	XML_SCHEMAV_CVC_ATTRIBUTE_1
	XML_SCHEMAV_CVC_ATTRIBUTE_2
	XML_SCHEMAV_CVC_ATTRIBUTE_3
	XML_SCHEMAV_CVC_ATTRIBUTE_4
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_2_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_2_2
	XML_SCHEMAV_CVC_COMPLEX_TYPE_4
	XML_SCHEMAV_CVC_COMPLEX_TYPE_5_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_5_2
	XML_SCHEMAV_ELEMENT_CONTENT
	XML_SCHEMAV_DOCUMENT_ELEMENT_MISSING
	XML_SCHEMAV_CVC_COMPLEX_TYPE_1
	XML_SCHEMAV_CVC_AU
	XML_SCHEMAV_CVC_TYPE_1
	XML_SCHEMAV_CVC_TYPE_2
	XML_SCHEMAV_CVC_IDC
	XML_SCHEMAV_CVC_WILDCARD
	XML_SCHEMAV_MISC
)
const (
	XML_XPTR_UNKNOWN_SCHEME XmlParserErrors = iota + 1900
	XML_XPTR_CHILDSEQ_START
	XML_XPTR_EVAL_FAILED
	XML_XPTR_EXTRA_OBJECTS
)
const (
	XML_C14N_CREATE_CTXT XmlParserErrors = iota + 1950
	XML_C14N_REQUIRES_UTF8
	XML_C14N_CREATE_STACK
	XML_C14N_INVALID_NODE
	XML_C14N_UNKNOW_NODE
	XML_C14N_RELATIVE_NAMESPACE
)
const (
	XML_FTP_PASV_ANSWER XmlParserErrors = iota + 2000
	XML_FTP_EPSV_ANSWER
	XML_FTP_ACCNT
	XML_FTP_URL_SYNTAX
)
const (
	XML_HTTP_URL_SYNTAX XmlParserErrors = iota + 2020
	XML_HTTP_USE_IP
	XML_HTTP_UNKNOWN_HOST
)
const (
	XML_SCHEMAP_SRC_SIMPLE_TYPE_1 XmlParserErrors = iota + 3000
	XML_SCHEMAP_SRC_SIMPLE_TYPE_2
	XML_SCHEMAP_SRC_SIMPLE_TYPE_3
	XML_SCHEMAP_SRC_SIMPLE_TYPE_4
	XML_SCHEMAP_SRC_RESOLVE
	XML_SCHEMAP_SRC_RESTRICTION_BASE_OR_SIMPLETYPE
	XML_SCHEMAP_SRC_LIST_ITEMTYPE_OR_SIMPLETYPE
	XML_SCHEMAP_SRC_UNION_MEMBERTYPES_OR_SIMPLETYPES
	XML_SCHEMAP_ST_PROPS_CORRECT_1
	XML_SCHEMAP_ST_PROPS_CORRECT_2
	XML_SCHEMAP_ST_PROPS_CORRECT_3
	XML_SCHEMAP_COS_ST_RESTRICTS_1_1
	XML_SCHEMAP_COS_ST_RESTRICTS_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_1_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_1_3_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_1_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_3
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_4
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_5
	XML_SCHEMAP_COS_ST_RESTRICTS_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_2
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_3
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_4
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_5
	XML_SCHEMAP_COS_ST_DERIVED_OK_2_1
	XML_SCHEMAP_COS_ST_DERIVED_OK_2_2
	XML_SCHEMAP_S4S_ELEM_NOT_ALLOWED
	XML_SCHEMAP_S4S_ELEM_MISSING
	XML_SCHEMAP_S4S_ATTR_NOT_ALLOWED
	XML_SCHEMAP_S4S_ATTR_MISSING
	XML_SCHEMAP_S4S_ATTR_INVALID_VALUE
	XML_SCHEMAP_SRC_ELEMENT_1
	XML_SCHEMAP_SRC_ELEMENT_2_1
	XML_SCHEMAP_SRC_ELEMENT_2_2
	XML_SCHEMAP_SRC_ELEMENT_3
	XML_SCHEMAP_P_PROPS_CORRECT_1
	XML_SCHEMAP_P_PROPS_CORRECT_2_1
	XML_SCHEMAP_P_PROPS_CORRECT_2_2
	XML_SCHEMAP_E_PROPS_CORRECT_2
	XML_SCHEMAP_E_PROPS_CORRECT_3
	XML_SCHEMAP_E_PROPS_CORRECT_4
	XML_SCHEMAP_E_PROPS_CORRECT_5
	XML_SCHEMAP_E_PROPS_CORRECT_6
	XML_SCHEMAP_SRC_INCLUDE
	XML_SCHEMAP_SRC_ATTRIBUTE_1
	XML_SCHEMAP_SRC_ATTRIBUTE_2
	XML_SCHEMAP_SRC_ATTRIBUTE_3_1
	XML_SCHEMAP_SRC_ATTRIBUTE_3_2
	XML_SCHEMAP_SRC_ATTRIBUTE_4
	XML_SCHEMAP_NO_XmlNs
	XML_SCHEMAP_NO_XSI
	XML_SCHEMAP_COS_VALID_DEFAULT_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_2_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_2_2
	XML_SCHEMAP_CVC_SIMPLE_TYPE
	XML_SCHEMAP_COS_CT_EXTENDS_1_1
	XML_SCHEMAP_SRC_IMPORT_1_1
	XML_SCHEMAP_SRC_IMPORT_1_2
	XML_SCHEMAP_SRC_IMPORT_2
	XML_SCHEMAP_SRC_IMPORT_2_1
	XML_SCHEMAP_SRC_IMPORT_2_2
	XML_SCHEMAP_INTERNAL
	XML_SCHEMAP_NOT_DETERMINISTIC
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_1
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_2
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_3
	XML_SCHEMAP_MG_PROPS_CORRECT_1
	XML_SCHEMAP_MG_PROPS_CORRECT_2
	XML_SCHEMAP_SRC_CT_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_3
	XML_SCHEMAP_AU_PROPS_CORRECT_2
	XML_SCHEMAP_A_PROPS_CORRECT_2
	XML_SCHEMAP_C_PROPS_CORRECT
	XML_SCHEMAP_SRC_REDEFINE
	XML_SCHEMAP_SRC_IMPORT
	XML_SCHEMAP_WARN_SKIP_SCHEMA
	XML_SCHEMAP_WARN_UNLOCATED_SCHEMA
	XML_SCHEMAP_WARN_ATTR_REDECL_PROH
	XML_SCHEMAP_WARN_ATTR_POINTLESS_PROH
	XML_SCHEMAP_AG_PROPS_CORRECT
	XML_SCHEMAP_COS_CT_EXTENDS_1_2
	XML_SCHEMAP_AU_PROPS_CORRECT
	XML_SCHEMAP_A_PROPS_CORRECT_3
	XML_SCHEMAP_COS_ALL_LIMITED
)
const (
	XML_SCHEMATRONV_ASSERT XmlParserErrors = iota + 4000
	XML_SCHEMATRONV_REPORT
)
const (
	XML_MODULE_OPEN XmlParserErrors = iota + 4900
	XML_MODULE_CLOSE
)
const (
	XML_CHECK_FOUND_ELEMENT XmlParserErrors = iota + 5000
	XML_CHECK_FOUND_ATTRIBUTE
	XML_CHECK_FOUND_TEXT
	XML_CHECK_FOUND_CDATA
	XML_CHECK_FOUND_ENTITYREF
	XML_CHECK_FOUND_ENTITY
	XML_CHECK_FOUND_PI
	XML_CHECK_FOUND_COMMENT
	XML_CHECK_FOUND_DOCTYPE
	XML_CHECK_FOUND_FRAGMENT
	XML_CHECK_FOUND_NOTATION
	XML_CHECK_UNKNOWN_NODE
	XML_CHECK_ENTITY_TYPE
	XML_CHECK_NO_PARENT
	XML_CHECK_NO_DOC
	XML_CHECK_NO_NAME
	XML_CHECK_NO_ELEM
	XML_CHECK_WRONG_DOC
	XML_CHECK_NO_PREV
	XML_CHECK_WRONG_PREV
	XML_CHECK_NO_NEXT
	XML_CHECK_WRONG_NEXT
	XML_CHECK_NOT_DTD
	XML_CHECK_NOT_ATTR
	XML_CHECK_NOT_ATTR_DECL
	XML_CHECK_NOT_ELEM_DECL
	XML_CHECK_NOT_ENTITY_DECL
	XML_CHECK_NOT_NS_DECL
	XML_CHECK_NO_HREF
	XML_CHECK_WRONG_PARENT
	XML_CHECK_NS_SCOPE
	XML_CHECK_NS_ANCESTOR
	XML_CHECK_NOT_UTF8
	XML_CHECK_NO_DICT
	XML_CHECK_NOT_NCNAME
	XML_CHECK_OUTSIDE_DICT
	XML_CHECK_WRONG_NAME
	XML_CHECK_NAME_NOT_NULL
)
const (
	XML_I18N_NO_NAME XmlParserErrors = iota + 6000
	XML_I18N_NO_HANDLER
	XML_I18N_EXCESS_HANDLER
	XML_I18N_CONV_FAILED
	XML_I18N_NO_OUTPUT
	XML_BUF_OVERFLOW XmlParserErrors = 7000
)

type XmlXPathObjectType Enum

const (
	XPATH_UNDEFINED XmlXPathObjectType = iota
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

type XmlSchemaTypeType Enum

const (
	XML_SCHEMA_TYPE_BASIC XmlSchemaTypeType = iota + 1
	XML_SCHEMA_TYPE_ANY
	XML_SCHEMA_TYPE_FACET
	XML_SCHEMA_TYPE_SIMPLE
	XML_SCHEMA_TYPE_COMPLEX
	XML_SCHEMA_TYPE_SEQUENCE
	XML_SCHEMA_TYPE_CHOICE
	XML_SCHEMA_TYPE_ALL
	XML_SCHEMA_TYPE_SIMPLE_CONTENT
	XML_SCHEMA_TYPE_COMPLEX_CONTENT
	XML_SCHEMA_TYPE_UR
	XML_SCHEMA_TYPE_RESTRICTION
	XML_SCHEMA_TYPE_EXTENSION
	XML_SCHEMA_TYPE_ELEMENT
	XML_SCHEMA_TYPE_ATTRIBUTE
	XML_SCHEMA_TYPE_ATTRIBUTEGROUP
	XML_SCHEMA_TYPE_GROUP
	XML_SCHEMA_TYPE_NOTATION
	XML_SCHEMA_TYPE_LIST
	XML_SCHEMA_TYPE_UNION
	XML_SCHEMA_TYPE_ANY_ATTRIBUTE
	XML_SCHEMA_TYPE_IDC_UNIQUE
	XML_SCHEMA_TYPE_IDC_KEY
	XML_SCHEMA_TYPE_IDC_KEYREF
	XML_SCHEMA_TYPE_PARTICLE
	XML_SCHEMA_TYPE_ATTRIBUTE_USE
)
const (
	XML_SCHEMA_FACET_MININCLUSIVE XmlSchemaTypeType = 1000
	XML_SCHEMA_FACET_MINEXCLUSIVE
	XML_SCHEMA_FACET_MAXINCLUSIVE
	XML_SCHEMA_FACET_MAXEXCLUSIVE
	XML_SCHEMA_FACET_TOTALDIGITS
	XML_SCHEMA_FACET_FRACTIONDIGITS
	XML_SCHEMA_FACET_PATTERN
	XML_SCHEMA_FACET_ENUMERATION
	XML_SCHEMA_FACET_WHITESPACE
	XML_SCHEMA_FACET_LENGTH
	XML_SCHEMA_FACET_MAXLENGTH
	XML_SCHEMA_FACET_MINLENGTH
)
const (
	XML_SCHEMA_EXTRA_QNAMEREF XmlSchemaTypeType = 2000
	XML_SCHEMA_EXTRA_ATTR_USE_PROHIB
)

type XmlSchemaWhitespaceValueType Enum

const (
	XML_SCHEMA_WHITESPACE_UNKNOWN XmlSchemaWhitespaceValueType = iota
	XML_SCHEMA_WHITESPACE_PRESERVE
	XML_SCHEMA_WHITESPACE_REPLACE
	XML_SCHEMA_WHITESPACE_COLLAPSE
)

type XmlSchemaValType Enum

const (
	XML_SCHEMAS_UNKNOWN XmlSchemaValType = iota
	XML_SCHEMAS_STRING
	XML_SCHEMAS_NORMSTRING
	XML_SCHEMAS_DECIMAL
	XML_SCHEMAS_TIME
	XML_SCHEMAS_GDAY
	XML_SCHEMAS_GMONTH
	XML_SCHEMAS_GMONTHDAY
	XML_SCHEMAS_GYEAR
	XML_SCHEMAS_GYEARMONTH
	XML_SCHEMAS_DATE
	XML_SCHEMAS_DATETIME
	XML_SCHEMAS_DURATION
	XML_SCHEMAS_FLOAT
	XML_SCHEMAS_DOUBLE
	XML_SCHEMAS_BOOLEAN
	XML_SCHEMAS_TOKEN
	XML_SCHEMAS_LANGUAGE
	XML_SCHEMAS_NMTOKEN
	XML_SCHEMAS_NMTOKENS
	XML_SCHEMAS_NAME
	XML_SCHEMAS_QNAME
	XML_SCHEMAS_NCNAME
	XML_SCHEMAS_ID
	XML_SCHEMAS_IDREF
	XML_SCHEMAS_IDREFS
	XML_SCHEMAS_ENTITY
	XML_SCHEMAS_ENTITIES
	XML_SCHEMAS_NOTATION
	XML_SCHEMAS_ANYURI
	XML_SCHEMAS_INTEGER
	XML_SCHEMAS_NPINTEGER
	XML_SCHEMAS_NINTEGER
	XML_SCHEMAS_NNINTEGER
	XML_SCHEMAS_PINTEGER
	XML_SCHEMAS_INT
	XML_SCHEMAS_UINT
	XML_SCHEMAS_LONG
	XML_SCHEMAS_ULONG
	XML_SCHEMAS_SHORT
	XML_SCHEMAS_USHORT
	XML_SCHEMAS_BYTE
	XML_SCHEMAS_UBYTE
	XML_SCHEMAS_HEXBINARY
	XML_SCHEMAS_BASE64BINARY
	XML_SCHEMAS_ANYTYPE
	XML_SCHEMAS_ANYSIMPLETYPE
)

type XmlParserSeverities Enum

const (
	XML_PARSER_SEVERITY_VALIDITY_WARNING XmlParserSeverities = iota + 1
	XML_PARSER_SEVERITY_VALIDITY_ERROR
	XML_PARSER_SEVERITY_WARNING
	XML_PARSER_SEVERITY_ERROR
)

type XmlCatalogPrefer Enum

const (
	XML_CATA_PREFER_NONE XmlCatalogPrefer = iota
	XML_CATA_PREFER_PUBLIC
	XML_CATA_PREFER_SYSTEM
)

type XmlCatalogAllow Enum

const (
	XML_CATA_ALLOW_NONE XmlCatalogAllow = iota
	XML_CATA_ALLOW_GLOBAL
	XML_CATA_ALLOW_DOCUMENT
	XML_CATA_ALLOW_ALL
)

type XmlModuleOption Enum

const (
	XML_MODULE_LAZY XmlModuleOption = iota + 1
	XML_MODULE_LOCAL
)

type XmlParserOption Enum

const (
	XML_PARSE_RECOVER XmlParserOption = 1 << iota
	XML_PARSE_NOENT
	XML_PARSE_DTDLOAD
	XML_PARSE_DTDATTR
	XML_PARSE_DTDVALID
	XML_PARSE_NOERROR
	XML_PARSE_NOWARNING
	XML_PARSE_PEDANTIC
	XML_PARSE_NOBLANKS
	XML_PARSE_SAX1
	XML_PARSE_XINCLUDE
	XML_PARSE_NONET
	XML_PARSE_NODICT
	XML_PARSE_NSCLEAN
	XML_PARSE_NOCDATA
	XML_PARSE_NOXINCNODE
	XML_PARSE_COMPACT
	XML_PARSE_OLD10
	XML_PARSE_NOBASEFIX
	XML_PARSE_HUGE
	XML_PARSE_OLDSAX
	XML_PARSE_IGNORE_ENC
	XML_PARSE_BIG_LINES
)

var (
	XmlCheckVersion func(version int)

	XmlStrdup func(cur string) string

	XmlStrndup func(cur string, leng int) string

	XmlCharStrndup func(cur string, leng int) string

	XmlCharStrdup func(cur string) string

	XmlStrsub func(str string, start int, leng int) string

	XmlStrchr func(str string, val XmlChar) string

	XmlStrstr func(str, val string) string

	XmlStrcasestr func(str, val string) string

	XmlStrcmp func(str1, str2 string) int

	XmlStrncmp func(str1, str2 string, leng int) int

	XmlStrcasecmp func(str1, str2 string) int

	XmlStrncasecmp func(str1, str2 string, leng int) int

	XmlStrEqual func(str1, str2 string) int

	XmlStrQEqual func(pref, name, str string) int

	XmlStrlen func(str string) int

	XmlStrcat func(cur, add string) string

	XmlStrncat func(cur, add string, leng int) string

	XmlStrncatNew func(str1, str2 string, leng int) string

	XmlStrPrintf func(
		buf string, leng int, msg string, v ...VArg) int

	XmlStrVPrintf func(
		buf string, leng int, msg string, ap Va_list) int

	XmlGetUTF8Char func(utf *Unsigned_char, leng *int) int

	XmlCheckUTF8 func(utf *Unsigned_char) int

	XmlUTF8Strsize func(utf string, leng int) int

	XmlUTF8Strndup func(utf string, leng int) string

	XmlUTF8Strpos func(utf string, pos int) string

	XmlUTF8Strloc func(utf string, utfchar string) int

	XmlUTF8Strsub func(utf string, start, leng int) string

	XmlUTF8Strlen func(utf string) int

	XmlUTF8Size func(utf string) int

	XmlUTF8Charcmp func(utf1, utf2 string) int

	XmlBufContent func(buf XmlBufPtr) string

	XmlBufEnd func(buf XmlBufPtr) string

	XmlBufUse func(buf XmlBufPtr) Size_t

	XmlBufShrink func(buf XmlBufPtr, leng Size_t) Size_t

	XmlInitializeDict func() int

	XmlDictCreate func() XmlDictPtr

	XmlDictSetLimit func(dict XmlDictPtr, limit Size_t) Size_t

	XmlDictGetUsage func(dict XmlDictPtr) Size_t

	XmlDictCreateSub func(sub XmlDictPtr) XmlDictPtr

	XmlDictReference func(dict XmlDictPtr) int

	XmlDictFree func(dict XmlDictPtr)

	XmlDictLookup func(
		dict XmlDictPtr, name string, leng int) string

	XmlDictExists func(
		dict XmlDictPtr, name string, leng int) string

	XmlDictQLookup func(
		dict XmlDictPtr, prefix, name string) string

	XmlDictOwns func(dict XmlDictPtr, str string) int

	XmlDictSize func(dict XmlDictPtr) int

	XmlDictCleanup func()

	XmlRegexpCompile func(regexp string) XmlRegexpPtr

	XmlRegFreeRegexp func(regexp XmlRegexpPtr)

	XmlRegexpExec func(comp XmlRegexpPtr, value string) int

	XmlRegexpPrint func(output *FILE, regexp XmlRegexpPtr)

	XmlRegexpIsDeterminist func(comp XmlRegexpPtr) int

	XmlRegNewExecCtxt func(
		comp XmlRegexpPtr,
		callback XmlRegExecCallbacks,
		data *Void) XmlRegExecCtxtPtr

	XmlRegFreeExecCtxt func(
		exec XmlRegExecCtxtPtr)

	XmlRegExecPushString func(
		exec XmlRegExecCtxtPtr,
		value string,
		data *Void) int

	XmlRegExecPushString2 func(
		exec XmlRegExecCtxtPtr,
		value string,
		value2 string,
		data *Void) int

	XmlRegExecNextValues func(
		exec XmlRegExecCtxtPtr,
		nbval *int,
		nbneg *int,
		values *string,
		terminal *int) int

	XmlRegExecErrInfo func(
		exec XmlRegExecCtxtPtr,
		string *string,
		nbval *int,
		nbneg *int,
		values *string,
		terminal *int) int

	XmlExpFreeCtxt func(ctxt XmlExpCtxtPtr)

	XmlExpNewCtxt func(
		maxNodes int, dict XmlDictPtr) XmlExpCtxtPtr

	XmlExpCtxtNbNodes func(ctxt XmlExpCtxtPtr) int

	XmlExpCtxtNbCons func(ctxt XmlExpCtxtPtr) int

	XmlExpFree func(ctxt XmlExpCtxtPtr, expr XmlExpNodePtr)

	XmlExpRef func(expr XmlExpNodePtr)

	XmlExpParse func(
		ctxt XmlExpCtxtPtr, expr string) XmlExpNodePtr

	XmlExpNewAtom func(ctxt XmlExpCtxtPtr,
		name string, leng int) XmlExpNodePtr

	XmlExpNewOr func(
		ctxt XmlExpCtxtPtr,
		left XmlExpNodePtr,
		right XmlExpNodePtr) XmlExpNodePtr

	XmlExpNewSeq func(
		ctxt XmlExpCtxtPtr,
		left XmlExpNodePtr,
		right XmlExpNodePtr) XmlExpNodePtr

	XmlExpNewRange func(
		ctxt XmlExpCtxtPtr,
		subset XmlExpNodePtr,
		min int,
		max int) XmlExpNodePtr

	XmlExpIsNillable func(
		expr XmlExpNodePtr) int

	XmlExpMaxToken func(
		expr XmlExpNodePtr) int

	XmlExpGetLanguage func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		langList *string,
		leng int) int

	XmlExpGetStart func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		tokList *string,
		leng int) int

	XmlExpStringDerive func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		str string,
		leng int) XmlExpNodePtr

	XmlExpExpDerive func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		sub XmlExpNodePtr) XmlExpNodePtr

	XmlExpSubsume func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		sub XmlExpNodePtr) int

	XmlExpDump func(
		buf XmlBufferPtr,
		expr XmlExpNodePtr)

	XmlValidateNCName func(
		value string,
		space int) int

	XmlValidateQName func(
		value string,
		space int) int

	XmlValidateName func(
		value string,
		space int) int

	XmlValidateNMToken func(
		value string,
		space int) int

	XmlBuildQName func(
		ncname string,
		prefix string,
		memory string,
		leng int) string

	XmlSplitQName2 func(
		name string,
		prefix *string) string

	XmlSplitQName3 func(
		name string,
		leng *int) string

	XmlSetBufferAllocationScheme func(
		scheme XmlBufferAllocationScheme)

	XmlGetBufferAllocationScheme func() XmlBufferAllocationScheme

	XmlBufferCreate func() XmlBufferPtr

	XmlBufferCreateSize func(size Size_t) XmlBufferPtr

	XmlBufferCreateStatic func(
		mem *Void, size Size_t) XmlBufferPtr

	XmlBufferResize func(buf XmlBufferPtr, size Unsigned_int) int

	XmlBufferFree func(buf XmlBufferPtr)

	XmlBufferDump func(
		file *FILE,
		buf XmlBufferPtr) int

	XmlBufferAdd func(
		buf XmlBufferPtr,
		str string,
		leng int) int

	XmlBufferAddHead func(
		buf XmlBufferPtr,
		str string,
		leng int) int

	XmlBufferCat func(
		buf XmlBufferPtr,
		str string) int

	XmlBufferCCat func(
		buf XmlBufferPtr,
		str string) int

	XmlBufferShrink func(
		buf XmlBufferPtr,
		leng Unsigned_int) int

	XmlBufferGrow func(
		buf XmlBufferPtr,
		leng Unsigned_int) int

	XmlBufferEmpty func(
		buf XmlBufferPtr)

	XmlBufferContent func(
		buf XmlBufferPtr) string

	XmlBufferDetach func(
		buf XmlBufferPtr) string

	XmlBufferSetAllocationScheme func(
		buf XmlBufferPtr,
		scheme XmlBufferAllocationScheme)

	XmlBufferLength func(
		buf XmlBufferPtr) int

	XmlCreateIntSubset func(
		doc XmlDocPtr,
		name string,
		ExternalID string,
		SystemID string) XmlDtdPtr

	XmlNewDtd func(
		doc XmlDocPtr,
		name string,
		ExternalID string,
		SystemID string) XmlDtdPtr

	XmlGetIntSubset func(
		doc XmlDocPtr) XmlDtdPtr

	XmlFreeDtd func(
		cur XmlDtdPtr)

	XmlNewGlobalNs func(
		doc XmlDocPtr,
		href string,
		prefix string) XmlNsPtr

	XmlNewNs func(
		node XmlNodePtr,
		href string,
		prefix string) XmlNsPtr

	XmlFreeNs func(
		cur XmlNsPtr)

	XmlFreeNsList func(
		cur XmlNsPtr)

	XmlNewDoc func(
		version string) XmlDocPtr

	XmlFreeDoc func(
		cur XmlDocPtr)

	XmlNewDocProp func(
		doc XmlDocPtr,
		name string,
		value string) XmlAttrPtr

	XmlNewProp func(
		node XmlNodePtr,
		name string,
		value string) XmlAttrPtr

	XmlNewNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name string,
		value string) XmlAttrPtr

	XmlNewNsPropEatName func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name string,
		value string) XmlAttrPtr

	XmlFreePropList func(
		cur XmlAttrPtr)

	XmlFreeProp func(
		cur XmlAttrPtr)

	XmlCopyProp func(
		target XmlNodePtr,
		cur XmlAttrPtr) XmlAttrPtr

	XmlCopyPropList func(
		target XmlNodePtr,
		cur XmlAttrPtr) XmlAttrPtr

	XmlCopyDtd func(
		dtd XmlDtdPtr) XmlDtdPtr

	XmlCopyDoc func(
		doc XmlDocPtr,
		recursive int) XmlDocPtr

	XmlNewDocNode func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name string,
		content string) XmlNodePtr

	XmlNewDocNodeEatName func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name string,
		content string) XmlNodePtr

	XmlNewNode func(
		ns XmlNsPtr,
		name string) XmlNodePtr

	XmlNewNodeEatName func(
		ns XmlNsPtr,
		name string) XmlNodePtr

	XmlNewChild func(
		parent XmlNodePtr,
		ns XmlNsPtr,
		name string,
		content string) XmlNodePtr

	XmlNewDocText func(
		doc XmlDocPtr,
		content string) XmlNodePtr

	XmlNewText func(
		content string) XmlNodePtr

	XmlNewDocPI func(
		doc XmlDocPtr,
		name string,
		content string) XmlNodePtr

	XmlNewPI func(
		name string,
		content string) XmlNodePtr

	XmlNewDocTextLen func(
		doc XmlDocPtr,
		content string,
		leng int) XmlNodePtr

	XmlNewTextLen func(
		content string,
		leng int) XmlNodePtr

	XmlNewDocComment func(
		doc XmlDocPtr,
		content string) XmlNodePtr

	XmlNewComment func(
		content string) XmlNodePtr

	XmlNewCDataBlock func(
		doc XmlDocPtr,
		content string,
		leng int) XmlNodePtr

	XmlNewCharRef func(
		doc XmlDocPtr,
		name string) XmlNodePtr

	XmlNewReference func(
		doc XmlDocPtr,
		name string) XmlNodePtr

	XmlCopyNode func(
		node XmlNodePtr,
		recursive int) XmlNodePtr

	XmlDocCopyNode func(
		node XmlNodePtr,
		doc XmlDocPtr,
		recursive int) XmlNodePtr

	XmlDocCopyNodeList func(
		doc XmlDocPtr,
		node XmlNodePtr) XmlNodePtr

	XmlCopyNodeList func(
		node XmlNodePtr) XmlNodePtr

	XmlNewTextChild func(
		parent XmlNodePtr,
		ns XmlNsPtr,
		name string,
		content string) XmlNodePtr

	XmlNewDocRawNode func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name string,
		content string) XmlNodePtr

	XmlNewDocFragment func(
		doc XmlDocPtr) XmlNodePtr

	XmlGetLineNo func(
		node XmlNodePtr) Long

	XmlGetNodePath func(
		node XmlNodePtr) string

	XmlDocGetRootElement func(
		doc XmlDocPtr) XmlNodePtr

	XmlGetLastChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlNodeIsText func(
		node XmlNodePtr) int

	XmlIsBlankNode func(
		node XmlNodePtr) int

	XmlDocSetRootElement func(
		doc XmlDocPtr,
		root XmlNodePtr) XmlNodePtr

	XmlNodeSetName func(
		cur XmlNodePtr,
		name string)

	XmlAddChild func(
		parent XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlAddChildList func(
		parent XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlReplaceNode func(
		old XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlAddPrevSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlAddSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlAddNextSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlUnlinkNode func(
		cur XmlNodePtr)

	XmlTextMerge func(
		first XmlNodePtr,
		second XmlNodePtr) XmlNodePtr

	XmlTextConcat func(
		node XmlNodePtr,
		content string,
		leng int) int

	XmlFreeNodeList func(
		cur XmlNodePtr)

	XmlFreeNode func(
		cur XmlNodePtr)

	XmlSetTreeDoc func(
		tree XmlNodePtr,
		doc XmlDocPtr)

	XmlSetListDoc func(
		list XmlNodePtr,
		doc XmlDocPtr)

	XmlSearchNs func(
		doc XmlDocPtr,
		node XmlNodePtr,
		nameSpace string) XmlNsPtr

	XmlSearchNsByHref func(
		doc XmlDocPtr,
		node XmlNodePtr,
		href string) XmlNsPtr

	XmlGetNsList func(
		doc XmlDocPtr,
		node XmlNodePtr) *XmlNsPtr

	XmlSetNs func(
		node XmlNodePtr,
		ns XmlNsPtr)

	XmlCopyNamespace func(
		cur XmlNsPtr) XmlNsPtr

	XmlCopyNamespaceList func(
		cur XmlNsPtr) XmlNsPtr

	XmlSetProp func(
		node XmlNodePtr,
		name string,
		value string) XmlAttrPtr

	XmlSetNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name string,
		value string) XmlAttrPtr

	XmlGetNoNsProp func(
		node XmlNodePtr,
		name string) string

	XmlGetProp func(
		node XmlNodePtr,
		name string) string

	XmlHasProp func(
		node XmlNodePtr,
		name string) XmlAttrPtr

	XmlHasNsProp func(
		node XmlNodePtr,
		name string,
		nameSpace string) XmlAttrPtr

	XmlGetNsProp func(
		node XmlNodePtr,
		name string,
		nameSpace string) string

	XmlStringGetNodeList func(
		doc XmlDocPtr,
		value string) XmlNodePtr

	XmlStringLenGetNodeList func(
		doc XmlDocPtr,
		value string,
		leng int) XmlNodePtr

	XmlNodeListGetString func(
		doc XmlDocPtr,
		list XmlNodePtr,
		inLine int) string

	XmlNodeListGetRawString func(
		doc XmlDocPtr,
		list XmlNodePtr,
		inLine int) string

	XmlNodeSetContent func(
		cur XmlNodePtr,
		content string)

	XmlNodeSetContentLen func(
		cur XmlNodePtr,
		content string,
		leng int)

	XmlNodeAddContent func(
		cur XmlNodePtr,
		content string)

	XmlNodeAddContentLen func(
		cur XmlNodePtr,
		content string,
		leng int)

	XmlNodeGetContent func(
		cur XmlNodePtr) string

	XmlNodeBufGetContent func(
		buffer XmlBufferPtr,
		cur XmlNodePtr) int

	XmlBufGetNodeContent func(
		buf XmlBufPtr,
		cur XmlNodePtr) int

	XmlNodeGetLang func(
		cur XmlNodePtr) string

	XmlNodeGetSpacePreserve func(
		cur XmlNodePtr) int

	XmlNodeSetLang func(
		cur XmlNodePtr,
		lang string)

	XmlNodeSetSpacePreserve func(
		cur XmlNodePtr,
		val int)

	XmlNodeGetBase func(
		doc XmlDocPtr,
		cur XmlNodePtr) string

	XmlNodeSetBase func(
		cur XmlNodePtr,
		uri string)

	XmlRemoveProp func(
		cur XmlAttrPtr) int

	XmlUnsetNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name string) int

	XmlUnsetProp func(
		node XmlNodePtr,
		name string) int

	XmlBufferWriteCHAR func(
		buf XmlBufferPtr,
		string string)

	XmlBufferWriteChar func(
		buf XmlBufferPtr,
		string string)

	XmlBufferWriteQuotedString func(
		buf XmlBufferPtr,
		string string)

	XmlAttrSerializeTxtContent func(
		buf XmlBufferPtr,
		doc XmlDocPtr,
		attr XmlAttrPtr,
		string string)

	XmlReconciliateNs func(
		doc XmlDocPtr,
		tree XmlNodePtr) int

	XmlDocDumpFormatMemory func(
		cur XmlDocPtr,
		mem *string,
		size *int,
		format int)

	XmlDocDumpMemory func(
		cur XmlDocPtr,
		mem *string,
		size *int)

	XmlDocDumpMemoryEnc func(
		out_doc XmlDocPtr,
		doc_txt_ptr *string,
		doc_txt_len *int,
		txt_encoding string)

	XmlDocDumpFormatMemoryEnc func(
		out_doc XmlDocPtr,
		doc_txt_ptr *string,
		doc_txt_len *int,
		txt_encoding string,
		format int)

	XmlDocFormatDump func(
		f *FILE,
		cur XmlDocPtr,
		format int) int

	XmlDocDump func(f *FILE, cur XmlDocPtr) int

	XmlElemDump func(f *FILE, doc XmlDocPtr, cur XmlNodePtr)

	XmlSaveFile func(filename string, cur XmlDocPtr) int

	XmlSaveFormatFile func(
		filename string, cur XmlDocPtr, format int) int

	XmlBufNodeDump func(buf XmlBufPtr, doc XmlDocPtr,
		cur XmlNodePtr, level, format int) Size_t

	XmlNodeDump func(buf XmlBufferPtr, doc XmlDocPtr,
		cur XmlNodePtr, level, format int) int

	XmlSaveFileTo func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding string) int

	XmlSaveFormatFileTo func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding string,
		format int) int

	XmlNodeDumpOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		level int,
		format int,
		encoding string)

	XmlSaveFormatFileEnc func(
		filename string,
		cur XmlDocPtr,
		encoding string,
		format int) int

	XmlSaveFileEnc func(
		filename string,
		cur XmlDocPtr,
		encoding string) int

	XmlIsXHTML func(
		systemID string,
		publicID string) int

	XmlGetDocCompressMode func(
		doc XmlDocPtr) int

	XmlSetDocCompressMode func(
		doc XmlDocPtr,
		mode int)

	XmlGetCompressMode func() int

	XmlSetCompressMode func(
		mode int)

	XmlDOMWrapNewCtxt func() XmlDOMWrapCtxtPtr

	XmlDOMWrapFreeCtxt func(
		ctxt XmlDOMWrapCtxtPtr)

	XmlDOMWrapReconcileNamespaces func(
		ctxt XmlDOMWrapCtxtPtr,
		elem XmlNodePtr,
		options int) int

	XmlDOMWrapAdoptNode func(
		ctxt XmlDOMWrapCtxtPtr,
		sourceDoc XmlDocPtr,
		node XmlNodePtr,
		destDoc XmlDocPtr,
		destParent XmlNodePtr,
		options int) int

	XmlDOMWrapRemoveNode func(
		ctxt XmlDOMWrapCtxtPtr,
		doc XmlDocPtr,
		node XmlNodePtr,
		options int) int

	XmlDOMWrapCloneNode func(
		ctxt XmlDOMWrapCtxtPtr,
		sourceDoc XmlDocPtr,
		node XmlNodePtr,
		clonedNode *XmlNodePtr,
		destDoc XmlDocPtr,
		destParent XmlNodePtr,
		deep int,
		options int) int

	XmlChildElementCount func(
		parent XmlNodePtr) Unsigned_long

	XmlNextElementSibling func(
		node XmlNodePtr) XmlNodePtr

	XmlFirstElementChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlLastElementChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlPreviousElementSibling func(
		node XmlNodePtr) XmlNodePtr

	XmlMemSetup func(
		freeFunc XmlFreeFunc,
		mallocFunc XmlMallocFunc,
		reallocFunc XmlReallocFunc,
		strdupFunc XmlStrdupFunc) int

	XmlMemGet func(
		freeFunc *XmlFreeFunc,
		mallocFunc *XmlMallocFunc,
		reallocFunc *XmlReallocFunc,
		strdupFunc *XmlStrdupFunc) int

	XmlGcMemSetup func(
		freeFunc XmlFreeFunc,
		mallocFunc XmlMallocFunc,
		mallocAtomicFunc XmlMallocFunc,
		reallocFunc XmlReallocFunc,
		strdupFunc XmlStrdupFunc) int

	XmlGcMemGet func(
		freeFunc *XmlFreeFunc,
		mallocFunc *XmlMallocFunc,
		mallocAtomicFunc *XmlMallocFunc,
		reallocFunc *XmlReallocFunc,
		strdupFunc *XmlStrdupFunc) int

	XmlInitMemory func() int

	XmlCleanupMemory func()

	XmlMemUsed func() int

	XmlMemBlocks func() int

	XmlMemDisplay func(
		fp *FILE)

	XmlMemDisplayLast func(fp *FILE, nbBytes Long)

	XmlMemShow func(fp *FILE, nr int)

	XmlMemoryDump func()

	XmlMemMalloc func(size Size_t) *Void

	XmlMemRealloc func(ptr *Void, size Size_t) *Void

	XmlMemFree func(ptr *Void)

	XmlMemoryStrdup func(str string) string

	XmlMallocLoc func(size Size_t, file string, line int) *Void

	XmlReallocLoc func(
		ptr *Void, size Size_t, file string, line int) *Void

	XmlMallocAtomicLoc func(
		size Size_t,
		file string,
		line int) *Void

	XmlMemStrdupLoc func(
		str string,
		file string,
		line int) string

	XmlHashCreate func(
		size int) XmlHashTablePtr

	XmlHashCreateDict func(
		size int,
		dict XmlDictPtr) XmlHashTablePtr

	XmlHashFree func(
		table XmlHashTablePtr,
		f XmlHashDeallocator)

	XmlHashAddEntry func(
		table XmlHashTablePtr,
		name string,
		userdata *Void) int

	XmlHashUpdateEntry func(
		table XmlHashTablePtr,
		name string,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashAddEntry2 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		userdata *Void) int

	XmlHashUpdateEntry2 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashAddEntry3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string,
		userdata *Void) int

	XmlHashUpdateEntry3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry func(
		table XmlHashTablePtr,
		name string,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry2 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string,
		f XmlHashDeallocator) int

	XmlHashLookup func(
		table XmlHashTablePtr,
		name string) *Void

	XmlHashLookup2 func(
		table XmlHashTablePtr,
		name string,
		name2 string) *Void

	XmlHashLookup3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string) *Void

	XmlHashQLookup func(
		table XmlHashTablePtr,
		name string,
		prefix string) *Void

	XmlHashQLookup2 func(
		table XmlHashTablePtr,
		name string,
		prefix string,
		name2 string,
		prefix2 string) *Void

	XmlHashQLookup3 func(
		table XmlHashTablePtr,
		name string,
		prefix string,
		name2 string,
		prefix2 string,
		name3 string,
		prefix3 string) *Void

	XmlHashCopy func(
		table XmlHashTablePtr,
		f XmlHashCopier) XmlHashTablePtr

	XmlHashSize func(
		table XmlHashTablePtr) int

	XmlHashScan func(
		table XmlHashTablePtr,
		f XmlHashScanner,
		data *Void)

	XmlHashScan3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string,
		f XmlHashScanner,
		data *Void)

	XmlHashScanFull func(
		table XmlHashTablePtr,
		f XmlHashScannerFull,
		data *Void)

	XmlHashScanFull3 func(
		table XmlHashTablePtr,
		name string,
		name2 string,
		name3 string,
		f XmlHashScannerFull,
		data *Void)

	XmlSetGenericErrorFunc func(
		ctx *Void,
		handler XmlGenericErrorFunc)

	InitGenericErrorDefaultFunc func(
		handler *XmlGenericErrorFunc)

	XmlSetStructuredErrorFunc func(
		ctx *Void,
		handler XmlStructuredErrorFunc)

	XmlParserError func(
		ctx *Void,
		msg string,
		v ...VArg)

	XmlParserWarning func(
		ctx *Void,
		msg string,
		v ...VArg)

	XmlParserValidityError func(
		ctx *Void, msg string, v ...VArg)

	XmlParserValidityWarning func(
		ctx *Void, msg string, v ...VArg)

	XmlParserPrintFileInfo func(input XmlParserInputPtr)

	XmlParserPrintFileContext func(input XmlParserInputPtr)

	XmlGetLastError func() XmlErrorPtr

	XmlResetLastError func()

	XmlCtxtGetLastError func(ctx *Void) XmlErrorPtr

	XmlCtxtResetLastError func(ctx *Void)

	XmlResetError func(err XmlErrorPtr)

	XmlCopyError func(from XmlErrorPtr, to XmlErrorPtr) int

	XmlListCreate func(
		deallocator XmlListDeallocator,
		compare XmlListDataCompare) XmlListPtr

	XmlListDelete func(
		l XmlListPtr)

	XmlListSearch func(
		l XmlListPtr,
		data *Void) *Void

	XmlListReverseSearch func(
		l XmlListPtr,
		data *Void) *Void

	XmlListInsert func(
		l XmlListPtr,
		data *Void) int

	XmlListAppend func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveFirst func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveLast func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveAll func(
		l XmlListPtr,
		data *Void) int

	XmlListClear func(
		l XmlListPtr)

	XmlListEmpty func(
		l XmlListPtr) int

	XmlListFront func(
		l XmlListPtr) XmlLinkPtr

	XmlListEnd func(
		l XmlListPtr) XmlLinkPtr

	XmlListSize func(
		l XmlListPtr) int

	XmlListPopFront func(
		l XmlListPtr)

	XmlListPopBack func(
		l XmlListPtr)

	XmlListPushFront func(
		l XmlListPtr,
		data *Void) int

	XmlListPushBack func(
		l XmlListPtr,
		data *Void) int

	XmlListReverse func(
		l XmlListPtr)

	XmlListSort func(
		l XmlListPtr)

	XmlListWalk func(
		l XmlListPtr,
		walker XmlListWalker,
		user *Void)

	XmlListReverseWalk func(
		l XmlListPtr,
		walker XmlListWalker,
		user *Void)

	XmlListMerge func(
		l1 XmlListPtr,
		l2 XmlListPtr)

	XmlListDup func(
		old XmlListPtr) XmlListPtr

	XmlListCopy func(
		cur XmlListPtr,
		old XmlListPtr) int

	XmlLinkGetData func(
		lk XmlLinkPtr) *Void

	XmlNewAutomata func() XmlAutomataPtr

	XmlFreeAutomata func(
		am XmlAutomataPtr)

	XmlAutomataGetInitState func(
		am XmlAutomataPtr) XmlAutomataStatePtr

	XmlAutomataSetFinalState func(
		am XmlAutomataPtr,
		state XmlAutomataStatePtr) int

	XmlAutomataNewState func(
		am XmlAutomataPtr) XmlAutomataStatePtr

	XmlAutomataNewTransition func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewTransition2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		token2 string,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewNegTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		token2 string,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewCountTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewCountTrans2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		token2 string,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewOnceTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewOnceTrans2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token string,
		token2 string,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewAllTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		lax int) XmlAutomataStatePtr

	XmlAutomataNewEpsilon func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr) XmlAutomataStatePtr

	XmlAutomataNewCountedTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		counter int) XmlAutomataStatePtr

	XmlAutomataNewCounterTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		counter int) XmlAutomataStatePtr

	XmlAutomataNewCounter func(
		am XmlAutomataPtr,
		min int,
		max int) int

	XmlAutomataCompile func(
		am XmlAutomataPtr) XmlRegexpPtr

	XmlAutomataIsDeterminist func(
		am XmlAutomataPtr) int

	XmlAddNotationDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		name string,
		PublicID string,
		SystemID string) XmlNotationPtr

	XmlCopyNotationTable func(
		table XmlNotationTablePtr) XmlNotationTablePtr

	XmlFreeNotationTable func(table XmlNotationTablePtr)

	XmlDumpNotationDecl func(
		buf XmlBufferPtr, nota XmlNotationPtr)

	XmlDumpNotationTable func(
		buf XmlBufferPtr, table XmlNotationTablePtr)

	XmlNewElementContent func(name string,
		t XmlElementContentType) XmlElementContentPtr

	XmlCopyElementContent func(
		content XmlElementContentPtr) XmlElementContentPtr

	XmlFreeElementContent func(cur XmlElementContentPtr)

	XmlNewDocElementContent func(doc XmlDocPtr,
		name string, t XmlElementContentType) XmlElementContentPtr

	XmlCopyDocElementContent func(doc XmlDocPtr,
		content XmlElementContentPtr) XmlElementContentPtr

	XmlFreeDocElementContent func(
		doc XmlDocPtr, cur XmlElementContentPtr)

	XmlSnprintfElementContent func(buf string, size int,
		content XmlElementContentPtr, englob int)

	XmlSprintfElementContent func(
		buf string, content XmlElementContentPtr, englob int)

	XmlAddElementDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		name string,
		t XmlElementTypeVal,
		content XmlElementContentPtr) XmlElementPtr

	XmlCopyElementTable func(
		table XmlElementTablePtr) XmlElementTablePtr

	XmlFreeElementTable func(table XmlElementTablePtr)

	XmlDumpElementTable func(
		buf XmlBufferPtr, table XmlElementTablePtr)

	XmlDumpElementDecl func(
		buf XmlBufferPtr, elem XmlElementPtr)

	XmlCreateEnumeration func(name string) XmlEnumerationPtr

	XmlFreeEnumeration func(cur XmlEnumerationPtr)

	XmlCopyEnumeration func(
		cur XmlEnumerationPtr) XmlEnumerationPtr

	XmlAddAttributeDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		elem, name, ns string,
		t XmlAttributeType,
		def XmlAttributeDefault,
		defaultValue string,
		tree XmlEnumerationPtr) XmlAttributePtr

	XmlCopyAttributeTable func(
		table XmlAttributeTablePtr) XmlAttributeTablePtr

	XmlFreeAttributeTable func(table XmlAttributeTablePtr)

	XmlDumpAttributeTable func(
		buf XmlBufferPtr, table XmlAttributeTablePtr)

	XmlDumpAttributeDecl func(
		buf XmlBufferPtr, attr XmlAttributePtr)

	XmlAddID func(ctxt XmlValidCtxtPtr, doc XmlDocPtr,
		value string, attr XmlAttrPtr) XmlIDPtr

	XmlFreeIDTable func(table XmlIDTablePtr)

	XmlGetID func(doc XmlDocPtr, ID string) XmlAttrPtr

	XmlIsID func(
		doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) int

	XmlRemoveID func(doc XmlDocPtr, attr XmlAttrPtr) int

	XmlAddRef func(ctxt XmlValidCtxtPtr, doc XmlDocPtr,
		value string, attr XmlAttrPtr) XmlRefPtr

	XmlFreeRefTable func(table XmlRefTablePtr)

	XmlIsRef func(doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) int

	XmlRemoveRef func(doc XmlDocPtr, attr XmlAttrPtr) int

	XmlGetRefs func(doc XmlDocPtr, ID string) XmlListPtr

	XmlNewValidCtxt func() XmlValidCtxtPtr

	XmlFreeValidCtxt func(XmlValidCtxtPtr)

	XmlValidateRoot func(ctxt XmlValidCtxtPtr, doc XmlDocPtr) int

	XmlValidateElementDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, elem XmlElementPtr) int

	XmlValidNormalizeAttributeValue func(doc XmlDocPtr,
		elem XmlNodePtr, name string, value string) string

	XmlValidCtxtNormalizeAttributeValue func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		name string,
		value string) string

	XmlValidateAttributeDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, attr XmlAttributePtr) int

	XmlValidateAttributeValue func(
		t XmlAttributeType, value string) int

	XmlValidateNotationDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, nota XmlNotationPtr) int

	XmlValidateDtd func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, dtd XmlDtdPtr) int

	XmlValidateDtdFinal func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateDocument func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlValidateOneElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlValidateOneAttribute func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		attr XmlAttrPtr,
		value string) int

	XmlValidateOneNamespace func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		prefix string,
		ns XmlNsPtr,
		value string) int

	XmlValidateDocumentFinal func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateNotationUse func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		notationName string) int

	XmlIsMixedElement func(
		doc XmlDocPtr,
		name string) int

	XmlGetDtdAttrDesc func(
		dtd XmlDtdPtr,
		elem string,
		name string) XmlAttributePtr

	XmlGetDtdQAttrDesc func(
		dtd XmlDtdPtr,
		elem string,
		name string,
		prefix string) XmlAttributePtr

	XmlGetDtdNotationDesc func(
		dtd XmlDtdPtr,
		name string) XmlNotationPtr

	XmlGetDtdQElementDesc func(
		dtd XmlDtdPtr,
		name string,
		prefix string) XmlElementPtr

	XmlGetDtdElementDesc func(
		dtd XmlDtdPtr,
		name string) XmlElementPtr

	XmlValidGetPotentialChildren func(
		ctree *XmlElementContent,
		names *string,
		leng *int,
		max int) int

	XmlValidGetValidElements func(
		prev, next *XmlNode,
		names *string,
		max int) int

	XmlValidateNameValue func(
		value string) int

	XmlValidateNamesValue func(
		value string) int

	XmlValidateNmtokenValue func(
		value string) int

	XmlValidateNmtokensValue func(
		value string) int

	XmlValidBuildContentModel func(
		ctxt XmlValidCtxtPtr,
		elem XmlElementPtr) int

	XmlValidatePushElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		qname string) int

	XmlValidatePushCData func(
		ctxt XmlValidCtxtPtr,
		data string,
		leng int) int

	XmlValidatePopElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		qname string) int

	XmlInitializePredefinedEntities func()

	XmlNewEntity func(
		doc XmlDocPtr,
		name string,
		typ int,
		ExternalID string,
		SystemID string,
		content string) XmlEntityPtr

	XmlAddDocEntity func(
		doc XmlDocPtr,
		name string,
		typ int,
		ExternalID string,
		SystemID string,
		content string) XmlEntityPtr

	XmlAddDtdEntity func(
		doc XmlDocPtr,
		name string,
		typ int,
		ExternalID string,
		SystemID string,
		content string) XmlEntityPtr

	XmlGetPredefinedEntity func(
		name string) XmlEntityPtr

	XmlGetDocEntity func(
		doc XmlDocPtr,
		name string) XmlEntityPtr

	XmlGetDtdEntity func(
		doc XmlDocPtr,
		name string) XmlEntityPtr

	XmlGetParameterEntity func(
		doc XmlDocPtr,
		name string) XmlEntityPtr

	XmlEncodeEntities func(
		doc XmlDocPtr,
		input string) string

	XmlEncodeEntitiesReentrant func(
		doc XmlDocPtr,
		input string) string

	XmlEncodeSpecialChars func(
		doc XmlDocPtr,
		input string) string

	XmlCreateEntitiesTable func() XmlEntitiesTablePtr

	XmlCopyEntitiesTable func(
		table XmlEntitiesTablePtr) XmlEntitiesTablePtr

	XmlFreeEntitiesTable func(
		table XmlEntitiesTablePtr)

	XmlDumpEntitiesTable func(
		buf XmlBufferPtr,
		table XmlEntitiesTablePtr)

	XmlDumpEntityDecl func(
		buf XmlBufferPtr,
		ent XmlEntityPtr)

	XmlCleanupPredefinedEntities func()

	XmlInitCharEncodingHandlers func()

	XmlCleanupCharEncodingHandlers func()

	XmlRegisterCharEncodingHandler func(
		handler XmlCharEncodingHandlerPtr)

	XmlGetCharEncodingHandler func(
		enc XmlCharEncoding) XmlCharEncodingHandlerPtr

	XmlFindCharEncodingHandler func(
		name string) XmlCharEncodingHandlerPtr

	XmlNewCharEncodingHandler func(
		name string,
		input XmlCharEncodingInputFunc,
		output XmlCharEncodingOutputFunc) XmlCharEncodingHandlerPtr

	XmlAddEncodingAlias func(
		name string,
		alias string) int

	XmlDelEncodingAlias func(
		alias string) int

	XmlGetEncodingAlias func(
		alias string) string

	XmlCleanupEncodingAliases func()

	XmlParseCharEncoding func(
		name string) XmlCharEncoding

	XmlGetCharEncodingName func(
		enc XmlCharEncoding) string

	XmlDetectCharEncoding func(
		in *Unsigned_char,
		leng int) XmlCharEncoding

	XmlCharEncOutFunc func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncInFunc func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncFirstLine func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncCloseFunc func(
		handler XmlCharEncodingHandler) int

	UTF8Toisolat1 func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	Isolat1ToUTF8 func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	XmlCleanupInputCallbacks func()

	XmlPopInputCallbacks func() int

	XmlRegisterDefaultInputCallbacks func()

	XmlAllocParserInputBuffer func(
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateFile func(
		file *FILE,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateFd func(
		fd int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateMem func(
		mem string,
		size int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateStatic func(
		mem string,
		size int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferRead func(
		in XmlParserInputBufferPtr,
		leng int) int

	XmlParserInputBufferGrow func(
		in XmlParserInputBufferPtr,
		leng int) int

	XmlParserInputBufferPush func(
		in XmlParserInputBufferPtr,
		leng int,
		buf string) int

	XmlFreeParserInputBuffer func(
		in XmlParserInputBufferPtr)

	XmlParserGetDirectory func(
		filename string) string

	XmlRegisterInputCallbacks func(
		matchFunc XmlInputMatchCallback,
		openFunc XmlInputOpenCallback,
		readFunc XmlInputReadCallback,
		closeFunc XmlInputCloseCallback) int

	XmlParserInputBufferCreateFilename func(
		URI string,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlCleanupOutputCallbacks func()

	XmlRegisterDefaultOutputCallbacks func()

	XmlAllocOutputBuffer func(
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateFilename func(
		URI string,
		encoder XmlCharEncodingHandlerPtr,
		compression int) XmlOutputBufferPtr

	XmlOutputBufferCreateFile func(
		file *FILE,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateBuffer func(
		buffer XmlBufferPtr,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateFd func(
		fd int,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateIO func(
		iowrite XmlOutputWriteCallback,
		ioclose XmlOutputCloseCallback,
		ioctx *Void,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferGetContent func(
		out XmlOutputBufferPtr) string

	XmlOutputBufferGetSize func(
		out XmlOutputBufferPtr) Size_t

	XmlOutputBufferWrite func(
		out XmlOutputBufferPtr,
		leng int,
		buf string) int

	XmlOutputBufferWriteString func(
		out XmlOutputBufferPtr,
		str string) int

	XmlOutputBufferWriteEscape func(
		out XmlOutputBufferPtr,
		str string,
		escaping XmlCharEncodingOutputFunc) int

	XmlOutputBufferFlush func(
		out XmlOutputBufferPtr) int

	XmlOutputBufferClose func(
		out XmlOutputBufferPtr) int

	XmlRegisterOutputCallbacks func(
		matchFunc XmlOutputMatchCallback,
		openFunc XmlOutputOpenCallback,
		writeFunc XmlOutputWriteCallback,
		closeFunc XmlOutputCloseCallback) int

	XmlRegisterHTTPPostCallbacks func()

	XmlCheckHTTPInput func(
		ctxt XmlParserCtxtPtr,
		ret XmlParserInputPtr) XmlParserInputPtr

	XmlNoNetExternalEntityLoader func(
		URL string,
		ID string,
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	XmlNormalizeWindowsPath func(
		path string) string

	XmlCheckFilename func(
		path string) int

	XmlFileMatch func(
		filename string) int

	XmlFileOpen func(
		filename string) *Void

	XmlFileRead func(
		context *Void,
		buffer string,
		leng int) int

	XmlFileClose func(context *Void) int

	XmlIOHTTPMatch func(filename string) int

	XmlIOHTTPOpen func(filename string) *Void

	XmlIOHTTPOpenW func(post_uri string, compression int) *Void

	XmlIOHTTPRead func(context *Void, buffer string, leng int) int

	XmlIOHTTPClose func(
		context *Void) int

	XmlIOFTPMatch func(
		filename string) int

	XmlIOFTPOpen func(
		filename string) *Void

	XmlIOFTPRead func(
		context *Void,
		buffer string,
		leng int) int

	XmlIOFTPClose func(
		context *Void) int

	XmlInitParser func()

	XmlCleanupParser func()

	XmlParserInputRead func(
		in XmlParserInputPtr,
		leng int) int

	XmlParserInputGrow func(
		in XmlParserInputPtr,
		leng int) int

	XmlParseDoc func(
		cur string) XmlDocPtr

	XmlParseFile func(
		filename string) XmlDocPtr

	XmlParseMemory func(
		buffer string,
		size int) XmlDocPtr

	XmlSubstituteEntitiesDefault func(
		val int) int

	XmlKeepBlanksDefault func(
		val int) int

	XmlStopParser func(
		ctxt XmlParserCtxtPtr)

	XmlPedanticParserDefault func(
		val int) int

	XmlLineNumbersDefault func(
		val int) int

	XmlRecoverDoc func(
		cur string) XmlDocPtr

	XmlRecoverMemory func(
		buffer string,
		size int) XmlDocPtr

	XmlRecoverFile func(
		filename string) XmlDocPtr

	XmlParseDocument func(
		ctxt XmlParserCtxtPtr) int

	XmlParseExtParsedEnt func(
		ctxt XmlParserCtxtPtr) int

	XmlSAXUserParseFile func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		filename string) int

	XmlSAXUserParseMemory func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		buffer string,
		size int) int

	XmlSAXParseDoc func(
		sax XmlSAXHandlerPtr,
		cur string,
		recovery int) XmlDocPtr

	XmlSAXParseMemory func(
		sax XmlSAXHandlerPtr,
		buffer string,
		size int,
		recovery int) XmlDocPtr

	XmlSAXParseMemoryWithData func(
		sax XmlSAXHandlerPtr,
		buffer string,
		size int,
		recovery int,
		data *Void) XmlDocPtr

	XmlSAXParseFile func(
		sax XmlSAXHandlerPtr,
		filename string,
		recovery int) XmlDocPtr

	XmlSAXParseFileWithData func(
		sax XmlSAXHandlerPtr,
		filename string,
		recovery int,
		data *Void) XmlDocPtr

	XmlSAXParseEntity func(
		sax XmlSAXHandlerPtr,
		filename string) XmlDocPtr

	XmlParseEntity func(
		filename string) XmlDocPtr

	XmlSAXParseDTD func(
		sax XmlSAXHandlerPtr,
		ExternalID string,
		SystemID string) XmlDtdPtr

	XmlParseDTD func(
		ExternalID string,
		SystemID string) XmlDtdPtr

	XmlIOParseDTD func(
		sax XmlSAXHandlerPtr,
		input XmlParserInputBufferPtr,
		enc XmlCharEncoding) XmlDtdPtr

	XmlParseBalancedChunkMemory func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		string string,
		lst *XmlNodePtr) int

	XmlParseInNodeContext func(
		node XmlNodePtr,
		data string,
		datalen int,
		options int,
		lst *XmlNodePtr) XmlParserErrors

	XmlParseBalancedChunkMemoryRecover func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		string string,
		lst *XmlNodePtr,
		recover int) int

	XmlParseExternalEntity func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		URL string,
		ID string,
		lst *XmlNodePtr) int

	XmlParseCtxtExternalEntity func(
		ctx XmlParserCtxtPtr,
		URL string,
		ID string,
		lst *XmlNodePtr) int

	XmlNewParserCtxt func() XmlParserCtxtPtr

	XmlInitParserCtxt func(
		ctxt XmlParserCtxtPtr) int

	XmlClearParserCtxt func(
		ctxt XmlParserCtxtPtr)

	XmlFreeParserCtxt func(
		ctxt XmlParserCtxtPtr)

	XmlSetupParserForBuffer func(
		ctxt XmlParserCtxtPtr,
		buffer string,
		filename string)

	XmlCreateDocParserCtxt func(
		cur string) XmlParserCtxtPtr

	XmlGetFeaturesList func(
		leng *int,
		result *string) int

	XmlGetFeature func(
		ctxt XmlParserCtxtPtr,
		name string,
		result *Void) int

	XmlSetFeature func(
		ctxt XmlParserCtxtPtr,
		name string,
		value *Void) int

	XmlCreatePushParserCtxt func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		chunk string,
		size int,
		filename string) XmlParserCtxtPtr

	XmlParseChunk func(
		ctxt XmlParserCtxtPtr,
		chunk string,
		size int,
		terminate int) int

	XmlCreateIOParserCtxt func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		enc XmlCharEncoding) XmlParserCtxtPtr

	XmlNewIOInputStream func(
		ctxt XmlParserCtxtPtr,
		input XmlParserInputBufferPtr,
		enc XmlCharEncoding) XmlParserInputPtr

	XmlParserFindNodeInfo func(
		ctxt XmlParserCtxtPtr,
		node XmlNodePtr) *XmlParserNodeInfo

	XmlInitNodeInfoSeq func(
		seq XmlParserNodeInfoSeqPtr)

	XmlClearNodeInfoSeq func(
		seq XmlParserNodeInfoSeqPtr)

	XmlParserFindNodeInfoIndex func(
		seq XmlParserNodeInfoSeqPtr,
		node XmlNodePtr) Unsigned_long

	XmlParserAddNodeInfo func(
		ctxt XmlParserCtxtPtr,
		info XmlParserNodeInfoPtr)

	XmlSetExternalEntityLoader func(
		f XmlExternalEntityLoader)

	XmlGetExternalEntityLoader func() XmlExternalEntityLoader

	XmlLoadExternalEntity func(
		URL string,
		ID string,
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	XmlByteConsumed func(
		ctxt XmlParserCtxtPtr) Long

	XmlCtxtReset func(
		ctxt XmlParserCtxtPtr)

	XmlCtxtResetPush func(
		ctxt XmlParserCtxtPtr,
		chunk string,
		size int,
		filename string,
		encoding string) int

	XmlCtxtUseOptions func(
		ctxt XmlParserCtxtPtr,
		options int) int

	XmlReadDoc func(
		cur string,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlReadFile func(
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlReadMemory func(
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlReadFd func(
		fd int,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlReadIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlCtxtReadDoc func(
		ctxt XmlParserCtxtPtr,
		cur string,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlCtxtReadFile func(
		ctxt XmlParserCtxtPtr,
		filename string,
		encoding string,
		options XmlParserOption /*was int*/) XmlDocPtr

	XmlCtxtReadMemory func(
		ctxt XmlParserCtxtPtr,
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlCtxtReadFd func(
		ctxt XmlParserCtxtPtr,
		fd int,
		URL string,
		encoding string,
		options int) XmlDocPtr

	XmlCtxtReadIO func(
		ctxt XmlParserCtxtPtr,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL, encoding string,
		options int) XmlDocPtr

	XmlHasFeature func(feature XmlFeature) int

	XlinkGetDefaultDetect func() XlinkNodeDetectFunc

	XlinkSetDefaultDetect func(f XlinkNodeDetectFunc)

	XlinkGetDefaultHandler func() XlinkHandlerPtr

	XlinkSetDefaultHandler func(handler XlinkHandlerPtr)

	XlinkIsLink func(doc, node XmlNodePtr) XlinkType

	GetPublicId func(ctx *Void) string

	GetSystemId func(ctx *Void) string

	SetDocumentLocator func(ctx *Void, loc XmlSAXLocatorPtr)

	GetLineNumber func(ctx *Void) int

	GetColumnNumber func(ctx *Void) int

	IsStandalone func(ctx *Void) int

	HasInternalSubset func(ctx *Void) int

	HasExternalSubset func(ctx *Void) int

	InternalSubset func(
		ctx *Void, name, ExternalID, SystemID string)

	ExternalSubset func(
		ctx *Void, name, ExternalID, SystemID string)

	GetEntity func(ctx *Void, name string) XmlEntityPtr

	GetParameterEntity func(
		ctx *Void, name string) XmlEntityPtr

	ResolveEntity func(ctx *Void,
		publicId, systemId string) XmlParserInputPtr

	EntityDecl func(ctx *Void, name string, typ int,
		publicId, systemId, content string)

	AttributeDecl func(ctx *Void, elem, fullname string,
		typ, def int, defaultValue string,
		tree XmlEnumerationPtr)

	ElementDecl func(ctx *Void, name string,
		typ int, content XmlElementContentPtr)

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

	GetNamespace func(ctx *Void) XmlNsPtr

	CheckNamespace func(ctx *Void, nameSpace string) int

	NamespaceDecl func(ctx *Void, href, prefix string)

	Comment func(ctx *Void, value string)

	CdataBlock func(ctx *Void, value string, leng int)

	InitXmlDefaultSAXHandler func(
		hdlr *XmlSAXHandlerV1, warning int)

	InithtmlDefaultSAXHandler func(hdlr *XmlSAXHandlerV1)

	InitdocbDefaultSAXHandler func(hdlr *XmlSAXHandlerV1)

	XmlSAX2GetPublicId func(ctx *Void) string

	XmlSAX2GetSystemId func(ctx *Void) string

	XmlSAX2SetDocumentLocator func(
		ctx *Void, loc XmlSAXLocatorPtr)

	XmlSAX2GetLineNumber func(ctx *Void) int

	XmlSAX2GetColumnNumber func(ctx *Void) int

	XmlSAX2IsStandalone func(ctx *Void) int

	XmlSAX2HasInternalSubset func(ctx *Void) int

	XmlSAX2HasExternalSubset func(ctx *Void) int

	XmlSAX2InternalSubset func(
		ctx *Void,
		name string,
		ExternalID string,
		SystemID string)

	XmlSAX2ExternalSubset func(
		ctx *Void,
		name string,
		ExternalID string,
		SystemID string)

	XmlSAX2GetEntity func(
		ctx *Void,
		name string) XmlEntityPtr

	XmlSAX2GetParameterEntity func(
		ctx *Void,
		name string) XmlEntityPtr

	XmlSAX2ResolveEntity func(
		ctx *Void,
		publicId string,
		systemId string) XmlParserInputPtr

	XmlSAX2EntityDecl func(
		ctx *Void,
		name string,
		typ int,
		publicId string,
		systemId string,
		content string)

	XmlSAX2AttributeDecl func(
		ctx *Void,
		elem string,
		fullname string,
		typ int,
		def int,
		defaultValue string,
		tree XmlEnumerationPtr)

	XmlSAX2ElementDecl func(
		ctx *Void,
		name string,
		typ int,
		content XmlElementContentPtr)

	XmlSAX2NotationDecl func(
		ctx *Void,
		name string,
		publicId string,
		systemId string)

	XmlSAX2UnparsedEntityDecl func(
		ctx *Void,
		name string,
		publicId string,
		systemId string,
		notationName string)

	XmlSAX2StartDocument func(
		ctx *Void)

	XmlSAX2EndDocument func(
		ctx *Void)

	XmlSAX2StartElement func(
		ctx *Void,
		fullname string,
		atts *string)

	XmlSAX2EndElement func(
		ctx *Void,
		name string)

	XmlSAX2StartElementNs func(
		ctx *Void,
		localname string,
		prefix string,
		URI string,
		nb_namespaces int,
		namespaces *string,
		nb_attributes int,
		nb_defaulted int,
		attributes *string)

	XmlSAX2EndElementNs func(
		ctx *Void,
		localname string,
		prefix string,
		URI string)

	XmlSAX2Reference func(
		ctx *Void,
		name string)

	XmlSAX2Characters func(
		ctx *Void,
		ch string,
		leng int)

	XmlSAX2IgnorableWhitespace func(
		ctx *Void,
		ch string,
		leng int)

	XmlSAX2ProcessingInstruction func(
		ctx *Void,
		target string,
		data string)

	XmlSAX2Comment func(
		ctx *Void,
		value string)

	XmlSAX2CDataBlock func(
		ctx *Void,
		value string,
		leng int)

	XmlSAXDefaultVersion func(
		version int) int

	XmlSAXVersion func(
		hdlr *XmlSAXHandler,
		version int) int

	XmlSAX2InitDefaultSAXHandler func(
		hdlr *XmlSAXHandler,
		warning int)

	XmlSAX2InitHtmlDefaultSAXHandler func(
		hdlr *XmlSAXHandler)

	HtmlDefaultSAXHandlerInit func()

	XmlSAX2InitDocbDefaultSAXHandler func(
		hdlr *XmlSAXHandler)

	DocbDefaultSAXHandlerInit func()

	XmlDefaultSAXHandlerInit func()

	XmlInitGlobals func()

	XmlCleanupGlobals func()

	XmlInitializeGlobalState func(
		gs XmlGlobalStatePtr)

	XmlThrDefSetGenericErrorFunc func(
		ctx *Void,
		handler XmlGenericErrorFunc)

	XmlThrDefSetStructuredErrorFunc func(
		ctx *Void,
		handler XmlStructuredErrorFunc)

	XmlRegisterNodeDefault func(
		f XmlRegisterNodeFunc) XmlRegisterNodeFunc

	XmlThrDefRegisterNodeDefault func(
		f XmlRegisterNodeFunc) XmlRegisterNodeFunc

	XmlDeregisterNodeDefault func(
		f XmlDeregisterNodeFunc) XmlDeregisterNodeFunc

	XmlThrDefDeregisterNodeDefault func(
		f XmlDeregisterNodeFunc) XmlDeregisterNodeFunc

	XmlThrDefOutputBufferCreateFilenameDefault func(
		f XmlOutputBufferCreateFilenameFunc) XmlOutputBufferCreateFilenameFunc

	XmlThrDefParserInputBufferCreateFilenameDefault func(
		f XmlParserInputBufferCreateFilenameFunc) XmlParserInputBufferCreateFilenameFunc

	DocbDefaultSAXHandler func() *XmlSAXHandlerV1

	HtmlDefaultSAXHandler func() *XmlSAXHandlerV1

	XmlLastError func() *XmlError

	OldXMLWDcompatibility func() *int

	XmlBufferAllocScheme func() *XmlBufferAllocationScheme

	XmlThrDefBufferAllocScheme func(
		v XmlBufferAllocationScheme) XmlBufferAllocationScheme

	XmlDefaultBufferSize func() *int

	XmlThrDefDefaultBufferSize func(v int) int

	XmlDefaultSAXHandler func() *XmlSAXHandlerV1

	XmlDefaultSAXLocator func() *XmlSAXLocator

	XmlDoValidityCheckingDefaultValue func() *int

	XmlThrDefDoValidityCheckingDefaultValue func(v int) int

	XmlGenericError func() *XmlGenericErrorFunc

	XmlStructuredError func() *XmlStructuredErrorFunc

	XmlGenericErrorContext func() **Void

	XmlStructuredErrorContext func() **Void

	XmlGetWarningsDefaultValue func() *int

	XmlThrDefGetWarningsDefaultValue func(v int) int

	XmlIndentTreeOutput func() *int

	XmlThrDefIndentTreeOutput func(v int) int

	XmlTreeIndentString func() *string

	XmlThrDefTreeIndentString func(v string) string

	XmlKeepBlanksDefaultValue func() *int

	XmlThrDefKeepBlanksDefaultValue func(v int) int

	XmlLineNumbersDefaultValue func() *int

	XmlThrDefLineNumbersDefaultValue func(v int) int

	XmlLoadExtDtdDefaultValue func() *int

	XmlThrDefLoadExtDtdDefaultValue func(v int) int

	XmlParserDebugEntities func() *int

	XmlThrDefParserDebugEntities func(v int) int

	XmlParserVersion func() *string

	XmlPedanticParserDefaultValue func() *int

	XmlThrDefPedanticParserDefaultValue func(v int) int

	XmlSaveNoEmptyTags func() *int

	XmlThrDefSaveNoEmptyTags func(v int) int

	XmlSubstituteEntitiesDefaultValue func() *int

	XmlThrDefSubstituteEntitiesDefaultValue func(v int) int

	XmlRegisterNodeDefaultValue func() *XmlRegisterNodeFunc

	XmlDeregisterNodeDefaultValue func() *XmlDeregisterNodeFunc

	XmlParserInputBufferCreateFilenameValue func() *XmlParserInputBufferCreateFilenameFunc

	XmlOutputBufferCreateFilenameValue func() *XmlOutputBufferCreateFilenameFunc

	XmlNewMutex func() XmlMutexPtr

	XmlMutexLock func(
		tok XmlMutexPtr)

	XmlMutexUnlock func(tok XmlMutexPtr)

	XmlFreeMutex func(tok XmlMutexPtr)

	XmlNewRMutex func() XmlRMutexPtr

	XmlRMutexLock func(
		tok XmlRMutexPtr)

	XmlRMutexUnlock func(
		tok XmlRMutexPtr)

	XmlFreeRMutex func(
		tok XmlRMutexPtr)

	XmlInitThreads func()

	XmlLockLibrary func()

	XmlUnlockLibrary func()

	XmlGetThreadId func() int

	XmlIsMainThread func() int

	XmlCleanupThreads func()

	XmlGetGlobalState func() XmlGlobalStatePtr

	HtmlTagLookup func(
		tag string) *HtmlElemDesc

	HtmlEntityLookup func(name string) *HtmlEntityDesc

	HtmlEntityValueLookup func(
		value Unsigned_int) *HtmlEntityDesc

	HtmlIsAutoClosed func(doc HtmlDocPtr, elem HtmlNodePtr) int

	HtmlAutoCloseTag func(
		doc HtmlDocPtr, name string, elem HtmlNodePtr) int

	HtmlParseEntityRef func(
		ctxt HtmlParserCtxtPtr, str *string) *HtmlEntityDesc

	HtmlParseCharRef func(ctxt HtmlParserCtxtPtr) int

	HtmlParseElement func(ctxt HtmlParserCtxtPtr)

	HtmlNewParserCtxt func() HtmlParserCtxtPtr

	HtmlCreateMemoryParserCtxt func(
		buffer string, size int) HtmlParserCtxtPtr

	HtmlParseDocument func(ctxt HtmlParserCtxtPtr) int

	HtmlSAXParseDoc func(
		cur string,
		encoding string,
		sax HtmlSAXHandlerPtr,
		userData *Void) HtmlDocPtr

	HtmlParseDoc func(
		cur string,
		encoding string) HtmlDocPtr

	HtmlSAXParseFile func(
		filename string,
		encoding string,
		sax HtmlSAXHandlerPtr,
		userData *Void) HtmlDocPtr

	HtmlParseFile func(
		filename string,
		encoding string) HtmlDocPtr
	UTF8ToHtml func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	HtmlEncodeEntities func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int,
		quoteChar int) int

	HtmlIsScriptAttribute func(
		name string) int

	HtmlHandleOmittedElem func(
		val int) int

	HtmlCreatePushParserCtxt func(
		sax HtmlSAXHandlerPtr,
		user_data *Void,
		chunk string,
		size int,
		filename string,
		enc XmlCharEncoding) HtmlParserCtxtPtr

	HtmlParseChunk func(
		ctxt HtmlParserCtxtPtr,
		chunk string,
		size int,
		terminate int) int

	HtmlFreeParserCtxt func(
		ctxt HtmlParserCtxtPtr)

	HtmlCtxtReset func(
		ctxt HtmlParserCtxtPtr)

	HtmlCtxtUseOptions func(
		ctxt HtmlParserCtxtPtr,
		options int) int

	HtmlReadDoc func(
		cur string,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlReadFile func(
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlReadMemory func(
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlReadFd func(
		fd int,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlReadIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlCtxtReadDoc func(
		ctxt XmlParserCtxtPtr,
		cur string,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlCtxtReadFile func(
		ctxt XmlParserCtxtPtr,
		filename string,
		encoding string,
		options int) HtmlDocPtr

	HtmlCtxtReadMemory func(
		ctxt XmlParserCtxtPtr,
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlCtxtReadFd func(
		ctxt XmlParserCtxtPtr,
		fd int,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlCtxtReadIO func(
		ctxt XmlParserCtxtPtr,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL string,
		encoding string,
		options int) HtmlDocPtr

	HtmlAttrAllowed func(
		*HtmlElemDesc,
		string,
		int) HtmlStatus

	HtmlElementAllowedHere func(
		*HtmlElemDesc,
		string) int

	HtmlElementStatusHere func(
		*HtmlElemDesc,
		*HtmlElemDesc) HtmlStatus

	HtmlNodeStatus func(
		HtmlNodePtr,
		int) HtmlStatus

	HtmlNewDoc func(
		URI string,
		ExternalID string) HtmlDocPtr

	HtmlNewDocNoDtD func(
		URI string,
		ExternalID string) HtmlDocPtr

	HtmlGetMetaEncoding func(
		doc HtmlDocPtr) string

	HtmlSetMetaEncoding func(
		doc HtmlDocPtr,
		encoding string) int

	HtmlDocDumpMemory func(
		cur XmlDocPtr,
		mem *string,
		size *int)

	HtmlDocDumpMemoryFormat func(
		cur XmlDocPtr,
		mem *string,
		size *int,
		format int)

	HtmlDocDump func(
		f *FILE,
		cur XmlDocPtr) int

	HtmlSaveFile func(
		filename string,
		cur XmlDocPtr) int

	HtmlNodeDump func(
		buf XmlBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr) int

	HtmlNodeDumpFile func(
		out *FILE,
		doc XmlDocPtr,
		cur XmlNodePtr)

	HtmlNodeDumpFileFormat func(
		out *FILE,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding string,
		format int) int

	HtmlSaveFileEnc func(
		filename string,
		cur XmlDocPtr,
		encoding string) int

	HtmlSaveFileFormat func(
		filename string,
		cur XmlDocPtr,
		encoding string,
		format int) int

	HtmlNodeDumpFormatOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding string,
		format int)

	HtmlDocContentDumpOutput func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding string)

	HtmlDocContentDumpFormatOutput func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding string,
		format int)

	HtmlNodeDumpOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding string)

	HtmlIsBooleanAttr func(
		name string) int

	XmlXPathFreeObject func(obj XmlXPathObjectPtr)

	XmlXPathNodeSetCreate func(val XmlNodePtr) XmlNodeSetPtr

	XmlXPathFreeNodeSetList func(obj XmlXPathObjectPtr)

	XmlXPathFreeNodeSet func(obj XmlNodeSetPtr)

	XmlXPathObjectCopy func(
		val XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPathCmpNodes func(node1, node2 XmlNodePtr) int

	XmlXPathCastNumberToBoolean func(val Double) int

	XmlXPathCastStringToBoolean func(val string) int

	XmlXPathCastNodeSetToBoolean func(ns XmlNodeSetPtr) int

	XmlXPathCastToBoolean func(val XmlXPathObjectPtr) int

	XmlXPathCastBooleanToNumber func(val int) Double

	XmlXPathCastStringToNumber func(val string) Double

	XmlXPathCastNodeToNumber func(node XmlNodePtr) Double

	XmlXPathCastNodeSetToNumber func(ns XmlNodeSetPtr) Double

	XmlXPathCastToNumber func(val XmlXPathObjectPtr) Double

	XmlXPathCastBooleanToString func(val int) string

	XmlXPathCastNumberToString func(val Double) string

	XmlXPathCastNodeToString func(node XmlNodePtr) string

	XmlXPathCastNodeSetToString func(ns XmlNodeSetPtr) string

	XmlXPathCastToString func(val XmlXPathObjectPtr) string

	XmlXPathConvertBoolean func(
		val XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPathConvertNumber func(
		val XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPathConvertString func(
		val XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPathNewContext func(doc XmlDocPtr) XmlXPathContextPtr

	XmlXPathFreeContext func(ctxt XmlXPathContextPtr)

	XmlXPathContextSetCache func(
		ctxt XmlXPathContextPtr, active, value, options int) int

	XmlXPathOrderDocElems func(doc XmlDocPtr) Long

	XmlXPathEval func(
		str string, ctx XmlXPathContextPtr) XmlXPathObjectPtr

	XmlXPathEvalExpression func(
		str string, ctxt XmlXPathContextPtr) XmlXPathObjectPtr

	XmlXPathEvalPredicate func(
		ctxt XmlXPathContextPtr, res XmlXPathObjectPtr) int

	XmlXPathCompile func(str string) XmlXPathCompExprPtr

	XmlXPathCtxtCompile func(ctxt XmlXPathContextPtr,
		str string) XmlXPathCompExprPtr

	XmlXPathCompiledEval func(comp XmlXPathCompExprPtr,
		ctx XmlXPathContextPtr) XmlXPathObjectPtr

	XmlXPathCompiledEvalToBoolean func(comp XmlXPathCompExprPtr,
		ctxt XmlXPathContextPtr) int

	XmlXPathFreeCompExpr func(comp XmlXPathCompExprPtr)

	XmlXPathInit func()

	XmlXPathIsNaN func(val Double) int

	XmlXPathIsInf func(val Double) int

	XmlXPathPopBoolean func(
		ctxt XmlXPathParserContextPtr) int

	XmlXPathPopNumber func(
		ctxt XmlXPathParserContextPtr) Double

	XmlXPathPopString func(
		ctxt XmlXPathParserContextPtr) string

	XmlXPathPopNodeSet func(
		ctxt XmlXPathParserContextPtr) XmlNodeSetPtr

	XmlXPathPopExternal func(
		ctxt XmlXPathParserContextPtr) *Void

	XmlXPathRegisterVariableLookup func(
		ctxt XmlXPathContextPtr,
		f XmlXPathVariableLookupFunc,
		data *Void)

	XmlXPathRegisterFuncLookup func(
		ctxt XmlXPathContextPtr,
		f XmlXPathFuncLookupFunc,
		funcCtxt *Void)

	XmlXPatherror func(
		ctxt XmlXPathParserContextPtr,
		file string,
		line int,
		no int)

	XmlXPathErr func(
		ctxt XmlXPathParserContextPtr,
		error int)

	XmlXPathDebugDumpObject func(
		output *FILE,
		cur XmlXPathObjectPtr,
		depth int)

	XmlXPathDebugDumpCompExpr func(
		output *FILE,
		comp XmlXPathCompExprPtr,
		depth int)

	XmlXPathNodeSetContains func(
		cur XmlNodeSetPtr,
		val XmlNodePtr) int

	XmlXPathDifference func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathIntersection func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathDistinctSorted func(
		nodes XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathDistinct func(
		nodes XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathHasSameNodes func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) int

	XmlXPathNodeLeadingSorted func(
		nodes XmlNodeSetPtr,
		node XmlNodePtr) XmlNodeSetPtr

	XmlXPathLeadingSorted func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathNodeLeading func(
		nodes XmlNodeSetPtr,
		node XmlNodePtr) XmlNodeSetPtr

	XmlXPathLeading func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathNodeTrailingSorted func(
		nodes XmlNodeSetPtr,
		node XmlNodePtr) XmlNodeSetPtr

	XmlXPathTrailingSorted func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathNodeTrailing func(
		nodes XmlNodeSetPtr,
		node XmlNodePtr) XmlNodeSetPtr

	XmlXPathTrailing func(
		nodes1 XmlNodeSetPtr,
		nodes2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathRegisterNs func(
		ctxt XmlXPathContextPtr,
		prefix string,
		ns_uri string) int

	XmlXPathNsLookup func(
		ctxt XmlXPathContextPtr,
		prefix string) string

	XmlXPathRegisteredNsCleanup func(
		ctxt XmlXPathContextPtr)

	XmlXPathRegisterFunc func(
		ctxt XmlXPathContextPtr,
		name string,
		f XmlXPathFunction) int

	XmlXPathRegisterFuncNS func(
		ctxt XmlXPathContextPtr,
		name string,
		ns_uri string,
		f XmlXPathFunction) int

	XmlXPathRegisterVariable func(
		ctxt XmlXPathContextPtr,
		name string,
		value XmlXPathObjectPtr) int

	XmlXPathRegisterVariableNS func(ctxt XmlXPathContextPtr,
		name, ns_uri string, value XmlXPathObjectPtr) int

	XmlXPathFunctionLookup func(
		ctxt XmlXPathContextPtr, name string) XmlXPathFunction

	XmlXPathFunctionLookupNS func(ctxt XmlXPathContextPtr,
		name, ns_uri string) XmlXPathFunction

	XmlXPathRegisteredFuncsCleanup func(ctxt XmlXPathContextPtr)

	XmlXPathVariableLookup func(
		ctxt XmlXPathContextPtr, name string) XmlXPathObjectPtr

	XmlXPathVariableLookupNS func(ctxt XmlXPathContextPtr,
		name string, ns_uri string) XmlXPathObjectPtr

	XmlXPathRegisteredVariablesCleanup func(
		ctxt XmlXPathContextPtr)

	XmlXPathNewParserContext func(str string,
		ctxt XmlXPathContextPtr) XmlXPathParserContextPtr

	XmlXPathFreeParserContext func(
		ctxt XmlXPathParserContextPtr)

	ValuePop func(
		ctxt XmlXPathParserContextPtr) XmlXPathObjectPtr

	ValuePush func(ctxt XmlXPathParserContextPtr,
		value XmlXPathObjectPtr) int

	XmlXPathNewString func(val string) XmlXPathObjectPtr

	XmlXPathNewCString func(val string) XmlXPathObjectPtr

	XmlXPathWrapString func(val string) XmlXPathObjectPtr

	XmlXPathWrapCString func(val string) XmlXPathObjectPtr

	XmlXPathNewFloat func(val Double) XmlXPathObjectPtr

	XmlXPathNewBoolean func(val int) XmlXPathObjectPtr

	XmlXPathNewNodeSet func(val XmlNodePtr) XmlXPathObjectPtr

	XmlXPathNewValueTree func(val XmlNodePtr) XmlXPathObjectPtr

	XmlXPathNodeSetAdd func(
		cur XmlNodeSetPtr, val XmlNodePtr) int

	XmlXPathNodeSetAddUnique func(
		cur XmlNodeSetPtr, val XmlNodePtr) int

	XmlXPathNodeSetAddNs func(
		cur XmlNodeSetPtr, node XmlNodePtr, ns XmlNsPtr) int

	XmlXPathNodeSetSort func(set XmlNodeSetPtr)

	XmlXPathRoot func(ctxt XmlXPathParserContextPtr)

	XmlXPathEvalExpr func(ctxt XmlXPathParserContextPtr)

	XmlXPathParseName func(
		ctxt XmlXPathParserContextPtr) string

	XmlXPathParseNCName func(
		ctxt XmlXPathParserContextPtr) string

	XmlXPathStringEvalNumber func(str string) Double

	XmlXPathEvaluatePredicateResult func(
		ctxt XmlXPathParserContextPtr, res XmlXPathObjectPtr) int

	XmlXPathRegisterAllFunctions func(ctxt XmlXPathContextPtr)

	XmlXPathNodeSetMerge func(
		val1, val2 XmlNodeSetPtr) XmlNodeSetPtr

	XmlXPathNodeSetDel func(cur XmlNodeSetPtr, val XmlNodePtr)

	XmlXPathNodeSetRemove func(cur XmlNodeSetPtr, val int)

	XmlXPathNewNodeSetList func(
		val XmlNodeSetPtr) XmlXPathObjectPtr

	XmlXPathWrapNodeSet func(
		val XmlNodeSetPtr) XmlXPathObjectPtr

	XmlXPathWrapExternal func(val *Void) XmlXPathObjectPtr

	XmlXPathEqualValues func(ctxt XmlXPathParserContextPtr) int

	XmlXPathNotEqualValues func(ctxt XmlXPathParserContextPtr) int

	XmlXPathCompareValues func(
		ctxt XmlXPathParserContextPtr, inf, strict int) int

	XmlXPathValueFlipSign func(ctxt XmlXPathParserContextPtr)

	XmlXPathAddValues func(ctxt XmlXPathParserContextPtr)

	XmlXPathSubValues func(ctxt XmlXPathParserContextPtr)

	XmlXPathMultValues func(ctxt XmlXPathParserContextPtr)

	XmlXPathDivValues func(ctxt XmlXPathParserContextPtr)

	XmlXPathModValues func(ctxt XmlXPathParserContextPtr)

	XmlXPathIsNodeType func(name string) int

	XmlXPathNextSelf func(
		ctxt XmlXPathParserContextPtr,
		cur XmlNodePtr) XmlNodePtr

	XmlXPathNextChild func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextDescendant func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextDescendantOrSelf func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextParent func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextAncestorOrSelf func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextFollowingSibling func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextFollowing func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextNamespace func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextAttribute func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextPreceding func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextAncestor func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathNextPrecedingSibling func(
		ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr

	XmlXPathLastFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathPositionFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathCountFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathIdFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathLocalNameFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathNamespaceURIFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathStringFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathStringLengthFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathConcatFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathContainsFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathStartsWithFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathSubstringFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathSubstringBeforeFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathSubstringAfterFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathNormalizeFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathTranslateFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathNotFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathTrueFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathFalseFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathLangFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathNumberFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathSumFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathFloorFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathCeilingFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathRoundFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathBooleanFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPathNodeSetFreeNs func(ns XmlNsPtr)

	XmlXIncludeProcess func(doc XmlDocPtr) int

	XmlXIncludeProcessFlags func(doc XmlDocPtr, flags int) int

	XmlXIncludeProcessFlagsData func(
		doc XmlDocPtr, flags int, data *Void) int

	XmlXIncludeProcessTreeFlagsData func(
		tree XmlNodePtr, flags int, data *Void) int

	XmlXIncludeProcessTree func(tree XmlNodePtr) int

	XmlXIncludeProcessTreeFlags func(
		tree XmlNodePtr, flags int) int

	XmlXIncludeNewContext func(doc XmlDocPtr) XmlXIncludeCtxtPtr

	XmlXIncludeSetFlags func(
		ctxt XmlXIncludeCtxtPtr, flags int) int

	XmlXIncludeFreeContext func(ctxt XmlXIncludeCtxtPtr)

	XmlXIncludeProcessNode func(
		ctxt XmlXIncludeCtxtPtr, tree XmlNodePtr) int

	XmlXPtrLocationSetCreate func(
		val XmlXPathObjectPtr) XmlLocationSetPtr

	XmlXPtrFreeLocationSet func(
		obj XmlLocationSetPtr)

	XmlXPtrLocationSetMerge func(
		val1 XmlLocationSetPtr,
		val2 XmlLocationSetPtr) XmlLocationSetPtr

	XmlXPtrNewRange func(
		start XmlNodePtr,
		startindex int,
		end XmlNodePtr,
		endindex int) XmlXPathObjectPtr

	XmlXPtrNewRangePoints func(
		start XmlXPathObjectPtr,
		end XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPtrNewRangeNodePoint func(
		start XmlNodePtr,
		end XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPtrNewRangePointNode func(
		start XmlXPathObjectPtr,
		end XmlNodePtr) XmlXPathObjectPtr

	XmlXPtrNewRangeNodes func(
		start XmlNodePtr,
		end XmlNodePtr) XmlXPathObjectPtr

	XmlXPtrNewLocationSetNodes func(
		start XmlNodePtr,
		end XmlNodePtr) XmlXPathObjectPtr

	XmlXPtrNewLocationSetNodeSet func(
		set XmlNodeSetPtr) XmlXPathObjectPtr

	XmlXPtrNewRangeNodeObject func(
		start XmlNodePtr,
		end XmlXPathObjectPtr) XmlXPathObjectPtr

	XmlXPtrNewCollapsedRange func(
		start XmlNodePtr) XmlXPathObjectPtr

	XmlXPtrLocationSetAdd func(
		cur XmlLocationSetPtr,
		val XmlXPathObjectPtr)

	XmlXPtrWrapLocationSet func(
		val XmlLocationSetPtr) XmlXPathObjectPtr

	XmlXPtrLocationSetDel func(
		cur XmlLocationSetPtr,
		val XmlXPathObjectPtr)

	XmlXPtrLocationSetRemove func(
		cur XmlLocationSetPtr,
		val int)

	XmlXPtrNewContext func(doc XmlDocPtr,
		here, origin XmlNodePtr) XmlXPathContextPtr

	XmlXPtrEval func(
		str string, ctx XmlXPathContextPtr) XmlXPathObjectPtr

	XmlXPtrRangeToFunction func(
		ctxt XmlXPathParserContextPtr, nargs int)

	XmlXPtrBuildNodeList func(obj XmlXPathObjectPtr) XmlNodePtr

	XmlXPtrEvalRangePredicate func(ctxt XmlXPathParserContextPtr)

	XmlUCSIsAegeanNumbers func(code int) int

	XmlUCSIsAlphabeticPresentationForms func(code int) int

	XmlUCSIsArabic func(code int) int

	XmlUCSIsArabicPresentationFormsA func(code int) int

	XmlUCSIsArabicPresentationFormsB func(code int) int

	XmlUCSIsArmenian func(code int) int

	XmlUCSIsArrows func(code int) int

	XmlUCSIsBasicLatin func(code int) int

	XmlUCSIsBengali func(code int) int

	XmlUCSIsBlockElements func(code int) int

	XmlUCSIsBopomofo func(code int) int

	XmlUCSIsBopomofoExtended func(code int) int

	XmlUCSIsBoxDrawing func(code int) int

	XmlUCSIsBraillePatterns func(code int) int

	XmlUCSIsBuhid func(code int) int

	XmlUCSIsByzantineMusicalSymbols func(code int) int

	XmlUCSIsCJKCompatibility func(code int) int

	XmlUCSIsCJKCompatibilityForms func(code int) int

	XmlUCSIsCJKCompatibilityIdeographs func(code int) int

	XmlUCSIsCJKCompatibilityIdeographsSupplement func(code int) int

	XmlUCSIsCJKRadicalsSupplement func(code int) int

	XmlUCSIsCJKSymbolsandPunctuation func(code int) int

	XmlUCSIsCJKUnifiedIdeographs func(code int) int

	XmlUCSIsCJKUnifiedIdeographsExtensionA func(code int) int

	XmlUCSIsCJKUnifiedIdeographsExtensionB func(code int) int

	XmlUCSIsCherokee func(code int) int

	XmlUCSIsCombiningDiacriticalMarks func(code int) int

	XmlUCSIsCombiningDiacriticalMarksforSymbols func(code int) int

	XmlUCSIsCombiningHalfMarks func(code int) int

	XmlUCSIsCombiningMarksforSymbols func(code int) int

	XmlUCSIsControlPictures func(code int) int

	XmlUCSIsCurrencySymbols func(code int) int

	XmlUCSIsCypriotSyllabary func(code int) int

	XmlUCSIsCyrillic func(code int) int

	XmlUCSIsCyrillicSupplement func(code int) int

	XmlUCSIsDeseret func(code int) int

	XmlUCSIsDevanagari func(code int) int

	XmlUCSIsDingbats func(code int) int

	XmlUCSIsEnclosedAlphanumerics func(code int) int

	XmlUCSIsEnclosedCJKLettersandMonths func(code int) int

	XmlUCSIsEthiopic func(code int) int

	XmlUCSIsGeneralPunctuation func(code int) int

	XmlUCSIsGeometricShapes func(code int) int

	XmlUCSIsGeorgian func(code int) int

	XmlUCSIsGothic func(code int) int

	XmlUCSIsGreek func(code int) int

	XmlUCSIsGreekExtended func(code int) int

	XmlUCSIsGreekandCoptic func(code int) int

	XmlUCSIsGujarati func(code int) int

	XmlUCSIsGurmukhi func(code int) int

	XmlUCSIsHalfwidthandFullwidthForms func(code int) int

	XmlUCSIsHangulCompatibilityJamo func(code int) int

	XmlUCSIsHangulJamo func(code int) int

	XmlUCSIsHangulSyllables func(code int) int

	XmlUCSIsHanunoo func(code int) int

	XmlUCSIsHebrew func(code int) int

	XmlUCSIsHighPrivateUseSurrogates func(code int) int

	XmlUCSIsHighSurrogates func(code int) int

	XmlUCSIsHiragana func(code int) int

	XmlUCSIsIPAExtensions func(code int) int

	XmlUCSIsIdeographicDescriptionCharacters func(code int) int

	XmlUCSIsKanbun func(code int) int

	XmlUCSIsKangxiRadicals func(code int) int

	XmlUCSIsKannada func(code int) int

	XmlUCSIsKatakana func(code int) int

	XmlUCSIsKatakanaPhoneticExtensions func(code int) int

	XmlUCSIsKhmer func(code int) int

	XmlUCSIsKhmerSymbols func(code int) int

	XmlUCSIsLao func(code int) int

	XmlUCSIsLatin1Supplement func(code int) int

	XmlUCSIsLatinExtendedA func(code int) int

	XmlUCSIsLatinExtendedB func(code int) int

	XmlUCSIsLatinExtendedAdditional func(code int) int

	XmlUCSIsLetterlikeSymbols func(code int) int

	XmlUCSIsLimbu func(code int) int

	XmlUCSIsLinearBIdeograms func(code int) int

	XmlUCSIsLinearBSyllabary func(code int) int

	XmlUCSIsLowSurrogates func(code int) int

	XmlUCSIsMalayalam func(code int) int

	XmlUCSIsMathematicalAlphanumericSymbols func(code int) int

	XmlUCSIsMathematicalOperators func(code int) int

	XmlUCSIsMiscellaneousMathematicalSymbolsA func(code int) int

	XmlUCSIsMiscellaneousMathematicalSymbolsB func(code int) int

	XmlUCSIsMiscellaneousSymbols func(code int) int

	XmlUCSIsMiscellaneousSymbolsandArrows func(code int) int

	XmlUCSIsMiscellaneousTechnical func(code int) int

	XmlUCSIsMongolian func(code int) int

	XmlUCSIsMusicalSymbols func(code int) int

	XmlUCSIsMyanmar func(code int) int

	XmlUCSIsNumberForms func(code int) int

	XmlUCSIsOgham func(code int) int

	XmlUCSIsOldItalic func(code int) int

	XmlUCSIsOpticalCharacterRecognition func(code int) int

	XmlUCSIsOriya func(code int) int

	XmlUCSIsOsmanya func(code int) int

	XmlUCSIsPhoneticExtensions func(code int) int

	XmlUCSIsPrivateUse func(code int) int

	XmlUCSIsPrivateUseArea func(code int) int

	XmlUCSIsRunic func(code int) int

	XmlUCSIsShavian func(code int) int

	XmlUCSIsSinhala func(code int) int

	XmlUCSIsSmallFormVariants func(code int) int

	XmlUCSIsSpacingModifierLetters func(code int) int

	XmlUCSIsSpecials func(code int) int

	XmlUCSIsSuperscriptsandSubscripts func(code int) int

	XmlUCSIsSupplementalArrowsA func(code int) int

	XmlUCSIsSupplementalArrowsB func(code int) int

	XmlUCSIsSupplementalMathematicalOperators func(code int) int

	XmlUCSIsSupplementaryPrivateUseAreaA func(code int) int

	XmlUCSIsSupplementaryPrivateUseAreaB func(code int) int

	XmlUCSIsSyriac func(code int) int

	XmlUCSIsTagalog func(code int) int

	XmlUCSIsTagbanwa func(code int) int

	XmlUCSIsTags func(code int) int

	XmlUCSIsTaiLe func(code int) int

	XmlUCSIsTaiXuanJingSymbols func(code int) int

	XmlUCSIsTamil func(code int) int

	XmlUCSIsTelugu func(code int) int

	XmlUCSIsThaana func(code int) int

	XmlUCSIsThai func(code int) int

	XmlUCSIsTibetan func(code int) int

	XmlUCSIsUgaritic func(code int) int

	XmlUCSIsUnifiedCanadianAboriginalSyllabics func(code int) int

	XmlUCSIsVariationSelectors func(code int) int

	XmlUCSIsVariationSelectorsSupplement func(code int) int

	XmlUCSIsYiRadicals func(code int) int

	XmlUCSIsYiSyllables func(code int) int

	XmlUCSIsYijingHexagramSymbols func(code int) int

	XmlUCSIsBlock func(code int, block string) int

	XmlUCSIsCatC func(code int) int

	XmlUCSIsCatCc func(code int) int

	XmlUCSIsCatCf func(code int) int

	XmlUCSIsCatCo func(code int) int

	XmlUCSIsCatCs func(code int) int

	XmlUCSIsCatL func(code int) int

	XmlUCSIsCatLl func(code int) int

	XmlUCSIsCatLm func(code int) int

	XmlUCSIsCatLo func(code int) int

	XmlUCSIsCatLt func(code int) int

	XmlUCSIsCatLu func(code int) int

	XmlUCSIsCatM func(code int) int

	XmlUCSIsCatMc func(code int) int

	XmlUCSIsCatMe func(code int) int

	XmlUCSIsCatMn func(code int) int

	XmlUCSIsCatN func(code int) int

	XmlUCSIsCatNd func(code int) int

	XmlUCSIsCatNl func(code int) int

	XmlUCSIsCatNo func(code int) int

	XmlUCSIsCatP func(code int) int

	XmlUCSIsCatPc func(code int) int

	XmlUCSIsCatPd func(code int) int

	XmlUCSIsCatPe func(code int) int

	XmlUCSIsCatPf func(code int) int

	XmlUCSIsCatPi func(code int) int

	XmlUCSIsCatPo func(code int) int

	XmlUCSIsCatPs func(code int) int

	XmlUCSIsCatS func(code int) int

	XmlUCSIsCatSc func(code int) int

	XmlUCSIsCatSk func(code int) int

	XmlUCSIsCatSm func(code int) int

	XmlUCSIsCatSo func(code int) int

	XmlUCSIsCatZ func(code int) int

	XmlUCSIsCatZl func(code int) int

	XmlUCSIsCatZp func(code int) int

	XmlUCSIsCatZs func(code int) int

	XmlUCSIsCat func(code int, cat string) int

	XmlSchemaNewParserCtxt func(
		URL string) XmlSchemaParserCtxtPtr

	XmlSchemaNewMemParserCtxt func(buffer string,
		size int) XmlSchemaParserCtxtPtr

	XmlSchemaNewDocParserCtxt func(
		doc XmlDocPtr) XmlSchemaParserCtxtPtr

	XmlSchemaFreeParserCtxt func(
		ctxt XmlSchemaParserCtxtPtr)

	XmlSchemaSetParserErrors func(
		ctxt XmlSchemaParserCtxtPtr,
		err XmlSchemaValidityErrorFunc,
		warn XmlSchemaValidityWarningFunc,
		ctx *Void)

	XmlSchemaSetParserStructuredErrors func(
		ctxt XmlSchemaParserCtxtPtr,
		serror XmlStructuredErrorFunc,
		ctx *Void)

	XmlSchemaGetParserErrors func(
		ctxt XmlSchemaParserCtxtPtr,
		err *XmlSchemaValidityErrorFunc,
		warn *XmlSchemaValidityWarningFunc,
		ctx **Void) int

	XmlSchemaIsValid func(ctxt XmlSchemaValidCtxtPtr) int

	XmlSchemaParse func(ctxt XmlSchemaParserCtxtPtr) XmlSchemaPtr

	XmlSchemaFree func(schema XmlSchemaPtr)

	XmlSchemaDump func(output *FILE, schema XmlSchemaPtr)

	XmlSchemaSetValidErrors func(
		ctxt XmlSchemaValidCtxtPtr,
		err XmlSchemaValidityErrorFunc,
		warn XmlSchemaValidityWarningFunc,
		ctx *Void)

	XmlSchemaSetValidStructuredErrors func(
		ctxt XmlSchemaValidCtxtPtr,
		serror XmlStructuredErrorFunc,
		ctx *Void)

	XmlSchemaGetValidErrors func(
		ctxt XmlSchemaValidCtxtPtr,
		err *XmlSchemaValidityErrorFunc,
		warn *XmlSchemaValidityWarningFunc,
		ctx **Void) int

	XmlSchemaSetValidOptions func(
		ctxt XmlSchemaValidCtxtPtr,
		options int) int

	XmlSchemaValidateSetFilename func(
		vctxt XmlSchemaValidCtxtPtr,
		filename string)

	XmlSchemaValidCtxtGetOptions func(
		ctxt XmlSchemaValidCtxtPtr) int

	XmlSchemaNewValidCtxt func(
		schema XmlSchemaPtr) XmlSchemaValidCtxtPtr

	XmlSchemaFreeValidCtxt func(
		ctxt XmlSchemaValidCtxtPtr)

	XmlSchemaValidateDoc func(
		ctxt XmlSchemaValidCtxtPtr,
		instance XmlDocPtr) int

	XmlSchemaValidateOneElement func(
		ctxt XmlSchemaValidCtxtPtr,
		elem XmlNodePtr) int

	XmlSchemaValidateStream func(
		ctxt XmlSchemaValidCtxtPtr,
		input XmlParserInputBufferPtr,
		enc XmlCharEncoding,
		sax XmlSAXHandlerPtr,
		user_data *Void) int

	XmlSchemaValidateFile func(
		ctxt XmlSchemaValidCtxtPtr,
		filename string,
		options int) int

	XmlSchemaValidCtxtGetParserCtxt func(
		ctxt XmlSchemaValidCtxtPtr) XmlParserCtxtPtr

	XmlSchemaSAXUnplug func(
		plug XmlSchemaSAXPlugPtr) int

	XmlSchemaValidateSetLocator func(
		vctxt XmlSchemaValidCtxtPtr,
		f XmlSchemaValidityLocatorFunc,
		ctxt *Void)

	XmlSchemaSAXPlug func(
		ctxt XmlSchemaValidCtxtPtr,
		sax *XmlSAXHandlerPtr,
		user_data **Void) XmlSchemaSAXPlugPtr

	XmlSchemaFreeType func(
		typ XmlSchemaTypePtr)

	XmlSchemaFreeWildcard func(
		wildcard XmlSchemaWildcardPtr)

	XmlSchemaInitTypes func()

	XmlSchemaCleanupTypes func()

	XmlSchemaGetPredefinedType func(
		name string,
		ns string) XmlSchemaTypePtr

	XmlSchemaValidatePredefinedType func(
		typ XmlSchemaTypePtr,
		value string,
		val *XmlSchemaValPtr) int

	XmlSchemaValPredefTypeNode func(
		typ XmlSchemaTypePtr,
		value string,
		val *XmlSchemaValPtr,
		node XmlNodePtr) int

	XmlSchemaValidateFacet func(
		base XmlSchemaTypePtr,
		facet XmlSchemaFacetPtr,
		value string,
		val XmlSchemaValPtr) int

	XmlSchemaValidateFacetWhtsp func(
		facet XmlSchemaFacetPtr,
		fws XmlSchemaWhitespaceValueType,
		valType XmlSchemaValType,
		value string,
		val XmlSchemaValPtr,
		ws XmlSchemaWhitespaceValueType) int

	XmlSchemaFreeValue func(
		val XmlSchemaValPtr)

	XmlSchemaNewFacet func() XmlSchemaFacetPtr

	XmlSchemaCheckFacet func(
		facet XmlSchemaFacetPtr,
		typeDecl XmlSchemaTypePtr,
		ctxt XmlSchemaParserCtxtPtr,
		name string) int

	XmlSchemaFreeFacet func(
		facet XmlSchemaFacetPtr)

	XmlSchemaCompareValues func(
		x XmlSchemaValPtr,
		y XmlSchemaValPtr) int

	XmlSchemaGetBuiltInListSimpleTypeItemType func(
		typ XmlSchemaTypePtr) XmlSchemaTypePtr

	XmlSchemaValidateListSimpleTypeFacet func(
		facet XmlSchemaFacetPtr,
		value string,
		actualLen Unsigned_long,
		expectedLen *Unsigned_long) int

	XmlSchemaGetBuiltInType func(
		typ XmlSchemaValType) XmlSchemaTypePtr

	XmlSchemaIsBuiltInTypeFacet func(
		typ XmlSchemaTypePtr,
		facetType int) int

	XmlSchemaCollapseString func(
		value string) string

	XmlSchemaWhiteSpaceReplace func(
		value string) string

	XmlSchemaGetFacetValueAsULong func(
		facet XmlSchemaFacetPtr) Unsigned_long

	XmlSchemaValidateLengthFacet func(
		typ XmlSchemaTypePtr,
		facet XmlSchemaFacetPtr,
		value string,
		val XmlSchemaValPtr,
		length *Unsigned_long) int

	XmlSchemaValidateLengthFacetWhtsp func(
		facet XmlSchemaFacetPtr,
		valType XmlSchemaValType,
		value string,
		val XmlSchemaValPtr,
		length *Unsigned_long,
		ws XmlSchemaWhitespaceValueType) int

	XmlSchemaValPredefTypeNodeNoNorm func(
		typ XmlSchemaTypePtr,
		value string,
		val *XmlSchemaValPtr,
		node XmlNodePtr) int

	XmlSchemaGetCanonValue func(
		val XmlSchemaValPtr,
		retValue *string) int

	XmlSchemaGetCanonValueWhtsp func(
		val XmlSchemaValPtr,
		retValue *string,
		ws XmlSchemaWhitespaceValueType) int

	XmlSchemaValueAppend func(
		prev XmlSchemaValPtr,
		cur XmlSchemaValPtr) int

	XmlSchemaValueGetNext func(
		cur XmlSchemaValPtr) XmlSchemaValPtr

	XmlSchemaValueGetAsString func(
		val XmlSchemaValPtr) string

	XmlSchemaValueGetAsBoolean func(
		val XmlSchemaValPtr) int

	XmlSchemaNewStringValue func(
		typ XmlSchemaValType,
		value string) XmlSchemaValPtr

	XmlSchemaNewNOTATIONValue func(
		name string,
		ns string) XmlSchemaValPtr

	XmlSchemaNewQNameValue func(
		namespaceName string,
		localName string) XmlSchemaValPtr

	XmlSchemaCompareValuesWhtsp func(
		x XmlSchemaValPtr,
		xws XmlSchemaWhitespaceValueType,
		y XmlSchemaValPtr,
		yws XmlSchemaWhitespaceValueType) int

	XmlSchemaCopyValue func(
		val XmlSchemaValPtr) XmlSchemaValPtr

	XmlSchemaGetValType func(
		val XmlSchemaValPtr) XmlSchemaValType

	XmlSchematronNewParserCtxt func(
		URL string) XmlSchematronParserCtxtPtr

	XmlSchematronNewMemParserCtxt func(
		buffer string,
		size int) XmlSchematronParserCtxtPtr

	XmlSchematronNewDocParserCtxt func(
		doc XmlDocPtr) XmlSchematronParserCtxtPtr

	XmlSchematronFreeParserCtxt func(
		ctxt XmlSchematronParserCtxtPtr)

	XmlSchematronParse func(
		ctxt XmlSchematronParserCtxtPtr) XmlSchematronPtr

	XmlSchematronFree func(
		schema XmlSchematronPtr)

	XmlSchematronSetValidStructuredErrors func(
		ctxt XmlSchematronValidCtxtPtr,
		serror XmlStructuredErrorFunc,
		ctx *Void)

	XmlSchematronNewValidCtxt func(
		schema XmlSchematronPtr,
		options int) XmlSchematronValidCtxtPtr

	XmlSchematronFreeValidCtxt func(
		ctxt XmlSchematronValidCtxtPtr)

	XmlSchematronValidateDoc func(
		ctxt XmlSchematronValidCtxtPtr,
		instance XmlDocPtr) int

	XmlCharInRange func(
		val Unsigned_int,
		group *XmlChRangeGroup) int

	XmlIsLetter func(
		c int) int

	XmlCreateFileParserCtxt func(
		filename string) XmlParserCtxtPtr

	XmlCreateURLParserCtxt func(
		filename string,
		options int) XmlParserCtxtPtr

	XmlCreateMemoryParserCtxt func(
		buffer string,
		size int) XmlParserCtxtPtr

	XmlCreateEntityParserCtxt func(
		URL string,
		ID string,
		base string) XmlParserCtxtPtr

	XmlSwitchEncoding func(
		ctxt XmlParserCtxtPtr,
		enc XmlCharEncoding) int

	XmlSwitchToEncoding func(
		ctxt XmlParserCtxtPtr,
		handler XmlCharEncodingHandlerPtr) int

	XmlSwitchInputEncoding func(
		ctxt XmlParserCtxtPtr,
		input XmlParserInputPtr,
		handler XmlCharEncodingHandlerPtr) int

	XmlNewStringInputStream func(
		ctxt XmlParserCtxtPtr,
		buffer string) XmlParserInputPtr

	XmlNewEntityInputStream func(
		ctxt XmlParserCtxtPtr,
		entity XmlEntityPtr) XmlParserInputPtr

	XmlPushInput func(
		ctxt XmlParserCtxtPtr,
		input XmlParserInputPtr) int

	XmlPopInput func(
		ctxt XmlParserCtxtPtr) XmlChar

	XmlFreeInputStream func(
		input XmlParserInputPtr)

	XmlNewInputFromFile func(
		ctxt XmlParserCtxtPtr,
		filename string) XmlParserInputPtr

	XmlNewInputStream func(
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	XmlSplitQName func(
		ctxt XmlParserCtxtPtr,
		name string,
		prefix *string) string

	XmlParseName func(
		ctxt XmlParserCtxtPtr) string

	XmlParseNmtoken func(
		ctxt XmlParserCtxtPtr) string

	XmlParseEntityValue func(
		ctxt XmlParserCtxtPtr,
		orig *string) string

	XmlParseAttValue func(
		ctxt XmlParserCtxtPtr) string

	XmlParseSystemLiteral func(
		ctxt XmlParserCtxtPtr) string

	XmlParsePubidLiteral func(
		ctxt XmlParserCtxtPtr) string

	XmlParseCharData func(
		ctxt XmlParserCtxtPtr,
		cdata int)

	XmlParseExternalID func(
		ctxt XmlParserCtxtPtr,
		publicID *string,
		strict int) string

	XmlParseComment func(
		ctxt XmlParserCtxtPtr)

	XmlParsePITarget func(
		ctxt XmlParserCtxtPtr) string

	XmlParsePI func(
		ctxt XmlParserCtxtPtr)

	XmlParseNotationDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseEntityDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseDefaultDecl func(
		ctxt XmlParserCtxtPtr,
		value *string) int

	XmlParseNotationType func(
		ctxt XmlParserCtxtPtr) XmlEnumerationPtr

	XmlParseEnumerationType func(
		ctxt XmlParserCtxtPtr) XmlEnumerationPtr

	XmlParseEnumeratedType func(
		ctxt XmlParserCtxtPtr,
		tree *XmlEnumerationPtr) int

	XmlParseAttributeType func(
		ctxt XmlParserCtxtPtr,
		tree *XmlEnumerationPtr) int

	XmlParseAttributeListDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseElementMixedContentDecl func(
		ctxt XmlParserCtxtPtr,
		inputchk int) XmlElementContentPtr

	XmlParseElementChildrenContentDecl func(
		ctxt XmlParserCtxtPtr,
		inputchk int) XmlElementContentPtr

	XmlParseElementContentDecl func(
		ctxt XmlParserCtxtPtr,
		name string,
		result *XmlElementContentPtr) int

	XmlParseElementDecl func(
		ctxt XmlParserCtxtPtr) int

	XmlParseMarkupDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseCharRef func(
		ctxt XmlParserCtxtPtr) int

	XmlParseEntityRef func(
		ctxt XmlParserCtxtPtr) XmlEntityPtr

	XmlParseReference func(
		ctxt XmlParserCtxtPtr)

	XmlParsePEReference func(
		ctxt XmlParserCtxtPtr)

	XmlParseDocTypeDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseAttribute func(
		ctxt XmlParserCtxtPtr,
		value *string) string

	XmlParseStartTag func(
		ctxt XmlParserCtxtPtr) string

	XmlParseEndTag func(
		ctxt XmlParserCtxtPtr)

	XmlParseCDSect func(
		ctxt XmlParserCtxtPtr)

	XmlParseContent func(
		ctxt XmlParserCtxtPtr)

	XmlParseElement func(
		ctxt XmlParserCtxtPtr)

	XmlParseVersionNum func(
		ctxt XmlParserCtxtPtr) string

	XmlParseVersionInfo func(
		ctxt XmlParserCtxtPtr) string

	XmlParseEncName func(
		ctxt XmlParserCtxtPtr) string

	XmlParseEncodingDecl func(
		ctxt XmlParserCtxtPtr) string

	XmlParseSDDecl func(
		ctxt XmlParserCtxtPtr) int

	XmlParseXMLDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseTextDecl func(
		ctxt XmlParserCtxtPtr)

	XmlParseMisc func(
		ctxt XmlParserCtxtPtr)

	XmlParseExternalSubset func(
		ctxt XmlParserCtxtPtr,
		ExternalID string,
		SystemID string)

	XmlStringDecodeEntities func(
		ctxt XmlParserCtxtPtr,
		str string,
		what int,
		end XmlChar,
		end2 XmlChar,
		end3 XmlChar) string

	XmlStringLenDecodeEntities func(
		ctxt XmlParserCtxtPtr,
		str string,
		len int,
		what int,
		end XmlChar,
		end2 XmlChar,
		end3 XmlChar) string

	NodePush func(
		ctxt XmlParserCtxtPtr,
		value XmlNodePtr) int

	NodePop func(
		ctxt XmlParserCtxtPtr) XmlNodePtr

	InputPush func(
		ctxt XmlParserCtxtPtr,
		value XmlParserInputPtr) int

	InputPop func(
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	NamePop func(
		ctxt XmlParserCtxtPtr) string

	NamePush func(
		ctxt XmlParserCtxtPtr,
		value string) int

	XmlSkipBlankChars func(
		ctxt XmlParserCtxtPtr) int

	XmlStringCurrentChar func(
		ctxt XmlParserCtxtPtr,
		cur string,
		len *int) int

	XmlParserHandlePEReference func(
		ctxt XmlParserCtxtPtr)

	XmlCheckLanguageID func(
		lang string) int

	XmlCurrentChar func(
		ctxt XmlParserCtxtPtr,
		len *int) int

	XmlCopyCharMultiByte func(
		out string,
		val int) int

	XmlCopyChar func(
		len int,
		out string,
		val int) int

	XmlNextChar func(
		ctxt XmlParserCtxtPtr)

	XmlParserInputShrink func(
		in XmlParserInputPtr)

	HtmlInitAutoClose func()

	HtmlCreateFileParserCtxt func(
		filename, encoding string) HtmlParserCtxtPtr

	XmlSetEntityReferenceFunc func(
		f XmlEntityReferenceFunc)

	XmlParseQuotedString func(
		ctxt XmlParserCtxtPtr) string

	XmlParseNamespace func(
		ctxt XmlParserCtxtPtr)

	XmlNamespaceParseNSDef func(
		ctxt XmlParserCtxtPtr) string

	XmlScanName func(
		ctxt XmlParserCtxtPtr) string

	XmlNamespaceParseNCName func(
		ctxt XmlParserCtxtPtr) string

	XmlParserHandleReference func(
		ctxt XmlParserCtxtPtr)

	XmlNamespaceParseQName func(
		ctxt XmlParserCtxtPtr,
		prefix *string) string

	XmlDecodeEntities func(
		ctxt XmlParserCtxtPtr,
		len int,
		what int,
		end XmlChar,
		end2 XmlChar,
		end3 XmlChar) string

	XmlHandleEntity func(
		ctxt XmlParserCtxtPtr,
		entity XmlEntityPtr)

	XmlRelaxNGInitTypes func(
		Void) int

	XmlRelaxNGCleanupTypes func(
		Void)

	XmlRelaxNGNewParserCtxt func(
		URL string) XmlRelaxNGParserCtxtPtr

	XmlRelaxNGNewMemParserCtxt func(
		buffer string,
		size int) XmlRelaxNGParserCtxtPtr

	XmlRelaxNGNewDocParserCtxt func(
		doc XmlDocPtr) XmlRelaxNGParserCtxtPtr

	XmlRelaxParserSetFlag func(
		ctxt XmlRelaxNGParserCtxtPtr,
		flag int) int

	XmlRelaxNGFreeParserCtxt func(
		ctxt XmlRelaxNGParserCtxtPtr)

	XmlRelaxNGSetParserErrors func(
		ctxt XmlRelaxNGParserCtxtPtr,
		err XmlRelaxNGValidityErrorFunc,
		warn XmlRelaxNGValidityWarningFunc,
		ctx *Void)

	XmlRelaxNGGetParserErrors func(
		ctxt XmlRelaxNGParserCtxtPtr,
		err *XmlRelaxNGValidityErrorFunc,
		warn *XmlRelaxNGValidityWarningFunc,
		ctx **Void) int

	XmlRelaxNGSetParserStructuredErrors func(
		ctxt XmlRelaxNGParserCtxtPtr,
		serror XmlStructuredErrorFunc,
		ctx *Void)

	XmlRelaxNGParse func(
		ctxt XmlRelaxNGParserCtxtPtr) XmlRelaxNGPtr

	XmlRelaxNGFree func(
		schema XmlRelaxNGPtr)

	XmlRelaxNGDump func(
		output *FILE,
		schema XmlRelaxNGPtr)

	XmlRelaxNGDumpTree func(
		output *FILE,
		schema XmlRelaxNGPtr)

	XmlRelaxNGSetValidErrors func(
		ctxt XmlRelaxNGValidCtxtPtr,
		err XmlRelaxNGValidityErrorFunc,
		warn XmlRelaxNGValidityWarningFunc,
		ctx *Void)

	XmlRelaxNGGetValidErrors func(
		ctxt XmlRelaxNGValidCtxtPtr,
		err *XmlRelaxNGValidityErrorFunc,
		warn *XmlRelaxNGValidityWarningFunc,
		ctx **Void) int

	XmlRelaxNGSetValidStructuredErrors func(
		ctxt XmlRelaxNGValidCtxtPtr,
		serror XmlStructuredErrorFunc,
		ctx *Void)

	XmlRelaxNGNewValidCtxt func(
		schema XmlRelaxNGPtr) XmlRelaxNGValidCtxtPtr

	XmlRelaxNGFreeValidCtxt func(
		ctxt XmlRelaxNGValidCtxtPtr)

	XmlRelaxNGValidateDoc func(
		ctxt XmlRelaxNGValidCtxtPtr,
		doc XmlDocPtr) int

	XmlRelaxNGValidatePushElement func(
		ctxt XmlRelaxNGValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlRelaxNGValidatePushCData func(
		ctxt XmlRelaxNGValidCtxtPtr,
		data string,
		len int) int

	XmlRelaxNGValidatePopElement func(
		ctxt XmlRelaxNGValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlRelaxNGValidateFullElement func(
		ctxt XmlRelaxNGValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlNewTextReader func(
		input XmlParserInputBufferPtr,
		URI string) XmlTextReaderPtr

	XmlNewTextReaderFilename func(
		URI string) XmlTextReaderPtr

	XmlFreeTextReader func(
		reader XmlTextReaderPtr)

	XmlTextReaderSetup func(
		reader XmlTextReaderPtr,
		input XmlParserInputBufferPtr,
		URL string,
		encoding string,
		options int) int

	XmlTextReaderRead func(
		reader XmlTextReaderPtr) int

	XmlTextReaderReadInnerXml func(
		reader XmlTextReaderPtr) string

	XmlTextReaderReadOuterXml func(
		reader XmlTextReaderPtr) string

	XmlTextReaderReadString func(
		reader XmlTextReaderPtr) string

	XmlTextReaderReadAttributeValue func(
		reader XmlTextReaderPtr) int

	XmlTextReaderAttributeCount func(
		reader XmlTextReaderPtr) int

	XmlTextReaderDepth func(
		reader XmlTextReaderPtr) int

	XmlTextReaderHasAttributes func(
		reader XmlTextReaderPtr) int

	XmlTextReaderHasValue func(
		reader XmlTextReaderPtr) int

	XmlTextReaderIsDefault func(
		reader XmlTextReaderPtr) int

	XmlTextReaderIsEmptyElement func(
		reader XmlTextReaderPtr) int

	XmlTextReaderNodeType func(
		reader XmlTextReaderPtr) int

	XmlTextReaderQuoteChar func(
		reader XmlTextReaderPtr) int

	XmlTextReaderReadState func(
		reader XmlTextReaderPtr) int

	XmlTextReaderIsNamespaceDecl func(
		reader XmlTextReaderPtr) int

	XmlTextReaderConstBaseUri func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstLocalName func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstName func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstNamespaceUri func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstPrefix func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstXmlLang func(
		reader XmlTextReaderPtr) string

	XmlTextReaderConstString func(
		reader XmlTextReaderPtr,
		str string) string

	XmlTextReaderConstValue func(
		reader XmlTextReaderPtr) string

	XmlTextReaderBaseUri func(
		reader XmlTextReaderPtr) string

	XmlTextReaderLocalName func(
		reader XmlTextReaderPtr) string

	XmlTextReaderName func(
		reader XmlTextReaderPtr) string

	XmlTextReaderNamespaceUri func(
		reader XmlTextReaderPtr) string

	XmlTextReaderPrefix func(
		reader XmlTextReaderPtr) string

	XmlTextReaderXmlLang func(
		reader XmlTextReaderPtr) string

	XmlTextReaderValue func(
		reader XmlTextReaderPtr) string

	XmlTextReaderClose func(
		reader XmlTextReaderPtr) int

	XmlTextReaderGetAttributeNo func(
		reader XmlTextReaderPtr,
		no int) string

	XmlTextReaderGetAttribute func(
		reader XmlTextReaderPtr,
		name string) string

	XmlTextReaderGetAttributeNs func(
		reader XmlTextReaderPtr,
		localName string,
		namespaceURI string) string

	XmlTextReaderGetRemainder func(
		reader XmlTextReaderPtr) XmlParserInputBufferPtr

	XmlTextReaderLookupNamespace func(
		reader XmlTextReaderPtr,
		prefix string) string

	XmlTextReaderMoveToAttributeNo func(
		reader XmlTextReaderPtr,
		no int) int

	XmlTextReaderMoveToAttribute func(
		reader XmlTextReaderPtr,
		name string) int

	XmlTextReaderMoveToAttributeNs func(
		reader XmlTextReaderPtr,
		localName string,
		namespaceURI string) int

	XmlTextReaderMoveToFirstAttribute func(
		reader XmlTextReaderPtr) int

	XmlTextReaderMoveToNextAttribute func(
		reader XmlTextReaderPtr) int

	XmlTextReaderMoveToElement func(
		reader XmlTextReaderPtr) int

	XmlTextReaderNormalization func(
		reader XmlTextReaderPtr) int

	XmlTextReaderConstEncoding func(
		reader XmlTextReaderPtr) string

	XmlTextReaderSetParserProp func(
		reader XmlTextReaderPtr,
		prop int,
		value int) int

	XmlTextReaderGetParserProp func(
		reader XmlTextReaderPtr,
		prop int) int

	XmlTextReaderCurrentNode func(
		reader XmlTextReaderPtr) XmlNodePtr

	XmlTextReaderGetParserLineNumber func(
		reader XmlTextReaderPtr) int

	XmlTextReaderGetParserColumnNumber func(
		reader XmlTextReaderPtr) int

	XmlTextReaderPreserve func(
		reader XmlTextReaderPtr) XmlNodePtr

	XmlTextReaderPreservePattern func(
		reader XmlTextReaderPtr,
		pattern string,
		namespaces *string) int

	XmlTextReaderCurrentDoc func(
		reader XmlTextReaderPtr) XmlDocPtr

	XmlTextReaderExpand func(
		reader XmlTextReaderPtr) XmlNodePtr

	XmlTextReaderNext func(
		reader XmlTextReaderPtr) int

	XmlTextReaderNextSibling func(
		reader XmlTextReaderPtr) int

	XmlTextReaderIsValid func(
		reader XmlTextReaderPtr) int

	XmlTextReaderRelaxNGValidate func(
		reader XmlTextReaderPtr,
		rng string) int

	XmlTextReaderRelaxNGValidateCtxt func(
		reader XmlTextReaderPtr,
		ctxt XmlRelaxNGValidCtxtPtr,
		options int) int

	XmlTextReaderRelaxNGSetSchema func(
		reader XmlTextReaderPtr,
		schema XmlRelaxNGPtr) int

	XmlTextReaderSchemaValidate func(
		reader XmlTextReaderPtr,
		xsd string) int

	XmlTextReaderSchemaValidateCtxt func(
		reader XmlTextReaderPtr,
		ctxt XmlSchemaValidCtxtPtr,
		options int) int

	XmlTextReaderSetSchema func(
		reader XmlTextReaderPtr,
		schema XmlSchemaPtr) int

	XmlTextReaderConstXmlVersion func(
		reader XmlTextReaderPtr) string

	XmlTextReaderStandalone func(
		reader XmlTextReaderPtr) int

	XmlTextReaderByteConsumed func(
		reader XmlTextReaderPtr) Long

	XmlReaderWalker func(
		doc XmlDocPtr) XmlTextReaderPtr

	XmlReaderForDoc func(
		cur string,
		URL string,
		encoding string,
		options int) XmlTextReaderPtr

	XmlReaderForFile func(
		filename string,
		encoding string,
		options int) XmlTextReaderPtr

	XmlReaderForMemory func(
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) XmlTextReaderPtr

	XmlReaderForFd func(
		fd int,
		URL string,
		encoding string,
		options int) XmlTextReaderPtr

	XmlReaderForIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL string,
		encoding string,
		options int) XmlTextReaderPtr

	XmlReaderNewWalker func(
		reader XmlTextReaderPtr,
		doc XmlDocPtr) int

	XmlReaderNewDoc func(
		reader XmlTextReaderPtr,
		cur string,
		URL string,
		encoding string,
		options int) int

	XmlReaderNewFile func(
		reader XmlTextReaderPtr,
		filename string,
		encoding string,
		options int) int

	XmlReaderNewMemory func(
		reader XmlTextReaderPtr,
		buffer string,
		size int,
		URL string,
		encoding string,
		options int) int

	XmlReaderNewFd func(
		reader XmlTextReaderPtr,
		fd int,
		URL string,
		encoding string,
		options int) int

	XmlReaderNewIO func(
		reader XmlTextReaderPtr,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL string,
		encoding string,
		options int) int

	XmlTextReaderLocatorLineNumber func(
		locator XmlTextReaderLocatorPtr) int

	XmlTextReaderLocatorBaseURI func(
		locator XmlTextReaderLocatorPtr) string

	XmlTextReaderSetErrorHandler func(
		reader XmlTextReaderPtr,
		f XmlTextReaderErrorFunc,
		arg *Void)

	XmlTextReaderSetStructuredErrorHandler func(
		reader XmlTextReaderPtr,
		f XmlStructuredErrorFunc,
		arg *Void)

	XmlTextReaderGetErrorHandler func(
		reader XmlTextReaderPtr,
		f *XmlTextReaderErrorFunc,
		arg **Void)

	XmlNewCatalog func(
		sgml int) XmlCatalogPtr

	XmlLoadACatalog func(
		filename string) XmlCatalogPtr

	XmlLoadSGMLSuperCatalog func(
		filename string) XmlCatalogPtr

	XmlConvertSGMLCatalog func(
		catal XmlCatalogPtr) int

	XmlACatalogAdd func(
		catal XmlCatalogPtr,
		typ string,
		orig string,
		replace string) int

	XmlACatalogRemove func(
		catal XmlCatalogPtr,
		value string) int

	XmlACatalogResolve func(
		catal XmlCatalogPtr,
		pubID string,
		sysID string) string

	XmlACatalogResolveSystem func(
		catal XmlCatalogPtr,
		sysID string) string

	XmlACatalogResolvePublic func(
		catal XmlCatalogPtr,
		pubID string) string

	XmlACatalogResolveURI func(
		catal XmlCatalogPtr,
		URI string) string

	XmlACatalogDump func(
		catal XmlCatalogPtr,
		out *FILE)

	XmlFreeCatalog func(
		catal XmlCatalogPtr)

	XmlCatalogIsEmpty func(
		catal XmlCatalogPtr) int

	XmlInitializeCatalog func(
		Void)

	XmlLoadCatalog func(
		filename string) int

	XmlLoadCatalogs func(
		paths string)

	XmlCatalogCleanup func(
		Void)

	XmlCatalogDump func(
		out *FILE)

	XmlCatalogResolve func(
		pubID string,
		sysID string) string

	XmlCatalogResolveSystem func(
		sysID string) string

	XmlCatalogResolvePublic func(
		pubID string) string

	XmlCatalogResolveURI func(
		URI string) string

	XmlCatalogAdd func(
		typ string,
		orig string,
		replace string) int

	XmlCatalogRemove func(
		value string) int

	XmlParseCatalogFile func(
		filename string) XmlDocPtr

	XmlCatalogConvert func(
		Void) int

	XmlCatalogFreeLocal func(
		catalogs *Void)

	XmlCatalogAddLocal func(
		catalogs *Void,
		URL string) *Void

	XmlCatalogLocalResolve func(
		catalogs *Void,
		pubID string,
		sysID string) string

	XmlCatalogLocalResolveURI func(
		catalogs *Void,
		URI string) string

	XmlCatalogSetDebug func(
		level int) int

	XmlCatalogSetDefaultPrefer func(
		prefer XmlCatalogPrefer) XmlCatalogPrefer

	XmlCatalogSetDefaults func(
		allow XmlCatalogAllow)

	XmlCatalogGetDefaults func(
		Void) XmlCatalogAllow

	XmlCatalogGetSystem func(
		sysID string) string

	XmlCatalogGetPublic func(
		pubID string) string

	XmlDebugDumpString func(
		output *FILE,
		str string)

	XmlDebugDumpAttr func(
		output *FILE,
		attr XmlAttrPtr,
		depth int)

	XmlDebugDumpAttrList func(
		output *FILE,
		attr XmlAttrPtr,
		depth int)

	XmlDebugDumpOneNode func(
		output *FILE,
		node XmlNodePtr,
		depth int)

	XmlDebugDumpNode func(
		output *FILE,
		node XmlNodePtr,
		depth int)

	XmlDebugDumpNodeList func(
		output *FILE,
		node XmlNodePtr,
		depth int)

	XmlDebugDumpDocumentHead func(
		output *FILE,
		doc XmlDocPtr)

	XmlDebugDumpDocument func(
		output *FILE,
		doc XmlDocPtr)

	XmlDebugDumpDTD func(
		output *FILE,
		dtd XmlDtdPtr)

	XmlDebugDumpEntities func(
		output *FILE,
		doc XmlDocPtr)

	XmlDebugCheckDocument func(
		output *FILE,
		doc XmlDocPtr) int

	XmlLsOneNode func(
		output *FILE,
		node XmlNodePtr)

	XmlLsCountNode func(
		node XmlNodePtr) int

	XmlBoolToText func(
		boolval int) string

	XmlShellPrintXPathError func(
		errorType int,
		arg string)

	XmlShellPrintXPathResult func(
		list XmlXPathObjectPtr)

	XmlShellList func(
		ctxt XmlShellCtxtPtr,
		arg string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellBase func(
		ctxt XmlShellCtxtPtr,
		arg string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellDir func(
		ctxt XmlShellCtxtPtr,
		arg string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellLoad func(
		ctxt XmlShellCtxtPtr,
		filename string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellPrintNode func(
		node XmlNodePtr)

	XmlShellCat func(
		ctxt XmlShellCtxtPtr,
		arg string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellWrite func(
		ctxt XmlShellCtxtPtr,
		filename string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellSave func(
		ctxt XmlShellCtxtPtr,
		filename string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellValidate func(
		ctxt XmlShellCtxtPtr,
		dtd string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellDu func(
		ctxt XmlShellCtxtPtr,
		arg string,
		tree XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShellPwd func(
		ctxt XmlShellCtxtPtr,
		buffer string,
		node XmlNodePtr,
		node2 XmlNodePtr) int

	XmlShell func(
		doc XmlDocPtr,
		filename string,
		input XmlShellReadlineFunc,
		output *FILE)

	XmlNewTextWriter func(
		out XmlOutputBufferPtr) XmlTextWriterPtr

	XmlNewTextWriterFilename func(
		uri string,
		compression int) XmlTextWriterPtr

	XmlNewTextWriterMemory func(
		buf XmlBufferPtr,
		compression int) XmlTextWriterPtr

	XmlNewTextWriterPushParser func(
		ctxt XmlParserCtxtPtr,
		compression int) XmlTextWriterPtr

	XmlNewTextWriterDoc func(
		doc *XmlDocPtr,
		compression int) XmlTextWriterPtr

	XmlNewTextWriterTree func(
		doc XmlDocPtr,
		node XmlNodePtr,
		compression int) XmlTextWriterPtr

	XmlFreeTextWriter func(
		writer XmlTextWriterPtr)

	XmlTextWriterStartDocument func(
		writer XmlTextWriterPtr,
		version string,
		encoding string,
		standalone string) int

	XmlTextWriterEndDocument func(
		writer XmlTextWriterPtr) int

	XmlTextWriterStartComment func(
		writer XmlTextWriterPtr) int

	XmlTextWriterEndComment func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatComment func(
		writer XmlTextWriterPtr,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatComment func(
		writer XmlTextWriterPtr,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteComment func(
		writer XmlTextWriterPtr,
		Content string) int

	XmlTextWriterStartElement func(
		writer XmlTextWriterPtr,
		name string) int

	XmlTextWriterStartElementNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string) int

	XmlTextWriterEndElement func(
		writer XmlTextWriterPtr) int

	XmlTextWriterFullEndElement func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatElement func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatElement func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteElement func(
		writer XmlTextWriterPtr,
		name string,
		Content string) int

	XmlTextWriterWriteFormatElementNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatElementNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteElementNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		Content string) int

	XmlTextWriterWriteFormatRaw func(
		writer XmlTextWriterPtr,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatRaw func(
		writer XmlTextWriterPtr,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteRawLen func(
		writer XmlTextWriterPtr,
		content string,
		len int) int

	XmlTextWriterWriteRaw func(
		writer XmlTextWriterPtr,
		content string) int

	XmlTextWriterWriteFormatString func(
		writer XmlTextWriterPtr,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatString func(
		writer XmlTextWriterPtr,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteString func(
		writer XmlTextWriterPtr,
		Content string) int

	XmlTextWriterWriteBase64 func(
		writer XmlTextWriterPtr,
		data string,
		start int,
		len int) int

	XmlTextWriterWriteBinHex func(
		writer XmlTextWriterPtr,
		data string,
		start int,
		len int) int

	XmlTextWriterStartAttribute func(
		writer XmlTextWriterPtr,
		name string) int

	XmlTextWriterStartAttributeNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string) int

	XmlTextWriterEndAttribute func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatAttribute func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatAttribute func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteAttribute func(
		writer XmlTextWriterPtr,
		name string,
		Content string) int

	XmlTextWriterWriteFormatAttributeNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatAttributeNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteAttributeNS func(
		writer XmlTextWriterPtr,
		prefix string,
		name string,
		namespaceURI string,
		Content string) int

	XmlTextWriterStartPI func(
		writer XmlTextWriterPtr,
		target string) int

	XmlTextWriterEndPI func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatPI func(
		writer XmlTextWriterPtr,
		target string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatPI func(
		writer XmlTextWriterPtr,
		target string,
		format string,
		argptr Va_list) int

	XmlTextWriterWritePI func(
		writer XmlTextWriterPtr,
		target string,
		content string) int

	XmlTextWriterStartCDATA func(
		writer XmlTextWriterPtr) int

	XmlTextWriterEndCDATA func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatCDATA func(
		writer XmlTextWriterPtr,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatCDATA func(
		writer XmlTextWriterPtr,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteCDATA func(
		writer XmlTextWriterPtr,
		content string) int

	XmlTextWriterStartDTD func(
		writer XmlTextWriterPtr,
		name string,
		pubid string,
		sysid string) int

	XmlTextWriterEndDTD func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatDTD func(
		writer XmlTextWriterPtr,
		name string,
		pubid string,
		sysid string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatDTD func(
		writer XmlTextWriterPtr,
		name string,
		pubid string,
		sysid string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteDTD func(
		writer XmlTextWriterPtr,
		name string,
		pubid string,
		sysid string,
		subset string) int

	XmlTextWriterStartDTDElement func(
		writer XmlTextWriterPtr,
		name string) int

	XmlTextWriterEndDTDElement func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatDTDElement func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatDTDElement func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteDTDElement func(
		writer XmlTextWriterPtr,
		name string,
		Content string) int

	XmlTextWriterStartDTDAttlist func(
		writer XmlTextWriterPtr,
		name string) int

	XmlTextWriterEndDTDAttlist func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatDTDAttlist func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatDTDAttlist func(
		writer XmlTextWriterPtr,
		name string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteDTDAttlist func(
		writer XmlTextWriterPtr,
		name string,
		Content string) int

	XmlTextWriterStartDTDEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string) int

	XmlTextWriterEndDTDEntity func(
		writer XmlTextWriterPtr) int

	XmlTextWriterWriteFormatDTDInternalEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string,
		format string,
		v ...VArg) int

	XmlTextWriterWriteVFormatDTDInternalEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string,
		format string,
		argptr Va_list) int

	XmlTextWriterWriteDTDInternalEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string,
		content string) int

	XmlTextWriterWriteDTDExternalEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string,
		pubid string,
		sysid string,
		ndataid string) int

	XmlTextWriterWriteDTDExternalEntityContents func(
		writer XmlTextWriterPtr,
		pubid string,
		sysid string,
		ndataid string) int

	XmlTextWriterWriteDTDEntity func(
		writer XmlTextWriterPtr,
		pe int,
		name string,
		pubid string,
		sysid string,
		ndataid string,
		content string) int

	XmlTextWriterWriteDTDNotation func(
		writer XmlTextWriterPtr,
		name string,
		pubid string,
		sysid string) int

	XmlTextWriterSetIndent func(
		writer XmlTextWriterPtr,
		indent int) int

	XmlTextWriterSetIndentString func(
		writer XmlTextWriterPtr,
		str string) int

	XmlTextWriterSetQuoteChar func(
		writer XmlTextWriterPtr,
		quotechar XmlChar) int

	XmlTextWriterFlush func(
		writer XmlTextWriterPtr) int

	XmlNanoFTPInit func()

	XmlNanoFTPCleanup func()

	XmlNanoFTPNewCtxt func(
		URL string) *Void

	XmlNanoFTPFreeCtxt func(
		ctx *Void)

	XmlNanoFTPConnectTo func(
		server string,
		port int) *Void

	XmlNanoFTPOpen func(
		URL string) *Void

	XmlNanoFTPConnect func(
		ctx *Void) int

	XmlNanoFTPClose func(
		ctx *Void) int

	XmlNanoFTPQuit func(
		ctx *Void) int

	XmlNanoFTPScanProxy func(
		URL string)

	XmlNanoFTPProxy func(
		host string,
		port int,
		user string,
		passwd string,
		typ int)

	XmlNanoFTPUpdateURL func(
		ctx *Void,
		URL string) int

	XmlNanoFTPGetResponse func(
		ctx *Void) int

	XmlNanoFTPCheckResponse func(
		ctx *Void) int

	XmlNanoFTPCwd func(
		ctx *Void,
		directory string) int

	XmlNanoFTPDele func(
		ctx *Void,
		file string) int

	XmlNanoFTPGetConnection func(
		ctx *Void) SOCKET

	XmlNanoFTPCloseConnection func(
		ctx *Void) int

	XmlNanoFTPList func(
		ctx *Void,
		callback FtpListCallback,
		userData *Void,
		filename string) int

	XmlNanoFTPGetSocket func(
		ctx *Void,
		filename string) SOCKET

	XmlNanoFTPGet func(
		ctx *Void,
		callback FtpDataCallback,
		userData *Void,
		filename string) int

	XmlNanoFTPRead func(
		ctx *Void,
		dest *Void,
		len int) int

	XmlNanoHTTPInit func(
		Void)

	XmlNanoHTTPCleanup func(
		Void)

	XmlNanoHTTPScanProxy func(
		URL string)

	XmlNanoHTTPFetch func(
		URL string,
		filename string,
		contentType *string) int

	XmlNanoHTTPMethod func(
		URL string,
		method string,
		input string,
		contentType *string,
		headers string,
		ilen int) *Void

	XmlNanoHTTPMethodRedir func(
		URL string,
		method string,
		input string,
		contentType *string,
		redir *string,
		headers string,
		ilen int) *Void

	XmlNanoHTTPOpen func(
		URL string,
		contentType *string) *Void

	XmlNanoHTTPOpenRedir func(
		URL string,
		contentType *string,
		redir *string) *Void

	XmlNanoHTTPReturnCode func(
		ctx *Void) int

	XmlNanoHTTPAuthHeader func(
		ctx *Void) string

	XmlNanoHTTPRedir func(
		ctx *Void) string

	XmlNanoHTTPContentLength func(
		ctx *Void) int

	XmlNanoHTTPEncoding func(
		ctx *Void) string

	XmlNanoHTTPMimeType func(
		ctx *Void) string

	XmlNanoHTTPRead func(
		ctx *Void,
		dest *Void,
		len int) int

	XmlNanoHTTPSave func(
		ctxt *Void,
		filename string) int

	XmlNanoHTTPClose func(
		ctx *Void)

	XmlFreePattern func(comp XmlPatternPtr)

	XmlFreePatternList func(
		comp XmlPatternPtr)

	XmlPatterncompile func(
		pattern string,
		dict *XmlDict,
		flags int,
		namespaces *string) XmlPatternPtr

	XmlPatternMatch func(
		comp XmlPatternPtr,
		node XmlNodePtr) int

	XmlPatternStreamable func(
		comp XmlPatternPtr) int

	XmlPatternMaxDepth func(
		comp XmlPatternPtr) int

	XmlPatternMinDepth func(
		comp XmlPatternPtr) int

	XmlPatternFromRoot func(
		comp XmlPatternPtr) int

	XmlPatternGetStreamCtxt func(
		comp XmlPatternPtr) XmlStreamCtxtPtr

	XmlFreeStreamCtxt func(
		stream XmlStreamCtxtPtr)

	XmlStreamPushNode func(
		stream XmlStreamCtxtPtr,
		name string,
		ns string,
		nodeType int) int

	XmlStreamPush func(
		stream XmlStreamCtxtPtr,
		name string,
		ns string) int

	XmlStreamPushAttr func(
		stream XmlStreamCtxtPtr,
		name string,
		ns string) int

	XmlStreamPop func(
		stream XmlStreamCtxtPtr) int

	XmlStreamWantsAnyNode func(
		stream XmlStreamCtxtPtr) int

	XmlSaveToFd func(
		fd int,
		encoding string,
		options int) XmlSaveCtxtPtr

	XmlSaveToFilename func(
		filename string,
		encoding string,
		options int) XmlSaveCtxtPtr

	XmlSaveToBuffer func(
		buffer XmlBufferPtr,
		encoding string,
		options int) XmlSaveCtxtPtr

	XmlSaveToIO func(
		iowrite XmlOutputWriteCallback,
		ioclose XmlOutputCloseCallback,
		ioctx *Void,
		encoding string,
		options int) XmlSaveCtxtPtr

	XmlSaveDoc func(
		ctxt XmlSaveCtxtPtr,
		doc XmlDocPtr) Long

	XmlSaveTree func(
		ctxt XmlSaveCtxtPtr,
		node XmlNodePtr) Long

	XmlSaveFlush func(
		ctxt XmlSaveCtxtPtr) int

	XmlSaveClose func(
		ctxt XmlSaveCtxtPtr) int

	XmlSaveSetEscape func(
		ctxt XmlSaveCtxtPtr,
		escape XmlCharEncodingOutputFunc) int

	XmlSaveSetAttrEscape func(
		ctxt XmlSaveCtxtPtr,
		escape XmlCharEncodingOutputFunc) int

	XmlCreateURI func() XmlURIPtr

	XmlBuildURI func(URI, base string) string

	XmlBuildRelativeURI func(URI, base string) string

	XmlParseURI func(str string) XmlURIPtr

	XmlParseURIRaw func(str string, raw int) XmlURIPtr

	XmlParseURIReference func(uri XmlURIPtr, str string) int

	XmlSaveUri func(uri XmlURIPtr) string

	XmlPrintURI func(stream *FILE, uri XmlURIPtr)

	XmlURIEscapeStr func(str, list string) string

	XmlURIUnescapeString func(
		str string, len int, target string) string

	XmlNormalizeURIPath func(path string) int

	XmlURIEscape func(str string) string

	XmlFreeURI func(uri XmlURIPtr)

	XmlCanonicPath func(path string) string

	XmlPathToURI func(path string) string

	XmlRaiseError func(
		schannel XmlStructuredErrorFunc,
		channel XmlGenericErrorFunc,
		data, ctx, node *Void,
		domain, code int,
		level XmlErrorLevel,
		file string,
		line int,
		str1, str2, str3 string,
		int1, col int,
		msg string,
		v ...VArg)

	XmlSimpleError func(
		domain, code int, node XmlNodePtr, msg, extra string)

	XmlErrEncoding func(
		ctxt XmlParserCtxtPtr,
		xmlerr XmlParserErrors,
		msg, str1, str2 string)

	XmlModuleOpen func(filename string,
		XmlModuleOption int) XmlModulePtr

	XmlModuleSymbol func(module XmlModulePtr, name string,
		result **Void) int

	XmlModuleClose func(
		module XmlModulePtr) int

	XmlModuleFree func(
		module XmlModulePtr) int

	DocbEncodeEntities func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int,
		quoteChar int) int

	DocbSAXParseDoc func(
		cur *XmlChar,
		encoding *Char,
		sax DocbSAXHandlerPtr,
		userData *Void) DocbDocPtr

	DocbParseDoc func(
		cur *XmlChar,
		encoding *Char) DocbDocPtr

	DocbSAXParseFile func(
		filename *Char,
		encoding *Char,
		sax DocbSAXHandlerPtr,
		userData *Void) DocbDocPtr

	DocbParseFile func(
		filename *Char,
		encoding *Char) DocbDocPtr

	DocbFreeParserCtxt func(
		ctxt DocbParserCtxtPtr)

	DocbCreatePushParserCtxt func(
		sax DocbSAXHandlerPtr,
		user_data *Void,
		chunk *Char,
		size int,
		filename *Char,
		enc XmlCharEncoding) DocbParserCtxtPtr

	DocbParseChunk func(
		ctxt DocbParserCtxtPtr,
		chunk *Char,
		size int,
		terminate int) int

	DocbCreateFileParserCtxt func(
		filename *Char,
		encoding *Char) DocbParserCtxtPtr

	DocbParseDocument func(
		ctxt DocbParserCtxtPtr) int

	InitxmlDefaultSAXHandler func(
		hdlr *XmlSAXHandlerV1,
		warning int)

	XmlC14NDocSaveTo func(
		doc XmlDocPtr,
		nodes XmlNodeSetPtr,
		mode int,
		inclusive_ns_prefixes **XmlChar,
		with_comments int,
		buf XmlOutputBufferPtr) int

	XmlC14NDocDumpMemory func(
		doc XmlDocPtr,
		nodes XmlNodeSetPtr,
		mode int,
		inclusive_ns_prefixes **XmlChar,
		with_comments int,
		doc_txt_ptr **XmlChar) int

	XmlC14NDocSave func(
		doc XmlDocPtr,
		nodes XmlNodeSetPtr,
		mode int,
		inclusive_ns_prefixes **XmlChar,
		with_comments int,
		filename *Char,
		compression int) int

	XmlC14NExecute func(
		doc XmlDocPtr,
		is_visible_callback XmlC14NIsVisibleCallback,
		user_data *Void,
		mode int,
		inclusive_ns_prefixes **XmlChar,
		with_comments int,
		buf XmlOutputBufferPtr) int

	XmlErrMemory func(
		ctxt XmlParserCtxtPtr,
		extra *Char)

	XmlParserInputBufferCreateFilenameDefault func(
		XmlParserInputBufferCreateFilenameFunc) XmlParserInputBufferCreateFilenameFunc

	XmlOutputBufferCreateFilenameDefault func(
		XmlOutputBufferCreateFilenameFunc) XmlOutputBufferCreateFilenameFunc

	XmlIsBaseChar func(ch Unsigned_int) int

	XmlIsBlank func(ch Unsigned_int) int

	XmlIsChar func(ch Unsigned_int) int

	XmlIsCombining func(ch Unsigned_int) int

	XmlIsDigit func(ch Unsigned_int) int

	XmlIsExtender func(ch Unsigned_int) int

	XmlIsIdeographic func(ch Unsigned_int) int

	XmlIsPubidChar func(ch Unsigned_int) int
)

var dll = "libxml2-2.dll"

var apiList = Apis{
	{"UTF8ToHtml", &UTF8ToHtml},
	{"UTF8Toisolat1", &UTF8Toisolat1},
	{"__docbDefaultSAXHandler", &DocbDefaultSAXHandler},
	{"__htmlDefaultSAXHandler", &HtmlDefaultSAXHandler},
	{"__oldXMLWDcompatibility", &OldXMLWDcompatibility},
	{"__xmlBufferAllocScheme", &XmlBufferAllocScheme},
	{"__xmlDefaultBufferSize", &XmlDefaultBufferSize},
	{"__xmlDefaultSAXHandler", &XmlDefaultSAXHandler},
	{"__xmlDefaultSAXLocator", &XmlDefaultSAXLocator},
	{"__xmlDeregisterNodeDefaultValue", &XmlDeregisterNodeDefaultValue},
	{"__xmlDoValidityCheckingDefaultValue", &XmlDoValidityCheckingDefaultValue},
	{"__xmlErrEncoding", &XmlErrEncoding},
	{"__xmlGenericError", &XmlGenericError},
	{"__xmlGenericErrorContext", &XmlGenericErrorContext},
	{"__xmlGetWarningsDefaultValue", &XmlGetWarningsDefaultValue},
	{"__xmlIndentTreeOutput", &XmlIndentTreeOutput},
	{"__xmlKeepBlanksDefaultValue", &XmlKeepBlanksDefaultValue},
	{"__xmlLastError", &XmlLastError},
	{"__xmlLineNumbersDefaultValue", &XmlLineNumbersDefaultValue},
	{"__xmlLoadExtDtdDefaultValue", &XmlLoadExtDtdDefaultValue},
	{"__xmlOutputBufferCreateFilenameValue", &XmlOutputBufferCreateFilenameValue},
	{"__xmlParserDebugEntities", &XmlParserDebugEntities},
	{"__xmlParserInputBufferCreateFilenameValue", &XmlParserInputBufferCreateFilenameValue},
	{"__xmlParserVersion", &XmlParserVersion},
	{"__xmlPedanticParserDefaultValue", &XmlPedanticParserDefaultValue},
	{"__xmlRaiseError", &XmlRaiseError},
	{"__xmlRegisterNodeDefaultValue", &XmlRegisterNodeDefaultValue},
	{"__xmlSaveNoEmptyTags", &XmlSaveNoEmptyTags},
	{"__xmlSimpleError", &XmlSimpleError},
	{"__xmlStructuredError", &XmlStructuredError},
	{"__xmlStructuredErrorContext", &XmlStructuredErrorContext},
	{"__xmlSubstituteEntitiesDefaultValue", &XmlSubstituteEntitiesDefaultValue},
	{"__xmlTreeIndentString", &XmlTreeIndentString},
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
	{"xmlACatalogAdd", &XmlACatalogAdd},
	{"xmlACatalogDump", &XmlACatalogDump},
	{"xmlACatalogRemove", &XmlACatalogRemove},
	{"xmlACatalogResolve", &XmlACatalogResolve},
	{"xmlACatalogResolvePublic", &XmlACatalogResolvePublic},
	{"xmlACatalogResolveSystem", &XmlACatalogResolveSystem},
	{"xmlACatalogResolveURI", &XmlACatalogResolveURI},
	{"xmlAddAttributeDecl", &XmlAddAttributeDecl},
	{"xmlAddChild", &XmlAddChild},
	{"xmlAddChildList", &XmlAddChildList},
	{"xmlAddDocEntity", &XmlAddDocEntity},
	{"xmlAddDtdEntity", &XmlAddDtdEntity},
	{"xmlAddElementDecl", &XmlAddElementDecl},
	{"xmlAddEncodingAlias", &XmlAddEncodingAlias},
	{"xmlAddID", &XmlAddID},
	{"xmlAddNextSibling", &XmlAddNextSibling},
	{"xmlAddNotationDecl", &XmlAddNotationDecl},
	{"xmlAddPrevSibling", &XmlAddPrevSibling},
	{"xmlAddRef", &XmlAddRef},
	{"xmlAddSibling", &XmlAddSibling},
	{"xmlAllocOutputBuffer", &XmlAllocOutputBuffer},
	{"xmlAllocParserInputBuffer", &XmlAllocParserInputBuffer},
	{"xmlAttrSerializeTxtContent", &XmlAttrSerializeTxtContent},
	{"xmlAutomataCompile", &XmlAutomataCompile},
	{"xmlAutomataGetInitState", &XmlAutomataGetInitState},
	{"xmlAutomataIsDeterminist", &XmlAutomataIsDeterminist},
	{"xmlAutomataNewAllTrans", &XmlAutomataNewAllTrans},
	{"xmlAutomataNewCountTrans", &XmlAutomataNewCountTrans},
	{"xmlAutomataNewCountTrans2", &XmlAutomataNewCountTrans2},
	{"xmlAutomataNewCountedTrans", &XmlAutomataNewCountedTrans},
	{"xmlAutomataNewCounter", &XmlAutomataNewCounter},
	{"xmlAutomataNewCounterTrans", &XmlAutomataNewCounterTrans},
	{"xmlAutomataNewEpsilon", &XmlAutomataNewEpsilon},
	{"xmlAutomataNewNegTrans", &XmlAutomataNewNegTrans},
	{"xmlAutomataNewOnceTrans", &XmlAutomataNewOnceTrans},
	{"xmlAutomataNewOnceTrans2", &XmlAutomataNewOnceTrans2},
	{"xmlAutomataNewState", &XmlAutomataNewState},
	{"xmlAutomataNewTransition", &XmlAutomataNewTransition},
	{"xmlAutomataNewTransition2", &XmlAutomataNewTransition2},
	{"xmlAutomataSetFinalState", &XmlAutomataSetFinalState},
	{"xmlBoolToText", &XmlBoolToText},
	{"xmlBufferAdd", &XmlBufferAdd},
	{"xmlBufferAddHead", &XmlBufferAddHead},
	{"xmlBufferCCat", &XmlBufferCCat},
	{"xmlBufferCat", &XmlBufferCat},
	{"xmlBufferContent", &XmlBufferContent},
	{"xmlBufferCreate", &XmlBufferCreate},
	{"xmlBufferCreateSize", &XmlBufferCreateSize},
	{"xmlBufferCreateStatic", &XmlBufferCreateStatic},
	{"xmlBufferDump", &XmlBufferDump},
	{"xmlBufferEmpty", &XmlBufferEmpty},
	{"xmlBufferFree", &XmlBufferFree},
	{"xmlBufferGrow", &XmlBufferGrow},
	{"xmlBufferLength", &XmlBufferLength},
	{"xmlBufferResize", &XmlBufferResize},
	{"xmlBufferSetAllocationScheme", &XmlBufferSetAllocationScheme},
	{"xmlBufferShrink", &XmlBufferShrink},
	{"xmlBufferWriteCHAR", &XmlBufferWriteCHAR},
	{"xmlBufferWriteChar", &XmlBufferWriteChar},
	{"xmlBufferWriteQuotedString", &XmlBufferWriteQuotedString},
	{"xmlBuildQName", &XmlBuildQName},
	{"xmlBuildRelativeURI", &XmlBuildRelativeURI},
	{"xmlBuildURI", &XmlBuildURI},
	{"xmlByteConsumed", &XmlByteConsumed},
	{"xmlC14NDocDumpMemory", &XmlC14NDocDumpMemory},
	{"xmlC14NDocSave", &XmlC14NDocSave},
	{"xmlC14NDocSaveTo", &XmlC14NDocSaveTo},
	{"xmlC14NExecute", &XmlC14NExecute},
	{"xmlCanonicPath", &XmlCanonicPath},
	{"xmlCatalogAdd", &XmlCatalogAdd},
	{"xmlCatalogAddLocal", &XmlCatalogAddLocal},
	{"xmlCatalogCleanup", &XmlCatalogCleanup},
	{"xmlCatalogConvert", &XmlCatalogConvert},
	{"xmlCatalogDump", &XmlCatalogDump},
	{"xmlCatalogFreeLocal", &XmlCatalogFreeLocal},
	{"xmlCatalogGetDefaults", &XmlCatalogGetDefaults},
	{"xmlCatalogGetPublic", &XmlCatalogGetPublic},
	{"xmlCatalogGetSystem", &XmlCatalogGetSystem},
	{"xmlCatalogIsEmpty", &XmlCatalogIsEmpty},
	{"xmlCatalogLocalResolve", &XmlCatalogLocalResolve},
	{"xmlCatalogLocalResolveURI", &XmlCatalogLocalResolveURI},
	{"xmlCatalogRemove", &XmlCatalogRemove},
	{"xmlCatalogResolve", &XmlCatalogResolve},
	{"xmlCatalogResolvePublic", &XmlCatalogResolvePublic},
	{"xmlCatalogResolveSystem", &XmlCatalogResolveSystem},
	{"xmlCatalogResolveURI", &XmlCatalogResolveURI},
	{"xmlCatalogSetDebug", &XmlCatalogSetDebug},
	{"xmlCatalogSetDefaultPrefer", &XmlCatalogSetDefaultPrefer},
	{"xmlCatalogSetDefaults", &XmlCatalogSetDefaults},
	{"xmlCharEncCloseFunc", &XmlCharEncCloseFunc},
	{"xmlCharEncFirstLine", &XmlCharEncFirstLine},
	{"xmlCharEncInFunc", &XmlCharEncInFunc},
	{"xmlCharEncOutFunc", &XmlCharEncOutFunc},
	{"xmlCharInRange", &XmlCharInRange},
	{"xmlCharStrdup", &XmlCharStrdup},
	{"xmlCharStrndup", &XmlCharStrndup},
	{"xmlCheckFilename", &XmlCheckFilename},
	{"xmlCheckHTTPInput", &XmlCheckHTTPInput},
	{"xmlCheckLanguageID", &XmlCheckLanguageID},
	{"xmlCheckUTF8", &XmlCheckUTF8},
	{"xmlCheckVersion", &XmlCheckVersion},
	{"xmlChildElementCount", &XmlChildElementCount},
	{"xmlCleanupCharEncodingHandlers", &XmlCleanupCharEncodingHandlers},
	{"xmlCleanupEncodingAliases", &XmlCleanupEncodingAliases},
	{"xmlCleanupGlobals", &XmlCleanupGlobals},
	{"xmlCleanupInputCallbacks", &XmlCleanupInputCallbacks},
	{"xmlCleanupMemory", &XmlCleanupMemory},
	{"xmlCleanupOutputCallbacks", &XmlCleanupOutputCallbacks},
	{"xmlCleanupParser", &XmlCleanupParser},
	{"xmlCleanupPredefinedEntities", &XmlCleanupPredefinedEntities},
	{"xmlCleanupThreads", &XmlCleanupThreads},
	{"xmlClearNodeInfoSeq", &XmlClearNodeInfoSeq},
	{"xmlClearParserCtxt", &XmlClearParserCtxt},
	{"xmlConvertSGMLCatalog", &XmlConvertSGMLCatalog},
	{"xmlCopyAttributeTable", &XmlCopyAttributeTable},
	{"xmlCopyChar", &XmlCopyChar},
	{"xmlCopyCharMultiByte", &XmlCopyCharMultiByte},
	{"xmlCopyDoc", &XmlCopyDoc},
	{"xmlCopyDocElementContent", &XmlCopyDocElementContent},
	{"xmlCopyDtd", &XmlCopyDtd},
	{"xmlCopyElementContent", &XmlCopyElementContent},
	{"xmlCopyElementTable", &XmlCopyElementTable},
	{"xmlCopyEntitiesTable", &XmlCopyEntitiesTable},
	{"xmlCopyEnumeration", &XmlCopyEnumeration},
	{"xmlCopyError", &XmlCopyError},
	{"xmlCopyNamespace", &XmlCopyNamespace},
	{"xmlCopyNamespaceList", &XmlCopyNamespaceList},
	{"xmlCopyNode", &XmlCopyNode},
	{"xmlCopyNodeList", &XmlCopyNodeList},
	{"xmlCopyNotationTable", &XmlCopyNotationTable},
	{"xmlCopyProp", &XmlCopyProp},
	{"xmlCopyPropList", &XmlCopyPropList},
	{"xmlCreateDocParserCtxt", &XmlCreateDocParserCtxt},
	{"xmlCreateEntitiesTable", &XmlCreateEntitiesTable},
	{"xmlCreateEntityParserCtxt", &XmlCreateEntityParserCtxt},
	{"xmlCreateEnumeration", &XmlCreateEnumeration},
	{"xmlCreateFileParserCtxt", &XmlCreateFileParserCtxt},
	{"xmlCreateIOParserCtxt", &XmlCreateIOParserCtxt},
	{"xmlCreateIntSubset", &XmlCreateIntSubset},
	{"xmlCreateMemoryParserCtxt", &XmlCreateMemoryParserCtxt},
	{"xmlCreatePushParserCtxt", &XmlCreatePushParserCtxt},
	{"xmlCreateURI", &XmlCreateURI},
	{"xmlCreateURLParserCtxt", &XmlCreateURLParserCtxt},
	{"xmlCtxtGetLastError", &XmlCtxtGetLastError},
	{"xmlCtxtReadDoc", &XmlCtxtReadDoc},
	{"xmlCtxtReadFd", &XmlCtxtReadFd},
	{"xmlCtxtReadFile", &XmlCtxtReadFile},
	{"xmlCtxtReadIO", &XmlCtxtReadIO},
	{"xmlCtxtReadMemory", &XmlCtxtReadMemory},
	{"xmlCtxtReset", &XmlCtxtReset},
	{"xmlCtxtResetLastError", &XmlCtxtResetLastError},
	{"xmlCtxtResetPush", &XmlCtxtResetPush},
	{"xmlCtxtUseOptions", &XmlCtxtUseOptions},
	{"xmlCurrentChar", &XmlCurrentChar},
	{"xmlDOMWrapAdoptNode", &XmlDOMWrapAdoptNode},
	{"xmlDOMWrapCloneNode", &XmlDOMWrapCloneNode},
	{"xmlDOMWrapFreeCtxt", &XmlDOMWrapFreeCtxt},
	{"xmlDOMWrapNewCtxt", &XmlDOMWrapNewCtxt},
	{"xmlDOMWrapReconcileNamespaces", &XmlDOMWrapReconcileNamespaces},
	{"xmlDOMWrapRemoveNode", &XmlDOMWrapRemoveNode},
	{"xmlDebugCheckDocument", &XmlDebugCheckDocument},
	{"xmlDebugDumpAttr", &XmlDebugDumpAttr},
	{"xmlDebugDumpAttrList", &XmlDebugDumpAttrList},
	{"xmlDebugDumpDTD", &XmlDebugDumpDTD},
	{"xmlDebugDumpDocument", &XmlDebugDumpDocument},
	{"xmlDebugDumpDocumentHead", &XmlDebugDumpDocumentHead},
	{"xmlDebugDumpEntities", &XmlDebugDumpEntities},
	{"xmlDebugDumpNode", &XmlDebugDumpNode},
	{"xmlDebugDumpNodeList", &XmlDebugDumpNodeList},
	{"xmlDebugDumpOneNode", &XmlDebugDumpOneNode},
	{"xmlDebugDumpString", &XmlDebugDumpString},
	{"xmlDecodeEntities", &XmlDecodeEntities},
	{"xmlDefaultSAXHandlerInit", &XmlDefaultSAXHandlerInit},
	{"xmlDelEncodingAlias", &XmlDelEncodingAlias},
	{"xmlDeregisterNodeDefault", &XmlDeregisterNodeDefault},
	{"xmlDetectCharEncoding", &XmlDetectCharEncoding},
	{"xmlDictCleanup", &XmlDictCleanup},
	{"xmlDictCreate", &XmlDictCreate},
	{"xmlDictCreateSub", &XmlDictCreateSub},
	{"xmlDictExists", &XmlDictExists},
	{"xmlDictFree", &XmlDictFree},
	{"xmlDictLookup", &XmlDictLookup},
	{"xmlDictOwns", &XmlDictOwns},
	{"xmlDictQLookup", &XmlDictQLookup},
	{"xmlDictReference", &XmlDictReference},
	{"xmlDictSize", &XmlDictSize},
	{"xmlDocCopyNode", &XmlDocCopyNode},
	{"xmlDocCopyNodeList", &XmlDocCopyNodeList},
	{"xmlDocDump", &XmlDocDump},
	{"xmlDocDumpFormatMemory", &XmlDocDumpFormatMemory},
	{"xmlDocDumpFormatMemoryEnc", &XmlDocDumpFormatMemoryEnc},
	{"xmlDocDumpMemory", &XmlDocDumpMemory},
	{"xmlDocDumpMemoryEnc", &XmlDocDumpMemoryEnc},
	{"xmlDocFormatDump", &XmlDocFormatDump},
	{"xmlDocGetRootElement", &XmlDocGetRootElement},
	{"xmlDocSetRootElement", &XmlDocSetRootElement},
	{"xmlDumpAttributeDecl", &XmlDumpAttributeDecl},
	{"xmlDumpAttributeTable", &XmlDumpAttributeTable},
	{"xmlDumpElementDecl", &XmlDumpElementDecl},
	{"xmlDumpElementTable", &XmlDumpElementTable},
	{"xmlDumpEntitiesTable", &XmlDumpEntitiesTable},
	{"xmlDumpEntityDecl", &XmlDumpEntityDecl},
	{"xmlDumpNotationDecl", &XmlDumpNotationDecl},
	{"xmlDumpNotationTable", &XmlDumpNotationTable},
	{"xmlElemDump", &XmlElemDump},
	{"xmlEncodeEntities", &XmlEncodeEntities},
	{"xmlEncodeEntitiesReentrant", &XmlEncodeEntitiesReentrant},
	{"xmlEncodeSpecialChars", &XmlEncodeSpecialChars},
	{"xmlErrMemory", &XmlErrMemory},
	{"xmlExpCtxtNbCons", &XmlExpCtxtNbCons},
	{"xmlExpCtxtNbNodes", &XmlExpCtxtNbNodes},
	{"xmlExpDump", &XmlExpDump},
	{"xmlExpExpDerive", &XmlExpExpDerive},
	{"xmlExpFree", &XmlExpFree},
	{"xmlExpFreeCtxt", &XmlExpFreeCtxt},
	{"xmlExpGetLanguage", &XmlExpGetLanguage},
	{"xmlExpGetStart", &XmlExpGetStart},
	{"xmlExpIsNillable", &XmlExpIsNillable},
	{"xmlExpMaxToken", &XmlExpMaxToken},
	{"xmlExpNewAtom", &XmlExpNewAtom},
	{"xmlExpNewCtxt", &XmlExpNewCtxt},
	{"xmlExpNewOr", &XmlExpNewOr},
	{"xmlExpNewRange", &XmlExpNewRange},
	{"xmlExpNewSeq", &XmlExpNewSeq},
	{"xmlExpParse", &XmlExpParse},
	{"xmlExpRef", &XmlExpRef},
	{"xmlExpStringDerive", &XmlExpStringDerive},
	{"xmlExpSubsume", &XmlExpSubsume},
	{"xmlFileClose", &XmlFileClose},
	{"xmlFileMatch", &XmlFileMatch},
	{"xmlFileOpen", &XmlFileOpen},
	{"xmlFileRead", &XmlFileRead},
	{"xmlFindCharEncodingHandler", &XmlFindCharEncodingHandler},
	{"xmlFirstElementChild", &XmlFirstElementChild},
	{"xmlFreeAttributeTable", &XmlFreeAttributeTable},
	{"xmlFreeAutomata", &XmlFreeAutomata},
	{"xmlFreeCatalog", &XmlFreeCatalog},
	{"xmlFreeDoc", &XmlFreeDoc},
	{"xmlFreeDocElementContent", &XmlFreeDocElementContent},
	{"xmlFreeDtd", &XmlFreeDtd},
	{"xmlFreeElementContent", &XmlFreeElementContent},
	{"xmlFreeElementTable", &XmlFreeElementTable},
	{"xmlFreeEntitiesTable", &XmlFreeEntitiesTable},
	{"xmlFreeEnumeration", &XmlFreeEnumeration},
	{"xmlFreeIDTable", &XmlFreeIDTable},
	{"xmlFreeInputStream", &XmlFreeInputStream},
	{"xmlFreeMutex", &XmlFreeMutex},
	{"xmlFreeNode", &XmlFreeNode},
	{"xmlFreeNodeList", &XmlFreeNodeList},
	{"xmlFreeNotationTable", &XmlFreeNotationTable},
	{"xmlFreeNs", &XmlFreeNs},
	{"xmlFreeNsList", &XmlFreeNsList},
	{"xmlFreeParserCtxt", &XmlFreeParserCtxt},
	{"xmlFreeParserInputBuffer", &XmlFreeParserInputBuffer},
	{"xmlFreePattern", &XmlFreePattern},
	{"xmlFreePatternList", &XmlFreePatternList},
	{"xmlFreeProp", &XmlFreeProp},
	{"xmlFreePropList", &XmlFreePropList},
	{"xmlFreeRMutex", &XmlFreeRMutex},
	{"xmlFreeRefTable", &XmlFreeRefTable},
	{"xmlFreeStreamCtxt", &XmlFreeStreamCtxt},
	{"xmlFreeTextReader", &XmlFreeTextReader},
	{"xmlFreeTextWriter", &XmlFreeTextWriter},
	{"xmlFreeURI", &XmlFreeURI},
	{"xmlFreeValidCtxt", &XmlFreeValidCtxt},
	{"xmlGcMemGet", &XmlGcMemGet},
	{"xmlGcMemSetup", &XmlGcMemSetup},
	{"xmlGetBufferAllocationScheme", &XmlGetBufferAllocationScheme},
	{"xmlGetCharEncodingHandler", &XmlGetCharEncodingHandler},
	{"xmlGetCharEncodingName", &XmlGetCharEncodingName},
	{"xmlGetCompressMode", &XmlGetCompressMode},
	{"xmlGetDocCompressMode", &XmlGetDocCompressMode},
	{"xmlGetDocEntity", &XmlGetDocEntity},
	{"xmlGetDtdAttrDesc", &XmlGetDtdAttrDesc},
	{"xmlGetDtdElementDesc", &XmlGetDtdElementDesc},
	{"xmlGetDtdEntity", &XmlGetDtdEntity},
	{"xmlGetDtdNotationDesc", &XmlGetDtdNotationDesc},
	{"xmlGetDtdQAttrDesc", &XmlGetDtdQAttrDesc},
	{"xmlGetDtdQElementDesc", &XmlGetDtdQElementDesc},
	{"xmlGetEncodingAlias", &XmlGetEncodingAlias},
	{"xmlGetExternalEntityLoader", &XmlGetExternalEntityLoader},
	{"xmlGetFeature", &XmlGetFeature},
	{"xmlGetFeaturesList", &XmlGetFeaturesList},
	{"xmlGetGlobalState", &XmlGetGlobalState},
	{"xmlGetID", &XmlGetID},
	{"xmlGetIntSubset", &XmlGetIntSubset},
	{"xmlGetLastChild", &XmlGetLastChild},
	{"xmlGetLastError", &XmlGetLastError},
	{"xmlGetLineNo", &XmlGetLineNo},
	{"xmlGetNoNsProp", &XmlGetNoNsProp},
	{"xmlGetNodePath", &XmlGetNodePath},
	{"xmlGetNsList", &XmlGetNsList},
	{"xmlGetNsProp", &XmlGetNsProp},
	{"xmlGetParameterEntity", &XmlGetParameterEntity},
	{"xmlGetPredefinedEntity", &XmlGetPredefinedEntity},
	{"xmlGetProp", &XmlGetProp},
	{"xmlGetRefs", &XmlGetRefs},
	{"xmlGetThreadId", &XmlGetThreadId},
	{"xmlGetUTF8Char", &XmlGetUTF8Char},
	{"xmlHandleEntity", &XmlHandleEntity},
	{"xmlHasFeature", &XmlHasFeature},
	{"xmlHasNsProp", &XmlHasNsProp},
	{"xmlHasProp", &XmlHasProp},
	{"xmlHashAddEntry", &XmlHashAddEntry},
	{"xmlHashAddEntry2", &XmlHashAddEntry2},
	{"xmlHashAddEntry3", &XmlHashAddEntry3},
	{"xmlHashCopy", &XmlHashCopy},
	{"xmlHashCreate", &XmlHashCreate},
	{"xmlHashCreateDict", &XmlHashCreateDict},
	{"xmlHashFree", &XmlHashFree},
	{"xmlHashLookup", &XmlHashLookup},
	{"xmlHashLookup2", &XmlHashLookup2},
	{"xmlHashLookup3", &XmlHashLookup3},
	{"xmlHashQLookup", &XmlHashQLookup},
	{"xmlHashQLookup2", &XmlHashQLookup2},
	{"xmlHashQLookup3", &XmlHashQLookup3},
	{"xmlHashRemoveEntry", &XmlHashRemoveEntry},
	{"xmlHashRemoveEntry2", &XmlHashRemoveEntry2},
	{"xmlHashRemoveEntry3", &XmlHashRemoveEntry3},
	{"xmlHashScan", &XmlHashScan},
	{"xmlHashScan3", &XmlHashScan3},
	{"xmlHashScanFull", &XmlHashScanFull},
	{"xmlHashScanFull3", &XmlHashScanFull3},
	{"xmlHashSize", &XmlHashSize},
	{"xmlHashUpdateEntry", &XmlHashUpdateEntry},
	{"xmlHashUpdateEntry2", &XmlHashUpdateEntry2},
	{"xmlHashUpdateEntry3", &XmlHashUpdateEntry3},
	{"xmlIOFTPClose", &XmlIOFTPClose},
	{"xmlIOFTPMatch", &XmlIOFTPMatch},
	{"xmlIOFTPOpen", &XmlIOFTPOpen},
	{"xmlIOFTPRead", &XmlIOFTPRead},
	{"xmlIOHTTPClose", &XmlIOHTTPClose},
	{"xmlIOHTTPMatch", &XmlIOHTTPMatch},
	{"xmlIOHTTPOpen", &XmlIOHTTPOpen},
	{"xmlIOHTTPOpenW", &XmlIOHTTPOpenW},
	{"xmlIOHTTPRead", &XmlIOHTTPRead},
	{"xmlIOParseDTD", &XmlIOParseDTD},
	{"xmlInitCharEncodingHandlers", &XmlInitCharEncodingHandlers},
	{"xmlInitGlobals", &XmlInitGlobals},
	{"xmlInitMemory", &XmlInitMemory},
	{"xmlInitNodeInfoSeq", &XmlInitNodeInfoSeq},
	{"xmlInitParser", &XmlInitParser},
	{"xmlInitParserCtxt", &XmlInitParserCtxt},
	{"xmlInitThreads", &XmlInitThreads},
	{"xmlInitializeCatalog", &XmlInitializeCatalog},
	{"xmlInitializeGlobalState", &XmlInitializeGlobalState},
	{"xmlInitializePredefinedEntities", &XmlInitializePredefinedEntities},
	{"xmlIsBaseChar", &XmlIsBaseChar},
	{"xmlIsBlank", &XmlIsBlank},
	{"xmlIsBlankNode", &XmlIsBlankNode},
	{"xmlIsChar", &XmlIsChar},
	{"xmlIsCombining", &XmlIsCombining},
	{"xmlIsDigit", &XmlIsDigit},
	{"xmlIsExtender", &XmlIsExtender},
	{"xmlIsID", &XmlIsID},
	{"xmlIsIdeographic", &XmlIsIdeographic},
	{"xmlIsLetter", &XmlIsLetter},
	{"xmlIsMainThread", &XmlIsMainThread},
	{"xmlIsMixedElement", &XmlIsMixedElement},
	{"xmlIsPubidChar", &XmlIsPubidChar},
	{"xmlIsRef", &XmlIsRef},
	{"xmlIsXHTML", &XmlIsXHTML},
	{"xmlKeepBlanksDefault", &XmlKeepBlanksDefault},
	{"xmlLastElementChild", &XmlLastElementChild},
	{"xmlLineNumbersDefault", &XmlLineNumbersDefault},
	{"xmlLinkGetData", &XmlLinkGetData},
	{"xmlListAppend", &XmlListAppend},
	{"xmlListClear", &XmlListClear},
	{"xmlListCopy", &XmlListCopy},
	{"xmlListCreate", &XmlListCreate},
	{"xmlListDelete", &XmlListDelete},
	{"xmlListDup", &XmlListDup},
	{"xmlListEmpty", &XmlListEmpty},
	{"xmlListEnd", &XmlListEnd},
	{"xmlListFront", &XmlListFront},
	{"xmlListInsert", &XmlListInsert},
	{"xmlListMerge", &XmlListMerge},
	{"xmlListPopBack", &XmlListPopBack},
	{"xmlListPopFront", &XmlListPopFront},
	{"xmlListPushBack", &XmlListPushBack},
	{"xmlListPushFront", &XmlListPushFront},
	{"xmlListRemoveAll", &XmlListRemoveAll},
	{"xmlListRemoveFirst", &XmlListRemoveFirst},
	{"xmlListRemoveLast", &XmlListRemoveLast},
	{"xmlListReverse", &XmlListReverse},
	{"xmlListReverseSearch", &XmlListReverseSearch},
	{"xmlListReverseWalk", &XmlListReverseWalk},
	{"xmlListSearch", &XmlListSearch},
	{"xmlListSize", &XmlListSize},
	{"xmlListSort", &XmlListSort},
	{"xmlListWalk", &XmlListWalk},
	{"xmlLoadACatalog", &XmlLoadACatalog},
	{"xmlLoadCatalog", &XmlLoadCatalog},
	{"xmlLoadCatalogs", &XmlLoadCatalogs},
	{"xmlLoadExternalEntity", &XmlLoadExternalEntity},
	{"xmlLoadSGMLSuperCatalog", &XmlLoadSGMLSuperCatalog},
	{"xmlLockLibrary", &XmlLockLibrary},
	{"xmlLsCountNode", &XmlLsCountNode},
	{"xmlLsOneNode", &XmlLsOneNode},
	{"xmlMallocAtomicLoc", &XmlMallocAtomicLoc},
	{"xmlMallocLoc", &XmlMallocLoc},
	{"xmlMemBlocks", &XmlMemBlocks},
	{"xmlMemDisplay", &XmlMemDisplay},
	{"xmlMemDisplayLast", &XmlMemDisplayLast},
	{"xmlMemFree", &XmlMemFree},
	{"xmlMemGet", &XmlMemGet},
	{"xmlMemMalloc", &XmlMemMalloc},
	{"xmlMemRealloc", &XmlMemRealloc},
	{"xmlMemSetup", &XmlMemSetup},
	{"xmlMemShow", &XmlMemShow},
	{"xmlMemStrdupLoc", &XmlMemStrdupLoc},
	{"xmlMemUsed", &XmlMemUsed},
	{"xmlMemoryDump", &XmlMemoryDump},
	{"xmlMemoryStrdup", &XmlMemoryStrdup},
	{"xmlModuleClose", &XmlModuleClose},
	{"xmlModuleFree", &XmlModuleFree},
	{"xmlModuleOpen", &XmlModuleOpen},
	{"xmlModuleSymbol", &XmlModuleSymbol},
	{"xmlMutexLock", &XmlMutexLock},
	{"xmlMutexUnlock", &XmlMutexUnlock},
	{"xmlNamespaceParseNCName", &XmlNamespaceParseNCName},
	{"xmlNamespaceParseNSDef", &XmlNamespaceParseNSDef},
	{"xmlNamespaceParseQName", &XmlNamespaceParseQName},
	{"xmlNanoFTPCheckResponse", &XmlNanoFTPCheckResponse},
	{"xmlNanoFTPCleanup", &XmlNanoFTPCleanup},
	{"xmlNanoFTPClose", &XmlNanoFTPClose},
	{"xmlNanoFTPCloseConnection", &XmlNanoFTPCloseConnection},
	{"xmlNanoFTPConnect", &XmlNanoFTPConnect},
	{"xmlNanoFTPConnectTo", &XmlNanoFTPConnectTo},
	{"xmlNanoFTPCwd", &XmlNanoFTPCwd},
	{"xmlNanoFTPDele", &XmlNanoFTPDele},
	{"xmlNanoFTPFreeCtxt", &XmlNanoFTPFreeCtxt},
	{"xmlNanoFTPGet", &XmlNanoFTPGet},
	{"xmlNanoFTPGetConnection", &XmlNanoFTPGetConnection},
	{"xmlNanoFTPGetResponse", &XmlNanoFTPGetResponse},
	{"xmlNanoFTPGetSocket", &XmlNanoFTPGetSocket},
	{"xmlNanoFTPInit", &XmlNanoFTPInit},
	{"xmlNanoFTPList", &XmlNanoFTPList},
	{"xmlNanoFTPNewCtxt", &XmlNanoFTPNewCtxt},
	{"xmlNanoFTPOpen", &XmlNanoFTPOpen},
	{"xmlNanoFTPProxy", &XmlNanoFTPProxy},
	{"xmlNanoFTPQuit", &XmlNanoFTPQuit},
	{"xmlNanoFTPRead", &XmlNanoFTPRead},
	{"xmlNanoFTPScanProxy", &XmlNanoFTPScanProxy},
	{"xmlNanoFTPUpdateURL", &XmlNanoFTPUpdateURL},
	{"xmlNanoHTTPAuthHeader", &XmlNanoHTTPAuthHeader},
	{"xmlNanoHTTPCleanup", &XmlNanoHTTPCleanup},
	{"xmlNanoHTTPClose", &XmlNanoHTTPClose},
	{"xmlNanoHTTPContentLength", &XmlNanoHTTPContentLength},
	{"xmlNanoHTTPEncoding", &XmlNanoHTTPEncoding},
	{"xmlNanoHTTPFetch", &XmlNanoHTTPFetch},
	{"xmlNanoHTTPInit", &XmlNanoHTTPInit},
	{"xmlNanoHTTPMethod", &XmlNanoHTTPMethod},
	{"xmlNanoHTTPMethodRedir", &XmlNanoHTTPMethodRedir},
	{"xmlNanoHTTPMimeType", &XmlNanoHTTPMimeType},
	{"xmlNanoHTTPOpen", &XmlNanoHTTPOpen},
	{"xmlNanoHTTPOpenRedir", &XmlNanoHTTPOpenRedir},
	{"xmlNanoHTTPRead", &XmlNanoHTTPRead},
	{"xmlNanoHTTPRedir", &XmlNanoHTTPRedir},
	{"xmlNanoHTTPReturnCode", &XmlNanoHTTPReturnCode},
	{"xmlNanoHTTPSave", &XmlNanoHTTPSave},
	{"xmlNanoHTTPScanProxy", &XmlNanoHTTPScanProxy},
	{"xmlNewAutomata", &XmlNewAutomata},
	{"xmlNewCDataBlock", &XmlNewCDataBlock},
	{"xmlNewCatalog", &XmlNewCatalog},
	{"xmlNewCharEncodingHandler", &XmlNewCharEncodingHandler},
	{"xmlNewCharRef", &XmlNewCharRef},
	{"xmlNewChild", &XmlNewChild},
	{"xmlNewComment", &XmlNewComment},
	{"xmlNewDoc", &XmlNewDoc},
	{"xmlNewDocComment", &XmlNewDocComment},
	{"xmlNewDocElementContent", &XmlNewDocElementContent},
	{"xmlNewDocFragment", &XmlNewDocFragment},
	{"xmlNewDocNode", &XmlNewDocNode},
	{"xmlNewDocNodeEatName", &XmlNewDocNodeEatName},
	{"xmlNewDocPI", &XmlNewDocPI},
	{"xmlNewDocProp", &XmlNewDocProp},
	{"xmlNewDocRawNode", &XmlNewDocRawNode},
	{"xmlNewDocText", &XmlNewDocText},
	{"xmlNewDocTextLen", &XmlNewDocTextLen},
	{"xmlNewDtd", &XmlNewDtd},
	{"xmlNewElementContent", &XmlNewElementContent},
	{"xmlNewEntity", &XmlNewEntity},
	{"xmlNewEntityInputStream", &XmlNewEntityInputStream},
	{"xmlNewGlobalNs", &XmlNewGlobalNs},
	{"xmlNewIOInputStream", &XmlNewIOInputStream},
	{"xmlNewInputFromFile", &XmlNewInputFromFile},
	{"xmlNewInputStream", &XmlNewInputStream},
	{"xmlNewMutex", &XmlNewMutex},
	{"xmlNewNode", &XmlNewNode},
	{"xmlNewNodeEatName", &XmlNewNodeEatName},
	{"xmlNewNs", &XmlNewNs},
	{"xmlNewNsProp", &XmlNewNsProp},
	{"xmlNewNsPropEatName", &XmlNewNsPropEatName},
	{"xmlNewPI", &XmlNewPI},
	{"xmlNewParserCtxt", &XmlNewParserCtxt},
	{"xmlNewProp", &XmlNewProp},
	{"xmlNewRMutex", &XmlNewRMutex},
	{"xmlNewReference", &XmlNewReference},
	{"xmlNewStringInputStream", &XmlNewStringInputStream},
	{"xmlNewText", &XmlNewText},
	{"xmlNewTextChild", &XmlNewTextChild},
	{"xmlNewTextLen", &XmlNewTextLen},
	{"xmlNewTextReader", &XmlNewTextReader},
	{"xmlNewTextReaderFilename", &XmlNewTextReaderFilename},
	{"xmlNewTextWriter", &XmlNewTextWriter},
	{"xmlNewTextWriterDoc", &XmlNewTextWriterDoc},
	{"xmlNewTextWriterFilename", &XmlNewTextWriterFilename},
	{"xmlNewTextWriterMemory", &XmlNewTextWriterMemory},
	{"xmlNewTextWriterPushParser", &XmlNewTextWriterPushParser},
	{"xmlNewTextWriterTree", &XmlNewTextWriterTree},
	{"xmlNewValidCtxt", &XmlNewValidCtxt},
	{"xmlNextChar", &XmlNextChar},
	{"xmlNextElementSibling", &XmlNextElementSibling},
	{"xmlNoNetExternalEntityLoader", &XmlNoNetExternalEntityLoader},
	{"xmlNodeAddContent", &XmlNodeAddContent},
	{"xmlNodeAddContentLen", &XmlNodeAddContentLen},
	{"xmlNodeBufGetContent", &XmlNodeBufGetContent},
	{"xmlNodeDump", &XmlNodeDump},
	{"xmlNodeDumpOutput", &XmlNodeDumpOutput},
	{"xmlNodeGetBase", &XmlNodeGetBase},
	{"xmlNodeGetContent", &XmlNodeGetContent},
	{"xmlNodeGetLang", &XmlNodeGetLang},
	{"xmlNodeGetSpacePreserve", &XmlNodeGetSpacePreserve},
	{"xmlNodeIsText", &XmlNodeIsText},
	{"xmlNodeListGetRawString", &XmlNodeListGetRawString},
	{"xmlNodeListGetString", &XmlNodeListGetString},
	{"xmlNodeSetBase", &XmlNodeSetBase},
	{"xmlNodeSetContent", &XmlNodeSetContent},
	{"xmlNodeSetContentLen", &XmlNodeSetContentLen},
	{"xmlNodeSetLang", &XmlNodeSetLang},
	{"xmlNodeSetName", &XmlNodeSetName},
	{"xmlNodeSetSpacePreserve", &XmlNodeSetSpacePreserve},
	{"xmlNormalizeURIPath", &XmlNormalizeURIPath},
	{"xmlNormalizeWindowsPath", &XmlNormalizeWindowsPath},
	{"xmlOutputBufferClose", &XmlOutputBufferClose},
	{"xmlOutputBufferCreateBuffer", &XmlOutputBufferCreateBuffer},
	{"xmlOutputBufferCreateFd", &XmlOutputBufferCreateFd},
	{"xmlOutputBufferCreateFile", &XmlOutputBufferCreateFile},
	{"xmlOutputBufferCreateFilename", &XmlOutputBufferCreateFilename},
	{"xmlOutputBufferCreateFilenameDefault", &XmlOutputBufferCreateFilenameDefault},
	{"xmlOutputBufferCreateIO", &XmlOutputBufferCreateIO},
	{"xmlOutputBufferFlush", &XmlOutputBufferFlush},
	{"xmlOutputBufferWrite", &XmlOutputBufferWrite},
	{"xmlOutputBufferWriteEscape", &XmlOutputBufferWriteEscape},
	{"xmlOutputBufferWriteString", &XmlOutputBufferWriteString},
	{"xmlParseAttValue", &XmlParseAttValue},
	{"xmlParseAttribute", &XmlParseAttribute},
	{"xmlParseAttributeListDecl", &XmlParseAttributeListDecl},
	{"xmlParseAttributeType", &XmlParseAttributeType},
	{"xmlParseBalancedChunkMemory", &XmlParseBalancedChunkMemory},
	{"xmlParseBalancedChunkMemoryRecover", &XmlParseBalancedChunkMemoryRecover},
	{"xmlParseCDSect", &XmlParseCDSect},
	{"xmlParseCatalogFile", &XmlParseCatalogFile},
	{"xmlParseCharData", &XmlParseCharData},
	{"xmlParseCharEncoding", &XmlParseCharEncoding},
	{"xmlParseCharRef", &XmlParseCharRef},
	{"xmlParseChunk", &XmlParseChunk},
	{"xmlParseComment", &XmlParseComment},
	{"xmlParseContent", &XmlParseContent},
	{"xmlParseCtxtExternalEntity", &XmlParseCtxtExternalEntity},
	{"xmlParseDTD", &XmlParseDTD},
	{"xmlParseDefaultDecl", &XmlParseDefaultDecl},
	{"xmlParseDoc", &XmlParseDoc},
	{"xmlParseDocTypeDecl", &XmlParseDocTypeDecl},
	{"xmlParseDocument", &XmlParseDocument},
	{"xmlParseElement", &XmlParseElement},
	{"xmlParseElementChildrenContentDecl", &XmlParseElementChildrenContentDecl},
	{"xmlParseElementContentDecl", &XmlParseElementContentDecl},
	{"xmlParseElementDecl", &XmlParseElementDecl},
	{"xmlParseElementMixedContentDecl", &XmlParseElementMixedContentDecl},
	{"xmlParseEncName", &XmlParseEncName},
	{"xmlParseEncodingDecl", &XmlParseEncodingDecl},
	{"xmlParseEndTag", &XmlParseEndTag},
	{"xmlParseEntity", &XmlParseEntity},
	{"xmlParseEntityDecl", &XmlParseEntityDecl},
	{"xmlParseEntityRef", &XmlParseEntityRef},
	{"xmlParseEntityValue", &XmlParseEntityValue},
	{"xmlParseEnumeratedType", &XmlParseEnumeratedType},
	{"xmlParseEnumerationType", &XmlParseEnumerationType},
	{"xmlParseExtParsedEnt", &XmlParseExtParsedEnt},
	{"xmlParseExternalEntity", &XmlParseExternalEntity},
	{"xmlParseExternalID", &XmlParseExternalID},
	{"xmlParseExternalSubset", &XmlParseExternalSubset},
	{"xmlParseFile", &XmlParseFile},
	{"xmlParseInNodeContext", &XmlParseInNodeContext},
	{"xmlParseMarkupDecl", &XmlParseMarkupDecl},
	{"xmlParseMemory", &XmlParseMemory},
	{"xmlParseMisc", &XmlParseMisc},
	{"xmlParseName", &XmlParseName},
	{"xmlParseNamespace", &XmlParseNamespace},
	{"xmlParseNmtoken", &XmlParseNmtoken},
	{"xmlParseNotationDecl", &XmlParseNotationDecl},
	{"xmlParseNotationType", &XmlParseNotationType},
	{"xmlParsePEReference", &XmlParsePEReference},
	{"xmlParsePI", &XmlParsePI},
	{"xmlParsePITarget", &XmlParsePITarget},
	{"xmlParsePubidLiteral", &XmlParsePubidLiteral},
	{"xmlParseQuotedString", &XmlParseQuotedString},
	{"xmlParseReference", &XmlParseReference},
	{"xmlParseSDDecl", &XmlParseSDDecl},
	{"xmlParseStartTag", &XmlParseStartTag},
	{"xmlParseSystemLiteral", &XmlParseSystemLiteral},
	{"xmlParseTextDecl", &XmlParseTextDecl},
	{"xmlParseURI", &XmlParseURI},
	{"xmlParseURIRaw", &XmlParseURIRaw},
	{"xmlParseURIReference", &XmlParseURIReference},
	{"xmlParseVersionInfo", &XmlParseVersionInfo},
	{"xmlParseVersionNum", &XmlParseVersionNum},
	{"xmlParseXMLDecl", &XmlParseXMLDecl},
	{"xmlParserAddNodeInfo", &XmlParserAddNodeInfo},
	{"xmlParserError", &XmlParserError},
	{"xmlParserFindNodeInfo", &XmlParserFindNodeInfo},
	{"xmlParserFindNodeInfoIndex", &XmlParserFindNodeInfoIndex},
	{"xmlParserGetDirectory", &XmlParserGetDirectory},
	{"xmlParserHandlePEReference", &XmlParserHandlePEReference},
	{"xmlParserHandleReference", &XmlParserHandleReference},
	{"xmlParserInputBufferCreateFd", &XmlParserInputBufferCreateFd},
	{"xmlParserInputBufferCreateFile", &XmlParserInputBufferCreateFile},
	{"xmlParserInputBufferCreateFilename", &XmlParserInputBufferCreateFilename},
	{"xmlParserInputBufferCreateFilenameDefault", &XmlParserInputBufferCreateFilenameDefault},
	{"xmlParserInputBufferCreateIO", &XmlParserInputBufferCreateIO},
	{"xmlParserInputBufferCreateMem", &XmlParserInputBufferCreateMem},
	{"xmlParserInputBufferCreateStatic", &XmlParserInputBufferCreateStatic},
	{"xmlParserInputBufferGrow", &XmlParserInputBufferGrow},
	{"xmlParserInputBufferPush", &XmlParserInputBufferPush},
	{"xmlParserInputBufferRead", &XmlParserInputBufferRead},
	{"xmlParserInputGrow", &XmlParserInputGrow},
	{"xmlParserInputRead", &XmlParserInputRead},
	{"xmlParserInputShrink", &XmlParserInputShrink},
	{"xmlParserPrintFileContext", &XmlParserPrintFileContext},
	{"xmlParserPrintFileInfo", &XmlParserPrintFileInfo},
	{"xmlParserValidityError", &XmlParserValidityError},
	{"xmlParserValidityWarning", &XmlParserValidityWarning},
	{"xmlParserWarning", &XmlParserWarning},
	{"xmlPathToURI", &XmlPathToURI},
	{"xmlPatternFromRoot", &XmlPatternFromRoot},
	{"xmlPatternGetStreamCtxt", &XmlPatternGetStreamCtxt},
	{"xmlPatternMatch", &XmlPatternMatch},
	{"xmlPatternMaxDepth", &XmlPatternMaxDepth},
	{"xmlPatternMinDepth", &XmlPatternMinDepth},
	{"xmlPatternStreamable", &XmlPatternStreamable},
	{"xmlPatterncompile", &XmlPatterncompile},
	{"xmlPedanticParserDefault", &XmlPedanticParserDefault},
	{"xmlPopInput", &XmlPopInput},
	{"xmlPopInputCallbacks", &XmlPopInputCallbacks},
	{"xmlPreviousElementSibling", &XmlPreviousElementSibling},
	{"xmlPrintURI", &XmlPrintURI},
	{"xmlPushInput", &XmlPushInput},
	{"xmlRMutexLock", &XmlRMutexLock},
	{"xmlRMutexUnlock", &XmlRMutexUnlock},
	{"xmlReadDoc", &XmlReadDoc},
	{"xmlReadFd", &XmlReadFd},
	{"xmlReadFile", &XmlReadFile},
	{"xmlReadIO", &XmlReadIO},
	{"xmlReadMemory", &XmlReadMemory},
	{"xmlReaderForDoc", &XmlReaderForDoc},
	{"xmlReaderForFd", &XmlReaderForFd},
	{"xmlReaderForFile", &XmlReaderForFile},
	{"xmlReaderForIO", &XmlReaderForIO},
	{"xmlReaderForMemory", &XmlReaderForMemory},
	{"xmlReaderNewDoc", &XmlReaderNewDoc},
	{"xmlReaderNewFd", &XmlReaderNewFd},
	{"xmlReaderNewFile", &XmlReaderNewFile},
	{"xmlReaderNewIO", &XmlReaderNewIO},
	{"xmlReaderNewMemory", &XmlReaderNewMemory},
	{"xmlReaderNewWalker", &XmlReaderNewWalker},
	{"xmlReaderWalker", &XmlReaderWalker},
	{"xmlReallocLoc", &XmlReallocLoc},
	{"xmlReconciliateNs", &XmlReconciliateNs},
	{"xmlRecoverDoc", &XmlRecoverDoc},
	{"xmlRecoverFile", &XmlRecoverFile},
	{"xmlRecoverMemory", &XmlRecoverMemory},
	{"xmlRegExecErrInfo", &XmlRegExecErrInfo},
	{"xmlRegExecNextValues", &XmlRegExecNextValues},
	{"xmlRegExecPushString", &XmlRegExecPushString},
	{"xmlRegExecPushString2", &XmlRegExecPushString2},
	{"xmlRegFreeExecCtxt", &XmlRegFreeExecCtxt},
	{"xmlRegFreeRegexp", &XmlRegFreeRegexp},
	{"xmlRegNewExecCtxt", &XmlRegNewExecCtxt},
	{"xmlRegexpCompile", &XmlRegexpCompile},
	{"xmlRegexpExec", &XmlRegexpExec},
	{"xmlRegexpIsDeterminist", &XmlRegexpIsDeterminist},
	{"xmlRegexpPrint", &XmlRegexpPrint},
	{"xmlRegisterCharEncodingHandler", &XmlRegisterCharEncodingHandler},
	{"xmlRegisterDefaultInputCallbacks", &XmlRegisterDefaultInputCallbacks},
	{"xmlRegisterDefaultOutputCallbacks", &XmlRegisterDefaultOutputCallbacks},
	{"xmlRegisterHTTPPostCallbacks", &XmlRegisterHTTPPostCallbacks},
	{"xmlRegisterInputCallbacks", &XmlRegisterInputCallbacks},
	{"xmlRegisterNodeDefault", &XmlRegisterNodeDefault},
	{"xmlRegisterOutputCallbacks", &XmlRegisterOutputCallbacks},
	{"xmlRelaxNGCleanupTypes", &XmlRelaxNGCleanupTypes},
	{"xmlRelaxNGDump", &XmlRelaxNGDump},
	{"xmlRelaxNGDumpTree", &XmlRelaxNGDumpTree},
	{"xmlRelaxNGFree", &XmlRelaxNGFree},
	{"xmlRelaxNGFreeParserCtxt", &XmlRelaxNGFreeParserCtxt},
	{"xmlRelaxNGFreeValidCtxt", &XmlRelaxNGFreeValidCtxt},
	{"xmlRelaxNGGetParserErrors", &XmlRelaxNGGetParserErrors},
	{"xmlRelaxNGGetValidErrors", &XmlRelaxNGGetValidErrors},
	{"xmlRelaxNGInitTypes", &XmlRelaxNGInitTypes},
	{"xmlRelaxNGNewDocParserCtxt", &XmlRelaxNGNewDocParserCtxt},
	{"xmlRelaxNGNewMemParserCtxt", &XmlRelaxNGNewMemParserCtxt},
	{"xmlRelaxNGNewParserCtxt", &XmlRelaxNGNewParserCtxt},
	{"xmlRelaxNGNewValidCtxt", &XmlRelaxNGNewValidCtxt},
	{"xmlRelaxNGParse", &XmlRelaxNGParse},
	{"xmlRelaxNGSetParserErrors", &XmlRelaxNGSetParserErrors},
	{"xmlRelaxNGSetParserStructuredErrors", &XmlRelaxNGSetParserStructuredErrors},
	{"xmlRelaxNGSetValidErrors", &XmlRelaxNGSetValidErrors},
	{"xmlRelaxNGSetValidStructuredErrors", &XmlRelaxNGSetValidStructuredErrors},
	{"xmlRelaxNGValidateDoc", &XmlRelaxNGValidateDoc},
	{"xmlRelaxNGValidateFullElement", &XmlRelaxNGValidateFullElement},
	{"xmlRelaxNGValidatePopElement", &XmlRelaxNGValidatePopElement},
	{"xmlRelaxNGValidatePushCData", &XmlRelaxNGValidatePushCData},
	{"xmlRelaxNGValidatePushElement", &XmlRelaxNGValidatePushElement},
	{"xmlRelaxParserSetFlag", &XmlRelaxParserSetFlag},
	{"xmlRemoveID", &XmlRemoveID},
	{"xmlRemoveProp", &XmlRemoveProp},
	{"xmlRemoveRef", &XmlRemoveRef},
	{"xmlReplaceNode", &XmlReplaceNode},
	{"xmlResetError", &XmlResetError},
	{"xmlResetLastError", &XmlResetLastError},
	{"xmlSAX2AttributeDecl", &XmlSAX2AttributeDecl},
	{"xmlSAX2CDataBlock", &XmlSAX2CDataBlock},
	{"xmlSAX2Characters", &XmlSAX2Characters},
	{"xmlSAX2Comment", &XmlSAX2Comment},
	{"xmlSAX2ElementDecl", &XmlSAX2ElementDecl},
	{"xmlSAX2EndDocument", &XmlSAX2EndDocument},
	{"xmlSAX2EndElement", &XmlSAX2EndElement},
	{"xmlSAX2EndElementNs", &XmlSAX2EndElementNs},
	{"xmlSAX2EntityDecl", &XmlSAX2EntityDecl},
	{"xmlSAX2ExternalSubset", &XmlSAX2ExternalSubset},
	{"xmlSAX2GetColumnNumber", &XmlSAX2GetColumnNumber},
	{"xmlSAX2GetEntity", &XmlSAX2GetEntity},
	{"xmlSAX2GetLineNumber", &XmlSAX2GetLineNumber},
	{"xmlSAX2GetParameterEntity", &XmlSAX2GetParameterEntity},
	{"xmlSAX2GetPublicId", &XmlSAX2GetPublicId},
	{"xmlSAX2GetSystemId", &XmlSAX2GetSystemId},
	{"xmlSAX2HasExternalSubset", &XmlSAX2HasExternalSubset},
	{"xmlSAX2HasInternalSubset", &XmlSAX2HasInternalSubset},
	{"xmlSAX2IgnorableWhitespace", &XmlSAX2IgnorableWhitespace},
	{"xmlSAX2InitDefaultSAXHandler", &XmlSAX2InitDefaultSAXHandler},
	{"xmlSAX2InitDocbDefaultSAXHandler", &XmlSAX2InitDocbDefaultSAXHandler},
	{"xmlSAX2InitHtmlDefaultSAXHandler", &XmlSAX2InitHtmlDefaultSAXHandler},
	{"xmlSAX2InternalSubset", &XmlSAX2InternalSubset},
	{"xmlSAX2IsStandalone", &XmlSAX2IsStandalone},
	{"xmlSAX2NotationDecl", &XmlSAX2NotationDecl},
	{"xmlSAX2ProcessingInstruction", &XmlSAX2ProcessingInstruction},
	{"xmlSAX2Reference", &XmlSAX2Reference},
	{"xmlSAX2ResolveEntity", &XmlSAX2ResolveEntity},
	{"xmlSAX2SetDocumentLocator", &XmlSAX2SetDocumentLocator},
	{"xmlSAX2StartDocument", &XmlSAX2StartDocument},
	{"xmlSAX2StartElement", &XmlSAX2StartElement},
	{"xmlSAX2StartElementNs", &XmlSAX2StartElementNs},
	{"xmlSAX2UnparsedEntityDecl", &XmlSAX2UnparsedEntityDecl},
	{"xmlSAXDefaultVersion", &XmlSAXDefaultVersion},
	{"xmlSAXParseDTD", &XmlSAXParseDTD},
	{"xmlSAXParseDoc", &XmlSAXParseDoc},
	{"xmlSAXParseEntity", &XmlSAXParseEntity},
	{"xmlSAXParseFile", &XmlSAXParseFile},
	{"xmlSAXParseFileWithData", &XmlSAXParseFileWithData},
	{"xmlSAXParseMemory", &XmlSAXParseMemory},
	{"xmlSAXParseMemoryWithData", &XmlSAXParseMemoryWithData},
	{"xmlSAXUserParseFile", &XmlSAXUserParseFile},
	{"xmlSAXUserParseMemory", &XmlSAXUserParseMemory},
	{"xmlSAXVersion", &XmlSAXVersion},
	{"xmlSaveClose", &XmlSaveClose},
	{"xmlSaveDoc", &XmlSaveDoc},
	{"xmlSaveFile", &XmlSaveFile},
	{"xmlSaveFileEnc", &XmlSaveFileEnc},
	{"xmlSaveFileTo", &XmlSaveFileTo},
	{"xmlSaveFlush", &XmlSaveFlush},
	{"xmlSaveFormatFile", &XmlSaveFormatFile},
	{"xmlSaveFormatFileEnc", &XmlSaveFormatFileEnc},
	{"xmlSaveFormatFileTo", &XmlSaveFormatFileTo},
	{"xmlSaveSetAttrEscape", &XmlSaveSetAttrEscape},
	{"xmlSaveSetEscape", &XmlSaveSetEscape},
	{"xmlSaveToBuffer", &XmlSaveToBuffer},
	{"xmlSaveToFd", &XmlSaveToFd},
	{"xmlSaveToFilename", &XmlSaveToFilename},
	{"xmlSaveToIO", &XmlSaveToIO},
	{"xmlSaveTree", &XmlSaveTree},
	{"xmlSaveUri", &XmlSaveUri},
	{"xmlScanName", &XmlScanName},
	{"xmlSchemaCheckFacet", &XmlSchemaCheckFacet},
	{"xmlSchemaCleanupTypes", &XmlSchemaCleanupTypes},
	{"xmlSchemaCollapseString", &XmlSchemaCollapseString},
	{"xmlSchemaCompareValues", &XmlSchemaCompareValues},
	{"xmlSchemaCompareValuesWhtsp", &XmlSchemaCompareValuesWhtsp},
	{"xmlSchemaCopyValue", &XmlSchemaCopyValue},
	{"xmlSchemaDump", &XmlSchemaDump},
	{"xmlSchemaFree", &XmlSchemaFree},
	{"xmlSchemaFreeFacet", &XmlSchemaFreeFacet},
	{"xmlSchemaFreeParserCtxt", &XmlSchemaFreeParserCtxt},
	{"xmlSchemaFreeType", &XmlSchemaFreeType},
	{"xmlSchemaFreeValidCtxt", &XmlSchemaFreeValidCtxt},
	{"xmlSchemaFreeValue", &XmlSchemaFreeValue},
	{"xmlSchemaFreeWildcard", &XmlSchemaFreeWildcard},
	{"xmlSchemaGetBuiltInListSimpleTypeItemType", &XmlSchemaGetBuiltInListSimpleTypeItemType},
	{"xmlSchemaGetBuiltInType", &XmlSchemaGetBuiltInType},
	{"xmlSchemaGetCanonValue", &XmlSchemaGetCanonValue},
	{"xmlSchemaGetCanonValueWhtsp", &XmlSchemaGetCanonValueWhtsp},
	{"xmlSchemaGetFacetValueAsULong", &XmlSchemaGetFacetValueAsULong},
	{"xmlSchemaGetParserErrors", &XmlSchemaGetParserErrors},
	{"xmlSchemaGetPredefinedType", &XmlSchemaGetPredefinedType},
	{"xmlSchemaGetValType", &XmlSchemaGetValType},
	{"xmlSchemaGetValidErrors", &XmlSchemaGetValidErrors},
	{"xmlSchemaInitTypes", &XmlSchemaInitTypes},
	{"xmlSchemaIsBuiltInTypeFacet", &XmlSchemaIsBuiltInTypeFacet},
	{"xmlSchemaIsValid", &XmlSchemaIsValid},
	{"xmlSchemaNewDocParserCtxt", &XmlSchemaNewDocParserCtxt},
	{"xmlSchemaNewFacet", &XmlSchemaNewFacet},
	{"xmlSchemaNewMemParserCtxt", &XmlSchemaNewMemParserCtxt},
	{"xmlSchemaNewNOTATIONValue", &XmlSchemaNewNOTATIONValue},
	{"xmlSchemaNewParserCtxt", &XmlSchemaNewParserCtxt},
	{"xmlSchemaNewQNameValue", &XmlSchemaNewQNameValue},
	{"xmlSchemaNewStringValue", &XmlSchemaNewStringValue},
	{"xmlSchemaNewValidCtxt", &XmlSchemaNewValidCtxt},
	{"xmlSchemaParse", &XmlSchemaParse},
	{"xmlSchemaSAXPlug", &XmlSchemaSAXPlug},
	{"xmlSchemaSAXUnplug", &XmlSchemaSAXUnplug},
	{"xmlSchemaSetParserErrors", &XmlSchemaSetParserErrors},
	{"xmlSchemaSetParserStructuredErrors", &XmlSchemaSetParserStructuredErrors},
	{"xmlSchemaSetValidErrors", &XmlSchemaSetValidErrors},
	{"xmlSchemaSetValidOptions", &XmlSchemaSetValidOptions},
	{"xmlSchemaSetValidStructuredErrors", &XmlSchemaSetValidStructuredErrors},
	{"xmlSchemaValPredefTypeNode", &XmlSchemaValPredefTypeNode},
	{"xmlSchemaValPredefTypeNodeNoNorm", &XmlSchemaValPredefTypeNodeNoNorm},
	{"xmlSchemaValidCtxtGetOptions", &XmlSchemaValidCtxtGetOptions},
	{"xmlSchemaValidCtxtGetParserCtxt", &XmlSchemaValidCtxtGetParserCtxt},
	{"xmlSchemaValidateDoc", &XmlSchemaValidateDoc},
	{"xmlSchemaValidateFacet", &XmlSchemaValidateFacet},
	{"xmlSchemaValidateFacetWhtsp", &XmlSchemaValidateFacetWhtsp},
	{"xmlSchemaValidateFile", &XmlSchemaValidateFile},
	{"xmlSchemaValidateLengthFacet", &XmlSchemaValidateLengthFacet},
	{"xmlSchemaValidateLengthFacetWhtsp", &XmlSchemaValidateLengthFacetWhtsp},
	{"xmlSchemaValidateListSimpleTypeFacet", &XmlSchemaValidateListSimpleTypeFacet},
	{"xmlSchemaValidateOneElement", &XmlSchemaValidateOneElement},
	{"xmlSchemaValidatePredefinedType", &XmlSchemaValidatePredefinedType},
	{"xmlSchemaValidateStream", &XmlSchemaValidateStream},
	{"xmlSchemaValueAppend", &XmlSchemaValueAppend},
	{"xmlSchemaValueGetAsBoolean", &XmlSchemaValueGetAsBoolean},
	{"xmlSchemaValueGetAsString", &XmlSchemaValueGetAsString},
	{"xmlSchemaValueGetNext", &XmlSchemaValueGetNext},
	{"xmlSchemaWhiteSpaceReplace", &XmlSchemaWhiteSpaceReplace},
	{"xmlSchematronFree", &XmlSchematronFree},
	{"xmlSchematronFreeParserCtxt", &XmlSchematronFreeParserCtxt},
	{"xmlSchematronFreeValidCtxt", &XmlSchematronFreeValidCtxt},
	{"xmlSchematronNewDocParserCtxt", &XmlSchematronNewDocParserCtxt},
	{"xmlSchematronNewMemParserCtxt", &XmlSchematronNewMemParserCtxt},
	{"xmlSchematronNewParserCtxt", &XmlSchematronNewParserCtxt},
	{"xmlSchematronNewValidCtxt", &XmlSchematronNewValidCtxt},
	{"xmlSchematronParse", &XmlSchematronParse},
	{"xmlSchematronSetValidStructuredErrors", &XmlSchematronSetValidStructuredErrors},
	{"xmlSchematronValidateDoc", &XmlSchematronValidateDoc},
	{"xmlSearchNs", &XmlSearchNs},
	{"xmlSearchNsByHref", &XmlSearchNsByHref},
	{"xmlSetBufferAllocationScheme", &XmlSetBufferAllocationScheme},
	{"xmlSetCompressMode", &XmlSetCompressMode},
	{"xmlSetDocCompressMode", &XmlSetDocCompressMode},
	{"xmlSetEntityReferenceFunc", &XmlSetEntityReferenceFunc},
	{"xmlSetExternalEntityLoader", &XmlSetExternalEntityLoader},
	{"xmlSetFeature", &XmlSetFeature},
	{"xmlSetGenericErrorFunc", &XmlSetGenericErrorFunc},
	{"xmlSetListDoc", &XmlSetListDoc},
	{"xmlSetNs", &XmlSetNs},
	{"xmlSetNsProp", &XmlSetNsProp},
	{"xmlSetProp", &XmlSetProp},
	{"xmlSetStructuredErrorFunc", &XmlSetStructuredErrorFunc},
	{"xmlSetTreeDoc", &XmlSetTreeDoc},
	{"xmlSetupParserForBuffer", &XmlSetupParserForBuffer},
	{"xmlShell", &XmlShell},
	{"xmlShellBase", &XmlShellBase},
	{"xmlShellCat", &XmlShellCat},
	{"xmlShellDir", &XmlShellDir},
	{"xmlShellDu", &XmlShellDu},
	{"xmlShellList", &XmlShellList},
	{"xmlShellLoad", &XmlShellLoad},
	{"xmlShellPrintNode", &XmlShellPrintNode},
	{"xmlShellPrintXPathError", &XmlShellPrintXPathError},
	{"xmlShellPrintXPathResult", &XmlShellPrintXPathResult},
	{"xmlShellPwd", &XmlShellPwd},
	{"xmlShellSave", &XmlShellSave},
	{"xmlShellValidate", &XmlShellValidate},
	{"xmlShellWrite", &XmlShellWrite},
	{"xmlSkipBlankChars", &XmlSkipBlankChars},
	{"xmlSnprintfElementContent", &XmlSnprintfElementContent},
	{"xmlSplitQName", &XmlSplitQName},
	{"xmlSplitQName2", &XmlSplitQName2},
	{"xmlSplitQName3", &XmlSplitQName3},
	{"xmlSprintfElementContent", &XmlSprintfElementContent},
	{"xmlStopParser", &XmlStopParser},
	{"xmlStrEqual", &XmlStrEqual},
	{"xmlStrPrintf", &XmlStrPrintf},
	{"xmlStrQEqual", &XmlStrQEqual},
	{"xmlStrVPrintf", &XmlStrVPrintf},
	{"xmlStrcasecmp", &XmlStrcasecmp},
	{"xmlStrcasestr", &XmlStrcasestr},
	{"xmlStrcat", &XmlStrcat},
	{"xmlStrchr", &XmlStrchr},
	{"xmlStrcmp", &XmlStrcmp},
	{"xmlStrdup", &XmlStrdup},
	{"xmlStreamPop", &XmlStreamPop},
	{"xmlStreamPush", &XmlStreamPush},
	{"xmlStreamPushAttr", &XmlStreamPushAttr},
	{"xmlStreamPushNode", &XmlStreamPushNode},
	{"xmlStreamWantsAnyNode", &XmlStreamWantsAnyNode},
	{"xmlStringCurrentChar", &XmlStringCurrentChar},
	{"xmlStringDecodeEntities", &XmlStringDecodeEntities},
	{"xmlStringGetNodeList", &XmlStringGetNodeList},
	{"xmlStringLenDecodeEntities", &XmlStringLenDecodeEntities},
	{"xmlStringLenGetNodeList", &XmlStringLenGetNodeList},
	{"xmlStrlen", &XmlStrlen},
	{"xmlStrncasecmp", &XmlStrncasecmp},
	{"xmlStrncat", &XmlStrncat},
	{"xmlStrncatNew", &XmlStrncatNew},
	{"xmlStrncmp", &XmlStrncmp},
	{"xmlStrndup", &XmlStrndup},
	{"xmlStrstr", &XmlStrstr},
	{"xmlStrsub", &XmlStrsub},
	{"xmlSubstituteEntitiesDefault", &XmlSubstituteEntitiesDefault},
	{"xmlSwitchEncoding", &XmlSwitchEncoding},
	{"xmlSwitchInputEncoding", &XmlSwitchInputEncoding},
	{"xmlSwitchToEncoding", &XmlSwitchToEncoding},
	{"xmlTextConcat", &XmlTextConcat},
	{"xmlTextMerge", &XmlTextMerge},
	{"xmlTextReaderAttributeCount", &XmlTextReaderAttributeCount},
	{"xmlTextReaderBaseUri", &XmlTextReaderBaseUri},
	{"xmlTextReaderByteConsumed", &XmlTextReaderByteConsumed},
	{"xmlTextReaderClose", &XmlTextReaderClose},
	{"xmlTextReaderConstBaseUri", &XmlTextReaderConstBaseUri},
	{"xmlTextReaderConstEncoding", &XmlTextReaderConstEncoding},
	{"xmlTextReaderConstLocalName", &XmlTextReaderConstLocalName},
	{"xmlTextReaderConstName", &XmlTextReaderConstName},
	{"xmlTextReaderConstNamespaceUri", &XmlTextReaderConstNamespaceUri},
	{"xmlTextReaderConstPrefix", &XmlTextReaderConstPrefix},
	{"xmlTextReaderConstString", &XmlTextReaderConstString},
	{"xmlTextReaderConstValue", &XmlTextReaderConstValue},
	{"xmlTextReaderConstXmlLang", &XmlTextReaderConstXmlLang},
	{"xmlTextReaderConstXmlVersion", &XmlTextReaderConstXmlVersion},
	{"xmlTextReaderCurrentDoc", &XmlTextReaderCurrentDoc},
	{"xmlTextReaderCurrentNode", &XmlTextReaderCurrentNode},
	{"xmlTextReaderDepth", &XmlTextReaderDepth},
	{"xmlTextReaderExpand", &XmlTextReaderExpand},
	{"xmlTextReaderGetAttribute", &XmlTextReaderGetAttribute},
	{"xmlTextReaderGetAttributeNo", &XmlTextReaderGetAttributeNo},
	{"xmlTextReaderGetAttributeNs", &XmlTextReaderGetAttributeNs},
	{"xmlTextReaderGetErrorHandler", &XmlTextReaderGetErrorHandler},
	{"xmlTextReaderGetParserColumnNumber", &XmlTextReaderGetParserColumnNumber},
	{"xmlTextReaderGetParserLineNumber", &XmlTextReaderGetParserLineNumber},
	{"xmlTextReaderGetParserProp", &XmlTextReaderGetParserProp},
	{"xmlTextReaderGetRemainder", &XmlTextReaderGetRemainder},
	{"xmlTextReaderHasAttributes", &XmlTextReaderHasAttributes},
	{"xmlTextReaderHasValue", &XmlTextReaderHasValue},
	{"xmlTextReaderIsDefault", &XmlTextReaderIsDefault},
	{"xmlTextReaderIsEmptyElement", &XmlTextReaderIsEmptyElement},
	{"xmlTextReaderIsNamespaceDecl", &XmlTextReaderIsNamespaceDecl},
	{"xmlTextReaderIsValid", &XmlTextReaderIsValid},
	{"xmlTextReaderLocalName", &XmlTextReaderLocalName},
	{"xmlTextReaderLocatorBaseURI", &XmlTextReaderLocatorBaseURI},
	{"xmlTextReaderLocatorLineNumber", &XmlTextReaderLocatorLineNumber},
	{"xmlTextReaderLookupNamespace", &XmlTextReaderLookupNamespace},
	{"xmlTextReaderMoveToAttribute", &XmlTextReaderMoveToAttribute},
	{"xmlTextReaderMoveToAttributeNo", &XmlTextReaderMoveToAttributeNo},
	{"xmlTextReaderMoveToAttributeNs", &XmlTextReaderMoveToAttributeNs},
	{"xmlTextReaderMoveToElement", &XmlTextReaderMoveToElement},
	{"xmlTextReaderMoveToFirstAttribute", &XmlTextReaderMoveToFirstAttribute},
	{"xmlTextReaderMoveToNextAttribute", &XmlTextReaderMoveToNextAttribute},
	{"xmlTextReaderName", &XmlTextReaderName},
	{"xmlTextReaderNamespaceUri", &XmlTextReaderNamespaceUri},
	{"xmlTextReaderNext", &XmlTextReaderNext},
	{"xmlTextReaderNextSibling", &XmlTextReaderNextSibling},
	{"xmlTextReaderNodeType", &XmlTextReaderNodeType},
	{"xmlTextReaderNormalization", &XmlTextReaderNormalization},
	{"xmlTextReaderPrefix", &XmlTextReaderPrefix},
	{"xmlTextReaderPreserve", &XmlTextReaderPreserve},
	{"xmlTextReaderPreservePattern", &XmlTextReaderPreservePattern},
	{"xmlTextReaderQuoteChar", &XmlTextReaderQuoteChar},
	{"xmlTextReaderRead", &XmlTextReaderRead},
	{"xmlTextReaderReadAttributeValue", &XmlTextReaderReadAttributeValue},
	{"xmlTextReaderReadInnerXml", &XmlTextReaderReadInnerXml},
	{"xmlTextReaderReadOuterXml", &XmlTextReaderReadOuterXml},
	{"xmlTextReaderReadState", &XmlTextReaderReadState},
	{"xmlTextReaderReadString", &XmlTextReaderReadString},
	{"xmlTextReaderRelaxNGSetSchema", &XmlTextReaderRelaxNGSetSchema},
	{"xmlTextReaderRelaxNGValidate", &XmlTextReaderRelaxNGValidate},
	{"xmlTextReaderSchemaValidate", &XmlTextReaderSchemaValidate},
	{"xmlTextReaderSchemaValidateCtxt", &XmlTextReaderSchemaValidateCtxt},
	{"xmlTextReaderSetErrorHandler", &XmlTextReaderSetErrorHandler},
	{"xmlTextReaderSetParserProp", &XmlTextReaderSetParserProp},
	{"xmlTextReaderSetSchema", &XmlTextReaderSetSchema},
	{"xmlTextReaderSetStructuredErrorHandler", &XmlTextReaderSetStructuredErrorHandler},
	{"xmlTextReaderSetup", &XmlTextReaderSetup},
	{"xmlTextReaderStandalone", &XmlTextReaderStandalone},
	{"xmlTextReaderValue", &XmlTextReaderValue},
	{"xmlTextReaderXmlLang", &XmlTextReaderXmlLang},
	{"xmlTextWriterEndAttribute", &XmlTextWriterEndAttribute},
	{"xmlTextWriterEndCDATA", &XmlTextWriterEndCDATA},
	{"xmlTextWriterEndComment", &XmlTextWriterEndComment},
	{"xmlTextWriterEndDTD", &XmlTextWriterEndDTD},
	{"xmlTextWriterEndDTDAttlist", &XmlTextWriterEndDTDAttlist},
	{"xmlTextWriterEndDTDElement", &XmlTextWriterEndDTDElement},
	{"xmlTextWriterEndDTDEntity", &XmlTextWriterEndDTDEntity},
	{"xmlTextWriterEndDocument", &XmlTextWriterEndDocument},
	{"xmlTextWriterEndElement", &XmlTextWriterEndElement},
	{"xmlTextWriterEndPI", &XmlTextWriterEndPI},
	{"xmlTextWriterFlush", &XmlTextWriterFlush},
	{"xmlTextWriterFullEndElement", &XmlTextWriterFullEndElement},
	{"xmlTextWriterSetIndent", &XmlTextWriterSetIndent},
	{"xmlTextWriterSetIndentString", &XmlTextWriterSetIndentString},
	{"xmlTextWriterStartAttribute", &XmlTextWriterStartAttribute},
	{"xmlTextWriterStartAttributeNS", &XmlTextWriterStartAttributeNS},
	{"xmlTextWriterStartCDATA", &XmlTextWriterStartCDATA},
	{"xmlTextWriterStartComment", &XmlTextWriterStartComment},
	{"xmlTextWriterStartDTD", &XmlTextWriterStartDTD},
	{"xmlTextWriterStartDTDAttlist", &XmlTextWriterStartDTDAttlist},
	{"xmlTextWriterStartDTDElement", &XmlTextWriterStartDTDElement},
	{"xmlTextWriterStartDTDEntity", &XmlTextWriterStartDTDEntity},
	{"xmlTextWriterStartDocument", &XmlTextWriterStartDocument},
	{"xmlTextWriterStartElement", &XmlTextWriterStartElement},
	{"xmlTextWriterStartElementNS", &XmlTextWriterStartElementNS},
	{"xmlTextWriterStartPI", &XmlTextWriterStartPI},
	{"xmlTextWriterWriteAttribute", &XmlTextWriterWriteAttribute},
	{"xmlTextWriterWriteAttributeNS", &XmlTextWriterWriteAttributeNS},
	{"xmlTextWriterWriteBase64", &XmlTextWriterWriteBase64},
	{"xmlTextWriterWriteBinHex", &XmlTextWriterWriteBinHex},
	{"xmlTextWriterWriteCDATA", &XmlTextWriterWriteCDATA},
	{"xmlTextWriterWriteComment", &XmlTextWriterWriteComment},
	{"xmlTextWriterWriteDTD", &XmlTextWriterWriteDTD},
	{"xmlTextWriterWriteDTDAttlist", &XmlTextWriterWriteDTDAttlist},
	{"xmlTextWriterWriteDTDElement", &XmlTextWriterWriteDTDElement},
	{"xmlTextWriterWriteDTDEntity", &XmlTextWriterWriteDTDEntity},
	{"xmlTextWriterWriteDTDExternalEntity", &XmlTextWriterWriteDTDExternalEntity},
	{"xmlTextWriterWriteDTDExternalEntityContents", &XmlTextWriterWriteDTDExternalEntityContents},
	{"xmlTextWriterWriteDTDInternalEntity", &XmlTextWriterWriteDTDInternalEntity},
	{"xmlTextWriterWriteDTDNotation", &XmlTextWriterWriteDTDNotation},
	{"xmlTextWriterWriteElement", &XmlTextWriterWriteElement},
	{"xmlTextWriterWriteElementNS", &XmlTextWriterWriteElementNS},
	{"xmlTextWriterWriteFormatAttribute", &XmlTextWriterWriteFormatAttribute},
	{"xmlTextWriterWriteFormatAttributeNS", &XmlTextWriterWriteFormatAttributeNS},
	{"xmlTextWriterWriteFormatCDATA", &XmlTextWriterWriteFormatCDATA},
	{"xmlTextWriterWriteFormatComment", &XmlTextWriterWriteFormatComment},
	{"xmlTextWriterWriteFormatDTD", &XmlTextWriterWriteFormatDTD},
	{"xmlTextWriterWriteFormatDTDAttlist", &XmlTextWriterWriteFormatDTDAttlist},
	{"xmlTextWriterWriteFormatDTDElement", &XmlTextWriterWriteFormatDTDElement},
	{"xmlTextWriterWriteFormatDTDInternalEntity", &XmlTextWriterWriteFormatDTDInternalEntity},
	{"xmlTextWriterWriteFormatElement", &XmlTextWriterWriteFormatElement},
	{"xmlTextWriterWriteFormatElementNS", &XmlTextWriterWriteFormatElementNS},
	{"xmlTextWriterWriteFormatPI", &XmlTextWriterWriteFormatPI},
	{"xmlTextWriterWriteFormatRaw", &XmlTextWriterWriteFormatRaw},
	{"xmlTextWriterWriteFormatString", &XmlTextWriterWriteFormatString},
	{"xmlTextWriterWritePI", &XmlTextWriterWritePI},
	{"xmlTextWriterWriteRaw", &XmlTextWriterWriteRaw},
	{"xmlTextWriterWriteRawLen", &XmlTextWriterWriteRawLen},
	{"xmlTextWriterWriteString", &XmlTextWriterWriteString},
	{"xmlTextWriterWriteVFormatAttribute", &XmlTextWriterWriteVFormatAttribute},
	{"xmlTextWriterWriteVFormatAttributeNS", &XmlTextWriterWriteVFormatAttributeNS},
	{"xmlTextWriterWriteVFormatCDATA", &XmlTextWriterWriteVFormatCDATA},
	{"xmlTextWriterWriteVFormatComment", &XmlTextWriterWriteVFormatComment},
	{"xmlTextWriterWriteVFormatDTD", &XmlTextWriterWriteVFormatDTD},
	{"xmlTextWriterWriteVFormatDTDAttlist", &XmlTextWriterWriteVFormatDTDAttlist},
	{"xmlTextWriterWriteVFormatDTDElement", &XmlTextWriterWriteVFormatDTDElement},
	{"xmlTextWriterWriteVFormatDTDInternalEntity", &XmlTextWriterWriteVFormatDTDInternalEntity},
	{"xmlTextWriterWriteVFormatElement", &XmlTextWriterWriteVFormatElement},
	{"xmlTextWriterWriteVFormatElementNS", &XmlTextWriterWriteVFormatElementNS},
	{"xmlTextWriterWriteVFormatPI", &XmlTextWriterWriteVFormatPI},
	{"xmlTextWriterWriteVFormatRaw", &XmlTextWriterWriteVFormatRaw},
	{"xmlTextWriterWriteVFormatString", &XmlTextWriterWriteVFormatString},
	{"xmlThrDefBufferAllocScheme", &XmlThrDefBufferAllocScheme},
	{"xmlThrDefDefaultBufferSize", &XmlThrDefDefaultBufferSize},
	{"xmlThrDefDeregisterNodeDefault", &XmlThrDefDeregisterNodeDefault},
	{"xmlThrDefDoValidityCheckingDefaultValue", &XmlThrDefDoValidityCheckingDefaultValue},
	{"xmlThrDefGetWarningsDefaultValue", &XmlThrDefGetWarningsDefaultValue},
	{"xmlThrDefIndentTreeOutput", &XmlThrDefIndentTreeOutput},
	{"xmlThrDefKeepBlanksDefaultValue", &XmlThrDefKeepBlanksDefaultValue},
	{"xmlThrDefLineNumbersDefaultValue", &XmlThrDefLineNumbersDefaultValue},
	{"xmlThrDefLoadExtDtdDefaultValue", &XmlThrDefLoadExtDtdDefaultValue},
	{"xmlThrDefOutputBufferCreateFilenameDefault", &XmlThrDefOutputBufferCreateFilenameDefault},
	{"xmlThrDefParserDebugEntities", &XmlThrDefParserDebugEntities},
	{"xmlThrDefParserInputBufferCreateFilenameDefault", &XmlThrDefParserInputBufferCreateFilenameDefault},
	{"xmlThrDefPedanticParserDefaultValue", &XmlThrDefPedanticParserDefaultValue},
	{"xmlThrDefRegisterNodeDefault", &XmlThrDefRegisterNodeDefault},
	{"xmlThrDefSaveNoEmptyTags", &XmlThrDefSaveNoEmptyTags},
	{"xmlThrDefSetGenericErrorFunc", &XmlThrDefSetGenericErrorFunc},
	{"xmlThrDefSetStructuredErrorFunc", &XmlThrDefSetStructuredErrorFunc},
	{"xmlThrDefSubstituteEntitiesDefaultValue", &XmlThrDefSubstituteEntitiesDefaultValue},
	{"xmlThrDefTreeIndentString", &XmlThrDefTreeIndentString},
	{"xmlUCSIsAegeanNumbers", &XmlUCSIsAegeanNumbers},
	{"xmlUCSIsAlphabeticPresentationForms", &XmlUCSIsAlphabeticPresentationForms},
	{"xmlUCSIsArabic", &XmlUCSIsArabic},
	{"xmlUCSIsArabicPresentationFormsA", &XmlUCSIsArabicPresentationFormsA},
	{"xmlUCSIsArabicPresentationFormsB", &XmlUCSIsArabicPresentationFormsB},
	{"xmlUCSIsArmenian", &XmlUCSIsArmenian},
	{"xmlUCSIsArrows", &XmlUCSIsArrows},
	{"xmlUCSIsBasicLatin", &XmlUCSIsBasicLatin},
	{"xmlUCSIsBengali", &XmlUCSIsBengali},
	{"xmlUCSIsBlock", &XmlUCSIsBlock},
	{"xmlUCSIsBlockElements", &XmlUCSIsBlockElements},
	{"xmlUCSIsBopomofo", &XmlUCSIsBopomofo},
	{"xmlUCSIsBopomofoExtended", &XmlUCSIsBopomofoExtended},
	{"xmlUCSIsBoxDrawing", &XmlUCSIsBoxDrawing},
	{"xmlUCSIsBraillePatterns", &XmlUCSIsBraillePatterns},
	{"xmlUCSIsBuhid", &XmlUCSIsBuhid},
	{"xmlUCSIsByzantineMusicalSymbols", &XmlUCSIsByzantineMusicalSymbols},
	{"xmlUCSIsCJKCompatibility", &XmlUCSIsCJKCompatibility},
	{"xmlUCSIsCJKCompatibilityForms", &XmlUCSIsCJKCompatibilityForms},
	{"xmlUCSIsCJKCompatibilityIdeographs", &XmlUCSIsCJKCompatibilityIdeographs},
	{"xmlUCSIsCJKCompatibilityIdeographsSupplement", &XmlUCSIsCJKCompatibilityIdeographsSupplement},
	{"xmlUCSIsCJKRadicalsSupplement", &XmlUCSIsCJKRadicalsSupplement},
	{"xmlUCSIsCJKSymbolsandPunctuation", &XmlUCSIsCJKSymbolsandPunctuation},
	{"xmlUCSIsCJKUnifiedIdeographs", &XmlUCSIsCJKUnifiedIdeographs},
	{"xmlUCSIsCJKUnifiedIdeographsExtensionA", &XmlUCSIsCJKUnifiedIdeographsExtensionA},
	{"xmlUCSIsCJKUnifiedIdeographsExtensionB", &XmlUCSIsCJKUnifiedIdeographsExtensionB},
	{"xmlUCSIsCat", &XmlUCSIsCat},
	{"xmlUCSIsCatC", &XmlUCSIsCatC},
	{"xmlUCSIsCatCc", &XmlUCSIsCatCc},
	{"xmlUCSIsCatCf", &XmlUCSIsCatCf},
	{"xmlUCSIsCatCo", &XmlUCSIsCatCo},
	{"xmlUCSIsCatCs", &XmlUCSIsCatCs},
	{"xmlUCSIsCatL", &XmlUCSIsCatL},
	{"xmlUCSIsCatLl", &XmlUCSIsCatLl},
	{"xmlUCSIsCatLm", &XmlUCSIsCatLm},
	{"xmlUCSIsCatLo", &XmlUCSIsCatLo},
	{"xmlUCSIsCatLt", &XmlUCSIsCatLt},
	{"xmlUCSIsCatLu", &XmlUCSIsCatLu},
	{"xmlUCSIsCatM", &XmlUCSIsCatM},
	{"xmlUCSIsCatMc", &XmlUCSIsCatMc},
	{"xmlUCSIsCatMe", &XmlUCSIsCatMe},
	{"xmlUCSIsCatMn", &XmlUCSIsCatMn},
	{"xmlUCSIsCatN", &XmlUCSIsCatN},
	{"xmlUCSIsCatNd", &XmlUCSIsCatNd},
	{"xmlUCSIsCatNl", &XmlUCSIsCatNl},
	{"xmlUCSIsCatNo", &XmlUCSIsCatNo},
	{"xmlUCSIsCatP", &XmlUCSIsCatP},
	{"xmlUCSIsCatPc", &XmlUCSIsCatPc},
	{"xmlUCSIsCatPd", &XmlUCSIsCatPd},
	{"xmlUCSIsCatPe", &XmlUCSIsCatPe},
	{"xmlUCSIsCatPf", &XmlUCSIsCatPf},
	{"xmlUCSIsCatPi", &XmlUCSIsCatPi},
	{"xmlUCSIsCatPo", &XmlUCSIsCatPo},
	{"xmlUCSIsCatPs", &XmlUCSIsCatPs},
	{"xmlUCSIsCatS", &XmlUCSIsCatS},
	{"xmlUCSIsCatSc", &XmlUCSIsCatSc},
	{"xmlUCSIsCatSk", &XmlUCSIsCatSk},
	{"xmlUCSIsCatSm", &XmlUCSIsCatSm},
	{"xmlUCSIsCatSo", &XmlUCSIsCatSo},
	{"xmlUCSIsCatZ", &XmlUCSIsCatZ},
	{"xmlUCSIsCatZl", &XmlUCSIsCatZl},
	{"xmlUCSIsCatZp", &XmlUCSIsCatZp},
	{"xmlUCSIsCatZs", &XmlUCSIsCatZs},
	{"xmlUCSIsCherokee", &XmlUCSIsCherokee},
	{"xmlUCSIsCombiningDiacriticalMarks", &XmlUCSIsCombiningDiacriticalMarks},
	{"xmlUCSIsCombiningDiacriticalMarksforSymbols", &XmlUCSIsCombiningDiacriticalMarksforSymbols},
	{"xmlUCSIsCombiningHalfMarks", &XmlUCSIsCombiningHalfMarks},
	{"xmlUCSIsCombiningMarksforSymbols", &XmlUCSIsCombiningMarksforSymbols},
	{"xmlUCSIsControlPictures", &XmlUCSIsControlPictures},
	{"xmlUCSIsCurrencySymbols", &XmlUCSIsCurrencySymbols},
	{"xmlUCSIsCypriotSyllabary", &XmlUCSIsCypriotSyllabary},
	{"xmlUCSIsCyrillic", &XmlUCSIsCyrillic},
	{"xmlUCSIsCyrillicSupplement", &XmlUCSIsCyrillicSupplement},
	{"xmlUCSIsDeseret", &XmlUCSIsDeseret},
	{"xmlUCSIsDevanagari", &XmlUCSIsDevanagari},
	{"xmlUCSIsDingbats", &XmlUCSIsDingbats},
	{"xmlUCSIsEnclosedAlphanumerics", &XmlUCSIsEnclosedAlphanumerics},
	{"xmlUCSIsEnclosedCJKLettersandMonths", &XmlUCSIsEnclosedCJKLettersandMonths},
	{"xmlUCSIsEthiopic", &XmlUCSIsEthiopic},
	{"xmlUCSIsGeneralPunctuation", &XmlUCSIsGeneralPunctuation},
	{"xmlUCSIsGeometricShapes", &XmlUCSIsGeometricShapes},
	{"xmlUCSIsGeorgian", &XmlUCSIsGeorgian},
	{"xmlUCSIsGothic", &XmlUCSIsGothic},
	{"xmlUCSIsGreek", &XmlUCSIsGreek},
	{"xmlUCSIsGreekExtended", &XmlUCSIsGreekExtended},
	{"xmlUCSIsGreekandCoptic", &XmlUCSIsGreekandCoptic},
	{"xmlUCSIsGujarati", &XmlUCSIsGujarati},
	{"xmlUCSIsGurmukhi", &XmlUCSIsGurmukhi},
	{"xmlUCSIsHalfwidthandFullwidthForms", &XmlUCSIsHalfwidthandFullwidthForms},
	{"xmlUCSIsHangulCompatibilityJamo", &XmlUCSIsHangulCompatibilityJamo},
	{"xmlUCSIsHangulJamo", &XmlUCSIsHangulJamo},
	{"xmlUCSIsHangulSyllables", &XmlUCSIsHangulSyllables},
	{"xmlUCSIsHanunoo", &XmlUCSIsHanunoo},
	{"xmlUCSIsHebrew", &XmlUCSIsHebrew},
	{"xmlUCSIsHighPrivateUseSurrogates", &XmlUCSIsHighPrivateUseSurrogates},
	{"xmlUCSIsHighSurrogates", &XmlUCSIsHighSurrogates},
	{"xmlUCSIsHiragana", &XmlUCSIsHiragana},
	{"xmlUCSIsIPAExtensions", &XmlUCSIsIPAExtensions},
	{"xmlUCSIsIdeographicDescriptionCharacters", &XmlUCSIsIdeographicDescriptionCharacters},
	{"xmlUCSIsKanbun", &XmlUCSIsKanbun},
	{"xmlUCSIsKangxiRadicals", &XmlUCSIsKangxiRadicals},
	{"xmlUCSIsKannada", &XmlUCSIsKannada},
	{"xmlUCSIsKatakana", &XmlUCSIsKatakana},
	{"xmlUCSIsKatakanaPhoneticExtensions", &XmlUCSIsKatakanaPhoneticExtensions},
	{"xmlUCSIsKhmer", &XmlUCSIsKhmer},
	{"xmlUCSIsKhmerSymbols", &XmlUCSIsKhmerSymbols},
	{"xmlUCSIsLao", &XmlUCSIsLao},
	{"xmlUCSIsLatin1Supplement", &XmlUCSIsLatin1Supplement},
	{"xmlUCSIsLatinExtendedA", &XmlUCSIsLatinExtendedA},
	{"xmlUCSIsLatinExtendedAdditional", &XmlUCSIsLatinExtendedAdditional},
	{"xmlUCSIsLatinExtendedB", &XmlUCSIsLatinExtendedB},
	{"xmlUCSIsLetterlikeSymbols", &XmlUCSIsLetterlikeSymbols},
	{"xmlUCSIsLimbu", &XmlUCSIsLimbu},
	{"xmlUCSIsLinearBIdeograms", &XmlUCSIsLinearBIdeograms},
	{"xmlUCSIsLinearBSyllabary", &XmlUCSIsLinearBSyllabary},
	{"xmlUCSIsLowSurrogates", &XmlUCSIsLowSurrogates},
	{"xmlUCSIsMalayalam", &XmlUCSIsMalayalam},
	{"xmlUCSIsMathematicalAlphanumericSymbols", &XmlUCSIsMathematicalAlphanumericSymbols},
	{"xmlUCSIsMathematicalOperators", &XmlUCSIsMathematicalOperators},
	{"xmlUCSIsMiscellaneousMathematicalSymbolsA", &XmlUCSIsMiscellaneousMathematicalSymbolsA},
	{"xmlUCSIsMiscellaneousMathematicalSymbolsB", &XmlUCSIsMiscellaneousMathematicalSymbolsB},
	{"xmlUCSIsMiscellaneousSymbols", &XmlUCSIsMiscellaneousSymbols},
	{"xmlUCSIsMiscellaneousSymbolsandArrows", &XmlUCSIsMiscellaneousSymbolsandArrows},
	{"xmlUCSIsMiscellaneousTechnical", &XmlUCSIsMiscellaneousTechnical},
	{"xmlUCSIsMongolian", &XmlUCSIsMongolian},
	{"xmlUCSIsMusicalSymbols", &XmlUCSIsMusicalSymbols},
	{"xmlUCSIsMyanmar", &XmlUCSIsMyanmar},
	{"xmlUCSIsNumberForms", &XmlUCSIsNumberForms},
	{"xmlUCSIsOgham", &XmlUCSIsOgham},
	{"xmlUCSIsOldItalic", &XmlUCSIsOldItalic},
	{"xmlUCSIsOpticalCharacterRecognition", &XmlUCSIsOpticalCharacterRecognition},
	{"xmlUCSIsOriya", &XmlUCSIsOriya},
	{"xmlUCSIsOsmanya", &XmlUCSIsOsmanya},
	{"xmlUCSIsPhoneticExtensions", &XmlUCSIsPhoneticExtensions},
	{"xmlUCSIsPrivateUse", &XmlUCSIsPrivateUse},
	{"xmlUCSIsPrivateUseArea", &XmlUCSIsPrivateUseArea},
	{"xmlUCSIsRunic", &XmlUCSIsRunic},
	{"xmlUCSIsShavian", &XmlUCSIsShavian},
	{"xmlUCSIsSinhala", &XmlUCSIsSinhala},
	{"xmlUCSIsSmallFormVariants", &XmlUCSIsSmallFormVariants},
	{"xmlUCSIsSpacingModifierLetters", &XmlUCSIsSpacingModifierLetters},
	{"xmlUCSIsSpecials", &XmlUCSIsSpecials},
	{"xmlUCSIsSuperscriptsandSubscripts", &XmlUCSIsSuperscriptsandSubscripts},
	{"xmlUCSIsSupplementalArrowsA", &XmlUCSIsSupplementalArrowsA},
	{"xmlUCSIsSupplementalArrowsB", &XmlUCSIsSupplementalArrowsB},
	{"xmlUCSIsSupplementalMathematicalOperators", &XmlUCSIsSupplementalMathematicalOperators},
	{"xmlUCSIsSupplementaryPrivateUseAreaA", &XmlUCSIsSupplementaryPrivateUseAreaA},
	{"xmlUCSIsSupplementaryPrivateUseAreaB", &XmlUCSIsSupplementaryPrivateUseAreaB},
	{"xmlUCSIsSyriac", &XmlUCSIsSyriac},
	{"xmlUCSIsTagalog", &XmlUCSIsTagalog},
	{"xmlUCSIsTagbanwa", &XmlUCSIsTagbanwa},
	{"xmlUCSIsTags", &XmlUCSIsTags},
	{"xmlUCSIsTaiLe", &XmlUCSIsTaiLe},
	{"xmlUCSIsTaiXuanJingSymbols", &XmlUCSIsTaiXuanJingSymbols},
	{"xmlUCSIsTamil", &XmlUCSIsTamil},
	{"xmlUCSIsTelugu", &XmlUCSIsTelugu},
	{"xmlUCSIsThaana", &XmlUCSIsThaana},
	{"xmlUCSIsThai", &XmlUCSIsThai},
	{"xmlUCSIsTibetan", &XmlUCSIsTibetan},
	{"xmlUCSIsUgaritic", &XmlUCSIsUgaritic},
	{"xmlUCSIsUnifiedCanadianAboriginalSyllabics", &XmlUCSIsUnifiedCanadianAboriginalSyllabics},
	{"xmlUCSIsVariationSelectors", &XmlUCSIsVariationSelectors},
	{"xmlUCSIsVariationSelectorsSupplement", &XmlUCSIsVariationSelectorsSupplement},
	{"xmlUCSIsYiRadicals", &XmlUCSIsYiRadicals},
	{"xmlUCSIsYiSyllables", &XmlUCSIsYiSyllables},
	{"xmlUCSIsYijingHexagramSymbols", &XmlUCSIsYijingHexagramSymbols},
	{"xmlURIEscape", &XmlURIEscape},
	{"xmlURIEscapeStr", &XmlURIEscapeStr},
	{"xmlURIUnescapeString", &XmlURIUnescapeString},
	{"xmlUTF8Charcmp", &XmlUTF8Charcmp},
	{"xmlUTF8Size", &XmlUTF8Size},
	{"xmlUTF8Strlen", &XmlUTF8Strlen},
	{"xmlUTF8Strloc", &XmlUTF8Strloc},
	{"xmlUTF8Strndup", &XmlUTF8Strndup},
	{"xmlUTF8Strpos", &XmlUTF8Strpos},
	{"xmlUTF8Strsize", &XmlUTF8Strsize},
	{"xmlUTF8Strsub", &XmlUTF8Strsub},
	{"xmlUnlinkNode", &XmlUnlinkNode},
	{"xmlUnlockLibrary", &XmlUnlockLibrary},
	{"xmlUnsetNsProp", &XmlUnsetNsProp},
	{"xmlUnsetProp", &XmlUnsetProp},
	{"xmlValidBuildContentModel", &XmlValidBuildContentModel},
	{"xmlValidCtxtNormalizeAttributeValue", &XmlValidCtxtNormalizeAttributeValue},
	{"xmlValidGetPotentialChildren", &XmlValidGetPotentialChildren},
	{"xmlValidGetValidElements", &XmlValidGetValidElements},
	{"xmlValidNormalizeAttributeValue", &XmlValidNormalizeAttributeValue},
	{"xmlValidateAttributeDecl", &XmlValidateAttributeDecl},
	{"xmlValidateAttributeValue", &XmlValidateAttributeValue},
	{"xmlValidateDocument", &XmlValidateDocument},
	{"xmlValidateDocumentFinal", &XmlValidateDocumentFinal},
	{"xmlValidateDtd", &XmlValidateDtd},
	{"xmlValidateDtdFinal", &XmlValidateDtdFinal},
	{"xmlValidateElement", &XmlValidateElement},
	{"xmlValidateElementDecl", &XmlValidateElementDecl},
	{"xmlValidateNCName", &XmlValidateNCName},
	{"xmlValidateNMToken", &XmlValidateNMToken},
	{"xmlValidateName", &XmlValidateName},
	{"xmlValidateNameValue", &XmlValidateNameValue},
	{"xmlValidateNamesValue", &XmlValidateNamesValue},
	{"xmlValidateNmtokenValue", &XmlValidateNmtokenValue},
	{"xmlValidateNmtokensValue", &XmlValidateNmtokensValue},
	{"xmlValidateNotationDecl", &XmlValidateNotationDecl},
	{"xmlValidateNotationUse", &XmlValidateNotationUse},
	{"xmlValidateOneAttribute", &XmlValidateOneAttribute},
	{"xmlValidateOneElement", &XmlValidateOneElement},
	{"xmlValidateOneNamespace", &XmlValidateOneNamespace},
	{"xmlValidatePopElement", &XmlValidatePopElement},
	{"xmlValidatePushCData", &XmlValidatePushCData},
	{"xmlValidatePushElement", &XmlValidatePushElement},
	{"xmlValidateQName", &XmlValidateQName},
	{"xmlValidateRoot", &XmlValidateRoot},
	{"xmlXIncludeFreeContext", &XmlXIncludeFreeContext},
	{"xmlXIncludeNewContext", &XmlXIncludeNewContext},
	{"xmlXIncludeProcess", &XmlXIncludeProcess},
	{"xmlXIncludeProcessFlags", &XmlXIncludeProcessFlags},
	{"xmlXIncludeProcessFlagsData", &XmlXIncludeProcessFlagsData},
	{"xmlXIncludeProcessNode", &XmlXIncludeProcessNode},
	{"xmlXIncludeProcessTree", &XmlXIncludeProcessTree},
	{"xmlXIncludeProcessTreeFlags", &XmlXIncludeProcessTreeFlags},
	{"xmlXIncludeProcessTreeFlagsData", &XmlXIncludeProcessTreeFlagsData},
	{"xmlXIncludeSetFlags", &XmlXIncludeSetFlags},
	{"xmlXPathAddValues", &XmlXPathAddValues},
	{"xmlXPathBooleanFunction", &XmlXPathBooleanFunction},
	{"xmlXPathCastBooleanToNumber", &XmlXPathCastBooleanToNumber},
	{"xmlXPathCastBooleanToString", &XmlXPathCastBooleanToString},
	{"xmlXPathCastNodeSetToBoolean", &XmlXPathCastNodeSetToBoolean},
	{"xmlXPathCastNodeSetToNumber", &XmlXPathCastNodeSetToNumber},
	{"xmlXPathCastNodeSetToString", &XmlXPathCastNodeSetToString},
	{"xmlXPathCastNodeToNumber", &XmlXPathCastNodeToNumber},
	{"xmlXPathCastNodeToString", &XmlXPathCastNodeToString},
	{"xmlXPathCastNumberToBoolean", &XmlXPathCastNumberToBoolean},
	{"xmlXPathCastNumberToString", &XmlXPathCastNumberToString},
	{"xmlXPathCastStringToBoolean", &XmlXPathCastStringToBoolean},
	{"xmlXPathCastStringToNumber", &XmlXPathCastStringToNumber},
	{"xmlXPathCastToBoolean", &XmlXPathCastToBoolean},
	{"xmlXPathCastToNumber", &XmlXPathCastToNumber},
	{"xmlXPathCastToString", &XmlXPathCastToString},
	{"xmlXPathCeilingFunction", &XmlXPathCeilingFunction},
	{"xmlXPathCmpNodes", &XmlXPathCmpNodes},
	{"xmlXPathCompareValues", &XmlXPathCompareValues},
	{"xmlXPathCompile", &XmlXPathCompile},
	{"xmlXPathCompiledEval", &XmlXPathCompiledEval},
	{"xmlXPathCompiledEvalToBoolean", &XmlXPathCompiledEvalToBoolean},
	{"xmlXPathConcatFunction", &XmlXPathConcatFunction},
	{"xmlXPathContainsFunction", &XmlXPathContainsFunction},
	{"xmlXPathContextSetCache", &XmlXPathContextSetCache},
	{"xmlXPathConvertBoolean", &XmlXPathConvertBoolean},
	{"xmlXPathConvertNumber", &XmlXPathConvertNumber},
	{"xmlXPathConvertString", &XmlXPathConvertString},
	{"xmlXPathCountFunction", &XmlXPathCountFunction},
	{"xmlXPathCtxtCompile", &XmlXPathCtxtCompile},
	{"xmlXPathDebugDumpCompExpr", &XmlXPathDebugDumpCompExpr},
	{"xmlXPathDebugDumpObject", &XmlXPathDebugDumpObject},
	{"xmlXPathDifference", &XmlXPathDifference},
	{"xmlXPathDistinct", &XmlXPathDistinct},
	{"xmlXPathDistinctSorted", &XmlXPathDistinctSorted},
	{"xmlXPathDivValues", &XmlXPathDivValues},
	{"xmlXPathEqualValues", &XmlXPathEqualValues},
	{"xmlXPathErr", &XmlXPathErr},
	{"xmlXPathEval", &XmlXPathEval},
	{"xmlXPathEvalExpr", &XmlXPathEvalExpr},
	{"xmlXPathEvalExpression", &XmlXPathEvalExpression},
	{"xmlXPathEvalPredicate", &XmlXPathEvalPredicate},
	{"xmlXPathEvaluatePredicateResult", &XmlXPathEvaluatePredicateResult},
	{"xmlXPathFalseFunction", &XmlXPathFalseFunction},
	{"xmlXPathFloorFunction", &XmlXPathFloorFunction},
	{"xmlXPathFreeCompExpr", &XmlXPathFreeCompExpr},
	{"xmlXPathFreeContext", &XmlXPathFreeContext},
	{"xmlXPathFreeNodeSet", &XmlXPathFreeNodeSet},
	{"xmlXPathFreeNodeSetList", &XmlXPathFreeNodeSetList},
	{"xmlXPathFreeObject", &XmlXPathFreeObject},
	{"xmlXPathFreeParserContext", &XmlXPathFreeParserContext},
	{"xmlXPathFunctionLookup", &XmlXPathFunctionLookup},
	{"xmlXPathFunctionLookupNS", &XmlXPathFunctionLookupNS},
	{"xmlXPathHasSameNodes", &XmlXPathHasSameNodes},
	{"xmlXPathIdFunction", &XmlXPathIdFunction},
	{"xmlXPathInit", &XmlXPathInit},
	{"xmlXPathIntersection", &XmlXPathIntersection},
	{"xmlXPathIsInf", &XmlXPathIsInf},
	{"xmlXPathIsNaN", &XmlXPathIsNaN},
	{"xmlXPathIsNodeType", &XmlXPathIsNodeType},
	{"xmlXPathLangFunction", &XmlXPathLangFunction},
	{"xmlXPathLastFunction", &XmlXPathLastFunction},
	{"xmlXPathLeading", &XmlXPathLeading},
	{"xmlXPathLeadingSorted", &XmlXPathLeadingSorted},
	{"xmlXPathLocalNameFunction", &XmlXPathLocalNameFunction},
	{"xmlXPathModValues", &XmlXPathModValues},
	{"xmlXPathMultValues", &XmlXPathMultValues},
	{"xmlXPathNamespaceURIFunction", &XmlXPathNamespaceURIFunction},
	{"xmlXPathNewBoolean", &XmlXPathNewBoolean},
	{"xmlXPathNewCString", &XmlXPathNewCString},
	{"xmlXPathNewContext", &XmlXPathNewContext},
	{"xmlXPathNewFloat", &XmlXPathNewFloat},
	{"xmlXPathNewNodeSet", &XmlXPathNewNodeSet},
	{"xmlXPathNewNodeSetList", &XmlXPathNewNodeSetList},
	{"xmlXPathNewParserContext", &XmlXPathNewParserContext},
	{"xmlXPathNewString", &XmlXPathNewString},
	{"xmlXPathNewValueTree", &XmlXPathNewValueTree},
	{"xmlXPathNextAncestor", &XmlXPathNextAncestor},
	{"xmlXPathNextAncestorOrSelf", &XmlXPathNextAncestorOrSelf},
	{"xmlXPathNextAttribute", &XmlXPathNextAttribute},
	{"xmlXPathNextChild", &XmlXPathNextChild},
	{"xmlXPathNextDescendant", &XmlXPathNextDescendant},
	{"xmlXPathNextDescendantOrSelf", &XmlXPathNextDescendantOrSelf},
	{"xmlXPathNextFollowing", &XmlXPathNextFollowing},
	{"xmlXPathNextFollowingSibling", &XmlXPathNextFollowingSibling},
	{"xmlXPathNextNamespace", &XmlXPathNextNamespace},
	{"xmlXPathNextParent", &XmlXPathNextParent},
	{"xmlXPathNextPreceding", &XmlXPathNextPreceding},
	{"xmlXPathNextPrecedingSibling", &XmlXPathNextPrecedingSibling},
	{"xmlXPathNextSelf", &XmlXPathNextSelf},
	{"xmlXPathNodeLeading", &XmlXPathNodeLeading},
	{"xmlXPathNodeLeadingSorted", &XmlXPathNodeLeadingSorted},
	{"xmlXPathNodeSetAdd", &XmlXPathNodeSetAdd},
	{"xmlXPathNodeSetAddNs", &XmlXPathNodeSetAddNs},
	{"xmlXPathNodeSetAddUnique", &XmlXPathNodeSetAddUnique},
	{"xmlXPathNodeSetContains", &XmlXPathNodeSetContains},
	{"xmlXPathNodeSetCreate", &XmlXPathNodeSetCreate},
	{"xmlXPathNodeSetDel", &XmlXPathNodeSetDel},
	{"xmlXPathNodeSetFreeNs", &XmlXPathNodeSetFreeNs},
	{"xmlXPathNodeSetMerge", &XmlXPathNodeSetMerge},
	{"xmlXPathNodeSetRemove", &XmlXPathNodeSetRemove},
	{"xmlXPathNodeSetSort", &XmlXPathNodeSetSort},
	{"xmlXPathNodeTrailing", &XmlXPathNodeTrailing},
	{"xmlXPathNodeTrailingSorted", &XmlXPathNodeTrailingSorted},
	{"xmlXPathNormalizeFunction", &XmlXPathNormalizeFunction},
	{"xmlXPathNotEqualValues", &XmlXPathNotEqualValues},
	{"xmlXPathNotFunction", &XmlXPathNotFunction},
	{"xmlXPathNsLookup", &XmlXPathNsLookup},
	{"xmlXPathNumberFunction", &XmlXPathNumberFunction},
	{"xmlXPathObjectCopy", &XmlXPathObjectCopy},
	{"xmlXPathOrderDocElems", &XmlXPathOrderDocElems},
	{"xmlXPathParseNCName", &XmlXPathParseNCName},
	{"xmlXPathParseName", &XmlXPathParseName},
	{"xmlXPathPopBoolean", &XmlXPathPopBoolean},
	{"xmlXPathPopExternal", &XmlXPathPopExternal},
	{"xmlXPathPopNodeSet", &XmlXPathPopNodeSet},
	{"xmlXPathPopNumber", &XmlXPathPopNumber},
	{"xmlXPathPopString", &XmlXPathPopString},
	{"xmlXPathPositionFunction", &XmlXPathPositionFunction},
	{"xmlXPathRegisterAllFunctions", &XmlXPathRegisterAllFunctions},
	{"xmlXPathRegisterFunc", &XmlXPathRegisterFunc},
	{"xmlXPathRegisterFuncLookup", &XmlXPathRegisterFuncLookup},
	{"xmlXPathRegisterFuncNS", &XmlXPathRegisterFuncNS},
	{"xmlXPathRegisterNs", &XmlXPathRegisterNs},
	{"xmlXPathRegisterVariable", &XmlXPathRegisterVariable},
	{"xmlXPathRegisterVariableLookup", &XmlXPathRegisterVariableLookup},
	{"xmlXPathRegisterVariableNS", &XmlXPathRegisterVariableNS},
	{"xmlXPathRegisteredFuncsCleanup", &XmlXPathRegisteredFuncsCleanup},
	{"xmlXPathRegisteredNsCleanup", &XmlXPathRegisteredNsCleanup},
	{"xmlXPathRegisteredVariablesCleanup", &XmlXPathRegisteredVariablesCleanup},
	{"xmlXPathRoot", &XmlXPathRoot},
	{"xmlXPathRoundFunction", &XmlXPathRoundFunction},
	{"xmlXPathStartsWithFunction", &XmlXPathStartsWithFunction},
	{"xmlXPathStringEvalNumber", &XmlXPathStringEvalNumber},
	{"xmlXPathStringFunction", &XmlXPathStringFunction},
	{"xmlXPathStringLengthFunction", &XmlXPathStringLengthFunction},
	{"xmlXPathSubValues", &XmlXPathSubValues},
	{"xmlXPathSubstringAfterFunction", &XmlXPathSubstringAfterFunction},
	{"xmlXPathSubstringBeforeFunction", &XmlXPathSubstringBeforeFunction},
	{"xmlXPathSubstringFunction", &XmlXPathSubstringFunction},
	{"xmlXPathSumFunction", &XmlXPathSumFunction},
	{"xmlXPathTrailing", &XmlXPathTrailing},
	{"xmlXPathTrailingSorted", &XmlXPathTrailingSorted},
	{"xmlXPathTranslateFunction", &XmlXPathTranslateFunction},
	{"xmlXPathTrueFunction", &XmlXPathTrueFunction},
	{"xmlXPathValueFlipSign", &XmlXPathValueFlipSign},
	{"xmlXPathVariableLookup", &XmlXPathVariableLookup},
	{"xmlXPathVariableLookupNS", &XmlXPathVariableLookupNS},
	{"xmlXPathWrapCString", &XmlXPathWrapCString},
	{"xmlXPathWrapExternal", &XmlXPathWrapExternal},
	{"xmlXPathWrapNodeSet", &XmlXPathWrapNodeSet},
	{"xmlXPathWrapString", &XmlXPathWrapString},
	{"xmlXPatherror", &XmlXPatherror},
	{"xmlXPtrBuildNodeList", &XmlXPtrBuildNodeList},
	{"xmlXPtrEval", &XmlXPtrEval},
	{"xmlXPtrEvalRangePredicate", &XmlXPtrEvalRangePredicate},
	{"xmlXPtrFreeLocationSet", &XmlXPtrFreeLocationSet},
	{"xmlXPtrLocationSetAdd", &XmlXPtrLocationSetAdd},
	{"xmlXPtrLocationSetCreate", &XmlXPtrLocationSetCreate},
	{"xmlXPtrLocationSetDel", &XmlXPtrLocationSetDel},
	{"xmlXPtrLocationSetMerge", &XmlXPtrLocationSetMerge},
	{"xmlXPtrLocationSetRemove", &XmlXPtrLocationSetRemove},
	{"xmlXPtrNewCollapsedRange", &XmlXPtrNewCollapsedRange},
	{"xmlXPtrNewContext", &XmlXPtrNewContext},
	{"xmlXPtrNewLocationSetNodeSet", &XmlXPtrNewLocationSetNodeSet},
	{"xmlXPtrNewLocationSetNodes", &XmlXPtrNewLocationSetNodes},
	{"xmlXPtrNewRange", &XmlXPtrNewRange},
	{"xmlXPtrNewRangeNodeObject", &XmlXPtrNewRangeNodeObject},
	{"xmlXPtrNewRangeNodePoint", &XmlXPtrNewRangeNodePoint},
	{"xmlXPtrNewRangeNodes", &XmlXPtrNewRangeNodes},
	{"xmlXPtrNewRangePointNode", &XmlXPtrNewRangePointNode},
	{"xmlXPtrNewRangePoints", &XmlXPtrNewRangePoints},
	{"xmlXPtrRangeToFunction", &XmlXPtrRangeToFunction},
	{"xmlXPtrWrapLocationSet", &XmlXPtrWrapLocationSet},
}

var dataList = Data{
// {"emptyExp", new(emptyExp)},
// {"forbiddenExp", new(forbiddenExp)},
// {"xmlFree", new(XmlFree)},
// {"xmlIsBaseCharGroup", new(XmlIsBaseCharGroup)},
// {"xmlIsCharGroup", new(XmlIsCharGroup)},
// {"xmlIsCombiningGroup", new(XmlIsCombiningGroup)},
// {"xmlIsDigitGroup", new(XmlIsDigitGroup)},
// {"xmlIsExtenderGroup", new(XmlIsExtenderGroup)},
// {"xmlIsIdeographicGroup", new(XmlIsIdeographicGroup)},
// {"xmlIsPubidChar_tab", new(XmlIsPubidChar_tab)},
// {"xmlMalloc", new(XmlMalloc)},
// {"xmlMallocAtomic", new(XmlMallocAtomic)},
// {"xmlMemStrdup", new(XmlMemStrdup)},
// {"xmlParserMaxDepth", new(XmlParserMaxDepth)},
// {"xmlRealloc", new(XmlRealloc)},
// {"xmlStringComment", new(XmlStringComment)},
// {"xmlStringText", new(XmlStringText)},
// {"xmlStringTextNoenc", new(XmlStringTextNoenc)},
// {"xmlXPathNAN", new(XmlXPathNAN)},
// {"xmlXPathNINF", new(XmlXPathNINF)},
// {"xmlXPathPINF", new(XmlXPathPINF)},
}
