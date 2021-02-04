// source: tendermint/rpc/grpc/types.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var tendermint_abci_types_pb = require('../../../tendermint/abci/types_pb.js');
goog.object.extend(proto, tendermint_abci_types_pb);
goog.exportSymbol('proto.tendermint.rpc.grpc.RequestBroadcastTx', null, global);
goog.exportSymbol('proto.tendermint.rpc.grpc.RequestPing', null, global);
goog.exportSymbol('proto.tendermint.rpc.grpc.ResponseBroadcastTx', null, global);
goog.exportSymbol('proto.tendermint.rpc.grpc.ResponsePing', null, global);
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
proto.tendermint.rpc.grpc.RequestPing = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.rpc.grpc.RequestPing, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.rpc.grpc.RequestPing.displayName = 'proto.tendermint.rpc.grpc.RequestPing';
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
proto.tendermint.rpc.grpc.RequestBroadcastTx = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.rpc.grpc.RequestBroadcastTx, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.rpc.grpc.RequestBroadcastTx.displayName = 'proto.tendermint.rpc.grpc.RequestBroadcastTx';
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
proto.tendermint.rpc.grpc.ResponsePing = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.rpc.grpc.ResponsePing, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.rpc.grpc.ResponsePing.displayName = 'proto.tendermint.rpc.grpc.ResponsePing';
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
proto.tendermint.rpc.grpc.ResponseBroadcastTx = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.rpc.grpc.ResponseBroadcastTx, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.rpc.grpc.ResponseBroadcastTx.displayName = 'proto.tendermint.rpc.grpc.ResponseBroadcastTx';
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
proto.tendermint.rpc.grpc.RequestPing.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.rpc.grpc.RequestPing.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.rpc.grpc.RequestPing} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.RequestPing.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.tendermint.rpc.grpc.RequestPing}
 */
proto.tendermint.rpc.grpc.RequestPing.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.rpc.grpc.RequestPing;
  return proto.tendermint.rpc.grpc.RequestPing.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.rpc.grpc.RequestPing} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.rpc.grpc.RequestPing}
 */
proto.tendermint.rpc.grpc.RequestPing.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.tendermint.rpc.grpc.RequestPing.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.rpc.grpc.RequestPing.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.rpc.grpc.RequestPing} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.RequestPing.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
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
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.rpc.grpc.RequestBroadcastTx.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.toObject = function(includeInstance, msg) {
  var f, obj = {
    tx: msg.getTx_asB64()
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
 * @return {!proto.tendermint.rpc.grpc.RequestBroadcastTx}
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.rpc.grpc.RequestBroadcastTx;
  return proto.tendermint.rpc.grpc.RequestBroadcastTx.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.rpc.grpc.RequestBroadcastTx}
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setTx(value);
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
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.rpc.grpc.RequestBroadcastTx.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTx_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
};


/**
 * optional bytes tx = 1;
 * @return {!(string|Uint8Array)}
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.getTx = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes tx = 1;
 * This is a type-conversion wrapper around `getTx()`
 * @return {string}
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.getTx_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getTx()));
};


/**
 * optional bytes tx = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getTx()`
 * @return {!Uint8Array}
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.getTx_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getTx()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.tendermint.rpc.grpc.RequestBroadcastTx} returns this
 */
proto.tendermint.rpc.grpc.RequestBroadcastTx.prototype.setTx = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
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
proto.tendermint.rpc.grpc.ResponsePing.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.rpc.grpc.ResponsePing.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.rpc.grpc.ResponsePing} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.ResponsePing.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.tendermint.rpc.grpc.ResponsePing}
 */
proto.tendermint.rpc.grpc.ResponsePing.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.rpc.grpc.ResponsePing;
  return proto.tendermint.rpc.grpc.ResponsePing.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.rpc.grpc.ResponsePing} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.rpc.grpc.ResponsePing}
 */
proto.tendermint.rpc.grpc.ResponsePing.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.tendermint.rpc.grpc.ResponsePing.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.rpc.grpc.ResponsePing.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.rpc.grpc.ResponsePing} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.ResponsePing.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
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
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.rpc.grpc.ResponseBroadcastTx.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.toObject = function(includeInstance, msg) {
  var f, obj = {
    checkTx: (f = msg.getCheckTx()) && tendermint_abci_types_pb.ResponseCheckTx.toObject(includeInstance, f),
    deliverTx: (f = msg.getDeliverTx()) && tendermint_abci_types_pb.ResponseDeliverTx.toObject(includeInstance, f)
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
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.rpc.grpc.ResponseBroadcastTx;
  return proto.tendermint.rpc.grpc.ResponseBroadcastTx.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new tendermint_abci_types_pb.ResponseCheckTx;
      reader.readMessage(value,tendermint_abci_types_pb.ResponseCheckTx.deserializeBinaryFromReader);
      msg.setCheckTx(value);
      break;
    case 2:
      var value = new tendermint_abci_types_pb.ResponseDeliverTx;
      reader.readMessage(value,tendermint_abci_types_pb.ResponseDeliverTx.deserializeBinaryFromReader);
      msg.setDeliverTx(value);
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
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.rpc.grpc.ResponseBroadcastTx.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCheckTx();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      tendermint_abci_types_pb.ResponseCheckTx.serializeBinaryToWriter
    );
  }
  f = message.getDeliverTx();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      tendermint_abci_types_pb.ResponseDeliverTx.serializeBinaryToWriter
    );
  }
};


/**
 * optional tendermint.abci.ResponseCheckTx check_tx = 1;
 * @return {?proto.tendermint.abci.ResponseCheckTx}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.getCheckTx = function() {
  return /** @type{?proto.tendermint.abci.ResponseCheckTx} */ (
    jspb.Message.getWrapperField(this, tendermint_abci_types_pb.ResponseCheckTx, 1));
};


/**
 * @param {?proto.tendermint.abci.ResponseCheckTx|undefined} value
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} returns this
*/
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.setCheckTx = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} returns this
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.clearCheckTx = function() {
  return this.setCheckTx(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.hasCheckTx = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional tendermint.abci.ResponseDeliverTx deliver_tx = 2;
 * @return {?proto.tendermint.abci.ResponseDeliverTx}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.getDeliverTx = function() {
  return /** @type{?proto.tendermint.abci.ResponseDeliverTx} */ (
    jspb.Message.getWrapperField(this, tendermint_abci_types_pb.ResponseDeliverTx, 2));
};


/**
 * @param {?proto.tendermint.abci.ResponseDeliverTx|undefined} value
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} returns this
*/
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.setDeliverTx = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.rpc.grpc.ResponseBroadcastTx} returns this
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.clearDeliverTx = function() {
  return this.setDeliverTx(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.rpc.grpc.ResponseBroadcastTx.prototype.hasDeliverTx = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.tendermint.rpc.grpc);