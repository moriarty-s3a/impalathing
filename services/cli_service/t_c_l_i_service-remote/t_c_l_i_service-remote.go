// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "github.com/chenjingping/impalathing/services/cli_service"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := cli_service.NewTCLIServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg58 := flag.Arg(1)
    mbTrans59 := thrift.NewTMemoryBufferLen(len(arg58))
    defer mbTrans59.Close()
    _, err60 := mbTrans59.WriteString(arg58)
    if err60 != nil {
      Usage()
      return
    }
    factory61 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt62 := factory61.GetProtocol(mbTrans59)
    argvalue0 := cli_service.NewTOpenSessionReq()
    err63 := argvalue0.Read(jsProt62)
    if err63 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg64 := flag.Arg(1)
    mbTrans65 := thrift.NewTMemoryBufferLen(len(arg64))
    defer mbTrans65.Close()
    _, err66 := mbTrans65.WriteString(arg64)
    if err66 != nil {
      Usage()
      return
    }
    factory67 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt68 := factory67.GetProtocol(mbTrans65)
    argvalue0 := cli_service.NewTCloseSessionReq()
    err69 := argvalue0.Read(jsProt68)
    if err69 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg70 := flag.Arg(1)
    mbTrans71 := thrift.NewTMemoryBufferLen(len(arg70))
    defer mbTrans71.Close()
    _, err72 := mbTrans71.WriteString(arg70)
    if err72 != nil {
      Usage()
      return
    }
    factory73 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt74 := factory73.GetProtocol(mbTrans71)
    argvalue0 := cli_service.NewTGetInfoReq()
    err75 := argvalue0.Read(jsProt74)
    if err75 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg76 := flag.Arg(1)
    mbTrans77 := thrift.NewTMemoryBufferLen(len(arg76))
    defer mbTrans77.Close()
    _, err78 := mbTrans77.WriteString(arg76)
    if err78 != nil {
      Usage()
      return
    }
    factory79 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt80 := factory79.GetProtocol(mbTrans77)
    argvalue0 := cli_service.NewTExecuteStatementReq()
    err81 := argvalue0.Read(jsProt80)
    if err81 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg82 := flag.Arg(1)
    mbTrans83 := thrift.NewTMemoryBufferLen(len(arg82))
    defer mbTrans83.Close()
    _, err84 := mbTrans83.WriteString(arg82)
    if err84 != nil {
      Usage()
      return
    }
    factory85 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt86 := factory85.GetProtocol(mbTrans83)
    argvalue0 := cli_service.NewTGetTypeInfoReq()
    err87 := argvalue0.Read(jsProt86)
    if err87 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg88 := flag.Arg(1)
    mbTrans89 := thrift.NewTMemoryBufferLen(len(arg88))
    defer mbTrans89.Close()
    _, err90 := mbTrans89.WriteString(arg88)
    if err90 != nil {
      Usage()
      return
    }
    factory91 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt92 := factory91.GetProtocol(mbTrans89)
    argvalue0 := cli_service.NewTGetCatalogsReq()
    err93 := argvalue0.Read(jsProt92)
    if err93 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg94 := flag.Arg(1)
    mbTrans95 := thrift.NewTMemoryBufferLen(len(arg94))
    defer mbTrans95.Close()
    _, err96 := mbTrans95.WriteString(arg94)
    if err96 != nil {
      Usage()
      return
    }
    factory97 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt98 := factory97.GetProtocol(mbTrans95)
    argvalue0 := cli_service.NewTGetSchemasReq()
    err99 := argvalue0.Read(jsProt98)
    if err99 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg100 := flag.Arg(1)
    mbTrans101 := thrift.NewTMemoryBufferLen(len(arg100))
    defer mbTrans101.Close()
    _, err102 := mbTrans101.WriteString(arg100)
    if err102 != nil {
      Usage()
      return
    }
    factory103 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt104 := factory103.GetProtocol(mbTrans101)
    argvalue0 := cli_service.NewTGetTablesReq()
    err105 := argvalue0.Read(jsProt104)
    if err105 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg106 := flag.Arg(1)
    mbTrans107 := thrift.NewTMemoryBufferLen(len(arg106))
    defer mbTrans107.Close()
    _, err108 := mbTrans107.WriteString(arg106)
    if err108 != nil {
      Usage()
      return
    }
    factory109 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt110 := factory109.GetProtocol(mbTrans107)
    argvalue0 := cli_service.NewTGetTableTypesReq()
    err111 := argvalue0.Read(jsProt110)
    if err111 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg112 := flag.Arg(1)
    mbTrans113 := thrift.NewTMemoryBufferLen(len(arg112))
    defer mbTrans113.Close()
    _, err114 := mbTrans113.WriteString(arg112)
    if err114 != nil {
      Usage()
      return
    }
    factory115 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt116 := factory115.GetProtocol(mbTrans113)
    argvalue0 := cli_service.NewTGetColumnsReq()
    err117 := argvalue0.Read(jsProt116)
    if err117 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg118 := flag.Arg(1)
    mbTrans119 := thrift.NewTMemoryBufferLen(len(arg118))
    defer mbTrans119.Close()
    _, err120 := mbTrans119.WriteString(arg118)
    if err120 != nil {
      Usage()
      return
    }
    factory121 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt122 := factory121.GetProtocol(mbTrans119)
    argvalue0 := cli_service.NewTGetFunctionsReq()
    err123 := argvalue0.Read(jsProt122)
    if err123 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg124 := flag.Arg(1)
    mbTrans125 := thrift.NewTMemoryBufferLen(len(arg124))
    defer mbTrans125.Close()
    _, err126 := mbTrans125.WriteString(arg124)
    if err126 != nil {
      Usage()
      return
    }
    factory127 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt128 := factory127.GetProtocol(mbTrans125)
    argvalue0 := cli_service.NewTGetOperationStatusReq()
    err129 := argvalue0.Read(jsProt128)
    if err129 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg130 := flag.Arg(1)
    mbTrans131 := thrift.NewTMemoryBufferLen(len(arg130))
    defer mbTrans131.Close()
    _, err132 := mbTrans131.WriteString(arg130)
    if err132 != nil {
      Usage()
      return
    }
    factory133 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt134 := factory133.GetProtocol(mbTrans131)
    argvalue0 := cli_service.NewTCancelOperationReq()
    err135 := argvalue0.Read(jsProt134)
    if err135 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg136 := flag.Arg(1)
    mbTrans137 := thrift.NewTMemoryBufferLen(len(arg136))
    defer mbTrans137.Close()
    _, err138 := mbTrans137.WriteString(arg136)
    if err138 != nil {
      Usage()
      return
    }
    factory139 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt140 := factory139.GetProtocol(mbTrans137)
    argvalue0 := cli_service.NewTCloseOperationReq()
    err141 := argvalue0.Read(jsProt140)
    if err141 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg142 := flag.Arg(1)
    mbTrans143 := thrift.NewTMemoryBufferLen(len(arg142))
    defer mbTrans143.Close()
    _, err144 := mbTrans143.WriteString(arg142)
    if err144 != nil {
      Usage()
      return
    }
    factory145 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt146 := factory145.GetProtocol(mbTrans143)
    argvalue0 := cli_service.NewTGetResultSetMetadataReq()
    err147 := argvalue0.Read(jsProt146)
    if err147 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg148 := flag.Arg(1)
    mbTrans149 := thrift.NewTMemoryBufferLen(len(arg148))
    defer mbTrans149.Close()
    _, err150 := mbTrans149.WriteString(arg148)
    if err150 != nil {
      Usage()
      return
    }
    factory151 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt152 := factory151.GetProtocol(mbTrans149)
    argvalue0 := cli_service.NewTFetchResultsReq()
    err153 := argvalue0.Read(jsProt152)
    if err153 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
