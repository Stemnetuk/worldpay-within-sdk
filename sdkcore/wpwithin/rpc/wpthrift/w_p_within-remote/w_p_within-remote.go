// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"wpthrift"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void setup(string name, string description)")
	fmt.Fprintln(os.Stderr, "  void addService(Service svc)")
	fmt.Fprintln(os.Stderr, "  void removeService(Service svc)")
	fmt.Fprintln(os.Stderr, "  void initHCE(HCECard hceCard)")
	fmt.Fprintln(os.Stderr, "  void initHTE(string merchantClientKey, string merchantServiceKey)")
	fmt.Fprintln(os.Stderr, "  void initConsumer(string scheme, string hostname, i32 port, string urlPrefix, string serviceId)")
	fmt.Fprintln(os.Stderr, "  void initProducer()")
	fmt.Fprintln(os.Stderr, "  Device getDevice()")
	fmt.Fprintln(os.Stderr, "  void startServiceBroadcast(i32 timeoutMillis)")
	fmt.Fprintln(os.Stderr, "  void stopServiceBroadcast()")
	fmt.Fprintln(os.Stderr, "   serviceDiscovery(i32 timeoutMillis)")
	fmt.Fprintln(os.Stderr, "   requestServices()")
	fmt.Fprintln(os.Stderr, "   getServicePrices(i32 serviceId)")
	fmt.Fprintln(os.Stderr, "  TotalPriceResponse selectService(i32 serviceId, i32 numberOfUnits, i32 priceId)")
	fmt.Fprintln(os.Stderr, "  PaymentResponse makePayment(TotalPriceResponse request)")
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
	client := wpthrift.NewWPWithinClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "setup":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Setup requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.Setup(value0, value1))
		fmt.Print("\n")
		break
	case "addService":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddService requires 1 args")
			flag.Usage()
		}
		arg37 := flag.Arg(1)
		mbTrans38 := thrift.NewTMemoryBufferLen(len(arg37))
		defer mbTrans38.Close()
		_, err39 := mbTrans38.WriteString(arg37)
		if err39 != nil {
			Usage()
			return
		}
		factory40 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt41 := factory40.GetProtocol(mbTrans38)
		argvalue0 := wpthrift.NewService()
		err42 := argvalue0.Read(jsProt41)
		if err42 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddService(value0))
		fmt.Print("\n")
		break
	case "removeService":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveService requires 1 args")
			flag.Usage()
		}
		arg43 := flag.Arg(1)
		mbTrans44 := thrift.NewTMemoryBufferLen(len(arg43))
		defer mbTrans44.Close()
		_, err45 := mbTrans44.WriteString(arg43)
		if err45 != nil {
			Usage()
			return
		}
		factory46 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt47 := factory46.GetProtocol(mbTrans44)
		argvalue0 := wpthrift.NewService()
		err48 := argvalue0.Read(jsProt47)
		if err48 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RemoveService(value0))
		fmt.Print("\n")
		break
	case "initHCE":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "InitHCE requires 1 args")
			flag.Usage()
		}
		arg49 := flag.Arg(1)
		mbTrans50 := thrift.NewTMemoryBufferLen(len(arg49))
		defer mbTrans50.Close()
		_, err51 := mbTrans50.WriteString(arg49)
		if err51 != nil {
			Usage()
			return
		}
		factory52 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt53 := factory52.GetProtocol(mbTrans50)
		argvalue0 := wpthrift.NewHCECard()
		err54 := argvalue0.Read(jsProt53)
		if err54 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.InitHCE(value0))
		fmt.Print("\n")
		break
	case "initHTE":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "InitHTE requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.InitHTE(value0, value1))
		fmt.Print("\n")
		break
	case "initConsumer":
		if flag.NArg()-1 != 5 {
			fmt.Fprintln(os.Stderr, "InitConsumer requires 5 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		tmp2, err59 := (strconv.Atoi(flag.Arg(3)))
		if err59 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		argvalue4 := flag.Arg(5)
		value4 := argvalue4
		fmt.Print(client.InitConsumer(value0, value1, value2, value3, value4))
		fmt.Print("\n")
		break
	case "initProducer":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "InitProducer requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.InitProducer())
		fmt.Print("\n")
		break
	case "getDevice":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetDevice requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetDevice())
		fmt.Print("\n")
		break
	case "startServiceBroadcast":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartServiceBroadcast requires 1 args")
			flag.Usage()
		}
		tmp0, err62 := (strconv.Atoi(flag.Arg(1)))
		if err62 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.StartServiceBroadcast(value0))
		fmt.Print("\n")
		break
	case "stopServiceBroadcast":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "StopServiceBroadcast requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.StopServiceBroadcast())
		fmt.Print("\n")
		break
	case "serviceDiscovery":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ServiceDiscovery requires 1 args")
			flag.Usage()
		}
		tmp0, err63 := (strconv.Atoi(flag.Arg(1)))
		if err63 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.ServiceDiscovery(value0))
		fmt.Print("\n")
		break
	case "requestServices":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "RequestServices requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.RequestServices())
		fmt.Print("\n")
		break
	case "getServicePrices":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetServicePrices requires 1 args")
			flag.Usage()
		}
		tmp0, err64 := (strconv.Atoi(flag.Arg(1)))
		if err64 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetServicePrices(value0))
		fmt.Print("\n")
		break
	case "selectService":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "SelectService requires 3 args")
			flag.Usage()
		}
		tmp0, err65 := (strconv.Atoi(flag.Arg(1)))
		if err65 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err66 := (strconv.Atoi(flag.Arg(2)))
		if err66 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		tmp2, err67 := (strconv.Atoi(flag.Arg(3)))
		if err67 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.SelectService(value0, value1, value2))
		fmt.Print("\n")
		break
	case "makePayment":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MakePayment requires 1 args")
			flag.Usage()
		}
		arg68 := flag.Arg(1)
		mbTrans69 := thrift.NewTMemoryBufferLen(len(arg68))
		defer mbTrans69.Close()
		_, err70 := mbTrans69.WriteString(arg68)
		if err70 != nil {
			Usage()
			return
		}
		factory71 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt72 := factory71.GetProtocol(mbTrans69)
		argvalue0 := wpthrift.NewTotalPriceResponse()
		err73 := argvalue0.Read(jsProt72)
		if err73 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MakePayment(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
