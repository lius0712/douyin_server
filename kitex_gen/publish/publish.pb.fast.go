// Code generated by FastPB v0.0.1. DO NOT EDIT.

package publish

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	feed "github.com/lius0712/douyin_server/kitex_gen/feed"
)

func (x *PublishActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishActionRequest[number], err)
}

func (x *PublishActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.File, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *PublishActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishActionResponse[number], err)
}

func (x *PublishActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PublishActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Msg = &tmp
	return offset, err
}

func (x *PublishListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListRequest[number], err)
}

func (x *PublishListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListResponse[number], err)
}

func (x *PublishListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PublishListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Msg = &tmp
	return offset, err
}

func (x *PublishListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v feed.Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *PublishActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishActionRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.File) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 1, x.File)
	return offset
}

func (x *PublishActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Token)
	return offset
}

func (x *PublishActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Title)
	return offset
}

func (x *PublishActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PublishActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.Status)
	return offset
}

func (x *PublishActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.Msg)
	return offset
}

func (x *PublishListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PublishListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *PublishListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Token)
	return offset
}

func (x *PublishListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.Status)
	return offset
}

func (x *PublishListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, *x.Msg)
	return offset
}

func (x *PublishListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *PublishActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishActionRequest) sizeField1() (n int) {
	if len(x.File) == 0 {
		return n
	}
	n += fastpb.SizeBytes(1, x.File)
	return n
}

func (x *PublishActionRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Token)
	return n
}

func (x *PublishActionRequest) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Title)
	return n
}

func (x *PublishActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PublishActionResponse) sizeField1() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.Status)
	return n
}

func (x *PublishActionResponse) sizeField2() (n int) {
	if x.Msg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.Msg)
	return n
}

func (x *PublishListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PublishListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *PublishListRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Token)
	return n
}

func (x *PublishListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishListResponse) sizeField1() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.Status)
	return n
}

func (x *PublishListResponse) sizeField2() (n int) {
	if x.Msg == nil {
		return n
	}
	n += fastpb.SizeString(2, *x.Msg)
	return n
}

func (x *PublishListResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

var fieldIDToName_PublishActionRequest = map[int32]string{
	1: "File",
	2: "Token",
	3: "Title",
}

var fieldIDToName_PublishActionResponse = map[int32]string{
	1: "Status",
	2: "Msg",
}

var fieldIDToName_PublishListRequest = map[int32]string{
	1: "UserId",
	2: "Token",
}

var fieldIDToName_PublishListResponse = map[int32]string{
	1: "Status",
	2: "Msg",
	3: "VideoList",
}

var _ = feed.File_feed_proto