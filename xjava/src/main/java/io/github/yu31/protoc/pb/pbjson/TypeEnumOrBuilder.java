// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public interface TypeEnumOrBuilder extends
    // @@protoc_insertion_point(interface_extends:json.TypeEnum)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Codec specifies encoding format for the field type of Enum.
   * The default is `Number` in plain value, repeated elements and map value.
   * </pre>
   *
   * <code>.json.TypeEnum.Codec codec = 1;</code>
   * @return The enum numeric value on the wire for codec.
   */
  int getCodecValue();
  /**
   * <pre>
   * Codec specifies encoding format for the field type of Enum.
   * The default is `Number` in plain value, repeated elements and map value.
   * </pre>
   *
   * <code>.json.TypeEnum.Codec codec = 1;</code>
   * @return The codec.
   */
  io.github.yu31.protoc.pb.pbjson.TypeEnum.Codec getCodec();
}
