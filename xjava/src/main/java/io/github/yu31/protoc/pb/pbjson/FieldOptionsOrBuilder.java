// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public interface FieldOptionsOrBuilder extends
    // @@protoc_insertion_point(interface_extends:json.FieldOptions)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return Whether the json field is set.
   */
  boolean hasJson();
  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return The json.
   */
  java.lang.String getJson();
  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return The bytes for json.
   */
  com.google.protobuf.ByteString
      getJsonBytes();

  /**
   * <pre>
   * The field will be ignored if true.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>bool ignore = 2;</code>
   * @return The ignore.
   */
  boolean getIgnore();

  /**
   * <pre>
   * Same as the go struct tag `json:"xxx,omitempty"`.
   * omitempty indicates whether omit the field if it is empty in MarshalJSON.
   * </pre>
   *
   * <code>bool omitempty = 3;</code>
   * @return The omitempty.
   */
  boolean getOmitempty();

  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   * @return Whether the reference field is set.
   */
  boolean hasReference();
  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   * @return The reference.
   */
  io.github.yu31.protoc.pb.pbjson.TypeReference getReference();
  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   */
  io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder getReferenceOrBuilder();
}
