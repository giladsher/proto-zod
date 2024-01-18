package main

import (
	"fmt"
	"math"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	Double          = "DoubleSchema"
	Float           = "FloatSchema"
	Int32           = "Int32Schema"
	Int64           = "Int64Schema"
	Uint32          = "Uint32Schema"
	Uint64          = "Uint64Schema"
	Sint32          = "Sint32Schema"
	Sint64          = "Sint64Schema"
	Fixed32         = "Fixed32Schema"
	Fixed64         = "Fixed64Schema"
	Sfixed32        = "Sfixed32Schema"
	Sfixed64        = "Sfixed64Schema"
	Boolean         = "BooleanSchema"
	String          = "StringSchema"
	Bytes           = "BytesSchema"
	Enum            = "EnumSchema"
	Timestamp       = "TimestampSchema"
	Any             = "AnySchema"
	Struct          = "StructSchema"
	ZodImport       = "import { z } from 'zod'"
	ZodNumber       = "number()"
	ZodString       = "string()"
	ZodBoolean      = "boolean()"
	ZodNativeEnum   = "nativeEnum(%s)"
	ZodDate         = "date()"
	ZodUndefined    = "undefined()"
	ZodNull         = "null()"
	ZodAny          = "any()"
	ZodUnknown      = "unknown()"
	ZodGreaterThan  = "gt(%d)"
	ZodLessThan     = "lt(%d)"
	ZodMin          = "min(%d)"
	ZodMax          = "max(%d)"
	ZodArray        = "array()"
	maxDoubleNumber = 9007199254740992
	maxStringLength = 232
	maxByteLength   = 232
)

type (
	ZodSchemaGenerator interface {
		ToString() string
		New(name string, zodType string, constraints []string) ZodSchema
	}

	ZodSchema struct {
		Name        string
		ZodType     string
		Constraints []string
	}
)

func NewZodSchema(name string, zodType string, constraints []string) ZodSchema {
	var zs ZodSchema
	return zs.New(name, zodType, constraints)
}

func (zs *ZodSchema) ToString() string {
	var constraints string
	if len(zs.Constraints) > 0 {
		constraints = fmt.Sprintf(".%v", strings.Join(zs.Constraints, "."))
	}
	// TODO: handle special enum case
	return fmt.Sprintf("export const %v = z.%v%v;\n", zs.Name, zs.ZodType, constraints)
}

func (zs *ZodSchema) New(name string, zodType string, constraints []string) ZodSchema {
	switch name {
	case Double:
		return ZodSchema{
			Name:        Double,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMax, maxDoubleNumber)},
		}
	case Float:
		return ZodSchema{
			Name:        Float,
			ZodType:     ZodNumber,
			Constraints: []string{},
		}
	case Int32:
		return ZodSchema{
			Name:        Int32,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt32), fmt.Sprintf(ZodMax, math.MaxInt32)},
		}

	case Int64:
		return ZodSchema{
			Name:        Int64,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt64), fmt.Sprintf(ZodMax, math.MaxInt64)},
		}

	case Uint32:
		return ZodSchema{
			Name:        Uint32,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, 0), fmt.Sprintf(ZodMax, math.MaxUint32)},
		}

	case Uint64:
		return ZodSchema{
			Name:        Uint64,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, 0), fmt.Sprintf(ZodMax, uint64(math.MaxUint64))},
		}

	case Sint32:
		return ZodSchema{
			Name:        Sint32,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt32), fmt.Sprintf(ZodMax, math.MaxInt32)},
		}

	case Sint64:
		return ZodSchema{
			Name:        Sint64,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt64), fmt.Sprintf(ZodMax, math.MaxInt64)},
		}

	case Fixed32:
		return ZodSchema{
			Name:        Fixed32,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, 0), fmt.Sprintf(ZodMax, math.MaxUint32)},
		}

	case Fixed64:
		return ZodSchema{
			Name:        Fixed64,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, 0), fmt.Sprintf(ZodMax, uint64(math.MaxUint64))},
		}

	case Sfixed32:
		return ZodSchema{
			Name:        Sfixed32,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt32), fmt.Sprintf(ZodMax, math.MaxInt32)},
		}

	case Sfixed64:
		return ZodSchema{
			Name:        Sfixed64,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MinInt64), fmt.Sprintf(ZodMax, math.MaxInt64)},
		}

	case Boolean:
		return ZodSchema{
			Name:        Boolean,
			ZodType:     ZodBoolean,
			Constraints: []string{},
		}

	case String:
		return ZodSchema{
			Name:        String,
			ZodType:     ZodString,
			Constraints: []string{fmt.Sprintf(ZodMax, maxStringLength)},
		}

	case Bytes:
		return ZodSchema{
			Name:        Bytes,
			ZodType:     ZodNumber,
			Constraints: []string{fmt.Sprintf(ZodMin, math.MaxUint8), ZodArray, ZodArray},
		}

	default:
		panic(fmt.Sprintf("Could not match name %s", name))
	}
}

func generateZodSchema(g *protogen.GeneratedFile, zs ZodSchemaGenerator) {
	g.P(zs.ToString())
	g.P()
}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_proto_zod.pb.ts"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by proto-zod. DO NOT EDIT.")
	g.P()
	g.P(ZodImport)
	g.P()
	// var zs ZodSchema
	// zs.New()

	return g
}
