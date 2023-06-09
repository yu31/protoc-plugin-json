// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

/**
 * Protobuf type {@code json.FieldOptions}
 */
public final class FieldOptions extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:json.FieldOptions)
    FieldOptionsOrBuilder {
private static final long serialVersionUID = 0L;
  // Use FieldOptions.newBuilder() to construct.
  private FieldOptions(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private FieldOptions() {
    json_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new FieldOptions();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private FieldOptions(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
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
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();
            bitField0_ |= 0x00000001;
            json_ = s;
            break;
          }
          case 16: {

            ignore_ = input.readBool();
            break;
          }
          case 24: {

            omitempty_ = input.readBool();
            break;
          }
          case 34: {
            io.github.yu31.protoc.pb.pbjson.TypeReference.Builder subBuilder = null;
            if (reference_ != null) {
              subBuilder = reference_.toBuilder();
            }
            reference_ = input.readMessage(io.github.yu31.protoc.pb.pbjson.TypeReference.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(reference_);
              reference_ = subBuilder.buildPartial();
            }

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
    return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_FieldOptions_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_FieldOptions_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.github.yu31.protoc.pb.pbjson.FieldOptions.class, io.github.yu31.protoc.pb.pbjson.FieldOptions.Builder.class);
  }

  private int bitField0_;
  public static final int JSON_FIELD_NUMBER = 1;
  private volatile java.lang.Object json_;
  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return Whether the json field is set.
   */
  @java.lang.Override
  public boolean hasJson() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return The json.
   */
  @java.lang.Override
  public java.lang.String getJson() {
    java.lang.Object ref = json_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      json_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * The key name in JSON content. Default is the field's name in proto file.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>optional string json = 1;</code>
   * @return The bytes for json.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getJsonBytes() {
    java.lang.Object ref = json_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      json_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int IGNORE_FIELD_NUMBER = 2;
  private boolean ignore_;
  /**
   * <pre>
   * The field will be ignored if true.
   * It effect the MarshalJSON and UnmarshalJSON.
   * </pre>
   *
   * <code>bool ignore = 2;</code>
   * @return The ignore.
   */
  @java.lang.Override
  public boolean getIgnore() {
    return ignore_;
  }

  public static final int OMITEMPTY_FIELD_NUMBER = 3;
  private boolean omitempty_;
  /**
   * <pre>
   * Same as the go struct tag `json:"xxx,omitempty"`.
   * omitempty indicates whether omit the field if it is empty in MarshalJSON.
   * </pre>
   *
   * <code>bool omitempty = 3;</code>
   * @return The omitempty.
   */
  @java.lang.Override
  public boolean getOmitempty() {
    return omitempty_;
  }

  public static final int REFERENCE_FIELD_NUMBER = 4;
  private io.github.yu31.protoc.pb.pbjson.TypeReference reference_;
  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   * @return Whether the reference field is set.
   */
  @java.lang.Override
  public boolean hasReference() {
    return reference_ != null;
  }
  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   * @return The reference.
   */
  @java.lang.Override
  public io.github.yu31.protoc.pb.pbjson.TypeReference getReference() {
    return reference_ == null ? io.github.yu31.protoc.pb.pbjson.TypeReference.getDefaultInstance() : reference_;
  }
  /**
   * <pre>
   * Format set for field.
   * </pre>
   *
   * <code>.json.TypeReference reference = 4;</code>
   */
  @java.lang.Override
  public io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder getReferenceOrBuilder() {
    return getReference();
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
    if (((bitField0_ & 0x00000001) != 0)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, json_);
    }
    if (ignore_ != false) {
      output.writeBool(2, ignore_);
    }
    if (omitempty_ != false) {
      output.writeBool(3, omitempty_);
    }
    if (reference_ != null) {
      output.writeMessage(4, getReference());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (((bitField0_ & 0x00000001) != 0)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, json_);
    }
    if (ignore_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(2, ignore_);
    }
    if (omitempty_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(3, omitempty_);
    }
    if (reference_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(4, getReference());
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
    if (!(obj instanceof io.github.yu31.protoc.pb.pbjson.FieldOptions)) {
      return super.equals(obj);
    }
    io.github.yu31.protoc.pb.pbjson.FieldOptions other = (io.github.yu31.protoc.pb.pbjson.FieldOptions) obj;

    if (hasJson() != other.hasJson()) return false;
    if (hasJson()) {
      if (!getJson()
          .equals(other.getJson())) return false;
    }
    if (getIgnore()
        != other.getIgnore()) return false;
    if (getOmitempty()
        != other.getOmitempty()) return false;
    if (hasReference() != other.hasReference()) return false;
    if (hasReference()) {
      if (!getReference()
          .equals(other.getReference())) return false;
    }
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
    if (hasJson()) {
      hash = (37 * hash) + JSON_FIELD_NUMBER;
      hash = (53 * hash) + getJson().hashCode();
    }
    hash = (37 * hash) + IGNORE_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getIgnore());
    hash = (37 * hash) + OMITEMPTY_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getOmitempty());
    if (hasReference()) {
      hash = (37 * hash) + REFERENCE_FIELD_NUMBER;
      hash = (53 * hash) + getReference().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbjson.FieldOptions parseFrom(
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
  public static Builder newBuilder(io.github.yu31.protoc.pb.pbjson.FieldOptions prototype) {
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
   * Protobuf type {@code json.FieldOptions}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:json.FieldOptions)
      io.github.yu31.protoc.pb.pbjson.FieldOptionsOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_FieldOptions_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_FieldOptions_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.github.yu31.protoc.pb.pbjson.FieldOptions.class, io.github.yu31.protoc.pb.pbjson.FieldOptions.Builder.class);
    }

    // Construct using io.github.yu31.protoc.pb.pbjson.FieldOptions.newBuilder()
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
      json_ = "";
      bitField0_ = (bitField0_ & ~0x00000001);
      ignore_ = false;

      omitempty_ = false;

      if (referenceBuilder_ == null) {
        reference_ = null;
      } else {
        reference_ = null;
        referenceBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.github.yu31.protoc.pb.pbjson.PBJSON.internal_static_json_FieldOptions_descriptor;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.FieldOptions getDefaultInstanceForType() {
      return io.github.yu31.protoc.pb.pbjson.FieldOptions.getDefaultInstance();
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.FieldOptions build() {
      io.github.yu31.protoc.pb.pbjson.FieldOptions result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbjson.FieldOptions buildPartial() {
      io.github.yu31.protoc.pb.pbjson.FieldOptions result = new io.github.yu31.protoc.pb.pbjson.FieldOptions(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        to_bitField0_ |= 0x00000001;
      }
      result.json_ = json_;
      result.ignore_ = ignore_;
      result.omitempty_ = omitempty_;
      if (referenceBuilder_ == null) {
        result.reference_ = reference_;
      } else {
        result.reference_ = referenceBuilder_.build();
      }
      result.bitField0_ = to_bitField0_;
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
      if (other instanceof io.github.yu31.protoc.pb.pbjson.FieldOptions) {
        return mergeFrom((io.github.yu31.protoc.pb.pbjson.FieldOptions)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.github.yu31.protoc.pb.pbjson.FieldOptions other) {
      if (other == io.github.yu31.protoc.pb.pbjson.FieldOptions.getDefaultInstance()) return this;
      if (other.hasJson()) {
        bitField0_ |= 0x00000001;
        json_ = other.json_;
        onChanged();
      }
      if (other.getIgnore() != false) {
        setIgnore(other.getIgnore());
      }
      if (other.getOmitempty() != false) {
        setOmitempty(other.getOmitempty());
      }
      if (other.hasReference()) {
        mergeReference(other.getReference());
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
      io.github.yu31.protoc.pb.pbjson.FieldOptions parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.github.yu31.protoc.pb.pbjson.FieldOptions) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object json_ = "";
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @return Whether the json field is set.
     */
    public boolean hasJson() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @return The json.
     */
    public java.lang.String getJson() {
      java.lang.Object ref = json_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        json_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @return The bytes for json.
     */
    public com.google.protobuf.ByteString
        getJsonBytes() {
      java.lang.Object ref = json_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        json_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @param value The json to set.
     * @return This builder for chaining.
     */
    public Builder setJson(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
      json_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearJson() {
      bitField0_ = (bitField0_ & ~0x00000001);
      json_ = getDefaultInstance().getJson();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The key name in JSON content. Default is the field's name in proto file.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>optional string json = 1;</code>
     * @param value The bytes for json to set.
     * @return This builder for chaining.
     */
    public Builder setJsonBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      bitField0_ |= 0x00000001;
      json_ = value;
      onChanged();
      return this;
    }

    private boolean ignore_ ;
    /**
     * <pre>
     * The field will be ignored if true.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>bool ignore = 2;</code>
     * @return The ignore.
     */
    @java.lang.Override
    public boolean getIgnore() {
      return ignore_;
    }
    /**
     * <pre>
     * The field will be ignored if true.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>bool ignore = 2;</code>
     * @param value The ignore to set.
     * @return This builder for chaining.
     */
    public Builder setIgnore(boolean value) {
      
      ignore_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The field will be ignored if true.
     * It effect the MarshalJSON and UnmarshalJSON.
     * </pre>
     *
     * <code>bool ignore = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearIgnore() {
      
      ignore_ = false;
      onChanged();
      return this;
    }

    private boolean omitempty_ ;
    /**
     * <pre>
     * Same as the go struct tag `json:"xxx,omitempty"`.
     * omitempty indicates whether omit the field if it is empty in MarshalJSON.
     * </pre>
     *
     * <code>bool omitempty = 3;</code>
     * @return The omitempty.
     */
    @java.lang.Override
    public boolean getOmitempty() {
      return omitempty_;
    }
    /**
     * <pre>
     * Same as the go struct tag `json:"xxx,omitempty"`.
     * omitempty indicates whether omit the field if it is empty in MarshalJSON.
     * </pre>
     *
     * <code>bool omitempty = 3;</code>
     * @param value The omitempty to set.
     * @return This builder for chaining.
     */
    public Builder setOmitempty(boolean value) {
      
      omitempty_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Same as the go struct tag `json:"xxx,omitempty"`.
     * omitempty indicates whether omit the field if it is empty in MarshalJSON.
     * </pre>
     *
     * <code>bool omitempty = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearOmitempty() {
      
      omitempty_ = false;
      onChanged();
      return this;
    }

    private io.github.yu31.protoc.pb.pbjson.TypeReference reference_;
    private com.google.protobuf.SingleFieldBuilderV3<
        io.github.yu31.protoc.pb.pbjson.TypeReference, io.github.yu31.protoc.pb.pbjson.TypeReference.Builder, io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder> referenceBuilder_;
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     * @return Whether the reference field is set.
     */
    public boolean hasReference() {
      return referenceBuilder_ != null || reference_ != null;
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     * @return The reference.
     */
    public io.github.yu31.protoc.pb.pbjson.TypeReference getReference() {
      if (referenceBuilder_ == null) {
        return reference_ == null ? io.github.yu31.protoc.pb.pbjson.TypeReference.getDefaultInstance() : reference_;
      } else {
        return referenceBuilder_.getMessage();
      }
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public Builder setReference(io.github.yu31.protoc.pb.pbjson.TypeReference value) {
      if (referenceBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        reference_ = value;
        onChanged();
      } else {
        referenceBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public Builder setReference(
        io.github.yu31.protoc.pb.pbjson.TypeReference.Builder builderForValue) {
      if (referenceBuilder_ == null) {
        reference_ = builderForValue.build();
        onChanged();
      } else {
        referenceBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public Builder mergeReference(io.github.yu31.protoc.pb.pbjson.TypeReference value) {
      if (referenceBuilder_ == null) {
        if (reference_ != null) {
          reference_ =
            io.github.yu31.protoc.pb.pbjson.TypeReference.newBuilder(reference_).mergeFrom(value).buildPartial();
        } else {
          reference_ = value;
        }
        onChanged();
      } else {
        referenceBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public Builder clearReference() {
      if (referenceBuilder_ == null) {
        reference_ = null;
        onChanged();
      } else {
        reference_ = null;
        referenceBuilder_ = null;
      }

      return this;
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public io.github.yu31.protoc.pb.pbjson.TypeReference.Builder getReferenceBuilder() {
      
      onChanged();
      return getReferenceFieldBuilder().getBuilder();
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    public io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder getReferenceOrBuilder() {
      if (referenceBuilder_ != null) {
        return referenceBuilder_.getMessageOrBuilder();
      } else {
        return reference_ == null ?
            io.github.yu31.protoc.pb.pbjson.TypeReference.getDefaultInstance() : reference_;
      }
    }
    /**
     * <pre>
     * Format set for field.
     * </pre>
     *
     * <code>.json.TypeReference reference = 4;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        io.github.yu31.protoc.pb.pbjson.TypeReference, io.github.yu31.protoc.pb.pbjson.TypeReference.Builder, io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder> 
        getReferenceFieldBuilder() {
      if (referenceBuilder_ == null) {
        referenceBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            io.github.yu31.protoc.pb.pbjson.TypeReference, io.github.yu31.protoc.pb.pbjson.TypeReference.Builder, io.github.yu31.protoc.pb.pbjson.TypeReferenceOrBuilder>(
                getReference(),
                getParentForChildren(),
                isClean());
        reference_ = null;
      }
      return referenceBuilder_;
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


    // @@protoc_insertion_point(builder_scope:json.FieldOptions)
  }

  // @@protoc_insertion_point(class_scope:json.FieldOptions)
  private static final io.github.yu31.protoc.pb.pbjson.FieldOptions DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.github.yu31.protoc.pb.pbjson.FieldOptions();
  }

  public static io.github.yu31.protoc.pb.pbjson.FieldOptions getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<FieldOptions>
      PARSER = new com.google.protobuf.AbstractParser<FieldOptions>() {
    @java.lang.Override
    public FieldOptions parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new FieldOptions(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<FieldOptions> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<FieldOptions> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.github.yu31.protoc.pb.pbjson.FieldOptions getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

