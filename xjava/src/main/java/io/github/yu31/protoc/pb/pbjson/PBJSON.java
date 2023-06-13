// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: json.proto

package io.github.yu31.protoc.pb.pbjson;

public final class PBJSON {
  private PBJSON() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
    registry.add(io.github.yu31.protoc.pb.pbjson.PBJSON.message);
    registry.add(io.github.yu31.protoc.pb.pbjson.PBJSON.field);
    registry.add(io.github.yu31.protoc.pb.pbjson.PBJSON.oneof);
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public static final int MESSAGE_FIELD_NUMBER = 64032;
  /**
   * <code>extend .google.protobuf.MessageOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.MessageOptions,
      io.github.yu31.protoc.pb.pbjson.MessageOptions> message = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        io.github.yu31.protoc.pb.pbjson.MessageOptions.class,
        io.github.yu31.protoc.pb.pbjson.MessageOptions.getDefaultInstance());
  public static final int FIELD_FIELD_NUMBER = 64033;
  /**
   * <code>extend .google.protobuf.FieldOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.FieldOptions,
      io.github.yu31.protoc.pb.pbjson.FieldOptions> field = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        io.github.yu31.protoc.pb.pbjson.FieldOptions.class,
        io.github.yu31.protoc.pb.pbjson.FieldOptions.getDefaultInstance());
  public static final int ONEOF_FIELD_NUMBER = 64034;
  /**
   * <code>extend .google.protobuf.OneofOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.OneofOptions,
      io.github.yu31.protoc.pb.pbjson.OneofOptions> oneof = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        io.github.yu31.protoc.pb.pbjson.OneofOptions.class,
        io.github.yu31.protoc.pb.pbjson.OneofOptions.getDefaultInstance());
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_MessageOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_MessageOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_OneofOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_OneofOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_FieldOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_FieldOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_TypeEnum_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_TypeEnum_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_TypeAny_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_TypeAny_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_TypeDuration_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_TypeDuration_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_TypeTimestamp_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_TypeTimestamp_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_json_TypeTimestamp_Layout_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_json_TypeTimestamp_Layout_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\njson.proto\022\004json\032 google/protobuf/desc" +
      "riptor.proto\"A\n\016MessageOptions\022\016\n\006ignore" +
      "\030\001 \001(\010\022\037\n\027disallow_unknown_fields\030\002 \001(\010\"" +
      "[\n\014OneofOptions\022\021\n\004json\030\001 \001(\tH\000\210\001\001\022\016\n\006ig" +
      "nore\030\002 \001(\010\022\021\n\tomitempty\030\003 \001(\010\022\014\n\004hide\030\004 " +
      "\001(\010B\007\n\005_json\"\201\002\n\014FieldOptions\022\021\n\004json\030\001 " +
      "\001(\tH\001\210\001\001\022\016\n\006ignore\030\002 \001(\010\022\021\n\tomitempty\030\003 " +
      "\001(\010\022\027\n\017use_enum_string\030\004 \001(\010\022\036\n\004enum\030\013 \001" +
      "(\0132\016.json.TypeEnumH\000\022\034\n\003any\030\014 \001(\0132\r.json" +
      ".TypeAnyH\000\022&\n\010duration\030\r \001(\0132\022.json.Type" +
      "DurationH\000\022(\n\ttimestamp\030\016 \001(\0132\023.json.Typ" +
      "eTimestampH\000B\t\n\007TypeSetB\007\n\005_json\"S\n\010Type" +
      "Enum\022%\n\006format\030\001 \001(\0162\025.json.TypeEnum.For" +
      "mat\" \n\006Format\022\n\n\006Number\020\000\022\n\n\006String\020\001\"P\n" +
      "\007TypeAny\022$\n\006format\030\001 \001(\0162\024.json.TypeAny." +
      "Format\"\037\n\006Format\022\n\n\006Native\020\000\022\t\n\005Proto\020\001\"" +
      "\265\001\n\014TypeDuration\022)\n\006format\030\001 \001(\0162\031.json." +
      "TypeDuration.Format\"z\n\006Format\022\n\n\006Native\020" +
      "\000\022\n\n\006String\020\001\022\017\n\013Nanoseconds\020\002\022\020\n\014Micros" +
      "econds\020\003\022\020\n\014Milliseconds\020\004\022\013\n\007Seconds\020\005\022" +
      "\013\n\007Minutes\020\006\022\t\n\005Hours\020\007\"\214\002\n\rTypeTimestam" +
      "p\022*\n\006format\030\001 \001(\0162\032.json.TypeTimestamp.F" +
      "ormat\022*\n\006layout\030\002 \001(\0132\032.json.TypeTimesta" +
      "mp.Layout\032D\n\006Layout\022\016\n\006golang\030\001 \001(\t\022\014\n\004j" +
      "ava\030\002 \001(\t\022\014\n\004rust\030\003 \001(\t\022\016\n\006python\030\004 \001(\t\"" +
      "]\n\006Format\022\n\n\006Native\020\000\022\016\n\nTimeLayout\020\001\022\014\n" +
      "\010UnixNano\020\002\022\r\n\tUnixMicro\020\003\022\r\n\tUnixMilli\020" +
      "\004\022\013\n\007UnixSec\020\005:H\n\007message\022\037.google.proto" +
      "buf.MessageOptions\030\240\364\003 \001(\0132\024.json.Messag" +
      "eOptions:B\n\005field\022\035.google.protobuf.Fiel" +
      "dOptions\030\241\364\003 \001(\0132\022.json.FieldOptions:B\n\005" +
      "oneof\022\035.google.protobuf.OneofOptions\030\242\364\003" +
      " \001(\0132\022.json.OneofOptionsB]\n\037io.github.yu" +
      "31.protoc.pb.pbjsonB\006PBJSONP\001Z0github.co" +
      "m/yu31/protoc-plugin-json/xgo/pb/pbjsonb" +
      "\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.DescriptorProtos.getDescriptor(),
        });
    internal_static_json_MessageOptions_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_json_MessageOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_MessageOptions_descriptor,
        new java.lang.String[] { "Ignore", "DisallowUnknownFields", });
    internal_static_json_OneofOptions_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_json_OneofOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_OneofOptions_descriptor,
        new java.lang.String[] { "Json", "Ignore", "Omitempty", "Hide", "Json", });
    internal_static_json_FieldOptions_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_json_FieldOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_FieldOptions_descriptor,
        new java.lang.String[] { "Json", "Ignore", "Omitempty", "UseEnumString", "Enum", "Any", "Duration", "Timestamp", "TypeSet", "Json", });
    internal_static_json_TypeEnum_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_json_TypeEnum_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_TypeEnum_descriptor,
        new java.lang.String[] { "Format", });
    internal_static_json_TypeAny_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_json_TypeAny_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_TypeAny_descriptor,
        new java.lang.String[] { "Format", });
    internal_static_json_TypeDuration_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_json_TypeDuration_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_TypeDuration_descriptor,
        new java.lang.String[] { "Format", });
    internal_static_json_TypeTimestamp_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_json_TypeTimestamp_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_TypeTimestamp_descriptor,
        new java.lang.String[] { "Format", "Layout", });
    internal_static_json_TypeTimestamp_Layout_descriptor =
      internal_static_json_TypeTimestamp_descriptor.getNestedTypes().get(0);
    internal_static_json_TypeTimestamp_Layout_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_json_TypeTimestamp_Layout_descriptor,
        new java.lang.String[] { "Golang", "Java", "Rust", "Python", });
    message.internalInit(descriptor.getExtensions().get(0));
    field.internalInit(descriptor.getExtensions().get(1));
    oneof.internalInit(descriptor.getExtensions().get(2));
    com.google.protobuf.DescriptorProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
