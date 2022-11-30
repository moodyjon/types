// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: metadata.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata_Version int32

const (
	Metadata_UNKNOWN_VERSION Metadata_Version = 0
	Metadata__0_0_1          Metadata_Version = 1
	Metadata__0_0_2          Metadata_Version = 2
	Metadata__0_0_3          Metadata_Version = 3
	Metadata__0_1_0          Metadata_Version = 4
)

// Enum value maps for Metadata_Version.
var (
	Metadata_Version_name = map[int32]string{
		0: "UNKNOWN_VERSION",
		1: "_0_0_1",
		2: "_0_0_2",
		3: "_0_0_3",
		4: "_0_1_0",
	}
	Metadata_Version_value = map[string]int32{
		"UNKNOWN_VERSION": 0,
		"_0_0_1":          1,
		"_0_0_2":          2,
		"_0_0_3":          3,
		"_0_1_0":          4,
	}
)

func (x Metadata_Version) Enum() *Metadata_Version {
	p := new(Metadata_Version)
	*p = x
	return p
}

func (x Metadata_Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_Version) Descriptor() protoreflect.EnumDescriptor {
	return file_metadata_proto_enumTypes[0].Descriptor()
}

func (Metadata_Version) Type() protoreflect.EnumType {
	return &file_metadata_proto_enumTypes[0]
}

func (x Metadata_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Metadata_Version) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Metadata_Version(num)
	return nil
}

// Deprecated: Use Metadata_Version.Descriptor instead.
func (Metadata_Version) EnumDescriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type Metadata_Language int32

const (
	Metadata_UNKNOWN_LANGUAGE Metadata_Language = 0
	Metadata_en               Metadata_Language = 1
	Metadata_aa               Metadata_Language = 2
	Metadata_ab               Metadata_Language = 3
	Metadata_ae               Metadata_Language = 4
	Metadata_af               Metadata_Language = 5
	Metadata_ak               Metadata_Language = 6
	Metadata_am               Metadata_Language = 7
	Metadata_an               Metadata_Language = 8
	Metadata_ar               Metadata_Language = 9
	Metadata_as               Metadata_Language = 10
	Metadata_av               Metadata_Language = 11
	Metadata_ay               Metadata_Language = 12
	Metadata_az               Metadata_Language = 13
	Metadata_ba               Metadata_Language = 14
	Metadata_be               Metadata_Language = 15
	Metadata_bg               Metadata_Language = 16
	Metadata_bh               Metadata_Language = 17
	Metadata_bi               Metadata_Language = 18
	Metadata_bm               Metadata_Language = 19
	Metadata_bn               Metadata_Language = 20
	Metadata_bo               Metadata_Language = 21
	Metadata_br               Metadata_Language = 22
	Metadata_bs               Metadata_Language = 23
	Metadata_ca               Metadata_Language = 24
	Metadata_ce               Metadata_Language = 25
	Metadata_ch               Metadata_Language = 26
	Metadata_co               Metadata_Language = 27
	Metadata_cr               Metadata_Language = 28
	Metadata_cs               Metadata_Language = 29
	Metadata_cu               Metadata_Language = 30
	Metadata_cv               Metadata_Language = 31
	Metadata_cy               Metadata_Language = 32
	Metadata_da               Metadata_Language = 33
	Metadata_de               Metadata_Language = 34
	Metadata_dv               Metadata_Language = 35
	Metadata_dz               Metadata_Language = 36
	Metadata_ee               Metadata_Language = 37
	Metadata_el               Metadata_Language = 38
	Metadata_eo               Metadata_Language = 39
	Metadata_es               Metadata_Language = 40
	Metadata_et               Metadata_Language = 41
	Metadata_eu               Metadata_Language = 42
	Metadata_fa               Metadata_Language = 43
	Metadata_ff               Metadata_Language = 44
	Metadata_fi               Metadata_Language = 45
	Metadata_fj               Metadata_Language = 46
	Metadata_fo               Metadata_Language = 47
	Metadata_fr               Metadata_Language = 48
	Metadata_fy               Metadata_Language = 49
	Metadata_ga               Metadata_Language = 50
	Metadata_gd               Metadata_Language = 51
	Metadata_gl               Metadata_Language = 52
	Metadata_gn               Metadata_Language = 53
	Metadata_gu               Metadata_Language = 54
	Metadata_gv               Metadata_Language = 55
	Metadata_ha               Metadata_Language = 56
	Metadata_he               Metadata_Language = 57
	Metadata_hi               Metadata_Language = 58
	Metadata_ho               Metadata_Language = 59
	Metadata_hr               Metadata_Language = 60
	Metadata_ht               Metadata_Language = 61
	Metadata_hu               Metadata_Language = 62
	Metadata_hy               Metadata_Language = 63
	Metadata_hz               Metadata_Language = 64
	Metadata_ia               Metadata_Language = 65
	Metadata_id               Metadata_Language = 66
	Metadata_ie               Metadata_Language = 67
	Metadata_ig               Metadata_Language = 68
	Metadata_ii               Metadata_Language = 69
	Metadata_ik               Metadata_Language = 70
	Metadata_io               Metadata_Language = 71
	Metadata_is               Metadata_Language = 72
	Metadata_it               Metadata_Language = 73
	Metadata_iu               Metadata_Language = 74
	Metadata_ja               Metadata_Language = 75
	Metadata_jv               Metadata_Language = 76
	Metadata_ka               Metadata_Language = 77
	Metadata_kg               Metadata_Language = 78
	Metadata_ki               Metadata_Language = 79
	Metadata_kj               Metadata_Language = 80
	Metadata_kk               Metadata_Language = 81
	Metadata_kl               Metadata_Language = 82
	Metadata_km               Metadata_Language = 83
	Metadata_kn               Metadata_Language = 84
	Metadata_ko               Metadata_Language = 85
	Metadata_kr               Metadata_Language = 86
	Metadata_ks               Metadata_Language = 87
	Metadata_ku               Metadata_Language = 88
	Metadata_kv               Metadata_Language = 89
	Metadata_kw               Metadata_Language = 90
	Metadata_ky               Metadata_Language = 91
	Metadata_la               Metadata_Language = 92
	Metadata_lb               Metadata_Language = 93
	Metadata_lg               Metadata_Language = 94
	Metadata_li               Metadata_Language = 95
	Metadata_ln               Metadata_Language = 96
	Metadata_lo               Metadata_Language = 97
	Metadata_lt               Metadata_Language = 98
	Metadata_lu               Metadata_Language = 99
	Metadata_lv               Metadata_Language = 100
	Metadata_mg               Metadata_Language = 101
	Metadata_mh               Metadata_Language = 102
	Metadata_mi               Metadata_Language = 103
	Metadata_mk               Metadata_Language = 104
	Metadata_ml               Metadata_Language = 105
	Metadata_mn               Metadata_Language = 106
	Metadata_mr               Metadata_Language = 107
	Metadata_ms               Metadata_Language = 108
	Metadata_mt               Metadata_Language = 109
	Metadata_my               Metadata_Language = 110
	Metadata_na               Metadata_Language = 111
	Metadata_nb               Metadata_Language = 112
	Metadata_nd               Metadata_Language = 113
	Metadata_ne               Metadata_Language = 114
	Metadata_ng               Metadata_Language = 115
	Metadata_nl               Metadata_Language = 116
	Metadata_nn               Metadata_Language = 117
	Metadata_no               Metadata_Language = 118
	Metadata_nr               Metadata_Language = 119
	Metadata_nv               Metadata_Language = 120
	Metadata_ny               Metadata_Language = 121
	Metadata_oc               Metadata_Language = 122
	Metadata_oj               Metadata_Language = 123
	Metadata_om               Metadata_Language = 124
	Metadata_or               Metadata_Language = 125
	Metadata_os               Metadata_Language = 126
	Metadata_pa               Metadata_Language = 127
	Metadata_pi               Metadata_Language = 128
	Metadata_pl               Metadata_Language = 129
	Metadata_ps               Metadata_Language = 130
	Metadata_pt               Metadata_Language = 131
	Metadata_qu               Metadata_Language = 132
	Metadata_rm               Metadata_Language = 133
	Metadata_rn               Metadata_Language = 134
	Metadata_ro               Metadata_Language = 135
	Metadata_ru               Metadata_Language = 136
	Metadata_rw               Metadata_Language = 137
	Metadata_sa               Metadata_Language = 138
	Metadata_sc               Metadata_Language = 139
	Metadata_sd               Metadata_Language = 140
	Metadata_se               Metadata_Language = 141
	Metadata_sg               Metadata_Language = 142
	Metadata_si               Metadata_Language = 143
	Metadata_sk               Metadata_Language = 144
	Metadata_sl               Metadata_Language = 145
	Metadata_sm               Metadata_Language = 146
	Metadata_sn               Metadata_Language = 147
	Metadata_so               Metadata_Language = 148
	Metadata_sq               Metadata_Language = 149
	Metadata_sr               Metadata_Language = 150
	Metadata_ss               Metadata_Language = 151
	Metadata_st               Metadata_Language = 152
	Metadata_su               Metadata_Language = 153
	Metadata_sv               Metadata_Language = 154
	Metadata_sw               Metadata_Language = 155
	Metadata_ta               Metadata_Language = 156
	Metadata_te               Metadata_Language = 157
	Metadata_tg               Metadata_Language = 158
	Metadata_th               Metadata_Language = 159
	Metadata_ti               Metadata_Language = 160
	Metadata_tk               Metadata_Language = 161
	Metadata_tl               Metadata_Language = 162
	Metadata_tn               Metadata_Language = 163
	Metadata_to               Metadata_Language = 164
	Metadata_tr               Metadata_Language = 165
	Metadata_ts               Metadata_Language = 166
	Metadata_tt               Metadata_Language = 167
	Metadata_tw               Metadata_Language = 168
	Metadata_ty               Metadata_Language = 169
	Metadata_ug               Metadata_Language = 170
	Metadata_uk               Metadata_Language = 171
	Metadata_ur               Metadata_Language = 172
	Metadata_uz               Metadata_Language = 173
	Metadata_ve               Metadata_Language = 174
	Metadata_vi               Metadata_Language = 175
	Metadata_vo               Metadata_Language = 176
	Metadata_wa               Metadata_Language = 177
	Metadata_wo               Metadata_Language = 178
	Metadata_xh               Metadata_Language = 179
	Metadata_yi               Metadata_Language = 180
	Metadata_yo               Metadata_Language = 181
	Metadata_za               Metadata_Language = 182
	Metadata_zh               Metadata_Language = 183
	Metadata_zu               Metadata_Language = 184
)

// Enum value maps for Metadata_Language.
var (
	Metadata_Language_name = map[int32]string{
		0:   "UNKNOWN_LANGUAGE",
		1:   "en",
		2:   "aa",
		3:   "ab",
		4:   "ae",
		5:   "af",
		6:   "ak",
		7:   "am",
		8:   "an",
		9:   "ar",
		10:  "as",
		11:  "av",
		12:  "ay",
		13:  "az",
		14:  "ba",
		15:  "be",
		16:  "bg",
		17:  "bh",
		18:  "bi",
		19:  "bm",
		20:  "bn",
		21:  "bo",
		22:  "br",
		23:  "bs",
		24:  "ca",
		25:  "ce",
		26:  "ch",
		27:  "co",
		28:  "cr",
		29:  "cs",
		30:  "cu",
		31:  "cv",
		32:  "cy",
		33:  "da",
		34:  "de",
		35:  "dv",
		36:  "dz",
		37:  "ee",
		38:  "el",
		39:  "eo",
		40:  "es",
		41:  "et",
		42:  "eu",
		43:  "fa",
		44:  "ff",
		45:  "fi",
		46:  "fj",
		47:  "fo",
		48:  "fr",
		49:  "fy",
		50:  "ga",
		51:  "gd",
		52:  "gl",
		53:  "gn",
		54:  "gu",
		55:  "gv",
		56:  "ha",
		57:  "he",
		58:  "hi",
		59:  "ho",
		60:  "hr",
		61:  "ht",
		62:  "hu",
		63:  "hy",
		64:  "hz",
		65:  "ia",
		66:  "id",
		67:  "ie",
		68:  "ig",
		69:  "ii",
		70:  "ik",
		71:  "io",
		72:  "is",
		73:  "it",
		74:  "iu",
		75:  "ja",
		76:  "jv",
		77:  "ka",
		78:  "kg",
		79:  "ki",
		80:  "kj",
		81:  "kk",
		82:  "kl",
		83:  "km",
		84:  "kn",
		85:  "ko",
		86:  "kr",
		87:  "ks",
		88:  "ku",
		89:  "kv",
		90:  "kw",
		91:  "ky",
		92:  "la",
		93:  "lb",
		94:  "lg",
		95:  "li",
		96:  "ln",
		97:  "lo",
		98:  "lt",
		99:  "lu",
		100: "lv",
		101: "mg",
		102: "mh",
		103: "mi",
		104: "mk",
		105: "ml",
		106: "mn",
		107: "mr",
		108: "ms",
		109: "mt",
		110: "my",
		111: "na",
		112: "nb",
		113: "nd",
		114: "ne",
		115: "ng",
		116: "nl",
		117: "nn",
		118: "no",
		119: "nr",
		120: "nv",
		121: "ny",
		122: "oc",
		123: "oj",
		124: "om",
		125: "or",
		126: "os",
		127: "pa",
		128: "pi",
		129: "pl",
		130: "ps",
		131: "pt",
		132: "qu",
		133: "rm",
		134: "rn",
		135: "ro",
		136: "ru",
		137: "rw",
		138: "sa",
		139: "sc",
		140: "sd",
		141: "se",
		142: "sg",
		143: "si",
		144: "sk",
		145: "sl",
		146: "sm",
		147: "sn",
		148: "so",
		149: "sq",
		150: "sr",
		151: "ss",
		152: "st",
		153: "su",
		154: "sv",
		155: "sw",
		156: "ta",
		157: "te",
		158: "tg",
		159: "th",
		160: "ti",
		161: "tk",
		162: "tl",
		163: "tn",
		164: "to",
		165: "tr",
		166: "ts",
		167: "tt",
		168: "tw",
		169: "ty",
		170: "ug",
		171: "uk",
		172: "ur",
		173: "uz",
		174: "ve",
		175: "vi",
		176: "vo",
		177: "wa",
		178: "wo",
		179: "xh",
		180: "yi",
		181: "yo",
		182: "za",
		183: "zh",
		184: "zu",
	}
	Metadata_Language_value = map[string]int32{
		"UNKNOWN_LANGUAGE": 0,
		"en":               1,
		"aa":               2,
		"ab":               3,
		"ae":               4,
		"af":               5,
		"ak":               6,
		"am":               7,
		"an":               8,
		"ar":               9,
		"as":               10,
		"av":               11,
		"ay":               12,
		"az":               13,
		"ba":               14,
		"be":               15,
		"bg":               16,
		"bh":               17,
		"bi":               18,
		"bm":               19,
		"bn":               20,
		"bo":               21,
		"br":               22,
		"bs":               23,
		"ca":               24,
		"ce":               25,
		"ch":               26,
		"co":               27,
		"cr":               28,
		"cs":               29,
		"cu":               30,
		"cv":               31,
		"cy":               32,
		"da":               33,
		"de":               34,
		"dv":               35,
		"dz":               36,
		"ee":               37,
		"el":               38,
		"eo":               39,
		"es":               40,
		"et":               41,
		"eu":               42,
		"fa":               43,
		"ff":               44,
		"fi":               45,
		"fj":               46,
		"fo":               47,
		"fr":               48,
		"fy":               49,
		"ga":               50,
		"gd":               51,
		"gl":               52,
		"gn":               53,
		"gu":               54,
		"gv":               55,
		"ha":               56,
		"he":               57,
		"hi":               58,
		"ho":               59,
		"hr":               60,
		"ht":               61,
		"hu":               62,
		"hy":               63,
		"hz":               64,
		"ia":               65,
		"id":               66,
		"ie":               67,
		"ig":               68,
		"ii":               69,
		"ik":               70,
		"io":               71,
		"is":               72,
		"it":               73,
		"iu":               74,
		"ja":               75,
		"jv":               76,
		"ka":               77,
		"kg":               78,
		"ki":               79,
		"kj":               80,
		"kk":               81,
		"kl":               82,
		"km":               83,
		"kn":               84,
		"ko":               85,
		"kr":               86,
		"ks":               87,
		"ku":               88,
		"kv":               89,
		"kw":               90,
		"ky":               91,
		"la":               92,
		"lb":               93,
		"lg":               94,
		"li":               95,
		"ln":               96,
		"lo":               97,
		"lt":               98,
		"lu":               99,
		"lv":               100,
		"mg":               101,
		"mh":               102,
		"mi":               103,
		"mk":               104,
		"ml":               105,
		"mn":               106,
		"mr":               107,
		"ms":               108,
		"mt":               109,
		"my":               110,
		"na":               111,
		"nb":               112,
		"nd":               113,
		"ne":               114,
		"ng":               115,
		"nl":               116,
		"nn":               117,
		"no":               118,
		"nr":               119,
		"nv":               120,
		"ny":               121,
		"oc":               122,
		"oj":               123,
		"om":               124,
		"or":               125,
		"os":               126,
		"pa":               127,
		"pi":               128,
		"pl":               129,
		"ps":               130,
		"pt":               131,
		"qu":               132,
		"rm":               133,
		"rn":               134,
		"ro":               135,
		"ru":               136,
		"rw":               137,
		"sa":               138,
		"sc":               139,
		"sd":               140,
		"se":               141,
		"sg":               142,
		"si":               143,
		"sk":               144,
		"sl":               145,
		"sm":               146,
		"sn":               147,
		"so":               148,
		"sq":               149,
		"sr":               150,
		"ss":               151,
		"st":               152,
		"su":               153,
		"sv":               154,
		"sw":               155,
		"ta":               156,
		"te":               157,
		"tg":               158,
		"th":               159,
		"ti":               160,
		"tk":               161,
		"tl":               162,
		"tn":               163,
		"to":               164,
		"tr":               165,
		"ts":               166,
		"tt":               167,
		"tw":               168,
		"ty":               169,
		"ug":               170,
		"uk":               171,
		"ur":               172,
		"uz":               173,
		"ve":               174,
		"vi":               175,
		"vo":               176,
		"wa":               177,
		"wo":               178,
		"xh":               179,
		"yi":               180,
		"yo":               181,
		"za":               182,
		"zh":               183,
		"zu":               184,
	}
)

func (x Metadata_Language) Enum() *Metadata_Language {
	p := new(Metadata_Language)
	*p = x
	return p
}

func (x Metadata_Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_Language) Descriptor() protoreflect.EnumDescriptor {
	return file_metadata_proto_enumTypes[1].Descriptor()
}

func (Metadata_Language) Type() protoreflect.EnumType {
	return &file_metadata_proto_enumTypes[1]
}

func (x Metadata_Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Metadata_Language) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Metadata_Language(num)
	return nil
}

// Deprecated: Use Metadata_Language.Descriptor instead.
func (Metadata_Language) EnumDescriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0, 1}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     *Metadata_Version  `protobuf:"varint,1,req,name=version,enum=legacy_pb.Metadata_Version" json:"version"`
	Language    *Metadata_Language `protobuf:"varint,2,req,name=language,enum=legacy_pb.Metadata_Language" json:"language"`
	Title       *string            `protobuf:"bytes,3,req,name=title" json:"title"`
	Description *string            `protobuf:"bytes,4,req,name=description" json:"description"`
	Author      *string            `protobuf:"bytes,5,req,name=author" json:"author"`
	License     *string            `protobuf:"bytes,6,req,name=license" json:"license"`
	Nsfw        *bool              `protobuf:"varint,7,req,name=nsfw" json:"nsfw"`
	Fee         *Fee               `protobuf:"bytes,8,opt,name=fee" json:"fee"`
	Thumbnail   *string            `protobuf:"bytes,9,opt,name=thumbnail" json:"thumbnail"`
	Preview     *string            `protobuf:"bytes,10,opt,name=preview" json:"preview"`
	LicenseUrl  *string            `protobuf:"bytes,11,opt,name=licenseUrl" json:"licenseUrl"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetVersion() Metadata_Version {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return Metadata_UNKNOWN_VERSION
}

func (x *Metadata) GetLanguage() Metadata_Language {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return Metadata_UNKNOWN_LANGUAGE
}

func (x *Metadata) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Metadata) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Metadata) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *Metadata) GetLicense() string {
	if x != nil && x.License != nil {
		return *x.License
	}
	return ""
}

func (x *Metadata) GetNsfw() bool {
	if x != nil && x.Nsfw != nil {
		return *x.Nsfw
	}
	return false
}

func (x *Metadata) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Metadata) GetThumbnail() string {
	if x != nil && x.Thumbnail != nil {
		return *x.Thumbnail
	}
	return ""
}

func (x *Metadata) GetPreview() string {
	if x != nil && x.Preview != nil {
		return *x.Preview
	}
	return ""
}

func (x *Metadata) GetLicenseUrl() string {
	if x != nil && x.LicenseUrl != nil {
		return *x.LicenseUrl
	}
	return ""
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x1a, 0x09, 0x66, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x0f, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18,
	0x06, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x73, 0x66, 0x77, 0x18, 0x07, 0x20, 0x02, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x73,
	0x66, 0x77, 0x12, 0x20, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x65, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1e, 0x0a, 0x0a,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x4e, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x5f, 0x30, 0x5f, 0x30, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x5f, 0x30, 0x5f, 0x30,
	0x5f, 0x32, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x5f, 0x30, 0x5f, 0x30, 0x5f, 0x33, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x5f, 0x30, 0x5f, 0x31, 0x5f, 0x30, 0x10, 0x04, 0x22, 0x99, 0x0c, 0x0a,
	0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x61, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x62, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x65, 0x10, 0x04, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x66, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x6b, 0x10, 0x06, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x6d, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x6e, 0x10, 0x08, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x72, 0x10, 0x09, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x73, 0x10, 0x0a, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x76, 0x10, 0x0b, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x79, 0x10, 0x0c, 0x12,
	0x06, 0x0a, 0x02, 0x61, 0x7a, 0x10, 0x0d, 0x12, 0x06, 0x0a, 0x02, 0x62, 0x61, 0x10, 0x0e, 0x12,
	0x06, 0x0a, 0x02, 0x62, 0x65, 0x10, 0x0f, 0x12, 0x06, 0x0a, 0x02, 0x62, 0x67, 0x10, 0x10, 0x12,
	0x06, 0x0a, 0x02, 0x62, 0x68, 0x10, 0x11, 0x12, 0x06, 0x0a, 0x02, 0x62, 0x69, 0x10, 0x12, 0x12,
	0x06, 0x0a, 0x02, 0x62, 0x6d, 0x10, 0x13, 0x12, 0x06, 0x0a, 0x02, 0x62, 0x6e, 0x10, 0x14, 0x12,
	0x06, 0x0a, 0x02, 0x62, 0x6f, 0x10, 0x15, 0x12, 0x06, 0x0a, 0x02, 0x62, 0x72, 0x10, 0x16, 0x12,
	0x06, 0x0a, 0x02, 0x62, 0x73, 0x10, 0x17, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x61, 0x10, 0x18, 0x12,
	0x06, 0x0a, 0x02, 0x63, 0x65, 0x10, 0x19, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x68, 0x10, 0x1a, 0x12,
	0x06, 0x0a, 0x02, 0x63, 0x6f, 0x10, 0x1b, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x72, 0x10, 0x1c, 0x12,
	0x06, 0x0a, 0x02, 0x63, 0x73, 0x10, 0x1d, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x75, 0x10, 0x1e, 0x12,
	0x06, 0x0a, 0x02, 0x63, 0x76, 0x10, 0x1f, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x79, 0x10, 0x20, 0x12,
	0x06, 0x0a, 0x02, 0x64, 0x61, 0x10, 0x21, 0x12, 0x06, 0x0a, 0x02, 0x64, 0x65, 0x10, 0x22, 0x12,
	0x06, 0x0a, 0x02, 0x64, 0x76, 0x10, 0x23, 0x12, 0x06, 0x0a, 0x02, 0x64, 0x7a, 0x10, 0x24, 0x12,
	0x06, 0x0a, 0x02, 0x65, 0x65, 0x10, 0x25, 0x12, 0x06, 0x0a, 0x02, 0x65, 0x6c, 0x10, 0x26, 0x12,
	0x06, 0x0a, 0x02, 0x65, 0x6f, 0x10, 0x27, 0x12, 0x06, 0x0a, 0x02, 0x65, 0x73, 0x10, 0x28, 0x12,
	0x06, 0x0a, 0x02, 0x65, 0x74, 0x10, 0x29, 0x12, 0x06, 0x0a, 0x02, 0x65, 0x75, 0x10, 0x2a, 0x12,
	0x06, 0x0a, 0x02, 0x66, 0x61, 0x10, 0x2b, 0x12, 0x06, 0x0a, 0x02, 0x66, 0x66, 0x10, 0x2c, 0x12,
	0x06, 0x0a, 0x02, 0x66, 0x69, 0x10, 0x2d, 0x12, 0x06, 0x0a, 0x02, 0x66, 0x6a, 0x10, 0x2e, 0x12,
	0x06, 0x0a, 0x02, 0x66, 0x6f, 0x10, 0x2f, 0x12, 0x06, 0x0a, 0x02, 0x66, 0x72, 0x10, 0x30, 0x12,
	0x06, 0x0a, 0x02, 0x66, 0x79, 0x10, 0x31, 0x12, 0x06, 0x0a, 0x02, 0x67, 0x61, 0x10, 0x32, 0x12,
	0x06, 0x0a, 0x02, 0x67, 0x64, 0x10, 0x33, 0x12, 0x06, 0x0a, 0x02, 0x67, 0x6c, 0x10, 0x34, 0x12,
	0x06, 0x0a, 0x02, 0x67, 0x6e, 0x10, 0x35, 0x12, 0x06, 0x0a, 0x02, 0x67, 0x75, 0x10, 0x36, 0x12,
	0x06, 0x0a, 0x02, 0x67, 0x76, 0x10, 0x37, 0x12, 0x06, 0x0a, 0x02, 0x68, 0x61, 0x10, 0x38, 0x12,
	0x06, 0x0a, 0x02, 0x68, 0x65, 0x10, 0x39, 0x12, 0x06, 0x0a, 0x02, 0x68, 0x69, 0x10, 0x3a, 0x12,
	0x06, 0x0a, 0x02, 0x68, 0x6f, 0x10, 0x3b, 0x12, 0x06, 0x0a, 0x02, 0x68, 0x72, 0x10, 0x3c, 0x12,
	0x06, 0x0a, 0x02, 0x68, 0x74, 0x10, 0x3d, 0x12, 0x06, 0x0a, 0x02, 0x68, 0x75, 0x10, 0x3e, 0x12,
	0x06, 0x0a, 0x02, 0x68, 0x79, 0x10, 0x3f, 0x12, 0x06, 0x0a, 0x02, 0x68, 0x7a, 0x10, 0x40, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x61, 0x10, 0x41, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x42, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x65, 0x10, 0x43, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x67, 0x10, 0x44, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x69, 0x10, 0x45, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x6b, 0x10, 0x46, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x6f, 0x10, 0x47, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x73, 0x10, 0x48, 0x12,
	0x06, 0x0a, 0x02, 0x69, 0x74, 0x10, 0x49, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x75, 0x10, 0x4a, 0x12,
	0x06, 0x0a, 0x02, 0x6a, 0x61, 0x10, 0x4b, 0x12, 0x06, 0x0a, 0x02, 0x6a, 0x76, 0x10, 0x4c, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x61, 0x10, 0x4d, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x67, 0x10, 0x4e, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x69, 0x10, 0x4f, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x6a, 0x10, 0x50, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x6b, 0x10, 0x51, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x6c, 0x10, 0x52, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x6d, 0x10, 0x53, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x6e, 0x10, 0x54, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x6f, 0x10, 0x55, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x72, 0x10, 0x56, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x73, 0x10, 0x57, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x75, 0x10, 0x58, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x76, 0x10, 0x59, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x77, 0x10, 0x5a, 0x12,
	0x06, 0x0a, 0x02, 0x6b, 0x79, 0x10, 0x5b, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x61, 0x10, 0x5c, 0x12,
	0x06, 0x0a, 0x02, 0x6c, 0x62, 0x10, 0x5d, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x67, 0x10, 0x5e, 0x12,
	0x06, 0x0a, 0x02, 0x6c, 0x69, 0x10, 0x5f, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x6e, 0x10, 0x60, 0x12,
	0x06, 0x0a, 0x02, 0x6c, 0x6f, 0x10, 0x61, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x74, 0x10, 0x62, 0x12,
	0x06, 0x0a, 0x02, 0x6c, 0x75, 0x10, 0x63, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x76, 0x10, 0x64, 0x12,
	0x06, 0x0a, 0x02, 0x6d, 0x67, 0x10, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x6d, 0x68, 0x10, 0x66, 0x12,
	0x06, 0x0a, 0x02, 0x6d, 0x69, 0x10, 0x67, 0x12, 0x06, 0x0a, 0x02, 0x6d, 0x6b, 0x10, 0x68, 0x12,
	0x06, 0x0a, 0x02, 0x6d, 0x6c, 0x10, 0x69, 0x12, 0x06, 0x0a, 0x02, 0x6d, 0x6e, 0x10, 0x6a, 0x12,
	0x06, 0x0a, 0x02, 0x6d, 0x72, 0x10, 0x6b, 0x12, 0x06, 0x0a, 0x02, 0x6d, 0x73, 0x10, 0x6c, 0x12,
	0x06, 0x0a, 0x02, 0x6d, 0x74, 0x10, 0x6d, 0x12, 0x06, 0x0a, 0x02, 0x6d, 0x79, 0x10, 0x6e, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x61, 0x10, 0x6f, 0x12, 0x06, 0x0a, 0x02, 0x6e, 0x62, 0x10, 0x70, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x64, 0x10, 0x71, 0x12, 0x06, 0x0a, 0x02, 0x6e, 0x65, 0x10, 0x72, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x67, 0x10, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x6e, 0x6c, 0x10, 0x74, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x6e, 0x10, 0x75, 0x12, 0x06, 0x0a, 0x02, 0x6e, 0x6f, 0x10, 0x76, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x72, 0x10, 0x77, 0x12, 0x06, 0x0a, 0x02, 0x6e, 0x76, 0x10, 0x78, 0x12,
	0x06, 0x0a, 0x02, 0x6e, 0x79, 0x10, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x63, 0x10, 0x7a, 0x12,
	0x06, 0x0a, 0x02, 0x6f, 0x6a, 0x10, 0x7b, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x6d, 0x10, 0x7c, 0x12,
	0x06, 0x0a, 0x02, 0x6f, 0x72, 0x10, 0x7d, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x73, 0x10, 0x7e, 0x12,
	0x06, 0x0a, 0x02, 0x70, 0x61, 0x10, 0x7f, 0x12, 0x07, 0x0a, 0x02, 0x70, 0x69, 0x10, 0x80, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x70, 0x6c, 0x10, 0x81, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x70, 0x73, 0x10,
	0x82, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x70, 0x74, 0x10, 0x83, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x71,
	0x75, 0x10, 0x84, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x72, 0x6d, 0x10, 0x85, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x72, 0x6e, 0x10, 0x86, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x72, 0x6f, 0x10, 0x87, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x72, 0x75, 0x10, 0x88, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x72, 0x77, 0x10, 0x89,
	0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x61, 0x10, 0x8a, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x63,
	0x10, 0x8b, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x64, 0x10, 0x8c, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x73, 0x65, 0x10, 0x8d, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x67, 0x10, 0x8e, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x73, 0x69, 0x10, 0x8f, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x6b, 0x10, 0x90, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x73, 0x6c, 0x10, 0x91, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x6d, 0x10,
	0x92, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x6e, 0x10, 0x93, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73,
	0x6f, 0x10, 0x94, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x71, 0x10, 0x95, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x73, 0x72, 0x10, 0x96, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x73, 0x10, 0x97, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x73, 0x74, 0x10, 0x98, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x75, 0x10, 0x99,
	0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x76, 0x10, 0x9a, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x73, 0x77,
	0x10, 0x9b, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x61, 0x10, 0x9c, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x74, 0x65, 0x10, 0x9d, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x67, 0x10, 0x9e, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x74, 0x68, 0x10, 0x9f, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x69, 0x10, 0xa0, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x74, 0x6b, 0x10, 0xa1, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x6c, 0x10,
	0xa2, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x6e, 0x10, 0xa3, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74,
	0x6f, 0x10, 0xa4, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x72, 0x10, 0xa5, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x74, 0x73, 0x10, 0xa6, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x74, 0x10, 0xa7, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x74, 0x77, 0x10, 0xa8, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x74, 0x79, 0x10, 0xa9,
	0x01, 0x12, 0x07, 0x0a, 0x02, 0x75, 0x67, 0x10, 0xaa, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x75, 0x6b,
	0x10, 0xab, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x75, 0x72, 0x10, 0xac, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x75, 0x7a, 0x10, 0xad, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x76, 0x65, 0x10, 0xae, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x76, 0x69, 0x10, 0xaf, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x76, 0x6f, 0x10, 0xb0, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x77, 0x61, 0x10, 0xb1, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x77, 0x6f, 0x10,
	0xb2, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x78, 0x68, 0x10, 0xb3, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x79,
	0x69, 0x10, 0xb4, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x79, 0x6f, 0x10, 0xb5, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x7a, 0x61, 0x10, 0xb6, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x7a, 0x68, 0x10, 0xb7, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x7a, 0x75, 0x10, 0xb8, 0x01,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_metadata_proto_goTypes = []interface{}{
	(Metadata_Version)(0),  // 0: legacy_pb.Metadata.Version
	(Metadata_Language)(0), // 1: legacy_pb.Metadata.Language
	(*Metadata)(nil),       // 2: legacy_pb.Metadata
	(*Fee)(nil),            // 3: legacy_pb.Fee
}
var file_metadata_proto_depIdxs = []int32{
	0, // 0: legacy_pb.Metadata.version:type_name -> legacy_pb.Metadata.Version
	1, // 1: legacy_pb.Metadata.language:type_name -> legacy_pb.Metadata.Language
	3, // 2: legacy_pb.Metadata.fee:type_name -> legacy_pb.Fee
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	file_fee_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		EnumInfos:         file_metadata_proto_enumTypes,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}
