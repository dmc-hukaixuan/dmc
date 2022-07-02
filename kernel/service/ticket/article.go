package ticket

func BackendForChannel() {

}

// Returns an filtered array of base article data for a ticket.
// Returns a list with base article data (no back end related data included)
func ArticleList(ticketID int64) {

}

func TicketIDLookup() {}

// Set article flags.
func ArticleFlagSet() {}

func ArticleFlagDelete() {}

func ArticleFlagGet() {}

func ArticleFlagsOfTicketGet() {}

func ArticleAccountedTimeGet(){}

func ArticleAccountedTimeDelete() {}

// List all article sender types.
func ArticleSenderTypeList(){}

// Lookup an article sender type id or name.
func ArticleSenderTypeLookup() {}

// Set the article flags to indicate if the article search index needs to be rebuilt.
func ArticleSearchIndexRebuildFlagSet() {}

// 
func ArticleSearchIndexRebuildFlagList() {}

// gets an article indexing status hash.
func ArticleSearchIndexStatus(){}

// Rebuilds the current article search index table content. Existing article entries will be replaced.
func ArticleSearchIndexBuild() {}

func ArticleSearchIndexDelete() {}

// Checks the given search parameters for used article backend fields.
func ArticleSearchIndexSQLJoinNeeded(){}

// Generates SQL string extensions, including the needed table joins for the article index search.
func ArticleSearchIndexSQLJoin() {}

// Generates SQL query conditions for the used article fields, that may be used in the WHERE clauses of main
func ArticleSearchIndexWhereCondition() {}

// Find stop words within given search string.
func SearchStringStopWordsFind() {}

func ArticleCreate() {
}

func ArticleGet() {}

func ArticleUpdate() {}

func ArticleSend() {}

func ArticleBounce() {}

func SendAutoResponse() {}

func ArticleIndex() {}

func ArticleAttachmentIndex() {}

// Get article attachment from storage. This is a delegate method from active backend.
func ArticleAttachment() {}

// Write an article attachment to storage.
func ArticleWriteAttachment() {}

// Returns count of article.
func ArticleCount() {}

// Returns count of article attachment.
func ArticleAttachmentCount() {}

// Get the stored content path of an article.
func ArticleContentPathGet() {}
