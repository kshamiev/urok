// This file generated by go generate; DO NOT EDIT.
// НЕ РЕДАКТИРОВАТЬ! Изменения будут перезаписаны при следующей генерации.

package resource

import "time"

func init() {
	_ = Get().Add(
		"www",
		"test.html",
		Resource{
			Size: 160,
			Time: func() time.Time {
				t, _ := time.ParseInLocation(time.RFC3339Nano, "2025-02-06T20:51:10.526174963+03:00", time.Local)
				return t
			}(),
			ContentType: "text/html; charset=utf-8",
			Content: []byte{
				0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61,
				0x6e, 0x67, 0x3d, 0x22, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74,
				0x61, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x22, 0x55, 0x54, 0x46, 0x2d, 0x38, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c,
				0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x65,
				0x61, 0x64, 0x3e, 0x0a, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x68, 0x31, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x73, 0x74,
				0x0a, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0x0a, 0x3c, 0x70, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x3c, 0x2f, 0x70, 0x3e,
				0x0a, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e,
			},
		})
}
