// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package haru

/*
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

const (
	hpdfVersion = "2.4.0dev"
)

type (
	hpdfDoc          C.HPDF_Doc
	hpdfPage         C.HPDF_Page
	hpdfPages        C.HPDF_Pages
	hpdfStream       C.HPDF_Stream
	hpdfImage        C.HPDF_Image
	hpdfFont         C.HPDF_Font
	hpdfOutline      C.HPDF_Outline
	hpdfEncoder      C.HPDF_Encoder
	hpdf3DMeasure    C.HPDF_3DMeasure
	hpdfExData       C.HPDF_ExData
	hpdfDestination  C.HPDF_Destination
	hpdfXObject      C.HPDF_XObject
	hpdfAnnotation   C.HPDF_Annotation
	hpdfExtGState    C.HPDF_ExtGState
	hpdfFontDef      C.HPDF_FontDef
	hpdfU3D          C.HPDF_U3D
	hpdfJavaScript   C.HPDF_JavaScript
	hpdfError        C.HPDF_Error
	hpdfMMgr         C.HPDF_MMgr
	hpdfDict         C.HPDF_Dict
	hpdfEmbeddedFile C.HPDF_EmbeddedFile
	hpdfOutputIntent C.HPDF_OutputIntent
	hpdfXref         C.HPDF_Xref
	hpdfStatus       uint
)

func hpdfGetVersion() string {
	p := C.HPDF_GetVersion()
	s := C.GoString(p)
	return s
}

func hpdfNewEx() hpdfDoc {
	panic("disable")
}

func hpdfNew() hpdfDoc {
	p := C.HPDF_New(nil, nil)
	return hpdfDoc(p)
}

func hpdfSetErrorHandler() {
	panic("disable")
}

func hpdfFree(p hpdfDoc) {
	C.HPDF_Free(C.HPDF_Doc(p))
}

func hpdfNewDoc(p hpdfDoc) hpdfStatus {
	t := C.HPDF_NewDoc(C.HPDF_Doc(p))
	return hpdfStatus(t)
}

func hpdfFreeDoc(p hpdfDoc) {
	C.HPDF_FreeDoc(C.HPDF_Doc(p))
}

func hpdfHasDoc(p hpdfDoc) bool {
	t := C.HPDF_HasDoc(C.HPDF_Doc(p))
	return t != 0
}

func hpdfFreeDocAll(p hpdfDoc) {
	C.HPDF_FreeDocAll(C.HPDF_Doc(p))
}

/*


HPDF_EXPORT(HPDF_STATUS)
HPDF_SaveToStream  (HPDF_Doc   pdf);

HPDF_EXPORT(HPDF_STATUS)
HPDF_GetContents   (HPDF_Doc   pdf,
                   HPDF_BYTE  *buf,
                 HPDF_UINT32  *size);

HPDF_EXPORT(HPDF_UINT32)
HPDF_GetStreamSize  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_ReadFromStream  (HPDF_Doc       pdf,
                      HPDF_BYTE     *buf,
                      HPDF_UINT32   *size);


HPDF_EXPORT(HPDF_STATUS)
HPDF_ResetStream  (HPDF_Doc     pdf);

*/

func hpdfSaveToFile(pdf hpdfDoc, filename string) hpdfStatus {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	t := C.HPDF_SaveToFile(C.HPDF_Doc(pdf), cs)
	return hpdfStatus(t)
}

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_GetError  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_GetErrorDetail  (HPDF_Doc   pdf);

HPDF_EXPORT(void)
HPDF_ResetError  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_CheckError  (HPDF_Error   error);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetPagesConfiguration  (HPDF_Doc    pdf,
                             HPDF_UINT   page_per_pages);


HPDF_EXPORT(HPDF_Page)
HPDF_GetPageByIndex  (HPDF_Doc    pdf,
                      HPDF_UINT   index);

*/

/*---------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_MMgr)
HPDF_GetPageMMgr    (HPDF_Page  page);

HPDF_EXPORT(HPDF_PageLayout)
HPDF_GetPageLayout  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetPageLayout  (HPDF_Doc          pdf,
                     HPDF_PageLayout   layout);


HPDF_EXPORT(HPDF_PageMode)
HPDF_GetPageMode  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetPageMode  (HPDF_Doc        pdf,
                   HPDF_PageMode   mode);


HPDF_EXPORT(HPDF_UINT)
HPDF_GetViewerPreference  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetViewerPreference  (HPDF_Doc     pdf,
                           HPDF_UINT    value);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetOpenAction  (HPDF_Doc           pdf,
                     HPDF_Destination   open_action);
*/

/*---------------------------------------------------------------------------*/
/*----- page handling -------------------------------------------------------*/

func hpdfGetCurrentPage(pdf hpdfDoc) hpdfPage {
	t := C.HPDF_GetCurrentPage(C.HPDF_Doc(pdf))
	return hpdfPage(t)
}

func hpdfAddPage(pdf hpdfDoc) hpdfPage {
	t := C.HPDF_AddPage(C.HPDF_Doc(pdf))
	return hpdfPage(t)
}

func hpdfInsertPage(pdf hpdfDoc, page hpdfPage) hpdfPage {
	t := C.HPDF_InsertPage(C.HPDF_Doc(pdf), C.HPDF_Page(page))
	return hpdfPage(t)
}

func hpdfPageSetWidth(page hpdfPage, value float32) hpdfStatus {
	t := C.HPDF_Page_SetWidth(C.HPDF_Page(page), C.HPDF_REAL(value))
	return hpdfStatus(t)
}

func hpdfPageSetHeight(page hpdfPage, value float32) hpdfStatus {
	t := C.HPDF_Page_SetHeight(C.HPDF_Page(page), C.HPDF_REAL(value))
	return hpdfStatus(t)
}

/*

HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetSize  (HPDF_Page            page,
                    HPDF_PageSizes       size,
                    HPDF_PageDirection   direction);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetRotate  (HPDF_Page     page,
                      HPDF_UINT16   angle);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetZoom  (HPDF_Page     page,
                    HPDF_REAL     zoom);
*/

/*---------------------------------------------------------------------------*/
/*----- font handling -------------------------------------------------------*/

func hpdfGetFont(pdf hpdfDoc, fontName, encodingName string) hpdfFont {
	font_name := C.CString(fontName)
	defer C.free(unsafe.Pointer(font_name))
	encoding_name := C.CString(encodingName)
	defer C.free(unsafe.Pointer(encoding_name))
	t := C.HPDF_GetFont(
		C.HPDF_Doc(pdf),
		font_name,
		encoding_name,
	)
	return hpdfFont(t)
}

/*

HPDF_EXPORT(const char*)
HPDF_LoadType1FontFromFile  (HPDF_Doc     pdf,
                             const char  *afm_file_name,
                             const char  *data_file_name);


HPDF_EXPORT(HPDF_FontDef)
HPDF_GetTTFontDefFromFile (HPDF_Doc     pdf,
                           const char  *file_name,
                           HPDF_BOOL    embedding);

HPDF_EXPORT(const char*)
HPDF_LoadTTFontFromFile (HPDF_Doc     pdf,
                         const char  *file_name,
                         HPDF_BOOL    embedding);


HPDF_EXPORT(const char*)
HPDF_LoadTTFontFromFile2 (HPDF_Doc     pdf,
                          const char  *file_name,
                          HPDF_UINT    index,
                          HPDF_BOOL    embedding);


HPDF_EXPORT(HPDF_STATUS)
HPDF_AddPageLabel  (HPDF_Doc            pdf,
                    HPDF_UINT           page_num,
                    HPDF_PageNumStyle   style,
                    HPDF_UINT           first_page,
                    const char         *prefix);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseJPFonts   (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseKRFonts   (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseCNSFonts   (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseCNTFonts   (HPDF_Doc   pdf);

*/

/*--------------------------------------------------------------------------*/
/*----- outline ------------------------------------------------------------*/

/*

HPDF_EXPORT(HPDF_Outline)
HPDF_CreateOutline  (HPDF_Doc       pdf,
                     HPDF_Outline   parent,
                     const char    *title,
                     HPDF_Encoder   encoder);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Outline_SetOpened  (HPDF_Outline  outline,
                         HPDF_BOOL     opened);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Outline_SetDestination (HPDF_Outline      outline,
                             HPDF_Destination  dst);

*/

/*--------------------------------------------------------------------------*/
/*----- destination --------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_Destination)
HPDF_Page_CreateDestination  (HPDF_Page   page);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetXYZ  (HPDF_Destination  dst,
                          HPDF_REAL         left,
                          HPDF_REAL         top,
                          HPDF_REAL         zoom);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFit  (HPDF_Destination  dst);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitH  (HPDF_Destination  dst,
                           HPDF_REAL         top);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitV  (HPDF_Destination  dst,
                           HPDF_REAL         left);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitR  (HPDF_Destination  dst,
                           HPDF_REAL         left,
                           HPDF_REAL         bottom,
                           HPDF_REAL         right,
                           HPDF_REAL         top);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitB  (HPDF_Destination  dst);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitBH  (HPDF_Destination  dst,
                            HPDF_REAL         top);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Destination_SetFitBV  (HPDF_Destination  dst,
                            HPDF_REAL         left);

*/

/*--------------------------------------------------------------------------*/
/*----- encoder ------------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_Encoder)
HPDF_GetEncoder  (HPDF_Doc     pdf,
                  const char  *encoding_name);


HPDF_EXPORT(HPDF_Encoder)
HPDF_GetCurrentEncoder  (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetCurrentEncoder  (HPDF_Doc     pdf,
                         const char  *encoding_name);


HPDF_EXPORT(HPDF_EncoderType)
HPDF_Encoder_GetType  (HPDF_Encoder   encoder);


HPDF_EXPORT(HPDF_ByteType)
HPDF_Encoder_GetByteType  (HPDF_Encoder    encoder,
                           const char     *text,
                           HPDF_UINT       index);


HPDF_EXPORT(HPDF_UNICODE)
HPDF_Encoder_GetUnicode  (HPDF_Encoder   encoder,
                          HPDF_UINT16    code);


HPDF_EXPORT(HPDF_WritingMode)
HPDF_Encoder_GetWritingMode (HPDF_Encoder   encoder);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseJPEncodings   (HPDF_Doc   pdf);



HPDF_EXPORT(HPDF_STATUS)
HPDF_UseKREncodings   (HPDF_Doc   pdf);



HPDF_EXPORT(HPDF_STATUS)
HPDF_UseCNSEncodings   (HPDF_Doc   pdf);



HPDF_EXPORT(HPDF_STATUS)
HPDF_UseCNTEncodings   (HPDF_Doc   pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_UseUTFEncodings   (HPDF_Doc   pdf);

*/

/*--------------------------------------------------------------------------*/
/*----- XObject ------------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_XObject)
HPDF_Page_CreateXObjectFromImage    (HPDF_Doc       pdf,
                                     HPDF_Page      page,
                                     HPDF_Rect      rect,
                                     HPDF_Image     image,
                                     HPDF_Boolean   zoom);

HPDF_EXPORT(HPDF_XObject)
HPDF_Page_CreateXObjectAsWhiteRect  (HPDF_Doc   pdf,
                                     HPDF_Page  page,
                                     HPDF_Rect  rect);
*/

/*--------------------------------------------------------------------------*/
/*----- annotation ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_Create3DAnnot    (HPDF_Page       page,
							HPDF_Rect       rect,
                            HPDF_BOOL       tb,
                            HPDF_BOOL       np,
                            HPDF_U3D        u3d,
                            HPDF_Image      ap);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateTextAnnot  (HPDF_Page       page,
                            HPDF_Rect       rect,
                            const char     *text,
                            HPDF_Encoder    encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateFreeTextAnnot  (HPDF_Page       page,
								HPDF_Rect       rect,
								const char     *text,
								HPDF_Encoder    encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateLineAnnot  (HPDF_Page       page,
							const char     *text,
							HPDF_Encoder    encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateWidgetAnnot_WhiteOnlyWhilePrint (HPDF_Doc   pdf,
                                                 HPDF_Page  page,
                                                 HPDF_Rect  rect);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateWidgetAnnot (HPDF_Page  page,
                             HPDF_Rect  rect);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateLinkAnnot  (HPDF_Page          page,
                            HPDF_Rect          rect,
                            HPDF_Destination   dst);


HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateURILinkAnnot  (HPDF_Page     page,
                               HPDF_Rect     rect,
                               const char   *uri);


HPDF_Annotation
HPDF_Page_CreateTextMarkupAnnot (HPDF_Page     page,
								HPDF_Rect      rect,
								const char     *text,
								HPDF_Encoder   encoder,
								HPDF_AnnotType subType);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateHighlightAnnot  (HPDF_Page   page,
								HPDF_Rect    rect,
								const char   *text,
								HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateUnderlineAnnot (HPDF_Page    page,
								HPDF_Rect    rect,
								const char   *text,
								HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateSquigglyAnnot  (HPDF_Page    page,
								HPDF_Rect    rect,
								const char   *text,
								HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateStrikeOutAnnot  (HPDF_Page   page,
								HPDF_Rect    rect,
								const char   *text,
								HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreatePopupAnnot  (	HPDF_Page          page,
								HPDF_Rect          rect,
								HPDF_Annotation	   parent);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateStampAnnot  (	HPDF_Page           page,
								HPDF_Rect           rect,
								HPDF_StampAnnotName name,
								const char*			text,
								HPDF_Encoder		encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateProjectionAnnot(HPDF_Page page,
								HPDF_Rect rect,
								const char* text,
								HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateSquareAnnot (HPDF_Page          page,
							 HPDF_Rect          rect,
							 const char			*text,
							 HPDF_Encoder       encoder);

HPDF_EXPORT(HPDF_Annotation)
HPDF_Page_CreateCircleAnnot (HPDF_Page          page,
							 HPDF_Rect          rect,
							 const char			*text,
							 HPDF_Encoder       encoder);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LinkAnnot_SetHighlightMode  (HPDF_Annotation           annot,
                                  HPDF_AnnotHighlightMode   mode);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LinkAnnot_SetJavaScript(HPDF_Annotation    annot,
                             HPDF_JavaScript    javascript);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LinkAnnot_SetBorderStyle  (HPDF_Annotation  annot,
                                HPDF_REAL        width,
                                HPDF_UINT16      dash_on,
                                HPDF_UINT16      dash_off);


HPDF_EXPORT(HPDF_STATUS)
HPDF_TextAnnot_SetIcon  (HPDF_Annotation   annot,
                         HPDF_AnnotIcon    icon);


HPDF_EXPORT(HPDF_STATUS)
HPDF_TextAnnot_SetOpened  (HPDF_Annotation   annot,
                          HPDF_BOOL          opened);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annot_SetRGBColor (HPDF_Annotation annot, HPDF_RGBColor color);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annot_SetCMYKColor (HPDF_Annotation annot, HPDF_CMYKColor color);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annot_SetGrayColor (HPDF_Annotation annot, HPDF_REAL color);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annot_SetNoColor (HPDF_Annotation annot);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetTitle (HPDF_Annotation annot, const char* name);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetSubject (HPDF_Annotation annot, const char* name);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetCreationDate (HPDF_Annotation annot, HPDF_Date value);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetTransparency (HPDF_Annotation annot, HPDF_REAL value);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetIntent (HPDF_Annotation  annot, HPDF_AnnotIntent  intent);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetPopup (HPDF_Annotation annot, HPDF_Annotation popup);

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetRectDiff (HPDF_Annotation  annot, HPDF_Rect  rect); // RD entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetCloudEffect (HPDF_Annotation  annot, HPDF_INT cloudIntensity); // BE entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetInteriorRGBColor (HPDF_Annotation  annot, HPDF_RGBColor color); // IC with RGB entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetInteriorCMYKColor (HPDF_Annotation  annot, HPDF_CMYKColor color); // IC with CMYK entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetInteriorGrayColor (HPDF_Annotation  annot, HPDF_REAL color); // IC with Gray entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_MarkupAnnot_SetInteriorTransparent (HPDF_Annotation  annot); // IC with No Color entry

HPDF_EXPORT(HPDF_STATUS)
HPDF_TextMarkupAnnot_SetQuadPoints ( HPDF_Annotation annot, HPDF_Point lb, HPDF_Point rb, HPDF_Point rt, HPDF_Point lt); // l-left, r-right, b-bottom, t-top positions

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annot_Set3DView  ( HPDF_MMgr mmgr,
					 	HPDF_Annotation	annot,
					 	HPDF_Annotation	annot3d,
					 	HPDF_Dict			view);

HPDF_EXPORT(HPDF_STATUS)
HPDF_PopupAnnot_SetOpened  (HPDF_Annotation   annot,
                            HPDF_BOOL         opened);

HPDF_EXPORT(HPDF_STATUS)
HPDF_FreeTextAnnot_SetLineEndingStyle (HPDF_Annotation annot, HPDF_LineAnnotEndingStyle startStyle, HPDF_LineAnnotEndingStyle endStyle);

HPDF_EXPORT(HPDF_STATUS)
HPDF_FreeTextAnnot_Set3PointCalloutLine (HPDF_Annotation annot, HPDF_Point startPoint, HPDF_Point kneePoint, HPDF_Point endPoint); // Callout line will be in default user space

HPDF_EXPORT(HPDF_STATUS)
HPDF_FreeTextAnnot_Set2PointCalloutLine (HPDF_Annotation annot, HPDF_Point startPoint, HPDF_Point endPoint); // Callout line will be in default user space

HPDF_EXPORT(HPDF_STATUS)
HPDF_FreeTextAnnot_SetDefaultStyle (HPDF_Annotation  annot, const char* style);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LineAnnot_SetPosition (HPDF_Annotation annot,
							HPDF_Point startPoint, HPDF_LineAnnotEndingStyle startStyle,
							HPDF_Point endPoint, HPDF_LineAnnotEndingStyle endStyle);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LineAnnot_SetLeader (HPDF_Annotation annot, HPDF_INT leaderLen, HPDF_INT leaderExtLen, HPDF_INT leaderOffsetLen);

HPDF_EXPORT(HPDF_STATUS)
HPDF_LineAnnot_SetCaption (HPDF_Annotation annot, HPDF_BOOL showCaption, HPDF_LineAnnotCapPosition position, HPDF_INT horzOffset, HPDF_INT vertOffset);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Annotation_SetBorderStyle  (HPDF_Annotation  annot,
                                 HPDF_BSSubtype   subtype,
                                 HPDF_REAL        width,
                                 HPDF_UINT16      dash_on,
                                 HPDF_UINT16      dash_off,
                                 HPDF_UINT16      dash_phase);

HPDF_EXPORT(HPDF_STATUS)
HPDF_ProjectionAnnot_SetExData(HPDF_Annotation annot, HPDF_ExData exdata);
*/

/*--------------------------------------------------------------------------*/
/*----- 3D Measure ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_3DMeasure)
HPDF_Page_Create3DC3DMeasure(HPDF_Page       page,
					         HPDF_Point3D    firstanchorpoint,
					         HPDF_Point3D    textanchorpoint
					         );

HPDF_EXPORT(HPDF_3DMeasure)
HPDF_Page_CreatePD33DMeasure(HPDF_Page       page,
					         HPDF_Point3D    annotationPlaneNormal,
					         HPDF_Point3D    firstAnchorPoint,
   					         HPDF_Point3D    secondAnchorPoint,
	        				 HPDF_Point3D    leaderLinesDirection,
   					         HPDF_Point3D    measurementValuePoint,
					         HPDF_Point3D    textYDirection,
					         HPDF_REAL       value,
							 const char*     unitsString
  					        );

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DMeasure_SetName(HPDF_3DMeasure measure,
					   const char* name);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DMeasure_SetColor(HPDF_3DMeasure measure,
						   HPDF_RGBColor color);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DMeasure_SetTextSize(HPDF_3DMeasure measure,
							  HPDF_REAL textsize);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DC3DMeasure_SetTextBoxSize(HPDF_3DMeasure measure,
							 HPDF_INT32 x,
							 HPDF_INT32 y);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DC3DMeasure_SetText(HPDF_3DMeasure measure,
						  const char* text,
						  HPDF_Encoder encoder);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DC3DMeasure_SetProjectionAnotation(HPDF_3DMeasure measure,
										 HPDF_Annotation projectionanotation);
*/

/*--------------------------------------------------------------------------*/
/*----- External Data ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_ExData)
HPDF_Page_Create3DAnnotExData(HPDF_Page       page );

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DAnnotExData_Set3DMeasurement(HPDF_ExData exdata, HPDF_3DMeasure measure);
*/

/*--------------------------------------------------------------------------*/
/*--------------------------------------------------------------------------*/
/*----- 3D View ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_Dict)
HPDF_Page_Create3DView    (HPDF_Page       page,
						   HPDF_U3D        u3d,
						   HPDF_Annotation	annot3d,
						   const char *name);

HPDF_EXPORT(HPDF_STATUS)
HPDF_3DView_Add3DC3DMeasure(HPDF_Dict       view,
							HPDF_3DMeasure measure);
*/

/*--------------------------------------------------------------------------*/
/*----- image data ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_Image)
HPDF_LoadPngImageFromMem  (HPDF_Doc     pdf,
                    const HPDF_BYTE    *buffer,
                          HPDF_UINT     size);

HPDF_EXPORT(HPDF_Image)
HPDF_LoadPngImageFromFile (HPDF_Doc      pdf,
                           const char    *filename);


HPDF_EXPORT(HPDF_Image)
HPDF_LoadPngImageFromFile2 (HPDF_Doc      pdf,
                            const char    *filename);


HPDF_EXPORT(HPDF_Image)
HPDF_LoadJpegImageFromFile (HPDF_Doc      pdf,
                            const char    *filename);

HPDF_EXPORT(HPDF_Image)
HPDF_LoadJpegImageFromMem   (HPDF_Doc      pdf,
                      const HPDF_BYTE     *buffer,
                            HPDF_UINT      size);

HPDF_EXPORT(HPDF_Image)
HPDF_LoadU3DFromFile (HPDF_Doc      pdf,
                            const char    *filename);

HPDF_EXPORT(HPDF_Image)
HPDF_LoadU3DFromMem  (HPDF_Doc      pdf,
               const HPDF_BYTE     *buffer,
                     HPDF_UINT      size);

HPDF_EXPORT(HPDF_Image)
HPDF_Image_LoadRaw1BitImageFromMem  (HPDF_Doc           pdf,
                           const HPDF_BYTE   *buf,
                          HPDF_UINT          width,
                          HPDF_UINT          height,
                          HPDF_UINT          line_width,
                          HPDF_BOOL          black_is1,
                          HPDF_BOOL          top_is_first);


HPDF_EXPORT(HPDF_Image)
HPDF_LoadRawImageFromFile  (HPDF_Doc           pdf,
                            const char         *filename,
                            HPDF_UINT          width,
                            HPDF_UINT          height,
                            HPDF_ColorSpace    color_space);


HPDF_EXPORT(HPDF_Image)
HPDF_LoadRawImageFromMem  (HPDF_Doc           pdf,
                           const HPDF_BYTE   *buf,
                           HPDF_UINT          width,
                           HPDF_UINT          height,
                           HPDF_ColorSpace    color_space,
                           HPDF_UINT          bits_per_component);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Image_AddSMask  (HPDF_Image    image,
                      HPDF_Image    smask);

HPDF_EXPORT(HPDF_Point)
HPDF_Image_GetSize (HPDF_Image  image);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Image_GetSize2 (HPDF_Image  image, HPDF_Point *size);


HPDF_EXPORT(HPDF_UINT)
HPDF_Image_GetWidth  (HPDF_Image   image);


HPDF_EXPORT(HPDF_UINT)
HPDF_Image_GetHeight  (HPDF_Image   image);


HPDF_EXPORT(HPDF_UINT)
HPDF_Image_GetBitsPerComponent (HPDF_Image  image);


HPDF_EXPORT(const char*)
HPDF_Image_GetColorSpace (HPDF_Image  image);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Image_SetColorMask (HPDF_Image   image,
                         HPDF_UINT    rmin,
                         HPDF_UINT    rmax,
                         HPDF_UINT    gmin,
                         HPDF_UINT    gmax,
                         HPDF_UINT    bmin,
                         HPDF_UINT    bmax);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Image_SetMaskImage  (HPDF_Image   image,
                          HPDF_Image   mask_image);
*/

/*--------------------------------------------------------------------------*/
/*----- info dictionary ----------------------------------------------------*/

/*

HPDF_EXPORT(HPDF_STATUS)
HPDF_SetInfoAttr (HPDF_Doc        pdf,
                  HPDF_InfoType   type,
                  const char     *value);


HPDF_EXPORT(const char*)
HPDF_GetInfoAttr (HPDF_Doc        pdf,
                  HPDF_InfoType   type);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetInfoDateAttr (HPDF_Doc        pdf,
                      HPDF_InfoType   type,
                      HPDF_Date       value);
*/

/*--------------------------------------------------------------------------*/
/*----- encryption ---------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_SetPassword  (HPDF_Doc      pdf,
                   const char   *owner_passwd,
                   const char   *user_passwd);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetPermission  (HPDF_Doc    pdf,
                     HPDF_UINT   permission);


HPDF_EXPORT(HPDF_STATUS)
HPDF_SetEncryptionMode  (HPDF_Doc           pdf,
                         HPDF_EncryptMode   mode,
                         HPDF_UINT          key_len);
*/

/*--------------------------------------------------------------------------*/
/*----- compression --------------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_SetCompressionMode  (HPDF_Doc    pdf,
                          HPDF_UINT   mode);
*/

/*--------------------------------------------------------------------------*/
/*----- font ---------------------------------------------------------------*/

/*
HPDF_EXPORT(const char*)
HPDF_Font_GetFontName  (HPDF_Font    font);


HPDF_EXPORT(const char*)
HPDF_Font_GetEncodingName  (HPDF_Font    font);


HPDF_EXPORT(HPDF_INT)
HPDF_Font_GetUnicodeWidth  (HPDF_Font       font,
                            HPDF_UNICODE    code);

HPDF_EXPORT(HPDF_Box)
HPDF_Font_GetBBox  (HPDF_Font    font);


HPDF_EXPORT(HPDF_INT)
HPDF_Font_GetAscent  (HPDF_Font  font);


HPDF_EXPORT(HPDF_INT)
HPDF_Font_GetDescent  (HPDF_Font  font);


HPDF_EXPORT(HPDF_UINT)
HPDF_Font_GetXHeight  (HPDF_Font  font);


HPDF_EXPORT(HPDF_UINT)
HPDF_Font_GetCapHeight  (HPDF_Font  font);


HPDF_EXPORT(HPDF_TextWidth)
HPDF_Font_TextWidth  (HPDF_Font          font,
                      const HPDF_BYTE   *text,
                      HPDF_UINT          len);


HPDF_EXPORT(HPDF_UINT)
HPDF_Font_MeasureText (HPDF_Font          font,
                       const HPDF_BYTE   *text,
                       HPDF_UINT          len,
                       HPDF_REAL          width,
                       HPDF_REAL          font_size,
                       HPDF_REAL          char_space,
                       HPDF_REAL          word_space,
                       HPDF_BOOL          wordwrap,
                       HPDF_REAL         *real_width);
*/

/*--------------------------------------------------------------------------*/
/*----- attachements -------------------------------------------------------*/

/*

HPDF_EXPORT(HPDF_EmbeddedFile)
HPDF_AttachFile  (HPDF_Doc    pdf,
                  const char *file);

*/

/*--------------------------------------------------------------------------*/
/*----- extended graphics state --------------------------------------------*/

/*
HPDF_EXPORT(HPDF_ExtGState)
HPDF_CreateExtGState  (HPDF_Doc  pdf);


HPDF_EXPORT(HPDF_STATUS)
HPDF_ExtGState_SetAlphaStroke  (HPDF_ExtGState   ext_gstate,
                                HPDF_REAL        value);


HPDF_EXPORT(HPDF_STATUS)
HPDF_ExtGState_SetAlphaFill  (HPDF_ExtGState   ext_gstate,
                              HPDF_REAL        value);



HPDF_EXPORT(HPDF_STATUS)
HPDF_ExtGState_SetBlendMode  (HPDF_ExtGState   ext_gstate,
                              HPDF_BlendMode   mode);
*/

/*--------------------------------------------------------------------------*/
/*--------------------------------------------------------------------------*/

func hpdfPageTextWidth(page hpdfPage, text string) float32 {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	t := C.HPDF_Page_TextWidth(C.HPDF_Page(page), cs)
	return float32(t)
}

/*

HPDF_EXPORT(HPDF_UINT)
HPDF_Page_MeasureText  (HPDF_Page    page,
                        const char  *text,
                        HPDF_REAL    width,
                        HPDF_BOOL    wordwrap,
                        HPDF_REAL   *real_width);
*/

func hpdfPageGetWidth(page hpdfPage) float32 {
	t := C.HPDF_Page_GetWidth(C.HPDF_Page(page))
	return float32(t)
}

func hpdfPageGetHeight(page hpdfPage) float32 {
	t := C.HPDF_Page_GetHeight(C.HPDF_Page(page))
	return float32(t)
}

/*
HPDF_EXPORT(HPDF_UINT16)
HPDF_Page_GetGMode  (HPDF_Page   page);


HPDF_EXPORT(HPDF_Point)
HPDF_Page_GetCurrentPos  (HPDF_Page   page);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_GetCurrentPos2  (HPDF_Page    page,
                           HPDF_Point  *pos);


HPDF_EXPORT(HPDF_Point)
HPDF_Page_GetCurrentTextPos (HPDF_Page   page);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_GetCurrentTextPos2  (HPDF_Page    page,
                               HPDF_Point  *pos);


HPDF_EXPORT(HPDF_Font)
HPDF_Page_GetCurrentFont  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetCurrentFontSize  (HPDF_Page   page);


HPDF_EXPORT(HPDF_TransMatrix)
HPDF_Page_GetTransMatrix  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetLineWidth  (HPDF_Page   page);


HPDF_EXPORT(HPDF_LineCap)
HPDF_Page_GetLineCap  (HPDF_Page   page);


HPDF_EXPORT(HPDF_LineJoin)
HPDF_Page_GetLineJoin  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetMiterLimit  (HPDF_Page   page);


HPDF_EXPORT(HPDF_DashMode)
HPDF_Page_GetDash  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetFlat  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetCharSpace  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetWordSpace  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetHorizontalScalling  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetTextLeading  (HPDF_Page   page);


HPDF_EXPORT(HPDF_TextRenderingMode)
HPDF_Page_GetTextRenderingMode  (HPDF_Page   page);


// This function is obsolete. Use HPDF_Page_GetTextRise.
HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetTextRaise  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetTextRise  (HPDF_Page   page);


HPDF_EXPORT(HPDF_RGBColor)
HPDF_Page_GetRGBFill  (HPDF_Page   page);


HPDF_EXPORT(HPDF_RGBColor)
HPDF_Page_GetRGBStroke  (HPDF_Page   page);


HPDF_EXPORT(HPDF_CMYKColor)
HPDF_Page_GetCMYKFill  (HPDF_Page   page);


HPDF_EXPORT(HPDF_CMYKColor)
HPDF_Page_GetCMYKStroke  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetGrayFill  (HPDF_Page   page);


HPDF_EXPORT(HPDF_REAL)
HPDF_Page_GetGrayStroke  (HPDF_Page   page);


HPDF_EXPORT(HPDF_ColorSpace)
HPDF_Page_GetStrokingColorSpace (HPDF_Page   page);


HPDF_EXPORT(HPDF_ColorSpace)
HPDF_Page_GetFillingColorSpace (HPDF_Page   page);


HPDF_EXPORT(HPDF_TransMatrix)
HPDF_Page_GetTextMatrix  (HPDF_Page   page);


HPDF_EXPORT(HPDF_UINT)
HPDF_Page_GetGStateDepth  (HPDF_Page   page);

*/

/*--------------------------------------------------------------------------*/
/*----- GRAPHICS OPERATORS -------------------------------------------------*/

/*--- General graphics state ---------------------------------------------*/

// w
func hpdfPageSetLineWidth(page hpdfPage, lineWidth float32) hpdfStatus {
	t := C.HPDF_Page_SetLineWidth(C.HPDF_Page(page), C.HPDF_REAL(lineWidth))
	return hpdfStatus(t)
}

/*
// J
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetLineCap  (HPDF_Page     page,
                       HPDF_LineCap  line_cap);

// j
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetLineJoin  (HPDF_Page      page,
                        HPDF_LineJoin  line_join);

// M
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetMiterLimit  (HPDF_Page  page,
                          HPDF_REAL  miter_limit);

// d
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetDash  (HPDF_Page           page,
                    const HPDF_UINT16  *dash_ptn,
                    HPDF_UINT           num_param,
                    HPDF_UINT           phase);



// ri --not implemented yet

// i
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetFlat  (HPDF_Page    page,
                    HPDF_REAL    flatness);

// gs

HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetExtGState  (HPDF_Page        page,
                         HPDF_ExtGState   ext_gstate);
*/

/*--- Special graphic state operator --------------------------------------*/

/*
// q
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_GSave  (HPDF_Page    page);

// Q
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_GRestore  (HPDF_Page    page);

// cm
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Concat  (HPDF_Page    page,
                   HPDF_REAL    a,
                   HPDF_REAL    b,
                   HPDF_REAL    c,
                   HPDF_REAL    d,
                   HPDF_REAL    x,
                   HPDF_REAL    y);
*/

/*--- Path construction operator ------------------------------------------*/

/*
// m
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_MoveTo  (HPDF_Page    page,
                   HPDF_REAL    x,
                   HPDF_REAL    y);

// l
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_LineTo  (HPDF_Page    page,
                   HPDF_REAL    x,
                   HPDF_REAL    y);

// c
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_CurveTo  (HPDF_Page    page,
                    HPDF_REAL    x1,
                    HPDF_REAL    y1,
                    HPDF_REAL    x2,
                    HPDF_REAL    y2,
                    HPDF_REAL    x3,
                    HPDF_REAL    y3);

// v
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_CurveTo2  (HPDF_Page    page,
                     HPDF_REAL    x2,
                     HPDF_REAL    y2,
                     HPDF_REAL    x3,
                     HPDF_REAL    y3);

// y
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_CurveTo3  (HPDF_Page  page,
                     HPDF_REAL  x1,
                     HPDF_REAL  y1,
                     HPDF_REAL  x3,
                     HPDF_REAL  y3);

// h
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ClosePath  (HPDF_Page  page);
*/

func hpdfPageRectangle(page hpdfPage, x, y, dx, dy float32) hpdfStatus {
	t := C.HPDF_Page_Rectangle(
		C.HPDF_Page(page),
		C.HPDF_REAL(x),
		C.HPDF_REAL(y),
		C.HPDF_REAL(dx),
		C.HPDF_REAL(dy),
	)
	return hpdfStatus(t)
}

/*--- Path painting operator ---------------------------------------------*/

// S
func hpdfPageStroke(page hpdfPage) hpdfStatus {
	t := C.HPDF_Page_Stroke(C.HPDF_Page(page))
	return hpdfStatus(t)
}

/*
// s
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ClosePathStroke  (HPDF_Page  page);

// f
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Fill  (HPDF_Page  page);

// f*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Eofill  (HPDF_Page  page);

// B
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_FillStroke  (HPDF_Page  page);

// B*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_EofillStroke  (HPDF_Page  page);

// b
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ClosePathFillStroke  (HPDF_Page  page);

// b*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ClosePathEofillStroke  (HPDF_Page  page);

// n
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_EndPath  (HPDF_Page  page);
*/

/*--- Clipping paths operator --------------------------------------------*/

/*
// W
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Clip  (HPDF_Page  page);

// W*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Eoclip  (HPDF_Page  page);
*/

/*--- Text object operator -----------------------------------------------*/

// BT
func hpdfPageBeginText(page hpdfPage) hpdfStatus {
	t := C.HPDF_Page_BeginText(C.HPDF_Page(page))
	return hpdfStatus(t)
}

// ET
func hpdfPageEndText(page hpdfPage) hpdfStatus {
	t := C.HPDF_Page_EndText(C.HPDF_Page(page))
	return hpdfStatus(t)
}

/*--- Text state ---------------------------------------------------------*/

/*
// Tc
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetCharSpace  (HPDF_Page  page,
                         HPDF_REAL  value);

// Tw
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetWordSpace  (HPDF_Page  page,
                         HPDF_REAL  value);

// Tz
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetHorizontalScalling  (HPDF_Page  page,
                                  HPDF_REAL  value);

// TL
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetTextLeading  (HPDF_Page  page,
                           HPDF_REAL  value);
*/

// Tf
func hpdfPageSetFontAndSize(page hpdfPage, font hpdfFont, size float32) hpdfStatus {
	t := C.HPDF_Page_SetFontAndSize(
		C.HPDF_Page(page),
		C.HPDF_Font(font),
		C.HPDF_REAL(size),
	)
	return hpdfStatus(t)
}

/*
// Tr
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetTextRenderingMode  (HPDF_Page               page,
                                 HPDF_TextRenderingMode  mode);

// Ts
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetTextRise  (HPDF_Page   page,
                        HPDF_REAL   value);

// This function is obsolete. Use HPDF_Page_SetTextRise.
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetTextRaise  (HPDF_Page   page,
                         HPDF_REAL   value);
*/

/*--- Text positioning ---------------------------------------------------*/

// Td
func hpdfPageMoveTextPos(page hpdfPage, x, y float32) hpdfStatus {
	t := C.HPDF_Page_MoveTextPos(
		C.HPDF_Page(page),
		C.HPDF_REAL(x),
		C.HPDF_REAL(y),
	)
	return hpdfStatus(t)
}

/*
// TD
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_MoveTextPos2  (HPDF_Page  page,
                         HPDF_REAL   x,
                         HPDF_REAL   y);

// Tm
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetTextMatrix  (HPDF_Page         page,
                          HPDF_REAL    a,
                          HPDF_REAL    b,
                          HPDF_REAL    c,
                          HPDF_REAL    d,
                          HPDF_REAL    x,
                          HPDF_REAL    y);


// T*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_MoveToNextLine  (HPDF_Page  page);
*/

/*--- Text showing -------------------------------------------------------*/

// Tj
func hpdfPageShowText(page hpdfPage, text string) hpdfStatus {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	t := C.HPDF_Page_ShowText(
		C.HPDF_Page(page),
		cs,
	)
	return hpdfStatus(t)
}

/*
// TJ

// '
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ShowTextNextLine  (HPDF_Page    page,
                             const char  *text);

// "
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ShowTextNextLineEx  (HPDF_Page    page,
                               HPDF_REAL    word_space,
                               HPDF_REAL    char_space,
                               const char  *text);
*/

/*--- Color showing ------------------------------------------------------*/

/*
// cs --not implemented yet
// CS --not implemented yet
// sc --not implemented yet
// scn --not implemented yet
// SC --not implemented yet
// SCN --not implemented yet

// g
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetGrayFill  (HPDF_Page   page,
                        HPDF_REAL   gray);

// G
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetGrayStroke  (HPDF_Page   page,
                          HPDF_REAL   gray);

// rg
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetRGBFill  (HPDF_Page  page,
                       HPDF_REAL  r,
                       HPDF_REAL  g,
                       HPDF_REAL  b);

// RG
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetRGBStroke  (HPDF_Page  page,
                         HPDF_REAL  r,
                         HPDF_REAL  g,
                         HPDF_REAL  b);

// k
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetCMYKFill  (HPDF_Page  page,
                        HPDF_REAL  c,
                        HPDF_REAL  m,
                        HPDF_REAL  y,
                        HPDF_REAL  k);

// K
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetCMYKStroke  (HPDF_Page  page,
                          HPDF_REAL  c,
                          HPDF_REAL  m,
                          HPDF_REAL  y,
                          HPDF_REAL  k);
*/

/*--- Shading patterns ---------------------------------------------------*/

// sh --not implemented yet

/*--- In-line images -----------------------------------------------------*/

// BI --not implemented yet
// ID --not implemented yet
// EI --not implemented yet

/*--- XObjects -----------------------------------------------------------*/

/*
// Do
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_ExecuteXObject  (HPDF_Page     page,
                           HPDF_XObject  obj);
*/

/*--- Content streams ----------------------------------------------------*/

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_New_Content_Stream  (HPDF_Page page,
                               HPDF_Dict* new_stream);

HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Insert_Shared_Content_Stream  (HPDF_Page page,
                                         HPDF_Dict shared_stream);
*/

/*--- Marked content -----------------------------------------------------*/

// BMC --not implemented yet
// BDC --not implemented yet
// EMC --not implemented yet
// MP --not implemented yet
// DP --not implemented yet

/*--- Compatibility ------------------------------------------------------*/

// BX --not implemented yet
// EX --not implemented yet

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_DrawImage  (HPDF_Page    page,
                      HPDF_Image   image,
                      HPDF_REAL    x,
                      HPDF_REAL    y,
                      HPDF_REAL    width,
                      HPDF_REAL    height);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Circle  (HPDF_Page     page,
                   HPDF_REAL     x,
                   HPDF_REAL     y,
                   HPDF_REAL     ray);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Ellipse  (HPDF_Page   page,
                    HPDF_REAL   x,
                    HPDF_REAL   y,
                    HPDF_REAL  xray,
                    HPDF_REAL  yray);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_Arc  (HPDF_Page    page,
                HPDF_REAL    x,
                HPDF_REAL    y,
                HPDF_REAL    ray,
                HPDF_REAL    ang1,
                HPDF_REAL    ang2);
*/

func hpdfPageTextOut(page hpdfPage, xpos, ypos float32, text string) hpdfStatus {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	t := C.HPDF_Page_TextOut(
		C.HPDF_Page(page),
		C.HPDF_REAL(xpos),
		C.HPDF_REAL(ypos),
		cs,
	)
	return hpdfStatus(t)
}

/*
HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_TextRect  (HPDF_Page            page,
                     HPDF_REAL            left,
                     HPDF_REAL            top,
                     HPDF_REAL            right,
                     HPDF_REAL            bottom,
                     const char          *text,
                     HPDF_TextAlignment   align,
                     HPDF_UINT           *len);


HPDF_EXPORT(HPDF_STATUS)
HPDF_Page_SetSlideShow  (HPDF_Page              page,
                         HPDF_TransitionStyle   type,
                         HPDF_REAL              disp_time,
                         HPDF_REAL              trans_time);


HPDF_EXPORT(HPDF_OutputIntent)
HPDF_ICC_LoadIccFromMem (HPDF_Doc   pdf,
                        HPDF_MMgr   mmgr,
                        HPDF_Stream iccdata,
                        HPDF_Xref   xref,
                        int         numcomponent);

HPDF_EXPORT(HPDF_OutputIntent)
HPDF_LoadIccProfileFromFile  (HPDF_Doc  pdf,
                            const char* icc_file_name,
                                   int  numcomponent);
*/
