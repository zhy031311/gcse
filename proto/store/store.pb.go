// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/store/store.proto
// DO NOT EDIT!

/*
Package stpb is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/store/store.proto

It has these top-level messages:
	PackageInfo
	PersonInfo
	Repository
*/
package stpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gcse_spider "github.com/daviddengcn/gcse/proto/spider"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PackageInfo struct {
	Name         string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Package      string                    `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	Author       string                    `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Stars        int32                     `protobuf:"varint,4,opt,name=stars" json:"stars,omitempty"`
	Synopsis     string                    `protobuf:"bytes,5,opt,name=synopsis" json:"synopsis,omitempty"`
	Description  string                    `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	ProjectUrl   string                    `protobuf:"bytes,7,opt,name=project_url,json=projectUrl" json:"project_url,omitempty"`
	ReadmeFn     string                    `protobuf:"bytes,8,opt,name=readme_fn,json=readmeFn" json:"readme_fn,omitempty"`
	ReadmeData   string                    `protobuf:"bytes,9,opt,name=readme_data,json=readmeData" json:"readme_data,omitempty"`
	Imports      []string                  `protobuf:"bytes,10,rep,name=imports" json:"imports,omitempty"`
	TestImports  []string                  `protobuf:"bytes,11,rep,name=test_imports,json=testImports" json:"test_imports,omitempty"`
	Exported     []string                  `protobuf:"bytes,12,rep,name=exported" json:"exported,omitempty"`
	References   []string                  `protobuf:"bytes,18,rep,name=references" json:"references,omitempty"`
	CrawlingInfo *gcse_spider.CrawlingInfo `protobuf:"bytes,17,opt,name=crawling_info,json=crawlingInfo" json:"crawling_info,omitempty"`
	// Available if the package is not the repo's root.
	FolderInfo *gcse_spider.FolderInfo `protobuf:"bytes,14,opt,name=folder_info,json=folderInfo" json:"folder_info,omitempty"`
	// Available if the package is the repo's root.
	RepoInfo *gcse_spider.RepoInfo `protobuf:"bytes,15,opt,name=repo_info,json=repoInfo" json:"repo_info,omitempty"`
}

func (m *PackageInfo) Reset()                    { *m = PackageInfo{} }
func (m *PackageInfo) String() string            { return proto.CompactTextString(m) }
func (*PackageInfo) ProtoMessage()               {}
func (*PackageInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PackageInfo) GetCrawlingInfo() *gcse_spider.CrawlingInfo {
	if m != nil {
		return m.CrawlingInfo
	}
	return nil
}

func (m *PackageInfo) GetFolderInfo() *gcse_spider.FolderInfo {
	if m != nil {
		return m.FolderInfo
	}
	return nil
}

func (m *PackageInfo) GetRepoInfo() *gcse_spider.RepoInfo {
	if m != nil {
		return m.RepoInfo
	}
	return nil
}

type PersonInfo struct {
	CrawlingInfo *gcse_spider.CrawlingInfo `protobuf:"bytes,1,opt,name=crawling_info,json=crawlingInfo" json:"crawling_info,omitempty"`
}

func (m *PersonInfo) Reset()                    { *m = PersonInfo{} }
func (m *PersonInfo) String() string            { return proto.CompactTextString(m) }
func (*PersonInfo) ProtoMessage()               {}
func (*PersonInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PersonInfo) GetCrawlingInfo() *gcse_spider.CrawlingInfo {
	if m != nil {
		return m.CrawlingInfo
	}
	return nil
}

type Repository struct {
	Branch       string                    `protobuf:"bytes,6,opt,name=branch" json:"branch,omitempty"`
	Signature    string                    `protobuf:"bytes,7,opt,name=signature" json:"signature,omitempty"`
	Packages     []*gcse_spider.Package    `protobuf:"bytes,1,rep,name=packages" json:"packages,omitempty"`
	ReadmeFn     string                    `protobuf:"bytes,2,opt,name=ReadmeFn,json=readmeFn" json:"ReadmeFn,omitempty"`
	ReadmeData   string                    `protobuf:"bytes,3,opt,name=ReadmeData,json=readmeData" json:"ReadmeData,omitempty"`
	Stars        int32                     `protobuf:"varint,4,opt,name=stars" json:"stars,omitempty"`
	CrawlingInfo *gcse_spider.CrawlingInfo `protobuf:"bytes,5,opt,name=crawling_info,json=crawlingInfo" json:"crawling_info,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repository) GetPackages() []*gcse_spider.Package {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *Repository) GetCrawlingInfo() *gcse_spider.CrawlingInfo {
	if m != nil {
		return m.CrawlingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageInfo)(nil), "gcse.store.PackageInfo")
	proto.RegisterType((*PersonInfo)(nil), "gcse.store.PersonInfo")
	proto.RegisterType((*Repository)(nil), "gcse.store.Repository")
}

var fileDescriptor0 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0x55, 0x9a, 0x3f, 0x4d, 0x66, 0xf3, 0xfb, 0x21, 0xac, 0x02, 0xa6, 0xa0, 0x12, 0x72, 0xe2,
	0x94, 0xa0, 0x56, 0x48, 0x9c, 0x38, 0x00, 0xaa, 0x54, 0x89, 0x43, 0xb5, 0x12, 0x17, 0x2e, 0x91,
	0xe3, 0x9d, 0x6c, 0x0c, 0x89, 0xbd, 0xb2, 0x1d, 0xa0, 0x5f, 0x80, 0x6f, 0xc3, 0x77, 0xc4, 0x1e,
	0x3b, 0x69, 0xa2, 0x72, 0xa0, 0x97, 0xc4, 0xf3, 0xde, 0xbc, 0x99, 0xb1, 0xe7, 0x2d, 0x5c, 0xd4,
	0xca, 0x2f, 0x37, 0xf3, 0x89, 0x34, 0xeb, 0x69, 0x25, 0xbe, 0xab, 0xaa, 0x42, 0x5d, 0x4b, 0x3d,
	0xad, 0xa5, 0xc3, 0x69, 0x63, 0x8d, 0x37, 0x53, 0xe7, 0x8d, 0xc5, 0xf4, 0x3b, 0x21, 0x84, 0x41,
	0x64, 0x27, 0x84, 0x9c, 0xbe, 0xf9, 0x87, 0x02, 0x8d, 0xaa, 0xd0, 0xe6, 0xbf, 0x54, 0x62, 0xfc,
	0xbb, 0x03, 0xc5, 0xb5, 0x90, 0xdf, 0x44, 0x8d, 0x57, 0x7a, 0x61, 0x18, 0x83, 0x8e, 0x16, 0x6b,
	0xe4, 0xad, 0x51, 0xeb, 0xd5, 0xa0, 0xa4, 0x33, 0xe3, 0x70, 0xdc, 0xa4, 0x14, 0x7e, 0x44, 0xf0,
	0x36, 0x64, 0x8f, 0xa1, 0x27, 0x36, 0x7e, 0x69, 0x2c, 0x6f, 0x13, 0x91, 0x23, 0x76, 0x02, 0x5d,
	0xe7, 0x85, 0x75, 0xbc, 0x13, 0xe0, 0x6e, 0x99, 0x02, 0x76, 0x0a, 0x7d, 0x77, 0xa3, 0x4d, 0xe3,
	0x94, 0xe3, 0x5d, 0xca, 0xdf, 0xc5, 0x6c, 0x04, 0x45, 0x85, 0x4e, 0x5a, 0xd5, 0x78, 0x65, 0x34,
	0xef, 0x11, 0xbd, 0x0f, 0xb1, 0x17, 0x50, 0x84, 0x91, 0xbf, 0xa2, 0xf4, 0xb3, 0x8d, 0x5d, 0xf1,
	0x63, 0xca, 0x80, 0x0c, 0x7d, 0xb6, 0x2b, 0xf6, 0x0c, 0x06, 0x16, 0x45, 0xb5, 0xc6, 0xd9, 0x42,
	0xf3, 0x7e, 0xaa, 0x9f, 0x80, 0x4b, 0x52, 0x67, 0xb2, 0x12, 0x5e, 0xf0, 0x41, 0x52, 0x27, 0xe8,
	0x63, 0x40, 0xe2, 0x25, 0xd5, 0xba, 0x31, 0xd6, 0x3b, 0x0e, 0xa3, 0x76, 0xbc, 0x64, 0x0e, 0xd9,
	0x4b, 0x18, 0x7a, 0x74, 0x7e, 0xb6, 0xa5, 0x0b, 0xa2, 0x8b, 0x88, 0x5d, 0xe5, 0x94, 0x70, 0x33,
	0xfc, 0x19, 0x8f, 0x58, 0xf1, 0x21, 0xd1, 0xbb, 0x98, 0x9d, 0x41, 0x68, 0xb3, 0x40, 0x8b, 0x5a,
	0xa2, 0xe3, 0x8c, 0xd8, 0x3d, 0x84, 0xbd, 0x83, 0xff, 0xa4, 0x15, 0x3f, 0x56, 0x4a, 0xd7, 0x33,
	0x15, 0x56, 0xc0, 0x1f, 0x86, 0xd9, 0x8a, 0xf3, 0xa7, 0x93, 0xb4, 0xdc, 0xb4, 0xac, 0x0f, 0x39,
	0x23, 0xee, 0xa8, 0x1c, 0xca, 0xbd, 0x88, 0xbd, 0x85, 0x62, 0x61, 0x56, 0x21, 0x29, 0xa9, 0xff,
	0x27, 0xf5, 0x93, 0x03, 0xf5, 0x25, 0xf1, 0xa4, 0x85, 0xc5, 0xee, 0xcc, 0xce, 0xe3, 0x83, 0x35,
	0x26, 0xe9, 0x1e, 0x90, 0xee, 0xd1, 0x81, 0xae, 0x0c, 0x2c, 0xa9, 0xfa, 0x36, 0x9f, 0xc6, 0x9f,
	0x00, 0xae, 0xd1, 0x3a, 0xa3, 0xa9, 0xc2, 0x9d, 0xd9, 0x5b, 0xf7, 0x9a, 0x7d, 0xfc, 0xeb, 0x08,
	0x20, 0x36, 0x71, 0x2a, 0x78, 0xf8, 0x26, 0xda, 0x69, 0x6e, 0x85, 0x96, 0xcb, 0xbc, 0xff, 0x1c,
	0xb1, 0xe7, 0x30, 0x70, 0xaa, 0xd6, 0xc2, 0x6f, 0x2c, 0xe6, 0xc5, 0xdf, 0x02, 0xec, 0x35, 0xf4,
	0xb3, 0x1f, 0x5d, 0xe8, 0xdf, 0x0e, 0xfd, 0x4f, 0x0e, 0xfa, 0x67, 0x7b, 0x97, 0xbb, 0xac, 0xb8,
	0xae, 0x32, 0x1b, 0x23, 0x3b, 0xfa, 0xd6, 0x28, 0x67, 0x71, 0xa2, 0xad, 0x2b, 0xb2, 0xad, 0xf7,
	0x7d, 0xf2, 0x77, 0x6b, 0xdf, 0x79, 0x88, 0xee, 0xbd, 0x1e, 0xe2, 0x7d, 0xef, 0x4b, 0xc7, 0xf9,
	0x66, 0x3e, 0xef, 0xd1, 0x57, 0x79, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x3b, 0x51, 0x43,
	0x0f, 0x04, 0x00, 0x00,
}
