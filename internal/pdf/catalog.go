package pdf

// TODO: implement

type Catalog struct {
	DictionaryElement
}

func (c *Catalog) Root() Object {
	return c.Dictionary().Key(KeyTreeRoot)
}

func (c *Catalog) MarkInfo() Object {
	return c.Dictionary().Key(KeyMarkInfo)
}

func (c *Catalog) Lang() Object {
	return c.Dictionary().Key(KeyLang)
}

func (c *Catalog) Metadata() Object {
	return c.Dictionary().Key(KeyMetadata)
}

func (c *Catalog) GetOrCreateMetadataObject() Object {
	dict := c.Dictionary()

	metadata := dict.Key(KeyMetadata)
	if metadata != nil {
		return metadata
	}

	// TODO: translate:
	// metadata = &GetDocument().GetObjects().CreateDictionaryObject("Metadata", "XML");

	// TODO? need to copy metadata?
	dict.AddKeyIndirect(KeyMetadata, metadata)

	panic("not implemented") // TODO: implement me
}

func (c *Catalog) MetadataStream() (stream []byte, err error) {
	metadata := c.Metadata()
	if metadata == nil {
		return nil, nil
	}

	// TODO: translate:

	//     auto stream = obj->GetStream();
	//     if (stream == nullptr)
	//         return ret;

	//     StringStreamDevice ouput(ret);
	//     stream->CopyTo(ouput);
	//     return ret;
	panic("not implemented") // TODO: implement me
}

func (c *Catalog) SetMetadataStream(value []byte) error {
	panic("not implemented") // TODO: implement me
}

func (c *Catalog) PageMode() PageMode {
	// TODO? what should Object->GetName() return?
	panic("not implemented") // TODO: implement me
}

func (c *Catalog) SetPageMode(mode PageMode) {
	panic("not implemented") // TODO: implement me
}

func (c *Catalog) SetPageLayout(layout PageLayout) {
	panic("not implemented") // TODO: implement me
}

// TODO: methods:
// - void SetUseFullScreen();
// - void SetHideToolbar();
// - void SetHideMenubar();
// - void SetHideWindowUI();
// - void SetFitWindow();
// - void SetCenterWindow();
// - void SetDisplayDocTitle();
// - void SetPrintScaling(const PdfName& scalingType);
// - void SetBaseURI(const std::string_view& baseURI);
// - void SetLanguage(const std::string_view& language);
// - void SetBindingDirection(const PdfName& direction);

// TODO: private methods:
// - void setViewerPreference(const PdfName& whichPref, const PdfObject& valueObj);
// - void setViewerPreference(const PdfName& whichPref, bool inValue);
