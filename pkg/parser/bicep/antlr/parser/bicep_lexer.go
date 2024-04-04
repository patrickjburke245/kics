// Code generated from bicep.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type bicepLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BicepLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func biceplexerLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'@'", "','", "'['", "']'", "'('", "')'", "'.'", "'|'", "':'",
		"'='", "'{'", "'}'", "'param'", "'var'", "'true'", "'false'", "'null'",
		"'object'", "'resource'", "", "", "", "", "'string'", "'int'", "'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "OBJECT", "RESOURCE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IDENTIFIER", "NUMBER", "NL", "SPACES", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "OBJECT", "RESOURCE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IDENTIFIER", "NUMBER", "NL", "SPACES", "UNKNOWN", "STRINGCHAR",
		"ESCAPE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 259, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 5, 0, 77, 8, 0, 10, 0, 12, 0, 80, 9, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 5, 20, 154, 8, 20, 10, 20, 12, 20, 157, 9, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 164, 8, 21, 10, 21, 12, 21, 167,
		9, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 174, 8, 22, 10, 22, 12,
		22, 177, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 5, 23, 183, 8, 23, 10, 23,
		12, 23, 186, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 5, 27, 208, 8, 27, 10, 27, 12, 27, 211, 9, 27, 1, 28, 4,
		28, 214, 8, 28, 11, 28, 12, 28, 215, 1, 28, 1, 28, 4, 28, 220, 8, 28, 11,
		28, 12, 28, 221, 3, 28, 224, 8, 28, 1, 29, 4, 29, 227, 8, 29, 11, 29, 12,
		29, 228, 1, 30, 4, 30, 232, 8, 30, 11, 30, 12, 30, 233, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 32, 1, 32, 3, 32, 242, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 4, 33, 250, 8, 33, 11, 33, 12, 33, 251, 1, 33, 1, 33, 3,
		33, 256, 8, 33, 1, 34, 1, 34, 1, 78, 0, 35, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 0, 67, 0, 69, 0, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2, 0,
		9, 9, 32, 32, 5, 0, 9, 10, 13, 13, 36, 36, 39, 39, 92, 92, 6, 0, 36, 36,
		39, 39, 92, 92, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97,
		102, 269, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0, 3, 85, 1, 0, 0, 0, 5,
		87, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93, 1, 0, 0,
		0, 13, 95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0, 19, 101,
		1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 107, 1, 0, 0,
		0, 27, 109, 1, 0, 0, 0, 29, 115, 1, 0, 0, 0, 31, 119, 1, 0, 0, 0, 33, 124,
		1, 0, 0, 0, 35, 130, 1, 0, 0, 0, 37, 135, 1, 0, 0, 0, 39, 142, 1, 0, 0,
		0, 41, 151, 1, 0, 0, 0, 43, 161, 1, 0, 0, 0, 45, 171, 1, 0, 0, 0, 47, 180,
		1, 0, 0, 0, 49, 189, 1, 0, 0, 0, 51, 196, 1, 0, 0, 0, 53, 200, 1, 0, 0,
		0, 55, 205, 1, 0, 0, 0, 57, 213, 1, 0, 0, 0, 59, 226, 1, 0, 0, 0, 61, 231,
		1, 0, 0, 0, 63, 237, 1, 0, 0, 0, 65, 241, 1, 0, 0, 0, 67, 243, 1, 0, 0,
		0, 69, 257, 1, 0, 0, 0, 71, 72, 5, 39, 0, 0, 72, 73, 5, 39, 0, 0, 73, 74,
		5, 39, 0, 0, 74, 78, 1, 0, 0, 0, 75, 77, 9, 0, 0, 0, 76, 75, 1, 0, 0, 0,
		77, 80, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 81, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 39, 0, 0, 82, 83, 5, 39, 0, 0,
		83, 84, 5, 39, 0, 0, 84, 2, 1, 0, 0, 0, 85, 86, 5, 64, 0, 0, 86, 4, 1,
		0, 0, 0, 87, 88, 5, 44, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 91, 0, 0, 90,
		8, 1, 0, 0, 0, 91, 92, 5, 93, 0, 0, 92, 10, 1, 0, 0, 0, 93, 94, 5, 40,
		0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 41, 0, 0, 96, 14, 1, 0, 0, 0, 97,
		98, 5, 46, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 124, 0, 0, 100, 18, 1,
		0, 0, 0, 101, 102, 5, 58, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104, 5, 61, 0,
		0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 123, 0, 0, 106, 24, 1, 0, 0, 0, 107,
		108, 5, 125, 0, 0, 108, 26, 1, 0, 0, 0, 109, 110, 5, 112, 0, 0, 110, 111,
		5, 97, 0, 0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5,
		109, 0, 0, 114, 28, 1, 0, 0, 0, 115, 116, 5, 118, 0, 0, 116, 117, 5, 97,
		0, 0, 117, 118, 5, 114, 0, 0, 118, 30, 1, 0, 0, 0, 119, 120, 5, 116, 0,
		0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 117, 0, 0, 122, 123, 5, 101, 0,
		0, 123, 32, 1, 0, 0, 0, 124, 125, 5, 102, 0, 0, 125, 126, 5, 97, 0, 0,
		126, 127, 5, 108, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 101, 0, 0,
		129, 34, 1, 0, 0, 0, 130, 131, 5, 110, 0, 0, 131, 132, 5, 117, 0, 0, 132,
		133, 5, 108, 0, 0, 133, 134, 5, 108, 0, 0, 134, 36, 1, 0, 0, 0, 135, 136,
		5, 111, 0, 0, 136, 137, 5, 98, 0, 0, 137, 138, 5, 106, 0, 0, 138, 139,
		5, 101, 0, 0, 139, 140, 5, 99, 0, 0, 140, 141, 5, 116, 0, 0, 141, 38, 1,
		0, 0, 0, 142, 143, 5, 114, 0, 0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 115,
		0, 0, 145, 146, 5, 111, 0, 0, 146, 147, 5, 117, 0, 0, 147, 148, 5, 114,
		0, 0, 148, 149, 5, 99, 0, 0, 149, 150, 5, 101, 0, 0, 150, 40, 1, 0, 0,
		0, 151, 155, 5, 39, 0, 0, 152, 154, 3, 65, 32, 0, 153, 152, 1, 0, 0, 0,
		154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156,
		158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 36, 0, 0, 159, 160,
		5, 123, 0, 0, 160, 42, 1, 0, 0, 0, 161, 165, 5, 125, 0, 0, 162, 164, 3,
		65, 32, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0,
		0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0,
		168, 169, 5, 36, 0, 0, 169, 170, 5, 123, 0, 0, 170, 44, 1, 0, 0, 0, 171,
		175, 5, 125, 0, 0, 172, 174, 3, 65, 32, 0, 173, 172, 1, 0, 0, 0, 174, 177,
		1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0,
		0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 5, 39, 0, 0, 179, 46, 1, 0, 0, 0,
		180, 184, 5, 39, 0, 0, 181, 183, 3, 65, 32, 0, 182, 181, 1, 0, 0, 0, 183,
		186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187,
		1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 188, 5, 39, 0, 0, 188, 48, 1, 0,
		0, 0, 189, 190, 5, 115, 0, 0, 190, 191, 5, 116, 0, 0, 191, 192, 5, 114,
		0, 0, 192, 193, 5, 105, 0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 103,
		0, 0, 195, 50, 1, 0, 0, 0, 196, 197, 5, 105, 0, 0, 197, 198, 5, 110, 0,
		0, 198, 199, 5, 116, 0, 0, 199, 52, 1, 0, 0, 0, 200, 201, 5, 98, 0, 0,
		201, 202, 5, 111, 0, 0, 202, 203, 5, 111, 0, 0, 203, 204, 5, 108, 0, 0,
		204, 54, 1, 0, 0, 0, 205, 209, 7, 0, 0, 0, 206, 208, 7, 1, 0, 0, 207, 206,
		1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0,
		0, 0, 210, 56, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 214, 7, 2, 0, 0,
		213, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215,
		216, 1, 0, 0, 0, 216, 223, 1, 0, 0, 0, 217, 219, 5, 46, 0, 0, 218, 220,
		7, 2, 0, 0, 219, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 217, 1, 0, 0, 0,
		223, 224, 1, 0, 0, 0, 224, 58, 1, 0, 0, 0, 225, 227, 7, 3, 0, 0, 226, 225,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0,
		0, 0, 229, 60, 1, 0, 0, 0, 230, 232, 7, 4, 0, 0, 231, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 236, 6, 30, 0, 0, 236, 62, 1, 0, 0, 0, 237, 238,
		9, 0, 0, 0, 238, 64, 1, 0, 0, 0, 239, 242, 8, 5, 0, 0, 240, 242, 3, 67,
		33, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 66, 1, 0, 0, 0,
		243, 255, 5, 92, 0, 0, 244, 256, 7, 6, 0, 0, 245, 246, 5, 117, 0, 0, 246,
		247, 5, 123, 0, 0, 247, 249, 1, 0, 0, 0, 248, 250, 3, 69, 34, 0, 249, 248,
		1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0,
		0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 5, 125, 0, 0, 254, 256, 1, 0, 0,
		0, 255, 244, 1, 0, 0, 0, 255, 245, 1, 0, 0, 0, 256, 68, 1, 0, 0, 0, 257,
		258, 7, 7, 0, 0, 258, 70, 1, 0, 0, 0, 15, 0, 78, 155, 165, 175, 184, 209,
		215, 221, 223, 228, 233, 241, 251, 255, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// bicepLexerInit initializes any static state used to implement bicepLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewbicepLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BicepLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.once.Do(biceplexerLexerInit)
}

// NewbicepLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewbicepLexer(input antlr.CharStream) *bicepLexer {
	BicepLexerInit()
	l := new(bicepLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BicepLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "bicep.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// bicepLexer tokens.
const (
	bicepLexerMULTILINE_STRING    = 1
	bicepLexerAT                  = 2
	bicepLexerCOMMA               = 3
	bicepLexerOBRACK              = 4
	bicepLexerCBRACK              = 5
	bicepLexerOPAR                = 6
	bicepLexerCPAR                = 7
	bicepLexerDOT                 = 8
	bicepLexerPIPE                = 9
	bicepLexerCOL                 = 10
	bicepLexerASSIGN              = 11
	bicepLexerOBRACE              = 12
	bicepLexerCBRACE              = 13
	bicepLexerPARAM               = 14
	bicepLexerVAR                 = 15
	bicepLexerTRUE                = 16
	bicepLexerFALSE               = 17
	bicepLexerNULL                = 18
	bicepLexerOBJECT              = 19
	bicepLexerRESOURCE            = 20
	bicepLexerSTRING_LEFT_PIECE   = 21
	bicepLexerSTRING_MIDDLE_PIECE = 22
	bicepLexerSTRING_RIGHT_PIECE  = 23
	bicepLexerSTRING_COMPLETE     = 24
	bicepLexerSTRING              = 25
	bicepLexerINT                 = 26
	bicepLexerBOOL                = 27
	bicepLexerIDENTIFIER          = 28
	bicepLexerNUMBER              = 29
	bicepLexerNL                  = 30
	bicepLexerSPACES              = 31
	bicepLexerUNKNOWN             = 32
)
