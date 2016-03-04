// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 05 Mar 2016 02:56:32 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package vorbis

/*
#cgo pkg-config: ogg vorbis
#include "ogg/ogg.h"
#include "vorbis/codec.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// OggStreamPacketin function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetin.html
func OggStreamPacketin(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetin(cos, cop)
	__v := (int32)(__ret)
	return __v
}

// OggStreamIovecin function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_iovecin.html
func OggStreamIovecin(os *OggStreamState, iov *OggIovec, count int32, eOS int, granulepos OggInt64) int32 {
	cos, _ := os.PassRef()
	ciov, _ := iov.PassRef()
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	ceOS, _ := (C.long)(eOS), cgoAllocsUnknown
	cgranulepos, _ := (C.ogg_int64_t)(granulepos), cgoAllocsUnknown
	__ret := C.ogg_stream_iovecin(cos, ciov, ccount, ceOS, cgranulepos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPageout function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pageout.html
func OggStreamPageout(os *OggStreamState, og *OggPage) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_stream_pageout(cos, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPageoutFill function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pageout_fill.html
func OggStreamPageoutFill(os *OggStreamState, og *OggPage, nfill int32) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	cnfill, _ := (C.int)(nfill), cgoAllocsUnknown
	__ret := C.ogg_stream_pageout_fill(cos, cog, cnfill)
	__v := (int32)(__ret)
	return __v
}

// OggStreamFlush function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_flush.html
func OggStreamFlush(os *OggStreamState, og *OggPage) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_stream_flush(cos, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamFlushFill function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_flush_fill.html
func OggStreamFlushFill(os *OggStreamState, og *OggPage, nfill int32) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	cnfill, _ := (C.int)(nfill), cgoAllocsUnknown
	__ret := C.ogg_stream_flush_fill(cos, cog, cnfill)
	__v := (int32)(__ret)
	return __v
}

// OggSyncInit function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_init.html
func OggSyncInit(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_init(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_clear.html
func OggSyncClear(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_clear(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncReset function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_reset.html
func OggSyncReset(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_reset(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncDestroy function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_destroy.html
func OggSyncDestroy(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_destroy(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncCheck function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_check.html
func OggSyncCheck(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_check(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncBuffer function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_buffer.html
func OggSyncBuffer(oy *OggSyncState, size int) []byte {
	coy, _ := oy.PassRef()
	csize, _ := (C.long)(size), cgoAllocsUnknown
	__ret := C.ogg_sync_buffer(coy, csize)
	__v := (*(*[0x7fffffff]byte)(unsafe.Pointer(__ret)))[:0]
	return __v
}

// OggSyncWrote function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_wrote.html
func OggSyncWrote(oy *OggSyncState, bytes int) int32 {
	coy, _ := oy.PassRef()
	cbytes, _ := (C.long)(bytes), cgoAllocsUnknown
	__ret := C.ogg_sync_wrote(coy, cbytes)
	__v := (int32)(__ret)
	return __v
}

// OggSyncPageseek function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_pageseek.html
func OggSyncPageseek(oy *OggSyncState, og *OggPage) int {
	coy, _ := oy.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_sync_pageseek(coy, cog)
	__v := (int)(__ret)
	return __v
}

// OggSyncPageout function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_pageout.html
func OggSyncPageout(oy *OggSyncState, og *OggPage) int32 {
	coy, _ := oy.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_sync_pageout(coy, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPagein function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pagein.html
func OggStreamPagein(os *OggStreamState, og *OggPage) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_stream_pagein(cos, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPacketout function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetout.html
func OggStreamPacketout(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetout(cos, cop)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPacketpeek function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetpeek.html
func OggStreamPacketpeek(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetpeek(cos, cop)
	__v := (int32)(__ret)
	return __v
}

// OggStreamInit function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_init.html
func OggStreamInit(os *OggStreamState, serialno int32) int32 {
	cos, _ := os.PassRef()
	cserialno, _ := (C.int)(serialno), cgoAllocsUnknown
	__ret := C.ogg_stream_init(cos, cserialno)
	__v := (int32)(__ret)
	return __v
}

// OggStreamClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_clear.html
func OggStreamClear(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_clear(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamReset function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_reset.html
func OggStreamReset(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_reset(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamResetSerialno function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_reset_serialno.html
func OggStreamResetSerialno(os *OggStreamState, serialno int32) int32 {
	cos, _ := os.PassRef()
	cserialno, _ := (C.int)(serialno), cgoAllocsUnknown
	__ret := C.ogg_stream_reset_serialno(cos, cserialno)
	__v := (int32)(__ret)
	return __v
}

// OggStreamDestroy function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_destroy.html
func OggStreamDestroy(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_destroy(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamCheck function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_check.html
func OggStreamCheck(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_check(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamEos function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_eos.html
func OggStreamEos(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_eos(cos)
	__v := (int32)(__ret)
	return __v
}

// OggPageChecksumSet function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_checksum_set.html
func OggPageChecksumSet(og *OggPage) {
	cog, _ := og.PassRef()
	C.ogg_page_checksum_set(cog)
}

// OggPageVersion function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_version.html
func OggPageVersion(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_version(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageContinued function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_continued.html
func OggPageContinued(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_continued(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageBos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_bos.html
func OggPageBos(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_bos(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageEos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_eos.html
func OggPageEos(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_eos(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageGranulepos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_granulepos.html
func OggPageGranulepos(og *OggPage) OggInt64 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_granulepos(cog)
	__v := (OggInt64)(__ret)
	return __v
}

// OggPageSerialno function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_serialno.html
func OggPageSerialno(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_serialno(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPagePageno function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_pageno.html
func OggPagePageno(og *OggPage) int {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_pageno(cog)
	__v := (int)(__ret)
	return __v
}

// OggPagePackets function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_packets.html
func OggPagePackets(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_packets(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPacketClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_packet_clear.html
func OggPacketClear(op *OggPacket) {
	cop, _ := op.PassRef()
	C.ogg_packet_clear(cop)
}

// InfoInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_init.html
func InfoInit(vi *Info) {
	cvi, _ := vi.PassRef()
	C.vorbis_info_init(cvi)
}

// InfoClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_clear.html
func InfoClear(vi *Info) {
	cvi, _ := vi.PassRef()
	C.vorbis_info_clear(cvi)
}

// InfoBlocksize function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_blocksize.html
func InfoBlocksize(vi *Info, zo int32) int32 {
	cvi, _ := vi.PassRef()
	czo, _ := (C.int)(zo), cgoAllocsUnknown
	__ret := C.vorbis_info_blocksize(cvi, czo)
	__v := (int32)(__ret)
	return __v
}

// CommentInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_init.html
func CommentInit(vc *Comment) {
	cvc, _ := vc.PassRef()
	C.vorbis_comment_init(cvc)
}

// CommentAdd function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_add.html
func CommentAdd(vc *Comment, comment string) {
	cvc, _ := vc.PassRef()
	ccomment, _ := unpackPCharString(comment)
	C.vorbis_comment_add(cvc, ccomment)
}

// CommentAddTag function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_add_tag.html
func CommentAddTag(vc *Comment, tag string, contents string) {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	ccontents, _ := unpackPCharString(contents)
	C.vorbis_comment_add_tag(cvc, ctag, ccontents)
}

// CommentQuery function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_query.html
func CommentQuery(vc *Comment, tag string, count int32) *byte {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	__ret := C.vorbis_comment_query(cvc, ctag, ccount)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// CommentQueryCount function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_query_count.html
func CommentQueryCount(vc *Comment, tag string) int32 {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	__ret := C.vorbis_comment_query_count(cvc, ctag)
	__v := (int32)(__ret)
	return __v
}

// CommentClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_clear.html
func CommentClear(vc *Comment) {
	cvc, _ := vc.PassRef()
	C.vorbis_comment_clear(cvc)
}

// BlockInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_block_init.html
func BlockInit(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_block_init(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// BlockClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_block_clear.html
func BlockClear(vb *Block) int32 {
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_block_clear(cvb)
	__v := (int32)(__ret)
	return __v
}

// DspClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_dsp_clear.html
func DspClear(v *DspState) {
	cv, _ := v.PassRef()
	C.vorbis_dsp_clear(cv)
}

// GranuleTime function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_granule_time.html
func GranuleTime(v *DspState, granulepos OggInt64) float64 {
	cv, _ := v.PassRef()
	cgranulepos, _ := (C.ogg_int64_t)(granulepos), cgoAllocsUnknown
	__ret := C.vorbis_granule_time(cv, cgranulepos)
	__v := (float64)(__ret)
	return __v
}

// VersionString function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_version_string.html
func VersionString() string {
	__ret := C.vorbis_version_string()
	__v := packPCharString(__ret)
	return __v
}

// AnalysisInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_init.html
func AnalysisInit(v *DspState, vi *Info) int32 {
	cv, _ := v.PassRef()
	cvi, _ := vi.PassRef()
	__ret := C.vorbis_analysis_init(cv, cvi)
	__v := (int32)(__ret)
	return __v
}

// CommentheaderOut function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_commentheader_out.html
func CommentheaderOut(vc *Comment, op *OggPacket) int32 {
	cvc, _ := vc.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_commentheader_out(cvc, cop)
	__v := (int32)(__ret)
	return __v
}

// AnalysisHeaderout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_headerout.html
func AnalysisHeaderout(v *DspState, vc *Comment, op *OggPacket, opComm []OggPacket, opCode []OggPacket) int32 {
	cv, _ := v.PassRef()
	cvc, _ := vc.PassRef()
	cop, _ := op.PassRef()
	copComm, _ := unpackArgSOggPacket(opComm)
	copCode, _ := unpackArgSOggPacket(opCode)
	__ret := C.vorbis_analysis_headerout(cv, cvc, cop, copComm, copCode)
	packSOggPacket(opCode, copCode)
	packSOggPacket(opComm, copComm)
	__v := (int32)(__ret)
	return __v
}

// AnalysisBuffer function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_buffer.html
func AnalysisBuffer(v *DspState, vals int32) **float32 {
	cv, _ := v.PassRef()
	cvals, _ := (C.int)(vals), cgoAllocsUnknown
	__ret := C.vorbis_analysis_buffer(cv, cvals)
	__v := *(***float32)(unsafe.Pointer(&__ret))
	return __v
}

// AnalysisWrote function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_wrote.html
func AnalysisWrote(v *DspState, vals int32) int32 {
	cv, _ := v.PassRef()
	cvals, _ := (C.int)(vals), cgoAllocsUnknown
	__ret := C.vorbis_analysis_wrote(cv, cvals)
	__v := (int32)(__ret)
	return __v
}

// AnalysisBlockout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_blockout.html
func AnalysisBlockout(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_analysis_blockout(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// Analysis function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis.html
func Analysis(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_analysis(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// BitrateAddblock function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_bitrate_addblock.html
func BitrateAddblock(vb *Block) int32 {
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_bitrate_addblock(cvb)
	__v := (int32)(__ret)
	return __v
}

// BitrateFlushpacket function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_bitrate_flushpacket.html
func BitrateFlushpacket(vd *DspState, op *OggPacket) int32 {
	cvd, _ := vd.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_bitrate_flushpacket(cvd, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisIdheader function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_idheader.html
func SynthesisIdheader(op *OggPacket) int32 {
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_idheader(cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisHeaderin function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_headerin.html
func SynthesisHeaderin(vi *Info, vc *Comment, op *OggPacket) int32 {
	cvi, _ := vi.PassRef()
	cvc, _ := vc.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_headerin(cvi, cvc, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_init.html
func SynthesisInit(v *DspState, vi *Info) int32 {
	cv, _ := v.PassRef()
	cvi, _ := vi.PassRef()
	__ret := C.vorbis_synthesis_init(cv, cvi)
	__v := (int32)(__ret)
	return __v
}

// SynthesisRestart function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_restart.html
func SynthesisRestart(v *DspState) int32 {
	cv, _ := v.PassRef()
	__ret := C.vorbis_synthesis_restart(cv)
	__v := (int32)(__ret)
	return __v
}

// Synthesis function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis.html
func Synthesis(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisTrackonly function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_trackonly.html
func SynthesisTrackonly(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_trackonly(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisBlockin function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_blockin.html
func SynthesisBlockin(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_synthesis_blockin(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// SynthesisPcmout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_pcmout.html
func SynthesisPcmout(v *DspState, pcm [][][]float32) int32 {
	cv, _ := v.PassRef()
	cpcm, _ := unpackArgSSSFloat(pcm)
	__ret := C.vorbis_synthesis_pcmout(cv, cpcm)
	packSSSFloat(pcm, cpcm)
	__v := (int32)(__ret)
	return __v
}

// SynthesisLapout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_lapout.html
func SynthesisLapout(v *DspState, pcm ***float32) int32 {
	cv, _ := v.PassRef()
	cpcm, _ := (***C.float)(unsafe.Pointer(pcm)), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_lapout(cv, cpcm)
	__v := (int32)(__ret)
	return __v
}

// SynthesisRead function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_read.html
func SynthesisRead(v *DspState, samples int32) int32 {
	cv, _ := v.PassRef()
	csamples, _ := (C.int)(samples), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_read(cv, csamples)
	__v := (int32)(__ret)
	return __v
}

// PacketBlocksize function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_packet_blocksize.html
func PacketBlocksize(vi *Info, op *OggPacket) int {
	cvi, _ := vi.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_packet_blocksize(cvi, cop)
	__v := (int)(__ret)
	return __v
}

// SynthesisHalfrate function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_halfrate.html
func SynthesisHalfrate(v *Info, flag int32) int32 {
	cv, _ := v.PassRef()
	cflag, _ := (C.int)(flag), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_halfrate(cv, cflag)
	__v := (int32)(__ret)
	return __v
}

// SynthesisHalfrateP function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_halfrate_p.html
func SynthesisHalfrateP(v *Info) int32 {
	cv, _ := v.PassRef()
	__ret := C.vorbis_synthesis_halfrate_p(cv)
	__v := (int32)(__ret)
	return __v
}
