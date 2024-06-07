// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

/**
 * <pre>
 * TypeInt32 declares the codec for field type uint32.
 * </pre>
 *
 * Protobuf type {@code json.TypeUint32}
 */
public final class TypeUint32 extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:json.TypeUint32)
    TypeUint32OrBuilder {
private static final long serialVersionUID = 0L;
  // Use TypeUint32.newBuilder() to construct.
  private TypeUint32(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private TypeUint32() {
    codec_ = 0;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new TypeUint32();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private TypeUint32(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {
            int rawValue = input.readEnum();

            codec_ = rawValue;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_TypeUint32_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_TypeUint32_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.github.yu31.protoc.pb.pbjson.TypeUint32.class, io.github.yu31.protoc.pb.pbjson.TypeUint32.Builder.class);
  }

  /**
   * Protobuf enum {@code json.TypeUint32.Codec}
   */
  public enum Codec
      implements com.google.protobuf.ProtocolMessageEnum {
    /**
     * <pre>
     * Unset represents the default value will be applied when encoding and decoding JSON.
     * </pre>
     *
     * <code>Unset = 0;</code>
     */
    Unset(0),
    /**
     * <pre>
     * Numeric represents convert the field as numeric when encoding and decoding JSON.
     * </pre>
     *
     * <code>Numeric = 1;</code>
     */
    Numeric(1),
    /**
     * <pre>
     * String represents convert the field as strings when encoding and decoding JSON.
     * </pre>
     *
     * <code>String = 2;</code>
     */
    String(2),
    UNRECOGNIZED(-1),
    ;

    /**
     * <pre>
     * Unset represents the default value will be applied when encoding and decoding JSON.
     * </pre>
     *
     * <code>Unset = 0;</code>
     */
    public static final int Unset_VALUE = 0;
    /**
     * <pre>
     * Numeric represents convert the field as numeric when encoding and decoding JSON.
     * </pre>
     *
     * <code>Numeric = 1;</code>
     */
    public static final int Numeric_VALUE = 1;
    /**
     * <pre>
     * String represents convert the field as strings when encoding and decoding JSON.
     * </pre>
     *
     * <code>String = 2;</code>
     */
    public static final int String_VALUE = 2;


    public final int getNumber() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalArgumentException(
            "Can't get the number of an unknown enum value.");
      }
      return value;
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     * @deprecated Use {@link #forNumber(int)} instead.
     */
    @java.lang.Deprecated
    public static Codec valueOf(int value) {
      return forNumber(value);
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     */
    public static Codec forNumber(int value) {
      switch (value) {
        case 0: return Unset;
        case 1: return Numeric;
        case 2: return String;
        default: return null;
      }
    }

    public static com.google.protobuf.Internal.EnumLiteMap<Codec>
        internalGetValueMap() {
      return internalValueMap;
    }
    private static final com.google.protobuf.Internal.EnumLiteMap<
        Codec> internalValueMap =
          new com.google.protobuf.Internal.EnumLiteMap<Codec>() {
            public Codec findValueByNumber(int number) {
              return Codec.forNumber(number);
            }
          };

    public final com.google.protobuf.Descriptors.EnumValueDescriptor
        getValueDescriptor() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalStateException(
            "Can't get the descriptor of an unrecognized enum value.");
      }
      return getDescriptor().getValues().get(ordinal());
    }
    public final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptorForType() {
      return getDescriptor();
    }
    public static final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbjson.TypeUint32.getDescriptor().getEnumTypes().get(0);
    }

    private static final Codec[] VALUES = values();

    public static Codec valueOf(
        com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
      if (desc.getType() != getDescriptor()) {
        throw new java.lang.IllegalArgumentException(
          "EnumValueDescriptor is not for this type.");
      }
      if (desc.getIndex() == -1) {
        return UNRECOGNIZED;
      }
      return VALUES[desc.getIndex()];
    }

    private final int value;

    private Codec(int value) {
      this.value = value;
    }

    // @@protoc_insertion_point(enum_scope:json.TypeUint32.Codec)
  }

  public static final int CODEC_FIELD_NUMBER = 1;
  private int codec_;
  /**
   * <pre>
   * Codec specifies encoding format for the field type of uint32.
   * The default is `Numeric` in plain value, repeated elements and map value.
   * And the default is `String` in map key.
   * </pre>
   *
   * <code>.json.TypeUint32.Codec codec = 1;</code>
   * @return The enum numeric value on the wire for codec.
   */
  @java.lang.Override public int getCodecValue() {
    return codec_;
  }
  /**
   * <pre>
   * Codec specifies encoding format for the field type of uint32.
   * The default is `Numeric` in plain value, repeated elements and map value.
   * And the default is `String` in map key.
   * </pre>
   *
   * <code>.json.TypeUint32.Codec codec = 1;</code>
   * @return The codec.
   */
  @java.lang.Override public io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec getCodec() {
    @SuppressWarnings("deprecation")
    io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec result = io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.valueOf(codec_);
    return result == null ? io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.UNRECOGNIZED : result;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (codec_ != io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.Unset.getNumber()) {
      output.writeEnum(1, codec_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (codec_ != io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.Unset.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, codec_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.github.yu31.protoc.pb.pbjson.TypeUint32)) {
      return super.equals(obj);
    }
    io.github.yu31.protoc.pb.pbjson.TypeUint32 other = (io.github.yu31.protoc.pb.pbjson.TypeUint32) obj;

    if (codec_ != other.codec_) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + CODEC_FIELD_NUMBER;
    hash = (53 * hash) + codec_;
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(io.github.yu31.protoc.pb.pbjson.TypeUint32 prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * TypeInt32 declares the codec for field type uint32.
   * </pre>
   *
   * Protobuf type {@code json.TypeUint32}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:json.TypeUint32)
      io.github.yu31.protoc.pb.pbjson.TypeUint32OrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_TypeUint32_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_TypeUint32_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.github.yu31.protoc.pb.pbjson.TypeUint32.class, io.github.yu31.protoc.pb.pbjson.TypeUint32.Builder.class);
    }

    // Construct using io.github.yu31.protoc.pb.pbjson.TypeUint32.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      codec_ = 0;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_TypeUint32_descriptor;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.TypeUint32 getDefaultInstanceForType() {
      return io.github.yu31.protoc.pb.pbjson.TypeUint32.getDefaultInstance();
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.TypeUint32 build() {
      io.github.yu31.protoc.pb.pbjson.TypeUint32 result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.TypeUint32 buildPartial() {
      io.github.yu31.protoc.pb.pbjson.TypeUint32 result = new io.github.yu31.protoc.pb.pbjson.TypeUint32(this);
      result.codec_ = codec_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof io.github.yu31.protoc.pb.pbjson.TypeUint32) {
        return mergeFrom((io.github.yu31.protoc.pb.pbjson.TypeUint32)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.github.yu31.protoc.pb.pbjson.TypeUint32 other) {
      if (other == io.github.yu31.protoc.pb.pbjson.TypeUint32.getDefaultInstance()) return this;
      if (other.codec_ != 0) {
        setCodecValue(other.getCodecValue());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      io.github.yu31.protoc.pb.pbjson.TypeUint32 parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.github.yu31.protoc.pb.pbjson.TypeUint32) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int codec_ = 0;
    /**
     * <pre>
     * Codec specifies encoding format for the field type of uint32.
     * The default is `Numeric` in plain value, repeated elements and map value.
     * And the default is `String` in map key.
     * </pre>
     *
     * <code>.json.TypeUint32.Codec codec = 1;</code>
     * @return The enum numeric value on the wire for codec.
     */
    @java.lang.Override public int getCodecValue() {
      return codec_;
    }
    /**
     * <pre>
     * Codec specifies encoding format for the field type of uint32.
     * The default is `Numeric` in plain value, repeated elements and map value.
     * And the default is `String` in map key.
     * </pre>
     *
     * <code>.json.TypeUint32.Codec codec = 1;</code>
     * @param value The enum numeric value on the wire for codec to set.
     * @return This builder for chaining.
     */
    public Builder setCodecValue(int value) {
      
      codec_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Codec specifies encoding format for the field type of uint32.
     * The default is `Numeric` in plain value, repeated elements and map value.
     * And the default is `String` in map key.
     * </pre>
     *
     * <code>.json.TypeUint32.Codec codec = 1;</code>
     * @return The codec.
     */
    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec getCodec() {
      @SuppressWarnings("deprecation")
      io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec result = io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.valueOf(codec_);
      return result == null ? io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec.UNRECOGNIZED : result;
    }
    /**
     * <pre>
     * Codec specifies encoding format for the field type of uint32.
     * The default is `Numeric` in plain value, repeated elements and map value.
     * And the default is `String` in map key.
     * </pre>
     *
     * <code>.json.TypeUint32.Codec codec = 1;</code>
     * @param value The codec to set.
     * @return This builder for chaining.
     */
    public Builder setCodec(io.github.yu31.protoc.pb.pbjson.TypeUint32.Codec value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      codec_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Codec specifies encoding format for the field type of uint32.
     * The default is `Numeric` in plain value, repeated elements and map value.
     * And the default is `String` in map key.
     * </pre>
     *
     * <code>.json.TypeUint32.Codec codec = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearCodec() {
      
      codec_ = 0;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:json.TypeUint32)
  }

  // @@protoc_insertion_point(class_scope:json.TypeUint32)
  private static final io.github.yu31.protoc.pb.pbjson.TypeUint32 DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.github.yu31.protoc.pb.pbjson.TypeUint32();
  }

  public static io.github.yu31.protoc.pb.pbjson.TypeUint32 getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<TypeUint32>
      PARSER = new com.google.protobuf.AbstractParser<TypeUint32>() {
    @java.lang.Override
    public TypeUint32 parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new TypeUint32(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<TypeUint32> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<TypeUint32> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.github.yu31.protoc.pb.pbjson.TypeUint32 getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

