package mimeBase

type artact interface {
	ArticleWriteAttachment()
}

func ArticleAttachment() artact {
	ArticleStorageBase := "DB"
	switch ArticleStorageBase {
	case "fs":
		return &FS{}
	case "db":
		return &DB{}
	default:
		return &DB{}
	}
}
