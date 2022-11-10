// source: extension.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

goog.exportSymbol('proto.pb.Extension', null, global);
goog.exportSymbol('proto.pb.Extension.TypeCase', null, global);
goog.exportSymbol('proto.pb.StringMap', null, global);
goog.exportSymbol('proto.pb.StringMap.Value', null, global);
goog.exportSymbol('proto.pb.StringMap.Value.TypeCase', null, global);
goog.exportSymbol('proto.pb.StringMap.Values', null, global);
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
proto.pb.StringMap = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.StringMap, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.StringMap.displayName = 'proto.pb.StringMap';
}
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
proto.pb.StringMap.Value = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.pb.StringMap.Value.oneofGroups_);
};
goog.inherits(proto.pb.StringMap.Value, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.StringMap.Value.displayName = 'proto.pb.StringMap.Value';
}
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
proto.pb.StringMap.Values = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.StringMap.Values.repeatedFields_, null);
};
goog.inherits(proto.pb.StringMap.Values, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.StringMap.Values.displayName = 'proto.pb.StringMap.Values';
}
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
proto.pb.Extension = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.pb.Extension.oneofGroups_);
};
goog.inherits(proto.pb.Extension, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.Extension.displayName = 'proto.pb.Extension';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.StringMap.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.StringMap.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.StringMap} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.toObject = function(includeInstance, msg) {
  var f, obj = {
    fieldsMap: (f = msg.getFieldsMap()) ? f.toObject(includeInstance, proto.pb.StringMap.Values.toObject) : []
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
 * @return {!proto.pb.StringMap}
 */
proto.pb.StringMap.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.StringMap;
  return proto.pb.StringMap.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.StringMap} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.StringMap}
 */
proto.pb.StringMap.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getFieldsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.pb.StringMap.Values.deserializeBinaryFromReader, "", new proto.pb.StringMap.Values());
         });
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
proto.pb.StringMap.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.StringMap.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.StringMap} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFieldsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.pb.StringMap.Values.serializeBinaryToWriter);
  }
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.pb.StringMap.Value.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.pb.StringMap.Value.TypeCase = {
  TYPE_NOT_SET: 0,
  STR: 1,
  INT: 2
};

/**
 * @return {proto.pb.StringMap.Value.TypeCase}
 */
proto.pb.StringMap.Value.prototype.getTypeCase = function() {
  return /** @type {proto.pb.StringMap.Value.TypeCase} */(jspb.Message.computeOneofCase(this, proto.pb.StringMap.Value.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.StringMap.Value.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.StringMap.Value.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.StringMap.Value} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.Value.toObject = function(includeInstance, msg) {
  var f, obj = {
    str: jspb.Message.getFieldWithDefault(msg, 1, ""),
    pb_int: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.pb.StringMap.Value}
 */
proto.pb.StringMap.Value.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.StringMap.Value;
  return proto.pb.StringMap.Value.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.StringMap.Value} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.StringMap.Value}
 */
proto.pb.StringMap.Value.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setStr(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInt(value);
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
proto.pb.StringMap.Value.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.StringMap.Value.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.StringMap.Value} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.Value.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional string str = 1;
 * @return {string}
 */
proto.pb.StringMap.Value.prototype.getStr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.pb.StringMap.Value} returns this
 */
proto.pb.StringMap.Value.prototype.setStr = function(value) {
  return jspb.Message.setOneofField(this, 1, proto.pb.StringMap.Value.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.pb.StringMap.Value} returns this
 */
proto.pb.StringMap.Value.prototype.clearStr = function() {
  return jspb.Message.setOneofField(this, 1, proto.pb.StringMap.Value.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.StringMap.Value.prototype.hasStr = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int64 int = 2;
 * @return {number}
 */
proto.pb.StringMap.Value.prototype.getInt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.pb.StringMap.Value} returns this
 */
proto.pb.StringMap.Value.prototype.setInt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.pb.StringMap.Value.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.pb.StringMap.Value} returns this
 */
proto.pb.StringMap.Value.prototype.clearInt = function() {
  return jspb.Message.setOneofField(this, 2, proto.pb.StringMap.Value.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.StringMap.Value.prototype.hasInt = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.StringMap.Values.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.StringMap.Values.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.StringMap.Values.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.StringMap.Values} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.Values.toObject = function(includeInstance, msg) {
  var f, obj = {
    vsList: jspb.Message.toObjectList(msg.getVsList(),
    proto.pb.StringMap.Value.toObject, includeInstance)
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
 * @return {!proto.pb.StringMap.Values}
 */
proto.pb.StringMap.Values.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.StringMap.Values;
  return proto.pb.StringMap.Values.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.StringMap.Values} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.StringMap.Values}
 */
proto.pb.StringMap.Values.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.StringMap.Value;
      reader.readMessage(value,proto.pb.StringMap.Value.deserializeBinaryFromReader);
      msg.addVs(value);
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
proto.pb.StringMap.Values.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.StringMap.Values.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.StringMap.Values} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.StringMap.Values.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.StringMap.Value.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Value vs = 1;
 * @return {!Array<!proto.pb.StringMap.Value>}
 */
proto.pb.StringMap.Values.prototype.getVsList = function() {
  return /** @type{!Array<!proto.pb.StringMap.Value>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.StringMap.Value, 1));
};


/**
 * @param {!Array<!proto.pb.StringMap.Value>} value
 * @return {!proto.pb.StringMap.Values} returns this
*/
proto.pb.StringMap.Values.prototype.setVsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.StringMap.Value=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.StringMap.Value}
 */
proto.pb.StringMap.Values.prototype.addVs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.StringMap.Value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.pb.StringMap.Values} returns this
 */
proto.pb.StringMap.Values.prototype.clearVsList = function() {
  return this.setVsList([]);
};


/**
 * map<string, Values> fields = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.pb.StringMap.Values>}
 */
proto.pb.StringMap.prototype.getFieldsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.pb.StringMap.Values>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      proto.pb.StringMap.Values));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.pb.StringMap} returns this
 */
proto.pb.StringMap.prototype.clearFieldsMap = function() {
  this.getFieldsMap().clear();
  return this;};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.pb.Extension.oneofGroups_ = [[1]];

/**
 * @enum {number}
 */
proto.pb.Extension.TypeCase = {
  TYPE_NOT_SET: 0,
  MAP: 1
};

/**
 * @return {proto.pb.Extension.TypeCase}
 */
proto.pb.Extension.prototype.getTypeCase = function() {
  return /** @type {proto.pb.Extension.TypeCase} */(jspb.Message.computeOneofCase(this, proto.pb.Extension.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.Extension.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.Extension.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.Extension} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Extension.toObject = function(includeInstance, msg) {
  var f, obj = {
    map: (f = msg.getMap()) && proto.pb.StringMap.toObject(includeInstance, f)
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
 * @return {!proto.pb.Extension}
 */
proto.pb.Extension.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.Extension;
  return proto.pb.Extension.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.Extension} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.Extension}
 */
proto.pb.Extension.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.StringMap;
      reader.readMessage(value,proto.pb.StringMap.deserializeBinaryFromReader);
      msg.setMap(value);
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
proto.pb.Extension.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.Extension.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.Extension} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.Extension.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMap();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.StringMap.serializeBinaryToWriter
    );
  }
};


/**
 * optional StringMap map = 1;
 * @return {?proto.pb.StringMap}
 */
proto.pb.Extension.prototype.getMap = function() {
  return /** @type{?proto.pb.StringMap} */ (
    jspb.Message.getWrapperField(this, proto.pb.StringMap, 1));
};


/**
 * @param {?proto.pb.StringMap|undefined} value
 * @return {!proto.pb.Extension} returns this
*/
proto.pb.Extension.prototype.setMap = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.pb.Extension.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.pb.Extension} returns this
 */
proto.pb.Extension.prototype.clearMap = function() {
  return this.setMap(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.Extension.prototype.hasMap = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.pb);
