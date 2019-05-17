package elasticsearch

import (
	"context"
	"fmt"
	"github.com/zppro/go-common/data"
	helper "github.com/zppro/go-common/syntax2func"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
)

var client *elastic.Client

type QueryBody interface {
	ToJSON ()
}

type ESQueryResult struct {
	total int64
	rows []data.Row
}

var defaultQueryOption = data.QueryOption{RowParser: data.DefaultRowParser()}

func (r *ESQueryResult) Total () int64 {
	return r.total
}

func (r *ESQueryResult) Rows () []data.Row {
	return r.rows
}

func DefaultQueryOption () data.QueryOption {
	return defaultQueryOption
}

func Init (host, port string) (err error) {
	client, err = elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("%s:%s", host, port)),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	return // err判断留给调用方
}

func rawQuery(eIndex, eType, body string, parser data.QueryRowParser) (result *ESQueryResult, err error) {
	rawQuery := elastic.NewRawStringQuery(body)
	res, err := client.Search().Index(eIndex).Type(eType).Query(rawQuery).DoC(context.Background())
	if err != nil {
		return
	}
	result = &ESQueryResult{}
	fmt.Println("-------->", result)

	result.total = res.Hits.TotalHits
	result.rows = make([]data.Row, 0, len(res.Hits.Hits))

	if len(res.Hits.Hits) > 0 {
		fmt.Printf("Found a total of %d record \n", res.Hits.TotalHits)
		for _, hit := range res.Hits.Hits {
			if t, ok := parser.ParseRow(*hit.Source); ok {
				//fmt.Printf("activity %v\n", t)
				result.rows = append(result.rows, t)
			}
		}
	}
	return
}

func Query (chRet chan data.QueryResult, eIndex, eType, body string, option data.QueryOption) {

	result, err := rawQuery(eIndex, eType, body, helper.Iif(helper.IsZero(option), defaultQueryOption, option).(data.QueryOption).RowParser)

	log.Println("aaa:")
	if err != nil {
		log.Println(err)
		close(chRet)
		return
	}

	chRet <- result
}


