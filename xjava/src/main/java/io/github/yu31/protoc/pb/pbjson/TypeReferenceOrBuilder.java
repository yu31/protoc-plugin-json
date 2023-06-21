// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public interface TypeReferenceOrBuilder extends
    // @@protoc_insertion_point(interface_extends:json.TypeReference)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * PlainOptions declares the tags applied to the plain field.
   * </pre>
   *
   * <code>.json.PlainOptions plain = 1;</code>
   * @return Whether the plain field is set.
   */
  boolean hasPlain();
  /**
   * <pre>
   * PlainOptions declares the tags applied to the plain field.
   * </pre>
   *
   * <code>.json.PlainOptions plain = 1;</code>
   * @return The plain.
   */
  io.github.yu31.protoc.pb.pbjson.PlainOptions getPlain();
  /**
   * <pre>
   * PlainOptions declares the tags applied to the plain field.
   * </pre>
   *
   * <code>.json.PlainOptions plain = 1;</code>
   */
  io.github.yu31.protoc.pb.pbjson.PlainOptionsOrBuilder getPlainOrBuilder();

  /**
   * <pre>
   * RepeatedOptions declares the tags applied to the repeated field.
   * </pre>
   *
   * <code>.json.RepeatedOptions repeated = 2;</code>
   * @return Whether the repeated field is set.
   */
  boolean hasRepeated();
  /**
   * <pre>
   * RepeatedOptions declares the tags applied to the repeated field.
   * </pre>
   *
   * <code>.json.RepeatedOptions repeated = 2;</code>
   * @return The repeated.
   */
  io.github.yu31.protoc.pb.pbjson.RepeatedOptions getRepeated();
  /**
   * <pre>
   * RepeatedOptions declares the tags applied to the repeated field.
   * </pre>
   *
   * <code>.json.RepeatedOptions repeated = 2;</code>
   */
  io.github.yu31.protoc.pb.pbjson.RepeatedOptionsOrBuilder getRepeatedOrBuilder();

  /**
   * <pre>
   * MapTags declares the tags applied to the map field.
   * </pre>
   *
   * <code>.json.MapOptions map = 3;</code>
   * @return Whether the map field is set.
   */
  boolean hasMap();
  /**
   * <pre>
   * MapTags declares the tags applied to the map field.
   * </pre>
   *
   * <code>.json.MapOptions map = 3;</code>
   * @return The map.
   */
  io.github.yu31.protoc.pb.pbjson.MapOptions getMap();
  /**
   * <pre>
   * MapTags declares the tags applied to the map field.
   * </pre>
   *
   * <code>.json.MapOptions map = 3;</code>
   */
  io.github.yu31.protoc.pb.pbjson.MapOptionsOrBuilder getMapOrBuilder();

  public io.github.yu31.protoc.pb.pbjson.TypeReference.KindCase getKindCase();
}
