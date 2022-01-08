package log

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type simpleEncoder struct {
	zapcore.Encoder
	pool buffer.Pool
}

// nolint
func SimpleEncoder(cfg zapcore.EncoderConfig) *simpleEncoder {
	return &simpleEncoder{Encoder: zapcore.NewJSONEncoder(cfg), pool: buffer.NewPool()}
}

func (e *simpleEncoder) Clone() zapcore.Encoder {
	return &simpleEncoder{Encoder: e.Encoder.Clone(), pool: e.pool}
}

// nolint
func (e *simpleEncoder) EncodeEntry(entry zapcore.Entry, _ []zapcore.Field) (*buffer.Buffer, error) {
	ret := e.pool.Get()
	m := levelToColor[entry.Level].Add(entry.Message)
	ret.AppendString(m)
	ret.AppendByte('\n')
	return ret, nil
}
