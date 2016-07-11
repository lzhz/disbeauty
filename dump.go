package disbeauty

import (
	"adserver/storage"
	"flag"
	"fmt"
	"strings"
	"tripod/zconf"
)

var Cfg struct {
	PcIndex string
	MbIndex string
	VdIndex string
}

var infos map[string]interface{}
var indexType string
var cmd string
var interactive bool

func init() {
	cfgPath := ""
	flag.StringVar(&cfgPath, "cfg", "./cfg.yaml", "config for dump")
	flag.StringVar(&indexType, "type", "pc", "index type")
	flag.BoolVar(&interactive, "i", false, "interactive mode")
	flag.StringVar(&cmd, "cmd", "", "K-/V- Index1>Index2>Index3...")
	flag.Parse()
	infos = make(map[string]interface{}, 20)
	fmt.Printf("cfgPath;%v\n", cfgPath)
	err := zconf.ParseYaml(cfgPath, &Cfg)
	fmt.Printf("cfg:%v, err:%v\n", Cfg, err)
}

func dump(index *storage.Index) {
	args := strings.Split(cmd, "-")
	if len(args) != 2 {
		fmt.Println("Invalid Cmd: ", cmd)
	}

	keylist := strings.TrimSpace(args[1])
	keys := strings.Split(keylist, ">")
	v := GetInValue(index, keys)

	if args[0] == "V" {
		fmt.Println(DumpString(v))
	} else if args[0] == "K" {
		fmt.Println(DumpSimple(v))
	}
}

func main() {
	var index *storage.Index = nil
	if indexType == "pc" {
		index = storage.NewIndex(Cfg.PcIndex, Cfg.PcIndex+"/incremental")
	} else if indexType == "mb" {
		index = storage.NewIndex(Cfg.MbIndex, Cfg.MbIndex+"/incremental")
	} else {
		index = storage.NewIndex(Cfg.VdIndex, Cfg.VdIndex+"/incremental")
	}

	if cmd != "" {
		dump(index)
	}

	if interactive {
		for {
			infos = make(map[string]interface{}, 20)
			fmt.Print("InputCmd:")
			_, err := fmt.Scanln(&cmd)
			if err != nil {
				fmt.Println("Get Cmd Fail:", err)
				break
			}

			dump(index)
			fmt.Println("")
		}
	}
}
