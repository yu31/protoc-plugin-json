// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public interface TypeFloatOrBuilder extends
    // @@protoc_insertion_point(interface_extends:json.TypeFloat)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Codec specifies encoding format for the field type of float.
   * The default is `Numeric` in plain value, repeated elements and map value.
   * </pre>
   *
   * <code>.json.TypeFloat.Codec codec = 1;</code>
   * @return The enum numeric value on the wire for codec.
   */
  int getCodecValue();
  /**
   * <pre>
   * Codec specifies encoding format for the field type of float.
   * The default is `Numeric` in plain value, repeated elements and map value.
   * </pre>
   *
   * <code>.json.TypeFloat.Codec codec = 1;</code>
   * @return The codec.
   */
  io.github.yu31.protoc.pb.pbjson.TypeFloat.Codec getCodec();
}
