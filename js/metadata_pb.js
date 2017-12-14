/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var fee_pb = require('./fee_pb.js');
goog.exportSymbol('proto.pb.Metadata', null, global);
goog.exportSymbol('proto.pb.Metadata.Language', null, global);
goog.exportSymbol('proto.pb.Metadata.Version', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.Metadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.Metadata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.pb.Metadata.displayName = 'proto.pb.Metadata';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.Metadata.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.Metadata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.Metadata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Metadata.toObject = function(includeInstance, msg) {
  var f, obj = {
    version: jspb.Message.getField(msg, 1),
    language: jspb.Message.getField(msg, 2),
    title: jspb.Message.getField(msg, 3),
    description: jspb.Message.getField(msg, 4),
    author: jspb.Message.getField(msg, 5),
    license: jspb.Message.getField(msg, 6),
    nsfw: jspb.Message.getField(msg, 7),
    fee: (f = msg.getFee()) && fee_pb.Fee.toObject(includeInstance, f),
    thumbnail: jspb.Message.getField(msg, 9),
    preview: jspb.Message.getField(msg, 10),
    licenseurl: jspb.Message.getField(msg, 11)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.Metadata}
 */
proto.pb.Metadata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.Metadata;
  return proto.pb.Metadata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.Metadata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.Metadata}
 */
proto.pb.Metadata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.pb.Metadata.Version} */ (reader.readEnum());
      msg.setVersion(value);
      break;
    case 2:
      var value = /** @type {!proto.pb.Metadata.Language} */ (reader.readEnum());
      msg.setLanguage(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAuthor(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setLicense(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setNsfw(value);
      break;
    case 8:
      var value = new fee_pb.Fee;
      reader.readMessage(value,fee_pb.Fee.deserializeBinaryFromReader);
      msg.setFee(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setThumbnail(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setPreview(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setLicenseurl(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.Metadata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.Metadata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.Metadata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Metadata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {!proto.pb.Metadata.Version} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = /** @type {!proto.pb.Metadata.Language} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeString(
      3,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeString(
      4,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeString(
      5,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeString(
      6,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getFee();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      fee_pb.Fee.serializeBinaryToWriter
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 9));
  if (f != null) {
    writer.writeString(
      9,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 10));
  if (f != null) {
    writer.writeString(
      10,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 11));
  if (f != null) {
    writer.writeString(
      11,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.pb.Metadata.Version = {
  UNKNOWN_VERSION: 0,
  _0_0_1: 1,
  _0_0_2: 2,
  _0_0_3: 3,
  _0_1_0: 4
};

/**
 * @enum {number}
 */
proto.pb.Metadata.Language = {
  UNKNOWN_LANGUAGE: 0,
  EN: 1,
  AA: 2,
  AB: 3,
  AE: 4,
  AF: 5,
  AK: 6,
  AM: 7,
  AN: 8,
  AR: 9,
  AS: 10,
  AV: 11,
  AY: 12,
  AZ: 13,
  BA: 14,
  BE: 15,
  BG: 16,
  BH: 17,
  BI: 18,
  BM: 19,
  BN: 20,
  BO: 21,
  BR: 22,
  BS: 23,
  CA: 24,
  CE: 25,
  CH: 26,
  CO: 27,
  CR: 28,
  CS: 29,
  CU: 30,
  CV: 31,
  CY: 32,
  DA: 33,
  DE: 34,
  DV: 35,
  DZ: 36,
  EE: 37,
  EL: 38,
  EO: 39,
  ES: 40,
  ET: 41,
  EU: 42,
  FA: 43,
  FF: 44,
  FI: 45,
  FJ: 46,
  FO: 47,
  FR: 48,
  FY: 49,
  GA: 50,
  GD: 51,
  GL: 52,
  GN: 53,
  GU: 54,
  GV: 55,
  HA: 56,
  HE: 57,
  HI: 58,
  HO: 59,
  HR: 60,
  HT: 61,
  HU: 62,
  HY: 63,
  HZ: 64,
  IA: 65,
  ID: 66,
  IE: 67,
  IG: 68,
  II: 69,
  IK: 70,
  IO: 71,
  IS: 72,
  IT: 73,
  IU: 74,
  JA: 75,
  JV: 76,
  KA: 77,
  KG: 78,
  KI: 79,
  KJ: 80,
  KK: 81,
  KL: 82,
  KM: 83,
  KN: 84,
  KO: 85,
  KR: 86,
  KS: 87,
  KU: 88,
  KV: 89,
  KW: 90,
  KY: 91,
  LA: 92,
  LB: 93,
  LG: 94,
  LI: 95,
  LN: 96,
  LO: 97,
  LT: 98,
  LU: 99,
  LV: 100,
  MG: 101,
  MH: 102,
  MI: 103,
  MK: 104,
  ML: 105,
  MN: 106,
  MR: 107,
  MS: 108,
  MT: 109,
  MY: 110,
  NA: 111,
  NB: 112,
  ND: 113,
  NE: 114,
  NG: 115,
  NL: 116,
  NN: 117,
  NO: 118,
  NR: 119,
  NV: 120,
  NY: 121,
  OC: 122,
  OJ: 123,
  OM: 124,
  OR: 125,
  OS: 126,
  PA: 127,
  PI: 128,
  PL: 129,
  PS: 130,
  PT: 131,
  QU: 132,
  RM: 133,
  RN: 134,
  RO: 135,
  RU: 136,
  RW: 137,
  SA: 138,
  SC: 139,
  SD: 140,
  SE: 141,
  SG: 142,
  SI: 143,
  SK: 144,
  SL: 145,
  SM: 146,
  SN: 147,
  SO: 148,
  SQ: 149,
  SR: 150,
  SS: 151,
  ST: 152,
  SU: 153,
  SV: 154,
  SW: 155,
  TA: 156,
  TE: 157,
  TG: 158,
  TH: 159,
  TI: 160,
  TK: 161,
  TL: 162,
  TN: 163,
  TO: 164,
  TR: 165,
  TS: 166,
  TT: 167,
  TW: 168,
  TY: 169,
  UG: 170,
  UK: 171,
  UR: 172,
  UZ: 173,
  VE: 174,
  VI: 175,
  VO: 176,
  WA: 177,
  WO: 178,
  XH: 179,
  YI: 180,
  YO: 181,
  ZA: 182,
  ZH: 183,
  ZU: 184
};

/**
 * required Version version = 1;
 * @return {!proto.pb.Metadata.Version}
 */
proto.pb.Metadata.prototype.getVersion = function() {
  return /** @type {!proto.pb.Metadata.Version} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.pb.Metadata.Version} value */
proto.pb.Metadata.prototype.setVersion = function(value) {
  jspb.Message.setField(this, 1, value);
};


proto.pb.Metadata.prototype.clearVersion = function() {
  jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasVersion = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * required Language language = 2;
 * @return {!proto.pb.Metadata.Language}
 */
proto.pb.Metadata.prototype.getLanguage = function() {
  return /** @type {!proto.pb.Metadata.Language} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.pb.Metadata.Language} value */
proto.pb.Metadata.prototype.setLanguage = function(value) {
  jspb.Message.setField(this, 2, value);
};


proto.pb.Metadata.prototype.clearLanguage = function() {
  jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasLanguage = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * required string title = 3;
 * @return {string}
 */
proto.pb.Metadata.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setTitle = function(value) {
  jspb.Message.setField(this, 3, value);
};


proto.pb.Metadata.prototype.clearTitle = function() {
  jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasTitle = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * required string description = 4;
 * @return {string}
 */
proto.pb.Metadata.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setDescription = function(value) {
  jspb.Message.setField(this, 4, value);
};


proto.pb.Metadata.prototype.clearDescription = function() {
  jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasDescription = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * required string author = 5;
 * @return {string}
 */
proto.pb.Metadata.prototype.getAuthor = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setAuthor = function(value) {
  jspb.Message.setField(this, 5, value);
};


proto.pb.Metadata.prototype.clearAuthor = function() {
  jspb.Message.setField(this, 5, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasAuthor = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * required string license = 6;
 * @return {string}
 */
proto.pb.Metadata.prototype.getLicense = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setLicense = function(value) {
  jspb.Message.setField(this, 6, value);
};


proto.pb.Metadata.prototype.clearLicense = function() {
  jspb.Message.setField(this, 6, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasLicense = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * required bool nsfw = 7;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.pb.Metadata.prototype.getNsfw = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 7, false));
};


/** @param {boolean} value */
proto.pb.Metadata.prototype.setNsfw = function(value) {
  jspb.Message.setField(this, 7, value);
};


proto.pb.Metadata.prototype.clearNsfw = function() {
  jspb.Message.setField(this, 7, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasNsfw = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional Fee fee = 8;
 * @return {?proto.pb.Fee}
 */
proto.pb.Metadata.prototype.getFee = function() {
  return /** @type{?proto.pb.Fee} */ (
    jspb.Message.getWrapperField(this, fee_pb.Fee, 8));
};


/** @param {?proto.pb.Fee|undefined} value */
proto.pb.Metadata.prototype.setFee = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.pb.Metadata.prototype.clearFee = function() {
  this.setFee(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasFee = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string thumbnail = 9;
 * @return {string}
 */
proto.pb.Metadata.prototype.getThumbnail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setThumbnail = function(value) {
  jspb.Message.setField(this, 9, value);
};


proto.pb.Metadata.prototype.clearThumbnail = function() {
  jspb.Message.setField(this, 9, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasThumbnail = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional string preview = 10;
 * @return {string}
 */
proto.pb.Metadata.prototype.getPreview = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setPreview = function(value) {
  jspb.Message.setField(this, 10, value);
};


proto.pb.Metadata.prototype.clearPreview = function() {
  jspb.Message.setField(this, 10, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasPreview = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional string licenseUrl = 11;
 * @return {string}
 */
proto.pb.Metadata.prototype.getLicenseurl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.pb.Metadata.prototype.setLicenseurl = function(value) {
  jspb.Message.setField(this, 11, value);
};


proto.pb.Metadata.prototype.clearLicenseurl = function() {
  jspb.Message.setField(this, 11, undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.pb.Metadata.prototype.hasLicenseurl = function() {
  return jspb.Message.getField(this, 11) != null;
};


goog.object.extend(exports, proto.pb);
