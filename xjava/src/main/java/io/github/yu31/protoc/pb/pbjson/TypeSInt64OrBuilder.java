// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public interface TypeSInt64OrBuilder extends
    // @@protoc_insertion_point(interface_extends:json.TypeSInt64)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Codec specifies encoding format for the field type of sint64.
   * The default is `Number` in plain value, repeated elements and map value.
   * And the default is `String` in map key.
   * </pre>
   *
   * <code>.json.TypeSInt64.Codec codec = 1;</code>
   * @return The enum numeric value on the wire for codec.
   */
  int getCodecValue();
  /**
   * <pre>
   * Codec specifies encoding format for the field type of sint64.
   * The default is `Number` in plain value, repeated elements and map value.
   * And the default is `String` in map key.
   * </pre>
   *
   * <code>.json.TypeSInt64.Codec codec = 1;</code>
   * @return The codec.
   */
  io.github.yu31.protoc.pb.pbjson.TypeSInt64.Codec getCodec();
}
