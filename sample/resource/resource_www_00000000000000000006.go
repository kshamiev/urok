// This file generated by go generate; DO NOT EDIT.
// НЕ РЕДАКТИРОВАТЬ! Изменения будут перезаписаны при следующей генерации.

package resource

import "time"

func init() {
	_ = Get().Add(
		"www",
		"profile_users.html",
		Resource{
			Size: 943,
			Time: func() time.Time {
				t, _ := time.ParseInLocation(time.RFC3339Nano, "2024-12-14T16:54:12.235166362+03:00", time.Local)
				return t
			}(),
			ContentType: "text/html; charset=utf-8",
			Content: []byte{
				0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61,
				0x6e, 0x67, 0x3d, 0x22, 0x72, 0x75, 0x22, 0x3e, 0x0a, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74,
				0x61, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x65, 0x71, 0x75, 0x69, 0x76, 0x3d, 0x22, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79,
				0x70, 0x65, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3b, 0x20,
				0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74,
				0x61, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x65, 0x71, 0x75, 0x69, 0x76, 0x3d, 0x22, 0x58, 0x2d, 0x55, 0x41, 0x2d, 0x43, 0x6f, 0x6d, 0x70, 0x61,
				0x74, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x49, 0x45, 0x3d, 0x65, 0x64, 0x67, 0x65, 0x22,
				0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x76, 0x69, 0x65, 0x77, 0x70, 0x6f,
				0x72, 0x74, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x64, 0x65, 0x76, 0x69, 0x63,
				0x65, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2c, 0x20, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x31,
				0x2e, 0x30, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x68, 0x74,
				0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6a, 0x73, 0x2f, 0x74, 0x65,
				0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x61, 0x70, 0x70, 0x2e, 0x6a, 0x73, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72,
				0x69, 0x70, 0x74, 0x3e, 0x20, 0x3c, 0x21, 0x2d, 0x2d, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xb4, 0xd0, 0xba, 0xd0, 0xbb, 0xd1, 0x8e, 0xd1, 0x87, 0xd0,
				0xb0, 0xd0, 0xb5, 0xd0, 0xbc, 0x20, 0xd1, 0x81, 0xd0, 0xba, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbf, 0xd1, 0x82, 0x20, 0xd0, 0xbe, 0xd1, 0x82, 0x20,
				0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd0, 0xb5, 0xd0, 0xb3, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xbc, 0x2d, 0x2d, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20,
				0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x68,
				0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20,
				0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69,
				0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x20, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x20, 0x68,
				0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x69, 0x63, 0x6f, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d,
				0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x78, 0x2d, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74,
				0x6c, 0x65, 0x3e, 0xd0, 0x92, 0xd1, 0x81, 0xd0, 0xb5, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd0, 0xbe, 0xd0, 0xb2,
				0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd0, 0xb8, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x65, 0x61,
				0x64, 0x3e, 0x0a, 0x0a, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x68, 0x31, 0x3e, 0xd0, 0x92, 0xd1, 0x81, 0xd0, 0xb5, 0x20, 0xd0, 0xbf,
				0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd0, 0xb8, 0x3c, 0x2f,
				0x68, 0x31, 0x3e, 0x0a, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x75, 0x73, 0x65, 0x72, 0x63, 0x61, 0x72, 0x64, 0x22, 0x3e, 0x3c,
				0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x3c, 0x70, 0x3e, 0xd0, 0xa0, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87, 0xd0, 0xb8, 0xd0, 0xb9, 0x20,
				0xd0, 0xbe, 0xd0, 0xb1, 0xd1, 0x8a, 0xd0, 0xb5, 0xd0, 0xba, 0xd1, 0x82, 0x3a, 0x20, 0x7b, 0x7b, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
				0x44, 0x7d, 0x7d, 0x3c, 0x2f, 0x70, 0x3e, 0x0a, 0x0a, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74,
				0x65, 0x78, 0x74, 0x2f, 0x6a, 0x61, 0x76, 0x61, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
				0x3d, 0x22, 0x4a, 0x61, 0x76, 0x61, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74,
				0x69, 0x63, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x6a, 0x73, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x0a,
				0x3c, 0x21, 0x2d, 0x2d, 0x20, 0x45, 0x72, 0x75, 0x64, 0x61, 0x20, 0x69, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x20, 0x66, 0x6f,
				0x72, 0x20, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x20, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x73, 0x20, 0x2d, 0x2d, 0x3e, 0x0a, 0x3c, 0x73,
				0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x64, 0x6e, 0x2e, 0x6a,
				0x73, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x6e, 0x70, 0x6d, 0x2f, 0x65, 0x72, 0x75, 0x64, 0x61, 0x22, 0x3e, 0x3c,
				0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x65, 0x72, 0x75, 0x64, 0x61, 0x2e, 0x69,
				0x6e, 0x69, 0x74, 0x28, 0x29, 0x3b, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a,
				0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e,
			},
		})
}
