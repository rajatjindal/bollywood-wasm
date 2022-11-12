// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

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

func easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(in *jlexer.Lexer, out *NewGame) {
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
		case "createdAt":
			out.CreatedAt = int64(in.Int64())
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
func easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(out *jwriter.Writer, in NewGame) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CreatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NewGame) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NewGame) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewGame) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NewGame) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes(l, v)
}
func easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(in *jlexer.Lexer, out *Meta) {
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
		case "RawMovieName":
			out.RawMovieName = string(in.String())
		case "NormalizedMovieName":
			out.NormalizedMovieName = string(in.String())
		case "MovieNameMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.MovieNameMap = make(map[int32]GuessResult)
				for !in.IsDelim('}') {
					key := int32(in.Int32Str())
					in.WantColon()
					var v1 GuessResult
					v1 = GuessResult(in.Int())
					(out.MovieNameMap)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(out *jwriter.Writer, in Meta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"RawMovieName\":"
		out.RawString(prefix[1:])
		out.String(string(in.RawMovieName))
	}
	{
		const prefix string = ",\"NormalizedMovieName\":"
		out.RawString(prefix)
		out.String(string(in.NormalizedMovieName))
	}
	{
		const prefix string = ",\"MovieNameMap\":"
		out.RawString(prefix)
		if in.MovieNameMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.MovieNameMap {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.Int32Str(int32(v2Name))
				out.RawByte(':')
				out.Int(int(v2Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Meta) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Meta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Meta) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Meta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes1(l, v)
}
func easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(in *jlexer.Lexer, out *Guess) {
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
		case "id":
			out.Id = string(in.String())
		case "key":
			out.Key = string(in.String())
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
func easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(out *jwriter.Writer, in Guess) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix)
		out.String(string(in.Key))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Guess) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Guess) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Guess) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Guess) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes2(l, v)
}
func easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(in *jlexer.Lexer, out *GameWithMeta) {
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
		case "game":
			if in.IsNull() {
				in.Skip()
				out.Game = nil
			} else {
				if out.Game == nil {
					out.Game = new(Game)
				}
				(*out.Game).UnmarshalEasyJSON(in)
			}
		case "meta":
			if in.IsNull() {
				in.Skip()
				out.Meta = nil
			} else {
				if out.Meta == nil {
					out.Meta = new(Meta)
				}
				(*out.Meta).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(out *jwriter.Writer, in GameWithMeta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"game\":"
		out.RawString(prefix[1:])
		if in.Game == nil {
			out.RawString("null")
		} else {
			(*in.Game).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"meta\":"
		out.RawString(prefix)
		if in.Meta == nil {
			out.RawString("null")
		} else {
			(*in.Meta).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameWithMeta) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameWithMeta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameWithMeta) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameWithMeta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes3(l, v)
}
func easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(in *jlexer.Lexer, out *Game) {
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
		case "id":
			out.Id = string(in.String())
		case "guessMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.GuessMap = make(map[int32]GuessResult)
				for !in.IsDelim('}') {
					key := int32(in.Int32Str())
					in.WantColon()
					var v3 GuessResult
					v3 = GuessResult(in.Int())
					(out.GuessMap)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "remainingGuesses":
			out.RemainingGuesses = int(in.Int())
		case "guessedSuccessfully":
			out.GuessedSuccessfully = bool(in.Bool())
		case "movieName":
			out.MovieName = string(in.String())
		case "plot":
			out.Plot = string(in.String())
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
func easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(out *jwriter.Writer, in Game) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"guessMap\":"
		out.RawString(prefix)
		if in.GuessMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.GuessMap {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.Int32Str(int32(v4Name))
				out.RawByte(':')
				out.Int(int(v4Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"remainingGuesses\":"
		out.RawString(prefix)
		out.Int(int(in.RemainingGuesses))
	}
	{
		const prefix string = ",\"guessedSuccessfully\":"
		out.RawString(prefix)
		out.Bool(bool(in.GuessedSuccessfully))
	}
	{
		const prefix string = ",\"movieName\":"
		out.RawString(prefix)
		out.String(string(in.MovieName))
	}
	{
		const prefix string = ",\"plot\":"
		out.RawString(prefix)
		out.String(string(in.Plot))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Game) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Game) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Game) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Game) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComRajatjindalBollywoodWasmBapiPkgTypes4(l, v)
}
