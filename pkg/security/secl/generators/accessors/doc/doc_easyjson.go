// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package doc

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(in *jlexer.Lexer, out *Documentation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "secl":
			if in.IsNull() {
				in.Skip()
				out.Kinds = nil
			} else {
				in.Delim('[')
				if out.Kinds == nil {
					if !in.IsDelim(']') {
						out.Kinds = make([]DocEventKind, 0, 0)
					} else {
						out.Kinds = []DocEventKind{}
					}
				} else {
					out.Kinds = (out.Kinds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 DocEventKind
					(v1).UnmarshalEasyJSON(in)
					out.Kinds = append(out.Kinds, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(out *jwriter.Writer, in Documentation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"secl\":"
		out.RawString(prefix[1:])
		if in.Kinds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Kinds {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Documentation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Documentation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Documentation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Documentation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc(l, v)
}
func easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(in *jlexer.Lexer, out *DocEventProperty) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "definition":
			out.Doc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(out *jwriter.Writer, in DocEventProperty) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"definition\":"
		out.RawString(prefix)
		out.String(string(in.Doc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocEventProperty) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocEventProperty) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocEventProperty) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocEventProperty) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc1(l, v)
}
func easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(in *jlexer.Lexer, out *DocEventKind) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "definition":
			out.Definition = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "from_agent_version":
			out.FromAgentVersion = string(in.String())
		case "properties":
			if in.IsNull() {
				in.Skip()
				out.Properties = nil
			} else {
				in.Delim('[')
				if out.Properties == nil {
					if !in.IsDelim(']') {
						out.Properties = make([]DocEventProperty, 0, 1)
					} else {
						out.Properties = []DocEventProperty{}
					}
				} else {
					out.Properties = (out.Properties)[:0]
				}
				for !in.IsDelim(']') {
					var v4 DocEventProperty
					(v4).UnmarshalEasyJSON(in)
					out.Properties = append(out.Properties, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(out *jwriter.Writer, in DocEventKind) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"definition\":"
		out.RawString(prefix)
		out.String(string(in.Definition))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"from_agent_version\":"
		out.RawString(prefix)
		out.String(string(in.FromAgentVersion))
	}
	{
		const prefix string = ",\"properties\":"
		out.RawString(prefix)
		if in.Properties == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Properties {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocEventKind) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocEventKind) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9972aa54EncodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocEventKind) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocEventKind) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9972aa54DecodeGithubComDataDogDatadogAgentPkgSecuritySeclGeneratorsAccessorsDoc2(l, v)
}
